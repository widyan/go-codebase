package worker

import (
	"encoding/json"
	"codebase/go-codebase/helper"
	"os"

	rdsWorker "github.com/garyburd/redigo/redis"
)

type Worker struct {
	Logger *helper.CustomLogger
}

func CreateWorker(logger *helper.CustomLogger) *Worker {
	return &Worker{logger}
}

type job struct {
	Class string        `json:"class"`
	Args  []interface{} `json:"args"`
}

type Client struct {
	conn rdsWorker.Conn
}

func (b Worker) SendOtpToEmailUser(fullName, email, otp string) (err error) {
	c, err := Dial(os.Getenv("REDIS"))
	if err != nil {
		b.Logger.Error(err.Error())
		return
	}

	body := map[string]interface{}{
		"task":   "Email",
		"action": "sendotp",
		"param": map[string]interface{}{
			"fullname": fullName,
			"email":    email,
			"otp":      otp,
		},
	}

	// Enqueue with no params
	if err = c.Enqueue("cli", "default", body); err != nil {
		b.Logger.Error(err.Error())
		return
	}
	return
}

// Dial establishes a connection to the redis instance at url
func Dial(url string) (*Client, error) {
	c, err := rdsWorker.Dial("tcp", url)
	if err != nil {
		return nil, err
	}
	return &Client{c}, nil
}

func (c *Client) Enqueue(class, queue string, args ...interface{}) error {
	var j = &job{class, makeJobArgs(args)}

	job_json, _ := json.Marshal(j)

	if err := c.conn.Send("LPUSH", "resque:queue:"+queue, job_json); err != nil {
		return err
	}
	return c.conn.Flush()
}

func (c *Client) Close() error {
	return c.conn.Close()
}

//A trick to make [{}] json struct for empty args
func makeJobArgs(args []interface{}) []interface{} {
	if len(args) == 0 {
		// NOTE: Dirty hack to make a [{}] JSON struct
		return append(make([]interface{}, 0), make(map[string]interface{}, 0))
	}

	return args
}
