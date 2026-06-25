package storage

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

func GenerateKey(folder string, filename string) string {
	ext := filepath.Ext(filename)

	name := strings.TrimSuffix(filename, ext)

	name = strings.ReplaceAll(name, " ", "-")
	name = strings.ToLower(name)

	return fmt.Sprintf(
		"%s/%d-%s%s",
		folder,
		time.Now().UnixNano(),
		name,
		ext,
	)
}
