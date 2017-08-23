package conditions

import (
	"github.com/tebeka/selenium"
)

// ElementIsLocated returns a condition that checks if theelement is found on page.
func ElementIsLocated(by, selector string) selenium.Condition {
	return func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.FindElement(by, selector)
		return err == nil, nil
	}
}
