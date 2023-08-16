# go-web-components

A simple template wrapper to define web components in go

```go
package main

import (
	wcomponent "github.com/zapling/go-web-components"
)

func main() {
    c := wcomponent.Component{
        Selector:     "my-thingy",
        HTML:         template.HTML(`<div>Hello World</div>`),
        UseShadowDOM: false,
    }

    js, err := c.GetJavascript()
    // ... 
}
```
