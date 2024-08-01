package pprof

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

type Option func(*option)

type option struct {
	prefix string
}

func WithPrefix(prefix string) Option {
	return func(o *option) {
		o.prefix = prefix
	}
}

func Register(r *gin.Engine, opts ...Option) {
	RouteRegister(&(r.RouterGroup), opts...)
}

func RouteRegister(rg gin.IRouter, opts ...Option) {
	o := &option{
		prefix: "/debug/pprof",
	}
	for _, opt := range opts {
		opt(o)
	}
	r := rg.Group(o.prefix)
	{
		r.GET("/", gin.WrapF(pprof.Index))
		r.GET("/cmdline", gin.WrapF(pprof.Cmdline))
		r.GET("/profile", gin.WrapF(pprof.Profile))
		r.POST("/symbol", gin.WrapF(pprof.Symbol))
		r.GET("/symbol", gin.WrapF(pprof.Symbol))
		r.GET("/trace", gin.WrapF(pprof.Trace))
		r.GET("/allocs", gin.WrapH(pprof.Handler("allocs")))
		r.GET("/block", gin.WrapH(pprof.Handler("block")))
		r.GET("/goroutine", gin.WrapH(pprof.Handler("goroutine")))
		r.GET("/heap", gin.WrapH(pprof.Handler("heap")))
		r.GET("/mutex", gin.WrapH(pprof.Handler("mutex")))
		r.GET("/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))
	}
}
