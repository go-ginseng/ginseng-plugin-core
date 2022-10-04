package static

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func htmlMiddleware(ctx *gin.Context) {
	if ctx.Request.URL.Path != "/" &&
		!strings.Contains(ctx.Request.URL.Path, "/api/") &&
		!strings.Contains(ctx.Request.URL.Path, "/static/") {
		ctx.Request.URL.Path = ctx.Request.URL.Path + ".html"
	}
}
