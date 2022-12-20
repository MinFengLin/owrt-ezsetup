package main

import (
	"fmt"

	http "github.com/Minfenglin/owrt-ezsetup/http"
)

func init() {
	fmt.Println("owrt-ezsetup start!!!!")
}

func main() {
	http.Server_run()
}
