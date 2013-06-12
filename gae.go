package obgae

import (
	"net/http"
	"sync"

	gae "appengine"

	ob "github.com/openbase/ob-core"
	obsrv "github.com/openbase/ob-core/server"
)

//	Call this in the init() func of your GAE app --- see demo-app/app.go.
//	While it returns an error, it nonetheless however also does binds a
//	default http handler serving the same error message to all clients.
//
//	So whether you handle err or not, Init() guarantees that, when it returns,
//	at least one http.Handler was registered to serve any and all web requests.
func Init(hiveDir string) (err error) {
	ob.Opt.Server, ob.Opt.Sandboxed = true, true
	if err = ob.Init(hiveDir, ob.NewLogger(nil)); err == nil {
		obsrv.Init()

		var logMutex sync.Mutex
		noopLog := ob.Opt.Log // we passed this above, but take from ob anyway

		obsrv.On.Request.Serving.Add(func(rc *obsrv.RequestContext) {
			ctx := gae.NewContext(rc.Req)
			rc.Ctx, rc.Log = ctx, newLogger(ctx)
			logMutex.Lock()
			defer logMutex.Unlock()
			ob.Opt.Log = rc.Log
		})

		obsrv.On.Request.Served.Add(func(rc *obsrv.RequestContext) {
			logMutex.Lock()
			defer logMutex.Unlock()
			// if the "global Log" is still ours, reset to original (noop-dummy in GAE case),
			//	else another request took over meanwhile: then ignore
			if ob.Opt.Log == rc.Log {
				ob.Opt.Log = noopLog
			}
		})

		http.Handle("/", obsrv.Router)
	}

	if err != nil {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, err.Error(), 500)
		})
	}
	return
}
