# uberserver

Test of all go-micro services running in 1 uberserver

## Getting Started

### Prerequisites

Install Consul
[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

Run Consul
```
$ consul agent -dev -advertise=127.0.0.1
```

### Run Service

```
$ go run main.go
```

### Run uberserver under gin (github.com/codegangsta/gin) - continouos uberserver rebuild, restart on file save

```
$ ./gin-uber.sh
```
