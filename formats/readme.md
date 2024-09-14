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

	"github.com/mentai-mayo/gotools/formats"
)

func main() {
	err := ReturnError()
	if err != nil {
		str, _ := formats.Sprintf("[main]", "こんなエラーがでたよ: %s", err)
		fmt.Println(str)
	}
}

func ReturnError() error {
	return errors.New("なんかエラーでたよ")
}

```
