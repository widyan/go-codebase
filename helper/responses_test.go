package helper

import (
	"codebase/go-codebase/model"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestJson(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	testResp := CreateCustomResponses("test")
	testResp.Json(c, http.StatusOK, "testing", "test")
}

func TestJsonWithErrorCode(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	testResp := CreateCustomResponses("test")
	testResp.JsonWithErrorCode(c, http.StatusBadRequest, InvalidToken)
}

func TestAbortWithStatusJSONAndInherited(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	testResp := CreateCustomResponses("test")
	testResp.AbortWithStatusJSONAndInherited(c, http.StatusOK, InvalidToken, "test", "testings")
}

func TestAbortWithStatusJSONAndErrorCode(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	testResp := CreateCustomResponses("test")
	testResp.AbortWithStatusJSONAndErrorCode(c, http.StatusOK, InvalidToken)
}

func TestStatusText(t *testing.T) {
	str := StatusText(http.StatusOK)
	assert.Equal(t, str, "OK")
}

func TestJsonWithCaptureError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	testResp := CreateCustomResponses("test")
	testResp.JsonWithCaptureError(c, fmt.Errorf("Testing Error"))
}

func TestJsonWithCaptureError1(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	cptrError := SetCaptureError(model.CaptureError{
		Type:      "capture error",
		HttpCode:  http.StatusBadRequest,
		ErrorCode: InvalidToken,
	})

	testResp := CreateCustomResponses("test")
	testResp.JsonWithCaptureError(c, cptrError)
}

func TestJsonWithCaptureError2(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	byteCapt, _ := json.Marshal(model.CaptureError{
		Type:      "",
		HttpCode:  http.StatusBadRequest,
		ErrorCode: InvalidToken,
	})

	testResp := CreateCustomResponses("test")
	testResp.JsonWithCaptureError(c, fmt.Errorf(string(byteCapt)))
}
