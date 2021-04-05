package tests

import (
	pages "automation-practice-webtests/pkg"
	utils "automation-practice-webtests/utils"
	"fmt"
	"log"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
)

var baseURL = "http://automationpractice.com/index.php"

var _ = Describe("Registration", func() {
	var driver *agouti.WebDriver
	var page *pages.AutomationPractice

	var _ = BeforeSuite(func() {
		driver = agouti.ChromeDriver()
		if err := driver.Start(); err != nil {
			log.Fatal("Failed to start driver..=>" + err.Error())
		}
	})

	var _ = AfterSuite(func() {
		if driver != nil {
			if err := driver.Stop(); err != nil {
				log.Fatal("Failed to stop driver..=>" + err.Error())
			}
		}
	})

	BeforeEach(func() {
		page = pages.NewAutomationPracticePage(driver)
		err := page.Open(baseURL)
		if err != nil {
			log.Fatal("Unable to open the Page...=>" + err.Error())
		}
	})

	Context("With an email address on the authentication page", func() {
		JustBeforeEach(func() {
			When("User Sign in from Home page", func() {
				page.ClickSignIn()
			})
		})

		It("Should redirect user to the secure authentication page", func() {
			p, _ := page.Title()
			fmt.Println("Page title here is:" + p)
			Expect(page.Title()).Should(Equal("Login - My Store"))
		})

		When("User enters a valid email address in the registration form", func() {
			JustBeforeEach(func() {
				page.EnterEmailAddressRegistration(utils.GenerateRandomString() + "@xyz.com")
				page.ClickSubmitRegistration()
			})
			It("Should redirect user to the registration page", func() {
				Eventually(func() bool {
					return page.IsAccountCreationFormDisplayed()
				}, 2*time.Second, 10*time.Millisecond).Should(BeTrue())
			})
		})

		When("User enters a blank email address in the registration form", func() {
			JustBeforeEach(func() {
				page.EnterEmailAddressRegistration("")
				page.ClickSubmitRegistration()
			})
			It("Should error as invalid email address", func() {
				Eventually(func() bool {
					return page.IsAccountErrorDisplayed()
				}, 2*time.Second, 10*time.Millisecond).Should(BeTrue())
				Expect(page.GetTextFromError()).To(Equal("Invalid email address."))
			})
		})

		When("User enters an email that is already registered in the registration form", func() {
			JustBeforeEach(func() {
				page.EnterEmailAddressRegistration("xyz@gmail.com")
				page.ClickSubmitRegistration()
			})
			It("Should error as previously registered", func() {
				Eventually(func() bool {
					return page.IsAccountErrorDisplayed()
				}, 2*time.Second, 10*time.Millisecond).Should(BeTrue())
				Expect(page.GetTextFromError()).To(ContainSubstring("email address has already been registered"))
			})
		})
	})
})
