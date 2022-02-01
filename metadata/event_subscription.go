package metadata

import (
	"path/filepath"
	"strings"
)

// see https://developer.salesforce.com/docs/atlas.en-us.api_meta.meta/api_meta/meta_eventsubscription.htm
type EventSubscription struct {
	FullName        string
	Active          bool `xml:"active"`
	EventParameters []struct {
		Name  string `xml:"parameterName"`
		Value string `xml:"parameterValue"`
	} `xml:"eventParameters"`
	EventType     string `xml:"eventType"`
	ReferenceData string `xml:"referenceData"`
}

// NewEventSubscriptionFromFile reads a metadata file and returns an EventSubscription.
func NewEventSubscriptionFromFile(src string) (EventSubscription, error) {
	var es EventSubscription
	if err := decodeXMLFile(src, &es); err != nil {
		return es, err
	}

	es.FullName = strings.TrimSuffix(filepath.Base(src), ".subscription-meta.xml")

	return es, nil
}

func IsEventSubscriptionFile(src string) bool {
	return strings.HasSuffix(src, ".subscription-meta.xml")
}
