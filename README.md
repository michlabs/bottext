### BotText
BotText is a Go package that help you:
- Translate your bot into multiple languages,
- Make your bot more interesting by expressing the same thing in different ways.

### Installation
Make sure you have [setup GOPATH](https://golang.org/doc/code.html#GOPATH)
```bash
go get github.com/michlabs/bottext
```

### Example
Create a new language file named `text.yml`:
```yaml
vi:
  welcome:
    - Xin chào %s!
    - Hi %s
en:
  welcome:
    - Hello %s!
    - Welcome %s!
    - Hi %s
```
Then in your source code:
```go
package main 

import (
    "fmt"
    "github.com/michlabs/bottext"
)

func main() {
    bottext.MustLoad("text.yml")
    var T1 bottext.BotTextFunc
    T1 = bottext.New("en")
    for i := 0; i < 10; i ++ {
        fmt.Printf(T1("welcome") + "\n", "Donald Trump")    
    }

    T2 := bottext.New("vi")
    for i := 0; i < 10; i ++ {
        fmt.Printf(T2("welcome"), "Đỗ Nam Trung")   
    }
}
```

### Document
Read in [Godoc](http://godoc.org/github.com/michlabs/bottext)
