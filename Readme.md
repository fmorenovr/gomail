# golang + Mail = GoMail (Gomail)

goMail (Golang for SMTP) is a Golang implementation for Simple Mail Transfer Protocol service.
You can see an extended doc in [godocs](https://godoc.org/github.com/Jenazads/goMail).

## SMTP

SMTP is a protocol that you can send secure text messages trough the internet.  
See more info [here](https://en.wikipedia.org/wiki/Simple_Mail_Transfer_Protocol).

## goMail

* First, You should download my library:

      go get github.com/jenazads/gomail/

* Then, you should use for differents Web Frameworks in Go.
        
    * First, Create a gomail object:
    
            gomailObject := gomail.NewGomail()
  
            // credentials
            gomailObject.Set("Username", "aduncus@gomail.com")
            gomailObject.Set("Password", "********")
      
            // servername
            gomailObject.Set("Servername", "smtp.gmail.com:465")
            
            // from
            gomailObject.Set("From", "example@gomail.com")
            gomailObject.Set("From_name", "gomail Service")

            // to
            gomailObject.Set("To", "person1@gomail.com")
            gomailObject.Set("To_name", "Person 1")
	
            // subject
            gomailObject.Set("Subject", "This is the email subject")

            // body message
            gomailObject.Set("BodyMessage", "This is the fabulous body message\n\nGood Luck!!")
    
* Finally, send a message:
        
       err := smtpServer.SendMessage(authUser, gomailObject)
       fmt.Println("Error: ", err)


