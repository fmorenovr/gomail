package gomail

import (
  "regexp";
  "net/smtp";
)

// instance of smtp.Auth
type unEncryptedAuth struct {
  smtp.Auth
}

// initialize encrypted smtp auth
func (a unEncryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
  s := *server
  s.TLS = true
  return a.Auth.Start(&s)
}

// Instance an auth with a specific host
func NewGomailAuth(a *UserCredentials, host string) (unEncryptedAuth){
  return unEncryptedAuth {
     smtp.PlainAuth("",a.GetUsername(), a.GetPasswd(), host),
   }
}

// verify format of email
func VerifyFormatEmail(email string) (bool){
  Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
 	return Re.MatchString(email)
}
