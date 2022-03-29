package helper

import (
	"bytes"
	"codebase/go-codebase/model"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmhttp"
)

// CallAPI is

var (
	Buf    bytes.Buffer
	Logger = log.New(&Buf, "logger: ", log.Lshortfile)
)

func CallAPI(ctx context.Context, logger *logrus.Logger, url, method string, payload interface{}, header []model.Header) (body []byte, err error) {
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

	client := apmhttp.WrapClient(http.DefaultClient)
	res, err := client.Do(req.WithContext(ctx))
	if err != nil {
		// apm.CaptureError(ctx, err).Send()
		// http.Error(w, "failed to query backend", 500)
		return
	}
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	return
}

func CallAPIFormData(logger *logrus.Logger, url, method string, formData []model.FormData, headers []model.Header) (body []byte, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for _, element := range formData {
		_ = writer.WriteField(element.Key, element.Value)
	}
	err = writer.Close()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())
	for _, e := range headers {
		req.Header.Add(e.Key, e.Value)
	}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	return
}
