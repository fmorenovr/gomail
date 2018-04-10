package gomail_test

import (
  "fmt"
  "github.com/jenazads/gomail"
)

func Example_sendingMessage() {
	gomailObject := gomail.NewGomail()
  
  // from
	gomailObject.Set("From", "example@gomail.com")
	gomailObject.Set("From_name", "GoMail Service")

  // to
	gomailObject.Set("To", "person1@gomail.com")
	gomailObject.Set("To_name", "Person 1")
	
	// subject
	gomailObject.Set("Subject", "This is the email subject")

  // body message
	gomailObject.Set("BodyMessage", "This is the fabulous body message\n\nGood Luck!!")
	
	// set a list of recipients
	gomailObject.SetListToIds([]gomail.Recipients{{"Person in list 1", "personlist_1@gomail.com"}, {"Person in list 2", "personlist_2@gomail.com"}})
	
	// add another recipient
	gomailObject.AddToIds("Person added", "personadded@gomail.com")

  // add other recipients
	gomailObject.AddListToIds([]gomail.Recipients{{"Person added in list 1", "personaddedlist_1@gomail.com"}, {"Person added in list 2", "personaddedlist_2@gomail.com"}})
	
	// delete any recipient
	gomailObject.DeleteToId("Email", "personlist_1@gomail.com")
	gomailObject.DeleteToId("Name", "Person to delete")

  // change info of any recipient
	gomailObject.ChangeToId("Email", "personadded@gomail.com", "personchangedMail@gomail.com")
	gomailObject.ChangeToId("Name", "Person added in list 1", "Person added in the list 1")
	
	smtpServer := gomail.NewSMTPServer()
	authUser := gomail.NewUserCredentials("gomail@aduncus.com", "******")
	
	err := smtpServer.SendMessage(authUser, gomailObject)
	fmt.Println("Error: ", err)
}
