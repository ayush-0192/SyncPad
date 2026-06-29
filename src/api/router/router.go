package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ayush-0192/SyncPad/src/api/document"
)

func Setup(docHandler *document.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "SyncPad API is running",
		})
	}) 

	r.POST("/create-document", docHandler.CreateDoc)
	r.GET("/get-document/:id", docHandler.GetDocByid)
	r.GET("/get-document-title-list", docHandler.GetAllDocTitleList)
	r.PATCH("/update-doc",docHandler.GetAllDocTitleList)
	return r
}