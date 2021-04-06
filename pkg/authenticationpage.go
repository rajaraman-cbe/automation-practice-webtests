package pkg

var (
	emailAddress       = "email_create"
	submitRegistration = "SubmitCreate"
	createAccountError = "create_account_error"
	createAccountForm  = "account-creation_form"
)

func (a *AutomationPractice) EnterEmailAddressRegistration(email string) error {
	err := a.FindByID(emailAddress).Fill(email)
	if check := a.CheckError(err); check != nil {
		return check
	}
	return err
}

func (a *AutomationPractice) ClickSubmitRegistration() error {
	err := a.FindByID(submitRegistration).Submit()
	if check := a.CheckError(err); check != nil {
		return check
	}
	return err
}

func (a *AutomationPractice) IsAccountErrorDisplayed() bool {
	displayed, _ := a.FindByID(createAccountError).Visible()
	return displayed
}

func (a *AutomationPractice) IsAccountCreationFormDisplayed() bool {
	displayed, _ := a.FindByID(createAccountForm).Visible()
	return displayed
}

func (a *AutomationPractice) GetTextFromError() string {
	str, _ := a.FindByID(createAccountError).Text()
	return str
}
