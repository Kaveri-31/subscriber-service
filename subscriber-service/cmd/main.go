// Package main Subscriber Service API
//
// @title Subscriber Service API
// @version 1.0
// @description Subscriber Management REST API using Gin, MongoDB, Kafka and Prometheus.
// @termsOfService http://swagger.io/terms/
//
// @contact.name Kaveri
// @contact.email kaveri@example.com
//
// @license.name MIT
//
// @host localhost:7070
// @BasePath /
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "subscriber-service/docs"

	"subscriber-service/internal/config"
	"subscriber-service/internal/database"
	"subscriber-service/internal/handler"
	"subscriber-service/internal/kafka"
	"subscriber-service/internal/logger"
	"subscriber-service/internal/metrics"
)

func main() {

	// Initialize Logger
	logger.Init()

	logrus.Info("Starting Subscriber Service")

	// Load Configuration
	cfg, err := config.Load(
		"configs/config.yaml",
	)

	if err != nil {

		logrus.WithError(err).
			Fatal("Failed to load configuration")
	}

	// Print Configuration
	logrus.Info("================================")

	logrus.Infof(
		"Application : %s",
		cfg.Application.Name,
	)

	logrus.Infof(
		"Environment : %s",
		cfg.Application.Environment,
	)

	logrus.Infof(
		"Port        : %d",
		cfg.Server.Port,
	)

	logrus.Infof(
		"Mongo URI   : %s",
		cfg.MongoDB.URI,
	)

	logrus.Infof(
		"Database    : %s",
		cfg.MongoDB.Database,
	)

	logrus.Infof(
		"Collection  : %s",
		cfg.MongoDB.Collection,
	)

	logrus.Infof(
		"Kafka Broker: %s",
		cfg.Kafka.Broker,
	)

	logrus.Infof(
		"Kafka Topic : %s",
		cfg.Kafka.Topic,
	)

	logrus.Info("================================")

	// Connect MongoDB
	database.ConnectMongoDB()

	// Register Prometheus Metrics
	metrics.RegisterMetrics()

	logrus.Info(
		"Prometheus Metrics Registered",
	)

	// Initialize Kafka Producer
	if err := kafka.InitProducer(); err != nil {

		logrus.WithError(err).
			Fatal("Failed to initialize Kafka Producer")
	}

	logrus.Info(
		"Kafka Producer Initialized",
	)

	// Create Router
	router := gin.Default()

	router.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)

	// Prometheus Endpoint
	router.GET(
		"/metrics",
		gin.WrapH(
			promhttp.Handler(),
		),
	)

	// Health Check Endpoints
	router.GET(
		"/health",
		handler.Health,
	)

	router.GET(
		"/ready",
		handler.Ready,
	)

	router.GET(
		"/live",
		handler.Live,
	)

	// Subscriber Endpoints
	router.GET(
		"/subscribers",
		handler.GetSubscribers,
	)

	router.GET(
		"/subscribers/:id",
		handler.GetSubscriberByID,
	)

	router.POST(
		"/subscribers",
		handler.CreateSubscriber,
	)

	router.PUT(
		"/subscribers/:id",
		handler.UpdateSubscriber,
	)

	router.DELETE(
		"/subscribers/:id",
		handler.DeleteSubscriber,
	)

	serverAddress := fmt.Sprintf(
		":%d",
		cfg.Server.Port,
	)

	server := &http.Server{
		Addr:    serverAddress,
		Handler: router,
	}

	go func() {

		logrus.Infof(
			"Subscriber Service running on %s",
			serverAddress,
		)

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {

			logrus.WithError(err).
				Fatal("Failed to start server")
		}

	}()

	// Graceful Shutdown

	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-quit

	logrus.Info(
		"Shutdown signal received",
	)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {

		logrus.WithError(err).
			Error("Server forced to shutdown")
	}

	logrus.Info(
		"HTTP Server Stopped",
	)

	kafka.CloseProducer()

	logrus.Info(
		"Kafka Producer Closed",
	)

	database.DisconnectMongoDB()

	logrus.Info(
		"MongoDB Connection Closed",
	)

	logrus.Info(
		"Subscriber Service Stopped Successfully",
	)
}
