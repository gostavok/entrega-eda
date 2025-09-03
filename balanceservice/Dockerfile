FROM golang:1.20
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o balanceservice-app ./cmd/balanceservice/main.go
CMD ["/app/balanceservice-app"]
EXPOSE 3003
