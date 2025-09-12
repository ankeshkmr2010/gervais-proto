package schema

import "github.com/google/uuid"

type Document struct {
	DocumentId uuid.UUID         `json:"doc_id"`
	MetaData   map[string]string `json:"meta_data"`
	Content    string            `json:"content"`
}

type ChunkedDocument struct {
	ChunkId     uuid.UUID         `json:"chunk_d"`
	ChunkSerial int               `json:"chunk_serial"`
	DocumentId  uuid.UUID         `json:"doc_id"`
	MetaData    map[string]string `json:"meta_data"`
	Content     string            `json:"content"`
	Embedding   []float64         `json:"embedding"`
}
