package file

import (
	"context"
	"mime/multipart"

	"github.com/sadaghiani/concurrent-file-processing/internal/repository"
)

type IFile interface {
	Upload(ctx context.Context, ff *multipart.FileHeader) error
}

type File struct {
	*repository.Repository
}

func NewFile(repo *repository.Repository) *File {
	return &File{
		Repository: repo,
	}
}
