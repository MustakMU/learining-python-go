package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	imap "github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
)

func main() {
	fetchIMAPAttachments()
}
func fetchIMAPAttachments() error {

	// connect to server
	c, err := client.DialTLS("imap.gmail.com:993", nil)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] imap: connected")
	defer func() {
		log.Printf("[DEBUG] imap: logout")
		if err := c.Logout(); err != nil {
			log.Printf("[ERROR] imap: logout error %v", err)
		}
	}()

	// log.Printf("[DEBUG] imap: enable debug")
	// c.SetDebug(os.Stdout)

	// login
	err = c.Login("mustaku055@gmail.com", "945$wyjgm")
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] imap: logged in")
	// criteria := imap.NewSearchCriteria()
	// criteria.
	// select mailbox
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		return err
	}

	// get all messages
	if mbox.Messages == 0 {
		return fmt.Errorf("no messages found in mailbox")
	}

	from := uint32(2)
	to := mbox.Messages
	if mbox.Messages > 3 {
		// We're using unsigned integers here, only subtract if the result is > 0
		from = mbox.Messages - 3
	}
	log.Printf("[INFO] imap: found %v messages, %v unseen", mbox.Messages, mbox.Unseen)

	// set for all messages
	seqSet := new(imap.SeqSet)
	seqSet.AddRange(from, to)

	// set for delete messages
	deleteSet := new(imap.SeqSet)

	// get the whole message body
	section := &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem(), imap.FetchRFC822, imap.FetchEnvelope, imap.FetchBody, imap.FetchFlags, imap.FetchInternalDate}

	messages := make(chan *imap.Message, 1)
	done := make(chan error, 1)

	go func() {
		done <- c.Fetch(seqSet, items, messages)
	}()

	downloadCount := 0
	for msg := range messages {
		if msg == nil {
			return fmt.Errorf("server didn't return message")
		}
		br := msg.GetBody(section)
		if br == nil {
			return fmt.Errorf("server didn't return message body")
		}
		fmt.Println(msg.Envelope.From)
		fmt.Println("SUBJECT ", msg.Envelope.Subject)
		fmt.Println("FROM ", msg.Envelope.From)
		fmt.Println("TO ", msg.Envelope.Date)

		body := make([]byte, 10000)
		br.Read(body)
		fmt.Println("BODY")
		fmt.Println(string(body))
		// create a new mail reader
		mr, err := mail.CreateReader(br)
		if err != nil {
			return err
		}

		// process each message's part
		isSuccess := false
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Printf("[ERROR] imap: can't read next part: %v, skip", err)
				isSuccess = false
				break
			}
			fmt.Println("SUB ", p.Header.Get("Subject"))
			switch h := p.Header.(type) {

			case *mail.Header:
				b, _ := ioutil.ReadAll(p.Body)
				log.Println("BB ", string(b))
			case *mail.AttachmentHeader:
				// this is an attachment
				filename, err := h.Filename()
				if err != nil {
					log.Printf("[ERROR] imap: %v, skip", err)
					continue
				}
				log.Printf("[INFO] imap: found attachment: %v", filename)

				outFile := filepath.Join(".", filename)
				log.Printf("[INFO] imap: save attachment to: %v", outFile)
				f, err := os.Create(outFile)
				if err != nil {
					log.Printf("[ERROR] imap: %v, skip", err)
					continue
				}

				_, err = io.Copy(f, p.Body)
				if err != nil {
					log.Printf("[ERROR] imap: %v, skip", err)
					continue
				}
				f.Close()
			}
		}

		if isSuccess {
			log.Printf("[DEBUG] imap: add SeqNum %v to delete set", msg.SeqNum)
			deleteSet.AddNum(msg.SeqNum)
		}
		downloadCount++
		fmt.Println()
		fmt.Println()

		fmt.Println()

	}
	log.Printf("[DEBUG] imap: %v attachments downloaded", downloadCount)

	if err := <-done; err != nil {
		return err
	}

	// if cfg.Input.IMAP.Delete {
	// 	log.Printf("[DEBUG] imap: delete emails after fetch")

	// 	// first, mark the messages as deleted
	// 	delItems := imap.FormatFlagsOp(imap.AddFlags, false)
	// 	delFlags := []interface{}{imap.DeletedFlag}

	// 	err := c.Store(deleteSet, delItems, delFlags, nil)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// then delete it
	// 	if err := c.Expunge(nil); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}
