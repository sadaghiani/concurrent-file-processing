package file

import (
	"github.com/sadaghiani/concurrent-file-processing/internal/models/file"
)

type File struct {
	file.IFile
}

func NewFile(f file.IFile) *File {
	return &File{
		f,
	}
}
