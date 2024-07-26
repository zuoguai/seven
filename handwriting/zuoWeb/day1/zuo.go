package day1

import (
	"net/http"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandleFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

func (e *Engine) addRoute(method string, pattern string, handle HandleFunc) {
	key := method + "-" + pattern
	e.router[key] = handle
}

func (e *Engine) GET(pattern string, handle HandleFunc) {
	e.addRoute("GET", pattern, handle)
}

func (e *Engine) PUT(pattern string, handle HandleFunc) {
	e.addRoute("PUT", pattern, handle)
}

func (e *Engine) DELETE(pattern string, handle HandleFunc) {
	e.addRoute("DELETE", pattern, handle)
}

func (e *Engine) POST(pattern string, handle HandleFunc) {
	e.addRoute("POST", pattern, handle)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handle, ok := e.router[key]; ok {
		handle(w, r)
	} else {
		http.NotFound(w, r)
	}
}
