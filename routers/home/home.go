package home

import (
	"net/http"
    	"github.com/ekersale/Github-API/models"
	"github.com/gin-gonic/gin"
    log "github.com/Sirupsen/logrus"
)

func PingHandler(c *gin.Context) {
	c.String(http.StatusFound, "pong")
}

func IndexHandler(c *gin.Context) {
    objects, err := models.GetProjects()
	if err != nil {
		log.Error("There was an error retrieving projects", err)
	}
    c.HTML(http.StatusOK, "template.html", gin.H {
        "objects" : objects,
    })
}

func LoginHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "connect.html", gin.H {})
}
