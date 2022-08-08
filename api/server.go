package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/imperiutx/nan_forum/db/sqlc"
	"github.com/imperiutx/nan_forum/middleware"
	"github.com/imperiutx/nan_forum/token"
	"github.com/imperiutx/nan_forum/utils"
)

// Server serves HTTP requests for our forum service.
type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.PasetoMaker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: *tokenMaker,
	}

	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	v.RegisterValidation("currency", validCurrency)
	//}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.Use(middleware.Recover())

	v1 := router.Group("v1")

	categories := v1.Group("/categories").Use(authMiddleware(server.tokenMaker))
	{
		categories.POST("", server.createCategory)

	}

	topics := v1.Group("/topics").Use(authMiddleware(server.tokenMaker))
	{
		topics.POST("", server.createTopic)
		topics.GET("", server.listTopics)
		topics.GET("/:id", server.getTopic)
	}

	comments := v1.Group("/comments").Use(authMiddleware(server.tokenMaker))
	{
		comments.POST("", server.createComment)

	}

	users := v1.Group("/users")
	{
		users.POST("", server.createUser)
		users.POST("/login", server.loginUser)

	}

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
