package fyaml

import (
	"github.com/go-yaml/yaml"
)

type Serializer struct{}

func NewSerializer() *Serializer {
	return &Serializer{}
}

func (s *Serializer) Marshal(input interface{}) ([]byte, error) {
	return yaml.Marshal(input)
}

func (s *Serializer) Unmarshal(data []byte, output interface{}) error {
	return yaml.Unmarshal(data, output)
}
