/**
 * @author: D-S
 * @date: 2021/1/27 10:36 上午
 */

package node

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

type Router struct {
	Middlewares []gin.HandlerFunc
	Engine      *gin.Engine
}

func NewRouter(middleware ...gin.HandlerFunc) *Router {
	r := defaultRouter()
	r.Engine.Use(middleware...)
	return r
}

func defaultRouter() *Router {
	gin.SetMode(gin.ReleaseMode)
	r := &Router{
		Middlewares: nil,
		Engine:      gin.Default(),
	}
	r.Engine.Use(r.cors())
	r.Engine.Use(r.ginHttp())
	return r
}

func (r *Router) cors() gin.HandlerFunc {
	midCors := cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "*"},
		AllowHeaders:     []string{"Origin", "*", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "*"},
		AllowCredentials: true,
	})
	return midCors
}

func (r *Router) ginHttp() gin.HandlerFunc {
	return ginhttp.Middleware(opentracing.GlobalTracer(), ginhttp.OperationNameFunc(func(r *http.Request) string {
		return r.URL.Path
	}))
}
