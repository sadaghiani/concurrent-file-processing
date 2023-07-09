package data

import (
	"github.com/sadaghiani/concurrent-file-processing/internal/models/data"
)

type Data struct {
	data.IDeta
}

func NewData(d data.IDeta) *Data {
	return &Data{
		d,
	}
}
