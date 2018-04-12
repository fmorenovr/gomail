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

type Gomail struct {
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
func NewGomail()(*Gomail, error){
  return &Gomail{
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
func (o *Gomail) Set(key, value string)(){
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
func (o *Gomail) FindByEmail(email string)(aux int){
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
func (o *Gomail) FindByName(name string)(aux int){
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
func (o *Gomail) ChangeToId(by, original, v string)(){
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
func (o *Gomail) changeToIdName(original, v string) (){
  aux := o.FindByName(original)
  if aux == -1 {
    return
  } else {
    o.to[aux].Name = v
  }
}

// set specific recipient address
func (o *Gomail) changeToIdEmail(original, v string)(){
  aux := o.FindByEmail(strings.ToLower(original))
  if aux == -1 {
    return
  } else {
    o.to[aux].Address = strings.ToLower(v)
  }
}

// menu to delete
func (o *Gomail) DeleteToId(by, original string)(){
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
func (o *Gomail) deleteToIdByName(original string)(){
  i := o.FindByName(original)
  if i == -1 {
    return
  } else {
    o.to = o.to[:i+copy(o.to[i:], o.to[i+1:])]
  }
}

// delete specific recipient address
func (o *Gomail) deleteToIdByEmail(original string)(){
  i := o.FindByEmail(strings.ToLower(original))
  if i == -1 {
    return
  } else {
    o.to = o.to[:i+copy(o.to[i:], o.to[i+1:])]
  }
}

// get specific recipient address
func (o *Gomail) GetToIdAddress(toId mail.Address)(string){
  return toId.Address
}

// get specific recipient name
func (o *Gomail) GetToIdName(toId mail.Address)(string){
  return toId.Name
}

// set unique recipient name
func (o *Gomail) SetToIdsName(v string)(){
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
func (o *Gomail) SetToIdsAddress(v string)(){
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
func (o *Gomail) SetToIds(name, email string)(){
  var aux []mail.Address
  aux = append(aux, mail.Address{name, strings.ToLower(email)})
  o.to = aux
}

// add new recipient
func (o *Gomail) AddToIds(name, email string)(){
  toId := mail.Address{name, strings.ToLower(email)}
  o.to = append(o.to, toId)
}

// set list of ids of recipients
func (o *Gomail) SetListToIds (toIdString []Recipients)(){
  var toIds []mail.Address
  for _, v:= range toIdString{
    toIds = append(toIds, mail.Address{v.Name, strings.ToLower(v.Email)})
  }
  o.to = toIds
}

// add list of ids of recipients
func (o *Gomail) AddListToIds (toIdString []Recipients)(){
  for _, v:= range toIdString{
    o.to = append(o.to, mail.Address{v.Name, strings.ToLower(v.Email)})
  }
}

// get ids of recipients
func (o *Gomail) GetToIds ()([]mail.Address){
  return o.to
}

// get ids of recipients
func (o *Gomail) GetToIdsToString ()([]string){
  var mails []string
  for _, v := range o.to {
    mails = append(mails, v.String())
  }
  return mails
}

// set from address
func (o *Gomail) SetFromAddress(v string)(){
  o.from.Address = v
}

// get from address
func (o *Gomail) GetFromAddress()(string){
  return o.from.Address
}

// set from name
func (o *Gomail) SetFromName(v string)(){
  o.from.Name = v
}

// get from name
func (o *Gomail) GetFromName()(string){
  return o.from.Name
}

// set from
func (o *Gomail) SetFrom(name, email string)(){
  v := mail.Address{name, email}
  o.from = v
}

// get from
func (o *Gomail) GetFrom()(string){
  return o.from.String()
}

// set subject message
func (o *Gomail) SetSubject(v string)(){
  o.subject = v
}

// get subject message
func (o *Gomail) GetSubject()(string){
  return o.subject
}

// set body message
func (o *Gomail) SetBodyMessage(v string)(){
  o.bodyMessage = v
}

// set boddy Message as Template
func (o *Gomail) SetBodyMessageAsTemplate(templateToResPath string, infoMail interface{}){
  t, err := template.ParseFiles(templateToResPath)
  if err != nil {
    fmt.Println(GomailErrNotFoundTemplate)
    panic(GomailErrNotFoundTemplate)
    return
  }
  
  // convertimos el template a bytes
  buffer := new(bytes.Buffer)
  err = t.Execute(buffer, infoMail)
  if err != nil {
    fmt.Println(GomailErrConvertToByte)
    panic(GomailErrConvertToByte)
    return
  }
  o.bodyMessage = buffer.String()
}

// get body message
func (o *Gomail) GetBodyMessage()(string){
  return o.bodyMessage
}

// set content type
func (o *Gomail) SetContentType(v string)(){
  o.contentType = v
}

// get content type
func (o *Gomail) GetContentType()(string){
  return o.contentType
}

// set username
func (o *Gomail) SetUsername(v string)(){
  o.username = v
}

// get username
func (o *Gomail) GetUsername()(string){
  return o.username
}

// set passwd
func (o *Gomail) SetPassword(v string)(){
  o.password = v
}

// get passwd
func (o *Gomail) GetPassword()(string){
  return o.password
}

// set servername
func (o *Gomail) SetServerName(v string)(){
  o.servername.NewSMTPServerName(v)
}

// get servername
func (o *Gomail) GetServerName()(string){
  return o.servername.GetServerName()
}

// remove element if is duplicate (email or name).
func (o *Gomail) VerifyDuplicityByEmail()(){
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
func (o *Gomail) VerifyDuplicityByName()(){
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
func (o *Gomail) VerifyEmails()(error){
  isFromCorrect := VerifyFormatEmail(o.GetFromAddress())
  if !isFromCorrect {
    return GomailErrBadFormatEmailFrom
  }
  for _, v := range o.GetToIds(){
    if !VerifyFormatEmail(v.Address) {
      return GomailErrBadFormatEmailToId
    }
  }
  return nil
}

// build the message
func (o *Gomail) PrepareMessage() string {
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
func (o *Gomail) SendMessage()(error){
  err := o.VerifyEmails()
	u := NewUserCredentials(o.GetUsername(), o.GetPassword())
  if err != nil{
    return err
  }
  messageToSend := o.PrepareMessage()
  servername := o.servername.GetServerName()
  host := o.servername.GetServerHost()

  //build an auth
  auth := NewGomailAuth(u, host)

  // TLS config
  tlsconfig := o.servername.SMTPServerTLS()
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
