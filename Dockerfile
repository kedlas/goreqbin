# Build the binary
FROM golang:latest
WORKDIR /goreqbin/cmd/
COPY . /goreqbin
RUN XDG_CACHE_HOME=/tmp/.cache CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./../build/goreqbin .

# Create runnable image with the binary only
FROM scratch
USER 65534:65534
WORKDIR /usr/local/bin
COPY --from=0 /goreqbin/build/goreqbin .
CMD ["./goreqbin"]
