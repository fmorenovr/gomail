package gomail

import(
  "errors";
)

// list of errors
var(
  GoMailErrBadFormatEmailFrom = errors.New("Bad format of email sender.\n")
  GoMailErrBadFormatEmailToId = errors.New("Bad format of email reciever.\n")
  GoMailErrMisbehavingServer  = errors.New("Server misbehaving.\n")
  GoMailErrWrongHostName      = errors.New("Wrong Host name.\n")
  GoMailErrSendMessage        = errors.New("Error sending message.\n")
  GoMailErrWriteClientClose   = errors.New("Closing Writer Client.\n")
  GoMailErrSMTPClient         = errors.New("Error Creating SMTP Client.\n")
  GoMailErrBadCredentials     = errors.New("Username and Password not accepted.\nCould be an error in credentials.\n")
  GoMailErrBadHostPortServer  = errors.New("Host or Port are incorrect and cannot connected with server. Or Server misbehaving.\n")
  GoMailErrSyntaxSender       = errors.New("Syntax error in sender.\n")
  GoMailErrSyntaxRecipients   = errors.New("Syntax error in recipients.\n.")
  GoMailErrMessageData        = errors.New("Error in data to send as message.\n")
  GoMailErrNotFoundTemplate   = errors.New("Not Found Template's Path to Response.\n")
  GoMailErrConvertToByte      = errors.New("Cannot Convert Template to bytes.\n")
)
