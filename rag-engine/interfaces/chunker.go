package interfaces

import (
	dm "rag-engine/app/models/schema"

	"github.com/jdkato/prose/v2"
)

type Chunker interface {
	ProcessDocument([]prose.Document, map[string]string) ([]dm.ChunkedDocument, error)
}
