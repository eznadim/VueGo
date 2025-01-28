# Build frontend
FROM node:18-alpine as frontend
WORKDIR /app
COPY frontend/package*.json ./
RUN npm install
COPY frontend .
RUN npm run build

# Build backend
FROM golang:1.20-alpine as backend
WORKDIR /app
COPY backend/go.mod .
COPY backend/go.sum .
RUN go mod download
COPY backend .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final image
FROM alpine:latest
WORKDIR /app
COPY --from=frontend /app/dist ./frontend/dist
COPY --from=backend /app/main .
EXPOSE 8080
CMD ["./main"]