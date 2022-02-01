package metadata

import (
	"path/filepath"
	"strings"
)

// see https://developer.salesforce.com/docs/atlas.en-us.api_meta.meta/api_meta/meta_classes.htm
type ApexClass struct {
	FullName        string
	APIVersion      float64          `xml:"apiVersion"`
	PackageVersions []PackageVersion `xml:"packageVersions"`
	Status          string           `xml:"status"`
}

type PackageVersion struct {
	Namespace   string `xml:"namespace"`
	MajorNumber int    `xml:"majorNumber"`
	MinorNumber int    `xml:"minorNumber"`
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

// IsApexClassFile returns true if the file is an ApexClass metadata file.
func IsApexClassFile(src string) bool {
	return strings.HasSuffix(src, ".cls-meta.xml")
}
