# selenium-until-extra

A library that provides conditions for `WebDriver.Wait()` function from `github.com/tebeka/selenium` package.

### Motivation

Recently `WebDriver.Wait()` function was added to Golang's Selenium binding (I've implemented it and I've made a PR on that, so I kind of know how it's working). It allows waiting for some condition to be true.

There was a decision to implement only `WebDriver.Wait()` in this `github.com/tebeka/selenium` package, and leave the implementation of conditions to another library. So, here it is!

### How does it work

There is a `type Condition func (selenium.WebDriver) (bool, error)` and `WebDriver.Wait(cond Condition, timeout, interval time.Duration)` in Golang's Selenium binding. The `WebDriver.Wait()`'s implementaion is pretty simple: it starts an endless loop, and return either `nil` if the condition would evaluate to true, or `error` if there would be an error executing a condition or on timeout.

So, what this package does is provides functions that returns `Condition`s, so you would be able to pass it as an argument for `WebDriver.Wait()` function.

### How to use it (aka example)

First you `import "github.com/serge1peshcoff/selenium-go-conditions"`, and then you use `conditions` package. See `example.go` for more real-life example.

## Contribution

All issues and PRs are welcomed and appreciated!

If you want to suggest something new, you can make an issue about that, and we'll figure that out!
