package fxml

import (
	"encoding/xml"
)

type Serializer struct{}

func NewSerializer() *Serializer {
	return &Serializer{}
}

func (s *Serializer) Marshal(input interface{}) ([]byte, error) {
	return xml.MarshalIndent(input, " ", "  ")
}

func (s *Serializer) Unmarshal(data []byte, output interface{}) error {
	return xml.Unmarshal(data, output)
}
