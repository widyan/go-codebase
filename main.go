package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	config "github.com/widyan/go-codebase/config/apm"
	"github.com/widyan/go-codebase/helper"
	"github.com/widyan/go-codebase/middleware"
	"github.com/widyan/go-codebase/modules/domain"
	"github.com/widyan/go-codebase/notification"
	"github.com/widyan/go-codebase/responses"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	validate "github.com/widyan/go-codebase/validator"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgin"
	"go.elastic.co/apm/module/apmhttp"
	"go.elastic.co/apm/transport"
)

func main() {

	// Aktivasi environment
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Sprintf("%s: %s", "Failed to load env", err))
	}

	response := responses.CreateCustomResponses(os.Getenv("PROJECT_NAME"))

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

	// Get dependency call API
	toolsAPI := helper.CreateToolsAPI(logger)

	// ************************************ Setting notification to telegram if become error ***********
	notifTelegram := notification.CreateNotification(toolsAPI, os.Getenv("PROJECT_NAME"), os.Getenv("TOKEN_BOT_TELEGRAM"), os.Getenv("CHAT_ID"))
	logger.AddHook(notifTelegram)
	// *************************************************************************************************

	// ************************************ Get dependency validator ***********************************
	validator := validator.New()
	vldt := validate.CreateValidator(validator)
	// *************************************************************************************************

	// ************************************ Config for implement DB ************************************
	cfg := config.CreateConfigImplAPM(logger)
	pq := cfg.Postgresql(os.Getenv("GORM_CONNECTION"), 20, 20)
	pqdbAuth := cfg.Postgresql(os.Getenv("POSTGRES_AUTH_CONNECTION"), 20, 20)
	// *************************************************************************************************

	// ************************************ Implement Gin Gonic as Framework ***************************
	routesGin := gin.New()
	if os.Getenv("MODE") == "development" {
		pprof.Register(routesGin)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	routesGin.Use(apmgin.Middleware(routesGin))
	// **************************************************************************************************

	// ************************************ Setting APM ************************************
	apm.DefaultTracer.Close()
	tracer, err := apm.NewTracer("", "")
	if err != nil {
		logger.Panic(err)
	}

	transport, err := transport.NewHTTPTransport()
	if err != nil {
		logger.Panic(err)
	}

	// transport.SetSecretToken(os.Getenv("ELASTIC_APM_SECRET_TOKEN"))
	u, err := url.Parse(os.Getenv("ELASTIC_APM_SERVER_URL"))
	if err != nil {
		logger.Panic(err)
	}

	transport.SetServerURL(u)
	tracer.Transport = transport

	auth := middleware.Init(routesGin, logger, pqdbAuth, vldt, response)
	domain.Init(routesGin, logger, vldt, pq, response, auth)
	s := &http.Server{
		Addr:         os.Getenv("PORT"),
		Handler:      apmhttp.Wrap(routesGin, apmhttp.WithTracer(tracer)),
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 30,
	}
	// *****************************************************************************************

	/*
		// ************************************ Setting DATADOG ************************************

		//import this librrary if you want setting datadog
		//"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
		//gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"

		tracer.Start()
		defer tracer.Stop()

		routesGin.Use(gintrace.Middleware("metanesia-payment"))
		s := &http.Server{
			Addr:         os.Getenv("PORT"),
			Handler:      routesGin,
			WriteTimeout: time.Second * 60,
			ReadTimeout:  time.Second * 30,
		}
		// *****************************************************************************************
	*/

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
	logger.Println("Close connection postgresql")
	pq.Close()
	pqdbAuth.Close()
	// logger.Println("Close connection amqp")
	// amqp.Close()

}
