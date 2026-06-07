package authlete

import (
	"encoding/json"
	"testing"
)

func TestGrantTypeEnumValues(t *testing.T) {
	for _, expected := range AllowedGrantTypeEnumValues {
		t.Run(string(expected), func(t *testing.T) {
			if !expected.IsValid() {
				t.Fatalf("%s should be a valid GrantType", expected)
			}

			grantType, err := NewGrantTypeFromValue(string(expected))
			if err != nil {
				t.Fatalf("NewGrantTypeFromValue returned error: %v", err)
			}
			if *grantType != expected {
				t.Fatalf("unexpected GrantType: %s", *grantType)
			}

			var unmarshaled GrantType
			if err := json.Unmarshal([]byte(`"`+string(expected)+`"`), &unmarshaled); err != nil {
				t.Fatalf("json.Unmarshal returned error: %v", err)
			}
			if unmarshaled != expected {
				t.Fatalf("unexpected unmarshaled GrantType: %s", unmarshaled)
			}
		})
	}
}
