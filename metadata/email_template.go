package metadata

import (
	"path/filepath"
	"strings"
)

// see https://developer.salesforce.com/docs/atlas.en-us.api_meta.meta/api_meta/meta_emailtemplate.htm
type EmailTemplate struct {
	FullName          string
	APIVersion        float64  `xml:"apiVersion"`
	AttachedDocuments []string `xml:"attachedDocuments"`
	Attachments       []struct {
		Name string `xml:"name"`
	} `xml:"attachments"`
	Available         bool             `xml:"available"`
	Description       string           `xml:"description"`
	EncodingKey       string           `xml:"encodingKey"`
	LetterHead        string           `xml:"letterhead"`
	PackageVersions   []PackageVersion `xml:"packageVersions"`
	RelatedEntityType string           `xml:"relatedEntityType"`
	Style             string           `xml:"style"`
	Subject           string           `xml:"subject"`
	TextOnly          string           `xml:"textOnly"`
	Type              string           `xml:"type"`
	UIType            string           `xml:"uiType"`
}

// NewEmailTemplateFromFile reads a metadata file and returns an EmailTemplate.
func NewEmailTemplateFromFile(src string) (EmailTemplate, error) {
	var et EmailTemplate
	if err := decodeXMLFile(src, &et); err != nil {
		return et, err
	}

	et.FullName = strings.TrimSuffix(filepath.Base(src), ".email-meta.xml")

	return et, nil
}

// IsEmailTemplateFile returns true if the file is an EmailTemplate metadata file.
func IsEmailTemplateFile(src string) bool {
	return strings.HasSuffix(src, ".email-meta.xml")
}
