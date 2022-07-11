package tools

import (
	"time"

	"github.com/google/uuid"
	"github.com/widyan/go-codebase/middleware/interfaces"
)

type ToolsImpl struct {
}

func CreateTools() interfaces.ToolsInterface {
	return &ToolsImpl{}
}

func (t *ToolsImpl) GetUUID() (uid string) {
	return uuid.New().String()
}

func (t *ToolsImpl) GetTimeNowUnix(hour int) int64 {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(loc).Add(time.Hour * time.Duration(hour)).Unix()
}

func (t *ToolsImpl) GetTimeNowUnixIssued() int64 {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(loc).Unix()
}
