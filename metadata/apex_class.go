package metadata

import (
	"path/filepath"
	"strings"
)

// see https://developer.salesforce.com/docs/atlas.en-us.api_meta.meta/api_meta/meta_classes.htm
type ApexClass struct {
	APIVersion      string `xml:"apiVersion"`
	FullName        string
	PackageVersions []struct {
		Namespace   string `xml:"namespace"`
		MajorNumber int    `xml:"majorNumber"`
		MinorNumber int    `xml:"minorNumber"`
	} `xml:"packageVersions"`
	Status string `xml:"status"`
}

// NewApexClassFromFile read Apex class metadata from a file.
func NewApexClassFromFile(src string) (ApexClass, error) {
	var class ApexClass

	if err := decodeXMLFile(src, &class); err != nil {
		return class, err
	}

	class.FullName = strings.TrimSuffix(filepath.Base(src), ".cls-meta.xml")

	return class, nil
}
