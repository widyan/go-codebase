package helper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/widyan/go-codebase/model"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
)

type ToolsAPI struct {
	Logger *logrus.Logger
}

func CreateToolsAPI(logger *logrus.Logger) API_Interface {
	return &ToolsAPI{logger}
}

type API_Interface interface {
	CallAPI(ctx context.Context, url, method string, payload interface{}, header []model.Header) (body []byte, err error)
	CallAPIFormData(ctx context.Context, url, method string, formData []model.FormData, headers []model.Header) (body []byte, err error)
	SendToTelegram(url, method, tokenBOT, chatID, text string, IsContainArstik bool)
}

// CallAPI is
func (t *ToolsAPI) CallAPI(ctx context.Context, url, method string, payload interface{}, header []model.Header) (body []byte, err error) {
	// var res *http.Response
	body, err = json.Marshal(payload)
	if err != nil {
		return
	}

	var req *http.Request
	// var w http.ResponseWriter
	// client := &http.Client{}

	req, err = http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	req.Header.Add("content-type", "application/json")
	for _, e := range header {
		req.Header.Add(e.Key, e.Value)
	}

	client := httptrace.WrapClient(http.DefaultClient)
	res, err := client.Do(req.WithContext(ctx))
	if err != nil {
		// apm.CaptureError(ctx, err).Send()
		// http.Error(w, "failed to query backend", 500)
		return
	}
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Logger.Error(err.Error())
		return
	}

	return
}

func (t *ToolsAPI) CallAPIFormData(ctx context.Context, url, method string, formData []model.FormData, headers []model.Header) (body []byte, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for _, element := range formData {
		_ = writer.WriteField(element.Key, element.Value)
	}
	err = writer.Close()
	if err != nil {
		t.Logger.Error(err.Error())
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		t.Logger.Error(err.Error())
		return
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())
	for _, e := range headers {
		req.Header.Add(e.Key, e.Value)
	}
	res, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Logger.Error(err.Error())
		return
	}

	if res.StatusCode > 399 {
		err = fmt.Errorf(string(body))
		t.Logger.Error(err.Error())
		return
	}

	return
}

func (t *ToolsAPI) SendToTelegram(url, method, tokenBOT, chatID, text string, IsContainArstik bool) {
	client := &http.Client{}
	payload := strings.NewReader("chat_id=" + chatID + "&text=" + text + "&parse_mode=Markdown")
	if IsContainArstik {
		payload = strings.NewReader("chat_id=" + chatID + "&text=" + text)
	}
	req, err := http.NewRequest("POST", url+tokenBOT+"/sendMessage", payload)
	if err != nil {
		t.Logger.Error(err.Error())
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		t.Logger.Error(err.Error())
		return
	}
	defer res.Body.Close()
}
