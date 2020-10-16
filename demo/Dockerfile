FROM golang:1.13.12 AS bulider
WORKDIR /build
COPY . .
ENV GOPROXY=https://goproxy.io CGO_ENABLED=0
RUN go mod download
RUN go build -o run-app .

FROM scratch
WORKDIR /app

COPY --from=bulider /build/run-app /app
EXPOSE 5000
CMD ["/app/run-app"]
