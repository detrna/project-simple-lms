package file

type CreateFileRequest struct {
	Name        string  `json:"name" binding:"required"`
	FileURL     string  `json:"fileUrl" binding:"required,url"`
	ContentType string  `json:"contentType" binding:"required"`
	Size        float64 `json:"size" binding:"required"`
}
