package xcaddybug

import "github.com/caddyserver/caddy/v2"

func init() {
	caddy.RegisterModule(Example{})
}

type Example struct {
}

func (Example) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "foo.example",
		New: func() caddy.Module { return new(Example) },
	}
}
