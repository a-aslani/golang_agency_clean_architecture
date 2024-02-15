package restapi

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (r *controller) metricRequestCounter() gin.HandlerFunc {
	return func(c *gin.Context) {
		r.reqCounter.Inc()
		c.Next()
	}
}

func (r *controller) metricLatencyRequestChecker() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := time.Now()

		c.Next()

		latency := time.Since(t)

		r.reqLatency.Observe(latency.Seconds())
	}
}

func (r *controller) authentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		// tokenInBytes, err := r.JwtToken.VerifyToken(c.GetHeader("token"))
		// if err != nil {
		// 	c.AbortWithStatus(http.StatusForbidden)
		// 	return
		// }
		//
		// var dataToken payload.DataToken
		// err = json.Unmarshal(tokenInBytes, &dataToken)
		// if err != nil {
		// 	c.AbortWithStatus(http.StatusForbidden)
		// 	return
		// }
		//
		// c.Set("data", dataToken)
		//
		// c.AbortWithStatus(http.StatusForbidden)
		// return

	}
}

func (r *controller) authorization() gin.HandlerFunc {

	return func(c *gin.Context) {

		authorized := true

		if !authorized {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
}
