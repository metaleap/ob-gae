package obgae

import (
	"appengine"
	"net/http"

	ob "github.com/openbase/ob-core"
	obsrv "github.com/openbase/ob-core/server"
)

//	Call this in the init() func of your GAE app --- see demo-app/app.go
func Init(hiveDir string) {
	ob.Opt.Server, ob.Opt.Sandboxed = true, true
	ob.Init(hiveDir, ob.NewLogger(nil))
	obsrv.Init()
	obsrv.On.Request.Serving.Add(func(rc *obsrv.RequestContext) {
		ctx := appengine.NewContext(rc.Req)
		rc.Ctx, rc.Log = ctx, newLogger(ctx)
	})
	http.Handle("/", obsrv.Mux)
}
