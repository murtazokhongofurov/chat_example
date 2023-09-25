package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kafka_example/api-gateway/api/models"
)


func (h *handlerV1) TestGet(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/user", h.GetUser)

	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatalf("error send request: %v", err)
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	//assert.Equal(t, http.StatusCreated, w.Code)
}

func (h *handlerV1) TestPost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/user", )
	user := models.UserRequest{
		
	}
	userJson, err := json.Marshal(&user)
	if err != nil {
		t.Fatalf("error marshaling: %v", err)
	}

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer([]byte(userJson)))
	if err != nil {
		t.Fatalf("error send request: %v", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

}
