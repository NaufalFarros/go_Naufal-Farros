package Middleware

import (
	"time"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		req := c.Request()
// 		res := c.Response()

// 		log.SetFormatter(&log.TextFormatter{
// 			ForceColors:   true,
// 			FullTimestamp: true,
// 		})

// 		log.WithFields(log.Fields{
// 			"method": req.Method,
// 			"uri":    req.RequestURI,
// 			"status": res.Status,
// 			"time":   time.Now().Format(time.RFC3339),
// 		}).Info()

// 		return next(c)
// 	}
// }

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)

		log.SetFormatter(&log.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})

		log.WithFields(log.Fields{
			"time":    time.Now().Format("2006-01-02 15:04:05"),
			"status":  c.Response().Status,
			"method":  c.Request().Method,
			"uri":     c.Request().RequestURI,
			"latency": time.Since(start).String(),
		}).Info("")
		return err
	}
}
