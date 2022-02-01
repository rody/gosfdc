package metadata

import (
	"fmt"
	"path/filepath"
	"strings"
)

type CustomField struct {
	CaseSensitive            bool   `xml:"caseSensitive"`
	CustomDataType           string `xml:"customDataType"`
	DefaultValue             string `xml:"defaultValue"`
	DeleteConstraint         string `xml:"deleteConstraint"`
	Deprecated               bool   `xml:"deprecated"`
	Description              string `xml:"description"`
	DisplayFormat            string `xml:"displayFormat"`
	DisplayLocationInDecimal bool   `xml:"displayLocationInDecimal"`
	Encrypted                bool   `xml:"encrypted"`
	ExternalDeveloperName    string `xml:"externalDeveloperName"`
	ExternalID               bool   `xml:"externalId"`
	FieldManageability       string `xml:"fieldManageability"`
	Formula                  string `xml:"formula"`
	FormulaTreatBlankAs      string `xml:"formulaTreatBlankAs"`
	FullName                 string
	GlobalPicklist           string `xml:"globalPicklist"`
	Indexed                  bool   `xml:"indexed"`
	InlineHelpText           string `xml:"inlineHelpText"`
	IsFilteringDisabled      bool   `xml:"isFilteringDisabled"`
	IsNameField              bool   `xml:"isNameField"`
	IsSortingDisabled        bool   `xml:"isSortingDisabled"`
	ReparentableMasterDetail bool   `xml:"reparentableMasterDetail"`
	Label                    string `xml:"label"`
	Length                   int    `xml:"length"`
	LookupFilter             struct {
		Active        bool   `xml:"active"`
		BooleanFilter string `xml:"booleanFilter"`
		Description   string `xml:"description"`
		ErrorMessage  string `xml:"errorMessage"`
		FilterItems   []struct {
			Field      string `xml:"field"`
			Operation  string `xml:"operation"`
			Value      string `xml:"value"`
			ValueField string `xml:"valueField"`
		} `xml:"filterItems"`
		InfoMessage string `xml:"infoMessage"`
		IsOptional  bool   `xml:"isOptional"`
	} `xml:"lookupFilters"`
	MaskChar                             string       `xml:"maskChar"`
	MaskType                             string       `xml:"maskType"`
	MetadataRelationshipControllingField string       `xml:"metadataRelationshipControllingField"`
	Picklist                             Picklist     `xml:"picklist"`
	PopulateExistingRows                 bool         `xml:"populateExistingRows"`
	Precision                            int          `xml:"precision"`
	ReferenceTargetField                 string       `xml:"referenceTargetField"`
	ReferenceTo                          string       `xml:"referenceTo"`
	RelationshipLabel                    string       `xml:"relationshipLabel"`
	RelationshipName                     string       `xml:"relationshipName"`
	RelationshipOrder                    int          `xml:"relationshipOrder"`
	Required                             bool         `xml:"required"`
	Scale                                int          `xml:"scale"`
	StartingNumber                       int          `xml:"startingNumber"`
	StripMarkup                          bool         `xml:"stripMarkup"`
	SummarizedField                      string       `xml:"summarizedField"`
	SummaryFilterItems                   []FilterItem `xml:"summaryFilterItems"`
	SummaryForeignKey                    string       `xml:"summaryForeignKey"`
	SummaryOperation                     string       `xml:"summaryOperation"`
	TrackFeedHistory                     bool         `xml:"trackFeedHistory"`
	TrackHistory                         bool         `xml:"trackHistory"`
	TrackTrending                        bool         `xml:"trackTrending"`
	TrueValueIndexed                     bool         `xml:"trueValueIndexed"`
	Type                                 string       `xml:"type"`
	Unique                               bool         `xml:"unique"`
	ValueSet                             string       `xml:"valueSet"` // TODO: verify mapping
	VisibleLines                         int          `xml:"visibleLines"`
	WriteRequiresMasterRead              bool         `xml:"writeRequiresMasterRead"`
}

type Picklist struct {
	ControllingField   string          `xml:"controllingField"`
	PicklistValues     []PicklistValue `xml:"picklistValues"` // TODO: verify mapping
	RestrictedPicklist bool            `xml:"restrictedPicklist"`
	Sorted             bool            `xml:"sorted"`
}

type PicklistValue struct {
	FullName               string   `xml:"fullName"`
	ControllingFieldValues []string `xml:"controllingFieldValues"`
	Default                bool     `xml:"default"`
}

// NewApexClassFromFile read Apex class metadata from a file.
func NewCustomFieldFromFile(src string) (CustomField, error) {
	var field CustomField

	if err := decodeXMLFile(src, &field); err != nil {
		return field, err
	}

	fieldName := strings.TrimSuffix(filepath.Base(src), ".field-meta.xml")

	objectDir, err := filepath.Abs(filepath.Join(filepath.Dir(src), ".."))
	if err != nil {
		return field, fmt.Errorf("could not determine object name for CustomField '%s': %w", fieldName, err)
	}

	objectDirElements := filepath.SplitList(objectDir)
	objectName := objectDirElements[len(objectDirElements)-1]
	field.FullName = fmt.Sprintf("%s.%s", objectName, fieldName)

	return field, nil
}

// IsApexClassFile returns true if the file is an ApexClass metadata file.
func IsCustomFieldFile(src string) bool {
	return strings.HasSuffix(src, ".field-meta.xml")
}
