package format

type Serializer interface {
	Marshal(input interface{}) ([]byte, error)
	Unmarshal(data []byte, output interface{}) error
}
