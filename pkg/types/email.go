package types

type Email struct {
	MessageId		string	`json:"message_id"`
	Date				string	`json:"date"`
	From				string	`json:"from"`
	To					string	`json:"to"`
	Subject			string	`json:"subject"`
	Content			string	`json:"content"`
}

type EmailData struct {
	Index		string	`json:"index"`
	Records	[]Email	`json:"records"`
}