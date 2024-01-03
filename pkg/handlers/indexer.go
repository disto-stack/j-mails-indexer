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
	"sync"

	"github.com/disto-stack/j-mails-indexer/pkg/services"
	"github.com/disto-stack/j-mails-indexer/pkg/types"
)

type IndexerHandler struct {
	configService     *services.ConfigService
	zincsearchService *services.ZincsearchService
}

var (
	wg sync.WaitGroup
	m  sync.Mutex
)

func (i *IndexerHandler) SetDependencies(c *services.ConfigService, z *services.ZincsearchService) {
	i.configService = c
	i.zincsearchService = z
}

func (ih *IndexerHandler) IndexFromDir(dir string) {
	emailDataSlice := []types.EmailData{
		{Index: "email", Records: []types.Email{}},
	}

	err := filepath.Walk(dir+"/maildir/", func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			wg.Add(1)

			go func(path string) error {
				email, err := readFile(path)
				if err != nil {
					return nil
				}

				m.Lock()

				lastEmailData := &emailDataSlice[len(emailDataSlice)-1]
				if len(lastEmailData.Records) >= 50000 {
					emailDataSlice = append(emailDataSlice, types.EmailData{
						Index: "email", Records: []types.Email{},
					})

					lastEmailData = &emailDataSlice[len(emailDataSlice)-1]
				}

				lastEmailData.Records = append(lastEmailData.Records, email)

				m.Unlock()

				return nil
			}(path)

		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()

	err = ih.zincsearchService.UploadIndex()
	ih.zincsearchService.UploadBulkData(emailDataSlice)
	if err != nil {
		log.Fatal(err)
	}
}

func readFile(path string) (types.Email, error) {
	defer wg.Done()
	file, err := os.ReadFile(path)

	b := bytes.NewReader(file)
	message, err := mail.ReadMessage(b)

	if err != nil {
		return types.Email{}, err
	}

	messageId := message.Header.Get("MESSAGE-ID")
	from := message.Header.Get("FROM")
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
		From:      from,
		To:        to,
		Subject:   subject,
		Content:   content,
	}

	return email, nil
}
