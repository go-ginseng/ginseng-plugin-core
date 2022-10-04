package auth_session

import "github.com/go-ginseng/ginseng"

var publicPaths = make(map[string]bool)

// AddPublicPath add the path to the public path list
func AddPublicPath(method string, path string) {
	key := method + ":" + path
	publicPaths[key] = true
}

var loggedInPaths = make(map[string]bool)

// AddLoggedInPath add the path to the logged in path list
func AddLoggedInPath(method string, path string) {
	key := method + ":" + path
	loggedInPaths[key] = true
}

var securityMiddlewareHandlers = make([]SecurityHandler, 0)

// AddSecurityMiddlewareHandler add a security middleware handler
func AddSecurityMiddlewareHandler(handler SecurityHandler) {
	securityMiddlewareHandlers = append(securityMiddlewareHandlers, handler)
}

// Install then security middleware to the engine
func Install(e *ginseng.Engine) {
	e.Gin().Use(buildSecurityMiddleware(e, securityMiddlewareHandlers))
}
