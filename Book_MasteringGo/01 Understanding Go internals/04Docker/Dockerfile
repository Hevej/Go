FROM golang:alpine

RUN mkdir /files
COPY HW.go /files
WORKDIR /files
RUN go build -o /files/HW HW.go
ENTRYPOINT ["/files/HW"]