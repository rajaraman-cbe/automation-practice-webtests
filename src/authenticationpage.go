package pkg

var (
	emailAddressRegistration = "email_create"
	submitRegistration       = "SubmitCreate"
	createAccountError       = "create_account_error"
	createAccountForm        = "account-creation_form"
	emailAddressLogin        = "email"
	passwd                   = "passwd"
	signInLogin              = "SubmitLogin"
)

//EnterEmailAddressRegistration enters the provided email address in the email field of the registration form
func (a *AutomationPractice) EnterEmailAddressRegistration(email string) error {
	err := a.FindByID(emailAddressRegistration).Fill(email)
	if check := a.CheckError(err); check != nil {
		return check
	}
	return err
}

//ClickSubmitRegistration submits the registration button on the registration form
func (a *AutomationPractice) ClickSubmitRegistration() error {
	err := a.FindByID(submitRegistration).Submit()
	if check := a.CheckError(err); check != nil {
		return check
	}
	return err
}

//IsAccountErrorDisplayed checks if the error is displayed in the registration form
func (a *AutomationPractice) IsAccountErrorDisplayed() bool {
	displayed, _ := a.FindByID(createAccountError).Visible()
	return displayed
}

//IsIsAccountCreationFormDisplayed checks if the account creation form is displayed
func (a *AutomationPractice) IsAccountCreationFormDisplayed() bool {
	displayed, _ := a.FindByID(createAccountForm).Visible()
	return displayed
}

//GetTextFromError retrieves the text() from the errors displayed
func (a *AutomationPractice) GetTextFromError() string {
	str, _ := a.FindByID(createAccountError).Text()
	return str
}
//EnterLoginCredentials enters the username and password in the login form
func (a *AutomationPractice) EnterLoginCredentials(username string, password string) error {
	err := a.FindByID(emailAddressLogin).Fill(username)
	if check := a.CheckError(err); check != nil {
		return check
	}
	err = a.FindByID(passwd).Fill(password)
	if check := a.CheckError(err); check != nil {
		return check
	}
	err = a.FindByID(signInLogin).Click()
	if check := a.CheckError(err); check != nil {
		return check
	}
	return nil
}
