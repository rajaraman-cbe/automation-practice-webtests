package pkg

import (
	"github.com/sclevine/agouti"
	"strconv"
)

type Registration struct {
	User Title
	Addr Address
}

type Title struct {
	Fname string
	Lname string
	Pwd   string
	Date  int
	Month string
	Year  int
	email string
}

type Address struct {
	Fname       string
	Lname       string
	Company     string
	Addrs       string
	City        string
	State       string
	PostCode    int64
	Country     string
	HomePhone   int64
	MobilePhone int64
	Alias       string
}

var (
	title       = "id_gender1"
	fname       = "customer_firstname"
	lname       = "customer_lastname"
	email       = "//input[@id='email']"
	pwd         = "passwd"
	date        = "//select[@id='days']/option[contains(text(),"
	month       = "//select[@id='months']/option[contains(text(),"
	year        = "//select[@id='years']/option[contains(text(),"
	company     = "company"
	addr        = "address1"
	city        = "city"
	state       = "//select[@id='id_state']/option[contains(text(),"
	postcode    = "postcode"
	homePhone   = "phone"
	mobilePhone = "phone_mobile"
	alias       = "alias"
	register    = "//button[@id='submitAccount']"
	errorAlert  = "//div[@class='alert alert-danger']//li[contains(.,"
)
//InputUserRegistrationDetails inputs the registration details as specified in the test
func (a *AutomationPractice) InputUserRegistrationDetails(registration Registration) error {
	err := a.FindByID(title).Click()
	if check := a.CheckError(err); check != nil {
		return check
	}
	err = a.FindByID(fname).Fill(registration.User.Fname)
	if check := a.CheckError(err); check != nil {
		return check
	}
	err = a.FindByID(lname).Fill(registration.User.Lname)
	if check := a.CheckError(err); check != nil {
		return check
	}
	err = a.FindByID(pwd).Fill(registration.User.Pwd)
	if check := a.CheckError(err); check != nil {
		return check
	}
	err = a.FindByXPath(email).Tap(agouti.SingleTap)
	err = a.FindByXPath(date + strconv.Itoa(registration.User.Date) + ")]").Click()
	if check := a.CheckError(err); check != nil {
		return check
	}
	err = a.FindByXPath(month + "'" + registration.User.Month + "')]").Click()
	if check := a.CheckError(err); check != nil {
		return check
	}
	err = a.FindByXPath(year + strconv.Itoa(registration.User.Year) + ")]").Click()
	if check := a.CheckError(err); check != nil {
		return check
	}
	err = a.FindByID(company).Fill(registration.Addr.Company)
	if check := a.CheckError(err); check != nil {
		return check
	}
	err = a.FindByID(addr).Fill(registration.Addr.Addrs)
	if check := a.CheckError(err); check != nil {
		return check
	}
	err = a.FindByID(city).Fill(registration.Addr.Company)
	if check := a.CheckError(err); check != nil {
		return check
	}
	err = a.FindByID(company).Fill(registration.Addr.Company)
	if check := a.CheckError(err); check != nil {
		return check
	}
	err = a.FindByXPath(state + "'" + registration.Addr.State + "')]").Click()
	if check := a.CheckError(err); check != nil {
		return check
	}
	err = a.FindByID(postcode).Fill(strconv.FormatInt(registration.Addr.PostCode, 10))
	if check := a.CheckError(err); check != nil {
		return check
	}
	err = a.FindByID(homePhone).Fill(strconv.FormatInt(registration.Addr.HomePhone, 10))
	if check := a.CheckError(err); check != nil {
		return check
	}
	err = a.FindByID(mobilePhone).Fill(strconv.FormatInt(registration.Addr.MobilePhone, 10))
	if check := a.CheckError(err); check != nil {
		return check
	}
	err = a.FindByID(alias).Fill(registration.Addr.Alias)
	if check := a.CheckError(err); check != nil {
		return check
	}
	return nil
}
//ClickRegister submits the register button in the account creation page
func (a *AutomationPractice) ClickRegister() error {
	err := a.FindByXPath(register).Click()
	if check := a.CheckError(err); check != nil {
		return check
	}
	return nil
}
//IsCreateAccountErrorDisplayed verifies if the account error is displayed based on the error msg specified in the test
func (a *AutomationPractice) IsCreateAccountErrorDisplayed(errorMsg string) (bool) {
	displayed,_ := a.FindByXPath(errorAlert + "'" + errorMsg + "')]").Visible()
	return displayed
}
