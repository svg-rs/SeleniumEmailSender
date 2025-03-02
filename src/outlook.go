package outlook

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

var (
	infoLogger    = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger   = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
)

func Send(subject, message, username, password string,
	recipients []string) (err error) {
	if message == "" || username == "" || password == "" || len(
		recipients) == 0 || subject == "" {
		errorLogger.Println("Message or username or password is missing!")
		return
	}
	if len(recipients) == 0 {
		errorLogger.Println("Recipients is missing!")
		return
	}
	infoLogger.Println("Starting Outlook chromedriver...")
	var service *selenium.Service
	service, err = selenium.NewChromeDriverService("chromedriver/chromedriver.exe", 4444)
	if err != nil {
		return err
	}
	defer service.Stop()
	var driver selenium.WebDriver
	driver, err = selenium.NewRemote(
		selenium.Capabilities{"browserName": "chrome"},
		"http://localhost:4444/wd/hub",
	)
	if err != nil {
		return err
	}
	defer driver.Quit()
	infoLogger.Println("Outlook chromedriver started!")
	infoLogger.Println("Opening Outlook...")
	err = driver.Get("https://go.microsoft.com/fwlink/p/?LinkID=2125442&deeplink=owa%2F")
	if err != nil {
		return err
	}
	infoLogger.Println("Outlook opened!")
	infoLogger.Println("Logging in...")
	time.Sleep(2 * time.Second)
	var loginButton selenium.WebElement
	loginButton, err = driver.FindElement(selenium.ByXPATH, "//input[@id='i0116']")
	if err != nil {
		return err
	}
	loginButton.Click()
	infoLogger.Println("Login page loaded!")
	time.Sleep(2 * time.Second)
	infoLogger.Println("Waiting for email input...")
	var emailField selenium.WebElement
	emailField, err = driver.FindElement(selenium.ByID, "i0116")
	if err != nil {
		return err
	}
	err = emailField.SendKeys(username)
	if err != nil {
		return err
	}
	infoLogger.Println("Email input filled!")
	time.Sleep(2 * time.Second)
	var loginNextButton selenium.WebElement
	loginNextButton, err = driver.FindElement(selenium.ByID, "idSIButton9")
	if err != nil {
		return err
	}
	loginNextButton.Click()
	infoLogger.Println("Login next button clicked!")
	time.Sleep(2 * time.Second)
	infoLogger.Println("Waiting for password input...")
	time.Sleep(7 * time.Second)
	var passwordField selenium.WebElement
	passwordField, err = driver.FindElement(selenium.ByID, "i0118")
	if err != nil {
		return err
	}
	err = passwordField.SendKeys(password)
	if err != nil {
		return err
	}
	infoLogger.Println("Password input filled!")
	time.Sleep(2 * time.Second)
	infoLogger.Println("Waiting for sign in button...")
	var signInButton selenium.WebElement
	signInButton, err = driver.FindElement(selenium.ByID, "idSIButton9")
	if err != nil {
		return err
	}
	time.Sleep(3 * time.Second)
	signInButton.Click()
	infoLogger.Println("Sign in button clicked!")
	time.Sleep(2500 * time.Millisecond)
	infoLogger.Println("Waiting for accept button...")
	var acceptButton selenium.WebElement
	acceptButton, err = driver.FindElement(selenium.ByID, "acceptButton")
	if err != nil {
		return err
	}
	acceptButton.Click()
	infoLogger.Println("Accept button clicked!")
	time.Sleep(3 * time.Second)
	infoLogger.Println("Waiting for new message button...")
	var newMessageButton selenium.WebElement
	newMessageButton, err = driver.FindElement(selenium.ByXPATH, "//button[contains(@class, 'splitPrimaryButton')]")
	if err != nil {
		return err
	}
	err = newMessageButton.Click()
	if err != nil {
		return err
	}
	infoLogger.Println("New message button clicked!")
	time.Sleep(8 * time.Second)
	infoLogger.Println("Waiting for message recipient...")
	var messageRecipient selenium.WebElement
	messageRecipient, err = driver.FindElement(selenium.ByCSSSelector, "div[aria-label='To']")
	if err != nil {
		return err
	}
	err = messageRecipient.Click()
	if err != nil {
		return err
	}

	for i, recipient := range recipients {
		infoLogger.Printf("Adding recipient: %s, %v\n", recipient, i)
		err = messageRecipient.SendKeys(recipient + ";")
		if err != nil {
			return err
		}
		var sleepTime time.Duration = time.Duration(rand.Intn(201)+100) * time.
			Millisecond
		time.Sleep(sleepTime)
	}

	err = messageRecipient.SendKeys(selenium.TabKey)
	if err != nil {
		return err
	}

	time.Sleep(200 * time.Millisecond)

	infoLogger.Println("Waiting for subject input...")
	var subjectField selenium.WebElement
	subjectField, err = driver.FindElement(selenium.ByCSSSelector, "input[aria-label='Add a subject']")
	err = subjectField.SendKeys(subject)
	if err != nil {
		return err
	}
	infoLogger.Println("Subject input filled!")
	time.Sleep(500 * time.Millisecond)
	infoLogger.Println("Waiting for message input...")
	var messageField selenium.WebElement
	messageField, err = driver.FindElement(selenium.ByXPATH, "//div[contains(@class, 'dFCbN') and contains(@class, 'dnzWM')]\n")
	if err != nil {
		return err
	}
	err = messageField.Click()
	if err != nil {
		return err
	}
	err = messageField.SendKeys(message)
	if err != nil {
		return err
	}
	infoLogger.Println("Message input filled!")
	time.Sleep(500 * time.Millisecond)
	infoLogger.Println("Waiting for send button...")
	var sendButton selenium.WebElement
	sendButton, err = driver.FindElement(selenium.ByCSSSelector, "button[aria-label='Send']")
	if err != nil {
		return err
	}
	err = sendButton.Click()
	if err != nil {
		return err
	}
	infoLogger.Println("Send button clicked!")

	var continu string
	fmt.Println("Press enter to continue and restart...")
	fmt.Scanln(&continu)
	if continu == "" {
		return Send(subject, message, username, password, recipients)
	}
	return nil
}
