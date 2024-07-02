package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/utils"
)

var searchBoxer = "lomachenko"

func main() {

	browser := rod.New().MustConnect().NoDefaultDevice().MustIncognito()
	page := browser.MustPage("https://www.tapology.com/").MustWindowFullscreen()

	page.MustElement("#siteSearch").MustInput(searchBoxer)

	popup := page.MustSearch("Button__StyledButton-a1qza5-0 gVUTik")
	if popup != nil {
		popup.MustClick()
	}

	page.MustElement("#search").MustClick()

	popup2 := page.MustSearch("Button__StyledButton-a1qza5-0 gVUTik")
	if popup2 != nil {
		popup2.MustClick()
	}

	page.MustSearch(".searchResultsFighter > .fcLeaderboard > tbody > tr > td > a ").MustClick()

	img := page.MustElement(".fighterImg > img")
	_ = utils.OutputFile("searchboxer.png", img.MustResource())
	el := page.MustElement("#stats")

	textinfo := el.MustText()
	os.WriteFile("boxer-def.txt", []byte(textinfo), 0666)

	section := page.MustElement("#proResults")

	// get children elements of an element
	combats := section.MustElements("li")

	for _, combat := range combats {
		//part := section.MustElements("li")
		strings.Split(combat.MustText(), " ")
		fmt.Printf(
			"Combat : '%s'",
			combat.MustText(),
		)
	}

	il := page.MustElement("#proResults")
	textResult := il.MustText()

	fmt.Println(textResult)

	os.Exit(0)

}
