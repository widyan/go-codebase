package responses

import (
	"encoding/json"
	"strconv"

	"github.com/widyan/go-codebase/model"

	"github.com/gin-gonic/gin"
)

type GinResponses interface {
	Json(c *gin.Context, code int, data interface{}, message string)
	JsonWithErrorCode(c *gin.Context, code int, errorCode int)
	AbortWithStatusJSONAndInherited(c *gin.Context, code int, errorCode int, data interface{}, message string)
	AbortWithStatusJSONAndErrorCode(c *gin.Context, code int, errorCode int)
	JsonWithCaptureError(c *gin.Context, err error)
}

type GinResponsesImpl struct {
	ProjectName string `json:"project_name"`
}

func CreateCustomResponses(projectName string) GinResponses {
	return &GinResponsesImpl{ProjectName: projectName}
}

// Json is
func (r *GinResponsesImpl) Json(c *gin.Context, code int, data interface{}, message string) {
	if code > 399 {
		data = nil
	}
	c.JSON(code, model.Responses{
		Code:      code,
		Status:    StatusText(code),
		ErrorCode: 0,
		Message:   message,
		Data:      data,
	})
}

// Json iss
func (r *GinResponsesImpl) JsonWithErrorCode(c *gin.Context, code int, errorCode int) {
	c.JSON(code, model.Responses{
		Code:      code,
		Status:    StatusText(code),
		ErrorCode: errorCode,
		Message:   r.ProjectName + " - " + ErrorCodeText[errorCode] + " - " + strconv.Itoa(errorCode),
		Data:      nil,
	})
}

// AbortWithStatusJSON is
func (r *GinResponsesImpl) AbortWithStatusJSONAndInherited(c *gin.Context, code int, errorCode int, data interface{}, message string) {
	c.AbortWithStatusJSON(code, model.Responses{
		Code:      code,
		Status:    StatusText(code),
		ErrorCode: errorCode,
		Message:   message,
		Data:      data,
	})
}

// AbortWithStatusJSON is
func (r *GinResponsesImpl) AbortWithStatusJSONAndErrorCode(c *gin.Context, code int, errorCode int) {
	c.AbortWithStatusJSON(code, model.Responses{
		Code:      code,
		Status:    StatusText(code),
		ErrorCode: errorCode,
		Message:   r.ProjectName + " - " + ErrorCodeText[errorCode] + " - " + strconv.Itoa(errorCode),
		Data:      nil,
	})
}

// Json is
func (r *GinResponsesImpl) JsonWithCaptureError(c *gin.Context, err error) {
	var pureError error = err
	var capt model.CaptureError
	err = json.Unmarshal([]byte(err.Error()), &capt)
	if err != nil {
		c.JSON(500, model.Responses{
			Code:      500,
			Status:    StatusText(500),
			ErrorCode: 0,
			Message:   pureError.Error(),
		})
	} else {
		if capt.Type == "capture error" {
			c.JSON(capt.HttpCode, model.Responses{
				Code:      capt.HttpCode,
				Status:    StatusText(capt.HttpCode),
				ErrorCode: capt.ErrorCode,
				Message:   r.ProjectName + " - " + ErrorCodeText[capt.ErrorCode] + " - " + strconv.Itoa(capt.ErrorCode),
			})
		} else {
			c.JSON(500, model.Responses{
				Code:      500,
				Status:    StatusText(500),
				ErrorCode: 0,
				Message:   pureError.Error(),
			})
		}
	}
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
