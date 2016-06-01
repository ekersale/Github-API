package home

import (
	"net/http"
    "github.com/glo2003/Team10/models"
	"github.com/gin-gonic/gin"
    log "github.com/Sirupsen/logrus"
    "fmt"
)

func PingHandler(c *gin.Context) {
	c.String(http.StatusFound, "pong")
}

func IndexHandler(c *gin.Context) {
    objects, err := models.GetProjects()
    fmt.Printf("%d\n",len(objects))
	if err != nil {
		log.Error("There was an error retrieving projects", err)
	}
    c.HTML(http.StatusOK, "template.html", gin.H {
        "objects" : objects,
    })
}
