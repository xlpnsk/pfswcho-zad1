# syntax=docker/dockerfile:1

# golang alpine as a builder image
FROM golang:1.17-alpine AS builder

# setup the workdir
WORKDIR /app

# download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# copy HTML templates and go file
COPY templates ./templates 
COPY *.go ./

# extldflags and tags for creating statically linked executable
RUN go build -ldflags '-w -extldflags "-static"' -tags netgo,osusergo -o /app/go-app

# scrach as final image base
FROM scratch
WORKDIR /app
# copy templates and compiled executable
COPY --from=builder /app/go-app .
COPY --from=builder /app/templates ./templates 
# expose port
EXPOSE 8080

# define container entrypoint
ENTRYPOINT [ "./go-app" ]