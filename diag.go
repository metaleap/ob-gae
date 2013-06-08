package obgae

import (
	"appengine"
)

type ctxLogger struct {
	appengine.Context
}

func newLogger(ctx appengine.Context) (me *ctxLogger) {
	me = &ctxLogger{}
	me.Context = ctx
	return
}

func (me *ctxLogger) Fatal(err error) {
	me.Criticalf("FATAL: %+v", err)
	panic(err)
}
