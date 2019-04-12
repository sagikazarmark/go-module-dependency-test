# Go Module Dependency Test

Testing dependency resolution with Go Modules.


## Summary

| Test Case       | `go mod tidy` time | `go mod tidy` size | `go mod download` time | `go mod download` size |
|-----------------|--------------------|--------------------|------------------------|------------------------|
| No proxy        | 21.253s            | 108M               | 22.299s                | 218M                   |
| Gocenter.io     | 8.215s             | 17M                | 11.990s                | 127M                   |
| My Athens Proxy | 5.851s             | 17M                | 6.811s                 | 126M                   |



## Test 1 - no proxy

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
go: finding github.com/goph/logur v0.11.0
go: finding github.com/InVisionApp/go-logger v1.0.1
go: finding github.com/go-kit/kit v0.8.0
go: finding github.com/gofrs/uuid v3.1.0+incompatible
go: finding github.com/rs/zerolog v1.11.0
go: finding github.com/hashicorp/go-hclog v0.0.0-20181001195459-61d530d6c27f
go: finding github.com/bugsnag/panicwrap v1.2.0
go: finding github.com/bugsnag/bugsnag-go v1.4.0
go: finding github.com/go-logfmt/logfmt v0.4.0
go: finding github.com/go-logr/logr v0.1.0
go: finding go.uber.org/zap v1.9.1
go: finding go.uber.org/atomic v1.3.2
go: finding github.com/pkg/errors v0.8.0
go: finding github.com/rollbar/rollbar-go v1.0.2
go: finding go.uber.org/multierr v1.1.0
go: finding github.com/go-sql-driver/mysql v1.4.1
go: finding github.com/ThreeDotsLabs/watermill v0.1.2
go: finding github.com/kardianos/osext v0.0.0-20170510131534-ae77be60afb1
go: finding github.com/kr/logfmt v0.0.0-20140226030751-b84e30acd515
go: finding google.golang.org/grpc v1.17.0
go: finding github.com/satori/go.uuid v1.2.0
go: finding github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a
go: finding github.com/go-chi/chi v3.3.3+incompatible
go: finding github.com/kr/pretty v0.1.0
go: finding github.com/confluentinc/confluent-kafka-go v0.11.6
go: finding gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
go: finding github.com/client9/misspell v0.3.4
go: finding github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
go: finding golang.org/x/net v0.0.0-20180826012351-8a410e7b638d
go: finding golang.org/x/tools v0.0.0-20180828015842-6cd1fcedba52
go: finding github.com/golang/protobuf v1.2.0
go: finding google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8
go: finding honnef.co/go/tools v0.0.0-20180728063816-88497007e858
go: finding golang.org/x/lint v0.0.0-20181026193005-c67002cb31c3
go: finding golang.org/x/text v0.3.0
go: finding cloud.google.com/go v0.26.0
go: finding google.golang.org/appengine v1.1.0
go: finding golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f
go: finding github.com/kisielk/gotool v1.0.0
go: finding github.com/golang/mock v1.1.1
go: finding golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
go: finding golang.org/x/sys v0.0.0-20180830151530-49385e6e1522
go: finding github.com/kr/text v0.1.0
go: finding github.com/kr/pty v1.1.1
go: downloading github.com/sirupsen/logrus v1.2.0
go: downloading github.com/goph/logur v0.11.0
go: extracting github.com/goph/logur v0.11.0
go: extracting github.com/sirupsen/logrus v1.2.0
go: downloading github.com/stretchr/testify v1.2.2
go: downloading golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
go: downloading github.com/konsorten/go-windows-terminal-sequences v1.0.1
go: extracting golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
go: extracting github.com/konsorten/go-windows-terminal-sequences v1.0.1
go: downloading golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33
go: extracting golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33
go: extracting github.com/stretchr/testify v1.2.2
go: downloading github.com/pmezard/go-difflib v1.0.0
go: downloading github.com/davecgh/go-spew v1.1.1
go: extracting github.com/pmezard/go-difflib v1.0.0
go: extracting github.com/davecgh/go-spew v1.1.1

real	0m21.253s
user	0m10.859s
sys 	0m6.873s


# du -sh /tmp/gopath-for-dependency-size-test-1
108M	/tmp/gopath-for-dependency-size-test-1

