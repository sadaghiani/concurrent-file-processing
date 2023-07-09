package information

import (
	"github.com/sadaghiani/concurrent-file-processing/internal/models/information"
)

type Information struct {
	information.IInformation
}

func NewInformation(i information.IInformation) *Information {
	return &Information{
		i,
	}
}
