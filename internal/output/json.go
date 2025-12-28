package output

import (
	"encoding/json"
	"os"

	"github.com/Narvdeshwar/go-logstreamer/pkg/model"
)

func WriteJSON(path string, summary model.Summary) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(summary)
}
