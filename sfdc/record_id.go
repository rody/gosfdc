package sfdc

import (
	"fmt"
	"regexp"
	"unicode"
)

var (
	checksumRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ012345")
	recordIDRegex = regexp.MustCompile("^[a-zA-Z0-9]{5}0[a-zA-Z0-9]{9}([a-zA-Z0-5]{3})?$")
)

// IsRecordID returns true if the given string represents a valid recordID.
//
// If the given record ID is 18 chars long, the checksum is verified as well
func IsRecordID(s string) bool {
	if !recordIDRegex.MatchString(s) {
		return false
	}

	if len(s) == 18 {
		csum, err := RecordIDChecksum(s)
		if err != nil {
			panic(err) // should not happen since we are checking for length prior to calling the function
		}

		return s[15:] == csum
	}

	return true
}

// RecordIDChecksum returns the checksum of the given record id
func RecordIDChecksum(s string) (string, error) {
	if len(s) != 15 && len(s) != 18 {
		return "", fmt.Errorf("record IDs must be 15 or 18 chars, got '%s'", s)
	}

	var csum []rune

	for part := 0; part < 3; part++ { // checksum has 1 char per group of 5 chars (15/5=3 parts)
		v := 0
		for pos, r := range s[part*5 : (part+1)*5] {
			if unicode.IsLetter(r) && unicode.IsUpper(r) {
				v += 1 << pos
			}
		}
		csum = append(csum, checksumRunes[v])
	}

	return string(csum), nil
}
