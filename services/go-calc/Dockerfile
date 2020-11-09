FROM golang:alpine as builder

RUN mkdir /build 
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o go-calc .

FROM scratch
COPY --from=builder /build/go-calc /app/
WORKDIR /app
CMD ["./go-calc"]
