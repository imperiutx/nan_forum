package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/imperiutx/nan_forum/db/sqlc"
	"github.com/lib/pq"
	"net/http"
)

type createCategoryRequest struct {
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
}

func (server *Server) createCategory(ctx *gin.Context) {
	var req createCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCategoryParams{
		Name:      req.Name,
		CreatedBy: req.CreatedBy,
	}

	category, err := server.store.CreateCategory(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}
