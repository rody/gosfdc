package metadata

import (
	"encoding/xml"
	"fmt"
	"os"
)

func decodeXMLFile(src string, data interface{}) error {
	file, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("could not open metadata file '%s' : %w", src, err)
	}
	defer file.Close()

	if err = xml.NewDecoder(file).Decode(data); err != nil {
		return fmt.Errorf("failed to decode metadata: %w", err)
	}

	return nil
}
