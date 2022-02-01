package metadata

import (
	"path/filepath"
	"strings"
)

// see https://developer.salesforce.com/docs/atlas.en-us.api_meta.meta/api_meta/meta_pages.htm
type ApexPage struct {
	FullName                  string
	APIVersion                float64          `xml:"apiVersion"`
	Description               string           `xml:"description"`
	AvailableInTouch          bool             `xml:"availableInTouch"`
	ConfirmationTokenRequired bool             `xml:"confirmationTokenRequired"`
	Label                     string           `xml:"label"`
	PackageVersions           []PackageVersion `xml:"packageVersions"`
}

// NewApexPageFromFile returns an ApexPage from data file.
func NewApexPageFromFile(src string) (ApexPage, error) {
	var ap ApexPage
	if err := decodeXMLFile(src, &ap); err != nil {
		return ap, err
	}

	ap.FullName = strings.TrimSuffix(filepath.Base(src), ".page-meta.xml")

	return ap, nil
}

func IsApexPageFile(src string) bool {
	return strings.HasSuffix(src, ".page-meta.xml")
}
