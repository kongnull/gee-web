package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandleFunc
}

func (engine *Engine) addRoute(method string, patter string, handler HandleFunc) {
	key := method + "-" + patter
	engine.router[key] = handler
}
func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

func (engine *Engine) Get(pattern string, handler HandleFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) Post(patter string, handler HandleFunc) {
	engine.addRoute("POST", patter, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := request.Method + "-" + request.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(writer, request)
	} else {
		_, err := fmt.Fprintf(writer, "404 NOT FOUND: %s\n", request.URL)
		if err != nil {
			return
		}
	}
}
