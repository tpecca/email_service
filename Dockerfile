 # Use an official Golang runtime as a parent image
 FROM golang:1.22.6

 # Set the working directory inside the container
 WORKDIR /app

 # Copy the local package files to the container's workspace
 COPY . /app

# CMD ["go run ./examples/twoproviders.go"]
 CMD ["go", "run", "examples/twoproviders.go"]