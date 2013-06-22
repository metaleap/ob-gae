package obgae

import (
	"net/http"

	gae "appengine"

	ob "github.com/openbase/ob-core"
	obsrv "github.com/openbase/ob-core/server"
)

//	Call this in the `init()` func of your GAE app --- see `demo-app/app.go`.
//	This always binds a working URL handler to `/`, even if an error was encountered during `Init()`.
func Init(hiveDir string) {
	var err error
	ob.Opt.Server, ob.Opt.Sandboxed = true, true
	if err = ob.Init(hiveDir, ob.NewLogger(nil)); err == nil {
		obsrv.Init()
		initLogHooks()
		http.Handle("/", obsrv.Router)
	} else {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			ctx := gae.NewContext(r)
			ctx.Errorf("App could not be initialized: %#v", err)
			http.Error(w, err.Error(), 500)
		})
	}
}
