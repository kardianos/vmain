package main // import "github.com/kardianos/vmain"

import (
	"fmt"

	// "github.com/kardianos/vtest"
	vtest2 "github.com/kardianos/vtest/v2"
	vtest3 "github.com/kardianos/vtest/v3"
)

func main() {
	fmt.Printf("I dream %s %s", vtest2.Peace(), vtest3.Peace())
}
