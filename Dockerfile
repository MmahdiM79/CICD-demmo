FROM golang:1.27-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY internal/ ./internal/
COPY cmd/ ./cmd/
RUN CGO_ENABLED=0 go build -trimpath -o /server ./cmd/server

FROM alpine:3.23
WORKDIR /app
COPY --from=build /server ./server
ARG APP_VERSION=dev
ENV APP_VERSION=$APP_VERSION PORT=3000
EXPOSE 3000
USER 65532:65532
HEALTHCHECK CMD wget -qO- http://localhost:3000/health || exit 1
CMD ["./server"]
