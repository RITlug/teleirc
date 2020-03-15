package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)


type DocumentObject struct {
	EnableLink bool
	Link string
	Object *tgbotapi.Document
}

/*
Constructs a new document and is used to decide under which format to formard
the document to the IRC client
*/

func newDocument(enabled bool, link string, u tgbotapi.Update) DocumentObject {
	doc := u.Message.Document
	return DocumentObject{enabled, link, doc}
}

/*
documentHandler once a document received, uses information about the document
to format a message to return to the client.
*/
func (doc DocumentObject) documentHandler()  string {

	message := " shared a file (" +
		doc.Object.MimeType + ") on Telegram with title" +
		"'" + doc.Object.FileName + "'."

	if doc.EnableLink {
		message += " View document at " + doc.Link
	}

	return message
}
