package conditions_test

import (
	"testing"
	"time"

	conditions "github.com/serge1peshcoff/selenium-go-conditions"
)

func TestTitleIs(t *testing.T) {
	// Testing successful selenium.Wait() call.
	err := wd.Get("http://localhost:3000/title_change")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/title_change: %v\n", err)
	}

	// This should not raise an error.
	if err = wd.Wait(conditions.TitleIs("Title changed.")); err != nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}

	// Testing unsuccessful selenium.Wait() call (this should raise error cause of timeout).
	err = wd.Get("http://localhost:3000/static")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/static: %v\n", err)
	}

	// This should raise an timeout error.
	if err = wd.WaitWithTimeout(conditions.TitleIs("Title changed."), 500*time.Millisecond); err == nil {
		t.Fatalf("wd.Wait() should raise an error, but it didn't.\n")
	}
}

func TestTitleIsNot(t *testing.T) {
	// Testing successful selenium.Wait() call.
	err := wd.Get("http://localhost:3000/title_change")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/title_change: %v\n", err)
	}

	// This should not raise an error.
	if err = wd.Wait(conditions.TitleIsNot("Selenium Test Page")); err != nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}

	// Testing unsuccessful selenium.Wait() call (this should raise error cause of timeout).
	err = wd.Get("http://localhost:3000/static")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/static: %v\n", err)
	}

	// This should raise an timeout error.
	if err = wd.WaitWithTimeout(conditions.TitleIsNot("Selenium Test Page"), 500*time.Millisecond); err == nil {
		t.Fatalf("wd.Wait() should raise an error, but it didn't.\n")
	}
}

func TestTitleContains(t *testing.T) {

	// Testing successful selenium.Wait() call.
	err := wd.Get("http://localhost:3000/title_change")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/title_change: %v\n", err)
	}

	// This should not raise an error.
	if err = wd.Wait(conditions.TitleContains("changed")); err != nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}

	// Testing unsuccessful selenium.Wait() call (this should raise error cause of timeout).
	err = wd.Get("http://localhost:3000/static")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/static: %v\n", err)
	}

	// This should raise an timeout error.
	if err = wd.WaitWithTimeout(conditions.TitleContains("changed"), 500*time.Millisecond); err == nil {
		t.Fatalf("wd.Wait() should raise an error, but it didn't.\n")
	}
}

func TestTitleNotContains(t *testing.T) {
	// Testing successful selenium.Wait() call.
	err := wd.Get("http://localhost:3000/title_change")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/title_change: %v\n", err)
	}

	// This should not raise an error.
	if err = wd.Wait(conditions.TitleNotContains("Selenium")); err != nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}

	// Testing unsuccessful selenium.Wait() call (this should raise error cause of timeout).
	err = wd.Get("http://localhost:3000/static")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/static: %v\n", err)
	}

	// This should raise an timeout error.
	if err = wd.WaitWithTimeout(conditions.TitleNotContains("Selenium"), 500*time.Millisecond); err == nil {
		t.Fatalf("wd.Wait() should raise an error, but it didn't.\n")
	}
}
