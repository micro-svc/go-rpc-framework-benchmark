## 压测列表

| name      | type  | package                                          | transport | codec    |     tps |
| --------- | ----- | ------------------------------------------------ | --------- | -------- | ------: |
| gin       | `WEB` | [gin](https://github.com/gin-gonic/gin)          | http      | json     | `28674` |
| grpc      | `RPC` | [grpc](https://github.com/grpc/grpc-go)          | grpc      | protobuf | `46998` |
| rpcx01    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | tcp       | json     | `36265` |
| rpcx02    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | tcp       | protobuf | `72328` |
| rpcx03    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | tcp       | msgpack  | `39970` |
| rpcx04    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | tcp       | jsoniter | `52698` |
| rpcx05    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | quic      | json     | `23173` |
| rpcx06    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | quic      | protobuf | `32436` |
| rpcx07    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | quic      | msgpack  | `23937` |
| rpcx08    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | quic      | jsoniter | `27995` |
| gomicro01 | `MIC` | [go-micro/v2](https://github.com/micro/go-micro) | grpc      | protobuf | `22380` |
| twirp01   | `RPC` | [twirp/v7](https://github.com/twitchtv/twirp)    | http      | json     |  `8360` |
| twirp02   | `RPC` | [twirp/v7](https://github.com/twitchtv/twirp)    | http      | protobuf | `41773` |

> 类型说明：`RPC` 表示只实现了简单的 RPC 功能；`MIC` 表示实现了微服务框架的大部分特性；`WEB` 表示传统的 WEB 框架，只是作为对比出现。

## 详细结果

### gin

```sh
$ bin/gin-cli
{"time":"2020-10-13T15:07:50.923+08:00","level":"info","caller":"main.go:38","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T15:07:54.441+08:00","level":"info","caller":"main.go:61","goid":1,"tps":28674,"min":"142.762µs","max":"28.851624ms","mean":"3.417126ms","median":"2.354175ms"}
```

### grpc

```sh
$ bin/grpc-cli
{"time":"2020-10-13T15:08:13.273+08:00","level":"info","caller":"main.go:34","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T15:08:15.432+08:00","level":"info","caller":"main.go:57","goid":1,"tps":46998,"min":"126.135µs","max":"20.396811ms","mean":"2.106846ms","median":"1.870091ms"}
```

### rpcx01

```sh
$ bin/rpcx-v5-cli --transport=tcp --codec=json
{"time":"2020-10-13T15:08:34.319+08:00","level":"info","caller":"main.go:65","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T15:08:37.111+08:00","level":"info","caller":"main.go:88","goid":1,"tps":36265,"min":"103.798µs","max":"31.059308ms","mean":"2.69807ms","median":"1.966065ms"}
```

### rpcx02

```sh
$ bin/rpcx-v5-cli --transport=tcp --codec=protobuf
{"time":"2020-10-13T15:08:55.790+08:00","level":"info","caller":"main.go:65","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T15:08:57.206+08:00","level":"info","caller":"main.go:88","goid":1,"tps":72328,"min":"74.125µs","max":"20.335373ms","mean":"1.353838ms","median":"1.088218ms"}
```

### rpcx03

```sh
$ bin/rpcx-v5-cli --transport=tcp --codec=msgpack
{"time":"2020-10-13T15:09:15.862+08:00","level":"info","caller":"main.go:65","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T15:09:18.396+08:00","level":"info","caller":"main.go:88","goid":1,"tps":39970,"min":"127.019µs","max":"32.079644ms","mean":"2.447469ms","median":"1.817598ms"}
```

### rpcx04

```sh
$ bin/rpcx-v5-cli --transport=tcp --codec=jsoniter
{"time":"2020-10-13T15:09:37.061+08:00","level":"info","caller":"main.go:65","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T15:09:38.990+08:00","level":"info","caller":"main.go:88","goid":1,"tps":52698,"min":"100.414µs","max":"21.543921ms","mean":"1.859224ms","median":"1.451236ms"}
```

### rpcx05

```sh
$ bin/rpcx-v5-cli --transport=quic --codec=json
{"time":"2020-10-13T15:09:57.966+08:00","level":"info","caller":"main.go:65","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T15:10:02.312+08:00","level":"info","caller":"main.go:88","goid":1,"tps":23173,"min":"170.903µs","max":"71.377275ms","mean":"4.229565ms","median":"3.138299ms"}
```

### rpcx06

```sh
$ bin/rpcx-v5-cli --transport=quic --codec=protobuf
{"time":"2020-10-13T15:10:21.304+08:00","level":"info","caller":"main.go:65","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T15:10:24.417+08:00","level":"info","caller":"main.go:88","goid":1,"tps":32436,"min":"138.89µs","max":"65.92925ms","mean":"3.025885ms","median":"2.378947ms"}
```

### rpcx07

```sh
$ bin/rpcx-v5-cli --transport=quic --codec=msgpack
{"time":"2020-10-13T15:10:43.398+08:00","level":"info","caller":"main.go:65","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T15:10:47.612+08:00","level":"info","caller":"main.go:88","goid":1,"tps":23937,"min":"191.618µs","max":"81.007022ms","mean":"4.101223ms","median":"3.133673ms"}
```

### rpcx08

```sh
$ bin/rpcx-v5-cli --transport=quic --codec=jsoniter
{"time":"2020-10-13T15:11:06.610+08:00","level":"info","caller":"main.go:65","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T15:11:10.215+08:00","level":"info","caller":"main.go:88","goid":1,"tps":27995,"min":"189.652µs","max":"61.616075ms","mean":"3.514571ms","median":"2.75031ms"}
```

### gomicro01

```sh
$ bin/go-micro-v2-cli
{"time":"2020-10-13T15:31:51.139+08:00","level":"info","caller":"main.go:46","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T15:31:55.645+08:00","level":"info","caller":"main.go:69","goid":1,"tps":22380,"min":"253.41µs","max":"31.187008ms","mean":"4.404891ms","median":"3.616454ms"}
```

### twirp01

```sh
$ bin/twirp-v7-cli --codec=json
{"time":"2020-10-13T15:11:28.995+08:00","level":"info","caller":"main.go:49","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T15:11:40.988+08:00","level":"info","caller":"main.go:72","goid":1,"tps":8360,"min":"349.523µs","max":"131.736047ms","mean":"11.75061ms","median":"8.009684ms"}
```

### twirp02

```sh
$ bin/twirp-v7-cli --codec=protobuf
{"time":"2020-10-13T15:11:59.640+08:00","level":"info","caller":"main.go:49","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T15:12:02.066+08:00","level":"info","caller":"main.go:72","goid":1,"tps":41773,"min":"105.569µs","max":"29.357838ms","mean":"2.342483ms","median":"1.614076ms"}
```

## TODO

- [kite](https://github.com/koding/kite)
- [go-chassis](https://github.com/go-chassis/go-chassis)
- [TarsGo](https://github.com/TarsCloud/TarsGo)
- [erpc](https://github.com/henrylee2cn/erpc)
- [go-zero](https://github.com/tal-tech/go-zero)
- [hprose](https://github.com/hprose/hprose-golang)
- [arpc](https://github.com/lesismal/arpc)
