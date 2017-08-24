package conditions

import (
	"github.com/tebeka/selenium"
	"strings"
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

// ElementTextIs returns a condition that checks if element's text equals to string.
func ElementTextIs(elt selenium.WebElement, text string) selenium.Condition {
	return func (wd selenium.WebDriver) (bool, error) {
		eltText, err := elt.Text()
		if err != nil {
			return false, err
		}

		return eltText == text, nil
	}
}

// ElementTextContains returns a condition that checks if element's text contains a string.
func ElementTextContains(elt selenium.WebElement, text string) selenium.Condition {
	return func (wd selenium.WebDriver) (bool, error) {
		eltText, err := elt.Text()
		if err != nil {
			return false, err
		}

		return strings.Contains(eltText, text), nil
	}
}