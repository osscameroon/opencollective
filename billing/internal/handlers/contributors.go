package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetContributors returns the list of contributors
// GET /contributors
// it accepts an organization name as a query parameter
func GetContributors(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "a list of contributors"})
}
