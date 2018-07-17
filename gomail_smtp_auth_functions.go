package gomail

import (
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
func NewGoMailAuth(a *UserCredentials, host string) (unEncryptedAuth){
  return unEncryptedAuth {
     smtp.PlainAuth("",a.GetUsername(), a.GetPasswd(), host),
   }
}
