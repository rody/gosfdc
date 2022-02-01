package metadata

import (
	"path/filepath"
	"strings"
)

// see https://developer.salesforce.com/docs/atlas.en-us.api_meta.meta/api_meta/meta_assignmentrule.htm
type AssignmentRules struct {
	FullName string
	Rules    []AssignmentRule `xml:"assignmentRule"`
}

// see https://developer.salesforce.com/docs/atlas.en-us.api_meta.meta/api_meta/meta_assignmentrule.htm#AssignmentRule_sub
type AssignmentRule struct {
	Active      bool        `xml:"active"`
	FullName    string      `xml:"fullName"`
	RuleEntries []RuleEntry `xml:"ruleEntry"`
}

// see https://developer.salesforce.com/docs/atlas.en-us.api_meta.meta/api_meta/meta_assignmentrule.htm#RuleEntry_title
type RuleEntry struct {
	AssignedTo            string       `xml:"assignedTo"`
	AssignedToType        string       `xml:"assignedToType"`
	BooleanFilter         string       `xml:"booleanFilter"`
	CriteriaItems         []FilterItem `xml:"criteriaItems"`
	Formula               string       `xml:"formula"`
	NotifyCCRecipient     bool         `xml:"notifyCcRecipients"`
	OverrideExistingTeams bool         `xml:"overrideExistingTeams"`
	Teams                 []string     `xml:"team"`
	Template              string       `xml:"template"`
}

// see https://developer.salesforce.com/docs/atlas.en-us.api_meta.meta/api_meta/customfield.htm#filteritem
type FilterItem struct {
	Field      string `xml:"field"`
	Operation  string `xml:"operation"`
	Value      string `xml:"value"`
	ValueField string `xml:"ValueField"`
}

func NewAssignmentRulesFromFile(src string) (AssignmentRules, error) {
	var ar AssignmentRules
	if err := decodeXMLFile(src, &ar); err != nil {
		return ar, err
	}

	ar.FullName = strings.TrimSuffix(filepath.Base(src), ".assignmentRules-meta.xml")

	return ar, nil
}

func IsAssignmentRulesFile(src string) bool {
	return strings.HasSuffix(src, ".assignmentRules-meta.xml")
}
