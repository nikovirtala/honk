FROM golang:alpine as build

WORKDIR /honk
COPY . .
RUN go mod init honk && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o honk .

FROM scratch

COPY --from=build /honk/honk /
EXPOSE 8080
ENTRYPOINT ["/honk"]