# start from a minimal base image
FROM golang:1.19.4
# add the source code to the container
ADD go-app/src /app
# set the working directory to /app
WORKDIR /app
# compile the application and call it app
RUN go build apiserver.go
RUN mv apiserver app
EXPOSE 8080
# start the application when the container starts
CMD ["./app"]
