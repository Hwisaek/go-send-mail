package sendmail

type Mail struct {
	Name  string
	Email string
}

type Attachment struct {
	Filename string
	Content  []byte
	Type     string
}
