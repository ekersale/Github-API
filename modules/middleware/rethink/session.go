package rethink

import (
	rethink "github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
)

const rethinkSessionKey string = "rethinkSession"

// Middleware returns a gin middleware that sets saves the rethinkDB Sessionthe HTTP Header `X-Git-Commit`
// to the latest git commit, it also sets it in the context so you can retrieve
// it with #FromContext
func Middleware(s *rethink.Session) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(rethinkSessionKey, s)
		c.Next()
	}
}

// FromContext return the gitCommit of the current request
func FromContext(ctx gin.Context) (*rethink.Session, bool) {
	session, ok := ctx.Value(rethinkSessionKey).(*rethink.Session)
	return session, ok
}
