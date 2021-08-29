package xcaddybug

import (
	"github.com/caddyserver/caddy/v2"
	stringutil1 "replacetest1/stringutil"
	stringutil2 "replacetest2/stringutil"
	stringutil3 "replacetest3/stringutil"
	stringutil4 "replacetest4/stringutil"
)

func init() {
	caddy.RegisterModule(Example{})
}

type Example struct {
}

func (Example) CaddyModule() caddy.ModuleInfo {
	ID := stringutil1.Reverse(stringutil2.Reverse(stringutil3.Reverse(stringutil4.Reverse("foo.example"))))

	return caddy.ModuleInfo{
		ID:  caddy.ModuleID(ID),
		New: func() caddy.Module { return new(Example) },
	}
}
