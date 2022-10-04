package auth_session

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ginseng/ginseng"
)

// The security middleware is composed by many security middleware handlers
// If the handler returns 0, the next handler will be executed
// If the handler returns 200, the rest of the handlers will be skipped
// If the handler returns 401 or 403, the rest of the handlers will be skipped and the request will be aborted
type SecurityHandler func(e *ginseng.Engine, ctx *gin.Context) int

func DeleteExpiredSessionHandler(e *ginseng.Engine, ctx *gin.Context) int {
	DeleteExpiredSession()
	return 0
}

func SkipOptionRequestHandler(e *ginseng.Engine, ctx *gin.Context) int {
	if ctx.Request.Method == "OPTIONS" {
		return 200
	}
	return 0
}

func SkipPublicPathHandler(e *ginseng.Engine, ctx *gin.Context) int {
	method := ctx.Request.Method
	pathname := GetRequestPath(ctx)
	_, ok := publicPaths[method+":"+pathname]
	if ok {
		return 200
	}
	return 0
}

func SkipLoggedInPathHandler(e *ginseng.Engine, ctx *gin.Context) int {
	s := CurrentSession(ctx)
	if s == nil {
		return 401
	}

	method := ctx.Request.Method
	pathname := GetRequestPath(ctx)
	_, ok := loggedInPaths[method+":"+pathname]
	if ok {
		return 200
	}
	return 0
}

func buildSecurityMiddleware(e *ginseng.Engine, additionalHandlers []SecurityHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, handler := range additionalHandlers {
			status := handler(e, ctx)
			switch status {
			case 0:
				continue
			case 200:
				ctx.Next()
				return
			case 401:
				ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
				return
			case 403:
				ctx.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
				return
			}
		}
	}
}

func GetRequestPath(ctx *gin.Context) string {
	fullPath := ctx.FullPath()
	if fullPath == "" {
		fullPath = ctx.Request.URL.Path
	}
	return fullPath
}
