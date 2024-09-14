# formats

# usage

```sh
go get github.com/mentai-mayo/gotools/formats
```


```go
package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mentai-mayo/gotools/formats"
)

func main() {
	err := ReturnError()
	if err != nil {
		list, _ := formats.Formats("[main]", "こんなエラーがでたよ: %s", err)
		fmt.Println(strings.Join(list, " "))
	}
}

func ReturnError() error {
	return errors.New("なんかエラーでたよ")
}

```
