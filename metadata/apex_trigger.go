package metadata

import (
	"path/filepath"
	"strings"
)

type ApexTrigger struct {
	APIVersion      string `xml:"apiVersion"`
	FullName        string
	PackageVersions []struct {
		Namespace   string `xml:"namespace"`
		MajorNumber int    `xml:"majorNumber"`
		MinorNumber int    `xml:"minorNumber"`
	} `xml:"packageVersions"`
	Status string `xml:"status"`
}

// NewApexTriggerFromFile reads Apex trigger metadata from a file.
func NewApexTriggerFromFile(src string) (ApexTrigger, error) {
	var trigger ApexTrigger

	if err := decodeXMLFile(src, &trigger); err != nil {
		return trigger, err
	}

	trigger.FullName = strings.TrimSuffix(filepath.Base(src), ".trigger-meta.xml")

	return trigger, nil
}
