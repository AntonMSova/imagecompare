# imagecompare

imagecompare accepts a CSV file with two fields (image1 and image2). It parses CSV, opens each pair of images and compares them. The result of the comparison is stored in a new file along with time taken to perform analysis


### Installing

Go to [antonsova.ca](https://antonsova.ca) and download a version for your OS

Run

```
$ sudo mkdir -p /usr/local/bin
```

then

```
$ sudo mv ~/Downloads/compare /usr/local/bin
```

set up permissions

```
$ sudo chmod 0755 /usr/local/bin/compare
```

Finally run

```
$ compare -h
```

to ensure that everything is set up properly

## Project

The project consists of two parts:
1. Cli that executes command
2. Web server that serves static html file and executables

To get the project clone repo inside `$GOPATH/src/github.com`

```
$ git clone  https://github.com/AntonMSova/imagecompare.git
```

## Running the tests

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