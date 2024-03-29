# imagecompare

imagecompare accepts a CSV file with two fields (image1 and image2). It parses CSV, opens each pair of images and compares them. The result of the comparison is stored in a new file along with time taken to perform analysis


### Installing

Go to [antonsova.ca](https://antonsova.ca) and download a version for your OS.

Then run

```
$ sudo mkdir -p /usr/local/bin
$ sudo mv ~/Downloads/compare /usr/local/bin
$ sudo chmod 0755 /usr/local/bin/compare
```

Verify that it is successfully installed

```
$ compare -h
```

## Project

The project consists of two parts:
1. Cli that executes command
2. Web server that serves static html file and executables

To get the project clone repo inside `$GOPATH/src/github.com`

```
$ git clone  https://github.com/AntonMSova/imagecompare.git
```

## Building project

```
$ go build -o compare cmd/cli/main.go
$ go build -o server cmd/server/main.go
```

## Running

Running cli

```
$ ./compare
```

Running server

```

$ ./server
```

Then navigate in your browser to [localhost:8080](localhost:8080)


## Running tests

Tests are run with `ginkgo` and `gomega`

1. Install ginkgo and gomega

```
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega/...
```

The from project's root folder run

```
$ ginkgo -r
```

## Deployment

Deployment is done through `Travis` CI/CD pipeline. Merging to master branch will automatically deploy it to GKE

## Authors

* **Anton Sova**