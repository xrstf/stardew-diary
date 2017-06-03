package xml

import "io"

func copyBOM(input io.Reader, output io.Writer) (err error) {
	bom := make([]byte, 3)

	_, err = input.Read(bom)
	if err != nil {
		return
	}

	_, err = output.Write(bom)
	return
}
