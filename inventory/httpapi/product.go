package httpapi

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type checkOrderRequest struct {
	Sku      string `json:"sku" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

func (server *Server) checkOrderHandler(ctx *gin.Context) {
	var req checkOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := server.store.CheckOrder(ctx, req.Sku, req.Quantity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

type getProductRequest struct {
	Sku string `uri:"sku" binding:"required"`
}

func (server *Server) getProductHandler(ctx *gin.Context) {
	var req getProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	product, err := server.store.GetProduct(ctx, req.Sku)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product.Quantity)

}
