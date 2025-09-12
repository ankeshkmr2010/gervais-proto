package dtos

type IngestorRequest struct {
	
}

type FileIngestionRequest struct {
	FileType string `json:"file_type"`
	FilePath string `json:"file_path"`
}
