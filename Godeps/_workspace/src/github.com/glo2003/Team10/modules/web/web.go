package web

import (
	"github.com/glo2003/Team10/modules/middleware/logger"
	"github.com/glo2003/Team10/modules/middleware/recovery"
	"github.com/glo2003/Team10/routers/home"
	"github.com/glo2003/Team10/routers/projects"
    "github.com/glo2003/Team10/modules/github"
	"github.com/gin-gonic/gin"
    "fmt"
)

type Server interface {
	Run()
}

type server struct {
	r *gin.Engine
}

func NewServer() Server {
	r := gin.New()
	r.Use(recovery.Middleware(),
		logger.Middleware())
        
    r.LoadHTMLFiles("templates/template.html")
	r.Use(Cors())
	r.GET("/", home.IndexHandler)
	r.GET("/ping", home.PingHandler)
	r.GET("/projects", projects.IndexHandler)
    r.OPTIONS("/projects", func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST,PUT,GET")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")
        c.Next()
    })
    client := github.NewClient()

    if (client == nil) {
        fmt.Println("client not connected")
    }
	return &server{r}
}

func (s *server) Run() {
	s.r.Run(":3456")
}

func Cors() gin.HandlerFunc {
 return func(c *gin.Context) {
 c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
 c.Next()
 }
}
