package gomail

// Names and Mails from recipients
type Recipients struct{
  Name string
  Email string
}

// username & passwd
type UserCredentials struct {
  username string
  passwd   string
}

// set username & passwd
func NewUserCredentials(u, v string)(*UserCredentials){
  return &UserCredentials{
           username: u,
           passwd: v,
         }
}

// set username
func (o *UserCredentials) SetUsername(v string)(){
  o.username = v
}

// get username
func (o *UserCredentials) GetUsername()(string){
  return o.username
}

// set passwd
func (o *UserCredentials) SetPasswd(v string)(){
  o.passwd = v
}

// get passwd
func (o *UserCredentials) GetPasswd()(string){
  return o.passwd
}
