package fjson

import (
	"encoding/json"
)

type Serializer struct{}

func NewSerializer() *Serializer {
	return &Serializer{}
}

func (s *Serializer) Marshal(input interface{}) ([]byte, error) {
	return json.Marshal(input)
}

func (s *Serializer) Unmarshal(data []byte, output interface{}) error {
	return json.Unmarshal(data, output)
}