# time go mod download
go: finding github.com/goph/logur v0.11.0
go: finding github.com/sirupsen/logrus v1.2.0
go: finding github.com/rs/zerolog v1.11.0
go: finding github.com/go-logfmt/logfmt v0.4.0
go: finding github.com/bugsnag/bugsnag-go v1.4.0
go: finding github.com/hashicorp/go-hclog v0.0.0-20181001195459-61d530d6c27f
go: finding github.com/rollbar/rollbar-go v1.0.2
go: finding github.com/InVisionApp/go-logger v1.0.1
go: finding github.com/go-kit/kit v0.8.0
go: finding github.com/pkg/errors v0.8.0
go: finding google.golang.org/grpc v1.17.0
go: finding go.uber.org/multierr v1.1.0
go: finding github.com/gofrs/uuid v3.1.0+incompatible
go: finding golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
go: finding github.com/go-sql-driver/mysql v1.4.1
go: finding github.com/stretchr/testify v1.2.2
go: finding go.uber.org/zap v1.9.1
go: finding github.com/stretchr/objx v0.1.1
go: finding github.com/go-logr/logr v0.1.0
go: finding github.com/davecgh/go-spew v1.1.1
go: finding golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33
go: finding github.com/bugsnag/panicwrap v1.2.0
go: finding github.com/golang/protobuf v1.2.0
go: finding golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f
go: finding github.com/client9/misspell v0.3.4
go: finding golang.org/x/sys v0.0.0-20180830151530-49385e6e1522
go: finding github.com/ThreeDotsLabs/watermill v0.1.2
go: finding golang.org/x/text v0.3.0
go: finding google.golang.org/appengine v1.1.0
go: finding github.com/konsorten/go-windows-terminal-sequences v1.0.1
go: finding golang.org/x/tools v0.0.0-20180828015842-6cd1fcedba52
go: finding github.com/kisielk/gotool v1.0.0
go: finding github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
go: finding honnef.co/go/tools v0.0.0-20180728063816-88497007e858
go: finding github.com/golang/mock v1.1.1
go: finding google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8
go: finding golang.org/x/lint v0.0.0-20181026193005-c67002cb31c3
go: finding github.com/kardianos/osext v0.0.0-20170510131534-ae77be60afb1
go: finding go.uber.org/atomic v1.3.2
go: finding github.com/pmezard/go-difflib v1.0.0
go: finding cloud.google.com/go v0.26.0
go: finding golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
go: finding github.com/confluentinc/confluent-kafka-go v0.11.6
go: finding github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a
go: finding github.com/satori/go.uuid v1.2.0
go: finding github.com/kr/pretty v0.1.0
go: finding github.com/kr/logfmt v0.0.0-20140226030751-b84e30acd515
go: finding golang.org/x/net v0.0.0-20180826012351-8a410e7b638d
go: finding github.com/go-chi/chi v3.3.3+incompatible
go: finding gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
go: finding github.com/kr/text v0.1.0
go: finding github.com/kr/pty v1.1.1

real	0m22.299s
user	0m18.516s
sys 	0m9.538s


