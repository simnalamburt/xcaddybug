package xcaddybug

import (
	"github.com/caddyserver/caddy/v2"
	"golang.org/x/example/stringutil"
)

func init() {
	caddy.RegisterModule(Example{})
}

type Example struct {
}

func (Example) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  caddy.ModuleID(stringutil.Reverse("elpmaxe.oof")),
		New: func() caddy.Module { return new(Example) },
	}
}
