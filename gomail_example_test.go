package gomail_test

import (
  "fmt";
  "time";
)

func Example_testGomailClient() {
  gomailObject, _ := gomail.NewGomail()
  
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
	
	fmt.Println(gomailObject)
  // Output: &{{GoMail Service example@gomail.com} [{Person 1 person1@gomail.com}] This is the email subject This is the fabulous body message Good Luck!!}
}
