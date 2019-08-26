FROM golang:latest
ENV GO111MODULE on
COPY . /app
WORKDIR /app
RUN go build -o /app/main .
CMD ["/app/main"]