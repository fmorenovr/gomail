package gomail

import (
  "fmt";
  "bytes";
  "strings";
  "net/mail";
  "net/smtp";
  "crypto/tls";
  "html/template";
)

// GoMail object
type GoMail struct {
  from        mail.Address
  to          []mail.Address
  subject     string
  bodyMessage string
  contentType string
  username    string
  password    string
  servername  SMTPServer
}

// Create a New Gomail instance
func NewGoMail()(*GoMail, error){
  return &GoMail{
           mail.Address{"", ""},
           nil,
           "",
           "",
           `text/html; charset="UTF-8"`,
           "",
           "",
           SMTPServer{"smtp.gmail.com","465"}}, nil
}

// menu to ser
func (o *GoMail) Set(key, value string)(){
  switch key{
    case "From":
      o.SetFromAddress(value)
    case "From_name":
      o.SetFromName(value)
    case "To":
      o.SetToIdsAddress(value)
    case "To_name":
      o.SetToIdsName(value)
    case "Subject":
      o.SetSubject(value)
    case "BodyMessage":
      o.SetBodyMessage(value)
    case "Content-Type":
      o.SetContentType(value)
    case "Username":
      o.SetUsername(value)
    case "Password":
      o.SetPassword(value)
    case "Servername":
      o.SetServerName(value)
  }
}

// find by name
func (o *GoMail) FindByEmail(email string)(aux int){
  aux = -1
  for k, v:= range o.GetToIds(){
    if v.Address == email {
      aux = k
      break
    }
    
  }
  return aux
}

// find by email
func (o *GoMail) FindByName(name string)(aux int){
  aux = -1
  for k, v:= range o.GetToIds(){
    if v.Name == name {
      aux = k
      break
    }
  }
  return aux
}

// menu to change
func (o *GoMail) ChangeToId(by, original, v string)(){
  switch by {
    case "Name":
      o.changeToIdName(original, v)
    case "name":
      o.changeToIdName(original, v)
    case "Email":
      o.changeToIdEmail(original, v)
    case "email":
      o.changeToIdEmail(original, v)
  }
}

// set specific recipient name
func (o *GoMail) changeToIdName(original, v string) (){
  aux := o.FindByName(original)
  if aux == -1 {
    return
  } else {
    o.to[aux].Name = v
  }
}

// set specific recipient address
func (o *GoMail) changeToIdEmail(original, v string)(){
  aux := o.FindByEmail(strings.ToLower(original))
  if aux == -1 {
    return
  } else {
    o.to[aux].Address = strings.ToLower(v)
  }
}

// menu to delete
func (o *GoMail) DeleteToId(by, original string)(){
  switch by {
    case "Name":
      o.deleteToIdByName(original)
    case "name":
      o.deleteToIdByName(original)
    case "Email":
      o.deleteToIdByEmail(original)
    case "email":
      o.deleteToIdByEmail(original)
  }
}

// delete specific recipient name
func (o *GoMail) deleteToIdByName(original string)(){
  i := o.FindByName(original)
  if i == -1 {
    return
  } else {
    o.to = o.to[:i+copy(o.to[i:], o.to[i+1:])]
  }
}

// delete specific recipient address
func (o *GoMail) deleteToIdByEmail(original string)(){
  i := o.FindByEmail(strings.ToLower(original))
  if i == -1 {
    return
  } else {
    o.to = o.to[:i+copy(o.to[i:], o.to[i+1:])]
  }
}

// get specific recipient address
func (o *GoMail) GetToIdAddress(toId mail.Address)(string){
  return toId.Address
}

// get specific recipient name
func (o *GoMail) GetToIdName(toId mail.Address)(string){
  return toId.Name
}

// set unique recipient name
func (o *GoMail) SetToIdsName(v string)(){
  if o.to == nil || len(o.to) > 1{
    var aux []mail.Address
    aux = append(aux, mail.Address{v, ""})
    o.to = aux
  } else {
    if len(o.to) == 1 {
      o.to[0].Name = v
    }
  }
}

// set unique recipient mail
func (o *GoMail) SetToIdsAddress(v string)(){
  if o.to == nil || len(o.to) > 1{
    var aux []mail.Address
    aux = append(aux, mail.Address{"", v})
    o.to = aux
  } else {
    if len(o.to) == 1 {
      o.to[0].Address = v
    }
  }
}

// set ids of recipients
func (o *GoMail) SetToIds(name, email string)(){
  var aux []mail.Address
  aux = append(aux, mail.Address{name, strings.ToLower(email)})
  o.to = aux
}

// add new recipient
func (o *GoMail) AddToIds(name, email string)(){
  toId := mail.Address{name, strings.ToLower(email)}
  o.to = append(o.to, toId)
}

// set list of ids of recipients
func (o *GoMail) SetListToIds (toIdString []Recipients)(){
  var toIds []mail.Address
  for _, v:= range toIdString{
    toIds = append(toIds, mail.Address{v.Name, strings.ToLower(v.Email)})
  }
  o.to = toIds
}

// add list of ids of recipients
func (o *GoMail) AddListToIds (toIdString []Recipients)(){
  for _, v:= range toIdString{
    o.to = append(o.to, mail.Address{v.Name, strings.ToLower(v.Email)})
  }
}

// get ids of recipients
func (o *GoMail) GetToIds ()([]mail.Address){
  return o.to
}

// get ids of recipients
func (o *GoMail) GetToIdsToString ()([]string){
  var mails []string
  for _, v := range o.to {
    mails = append(mails, v.String())
  }
  return mails
}

// set from address
func (o *GoMail) SetFromAddress(v string)(){
  o.from.Address = v
}

// get from address
func (o *GoMail) GetFromAddress()(string){
  return o.from.Address
}

