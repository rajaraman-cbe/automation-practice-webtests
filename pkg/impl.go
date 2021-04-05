package pkg

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sclevine/agouti"
)

var driver *agouti.WebDriver

type AutomationPractice struct {
	*agouti.Page
}

func NewAutomationPracticePage(driver *agouti.WebDriver) *AutomationPractice {
	page, err := driver.NewPage()
	if err != nil {
		log.Fatal("Failed to open Page")
	}
	return &AutomationPractice{page}
}

func (a *AutomationPractice) Open(url string) error {
	err := a.Navigate(url)
	if check := a.CheckError(err); check != nil {
		return check
	}
	return nil
}

func (a *AutomationPractice) Close() error {
	err := a.CloseWindow()
	if check := a.CheckError(err); check != nil {
		return check
	}
	return nil
}

func (a *AutomationPractice) TakeScreenShot(name string) (string, error) {
	folder := "../screencaptures/"
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.Mkdir(folder, os.FileMode(0522))
	}
	path := fmt.Sprintf("%s%s", folder, fmt.Sprintf("%s%d.jpg", name, time.Now().UnixNano()))
	err := a.Screenshot(path)
	if err != nil {
		log.Print("Unable to take screenshot of err..." + err.Error())
	}
	return path, nil
}

func (a *AutomationPractice) CheckError(e error) error {
	if e != nil {
		path, _ := a.TakeScreenShot(e.Error())
		return fmt.Errorf(fmt.Sprintf("%v %s", e.Error(), path))
	}
	return nil
}
