## 压测列表

| name      | package                                          | transport | codec    |     tps |
| --------- | ------------------------------------------------ | --------- | -------- | ------: |
| grpc      | [grpc](https://github.com/grpc/grpc-go)          | grpc      | protobuf | `47015` |
| rpcx01    | [rpcx/v5](https://github.com/smallnest/rpcx)     | tcp       | json     | `36211` |
| rpcx02    | [rpcx/v5](https://github.com/smallnest/rpcx)     | tcp       | protobuf | `70442` |
| rpcx03    | [rpcx/v5](https://github.com/smallnest/rpcx)     | tcp       | msgpack  | `39274` |
| rpcx04    | [rpcx/v5](https://github.com/smallnest/rpcx)     | quic      | json     | `21280` |
| rpcx05    | [rpcx/v5](https://github.com/smallnest/rpcx)     | quic      | protobuf | `31102` |
| rpcx06    | [rpcx/v5](https://github.com/smallnest/rpcx)     | quic      | msgpack  | `22943` |
| gomicro01 | [go-micro/v2](https://github.com/micro/go-micro) | grpc      | protobuf | `23336` |
| twirp01   | [twirp/v7](https://github.com/twitchtv/twirp)    | http      | json     |  `6251` |
| twirp02   | [twirp/v7](https://github.com/twitchtv/twirp)    | http      | protobuf | `21432` |

## 详细结果

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
2020-10-13T01:21:17.037+08:00 INF 1 main.go:62 >  clients=100 requests=1000 total=100000
2020-10-13T01:21:21.772+08:00 INF 1 main.go:85 >  tps=21280 min=198.343µs max=65.501679ms mean=4.606586ms median=3.483975ms
```

### rpcx05

```
2020-10-13T01:21:50.088+08:00 INF 1 main.go:62 >  clients=100 requests=1000 total=100000
2020-10-13T01:21:53.332+08:00 INF 1 main.go:85 >  tps=31102 min=148.555µs max=69.896592ms mean=3.158439ms median=2.449258ms
```

### rpcx06

```
2020-10-13T01:22:18.433+08:00 INF 1 main.go:62 >  clients=100 requests=1000 total=100000
2020-10-13T01:22:22.822+08:00 INF 1 main.go:85 >  tps=22943 min=191.919µs max=64.135991ms mean=4.285427ms median=3.329058ms
```

### gomicro01

```
2020-10-13T00:06:42.697+08:00 INF 1 main.go:47 >  clients=100 requests=1000 total=100000
2020-10-13T00:06:47.015+08:00 INF 1 main.go:70 >  tps=23336 min=218.587µs max=21.975576ms mean=4.224776ms median=3.562678ms
```

### twirp01

```
2020-10-13T00:07:47.758+08:00 INF 1 main.go:48 >  clients=100 requests=100 total=10000
2020-10-13T00:07:49.360+08:00 INF 1 main.go:71 >  tps=6251 min=455.483µs max=87.317299ms mean=15.293143ms median=12.795295ms
```

### twirp02

```
2020-10-13T00:08:35.644+08:00 INF 1 main.go:48 >  clients=100 requests=100 total=10000
2020-10-13T00:08:36.113+08:00 INF 1 main.go:71 >  tps=21432 min=130.911µs max=33.457038ms mean=4.420915ms median=3.550289ms
```

## TODO

- [kite](https://github.com/koding/kite)
- [go-chassis](https://github.com/go-chassis/go-chassis)
- [TarsGo](https://github.com/TarsCloud/TarsGo)
- [erpc](https://github.com/henrylee2cn/erpc)
- [go-zero](https://github.com/tal-tech/go-zero)
- [hprose](https://github.com/hprose/hprose-golang)
- [arpc](https://github.com/lesismal/arpc)
