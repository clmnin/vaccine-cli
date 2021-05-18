package main

import (
	"fmt"

	"github.com/clmno/vaccine-cli/cowin"
)

func main() {
	centers := cowin.GetAvailableSessions(307, 18)
	fmt.Println(centers)
}
