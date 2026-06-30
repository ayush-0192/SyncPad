package document

import (
	 "github.com/gin-gonic/gin"
	 "fmt"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler {
		service : service,
	}
}

func(h *Handler) CreateDoc(c *gin.Context) {
	var doc Document

	if err := c.ShouldBindJSON(&doc); err != nil {
		c.JSON(400, gin.H {
			"error" : err.Error(),
		})
		fmt.Printf("Error in binding json: %v\n", err)
		return
	}

	err := h.service.saveDocument(&doc)

	if err != nil {
		c.JSON(500, gin.H {
			"error" : err.Error(),
		})
		return
	}
	c.JSON(201,doc)
}

func (h *Handler) GetDocByid(c *gin.Context) {
	id := c.Param("id")

	doc, err := h.service.getDocumentById(id)

	if err != nil {
		c.JSON(500, gin.H {
			"error" : err.Error(),
		})
		return
	}

	if doc == nil {
		c.JSON(404, gin.H {
			"error" : "Document not found",
		})
		return
	}

	c.JSON(200, doc)
}

func(h *Handler) GetAllDocTitleList(c *gin.Context) {
	docTitleList, err := h.service.getDocumentTitleList()

	if err != nil {
		c.JSON(500,gin.H {
			"error": err.Error(),
		})
		return
	}
	c.JSON(200,docTitleList)
}

func (h *Handler) UpdateDoc(c *gin.Context) {
	 id := c.Param("id")
	 
	 var req UserDocumentUpdateRequest

	 if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400,gin.H {
			"error": err.Error(),
		})

		return
	 }

	 err := h.service.updateDocumentContentAndTitle(id,&req)

	 if err != nil {
		c.JSON(500,gin.H {
			"error" : err.Error(),
		})
	 }

	 c.JSON(200, gin.H{
        "message": "user updated successfully",
    })
}