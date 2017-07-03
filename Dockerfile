#
# Super simple example of a Dockerfile
#
FROM ubuntu:latest
MAINTAINER zixia zhen "jerryzhen01@gmail.com"

#RUN apt-get update
#RUN apt-get install -y python python-pip wget
#RUN pip install Flask

ADD simpleWebTest /usr/local/bin

EXPOSE 8080

WORKDIR /home
