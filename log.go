package obgae

import (
	"appengine"

	"github.com/go-utils/ugo"

	obsrv "github.com/openbase/ob-core/server"
)

func initLogHooks(ctx *obsrv.Ctx) {
	var logMutex ugo.MutexIf
	noopLog := ctx.Log // should be our logger that we passed in `Init`

	ctx.Http.Handler.On.Request.PreServe.Add(func(rc *obsrv.RequestContext) {
		rc.Log = newLogger(appengine.NewContext(rc.Req))
		defer logMutex.UnlockIf(logMutex.Lock())
		ctx.Log = rc.Log
	})

	ctx.Http.Handler.On.Request.PostServe.Add(func(rc *obsrv.RequestContext) {
		defer logMutex.UnlockIf(logMutex.Lock())
		// if the "global Log" is still ours, reset to original (noop-dummy in GAE case),
		//	else another request took over meanwhile: then ignore
		if ctx.Log == rc.Log {
			ctx.Log = noopLog
		}
	})
}

type logger struct {
	appengine.Context
}

func newLogger(gaeCtx appengine.Context) (me *logger) {
	me = &logger{Context: gaeCtx}
	return
}

func (me *logger) Error(err error) error {
	me.Errorf(err.Error())
	return err
}
