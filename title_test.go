package conditions_test

import (
	"testing"
	"time"

	conditions "github.com/serge1peshcoff/selenium-go-conditions"
	"github.com/tebeka/selenium"
)

func TestTitleIs(t *testing.T) {
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, "")
	defer wd.Quit()

	if err != nil {
		t.Fatalf("Cannot start selenium.NewRemote(): %v\n", err)
	}

	// Testing successful selenium.Wait() call.
	err = wd.Get("http://localhost:3000/index")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/index: %v\n", err)
	}
	element, err := wd.FindElement(selenium.ByID, "changeTitle")
	if err != nil {
		t.Fatalf("Cannot find element #changeTitle: %v\n", err)
	}
	if err = element.Click(); err != nil {
		t.Fatalf("Cannot click on  element #changeTitle: %v\n", err)
	}

	// This should not raise an error.
	if err = wd.Wait(conditions.TitleIs("Title changed.")); err != nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}

	// Testing unsuccessful selenium.Wait() call (this should raise error cause of timeout).
	err = wd.Get("http://localhost:3000/index")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/index: %v\n", err)
	}

	// This should raise an timeout error.
	if err = wd.WaitWithTimeout(conditions.TitleIs("Title changed."), 500*time.Millisecond); err == nil {
		t.Fatalf("wd.Wait() should raise an error, but it didn't.\n")
	}
}

func TestTitleIsNot(t *testing.T) {
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, "")
	defer wd.Quit()

	if err != nil {
		t.Fatalf("Cannot start selenium.NewRemote(): %v\n", err)
	}

	// Testing successful selenium.Wait() call.
	err = wd.Get("http://localhost:3000/index")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/index: %v\n", err)
	}
	element, err := wd.FindElement(selenium.ByID, "changeTitle")
	if err != nil {
		t.Fatalf("Cannot find element #changeTitle: %v\n", err)
	}
	if err = element.Click(); err != nil {
		t.Fatalf("Cannot click on  element #changeTitle: %v\n", err)
	}

	// This should not raise an error.
	if err = wd.Wait(conditions.TitleIsNot("Selenium Test Page")); err != nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}

	// Testing unsuccessful selenium.Wait() call (this should raise error cause of timeout).
	err = wd.Get("http://localhost:3000/index")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/index: %v\n", err)
	}

	// This should raise an timeout error.
	if err = wd.WaitWithTimeout(conditions.TitleIsNot("Selenium Test Page"), 500*time.Millisecond); err == nil {
		t.Fatalf("wd.Wait() should raise an error, but it didn't.\n")
	}
}
