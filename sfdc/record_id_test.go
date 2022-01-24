package sfdc

import "testing"

func TestChecksum(t *testing.T) {
	tests := map[string]string{
		"a085P0000037PdJQAU": "QAU",
		"0012O00000B0zyVQAR": "QAR",
		"0055K000000IujzQAC": "QAC",
		"0062O0000039NixQAE": "QAE",
		"0012O000008WieGQAS": "QAS",
		"0052O000001IyiAQAS": "QAS",
	}

	for id, want := range tests {
		have, err := RecordIDChecksum(id)
		if err != nil {
			t.Errorf("RecordIDChecksum returned an error: '%s'", err)
			continue
		}
		if have != want {
			t.Errorf("invalid checksum, wanted '%s' got '%s'", want, have)
		}
	}
}

func TestChecksumWithInvalidIDs(t *testing.T) {
	tests := []string{
		"a0000037PdJQAU",        // too short
		"0013324O00000B0zyVQAR", // too long
	}

	for _, id := range tests {
		_, err := RecordIDChecksum(id)
		if err == nil {
			t.Errorf("should have returned an error due to invalid lenght: '%s'", id)
		}
	}
}

func TestIsRecordID(t *testing.T) {
	tests := map[string]bool{
		"a085P0000037PdJ":    true,
		"0012O00000B0zyVQAR": true,
		"0055K000000IujzQAC": true,
		"0062O0000039Nix":    true,
		"0012O000008WieGQAS": true,
		"0052O000001IyiAQAS": true,
		"0052O000QAS":        false, // invalid length
		"0012O00000-WieGQAS": false, // invalid char
		"0055K000000IujzQAD": false, // wrong checksum
	}

	for id, want := range tests {
		got := IsRecordID(id)
		if got != want {
			t.Errorf("IsRecordID('%s') returned %t, wanted: %t", id, got, want)
		}
	}

}
