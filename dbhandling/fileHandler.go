package dbhandling

import (
	"fmt"
	"os"

	"github.com/artytheparty/project-0/acc"
)

//WriteToFile creates the file
func WriteToFile(c map[int]acc.AccountHolder) {
	f, err := os.Create("accounts.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 1; i < len(c)+1; i++ {
		f.WriteString((c[i]).Info())
	}
	err = f.Close()

}

// func main() {
// 	fmt.Println("I handle files")
// }
