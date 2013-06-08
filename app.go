package obgae

import (
	"net/http"

	ob "github.com/openbase/ob-core"
	obsrv "github.com/openbase/ob-core/server"
)

func init() {
	if !ob.Sandboxed {
		ob.Sandboxed = true
		ob.Init("")
		http.Handle("/", obsrv.Mux)
	}
}
