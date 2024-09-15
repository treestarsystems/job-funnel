package extract

import (
	"fmt"
	"io"
	"job-funnel/types"
	"mime"
	"mime/multipart"
	"net/mail"
	"strings"
)

// ParseEmailMessage parses an email message and returns an EmailMessage struct.
func ParseEmailMessage(rawMessage string) (*types.EmailMessage, error) {
	// Read the email message
	msg, err := mail.ReadMessage(strings.NewReader(rawMessage))
	if err != nil {
		return nil, fmt.Errorf("error - Email parsing message: %v", err)
	}

	// Extract the subject
	subject := msg.Header.Get("Subject")

	// Extract the date
	date := msg.Header.Get("Date")

	// Extract the sender
	from, err := mail.ParseAddress(msg.Header.Get("From"))
	if err != nil {
		return nil, fmt.Errorf("error - Email parsing sender: %v - %s:%s", err, subject, date)
	}

	// Extract the recipients
	to, err := msg.Header.AddressList("To")
	if err != nil {
		return nil, fmt.Errorf("error - Email parsing recipients: %v - %s:%s", err, subject, date)
	}

	// Parse the body and attachments
	emailPlainText, emailHtml, attachments, err := parseBodyAndAttachments(msg)
	if err != nil {
		return nil, fmt.Errorf("error - Email parsing body and attachments: %v - %s:%s", err, subject, date)
	}

	return &types.EmailMessage{
		EmailSubject:     subject,
		EmailFrom:        from,
		EmailTo:          to,
		EmailDate:        date,
		EmailHTML:        emailHtml,
		EmailPlainText:   emailPlainText,
		EmailAttachments: attachments,
	}, nil
}

// parseBodyAndAttachments parses the body and attachments of an email message.
func parseBodyAndAttachments(msg *mail.Message) (string, string, []types.EmailAttachment, error) {
	var plainText, html string
	var attachments []types.EmailAttachment

	mediaType, params, err := mime.ParseMediaType(msg.Header.Get("Content-Type"))
	if err != nil {
		return "", "", nil, fmt.Errorf("error - Email parsing media type: %v", err)
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		mr := multipart.NewReader(msg.Body, params["boundary"])
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				return "", "", nil, fmt.Errorf("Email reading multipart: %v", err)
			}

			slurp, err := io.ReadAll(p)
			if err != nil {
				return "", "", nil, fmt.Errorf("Email reading part: %v", err)
			}

			contentDisposition := p.Header.Get("Content-Disposition")
			contentType := p.Header.Get("Content-Type")
			if strings.HasPrefix(contentDisposition, "attachment") {
				filename := p.FileName()
				attachments = append(attachments, types.EmailAttachment{
					EmailAttachmentFilename: filename,
					EmailAttachmentContent:  slurp,
				})
			} else if strings.HasPrefix(contentType, "text/plain") {
				plainText += string(slurp)
			} else if strings.HasPrefix(contentType, "text/html") {
				html += string(slurp)
			}
		}
	} else {
		slurp, err := io.ReadAll(msg.Body)
		if err != nil {
			return "", "", nil, fmt.Errorf("Email reading body: %v", err)
		}
		if strings.HasPrefix(mediaType, "text/plain") {
			plainText = string(slurp)
		} else if strings.HasPrefix(mediaType, "text/html") {
			html = string(slurp)
		}
	}

	return plainText, html, attachments, nil
}
