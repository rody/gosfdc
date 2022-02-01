package metadata

import (
	"path/filepath"
	"strings"
)

// see https://developer.salesforce.com/docs/atlas.en-us.api_meta.meta/api_meta/meta_apextestsuite.htm
type ApexTestSuite struct {
	FullName       string
	TestClassNames []string `xml:"testClassName"`
}

func NewApexTestSuiteFromFile(src string) (ApexTestSuite, error) {
	var apt ApexTestSuite

	if err := decodeXMLFile(src, &apt); err != nil {
		return apt, err
	}

	apt.FullName = strings.TrimSuffix(filepath.Base(src), ".testSuite-meta.xml")

	return apt, nil
}

func IsApexTestSuiteFile(src string) bool {
	return strings.HasSuffix(src, ".testSuite-meta.xml")
}
