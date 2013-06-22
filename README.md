# obgae
--
    import "github.com/openbase/ob-gae"

To run OpenBase on AppEngine, import this package into yourapp.go, then call Init().

## Usage

#### func  Init

```go
func Init(hiveDir string) (err error)
```
Call this in the init() func of your GAE app --- see demo-app/app.go. While it
returns an error, it nonetheless however also does binds a default http handler
serving the same error message to all clients.

So whether you handle err or not, Init() guarantees that, when it returns, at
least one http.Handler was registered to serve any and all web requests.

--
**godocdown** http://github.com/robertkrimen/godocdown