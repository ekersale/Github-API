package web

import (
	"github.com/ekersale/Github-API/modules/middleware/logger"
	"github.com/ekersale/Github-API/modules/middleware/recovery"
	"github.com/ekersale/Github-API/routers/home"
	"github.com/ekersale/Github-API/routers/projects"
	"github.com/gin-gonic/gin"
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
        
    r.LoadHTMLFiles("templates/template.html", "templates/connect.html")
    r.Static("/css", "templates/css")
    r.Static("/js", "templates/js")
	r.Use(Cors())
    r.GET("/", home.LoginHandler)
	r.GET("/index", home.IndexHandler)
	r.GET("/ping", home.PingHandler)
	r.GET("/projects", projects.IndexHandler)
    r.POST("/login", projects.LoginHandler)
    r.OPTIONS("/projects", func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST,PUT,GET")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")
        c.Next()
    })
	return &server{r}
}

func (s *server) Run() {
	s.r.Run()
}

func Cors() gin.HandlerFunc {
 return func(c *gin.Context) {
 c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
 c.Next()
 }
}
