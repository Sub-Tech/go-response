# go-response

```
get get github.com/tomwright/go-response
```

## Usage

```
import response "github.com/tomwright/go-response"

v := response.New()
v.Vars["x"] = "y"
err = v.RenderJSON(w, r)
if err != nil {
    panic(err)
}
```