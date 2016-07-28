Middleware for [gin](https://github.com/gin-gonic/gin), when used with [fresh](https://github.com/pressly/fresh):

This will show any build errors:

```go
import showerrors "github.com/slava-vishnyakov/gin-fresh-showerrors-middleware"

router.Use(showerrors.FromFresh)
```

This will show any build errors and also will try to `go get` packages if they are missing:

```go
import showerrors "github.com/slava-vishnyakov/gin-fresh-showerrors-middleware"

router.Use(showerrors.FromFreshAndGoGet)
```
