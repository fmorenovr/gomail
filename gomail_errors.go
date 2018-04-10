package gomail

import(
  "errors";
)

// list of errors
var(
  GomailErrBadFormatEmailFrom = errors.New("Bad format of email sender.\n")
  GomailErrBadFormatEmailToId = errors.New("Bad format of email reciever.\n")
  GomailErrMisbehavingServer  = errors.New("Server misbehaving.\n")
  GomailErrWrongHostName      = errors.New("Wrong Host name.\n")
  GomailErrSendMessage        = errors.New("Error sending message.\n")
  GomailErrWriteClientClose   = errors.New("Closing Writer Client.\n")
  GomailErrSMTPClient         = errors.New("Error Creating SMTP Client.\n")
  GomailErrBadCredentials     = errors.New("Username and Password not accepted.\nCould be an error in credentials\n.")
  GomailErrBadHostPortServer  = errors.New("Host or Port are incorrect and cannot connected with server. Or Server misbehaving.")
  GomailErrSyntaxSender       = errors.New("Syntax error in sender\n.")
  GomailErrSyntaxRecipients   = errors.New("Syntax error in recipients\n.")
  GomailErrMessageData        = errors.New("Error in data to send as message\n.")
)
