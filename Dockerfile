FROM golang:alpine

RUN GOCACHE=OFF


WORKDIR /app
COPY . .

RUN go build -v ./...
EXPOSE 8080
CMD ["gotest"]