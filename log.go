package obgae

import (
	"sync"

	gae "appengine"

	ob "github.com/openbase/ob-core"
	obsrv "github.com/openbase/ob-core/server"
)

func initLogHooks() {
	var logMutex sync.Mutex
	noopLog := ob.Log // we passed this above, but take from ob anyway

	obsrv.On.Request.Serving.Add(func(rc *obsrv.RequestContext) {
		ctx := gae.NewContext(rc.Req)
		rc.Ctx, rc.Log = ctx, newLogger(ctx)
		logMutex.Lock()
		defer logMutex.Unlock()
		ob.Log = rc.Log
	})

	obsrv.On.Request.Served.Add(func(rc *obsrv.RequestContext) {
		logMutex.Lock()
		defer logMutex.Unlock()
		// if the "global Log" is still ours, reset to original (noop-dummy in GAE case),
		//	else another request took over meanwhile: then ignore
		if ob.Log == rc.Log {
			ob.Log = noopLog
		}
	})
}

type ctxLogger struct {
	gae.Context
}

func newLogger(ctx gae.Context) (me *ctxLogger) {
	me = &ctxLogger{}
	me.Context = ctx
	return
}

func (me *ctxLogger) Error(err error) error {
	me.Errorf(err.Error())
	return err
}
