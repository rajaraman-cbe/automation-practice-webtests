package pkg

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sclevine/agouti"
)

type AutomationPractice struct {
	*agouti.Page
}

//NewAutomationPracticePage creates a new page using the webdriver

func NewAutomationPracticePage(driver *agouti.WebDriver) *AutomationPractice {
	page, err := driver.NewPage()
	if err != nil {
		log.Fatal("Failed to open Page")
	}
	return &AutomationPractice{page}
}

//Open Navigate the created page to a specified URL
func (a *AutomationPractice) Open(url string) error {
	err := a.Navigate(url)
	if check := a.CheckError(err); check != nil {
		return check
	}
	return nil
}

//Close closes the current window
func (a *AutomationPractice) Close() error {
	err := a.CloseWindow()
	if check := a.CheckError(err); check != nil {
		return check
	}
	return nil
}
//TakeScreenShot takes a jpg image of the current window
func (a *AutomationPractice) TakeScreenShot(name string) (string, error) {
	folder := "../screencaptures/"
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		_ = os.Mkdir(folder, os.FileMode(0522))
	}
	path := fmt.Sprintf("%s%s", folder, fmt.Sprintf("%s%d.jpg", name, time.Now().UnixNano()))
	err := a.Screenshot(path)
	if err != nil {
		log.Print("Unable to take screenshot of err..." + err.Error())
	}
	return path, nil
}

//CheckError checks if an error occured and takes a screenshot
func (a *AutomationPractice) CheckError(e error) error {
	if e != nil {
		path, _ := a.TakeScreenShot(e.Error())
		return fmt.Errorf(fmt.Sprintf("%v %s", e.Error(), path))
	}
	return nil
}

//GetPageTitle gets the page title of the current page
func (a *AutomationPractice) GetPageTitle() (string) {
	title, err := a.Title()
	if check := a.CheckError(err); check != nil {
		return check.Error()
	}
	return title
}