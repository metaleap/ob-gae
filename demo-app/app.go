package obgae_demoapp

import (
	"net/http"
	"os"

	ob "github.com/openbase/ob-core"
	obsrv "github.com/openbase/ob-core/server"
)

func init() {
	ob.Opt.Server, ob.Opt.Sandboxed = true, true
	//	On GAE, the specified Hive-directory path MUST be empty ("")
	ob.Init("", os.Stdout)
	http.Handle("/", obsrv.Mux)
}
