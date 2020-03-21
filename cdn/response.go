package cdn

// UploadUUIDResponse represents a response for an upload with an auto generated filename
type UploadUUIDResponse struct {
	FileName string `json:"file_name"`
}

type errorResponse struct {
	Message string `json:"message"`
}

// NewUploadUUIDResponse creates a new UploadUUIDResponse instance
func NewUploadUUIDResponse(fileName string) *UploadUUIDResponse {
	return &UploadUUIDResponse{FileName: fileName}
}
