FROM golang:1.23.2

# Set the working directory inside the container
WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o coccoc main.go

# Command to run the executable
CMD ["./coccoc"]
