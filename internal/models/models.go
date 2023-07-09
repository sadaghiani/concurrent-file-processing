package models

import (
	"github.com/sadaghiani/concurrent-file-processing/internal/models/data"
	"github.com/sadaghiani/concurrent-file-processing/internal/models/file"
	"github.com/sadaghiani/concurrent-file-processing/internal/models/information"
	"github.com/sadaghiani/concurrent-file-processing/internal/repository"
)

type Models struct {
	data.IDeta
	file.IFile
	information.IInformation
}

func NewModels(repo *repository.Repository) *Models {

	return &Models{
		data.NewData(repo),
		file.NewFile(repo),
		information.NewInformation(repo),
	}
}
