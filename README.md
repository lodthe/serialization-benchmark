# Serialization benchmark

Comparison of different serialization methods in Go.

## Tested serialization methods:

- [json](https://pkg.go.dev/encoding/json)
- [xml](https://pkg.go.dev/encoding/xml)
- [gob](https://golang.org/pkg/encoding/gob)
- [yaml](https://pkg.go.dev/gopkg.in/yaml.v2)
- [MessagePack](https://github.com/vmihailenco/msgpack)
- [Avro (hamba)](https://github.com/hamba/avro)
- [Protobuf](https://github.com/golang/protobuf)

## Running

Run the following command to start the benchmark:
```shell
go run main.go
```

Benchmark results will be printed to the stdout. 

Also, marshalled data will be saved in the [output](./output) directory:
```shell
output/
├── avro.txt
├── gob.txt
├── json.txt
├── msgpack.txt
├── protobuf.txt
├── xml.txt
└── yaml.txt
```

## Data

There is a synthetically generated user object to measure performance of serialization methods. You can find the definition of the structure in the [sample](./sample/sample.go) package:
```go
type User struct {
	Name    string
	Phone   string
	Balance float64

	BirthDay     *time.Time
	RegisteredAt time.Time

	ShoppingCart ShoppingCart

	Blocked bool
}

type ShoppingCart map[string]CartItem

type CartItem struct {
	Item     Item
	Quantity int32
}

type Item struct {
	ID        string
	CreatedAt time.Time
	Visible   bool
	OwnerID   int32

	Name        string
	Description *string
	Keywords    []string

	Price  float64
	Weight float32
}
```

## Results

|  Method  | Marshalled data size (bytes) | Mean Marshal time (µs) | Mean Unmarshal time (µs) |
|:--------:|:----------------------------:|:-------------------------:|:---------------------------:|
|   avro   |             86302            |             94            |             159             |
| protobuf |             87389            |            130            |             238             |
|    gob   |             87130            |            281            |             288             |
|  msgpack |             94001            |            200            |             297             |
|   json   |            113024            |            523            |             1624            |
|    xml   |            132550            |            1355           |             5214            |
|   yaml   |            113957            |            5318           |             5151            |