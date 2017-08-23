package conditions_test

import (
	"testing"
	"time"

	conditions "github.com/serge1peshcoff/selenium-go-conditions"
)

func TestURLIs(t *testing.T) {
	// Testing successful selenium.Wait() call.
	err := wd.Get("http://localhost:3000/redirect")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/redirect: %v\n", err)
	}

	// This should not raise an error.
	if err = wd.Wait(conditions.URLIs("http://localhost:3000/static")); err != nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}

	// Testing unsuccessful selenium.Wait() call (this should raise error cause of timeout).
	err = wd.Get("http://localhost:3000/static")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/static: %v\n", err)
	}

	// This should raise an timeout error.
	if err = wd.WaitWithTimeout(conditions.URLIs("http://localhost:3000/redirect"), 500*time.Millisecond); err == nil {
		t.Fatalf("wd.Wait() should raise an error, but it didn't.\n")
	}
}
func TestURLIsNot(t *testing.T) {
	// Testing successful selenium.Wait() call.
	err := wd.Get("http://localhost:3000/redirect")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/redirect: %v\n", err)
	}

	// This should not raise an error.
	if err = wd.Wait(conditions.URLIsNot("http://localhost:3000/redirect")); err != nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}

	// Testing unsuccessful selenium.Wait() call (this should raise error cause of timeout).
	err = wd.Get("http://localhost:3000/static")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/static: %v\n", err)
	}

	// This should raise an timeout error.
	if err = wd.WaitWithTimeout(conditions.URLIsNot("http://localhost:3000/static"), 500*time.Millisecond); err == nil {
		t.Fatalf("wd.Wait() should raise an error, but it didn't.\n")
	}
}

func TestURLContains(t *testing.T) {
	// Testing successful selenium.Wait() call.
	err := wd.Get("http://localhost:3000/redirect")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/redirect: %v\n", err)
	}

	// This should not raise an error.
	if err = wd.Wait(conditions.URLContains("static")); err != nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}

	// Testing unsuccessful selenium.Wait() call (this should raise error cause of timeout).
	err = wd.Get("http://localhost:3000/static")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/static: %v\n", err)
	}

	// This should raise an timeout error.
	if err = wd.WaitWithTimeout(conditions.URLContains("redirect"), 500*time.Millisecond); err == nil {
		t.Fatalf("wd.Wait() should raise an error, but it didn't.\n")
	}
}

func TestURLNotContains(t *testing.T) {
	// Testing successful selenium.Wait() call.
	err := wd.Get("http://localhost:3000/redirect")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/redirect: %v\n", err)
	}

	// This should not raise an error.
	if err = wd.Wait(conditions.URLNotContains("redirect")); err != nil {
		t.Fatalf("Error while executing wd.Wait(): %v\n", err)
	}

	// Testing unsuccessful selenium.Wait() call (this should raise error cause of timeout).
	err = wd.Get("http://localhost:3000/static")
	if err != nil {
		t.Fatalf("Cannot get http://localhost:3000/static: %v\n", err)
	}

	// This should raise an timeout error.
	if err = wd.WaitWithTimeout(conditions.URLNotContains("static"), 500*time.Millisecond); err == nil {
		t.Fatalf("wd.Wait() should raise an error, but it didn't.\n")
	}
}
