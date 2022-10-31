# periodic-task

### Start Project

* go build periodic-task
* go run periodic-task

The App by default will run at port 8080, if you want you to change the http server port 
You have to set enviromental variable HTTP_PORT to whatever port you like

## Docker 

docker build --tag periodic-task .

docker run -p 8080:8080 periodic-task:latest

### if you want to change the port for the docker image:

docker run -e HTTP_PORT='3000' -p 3000:3000 periodic-task:latest