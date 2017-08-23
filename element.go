package conditions

import (
	"github.com/tebeka/selenium"
)

// ElementIsLocated returns a condition that checks if the element is found on page.
func ElementIsLocated(by, selector string) selenium.Condition {
	return func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.FindElement(by, selector)
		return err == nil, nil
	}
}

// ElementIsVisible returns a condition that checks if the element is visible.
func ElementIsVisible(elt selenium.WebElement) selenium.Condition {
	return func(wd selenium.WebDriver) (bool, error) {
		visible, err := elt.IsDisplayed()
		return visible, err
	}
}

// ElementIsLocatedAndVisible returns a condition that checks if the element is found on page and is visible.
func ElementIsLocatedAndVisible(by, selector string) selenium.Condition {
	return func(wd selenium.WebDriver) (bool, error) {
		element, err := wd.FindElement(by, selector)
		if err != nil {
			return false, nil
		}
		visible, err := element.IsDisplayed()
		return visible, err
	}
}