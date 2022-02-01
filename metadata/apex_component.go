package metadata

import (
	"path/filepath"
	"strings"
)

// see https://developer.salesforce.com/docs/atlas.en-us.api_meta.meta/api_meta/meta_component.htm
type ApexComponent struct {
	FullName        string
	APIVersion      float64          `xml:"apiVersion"`
	Description     string           `xml:"description"`
	Label           string           `xml:"label"`
	PackageVersions []PackageVersion `xml:"packageVersions"`
}

// NewApexComponentFromFile return an ApexComponent from file data.
func NewApexComponentFromFile(src string) (ApexComponent, error) {
	var ac ApexComponent
	if err := decodeXMLFile(src, &ac); err != nil {
		return ac, err
	}

	ac.FullName = strings.TrimSuffix(filepath.Base(src), ".component-meta.xml")

	return ac, nil
}

// IsApexComponent returns true if the file represent an ApexComponent.
func IsApexComponentFile(src string) bool {
	return strings.HasSuffix(src, ".component-meta.xml")
}
