FROM golang:alpine as build

WORKDIR /honk
COPY honk.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o honk .

FROM scratch

COPY --from=build /honk .
EXPOSE 8080
ENTRYPOINT ["./honk"]
