package information

import (
	"context"
	"net/http"

	"github.com/sadaghiani/concurrent-file-processing/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Summary healthcheck
// @Description healthcheck
// @Tags information
// @Accept json
// @Produce json
// @Success 200 {object} utils.GeneralResponse
// @Success 400 {object} utils.GeneralResponse
// @Router /health [get]
func (i *Information) Health(c *gin.Context) {

	if err := i.IInformation.Health(context.Background()); err != nil {
		utils.CreateAbortResponse(c, http.StatusInternalServerError, err)
		return
	}
	utils.CreateCompleteResponse(c, http.StatusOK, "service is healthy")
}