# du -sh /tmp/gopath-for-dependency-size-test-1
218M	/tmp/gopath-for-dependency-size-test-1
```


## Test 2 - gocenter.io

```
bash
export GOPATH=/tmp/gopath-for-dependency-size-test-2
export GOPROXY='https://gocenter.io'
time go mod tidy
du -sh /tmp/gopath-for-dependency-size-test-2
sudo rm -rf /tmp/gopath-for-dependency-size-test-2/*
time go mod download
du -sh /tmp/gopath-for-dependency-size-test-2
```

Output:

```
# time go mod tidy
go: finding github.com/goph/logur v0.11.0
go: finding github.com/sirupsen/logrus v1.2.0
go: finding github.com/bugsnag/bugsnag-go v1.4.0
go: finding github.com/rollbar/rollbar-go v1.0.2
go: finding github.com/go-logfmt/logfmt v0.4.0
go: finding github.com/go-kit/kit v0.8.0
go: finding go.uber.org/multierr v1.1.0
go: finding github.com/ThreeDotsLabs/watermill v0.1.2
go: finding google.golang.org/grpc v1.17.0
go: finding github.com/InVisionApp/go-logger v1.0.1
go: finding github.com/hashicorp/go-hclog v0.0.0-20181001195459-61d530d6c27f
go: finding github.com/pkg/errors v0.8.0
go: finding golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
go: finding github.com/gofrs/uuid v3.1.0+incompatible
go: finding go.uber.org/zap v1.9.1
go: finding github.com/go-sql-driver/mysql v1.4.1
go: finding github.com/stretchr/testify v1.2.2
go: finding github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a
go: finding golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33
go: finding golang.org/x/sys v0.0.0-20180830151530-49385e6e1522
go: finding github.com/go-logr/logr v0.1.0
go: finding google.golang.org/appengine v1.1.0
go: finding golang.org/x/tools v0.0.0-20180828015842-6cd1fcedba52
go: finding github.com/kardianos/osext v0.0.0-20170510131534-ae77be60afb1
go: finding github.com/kr/logfmt v0.0.0-20140226030751-b84e30acd515
go: finding github.com/go-chi/chi v3.3.3+incompatible
go: finding cloud.google.com/go v0.26.0
go: finding golang.org/x/net v0.0.0-20180826012351-8a410e7b638d
go: finding github.com/pmezard/go-difflib v1.0.0
go: finding golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f
go: finding github.com/stretchr/objx v0.1.1
go: finding golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
go: finding github.com/konsorten/go-windows-terminal-sequences v1.0.1
go: finding github.com/rs/zerolog v1.11.0
go: finding github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
go: finding github.com/satori/go.uuid v1.2.0
go: finding github.com/golang/protobuf v1.2.0
go: finding go.uber.org/atomic v1.3.2
go: finding golang.org/x/text v0.3.0
go: finding gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
go: finding github.com/confluentinc/confluent-kafka-go v0.11.6
go: finding google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8
go: finding github.com/bugsnag/panicwrap v1.2.0
go: finding github.com/davecgh/go-spew v1.1.1
go: finding github.com/kr/pretty v0.1.0
go: finding github.com/client9/misspell v0.3.4
go: finding github.com/kisielk/gotool v1.0.0
go: finding github.com/golang/mock v1.1.1
go: finding honnef.co/go/tools v0.0.0-20180728063816-88497007e858
go: finding golang.org/x/lint v0.0.0-20181026193005-c67002cb31c3
go: finding github.com/kr/text v0.1.0
go: finding github.com/kr/pty v1.1.1
go: downloading github.com/sirupsen/logrus v1.2.0
go: downloading github.com/goph/logur v0.11.0
go: extracting github.com/goph/logur v0.11.0
go: extracting github.com/sirupsen/logrus v1.2.0
go: downloading github.com/konsorten/go-windows-terminal-sequences v1.0.1
go: downloading golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
go: downloading github.com/stretchr/testify v1.2.2
go: extracting github.com/konsorten/go-windows-terminal-sequences v1.0.1
go: extracting github.com/stretchr/testify v1.2.2
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading github.com/pmezard/go-difflib v1.0.0
go: extracting github.com/pmezard/go-difflib v1.0.0
go: extracting github.com/davecgh/go-spew v1.1.1
go: extracting golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
go: downloading golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33
go: extracting golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33

real	0m8.215s
user	0m0.785s
sys 	0m0.765s

# du -sh /tmp/gopath-for-dependency-size-test-2
17M	/tmp/gopath-for-dependency-size-test-2

# time go mod download
go: finding github.com/sirupsen/logrus v1.2.0
go: finding github.com/goph/logur v0.11.0
go: finding github.com/go-kit/kit v0.8.0
go: finding github.com/rs/zerolog v1.11.0
go: finding github.com/InVisionApp/go-logger v1.0.1
go: finding github.com/hashicorp/go-hclog v0.0.0-20181001195459-61d530d6c27f
go: finding github.com/bugsnag/bugsnag-go v1.4.0
go: finding github.com/go-logfmt/logfmt v0.4.0
go: finding github.com/rollbar/rollbar-go v1.0.2
go: finding google.golang.org/grpc v1.17.0
go: finding go.uber.org/multierr v1.1.0
go: finding github.com/pkg/errors v0.8.0
go: finding golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
go: finding github.com/gofrs/uuid v3.1.0+incompatible
go: finding go.uber.org/atomic v1.3.2
go: finding github.com/go-sql-driver/mysql v1.4.1
go: finding github.com/stretchr/testify v1.2.2
go: finding github.com/stretchr/objx v0.1.1
go: finding golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33
go: finding golang.org/x/sys v0.0.0-20180830151530-49385e6e1522
go: finding golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
go: finding golang.org/x/text v0.3.0
go: finding cloud.google.com/go v0.26.0
go: finding github.com/kisielk/gotool v1.0.0
go: finding github.com/kr/logfmt v0.0.0-20140226030751-b84e30acd515
go: finding golang.org/x/lint v0.0.0-20181026193005-c67002cb31c3
go: finding golang.org/x/tools v0.0.0-20180828015842-6cd1fcedba52
go: finding golang.org/x/net v0.0.0-20180826012351-8a410e7b638d
go: finding golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f
go: finding github.com/ThreeDotsLabs/watermill v0.1.2
go: finding google.golang.org/appengine v1.1.0
go: finding github.com/konsorten/go-windows-terminal-sequences v1.0.1
go: finding github.com/client9/misspell v0.3.4
go: finding github.com/davecgh/go-spew v1.1.1
go: finding github.com/go-logr/logr v0.1.0
go: finding github.com/pmezard/go-difflib v1.0.0
go: finding google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8
go: finding github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
go: finding github.com/kardianos/osext v0.0.0-20170510131534-ae77be60afb1
go: finding github.com/go-chi/chi v3.3.3+incompatible
go: finding github.com/golang/mock v1.1.1
go: finding github.com/golang/protobuf v1.2.0
go: finding github.com/bugsnag/panicwrap v1.2.0
go: finding gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
go: finding github.com/kr/pretty v0.1.0
go: finding github.com/confluentinc/confluent-kafka-go v0.11.6
go: finding github.com/satori/go.uuid v1.2.0
go: finding github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a
go: finding honnef.co/go/tools v0.0.0-20180728063816-88497007e858
go: finding go.uber.org/zap v1.9.1
go: finding github.com/kr/text v0.1.0
go: finding github.com/kr/pty v1.1.1

real	0m11.990s
user	0m2.417s
sys 	0m2.211s

# du -sh /tmp/gopath-for-dependency-size-test-2
127M	/tmp/gopath-for-dependency-size-test-2
```


## Test 2 - MY PROXY

```
bash
export GOPATH=/tmp/gopath-for-dependency-size-test-3
export GOPROXY=$MYGOPROXY
time go mod tidy
du -sh /tmp/gopath-for-dependency-size-test-3
sudo rm -rf /tmp/gopath-for-dependency-size-test-3/*
time go mod download
du -sh /tmp/gopath-for-dependency-size-test-3
```

Output:

```
# time go mod tidy
go: finding github.com/goph/logur v0.11.0
go: finding github.com/sirupsen/logrus v1.2.0
go: finding github.com/rs/zerolog v1.11.0
go: finding github.com/go-kit/kit v0.8.0
go: finding github.com/bugsnag/bugsnag-go v1.4.0
go: finding github.com/rollbar/rollbar-go v1.0.2
go: finding github.com/hashicorp/go-hclog v0.0.0-20181001195459-61d530d6c27f
go: finding github.com/go-logfmt/logfmt v0.4.0
go: finding go.uber.org/multierr v1.1.0
go: finding github.com/InVisionApp/go-logger v1.0.1
go: finding google.golang.org/grpc v1.17.0
go: finding github.com/pkg/errors v0.8.0
go: finding golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
go: finding github.com/gofrs/uuid v3.1.0+incompatible
go: finding github.com/ThreeDotsLabs/watermill v0.1.2
go: finding github.com/go-sql-driver/mysql v1.4.1
go: finding cloud.google.com/go v0.26.0
go: finding golang.org/x/net v0.0.0-20180826012351-8a410e7b638d
go: finding golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33
go: finding google.golang.org/appengine v1.1.0
go: finding github.com/kisielk/gotool v1.0.0
go: finding golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
go: finding github.com/davecgh/go-spew v1.1.1
go: finding github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
go: finding google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8
go: finding github.com/pmezard/go-difflib v1.0.0
go: finding github.com/golang/mock v1.1.1
go: finding github.com/satori/go.uuid v1.2.0
go: finding go.uber.org/atomic v1.3.2
go: finding golang.org/x/lint v0.0.0-20181026193005-c67002cb31c3
go: finding github.com/stretchr/objx v0.1.1
go: finding gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
go: finding github.com/stretchr/testify v1.2.2
go: finding go.uber.org/zap v1.9.1
go: finding golang.org/x/sys v0.0.0-20180830151530-49385e6e1522
go: finding golang.org/x/text v0.3.0
go: finding github.com/go-chi/chi v3.3.3+incompatible
go: finding honnef.co/go/tools v0.0.0-20180728063816-88497007e858
go: finding github.com/go-logr/logr v0.1.0
go: finding github.com/konsorten/go-windows-terminal-sequences v1.0.1
go: finding github.com/kardianos/osext v0.0.0-20170510131534-ae77be60afb1
go: finding github.com/client9/misspell v0.3.4
go: finding github.com/bugsnag/panicwrap v1.2.0
go: finding golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f
go: finding github.com/golang/protobuf v1.2.0
go: finding github.com/kr/logfmt v0.0.0-20140226030751-b84e30acd515
go: finding github.com/kr/pretty v0.1.0
go: finding github.com/confluentinc/confluent-kafka-go v0.11.6
go: finding golang.org/x/tools v0.0.0-20180828015842-6cd1fcedba52
go: finding github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a
go: finding github.com/kr/text v0.1.0
go: finding github.com/kr/pty v1.1.1
go: downloading github.com/sirupsen/logrus v1.2.0
go: downloading github.com/goph/logur v0.11.0
go: extracting github.com/sirupsen/logrus v1.2.0
go: extracting github.com/goph/logur v0.11.0
go: downloading github.com/konsorten/go-windows-terminal-sequences v1.0.1
go: downloading golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
go: downloading github.com/stretchr/testify v1.2.2
go: extracting github.com/konsorten/go-windows-terminal-sequences v1.0.1
go: extracting github.com/stretchr/testify v1.2.2
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading github.com/pmezard/go-difflib v1.0.0
go: extracting github.com/davecgh/go-spew v1.1.1
go: extracting github.com/pmezard/go-difflib v1.0.0
go: extracting golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
go: downloading golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33
go: extracting golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33

real	0m5.851s
user	0m0.651s
sys 	0m0.749s

# du -sh /tmp/gopath-for-dependency-size-test-3
17M	/tmp/gopath-for-dependency-size-test-3

# time go mod download
go: finding github.com/sirupsen/logrus v1.2.0
go: finding github.com/goph/logur v0.11.0
go: finding github.com/davecgh/go-spew v1.1.1
go: finding github.com/pmezard/go-difflib v1.0.0
go: finding golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33
go: finding github.com/stretchr/objx v0.1.1
go: finding golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
go: finding github.com/konsorten/go-windows-terminal-sequences v1.0.1
go: finding github.com/stretchr/testify v1.2.2
go: finding go.uber.org/atomic v1.3.2
go: finding go.uber.org/multierr v1.1.0
go: finding github.com/bugsnag/bugsnag-go v1.4.0
go: finding google.golang.org/grpc v1.17.0
go: finding github.com/gofrs/uuid v3.1.0+incompatible
go: finding github.com/hashicorp/go-hclog v0.0.0-20181001195459-61d530d6c27f
go: finding github.com/bugsnag/panicwrap v1.2.0
go: finding github.com/go-logfmt/logfmt v0.4.0
go: finding github.com/go-sql-driver/mysql v1.4.1
go: finding github.com/kardianos/osext v0.0.0-20170510131534-ae77be60afb1
go: finding github.com/rollbar/rollbar-go v1.0.2
go: finding go.uber.org/zap v1.9.1
go: finding github.com/go-logr/logr v0.1.0
go: finding golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
go: finding google.golang.org/appengine v1.1.0
go: finding github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
go: finding golang.org/x/text v0.3.0
go: finding github.com/kisielk/gotool v1.0.0
go: finding golang.org/x/tools v0.0.0-20180828015842-6cd1fcedba52
go: finding google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8
go: finding github.com/ThreeDotsLabs/watermill v0.1.2
go: finding github.com/kr/logfmt v0.0.0-20140226030751-b84e30acd515
go: finding cloud.google.com/go v0.26.0
go: finding golang.org/x/lint v0.0.0-20181026193005-c67002cb31c3
go: finding github.com/golang/mock v1.1.1
go: finding github.com/client9/misspell v0.3.4
go: finding github.com/golang/protobuf v1.2.0
go: finding golang.org/x/sys v0.0.0-20180830151530-49385e6e1522
go: finding golang.org/x/net v0.0.0-20180826012351-8a410e7b638d
go: finding github.com/InVisionApp/go-logger v1.0.1
go: finding github.com/go-chi/chi v3.3.3+incompatible
go: finding honnef.co/go/tools v0.0.0-20180728063816-88497007e858
go: finding golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f
go: finding github.com/pkg/errors v0.8.0
go: finding gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
go: finding github.com/kr/pretty v0.1.0
go: finding github.com/confluentinc/confluent-kafka-go v0.11.6
go: finding github.com/satori/go.uuid v1.2.0
go: finding github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a
go: finding github.com/rs/zerolog v1.11.0
go: finding github.com/go-kit/kit v0.8.0
go: finding github.com/kr/text v0.1.0
go: finding github.com/kr/pty v1.1.1

real	0m6.811s
user	0m2.211s
sys 	0m1.717s

# du -sh /tmp/gopath-for-dependency-size-test-3
126M	/tmp/gopath-for-dependency-size-test-3
```
