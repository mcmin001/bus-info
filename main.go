package main

import (
	http_util "bus-info/util/http"
	"fmt"
)

func main() {
	http_util.Httpget()
	fmt.Println("hello")
}
