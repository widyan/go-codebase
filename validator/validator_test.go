package validator

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var usecaseService = CreateValidator(validator.New())

func TestValidateRequestWithGetBody(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString("{\"foo\":\"bar\", \"bar\":\"foo\"}"))
	c.Request.Header.Add("Content-Type", gin.MIMEXML) // set fake content-type

	var obj struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}

	usecaseService.ValidateRequestWithGetBody(c, &obj)
}

func TestValidateRequestWithGetBodyWithValidateStruct(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString("{\"foo\":\"bar\", \"bar\":\"\"}"))
	c.Request.Header.Add("Content-Type", gin.MIMEXML) // set fake content-type

	var obj struct {
		Foo string `json:"foo"`
		Bar string `json:"bar" validate:"required"`
	}

	usecaseService.ValidateRequestWithGetBody(c, &obj)
}

func TestValidateRequestWithGetBodyWithWringPayload(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("POST", "/", bytes.NewBufferString("{\"foo\"\"bar\"}"))
	c.Request.Header.Add("Content-Type", gin.MIMEXML) // set fake content-type

	var obj struct {
		Foo string `json:"foo"`
		Bar string `json:"bar" validate:"required"`
	}

	usecaseService.ValidateRequestWithGetBody(c, &obj)
}

func TestValidateRequest(t *testing.T) {
	type Obj struct {
		Foo string `json:"foo"`
		Bar string `json:"bar" validate:"required"`
	}

	var data = Obj{
		Foo: "bar",
		Bar: "foo",
	}

	usecaseService.ValidateRequest(data)
}

func TestValidateRequestWithAnyFieldRquired(t *testing.T) {
	type Obj struct {
		Foo string `json:"foo"`
		Bar string `json:"bar" validate:"required"`
	}

	var data = Obj{
		Foo: "bar",
	}

	usecaseService.ValidateRequest(data)
}
