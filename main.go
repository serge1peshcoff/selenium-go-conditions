// Package conditions is providing some conditions for WebDriver.Wait() function
// from the "github.com/tebeka/selenium" package.
package conditions

import "github.com/tebeka/selenium"

// TitleIs returns a condition that checks if the title matches the expectedTitle.
func TitleIs(expectedTitle string) selenium.Condition {
	return func(wd selenium.WebDriver) (bool, error) {
		title, err := wd.Title()
		if err != nil {
			return false, err
		}

		return title == expectedTitle, nil
	}
}

// TitleIsNot returns a condition that checks if the title doesn't match the expectedTitle.
func TitleIsNot(expectedTitle string) selenium.Condition {
	return func(wd selenium.WebDriver) (bool, error) {
		title, err := wd.Title()
		if err != nil {
			return false, err
		}

		return title != expectedTitle, nil
	}
}
