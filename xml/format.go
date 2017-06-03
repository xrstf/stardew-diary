package xml

import (
	"fmt"
	"os"

	"github.com/xrstf/ppxml/xml"
)

func FormatFile(filename string) error {
	// open file
	input, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer input.Close()

	// determine output
	dstFilename := fmt.Sprintf("%s.xml", filename)
	dstFile, err := os.Create(dstFilename)
	if err != nil {
		return err
	}

	copyBOM(input, dstFile)

	err = xml.Format(input, dstFile)
	if err != nil {
		dstFile.Close()
		os.Remove(dstFilename)
		return err
	}

	return dstFile.Close()
}