// set from name
func (o *GoMail) SetFromName(v string)(){
  o.from.Name = v
}

// get from name
func (o *GoMail) GetFromName()(string){
  return o.from.Name
}

// set from
func (o *GoMail) SetFrom(name, email string)(){
  v := mail.Address{name, email}
  o.from = v
}

// get from
func (o *GoMail) GetFrom()(string){
  return o.from.String()
}

// set subject message
func (o *GoMail) SetSubject(v string)(){
  o.subject = v
}

// get subject message
func (o *GoMail) GetSubject()(string){
  return o.subject
}

// set body message
func (o *GoMail) SetBodyMessage(v string)(){
  o.bodyMessage = v
}

// set boddy Message as Template
func (o *GoMail) SetBodyMessageAsTemplate(templateToResPath string, infoMail interface{}){
  t, err := template.ParseFiles(templateToResPath)
  if err != nil {
    fmt.Println(GoMailErrNotFoundTemplate)
    panic(GoMailErrNotFoundTemplate)
    return
  }
  
  // convertimos el template a bytes
  buffer := new(bytes.Buffer)
  err = t.Execute(buffer, infoMail)
  if err != nil {
    fmt.Println(GoMailErrConvertToByte)
    panic(GoMailErrConvertToByte)
    return
  }
  o.bodyMessage = buffer.String()
}

// get body message
func (o *GoMail) GetBodyMessage()(string){
  return o.bodyMessage
}

// set content type
func (o *GoMail) SetContentType(v string)(){
  o.contentType = v
}

// get content type
func (o *GoMail) GetContentType()(string){
  return o.contentType
}

// set username
func (o *GoMail) SetUsername(v string)(){
  o.username = v
}

// get username
func (o *GoMail) GetUsername()(string){
  return o.username
}

// set passwd
func (o *GoMail) SetPassword(v string)(){
  o.password = v
}

// get passwd
func (o *GoMail) GetPassword()(string){
  return o.password
}

// set servername
func (o *GoMail) SetServerName(v string)(){
  o.servername.NewSMTPServerName(v)
}

// get servername
func (o *GoMail) GetServerName()(string){
  return o.servername.GetServerName()
}

// remove element if is duplicate (email or name).
func (o *GoMail) VerifyDuplicityByEmail()(){
  result := []mail.Address{}
  for i := 0; i < len(o.to); i++ {
    // Scan slice for a previous element of the same value.
    exists := false
    for v := 0; v < i; v++ {
      if o.to[v].Address == o.to[i].Address {
        exists = true
        break
      }
    }
    // If no previous element exists, append this one.
    if !exists {
      result = append(result, o.to[i])
    }
  }
  o.to = result
}

// remove element if is duplicate (email or name).
func (o *GoMail) VerifyDuplicityByName()(){
  result := []mail.Address{}
  for i := 0; i < len(o.to); i++ {
    // Scan slice for a previous element of the same value.
    exists := false
    for v := 0; v < i; v++ {
      if o.to[v].Name == o.to[i].Name {
        exists = true
        break
      }
    }
    // If no previous element exists, append this one.
    if !exists {
      result = append(result, o.to[i])
    }
  }
  o.to = result
}

// verify all mails
func (o *GoMail) VerifyEmails()(error){
  isFromCorrect := VerifyFormatEmail(o.GetFromAddress())
  if !isFromCorrect {
    return GoMailErrBadFormatEmailFrom
  }
  for _, v := range o.GetToIds(){
    if !VerifyFormatEmail(v.Address) {
      return GoMailErrBadFormatEmailToId
    }
  }
  return nil
}

// build the message
func (o *GoMail) PrepareMessage() string {
  o.VerifyDuplicityByEmail()
   // Setup message
  message := ""
  message += fmt.Sprintf("Content-Type: %s\r\n", o.GetContentType())
  message += fmt.Sprintf("From: %s\r\n", o.GetFrom())
  if len(o.to) > 0 {
    message += fmt.Sprintf("To: %s\r\n", strings.Join(o.GetToIdsToString(), ";"))
  }

  message += fmt.Sprintf("Subject: %s\r\n", o.GetSubject())
  message += "\r\n" + o.GetBodyMessage()

  fmt.Println(message)
  return message
}

// Send Message
func (o *GoMail) SendMessage()(error){
  err := o.VerifyEmails()
	u := NewUserCredentials(o.GetUsername(), o.GetPassword())
  if err != nil{
    return err
  }
  messageToSend := o.PrepareMessage()
  servername := o.servername.GetServerName()
  host := o.servername.GetServerHost()

  //build an auth
  auth := NewGoMailAuth(u, host)

  // TLS config
  tlsconfig := o.servername.SMTPServerTLS()
  conn, err := tls.Dial("tcp", servername, tlsconfig)
  if err != nil {
    return GoMailErrBadHostPortServer
  }

  // create new client
  c, err := smtp.NewClient(conn, host)
  if err != nil {
    return GoMailErrSMTPClient
  }

  defer c.Quit()

  // step 1: Use Auth
  if err = c.Auth(auth); err != nil {
    return GoMailErrBadCredentials
  }

  // step 2: add all from and to
  if err = c.Mail(o.GetFromAddress()); err != nil {
    return GoMailErrSyntaxSender
  }

  for _, k := range o.GetToIds() {
    if err = c.Rcpt(o.GetToIdAddress(k)); err != nil {
      return GoMailErrSyntaxRecipients
    }
  }

  // Data
  w, err := c.Data()
  if err != nil {
    return GoMailErrMessageData
  }

  _, err = w.Write([]byte(messageToSend))
  if err != nil {
    return GoMailErrSendMessage
  }

  err = w.Close()
  if err != nil {
    return GoMailErrWriteClientClose
  }
  return nil
}
