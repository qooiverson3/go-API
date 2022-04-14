package handler

import (
	"ces-api/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InstanceHandler struct {
	Svc model.Service
}

func NewInstanceHandler(s model.Service) *InstanceHandler {
	return &InstanceHandler{
		Svc: s,
	}
}

// Router path
func (h *InstanceHandler) Router(e *gin.Engine) {
	e.GET("/api/v1/instance", h.GetInstanceList)
}

// GetInstance _
func (h *InstanceHandler) GetInstanceList(e *gin.Context) {
	header := &model.XHeader{}
	err := e.BindHeader(header)

	if err != nil {
		e.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"state": false,
			"message": []string{
				"wrong token.",
			},
		})
		return
	}

	token := e.GetHeader("token")
	if token != "4BEE7FAC2C9FE56685E61637248EF2A3" {
		e.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"state": false,
			"message": []string{
				"illegal token.",
			},
		})
		return
	}

	dept := e.Query("dept")
	page := e.Query("page")

	data := h.Svc.GetInstanceList(dept, page)
	e.JSON(http.StatusOK, gin.H{
		"state": true,
		"data":  data,
	})
}
