package main

import (
	"fmt"

	"github.com/serge1peshcoff/selenium-go-conditions"
	"github.com/tebeka/selenium"
)

func main() {
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	if err := wd.Get("https://csgowitch.com"); err != nil {
		panic(err)
	}

	if err := wd.Wait(conditions.TitleIs("CSGOWitch Gambling!")); err != nil {
		panic(err)
	}

	fmt.Printf("Page loaded\n")
}
