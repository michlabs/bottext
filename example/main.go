package main 

import (
	"fmt"
	"github.com/michlabs/bottext"
)

func main() {
	bottext.MustLoad("text.yml")
	T1 := bottext.New("en")
	for i := 0; i < 10; i ++ {
		fmt.Printf(T1("welcome") + "\n", "Donald Trump")	
	}

	T2 := bottext.New("vi")
	for i := 0; i < 10; i ++ {
		fmt.Printf(T2("welcome") + "\n", "Đỗ Nam Trung")	
	}
}