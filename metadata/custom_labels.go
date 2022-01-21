package metadata

import (
	"path/filepath"
	"strings"
)

// see https://developer.salesforce.com/docs/atlas.en-us.api_meta.meta/api_meta/meta_customlabels.htm
type CustomLabels struct {
	FullName string
	Labels   []CustomLabel `xml:"labels"`
}

// see https://developer.salesforce.com/docs/atlas.en-us.api_meta.meta/api_meta/meta_customlabels.htm
type CustomLabel struct {
	FullName         string `xml:"fullName"`
	Value            string `xml:"value"`
	Language         string `xml:"language"`
	Protected        bool   `xml:"protected"`
	ShortDescription string `xml:"shortDescription"`
	Categories       string `xml:"categories"`
}

// NewCustomLabelsFromFile reads CustomLablels metadata from a file.
func NewCustomLabelsFromFile(src string) (CustomLabels, error) {
	var labels CustomLabels

	if err := decodeXMLFile(src, &labels); err != nil {
		return labels, err
	}

	labels.FullName = strings.TrimSuffix(filepath.Base(src), "labels.-meta.xml")

	return labels, nil
}
