package metadata

import (
	"path/filepath"
	"strings"
)

type ApexTrigger struct {
	FullName        string
	APIVersion      float64 `xml:"apiVersion"`
	PackageVersions []PackageVersion
	Status          string `xml:"status"`
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

// IsApexTriggerFile returns true if the file is an ApexTrigger metadata file
func IsApexTriggerFile(src string) bool {
	return strings.HasSuffix(src, ".trigger-meta.xml")
}
