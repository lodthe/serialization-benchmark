package fmsgpack

import "github.com/vmihailenco/msgpack/v5"

type Serializer struct{}

func NewSerializer() *Serializer {
	return &Serializer{}
}

func (s *Serializer) Marshal(input interface{}) ([]byte, error) {
	return msgpack.Marshal(input)
}

func (s *Serializer) Unmarshal(data []byte, output interface{}) error {
	return msgpack.Unmarshal(data, output)
}
