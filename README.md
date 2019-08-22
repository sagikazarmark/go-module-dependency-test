# Go Module Dependency Test V2

Testing dependency resolution with Go Modules.

```bash
go version
go version go1.13rc1 darwin/amd64
```

Compared to the previous set of tests

- Go 1.13 RC1 is used here
- The new [logur library](https://github.com/logur) is used with separate modules for adapters.

(For adapter packages the amount of downloaded code is significantly less)

**Note:** Go 1.13 RC1 with the default proxy gives roughly the same results as Go 1.12 with a custom Athens proxy.


## Summary

| Test Case       | `go mod tidy` time | `go mod tidy` size | `go mod download` time | `go mod download` size |
|-----------------|--------------------|--------------------|------------------------|------------------------|
| Default proxy   | 1.053s             | 12M                | 1.250s                 | 13M                    |
| No proxy        | 5.896s             | 29M                | 5.918s                 | 30M                    |



## Test 1 - default proxy

```
bash
export GOPATH=/tmp/gopath-for-dependency-size-test-1
time go mod tidy
du -sh /tmp/gopath-for-dependency-size-test-1
sudo rm -rf /tmp/gopath-for-dependency-size-test-1/*
time go mod download
du -sh /tmp/gopath-for-dependency-size-test-1
```

Output:

```
# time go mod tidy
go: downloading github.com/sirupsen/logrus v1.4.2
go: downloading logur.dev/adapter/logrus v0.2.0
go: extracting github.com/sirupsen/logrus v1.4.2
go: downloading golang.org/x/sys v0.0.0-20190422165155-953cdadca894
go: downloading github.com/konsorten/go-windows-terminal-sequences v1.0.1
go: extracting logur.dev/adapter/logrus v0.2.0
go: downloading github.com/stretchr/testify v1.2.2
go: extracting github.com/konsorten/go-windows-terminal-sequences v1.0.1
go: downloading logur.dev/logur v0.15.0
go: extracting logur.dev/logur v0.15.0
go: extracting github.com/stretchr/testify v1.2.2
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading github.com/pmezard/go-difflib v1.0.0
go: extracting golang.org/x/sys v0.0.0-20190422165155-953cdadca894
go: extracting github.com/pmezard/go-difflib v1.0.0
go: extracting github.com/davecgh/go-spew v1.1.1

real	0m1.053s
user	0m0.464s
sys	0m0.399s


# du -sh /tmp/gopath-for-dependency-size-test-1
12M	/tmp/gopath-for-dependency-size-test-1

# time go mod download
go: finding github.com/davecgh/go-spew v1.1.1
go: finding github.com/konsorten/go-windows-terminal-sequences v1.0.1
go: finding github.com/pmezard/go-difflib v1.0.0
go: finding github.com/sirupsen/logrus v1.4.2
go: finding github.com/stretchr/objx v0.1.1
go: finding github.com/stretchr/testify v1.2.2
go: finding golang.org/x/sys v0.0.0-20190422165155-953cdadca894
go: finding logur.dev/adapter/logrus v0.2.0
go: finding logur.dev/logur v0.15.0

real	0m1.250s
user	0m0.338s
sys	0m0.245s


# du -sh /tmp/gopath-for-dependency-size-test-1
13M	/tmp/gopath-for-dependency-size-test-1
```


## Test 2 - no proxy

```
bash
export GOPATH=/tmp/gopath-for-dependency-size-test-2
export GOPROXY=direct
time go mod tidy
du -sh /tmp/gopath-for-dependency-size-test-2
sudo rm -rf /tmp/gopath-for-dependency-size-test-2/*
time go mod download
du -sh /tmp/gopath-for-dependency-size-test-2
```

Output:

```
# time go mod tidy
go: downloading github.com/sirupsen/logrus v1.4.2
go: downloading logur.dev/adapter/logrus v0.2.0
go: extracting logur.dev/adapter/logrus v0.2.0
go: downloading logur.dev/logur v0.15.0
go: extracting github.com/sirupsen/logrus v1.4.2
go: downloading github.com/stretchr/testify v1.2.2
go: downloading github.com/konsorten/go-windows-terminal-sequences v1.0.1
go: downloading golang.org/x/sys v0.0.0-20190422165155-953cdadca894
go: extracting logur.dev/logur v0.15.0
go: extracting github.com/konsorten/go-windows-terminal-sequences v1.0.1
go: extracting github.com/stretchr/testify v1.2.2
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading github.com/pmezard/go-difflib v1.0.0
go: extracting github.com/pmezard/go-difflib v1.0.0
go: extracting github.com/davecgh/go-spew v1.1.1
go: extracting golang.org/x/sys v0.0.0-20190422165155-953cdadca894

real	0m5.896s
user	0m2.728s
sys	0m1.409s


# du -sh /tmp/gopath-for-dependency-size-test-1
29M	/tmp/gopath-for-dependency-size-test-2

# time go mod download
go: finding github.com/davecgh/go-spew v1.1.1
go: finding github.com/konsorten/go-windows-terminal-sequences v1.0.1
go: finding github.com/pmezard/go-difflib v1.0.0
go: finding github.com/sirupsen/logrus v1.4.2
go: finding github.com/stretchr/objx v0.1.1
go: finding github.com/stretchr/testify v1.2.2
go: finding golang.org/x/sys v0.0.0-20190422165155-953cdadca894
go: finding logur.dev/adapter/logrus v0.2.0
go: finding logur.dev/logur v0.15.0

real	0m5.918s
user	0m2.663s
sys	0m1.457s


# du -sh /tmp/gopath-for-dependency-size-test-1
30M	/tmp/gopath-for-dependency-size-test-2
```
