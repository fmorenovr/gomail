package gomail

import (
  "strings";
  "net/smtp";
  "crypto/tls";
)

// SMTP struct with host & port
type SMTPServer struct {
  host string
  port string
}

// new instance by defulat of SMTP server
func NewSMTPServer()(*SMTPServer){
  return &SMTPServer{
           host: "smtp.gmail.com",
           port: "465",
         }
}

// new instance by defulat of SMTP server
func NewSMTPServerOptions(mhost, mport string)(*SMTPServer){
  return &SMTPServer{
           host: mhost,
           port: mport,
         }
}

// set servername
func (s *SMTPServer) SetServerName(v string) (){
  hostport := strings.Split(v, ":")
  s.SetServerHost(hostport[0])
  s.SetServerPort(hostport[1])
}

// set tls config to smtp server
func (s *SMTPServer) SMTPServerTLS()(*tls.Config){
 return &tls.Config{
          InsecureSkipVerify: true,
          ServerName: s.GetServerHost(),
        }
}

// get servername
func (s *SMTPServer) GetServerName() string {
  return s.host + ":" + s.port
}

// set server host
func (s *SMTPServer) SetServerHost(v string) (){
  s.host = v
}

// get server host
func (s *SMTPServer) GetServerHost() string {
  return s.host
}

// set server port
func (s *SMTPServer) SetServerPort(v string) () {
  s.port = v
}

// get server port
func (s *SMTPServer) GetServerPort() string {
  return s.port
}

// Send Message
func (s *SMTPServer) SendMessage(u *UserCredentials, o *Gomail)(error){
  err := o.VerifyEmails()
  if err != nil{
    return err
  }
  messageToSend := o.PrepareMessage()
  servername := s.GetServerName()
  host := s.GetServerHost()

  //build an auth
  auth := NewGomailAuth(u, host)

  // TLS config
  tlsconfig := s.SMTPServerTLS()
  conn, err := tls.Dial("tcp", servername, tlsconfig)
  if err != nil {
    return GomailErrBadHostPortServer
  }

  // create new client
  c, err := smtp.NewClient(conn, host)
  if err != nil {
    return GomailErrSMTPClient
  }

  defer c.Quit()

  // step 1: Use Auth
  if err = c.Auth(auth); err != nil {
    return GomailErrBadCredentials
  }

  // step 2: add all from and to
  if err = c.Mail(o.GetFromAddress()); err != nil {
    return GomailErrSyntaxSender
  }

  for _, k := range o.GetToIds() {
    if err = c.Rcpt(o.GetToIdAddress(k)); err != nil {
      return GomailErrSyntaxRecipients
    }
  }

  // Data
  w, err := c.Data()
  if err != nil {
    return GomailErrMessageData
  }

  _, err = w.Write([]byte(messageToSend))
  if err != nil {
    return GomailErrSendMessage
  }

  err = w.Close()
  if err != nil {
    return GomailErrWriteClientClose
  }
  return nil
}
