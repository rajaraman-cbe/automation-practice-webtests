package pkg

var (
	signIn = "login"
)

func (a *AutomationPractice) ClickSignIn() error {
	err := a.FindByClass(signIn).Click()
	if check := a.CheckError(err); check != nil {
		return check
	}
	return nil
}
