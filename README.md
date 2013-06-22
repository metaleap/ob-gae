# obgae
--
    import "github.com/openbase/ob-gae"

To run on App Engine, `import` this package into your_app.go and call `Init()`.

## Usage

#### func  Init

```go
func Init(hiveDir string)
```
Call this in the `init()` func of your GAE app --- see `demo-app/app.go`. This
always binds a working URL handler to `/`, even if an error was encountered
during `Init()`.

--
**godocdown** http://github.com/robertkrimen/godocdown