# obgae

To run on App Engine, `import` this package into `your_app.go` and call `Init`.

## Usage

#### func  Init

```go
func Init(hiveDirPath string)
```
Call this in the `init` func of your GAE app --- see `demo-app/app.go`.

This always binds a working URL handler to `/`, even if an `error` was
encountered during initialization -- in which case, the initialization error
message is rendered in plain-text with an HTTP 500 status to all client web
requests (while each time also logging an Error-level message with GAE during
such a request).
