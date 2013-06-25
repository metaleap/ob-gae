package obgae

import (
	"appengine"
	"net/http"

	ob "github.com/openbase/ob-core"
	obsrv "github.com/openbase/ob-core/server"
)

//	Call this in the `init` func of your GAE app --- see `demo-app/app.go`.
//
//	This always binds a working URL handler to `/`, even if an `error` was encountered
//	during initialization -- in which case, the initialization error message is rendered
//	in plain-text with an HTTP 500 status to all client web requests (while each time
//	also logging an Error-level message with GAE during such a request).
func Init(hiveDirPath string) {
	ctx, err := obsrv.NewCtx(hiveDirPath, ob.NewLogger(nil))
	if err == nil {
		handler := obsrv.NewHttpHandler(ctx)
		initLogHooks(handler)
		http.Handle("/", handler)
	} else {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			gaeCtx := appengine.NewContext(r)
			gaeCtx.Errorf("App could not be initialized: %#v", err)
			http.Error(w, err.Error(), 500)
		})
	}
}
