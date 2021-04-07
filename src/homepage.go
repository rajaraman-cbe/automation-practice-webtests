package pkg

var (
	signIn = "login"
)
//ClickSignIn function to click the sign in link from the homepage
func (a *AutomationPractice) ClickSignIn() error {
	err := a.FindByClass(signIn).Click()
	if check := a.CheckError(err); check != nil {
		return check
	}
	return nil
}
