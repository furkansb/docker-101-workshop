# Basic Docker Demo for Go webserver

This project demonstrates a simple Docker workflow for a Go http application using multi-stage build.

## Build and run the Docker Image
Run the following command to build/run the image:

```sh
docker build -t go-multistage .
docker run -d -p 8080:8080 go-multistage
```

## What are the upsides?

The final image size will be much smaller and with less dependencies meaning less attack surface for security as well.