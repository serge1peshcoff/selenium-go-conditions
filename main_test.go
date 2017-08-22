package conditions_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"
)

var cmd *exec.Cmd

func TestMain(m *testing.M) {
	err := runSelenium()
	if err != nil {
		fmt.Printf("Error staring Selenium.\n")
		panic(err)
	}

	fmt.Printf("Selenium started.\n")

	go startServer()

	retCode := m.Run()

	err = stopSelenium()
	if err != nil {
		fmt.Printf("Error stopping Selenium.\n")
		panic(err)
	}

	fmt.Printf("Selenium stopped.\n")

	os.Exit(retCode)

}

func runSelenium() error {
	// Running selenium-standalone.
	cmd = exec.Command("java", "-jar", "testing/selenium-server-standalone-3.5.1.jar")
	err := cmd.Start()
	if err != nil {
		return err
	}

	// Waiting for selenium-standalone server to run.
	for i := 0; i < 30; i++ {
		time.Sleep(time.Second)
		resp, err := http.Get("http://localhost:4444/status")
		if err == nil {
			resp.Body.Close()
			switch resp.StatusCode {
			// Selenium <3 returned Forbidden and BadRequest. ChromeDriver and
			// Selenium 3 return OK.
			case http.StatusForbidden, http.StatusBadRequest, http.StatusOK:
				return nil
			}
		}
	}
	return fmt.Errorf("server did not respond on port %d", 4444)

}

func handler(writer http.ResponseWriter, request *http.Request) {
	pwd, _ := os.Getwd()

	path := request.URL.Path
	bytes, err := ioutil.ReadFile(pwd + "/testing/static" + path + ".html")

	if err != nil {
		http.NotFound(writer, request)
		return
	}

	fmt.Fprintf(writer, string(bytes))
}

func startServer() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func stopSelenium() error {
	err := cmd.Process.Kill()
	return err
}
