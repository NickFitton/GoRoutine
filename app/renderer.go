package app

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

func HtmlResponse(comp templ.Component) func(cont *gin.Context) {
	return func(cont *gin.Context) {
		cont.HTML(http.StatusOK, "", comp)
	}
}

type TemplRender struct {
	Code int
	Data templ.Component
}

func (t TemplRender) Render(w http.ResponseWriter) error {
	t.WriteContentType(w)
	w.WriteHeader(t.Code)
	if t.Data != nil {
		return t.Data.Render(context.Background(), w)
	}
	return nil
}

func (t TemplRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func (t *TemplRender) Instance(_ string, data interface{}) render.Render {
	if templData, ok := data.(templ.Component); ok {
		return &TemplRender{
			Code: http.StatusOK,
			Data: templData,
		}
	}
	return nil
}
