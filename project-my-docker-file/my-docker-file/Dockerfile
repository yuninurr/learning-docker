FROM golang:1.22.3

# set destination for copy 
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux go build -o /my-docker-file

EXPOSE 78

CMD ["/my-docker-file"]