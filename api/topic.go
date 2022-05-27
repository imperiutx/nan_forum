package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "github.com/imperiutx/nan_forum/db/sqlc"
	"github.com/imperiutx/nan_forum/utils"
	"github.com/lib/pq"
	"net/http"
)

type createTopicRequest struct {
	CategoryID int64  `json:"category_id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	CreatedBy  string `json:"created_by"`
}

func (server *Server) createTopic(ctx *gin.Context) {
	var req createTopicRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTopicParams{
		CategoryID: req.CategoryID,
		Title:      req.Title,
		Body:       req.Body,
		CreatedBy:  req.CreatedBy,
	}

	topic, err := server.store.CreateTopic(ctx, arg)
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

	ctx.JSON(http.StatusOK, topic)
}

type getTopicRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type getTopicResponse struct {
	CategoryID    int64        `json:"category_id"`
	Title         string       `json:"title"`
	Body          string       `json:"body"`
	CreatedBy     string       `json:"created_by"`
	Points        int64        `json:"points"`
	TimeAgo       string       `json:"time_ago"`
	CommentsCount int64        `json:"comments_count"`
	Comments      []db.Comment `json:"comments"`
}

func (server *Server) getTopic(ctx *gin.Context) {
	var req getTopicRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	topic, err := server.store.GetTopicByID(ctx, int32(req.ID))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	count, err := server.store.CountCommentsByTopicID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	comments, err := server.store.ListCommentsByTopicID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := getTopicResponse{
		CategoryID:    topic.CategoryID,
		Title:         topic.Title,
		Body:          topic.Body,
		CreatedBy:     topic.CreatedBy,
		Points:        topic.Points,
		TimeAgo:       utils.TimeSince(topic.CreatedAt),
		CommentsCount: count,
		Comments:      comments,
	}

	ctx.JSON(http.StatusOK, rsp)
}

type listTopicRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listTopics(ctx *gin.Context) {
	//var req listTopicRequest
	//if err := ctx.ShouldBindQuery(&req); err != nil {
	//	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	//	return
	//}

	topics, err := server.store.ListTopics(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rspTopics := make([]getTopicResponse, len(topics))

	for _, tp := range topics {
		count, err := server.store.CountCommentsByTopicID(ctx, int64(tp.ID))
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}

			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		rsp := getTopicResponse{
			CategoryID:    tp.CategoryID,
			Title:         tp.Title,
			Body:          tp.Body,
			CreatedBy:     tp.CreatedBy,
			Points:        tp.Points,
			TimeAgo:       utils.TimeSince(tp.CreatedAt),
			CommentsCount: count,
		}
		rspTopics = append(rspTopics, rsp)
	}

	ctx.JSON(http.StatusOK, rspTopics)
}
