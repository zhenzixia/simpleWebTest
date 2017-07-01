package resource

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"simpleWebTest/pkg/utils"
	"strconv"
	"github.com/cactus/go-statsd-client/statsd"
	"log"
)


type Product struct {
	Id	string	`json:"id"`
	Name 	string	`json:"name"`
	Count 	int64 	`json:"count"`
	Geo	Geo	`json:"geo"`
}

type Geo struct {
	CityName	string	`json:"city_name"`
	ContinentCode	string	`json:"continent_code"`
	CountryIsoCode	string	`json:"country_iso_code"`
}

func NewItem() *Product {
	return &Product{}
}


type ProductResource struct {
	client *statsd.Client
}

func (s *ProductResource) Initialize(prefix, statsdHost string) {
	//statsdclient := statsd.NewStatsdClient(statsdHost, prefix)
	//client, err := statsd.NewClient("127.0.0.1:8125", "test-client")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//s.client = client.(statsd.Client)
}

func (s *ProductResource) Register() {
	service := new(restful.WebService)
	service.Path("/product")
	service.Consumes(restful.MIME_JSON)
	service.Produces(restful.MIME_JSON)

	service.Route(service.GET("name").To(s.GetOne))

	service.Route(service.POST("").To(s.PostOne)).
		Param(service.QueryParameter("count", "count count count").DataType("int"))

	service.Route(service.PUT("").To(s.CreateOne))

	service.Route(service.DELETE("/{id}").To(s.RemoveOne))

	restful.Add(service)
}

func (s *ProductResource) PostOne(request *restful.Request, response *restful.Response) {
	count, _ := strconv.Atoi(request.QueryParameter("count"))
	product := Product{
		Id : utils.GenerateUUID(),
		Count: int64(count),
	}

	log.Printf(">> Posting one record! Product ID: %v", product.Id)
	//create a new client
	client, err := statsd.NewClient("127.0.0.1:8125", "test-client")
	if err != nil {
		log.Fatal(err)
	}

	err = request.ReadEntity(&product)
	// here you would update the user with some persistence system
	if err == nil {
		//response.WriteEntity(product)
		client.Inc(product.Name, product.Count, 1.0)
		response.WriteEntity(product)
	} else {
		response.WriteError(http.StatusInternalServerError,err)
	}
}


func (s *ProductResource) GetAll(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")
	// here you would fetch user from some persistence system
	//usr := &User{Id: id, Name: "John Doe"}
	item := Product{Id:id}
	response.WriteEntity(item)
}

func (s *ProductResource) GetOne(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")
	//queryId := request.QueryParameter("id")
	// here you would fetch user from some persistence system
	item := Product{Name: name}
	response.WriteEntity(item)
}

func (s *ProductResource) CreateOne(request *restful.Request, response *restful.Response) {

	product := Product{}
	err := request.ReadEntity(&product)
	// here you would create the user with some persistence system
	if err == nil {
		response.WriteEntity(product)
	} else {
		response.WriteError(http.StatusInternalServerError,err)
	}
}

func (s *ProductResource) RemoveOne(request *restful.Request, response *restful.Response) {
	// here you would delete the user from some persistence system
}
