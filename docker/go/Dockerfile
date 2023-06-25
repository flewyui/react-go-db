FROM golang:alpine
WORKDIR /app/go
COPY ./docker/go/go-app /app/go
RUN pwd
RUN ls
RUN go build main.go
RUN touch go.html
RUN ls
CMD ./main
