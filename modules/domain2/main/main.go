package main

import (
	"codebase/go-codebase/modules/domain2"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgin"
	"go.elastic.co/apm/module/apmhttp"
	"go.elastic.co/apm/module/apmlogrus"
)

func main() {

	// Aktivasi environment
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Sprintf("%s: %s", "Failed to load env", err))
	}

	var logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			_, filename := path.Split(f.File)
			return funcname, filename
		},
	})
	logger.SetReportCaller(true)
	// logger.Hooks.Add(&apmlogrus.Hook{
	// 	LogLevels: logrus.AllLevels,
	// })
	logger.AddHook(&apmlogrus.Hook{
		LogLevels: logrus.AllLevels,
	})

	// var lo logger.Loggers
	routesGin := gin.New()
	if os.Getenv("MODE") == "development" {
		pprof.Register(routesGin)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	routesGin.Use(apmgin.Middleware(routesGin))

	routesGin, pq, redis, amqp := domain2.Init(routesGin, logger)
	s := &http.Server{
		Addr:         os.Getenv("PORT"),
		Handler:      apmhttp.Wrap(routesGin),
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 30,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			logger.Printf("listen: %s\n", err)
		}
		s.SetKeepAlivesEnabled(false)
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown:", err)
	}

	logger.Println("Server exiting")
	logger.Println("Close clonnection postgresql")
	logger.Println("Close clonnection redis")
	pq.Close()
	amqp.Close()
	redis.Close()
}
