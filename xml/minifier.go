package xml

import (
	"fmt"
	"os"

	"github.com/xrstf/ppxml/xml"
)

func MinifyFile(filename string) error {
	// open file
	input, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer input.Close()

	// determine output
	dstFilename := fmt.Sprintf("%s.mini", filename)
	dstFile, err := os.Create(dstFilename)
	if err != nil {
		return err
	}

	copyBOM(input, dstFile)

	err = xml.Minify(input, dstFile)
	if err != nil {
		dstFile.Close()
		os.Remove(dstFilename)
		return err
	}

	return dstFile.Close()
}
