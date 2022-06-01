package categorytransport

import (
	"github.com/gin-gonic/gin"
	"github.com/imperiutx/nan_forum/appctx"
	"github.com/imperiutx/nan_forum/common"
	db "github.com/imperiutx/nan_forum/db/sqlc"
	"github.com/imperiutx/nan_forum/modules/category/categorybusiness"
)

func CreateCategory() func(c *gin.Context) {
	return func(c *gin.Context) {
		var category db.CreateCategoryParams

		if err := c.ShouldBindJSON(&category); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		ac := appctx.AppContext
		daba := ac.GetDBConnection()

		store := db.NewStore(daba)

		bizcategory := categorybusiness.NewCreateCategoryBiz(store)

		if err :=
	}
}
