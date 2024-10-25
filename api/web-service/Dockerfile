FROM --platform=linux/amd64 golang:1.18-alpine
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY main.go .
RUN go mod download
COPY . .
RUN go build -o ./out/dist .
EXPOSE 80
CMD ./out/dist
