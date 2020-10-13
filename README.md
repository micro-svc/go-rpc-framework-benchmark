## 压测列表

| name      | type  | package                                          | transport | codec    |     tps |
| --------- | ----- | ------------------------------------------------ | --------- | -------- | ------: |
| gin       | `WEB` | [gin](https://github.com/gin-gonic/gin)          | http      | json     | `26694` |
| grpc      | `RPC` | [grpc](https://github.com/grpc/grpc-go)          | grpc      | protobuf | `47015` |
| rpcx01    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | tcp       | json     | `36211` |
| rpcx02    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | tcp       | protobuf | `70442` |
| rpcx03    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | tcp       | msgpack  | `39274` |
| rpcx04    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | tcp       | jsoniter | `50385` |
| rpcx05    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | quic      | json     | `21280` |
| rpcx06    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | quic      | protobuf | `31102` |
| rpcx07    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | quic      | msgpack  | `22943` |
| rpcx08    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | quic      | jsoniter | `26421` |
| gomicro01 | `MIC` | [go-micro/v2](https://github.com/micro/go-micro) | grpc      | protobuf | `23336` |
| twirp01   | `RPC` | [twirp/v7](https://github.com/twitchtv/twirp)    | http      | json     |  `7728` |
| twirp02   | `RPC` | [twirp/v7](https://github.com/twitchtv/twirp)    | http      | protobuf | `38688` |

> 类型说明：`RPC` 表示只实现了简单的 RPC 功能；`MIC` 表示实现了微服务框架的大部分特性；`WEB` 表示传统的 WEB 框架，只是作为对比出现。

## 详细结果

### gin

```
2020-10-13T13:06:35.941+08:00 INF 1 main.go:39 >  clients=100 requests=1000 total=100000
2020-10-13T13:06:39.721+08:00 INF 1 main.go:62 >  tps=26694 min=170.532µs max=39.744728ms mean=3.681064ms median=2.45578ms
```

### grpc

```
2020-10-13T00:03:32.134+08:00 INF 1 main.go:43 >  clients=100 requests=1000 total=100000
2020-10-13T00:03:34.293+08:00 INF 1 main.go:66 >  tps=47015 min=137.933µs max=21.070551ms mean=2.108709ms median=1.875183ms
```

### rpcx01

```
2020-10-13T00:04:32.646+08:00 INF 1 main.go:59 >  clients=100 requests=1000 total=100000
2020-10-13T00:04:35.441+08:00 INF 1 main.go:82 >  tps=36211 min=114.641µs max=44.761202ms mean=2.708693ms median=2.017694ms
```

### rpcx02

```
2020-10-13T00:05:12.396+08:00 INF 1 main.go:59 >  clients=100 requests=1000 total=100000
2020-10-13T00:05:13.849+08:00 INF 1 main.go:82 >  tps=70442 min=69.306µs max=27.193609ms mean=1.393627ms median=1.11162ms
```

### rpcx03

```
2020-10-13T00:05:59.292+08:00 INF 1 main.go:59 >  clients=100 requests=1000 total=100000
2020-10-13T00:06:01.874+08:00 INF 1 main.go:82 >  tps=39274 min=146.77µs max=39.616732ms mean=2.496386ms median=1.92411ms
```

### rpcx04

```
2020-10-13T11:02:01.346+08:00 INF 1 main.go:66 >  clients=100 requests=1000 total=100000
2020-10-13T11:02:03.364+08:00 INF 1 main.go:89 >  tps=50385 min=94.149µs max=25.531931ms mean=1.944109ms median=1.494184ms
```

### rpcx05

```
2020-10-13T01:21:17.037+08:00 INF 1 main.go:62 >  clients=100 requests=1000 total=100000
2020-10-13T01:21:21.772+08:00 INF 1 main.go:85 >  tps=21280 min=198.343µs max=65.501679ms mean=4.606586ms median=3.483975ms
```

### rpcx06

```
2020-10-13T01:21:50.088+08:00 INF 1 main.go:62 >  clients=100 requests=1000 total=100000
2020-10-13T01:21:53.332+08:00 INF 1 main.go:85 >  tps=31102 min=148.555µs max=69.896592ms mean=3.158439ms median=2.449258ms
```

### rpcx07

```
2020-10-13T01:22:18.433+08:00 INF 1 main.go:62 >  clients=100 requests=1000 total=100000
2020-10-13T01:22:22.822+08:00 INF 1 main.go:85 >  tps=22943 min=191.919µs max=64.135991ms mean=4.285427ms median=3.329058ms
```

### rpcx08

```
2020-10-13T11:05:07.709+08:00 INF 1 main.go:66 >  clients=100 requests=1000 total=100000
2020-10-13T11:05:11.525+08:00 INF 1 main.go:89 >  tps=26421 min=171.441µs max=60.401377ms mean=3.717221ms median=2.901175ms
```

### gomicro01

```
2020-10-13T00:06:42.697+08:00 INF 1 main.go:47 >  clients=100 requests=1000 total=100000
2020-10-13T00:06:47.015+08:00 INF 1 main.go:70 >  tps=23336 min=218.587µs max=21.975576ms mean=4.224776ms median=3.562678ms
```

### twirp01

```
2020-10-13T11:56:17.667+08:00 INF 1 main.go:50 >  clients=100 requests=1000 total=100000
2020-10-13T11:56:30.639+08:00 INF 1 main.go:73 >  tps=7728 min=438.821µs max=153.618607ms mean=12.732407ms median=8.769251ms
```

### twirp02

```
2020-10-13T11:55:23.837+08:00 INF 1 main.go:50 >  clients=100 requests=1000 total=100000
2020-10-13T11:55:26.453+08:00 INF 1 main.go:73 >  tps=38688 min=111.06µs max=26.905751ms mean=2.540243ms median=1.726098ms
```

## TODO

- [kite](https://github.com/koding/kite)
- [go-chassis](https://github.com/go-chassis/go-chassis)
- [TarsGo](https://github.com/TarsCloud/TarsGo)
- [erpc](https://github.com/henrylee2cn/erpc)
- [go-zero](https://github.com/tal-tech/go-zero)
- [hprose](https://github.com/hprose/hprose-golang)
- [arpc](https://github.com/lesismal/arpc)
