package domain

import "mime/multipart"

type File struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
	Base64   string `json:"base64"`
}

type IFileRepository interface {
	Save(file multipart.FileHeader) (string, error)
	GetFiles(paths []string) ([]File, error)
}
