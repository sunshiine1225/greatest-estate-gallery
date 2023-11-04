# Use the official Golang image as a parent image
FROM golang:1.19

# Set the working directory to /app
WORKDIR /app

# Copy the source code and directories into the container
COPY src/ /app/
COPY templates/ /app/templates/
COPY static/ /app/static/

# Build the Go application
RUN go build -o greatest-estate-gallery /app/main.go

# Expose port 8000
EXPOSE 8000

# Run the application
CMD ["/app/greatest-estate-gallery"]
