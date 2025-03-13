# Basic Docker Demo for Go webserver

This project demonstrates a simple Docker workflow for a Go http application.

## Build and run the Docker Image
Run the following command to build/run the image:

```sh
docker build -t go-single-stage .
docker run -d -p 8080:8080 go-single-stage 
```

## What are the downsides? What did we do wrong?

This image uses golang alpine as base image which means it golang toolkit(compiler) and ton of dependencies with it. So, how to make it better?