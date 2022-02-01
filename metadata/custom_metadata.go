package metadata

import (
	"path/filepath"
	"strings"
)

// see https://developer.salesforce.com/docs/atlas.en-us.api_meta.meta/api_meta/meta_custommetadata.htm
type CustomMetadata struct {
	FullName    string
	Description string `xml:"description"`
	Label       string `xml:"label"`
	Protected   bool   `xml:"protected"`
	Values      []CustomMetadataValue
}

// see https://developer.salesforce.com/docs/atlas.en-us.api_meta.meta/api_meta/meta_custommetadata.htm#CustomMetadataValues_title
type CustomMetadataValue struct {
	Field string `xml:"field"`
	Value string `xml:"value"`
	Type  string `xml:"xsd:type,attr"`
}

// NewCustomMetadataFromFile reads CustomMetadata info from a file.
func NewCustomMetadataFromFile(src string) (CustomMetadata, error) {
	var cm CustomMetadata

	if err := decodeXMLFile(src, &cm); err != nil {
		return cm, err
	}

	cm.FullName = strings.TrimSuffix(filepath.Base(src), ".md-meta.xml")

	return cm, nil
}

func IsCustomMetadataFile(src string) bool {
	return strings.HasSuffix(src, ".md-meta.xml")
}
