package diary

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

func addExt(filename string, ext string) string {
	return fmt.Sprintf("%s.%s", filename, ext)
}

func diaryFilename(saveGameID string) string {
	return addExt(saveGameID, "diary")
}

func DiaryIDs(storageDirectory string) ([]string, error) {
	pattern := filepath.Join(storageDirectory, "*.diary")

	files, err := filepath.Glob(pattern)
	if err == nil {
		for idx, file := range files {
			file = filepath.Base(file)
			file = strings.TrimSuffix(file, ".diary")

			files[idx] = file
		}

		sort.Strings(files)
	}

	return files, err
}
