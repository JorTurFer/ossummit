FROM golang:alpine AS build
WORKDIR /go/src/ossumit
COPY . .
RUN go build -o /go/bin/ossumit main.go

FROM ubuntu
COPY --from=build /go/bin/ossumit /go/bin/ossumit
COPY --from=build /go/src/ossumit/assets /assets
ENTRYPOINT ["/go/bin/ossumit"]