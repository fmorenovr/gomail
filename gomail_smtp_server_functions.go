package gomail

import (
  "strings";
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

// new instance servername
func (s *SMTPServer) NewSMTPServerName(v string) (){
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
