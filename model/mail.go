package model

type Mail struct {
	Name  string
	Email string
}

type Request struct {
	From    Mail
	ToList  []Mail
	CcList  []Mail
	BccList []Mail
	Subject string
	Content string
}
