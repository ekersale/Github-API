package projects

import (
	"net/http"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/glo2003/Team10/models"
    "github.com/glo2003/Team10/modules/github"
)

func IndexHandler(c *gin.Context) {
	projects, err := models.GetProjects()
	if (err != nil){
		error := models.Error{Message:err.Error()}
		c.JSON(http.StatusUnauthorized, error)
		log.Error("There was an error retrieving projects", err)
	} else if (len(projects) == 0) {
		c.JSON(http.StatusNoContent, "")
	} else {
		getprojects := models.GetProjectsFromRequest(projects)
		c.JSON(http.StatusOK, getprojects)
	}
}

// LoginHandler handle the login request
func LoginHandler(c *gin.Context) {
       github.SetAPIKey(c.Query("id"))
       client := github.NewClient()
       _, _, err := client.Users.Get("")
       if err != nil {
           c.JSON(401, gin.H{
            "status": "Bad token",
        })
        return;
       }
       c.JSON(200, gin.H{
           "status":  "ok",
       })
}