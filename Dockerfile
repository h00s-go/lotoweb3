# Use the official Go image as the base image
FROM golang:alpine AS backend

WORKDIR /app

COPY backend ./

RUN go mod download && \
    go build -o /out/lotoweb

FROM oven/bun:latest AS frontend

WORKDIR /app

COPY frontend ./

RUN bun run build

FROM alpine:latest

WORKDIR /app

# Copy the pre-built binary file from the previous stage
COPY --from=backend /out/lotoweb ./
COPY --from=frontend /app/build ./public

EXPOSE 3000

CMD ["./lotoweb"]