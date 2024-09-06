FROM golang:1.23.1-bookworm AS build-stage
    WORKDIR /app

    COPY go.mod go.sum ./

    RUN go mod download

    COPY . .

    RUN ls -al ./cmd/FlowManagerAPI/
    
    RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/FlowManagerAPI/main.go

    RUN ls -al

FROM scratch AS build-release-stage
    WORKDIR /
  
    COPY --from=build-stage /app/api ./api
  
    EXPOSE 8080
  
    ENTRYPOINT ["/api"]