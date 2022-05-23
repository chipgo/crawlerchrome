##
## Build
##
FROM golang:1.18-buster as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o main

##
## Deploy
##
FROM golang:1.18

# Install any required dependencies.
RUN apt -y install ca-certificates
RUN apt update && apt -y upgrade && apt -y install chromium

WORKDIR /

COPY --from=build /app/main /main
COPY /resources /resources

CMD ["/main"]