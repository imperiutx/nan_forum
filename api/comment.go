package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/imperiutx/nan_forum/db/sqlc"
	"github.com/lib/pq"
	"net/http"
)

type createCommentRequest struct {
	TopicID   int64  `json:"topic_id"`
	Body      string `json:"body"`
	CreatedBy string `json:"created_by"`
}

func (server *Server) createComment(ctx *gin.Context) {
	var req createCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCommentParams{
		TopicID:   req.TopicID,
		Body:      req.Body,
		CreatedBy: req.CreatedBy,
	}

	comment, err := server.store.CreateComment(ctx, arg)
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

	ctx.JSON(http.StatusOK, comment)
}

type listCommentsRequest struct {
	TopicID int64 `form:"topic_id" binding:"required,min=1"`
	//PageID   int32 `form:"page_id" binding:"required,min=1"`
	//PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type GetCommentResponse struct {
	ID        int64  `json:"id"`
	CreatedBy string `json:"created_by"`
	Body      string `json:"body"`
	Points    int64  `json:"points"`
	TimeAgo   string `json:"time_ago"`
	//RepliesCount int64 `json:"replies_count"`
}

func (server *Server) listComments(ctx *gin.Context) {
	var req listCommentsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	topics, err := server.store.ListCommentsByTopicID(ctx, req.TopicID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, topics)
}
