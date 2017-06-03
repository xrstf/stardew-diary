package diary

import "fmt"

func addExt(filename string, ext string) string {
	return fmt.Sprintf("%s.%s", filename, ext)
}
