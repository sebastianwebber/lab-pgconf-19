package web

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

var (
	router = gin.Default()
	dbConn *pg.DB
)

// GetRouter returns a routed gin Engine
func GetRouter(db *pg.DB) *gin.Engine {

	if db != nil {
		dbConn = db
	}

	v1 := router.Group("/api/v1/")
	{
		v1.POST("/load", loadData)
	}

	return router
}
