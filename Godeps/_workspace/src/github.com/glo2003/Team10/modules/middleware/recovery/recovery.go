package recovery

import "github.com/gin-gonic/gin"

// Middleware Initializes a new recovery handler middleware
func Middleware() gin.HandlerFunc {
	return gin.Recovery()
}
