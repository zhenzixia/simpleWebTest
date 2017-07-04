package resource

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"simpleWebTest/pkg/utils"
	"strconv"
	"github.com/cactus/go-statsd-client/statsd"
	"log"
	"time"
)

type Record struct {
	Id	string	`json:"id"`
	Count 	int64 	`json:"count"`
	Geo	Geo	`json:"geo"`
}

type Geo struct {
	CityName	string	`json:"city_name"`
	ContinentCode	string	`json:"continent_code"`
	CountryIsoCode	string	`json:"country_iso_code"`
}

func NewItem() *Record {
	return &Record{}
}

type RecordResource struct {
	client statsd.Statter
}

func (s *RecordResource) Initialize(statsdHost, prefix string) {
	/*
	client, err := statsd.NewClient(statsdHost, prefix)
	if err != nil {
		log.Fatal(err)
	}
	s.client = client.(statsd.Statter)
	*/

	//Use buffered client in improve performance
	bufferdClient, err := statsd.NewBufferedClient(statsdHost, prefix, 1*time.Second, 1024)
	if err != nil {
		log.Fatal(err)
	}
	s.client = bufferdClient.(statsd.Statter)

}

func (s *RecordResource) Register() {
	service := new(restful.WebService)
	service.Path("/record")
	service.Consumes(restful.MIME_JSON)
	service.Produces(restful.MIME_JSON)

	service.Route(service.GET("").To(s.GetOne))

	service.Route(service.POST("").To(s.PostOne)).
		Param(service.QueryParameter("count", "Description...").DataType("int"))

	service.Route(service.PUT("").To(s.CreateOne))

	service.Route(service.DELETE("/{id}").To(s.RemoveOne))

	restful.Add(service)
}

func (s *RecordResource) PostOne(request *restful.Request, response *restful.Response) {
	count, _ := strconv.Atoi(request.QueryParameter("count"))
	record := Record{
		Id : utils.GenerateUUID(),
		Count: int64(count),
	}

	err := request.ReadEntity(&record)

	log.Printf(">> Posting one record! Record: %+v ", record) 

	if err == nil {
		err = s.client.Inc(record.Geo.CityName, record.Count, 1.0)
		if err != nil {
			log.Printf(">> Error loading data to StatsD! Errer: %v", err.Error())
		}
		log.Printf(">> Success posting one record! Record: %v", record)
		response.WriteEntity(record)
	} else {
		response.WriteError(http.StatusInternalServerError,err)
	}
}


func (s *RecordResource) GetAll(request *restful.Request, response *restful.Response) {
}

func (s *RecordResource) GetOne(request *restful.Request, response *restful.Response) {
}

func (s *RecordResource) CreateOne(request *restful.Request, response *restful.Response) {
}

func (s *RecordResource) RemoveOne(request *restful.Request, response *restful.Response) {
}
