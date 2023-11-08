package somepackage

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Issue(c *gin.Context) {
	c.Redirect(http.StatusFound, "http://google.com")
}
