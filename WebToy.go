package WebToy

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct{
	r *router
}

//New create an instance of WebToy.Engine
func New() *Engine{
	return &Engine{r: newRouter()}
}

//addRoute add static router
func (engine *Engine) addRoute(method string,pattern string,handler HandlerFunc){
	engine.r.addRoute(method,pattern,handler)
}

//GET defines the method to add GET request
func (engine *Engine) GET(pattern string,handler HandlerFunc){
	engine.addRoute("GET",pattern,handler)
}

//POST defines the method to add POST request
func (engine *Engine) POST(pattern string,handler HandlerFunc){
	engine.addRoute("POST",pattern,handler)
}

// Run start a http service
func (engine *Engine) Run(addr string)(err error){
	return http.ListenAndServe(addr,engine)
}

// ServeHTTP Instance for http.ListenAndServe
func (engine *Engine)ServeHTTP(w http.ResponseWriter,req *http.Request){
	c:=newContext(w,req)
	engine.r.handle(c)
}

