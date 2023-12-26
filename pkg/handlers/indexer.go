package handlers

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"github.com/disto-stack/j-mails-indexer/pkg/services"
)

type IndexerHandler struct {
	configService	*services.Config
}


func (ih *IndexerHandler) IndexFromTgz(filename string)  {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error open the file:", err)
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println("Error in gzip:", err)
			return
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)
	counter := 0
	for {
		counter = counter + 1
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		
		if err != nil {
			fmt.Println("Error al leer el archivo tar:", err)
			return
		}

		fmt.Println("Nombre del archivo:", header.Name)
	}

	fmt.Println(ih.configService.ZincsearchUrl)
	fmt.Println("total:", counter)
}

func (ih *IndexerHandler) SetDependencies(config *services.Config)  {
	ih.configService = config;
}