package WebToy

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct{
	//origin content
	Writer http.ResponseWriter
	Req *http.Request
	//request info
	Path string
	Method string
	//response info
	StatusCode int
}

func newContext(w http.ResponseWriter,req *http.Request) *Context{
	return &Context{Writer:w,
					Req:req,
					Path:req.URL.Path,
					Method: req.Method,
			}
}

// Status set Response value--Status code
func (c *Context)Status(Status  int){
	c.StatusCode=Status
	c.Writer.WriteHeader(Status)
}

// SetHeader set Response value--Header
func (c *Context)SetHeader(key string,value string){
	c.Writer.Header().Set(key,value)
}

// Query get Request value--Query
func (c *Context)Query(key string) string{
	return c.Req.URL.Query().Get(key)
}

// PostForm get Request value--PostForm
func (c *Context)PostForm(key string) string{
	return c.Req.FormValue(key)
}

func (c *Context)HTML(code int,html string){
	c.SetHeader("Content-Type","text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

func (c *Context)Data(code int,data []byte){
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context)JSON(code int,obj interface{}){
	c.SetHeader("Content-Type","text/json")
	c.Status(code)
	encoder:=json.NewEncoder(c.Writer)
	if err:=encoder.Encode(obj);err!=nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context)String(code int,format string,value ...interface{}){
	c.SetHeader("Content-Type","text/string")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format,value...)))
}
