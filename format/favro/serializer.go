package favro

import (
	"log"

	"github.com/hamba/avro"
)

type Serializer struct{
	schema avro.Schema
}

func NewSerializer() *Serializer {
	parsed, err := avro.Parse(schema)
	if err != nil {
		log.Fatalf("failed to parse AVRO schema: %v\n", err)
	}

	return &Serializer{
		schema: parsed,
	}
}

func (s *Serializer) Marshal(input interface{}) ([]byte, error) {
	return avro.Marshal(s.schema, input)
}

func (s *Serializer) Unmarshal(data []byte, output interface{}) error {
	return avro.Unmarshal(s.schema, data, output)
}
