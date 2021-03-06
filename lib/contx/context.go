package contx

import (
	"github.com/go-macaron/cache"
	"github.com/go-macaron/session"
	"github.com/go-macaron/toolbox"
	"gopkg.in/macaron.v1"
	"net/http"
	"encoding/json"
    "log"
)

var ctx *Context

// Context representation
type Context struct {
	*macaron.Context
	Session session.Store
	Flash   *session.Flash
	Cache   cache.Cache
	Toolbox toolbox.Toolbox
}

// Contexter middleware
func Contexter() macaron.Handler {
	return func(c *macaron.Context, session session.Store, flash *session.Flash, cache cache.Cache, toolbox toolbox.Toolbox) {
		ctx = &Context{
			Context: c,
			Session: session,
			Flash:   flash,
			Cache:   cache,
			Toolbox: toolbox,
		}
		c.Map(ctx)
	}
}


// NativeHTML render using go engine
func (ctx *Context) NativeHTML(status int, name string) {
	ctx.Context.HTML(status, name, ctx.Data)
}

// JSONWithoutEscape render json without escape
func (ctx *Context) JSONWithoutEscape(status int, obj interface{}) {
	ctx.Header().Set("Content-Type", "application/json")
	ret, err := json.Marshal(&obj)
	if err != nil {
		log.Print("[JSONWithoutEscape]" + err.Error())
		http.Error(ctx.Resp, "{'errors':'JSON Marshaling Error = "+err.Error()+"'}", 500)
		return
	}
	ctx.Status(status)
	log.Println("[JSONWithoutEscape] Returned object: " + string(ret))
	ctx.Resp.Write(ret)
}

// I18n view func
func I18n(key string) string {
	return ctx.Tr(key)

}

/*
GetContext Get system context
*/
func GetContext() *Context {
    return ctx
}

