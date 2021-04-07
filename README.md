# automation-practice-webtests

UI automation tests using Go - Agouti, Ginkgo, Gomega

Language - Go (1.16)
Frameworks - Agouti, Ginkgo, Gomega
Website - automationpractice.com

# Setup & pre-requisites to run on local machine:
 - If not already installed, install Go for the specifc OS from here https://golang.org/dl/
 - If not already installed, install ChromeDriver for the specific OS & version of chrome from https://chromedriver.chromium.org/ - For the purpose of this test
   used chromedriver v89.xx 
      - Mac - brew install chromedriver 
      - Win - download the driver and add them to the %PATH% under the System Environment Variables
- git clone this project to a repository
- If running from a terminal, go to the location of the repo and type 'go mod download' (this will downlad all of the dependencies this repo uses)
- cd tests and type 'ginkgo'

# Scenarios covered:
- Test with invalid email
- Test with email already registered
- Test with valid email but missing mandatory registration details
- Test with valid email and successful registration
- Test login functionality with registered email and password

# Notes
- The pre-requisite detailed above assumes that the user is aware of the Go env, Webdriver setup etc.
- The tests developed follow the traditional Page object pattern where each page worked by the web driver is handled in a separate go file with the locators and the   specific actions on that page.
- Some common functions are separated out of the individual pages and implemented in the impl.go file
- The top level registration_test.go creates a single instance of the webdriver before the suite begins and creates a new page that access the home page and follow 
  the steps described in each of the tests

If i had missed any step in this process that is not clear or have missed any, happy to walk through the code!

Thanks.
