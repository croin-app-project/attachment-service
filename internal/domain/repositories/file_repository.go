package repositories

import (
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/croin-app-project/attachment-service/internal/domain"
	"github.com/google/uuid"
)

type FileRepository struct {
}

func NewFileRepository() domain.IFileRepository {
	return &FileRepository{}
}

func (r *FileRepository) Save(file multipart.FileHeader) (string, error) {
	workPath := "/AttachFiles"

	// Check OS for setting work path
	if os.Getenv("OS") == "Windows_NT" {
		workPath = "D:/AttachFiles"
	}

	// Ensure base workPath exists
	if err := os.MkdirAll(workPath, os.ModePerm); err != nil {
		return "", err
	}

	// Prepare directory structure based on current year
	now := time.Now()
	yearPath := filepath.Join(workPath, now.Format("2006"))
	if err := os.MkdirAll(yearPath, os.ModePerm); err != nil {
		return "", err
	}

	// Prepare directory structure based on current date
	datePath := filepath.Join(yearPath, now.Format("01"), now.Format("02"))
	if err := os.MkdirAll(datePath, os.ModePerm); err != nil {
		return "", err
	}

	_file, err := file.Open()
	if err != nil {
		return "", err
	}
	defer _file.Close()
	// Generate file name with timestamp
	fileExt := filepath.Ext(file.Filename)
	fileName := strings.TrimSuffix(file.Filename, fileExt) + "-" + now.Format("150405000") + fileExt
	fmt.Println(fileName)
	// Create full file path
	fullPath := filepath.Join(datePath, fileName)
	fmt.Println(fullPath)
	// Create the file
	out, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Copy the file content
	_, err = io.Copy(out, _file)
	if err != nil {
		return "", err
	}

	return fullPath, nil
}

func (r *FileRepository) GetFiles(paths []string) ([]domain.File, error) {
	var files []domain.File

	for _, path := range paths {
		fileBytes, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		fileName := filepath.Base(path)
		fileID := strings.ReplaceAll(strings.ToLower(uuid.New().String()), "-", "")

		files = append(files, domain.File{
			ID:       fileID,
			Filename: fileName,
			Base64:   base64.StdEncoding.EncodeToString(fileBytes),
		})
	}

	return files, nil
}

func (r *FileRepository) DeleteFiles(paths []string) error {
	for _, path := range paths {
		os.Remove(path)
	}
	return nil
}
