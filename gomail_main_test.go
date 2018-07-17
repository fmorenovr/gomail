package gomail_test

import (
  "fmt"
  "github.com/jenazads/gomail"
)

func Example_sendingMessage() {
  gomailObject, _ := gomail.NewGoMail()

  // credentials
  gomailObject.Set("Username", "aduncus@gomail.com")
  gomailObject.Set("Password", "********")
  
  // servername
  gomailObject.Set("Servername", "smtp.gmail.com:465")

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
  
  // repeat email with differents names
  gomailObject.AddToIds("Person added copy", "personadded@gomail.com")
  gomailObject.AddToIds("Person added no, im the copy", "personadded@gomail.com")
  gomailObject.AddToIds("Person added no, im the one", "personadded@gomail.com")

  // add other recipients
  gomailObject.AddListToIds([]gomail.Recipients{{"Person added in list 1", "personaddedlist_1@gomail.com"}, {"Person added in list 2", "personaddedlist_2@gomail.com"}})

  // delete any recipient
  gomailObject.DeleteToId("Email", "personlist_1@gomail.com")
  gomailObject.DeleteToId("Name", "Person to delete")

  // change info of any recipient
  gomailObject.ChangeToId("Email", "personadded@gomail.com", "personchangedMail@gomail.com")
  gomailObject.ChangeToId("Name", "Person added in list 1", "Person added in the list 1")

  err := gomailObject.SendMessage()
  fmt.Println("Error: ", err)
}
