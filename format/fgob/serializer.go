package fgob

import (
	"bytes"
	"encoding/gob"
)

type Serializer struct{}

func NewSerializer() *Serializer {
	return &Serializer{}
}

func (s *Serializer) Marshal(input interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)

	err := enc.Encode(input)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (s *Serializer) Unmarshal(data []byte, output interface{}) error {
	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(output)

	return err
}
