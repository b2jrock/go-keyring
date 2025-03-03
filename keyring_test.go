package keyring

import "testing"

const (
	service  = "test-service"
	user     = "test-user"
	password = "test-password"
)

// TestSet tests setting a user and password in the keyring.
func TestSet(t *testing.T) {
	err := Set(service, user, password)
	if err != nil {
		t.Errorf("Should not fail, got: %s", err)
	}
}

// TestGetMultiline tests getting a multi-line password from the keyring
func TestGetMultiLine(t *testing.T) {
	multilinePassword := `this password
has multiple
lines and will be
encoded by some keyring implementiations
like osx`
	err := Set(service, user, multilinePassword)
	if err != nil {
		t.Errorf("Should not fail, got: %s", err)
	}

	pw, err := Get(service, user)
	if err != nil {
		t.Errorf("Should not fail, got: %s", err)
	}

	if multilinePassword != pw {
		t.Errorf("Expected password %s, got %s", multilinePassword, pw)
	}
}

// TestGetMultiline tests getting a multi-line password from the keyring
func TestGetUmlaut(t *testing.T) {
	umlautPassword := "at least on OSX üöäÜÖÄß will be encoded"
	err := Set(service, user, umlautPassword)
	if err != nil {
		t.Errorf("Should not fail, got: %s", err)
	}

	pw, err := Get(service, user)
	if err != nil {
		t.Errorf("Should not fail, got: %s", err)
	}

	if umlautPassword != pw {
		t.Errorf("Expected password %s, got %s", umlautPassword, pw)
	}
}

// TestGetSingleLineHex tests getting a single line hex string password from the keyring.
func TestGetSingleLineHex(t *testing.T) {
	hexPassword := "abcdef123abcdef123"
	err := Set(service, user, hexPassword)
	if err != nil {
		t.Errorf("Should not fail, got: %s", err)
	}

	pw, err := Get(service, user)
	if err != nil {
		t.Errorf("Should not fail, got: %s", err)
	}

	if hexPassword != pw {
		t.Errorf("Expected password %s, got %s", hexPassword, pw)
	}
}

// TestGet tests getting a password from the keyring.
func TestGet(t *testing.T) {
	err := Set(service, user, password)
	if err != nil {
		t.Errorf("Should not fail, got: %s", err)
	}

	pw, err := Get(service, user)
	if err != nil {
		t.Errorf("Should not fail, got: %s", err)
	}

	if password != pw {
		t.Errorf("Expected password %s, got %s", password, pw)
	}
}

// TestGetNonExisting tests getting a secret not in the keyring.
func TestGetNonExisting(t *testing.T) {
	_, err := Get(service, user+"fake")
	if err != ErrNotFound {
		t.Errorf("Expected error ErrNotFound, got %s", err)
	}
}

// TestDelete tests deleting a secret from the keyring.
func TestDelete(t *testing.T) {
	err := Delete(service, user)
	if err != nil {
		t.Errorf("Should not fail, got: %s", err)
	}
}

// TestDeleteNonExisting tests deleting a secret not in the keyring.
func TestDeleteNonExisting(t *testing.T) {
	err := Delete(service, user+"fake")
	if err != ErrNotFound {
		t.Errorf("Expected error ErrNotFound, got %s", err)
	}
}
