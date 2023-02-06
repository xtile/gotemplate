FROM golang:alpine

RUN GOCACHE=OFF


WORKDIR /app
COPY . .

RUN go build -v -o app  ./cmd
EXPOSE 8000
CMD ["app"]
