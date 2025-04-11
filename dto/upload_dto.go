package dto

import "mime/multipart"

type UploadFileInput struct {
	FileHeader *multipart.FileHeader
	File       multipart.File
}

type UploadResponse struct {
	FileURL  string
	FileName string
}
