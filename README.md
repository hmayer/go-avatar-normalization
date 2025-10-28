# Go Avatar Normalization

A Go-based HTTP service for avatar image processing and normalization.
This service provides an endpoint to upload and process avatar images.

## Features

- Avatar image upload and processing via HTTP endpoint
- OpenCV integration for advanced image processing
- Docker support for easy deployment
- Lightweight Alpine-based Docker image

## Prerequisites

### Running Locally
- Go 1.25 or higher
- OpenCV development libraries
- HDF5 libraries
- VTK libraries
- GCC/G++ compiler

### Running with Docker
- Docker and Docker Compose installed on your system

## Installation

### Local Installation

1. Clone the repository:
```bash
git clone https://github.com/hmayer/go-avatar-normalization.git
cd go-avatar-normalization
```

2. Install system dependencies (Alpine/apk):
```bash
apk add --no-cache opencv-dev hdf5-dev vtk-dev g++
```
for Ubuntu/Debian (apt-get):
```bash
apt-get install libopencv-dev libhdf5-dev libvtk9-dev build-essential
```
or for Arch-based (pacman):
```bash
sudo pacman -S opencv hdf5 vtk
```

3. Download Go dependencies:
```bash
go mod download
```

4. Build the application:
```bash
go build -o go-avatar-normalization
```

### Docker Installation

Build the Docker image:
```bash
docker compose build --no-cache
```
## Usage

### Running Locally

1. Start the server:
```bash
./go-avatar-normalization
```
The server will start on port 8000.

### Running with Docker

Run the container:
```bash
docker compose up
```

## API Endpoints

### Upload Avatar

**Endpoint:** `POST /avatar`

**Description:** Upload and process an avatar image.

**Example:**
```bash
curl -X POST -F "avatar=@/path/to/image.jpg" http://localhost:8000/avatar
```
## Project Structure
```
.
├── main.go                 # Main application entry point
├── handlers/               # HTTP request handlers
├── resources/images/       # Cached image storage
├── Dockerfile             # Docker configuration
├── go.mod                 # Go module dependencies
└── go.sum                 # Go module checksums
```
## Automatic Cleanup

The service automatically manages cached files:
- On startup, removes all files older than 1 hour
- Runs cleanup every 10 minutes in the background
- Ensures disk space is not consumed by old cached images

## Docker Multi-Stage Build

The Dockerfile uses a multi-stage build to:
1. **Builder stage:** Compiles the application with all build dependencies
2. **Final stage:** Creates a minimal runtime image with only necessary libraries

This approach significantly reduces the final image size.

## Troubleshooting

### Port Already in Use
If port 8000 is already in use, either:
- Stop the service using that port
- Change the port in the application
- Use Docker port mapping to map to a different host port: `-p 8080:8000`

### OpenCV Dependencies
If you encounter OpenCV-related errors, ensure all required libraries are installed:
```bash
# Alpine
apk add opencv hdf5 vtk libopencv_aruco libopencv_photo libopencv_video libstdc++ libgcc

# Ubuntu/Debian
apt-get install libopencv-core-dev libopencv-imgproc-dev
```
## Author

Henrique Mayer
