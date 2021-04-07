package tests

import (
	pages "automation-practice-webtests/pkg"
	utils "automation-practice-webtests/utils"
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

	var _ = AfterSuite(func() {
		err := page.Close()
		if err != nil {
			log.Fatal("Unable to close the Page...=>" + err.Error())
		}
		if driver != nil {
			if err := driver.Stop(); err != nil {
				log.Fatal("Failed to stop driver..=>" + err.Error())
			}
		}
	})

	BeforeEach(func() {
		err := page.Open(baseURL)
		if err != nil {
			log.Fatal("Unable to open the Page...=>" + err.Error())
		}
		_ = page.ClickSignIn()
	})

	Context("With any email address on the Authentication page", func() {
		JustBeforeEach(func() {
			_ = page.EnterEmailAddressRegistration(utils.GenerateRandomString(10) + "@xyz.com")
			_ = page.ClickSubmitRegistration()
		})

		When("User enters a blank email address in the registration form", func() {
			JustBeforeEach(func() {
				_ = page.EnterEmailAddressRegistration("")
				_ = page.ClickSubmitRegistration()
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
				_ = page.EnterEmailAddressRegistration("xyz@gmail.com")
				_ = page.ClickSubmitRegistration()
			})
			It("Should error as previously registered", func() {
				Eventually(func() bool {
					return page.IsAccountErrorDisplayed()
				}, 5*time.Second, 10*time.Millisecond).Should(BeTrue())
				Expect(page.GetTextFromError()).To(ContainSubstring("email address has already been registered"))
			})
		})
	})

	Context("With a valid email address on the Authentication page", func() {
		emailAddr := utils.GenerateRandomString(10) + "@xyz.com"
		JustBeforeEach(func() {
			When("User enters a valid email address in the registration form", func() {
				_ = page.EnterEmailAddressRegistration(emailAddr)
				_ = page.ClickSubmitRegistration()
			})
		})

		When("User clicks register button without submitting few mandatory details", func() {
			It("Should error and advice user appropriately", func() {
				Eventually(func() bool {
					return page.IsAccountCreationFormDisplayed()
				}, 5*time.Second, 10*time.Millisecond).Should(BeTrue())
				err := page.ClickRegister()
				Expect(err).To(BeNil())
				Expect(page.IsCreateAccountErrorDisplayed("lastname is required")).To(BeTrue())
				Expect(page.IsCreateAccountErrorDisplayed("firstname is required")).To(BeTrue())
				Expect(page.IsCreateAccountErrorDisplayed("passwd is required")).To(BeTrue())
				Expect(page.IsCreateAccountErrorDisplayed("address1 is required")).To(BeTrue())
				Expect(page.IsCreateAccountErrorDisplayed("city is required")).To(BeTrue())
				Expect(page.IsCreateAccountErrorDisplayed("country requires you to choose a State")).To(BeTrue())
			})
		})

		When("User enters valid Registration details and submits", func() {
			It("Should be accepted and registered successfully", func() {
				Eventually(func() bool {
					return page.IsAccountCreationFormDisplayed()
				}, 5*time.Second, 10*time.Millisecond).Should(BeTrue())

				registration := pages.Registration{
					User: pages.Title{
						Fname: "Rajaraman",
						Lname: "Mahalingam",
						Pwd:   utils.GenerateRandomString(10),
						Date:  15,
						Month: "October",
						Year:  1981,
					},
					Addr: pages.Address{
						Company:     utils.GenerateRandomString(10) + " Company",
						Addrs:       "123, ABC street",
						City:        "SAN JOSE",
						State:       "California",
						PostCode:    utils.GenerateRandomNumber(5),
						HomePhone:   utils.GenerateRandomNumber(10),
						MobilePhone: utils.GenerateRandomNumber(10),
						Alias:       "Home",
					},
				}
				err := page.InputUserRegistrationDetails(registration)
				Expect(err).To(BeNil())
				err = page.ClickRegister()
				Expect(err).To(BeNil())
				Eventually(func() string {
					return page.GetPageTitle()
				}, 5*time.Second, 10*time.Millisecond).Should(Equal("My account - My Store"))
			})
		})
	})
})
