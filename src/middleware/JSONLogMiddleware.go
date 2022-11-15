package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sbecker/gin-api-demo/util"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
func JSONLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
		c.Request.Body = rdr2
		c.Next()

		time := time.Now()
		LogName := strconv.Itoa(time.Year()) + "_" + strconv.Itoa(int(time.Month())) + "_" + strconv.Itoa(time.Day()) + ".log"

		file, _ := os.OpenFile("./log/"+LogName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(file)
		log.SetLevel(log.InfoLevel)

		duration := util.GetDurationInMillseconds(time)
		newLog := log.WithFields(log.Fields{
			"client_ip":     util.GetClientIP(c),
			"duration_ms":   duration,
			"method":        c.Request.Method,
			"path":          c.Request.RequestURI,
			"status":        c.Writer.Status(),
			"request_body":  readBody(rdr1),
			"response_body": blw.body.String(),
		})
		if c.Writer.Status() >= 500 {
			newLog.Warn("server error")
		} else if c.Writer.Status() >= 400 {
			newLog.Warn("warn")
		} else {
			newLog.Info("default")
		}
	}
}
func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}
