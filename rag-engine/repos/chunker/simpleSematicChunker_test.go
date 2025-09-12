package chunker

import (
	"rag-engine/app/interfaces"
	"testing"

	"github.com/jdkato/prose/v2"
	"github.com/stretchr/testify/suite"
)

type ChunkerTestSuite struct {
	suite.Suite
	chunker interfaces.Chunker
	eg      *interfaces.MockEmbeddingsGetter
}

func (suite *ChunkerTestSuite) SetupTest() {
	suite.eg = &interfaces.MockEmbeddingsGetter{}
	suite.chunker = NewSimpleSemanticChunker(suite.eg)
}

func (suite *ChunkerTestSuite) TestProcessDocument() {

	suite.eg.EXPECT().GetEmbeddings("Hey there.").Return([]float64{0.1, 0.2, 0.3}, nil).Once()
	suite.eg.EXPECT().GetEmbeddings("My email id is ankesh.kmr@gmail.com.").Return([]float64{0.4, 0.5, 0.6}, nil).Once()
	suite.eg.EXPECT().GetEmbeddings("Hey there.My email id is ankesh.kmr@gmail.com.").Return([]float64{0.7, 0.8, 0.9}, nil).Once()
	doc, err := prose.NewDocument("Hey there. My email id is ankesh.kmr@gmail.com.")
	if err != nil {
		suite.Fail("Failed to create document")
	}
	docs := []prose.Document{*doc}
	chunks, err := suite.chunker.ProcessDocument(docs, nil)
	suite.NoError(err)
	suite.NotNil(chunks)
	suite.Greater(len(chunks), 0)
}

func (suite *ChunkerTestSuite) TearDownTest() {
}

func TestChunkerSuite(t *testing.T) {
	suite.Run(t, new(ChunkerTestSuite))
}
