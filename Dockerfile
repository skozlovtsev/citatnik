FROM golang:1.24-alpine

RUN mkdir -p citatnik

WORKDIR /citatnik

COPY . .
RUN go mod download

RUN go build ./cmd/main/citatnik.go

EXPOSE 8080

CMD [ "./citatnik" ]