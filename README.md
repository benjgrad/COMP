# COMP

### building the DockerFile

You can build the DockerFile found in the current directory `.` and tag it `comp`

```bash
docker build -t comp .
```

### Running the image

Now that the image is built, we can run it. Some notable flags are '-d' as detached, and '-p 1234:1234' for port mapping (eg. in web applications).

```bash
docker run --name Cmg.Tools -d comp
```

### Running code locally without docker

When in dev, it's useful to run the code using

```bash
go run src/comp/comp.go <linked list | other commands>
```

### Testing

```bash
go test -v src/comp/datastructures/linkedlist/linkedlist_test.go
```

For further information, refer to https://pkg.go.dev/testing
