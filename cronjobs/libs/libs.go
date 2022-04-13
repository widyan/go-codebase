package libs

import "log"

type Task struct {
	Name string
	Cron string
}

type Tasks struct {
	Project string
	Tasks   []Task
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
