package chunker

import (
	"errors"
	ifc "rag-engine/app/interfaces"
	scm "rag-engine/app/models/schema"
	utils "rag-engine/app/utils"

	"github.com/google/uuid"
	"github.com/jdkato/prose/v2"
)

const (
	MAX_CHUNK_SIZE = 2000 // max chunk size in characters
)

type SimpleSemanticChunker struct {
	eg ifc.EmbeddingsGetter
}

func NewSimpleSemanticChunker(eg ifc.EmbeddingsGetter) ifc.Chunker {
	return &SimpleSemanticChunker{
		eg: eg,
	}
}

func (ssc *SimpleSemanticChunker) ProcessDocument(docs []prose.Document, params map[string]string) ([]scm.ChunkedDocument, error) {
	if len(docs) == 0 {
		return nil, nil
	}
	if ssc.eg == nil {
		return nil, errors.New("embeddings getter is not initialized")
	}
	chunkedStrings := make([]scm.ChunkedDocument, 0)
	lastChunk := scm.ChunkedDocument{}

	for _, doc := range docs {
		// use token splitting logic here to split the document into chunks

		for _, line := range doc.Sentences() {
			embedding, err := ssc.eg.GetEmbeddings(line.Text)
			if err != nil {
				return nil, err
			}
			// first chunk
			if len(lastChunk.Content) == 0 {
				lastChunk = scm.ChunkedDocument{
					ChunkId:     uuid.New(),
					ChunkSerial: 1,
					Content:     line.Text,
					Embedding:   embedding,
				}
				continue
			}

			// chunk comparisions
			if len(lastChunk.Content)+len(line.Text) > MAX_CHUNK_SIZE {
				// create a new chunk
				// finalize the last chunk
				chunkedStrings = append(chunkedStrings, lastChunk)
				lastChunk = scm.ChunkedDocument{
					ChunkId:     uuid.New(),
					ChunkSerial: lastChunk.ChunkSerial + 1,
					Content:     line.Text,
					Embedding:   embedding,
				}
			} else {
				// cosine similarity check between lastChunk.Embedding and embedding
				// if cosine similarity is less than a threshold, create a new chunk
				// else append to the last chunk
				score, err := utils.CosineSimilarity(lastChunk.Embedding, embedding)
				if err != nil {
					return nil, err
				}
				if score < 0.8 {
					// create a new chunk
					chunkedStrings = append(chunkedStrings, lastChunk)
					lastChunk = scm.ChunkedDocument{
						ChunkId:     uuid.New(),
						ChunkSerial: lastChunk.ChunkSerial + 1,
						Content:     line.Text,
						Embedding:   embedding,
					}
				} else {
					// append to the last chunk
					// recompute the embedding for the last chunk
					lastChunk.Content += line.Text
					newEmbedding, err := ssc.eg.GetEmbeddings(lastChunk.Content)
					if err != nil {
						return nil, err
					}
					lastChunk.Embedding = newEmbedding
				}
			}
		}
	}
	chunkedStrings = append(chunkedStrings, lastChunk)
	return chunkedStrings, nil
}
