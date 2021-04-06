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
		page = pages.NewAutomationPracticePage(driver)
	})

	// var _ = AfterSuite(func() {
	// 	err := page.Close()
	// 	if err != nil {
	// 		log.Fatal("Unable to close the Page...=>" + err.Error())
	// 	}
	// 	if driver != nil {
	// 		if err := driver.Stop(); err != nil {
	// 			log.Fatal("Failed to stop driver..=>" + err.Error())
	// 		}
	// 	}
	// })

	BeforeEach(func() {
		err := page.Open(baseURL)
		if err != nil {
			log.Fatal("Unable to open the Page...=>" + err.Error())
		}
	})

	Context("With any email address on the authentication page", func() {
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
				}, 5*time.Second, 10*time.Millisecond).Should(BeTrue())
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
				}, 5*time.Second, 10*time.Millisecond).Should(BeTrue())
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
				}, 5*time.Second, 10*time.Millisecond).Should(BeTrue())
				Expect(page.GetTextFromError()).To(ContainSubstring("email address has already been registered"))
			})
		})
	})

	FContext("With a valid email address and user profile creation", func() {
		JustBeforeEach(func() {
			When("User Sign in from Home page", func() {
				page.ClickSignIn()
			})
			When("User enters a valid email address in the registration form", func() {
				page.EnterEmailAddressRegistration(utils.GenerateRandomString() + "@xyz.com")
				page.ClickSubmitRegistration()
			})
		})

		When("User enters valid Titular  details", func() {
			It("Should be accepted", func() {
				Eventually(func() bool {
					return page.IsAccountCreationFormDisplayed()
				}, 5*time.Second, 10*time.Millisecond).Should(BeTrue())

				registration := pages.Registration{
					User: pages.Title{
						Fname: "Rajaraman",
						Lname: "Mahalingam",
						Pwd:   utils.GenerateRandomString(),
						Date:  15,
						Month: "October",
						Year:  1981,
					},
				}
				page.InputUserRegistrationDetails(registration)
			})
		})

	})
})
