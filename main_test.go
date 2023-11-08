package somepackage_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	somepackage "example.com/m"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	somepackage.Issue(c)
	assert.Equal(t, http.StatusFound, w.Code)
	assert.NotEmpty(t, w.Result().Cookies())
}
