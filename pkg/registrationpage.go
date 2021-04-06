package pkg

import (
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
}

type Address struct {
	Fname       string
	Lname       string
	Company     string
	Addr        string
	City        string
	State       string
	PostCode    string
	Country     string
	HomePhone   int64
	MobilePhone int64
	Alias       string
}

var (
	title       = "id_gender1"
	fname       = "customer_firstname"
	lname       = "customer_lastname"
	email       = "email"
	pwd         = "passwd"
	date        = "//*[@id='days']/option"
	month       = "//*[@id='months']/option"
	year        = "//*[@id='year']/option"
	addrFname   = "firstname"
	addrLname   = "lastname"
	company     = "company"
	addr        = "address1"
	city        = "city"
	state       = "id_state"
	postcode    = "postcode"
	country     = "id_country"
	homePhone   = "phone"
	mobilePhone = "phone_mobile"
	alias       = "alias"
)

func (a *AutomationPractice) InputUserRegistrationDetails(registration Registration) error {
	err := a.FindByID(title).Click()
	err = a.FindByID(fname).Fill(registration.User.Fname)
	err = a.FindByID(lname).Fill(registration.User.Lname)
	err = a.FindByID(pwd).Fill(registration.User.Pwd)
	err = a.AllByXPath(date).Select(strconv.Itoa(registration.User.Date))
	// if err != nil {
	// 	var rows *agouti.MultiSelection
	// 	rows = a.AllByID(date)
	// 	count, _ := rows.Count()
	// 	log.Print(count)
	// 	log.Fatal(err.Error())
	// }
	err = a.AllByXPath(month).Select(registration.User.Month)
	err = a.AllByXPath(year).Select(strconv.Itoa(registration.User.Year))
	return err
}
