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

	ctx.JSON(http.StatusOK, topic.ID)
}

type getTopicRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type getTopicResponse struct {
	CategoryID    *int64        `json:"category_id,omitempty"`
	Title         *string       `json:"title,omitempty"`
	Body          *string       `json:"body,omitempty"`
	CreatedBy     *string       `json:"created_by,omitempty"`
	Points        *int64        `json:"points,omitempty"`
	TimeAgo       *string       `json:"time_ago,omitempty"`
	CommentsCount *int64        `json:"comments_count,omitempty"`
	Comments      *[]db.Comment `json:"comments,omitempty"`
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

	ta := utils.TimeSince(topic.CreatedAt)

	rsp := getTopicResponse{
		CategoryID:    &topic.CategoryID,
		Title:         &topic.Title,
		Body:          &topic.Body,
		CreatedBy:     &topic.CreatedBy,
		Points:        &topic.Points,
		TimeAgo:       &ta,
		CommentsCount: &count,
		Comments:      &comments,
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

	for i, tp := range topics {
		count, err := server.store.CountCommentsByTopicID(ctx, int64(tp.ID))
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}

			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		catID := tp.CategoryID
		tt := tp.Title
		tb := tp.Body
		ta := utils.TimeSince(tp.CreatedAt)
		tc := tp.CreatedBy
		tp := tp.Points

		rspTopics[i] = getTopicResponse{
			CategoryID:    &catID,
			Title:         &tt,
			Body:          &tb,
			CreatedBy:     &tc,
			Points:        &tp,
			TimeAgo:       &ta,
			CommentsCount: &count,
		}
	}

	ctx.JSON(http.StatusOK, rspTopics)
}
