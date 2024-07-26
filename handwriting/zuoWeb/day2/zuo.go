package day2

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: NewRouter()}
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handle HandlerFunc) {
	e.addRoute("GET", pattern, handle)
}

func (e *Engine) PUT(pattern string, handle HandlerFunc) {
	e.addRoute("PUT", pattern, handle)
}

func (e *Engine) DELETE(pattern string, handle HandlerFunc) {
	e.addRoute("DELETE", pattern, handle)
}

func (e *Engine) POST(pattern string, handle HandlerFunc) {
	e.addRoute("POST", pattern, handle)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.router.handle(c)
}
