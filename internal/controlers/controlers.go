package controlers

import (
	"github.com/sadaghiani/concurrent-file-processing/internal/controlers/data"
	"github.com/sadaghiani/concurrent-file-processing/internal/controlers/file"
	"github.com/sadaghiani/concurrent-file-processing/internal/controlers/information"
	"github.com/sadaghiani/concurrent-file-processing/internal/models"
)

type Controlers struct {
	Information *information.Information
	File        *file.File
	Data        *data.Data
}

func NewControlers(models *models.Models) *Controlers {
	return &Controlers{
		Information: information.NewInformation(models),
		File:        file.NewFile(models),
		Data:        data.NewData(models),
	}
}
