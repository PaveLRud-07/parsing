package main

import (
	"github.com/tebeka/selenium"
)

func mn() {

	// The paths to these binaries will be different on your machine!

	const (
		seleniumPath    = "C:/Users/Skrap/go/pkg/mod/github.com/tebeka/selenium@v0.9.9/vendor/selenium-server.jar"
		geckoDriverPath = "C:/Users/Skrap/go/pkg/mod/github.com/tebeka/selenium@v0.9.9/vendor/chromedriver.exe"
	)

	service, err := selenium.NewSeleniumService(
		seleniumPath,
		4444,
		selenium.GeckoDriver(geckoDriverPath))

	if err != nil {
		panic(err)
	}
	defer service.Stop()
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	err = wd.Get("https://www.packtpub.com/networking-and-servers/mastering-go")
	if err != nil {
		panic(err)
	}

	var elems []selenium.WebElement
	wd.Wait(func(wd2 selenium.WebDriver) (bool, error) {
		elems, err = wd.FindElements(selenium.ByCSSSelector, "div.product-reviews-review div.review-body")
		if err != nil {
			return false, err
		} else {
			return len(elems) > 0, nil
		}
	})
	for _, review := range elems {
		body, err := review.Text()
		if err != nil {
			panic(err)
		}
		println(body)
	}
}
