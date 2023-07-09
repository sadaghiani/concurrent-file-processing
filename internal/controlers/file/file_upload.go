package file

import (
	"context"
	"net/http"

	"github.com/sadaghiani/concurrent-file-processing/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Summary upload csv file
// @Description upload csv file
// @Tags file
// @Param file formData  file true " "
// @Accept mpfd
// @Produce json
// @Success 200 {object} utils.GeneralResponse
// @Success 400 {object} utils.GeneralResponse
// @Router /file/upload [post]
func (f *File) Upload(c *gin.Context) {

	newfile, err := c.FormFile("file")
	if err != nil {
		utils.CreateAbortResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	if err := f.IFile.Upload(context.Background(), newfile); err != nil {
		utils.CreateAbortResponse(c, http.StatusInternalServerError, err)
		return
	}
	utils.CreateCompleteResponse(c, http.StatusOK, "upload was successful")
}
