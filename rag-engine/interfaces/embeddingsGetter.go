package interfaces

//go:generate mockery
type EmbeddingsGetter interface {
	GetEmbeddings(text string) ([]float64, error)
}
