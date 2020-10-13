## 压测列表

> 类型说明：
> `RPC` 表示只实现了简单的 RPC 功能；
> `MIC` 表示实现了微服务框架的大部分特性；
> `WEB` 表示传统的 WEB 框架，只是作为对比出现。

| name      | type  | package                                          | transport | codec      |     tps |
| --------- | ----- | ------------------------------------------------ | --------- | ---------- | ------: |
| gin       | `WEB` | [gin](https://github.com/gin-gonic/gin)          | `http`    | `json`     | `27017` |
| grpc      | `RPC` | [grpc](https://github.com/grpc/grpc-go)          | `grpc`    | `protobuf` | `45241` |
| rpcx01    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | `tcp`     | `json`     | `27085` |
| rpcx02    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | `tcp`     | `protobuf` | `64934` |
| rpcx03    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | `tcp`     | `msgpack`  | `33644` |
| rpcx04    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | `tcp`     | `jsoniter` | `49558` |
| rpcx05    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | `quic`    | `json`     | `21907` |
| rpcx06    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | `quic`    | `protobuf` | `30908` |
| rpcx07    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | `quic`    | `msgpack`  | `22555` |
| rpcx08    | `MIC` | [rpcx/v5](https://github.com/smallnest/rpcx)     | `quic`    | `jsoniter` | `22487` |
| gomicro01 | `MIC` | [go-micro/v2](https://github.com/micro/go-micro) | `grpc`    | `json`     |  `5870` |
| gomicro02 | `MIC` | [go-micro/v2](https://github.com/micro/go-micro) | `grpc`    | `protobuf` | `21972` |
| gomicro03 | `MIC` | [go-micro/v2](https://github.com/micro/go-micro) | `http`    | `json`     |  `3388` |
| gomicro04 | `MIC` | [go-micro/v2](https://github.com/micro/go-micro) | `http`    | `protobuf` |  `3938` |
| twirp01   | `RPC` | [twirp/v7](https://github.com/twitchtv/twirp)    | `http`    | `json`     |  `7386` |
| twirp02   | `RPC` | [twirp/v7](https://github.com/twitchtv/twirp)    | `http`    | `protobuf` | `32878` |

### gin

```sh
$ bin/gin-cli
{"time":"2020-10-13T17:49:22.283+08:00","level":"info","caller":"main.go:38","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T17:49:26.018+08:00","level":"info","caller":"main.go:61","goid":1,"tps":27017,"min":"158.474µs","max":"41.386869ms","mean":"3.632799ms","median":"2.497086ms"}
```

### grpc

```sh
$ bin/grpc-cli
{"time":"2020-10-13T17:49:45.038+08:00","level":"info","caller":"main.go:34","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T17:49:47.282+08:00","level":"info","caller":"main.go:57","goid":1,"tps":45241,"min":"126.237µs","max":"34.400997ms","mean":"2.192001ms","median":"1.916892ms"}
```

### rpcx01

```sh
$ bin/rpcx-v5-cli --transport=tcp --codec=json
{"time":"2020-10-13T17:50:06.451+08:00","level":"info","caller":"main.go:65","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T17:50:10.175+08:00","level":"info","caller":"main.go:88","goid":1,"tps":27085,"min":"122.133µs","max":"77.277867ms","mean":"3.613228ms","median":"2.277533ms"}
```

### rpcx02

```sh
$ bin/rpcx-v5-cli --transport=tcp --codec=protobuf
{"time":"2020-10-13T17:50:28.869+08:00","level":"info","caller":"main.go:65","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T17:50:30.441+08:00","level":"info","caller":"main.go:88","goid":1,"tps":64934,"min":"80.605µs","max":"35.330013ms","mean":"1.501494ms","median":"1.153133ms"}
```

### rpcx03

```sh
$ bin/rpcx-v5-cli --transport=tcp --codec=msgpack
{"time":"2020-10-13T17:50:49.161+08:00","level":"info","caller":"main.go:65","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T17:50:52.165+08:00","level":"info","caller":"main.go:88","goid":1,"tps":33644,"min":"124.953µs","max":"42.136381ms","mean":"2.896914ms","median":"2.053496ms"}
```

### rpcx04

```sh
$ bin/rpcx-v5-cli --transport=tcp --codec=jsoniter
{"time":"2020-10-13T17:51:10.873+08:00","level":"info","caller":"main.go:65","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T17:51:12.924+08:00","level":"info","caller":"main.go:88","goid":1,"tps":49558,"min":"100.778µs","max":"19.068931ms","mean":"1.98144ms","median":"1.500624ms"}
```

### rpcx05

```sh
$ bin/rpcx-v5-cli --transport=quic --codec=json
{"time":"2020-10-13T17:51:31.955+08:00","level":"info","caller":"main.go:65","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T17:51:36.550+08:00","level":"info","caller":"main.go:88","goid":1,"tps":21907,"min":"199.796µs","max":"66.075861ms","mean":"4.473687ms","median":"3.408147ms"}
```

### rpcx06

```sh
$ bin/rpcx-v5-cli --transport=quic --codec=protobuf
{"time":"2020-10-13T17:51:55.552+08:00","level":"info","caller":"main.go:65","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T17:51:58.821+08:00","level":"info","caller":"main.go:88","goid":1,"tps":30908,"min":"139.266µs","max":"52.673411ms","mean":"3.18012ms","median":"2.511474ms"}
```

### rpcx07

```sh
$ bin/rpcx-v5-cli --transport=quic --codec=msgpack
{"time":"2020-10-13T17:52:17.893+08:00","level":"info","caller":"main.go:65","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T17:52:22.357+08:00","level":"info","caller":"main.go:88","goid":1,"tps":22555,"min":"185.701µs","max":"72.337652ms","mean":"4.34077ms","median":"3.329371ms"}
```

### rpcx08

```sh
$ bin/rpcx-v5-cli --transport=quic --codec=jsoniter
{"time":"2020-10-13T17:52:41.333+08:00","level":"info","caller":"main.go:65","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T17:52:45.827+08:00","level":"info","caller":"main.go:88","goid":1,"tps":22487,"min":"173.718µs","max":"91.865421ms","mean":"4.344471ms","median":"3.149399ms"}
```

### gomicro01

```sh
$ bin/go-micro-v2-cli --transport=grpc --codec=json
{"time":"2020-10-13T17:53:05.175+08:00","level":"info","caller":"main.go:79","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T17:53:22.241+08:00","level":"info","caller":"main.go:102","goid":1,"tps":5870,"min":"645.393µs","max":"140.172311ms","mean":"16.735734ms","median":"12.238477ms"}
```

### gomicro02

```sh
$ bin/go-micro-v2-cli --transport=grpc --codec=protobuf
{"time":"2020-10-13T17:53:41.064+08:00","level":"info","caller":"main.go:79","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T17:53:45.654+08:00","level":"info","caller":"main.go:102","goid":1,"tps":21972,"min":"254.284µs","max":"98.399354ms","mean":"4.470302ms","median":"3.576929ms"}
```

### gomicro03

```sh
$ bin/go-micro-v2-cli --transport=http --codec=json
{"time":"2020-10-13T17:54:04.688+08:00","level":"info","caller":"main.go:79","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T17:54:34.232+08:00","level":"info","caller":"main.go:102","goid":1,"tps":3388,"min":"1.022532ms","max":"570.629223ms","mean":"29.244503ms","median":"24.885394ms"}
```

### gomicro04

```sh
$ bin/go-micro-v2-cli --transport=http --codec=protobuf
{"time":"2020-10-13T17:54:53.281+08:00","level":"info","caller":"main.go:79","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T17:55:18.701+08:00","level":"info","caller":"main.go:102","goid":1,"tps":3938,"min":"690.177µs","max":"182.822138ms","mean":"25.330691ms","median":"25.16199ms"}
```

### twirp01

```sh
$ bin/twirp-v7-cli --codec=json
{"time":"2020-10-13T17:55:37.671+08:00","level":"info","caller":"main.go:49","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T17:55:51.240+08:00","level":"info","caller":"main.go:72","goid":1,"tps":7386,"min":"383.642µs","max":"164.931959ms","mean":"13.269275ms","median":"9.160097ms"}
```

### twirp02

```sh
$ bin/twirp-v7-cli --codec=protobuf
{"time":"2020-10-13T17:56:09.902+08:00","level":"info","caller":"main.go:49","goid":1,"clients":100,"requests":1000,"total":100000}
{"time":"2020-10-13T17:56:12.977+08:00","level":"info","caller":"main.go:72","goid":1,"tps":32878,"min":"103.053µs","max":"54.973644ms","mean":"2.973557ms","median":"1.743029ms"}
```

## TODO

- [go-chassis](https://github.com/go-chassis/go-chassis)
- [TarsGo](https://github.com/TarsCloud/TarsGo)
- [erpc](https://github.com/henrylee2cn/erpc)
- [go-zero](https://github.com/tal-tech/go-zero)
