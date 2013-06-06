package obgae

import (
	obcore "github.com/openbase/ob-core"
)

func init() {
	if !obcore.Sandboxed {
		obcore.Sandboxed = true
		obcore.Init("")
	}
}
