FROM golang:1.25-alpine AS builder
LABEL authors="Henrique Mayer"

## Preparing to build
RUN mkdir -p /opt/avatar
WORKDIR /opt/avatar

## Downloading and installing build dependencies
RUN apk add --no-cache opencv-dev hdf5-dev vtk-dev g++
COPY go.mod go.sum ./
RUN go mod download

## Copying the rest of the project
COPY . .

## Building
RUN GOOS=linux go build -o /opt/avatar/go-avatar-normalization

## Creating final image
FROM alpine:latest

## Installing runtime dependencies
RUN apk add --no-cache opencv hdf5 vtk libopencv_aruco libopencv_photo libopencv_video libstdc++ libgcc

## Copying built binaries to final image
WORKDIR /app
COPY --from=builder /opt/avatar .

## Exposing por 8000
EXPOSE 8000

## Setting initial command
CMD ["./go-avatar-normalization"]

## We are done