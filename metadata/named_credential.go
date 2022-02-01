package metadata

import (
	"path/filepath"
	"strings"
)

type NamedCredential struct {
	AllowMergeFieldsInBody      bool   `xml:"allowMergeFieldsInBody"`
	AllowMergeFieldsInHeader    bool   `xml:"allowMergeFieldsInHeader"`
	GenerateAuthorizationHeader bool   `xml:"generateAuthorizationHeader"`
	AWSAccessKey                string `xml:"awsAccesKey"`
	AWSAccessSecret             string `xml:"awsAccessSecret"`
	AWSRegion                   string `xml:"awsRegion"`
	AWSService                  string `xml:"awsService"`
	AuthProvider                string `xml:"authProvider"`
	AuthTokenEndpointURL        string `xml:"authTokenEndpointUrl"`
	Certificate                 string `xml:"certificate"`
	Endpoint                    string `xml:"endpoint"`
	FullName                    string
	JWTAudience                 string `xml:"jwtAudience"`
	JWTFormulaSubject           string `xml:"jwtFormulaSubject"`
	JWTIssuer                   string `xml:"jwtIssuer"`
	JWTSigningCertificate       string `xml:"jwtSigningCertificate"`
	JWTTextSubject              string `xml:"jwtTextSubject"`
	JWTValidityPeriodSeconds    string `xml:"jwtValidityPeriodSeconds"`
	Label                       string `xml:"label"`
	OAuthRefreshToken           string `xml:"oauthRefreshToken"`
	OAuthScope                  string `xml:"oauthScope"`
	OAuthToken                  string `xml:"oauthToken"`
	OutboundNetworkConnection   string `xml:"outboundNetworkConnection"`
	Password                    string `xml:"password"`
	PrincipalType               string `xml:"principalType"`
	Protocol                    string `xml:"protocol"`
	Username                    string `xml:"username"`
}

// NewNamedCredentialFromFile returns a NamedCredential from a file.
func NewNamedCredentialFromFile(src string) (NamedCredential, error) {
	var nc NamedCredential

	if err := decodeXMLFile(src, &nc); err != nil {
		return nc, err
	}

	nc.FullName = strings.TrimSuffix(filepath.Base(src), ".namedCredential-meta.xml")

	return nc, nil
}

// IsNamedCredentialFile returns true if src is a NamedCredential file.
func IsNamedCredentialFile(src string) bool {
	return strings.HasSuffix(src, ".namedCredential-meta.xml")
}
