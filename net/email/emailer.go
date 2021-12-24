package email

/**
邮件类接口
**/
type Emailer interface {
	Send(to,subject,body string)error
	SendTLS(to,subject,body string)error
	WithAppendByPath(path string)Emailer
	WithAppendBytes(name string,data []byte)Emailer
}
