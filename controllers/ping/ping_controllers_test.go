package ping

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {

	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)
	Health(context)

	if w.Code != 200 {
		t.Error("status code should be 200")
	}
}
