FROM golang:1.21.0

WORKDIR /app
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/app -buildvcs=false

EXPOSE 3000

# Run
CMD ["/build/app"]