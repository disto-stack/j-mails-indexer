package handlers

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/mail"
	"os"
	"path/filepath"

	"github.com/disto-stack/j-mails-indexer/pkg/services"
	"github.com/disto-stack/j-mails-indexer/pkg/types"
)

type IndexerHandler struct {
	configService	*services.Config
	zincsearchService	*services.ZincsearchService
}

func (i *IndexerHandler) SetDependencies(config *services.Config, zincsearchService *services.ZincsearchService)  {
	i.configService = config;
	i.zincsearchService = zincsearchService;
}

func (ih *IndexerHandler) IndexFromDir(dir string)  {
	emailDataSlice := []types.EmailData{
		{ Index: "email", Records: []types.Email{} },
	}
	fmt.Println(dir)
	err := filepath.Walk(dir + "/maildir/", func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			email, err := readFile(path)
			if err != nil {
				return nil
			}

			lastEmailData := &emailDataSlice[len(emailDataSlice) - 1];
			if len(lastEmailData.Records) >= 100000 {
				emailDataSlice = append(emailDataSlice, types.EmailData{
					Index: "email", Records: []types.Email{},
				})

				lastEmailData = &emailDataSlice[len(emailDataSlice) - 1];
			}

			lastEmailData.Records = append(lastEmailData.Records, email)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	
	err = ih.zincsearchService.UploadIndex()
	ih.zincsearchService.UploadBulkData(emailDataSlice)
	if err != nil {
		log.Fatal(err)
	}
}

func readFile(path string) (types.Email, error) {
	 file, err := os.ReadFile(path);

		b := bytes.NewReader(file)
		message, err := mail.ReadMessage(b)

		if err != nil {
			return types.Email{}, err
		}

		messageId := message.Header.Get("MESSAGE-ID")
		from := message.Header.Get("MESSAGE-ID")
		to := message.Header.Get("TO")
		subject := message.Header.Get("SUBJECT")

		body, err := io.ReadAll(message.Body)
		if err != nil {
			fmt.Println(err)
			return types.Email{}, err
		}

		content := string(body)

		email := types.Email{
			MessageId: messageId,
			From: from,
			To: to,
			Subject: subject,
			Content: content,
		}

		return email, nil
}