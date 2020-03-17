package cdn

// UploadUUIDResponse represents a response for an upload with an auto generated filename
type UploadUUIDResponse struct {
	FileName string `json:"file_name"`
}

// NewUploadUUIDResponse creates a new UploadUUIDResponse instance
func NewUploadUUIDResponse(fileName string) *UploadUUIDResponse {
	return &UploadUUIDResponse{FileName: fileName}
}
