# Use the official Golang image as the base image
FROM golang:1.20.6-alpine3.17

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the Golang code into the container
COPY . .

# Build the Golang binary
RUN go build -o main .

# Set the entry point to the Golang binary

EXPOSE 50051
CMD ["./main"]
