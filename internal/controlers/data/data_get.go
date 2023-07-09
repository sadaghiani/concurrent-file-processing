package data

import (
	"context"
	"net/http"

	"github.com/sadaghiani/concurrent-file-processing/internal/utils"

	"github.com/gin-gonic/gin"
)

type Query struct {
	Page  int `json:"page" binding:"required"`
	Limit int `json:"limit" binding:"required"`
}

// @Summary csv data
// @Description csv data
// @Tags data
// @Param Page query integer true " "
// @Param Limit query integer true " "
// @Accept json
// @Produce json
// @Success 200 {object} utils.GeneralResponse
// @Success 400 {object} utils.GeneralResponse
// @Router /data/ [get]
func (d *Data) Get(c *gin.Context) {

	var q Query
	if err := c.BindQuery(&q); err != nil {
		utils.CreateAbortResponse(c, http.StatusPreconditionRequired, err)
		return
	}

	data, err := d.IDeta.Get(context.Background(), q.Page, q.Limit)
	if err != nil {
		utils.CreateAbortResponse(c, http.StatusInternalServerError, err)
		return
	}
	utils.CreateCompleteResponse(c, http.StatusOK, data)
}
