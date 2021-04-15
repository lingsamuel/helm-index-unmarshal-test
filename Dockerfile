#syntax=docker/dockerfile:1.2

FROM golang:latest as golang
WORKDIR /go/src/app
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

# install deps first
COPY go.* ./
RUN go mod download
# Use .dockerignore to make sure unrelated changes won't invalidates cache
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build make build

# ---------------
# Generate Images
# ---------------

FROM scratch
COPY ./gen-index.yaml /helm-index.yaml
COPY --from=golang /go/src/app/output/main /
ENTRYPOINT ["/main"]
