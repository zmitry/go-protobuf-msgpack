# go-protbuf-msgpack
Compare the difference between protocol buffer and message pack in golang

## Benchmark

- Protocol Buffer
- Message Pack
- and JSON

### Serialization
```
BenchmarkProtobuf-8      1000000              1640 ns/op            1392 B/op         12 allocs/op
BenchmarkMsgpack-8        500000              3400 ns/op            2600 B/op         24 allocs/op
BenchmarkJSON-8           200000              8348 ns/op            2224 B/op         18 allocs/op
```

### Data length
```
Content-Type: application/protobuf
Date: Mon, 24 Sep 2018 14:49:51 GMT
Content-Length: 505
```

```
Content-Type: application/x-msgpack
Date: Mon, 24 Sep 2018 14:48:58 GMT
Content-Length: 539
```

```
Content-Type: application/json
Date: Mon, 24 Sep 2018 14:50:14 GMT
Content-Length: 563
```
