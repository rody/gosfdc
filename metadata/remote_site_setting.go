package metadata

import (
	"path/filepath"
	"strings"
)

// see https://developer.salesforce.com/docs/atlas.en-us.api_meta.meta/api_meta/meta_remotesitesetting.htm
type RemoteSiteSetting struct {
	FullName                string
	Description             string `xml:"description"`
	DisableProtocolSecurity bool   `xml:"disableProtocolSecurity"`
	IsActive                bool   `xml:"isActive"`
	URL                     string `xml:"url"`
}

func NewRemoteSiteSettingFromFile(src string) (RemoteSiteSetting, error) {
	var rss RemoteSiteSetting
	if err := decodeXMLFile(src, &rss); err != nil {
		return rss, err
	}

	rss.FullName = strings.TrimSuffix(filepath.Base(src), ".remoteSite-meta.xml")

	return rss, nil
}

func IsRemoteSiteSettingFile(src string) bool {
	return strings.HasSuffix(src, ".remoteSite-meta.xml")
}
