package auth

import (
	"net/http"
	"testing"
)

func TestHash(t *testing.T) {
	password := "BadPass123"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash the password: %v", err)
	}

	if hash == "" {
		t.Fatal("Hash is empty")
	}

	if hash == password {
		t.Fatal("Hash is the same as password, hashing failed")
	}
}

func TestCheckHash(t *testing.T) {
	password := "right_password"
	wrongPass := "wrong_password"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash the password: %v", err)
	}

	err = CheckPasswordHash(hash, password)
	if err != nil {
		t.Fatalf("Failed to verify the correct password: %v", err)
	}

	err = CheckPasswordHash(hash, wrongPass)
	if err == nil {
		t.Fatal("successfully failed with wrong password")
	}
}

func TestBearerToken(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expected      string
		expectedError bool
	}{
		{
			name:          "No auth header",
			headers:       http.Header{},
			expected:      "",
			expectedError: true,
		},
		{
			name: "Malformed Auth Header",
			headers: http.Header{
				"Authorization": []string{"Basic abcfef"},
			},
			expected:      "",
			expectedError: true,
		},
		{
			name: "Incomplete Bearer token",
			headers: http.Header{
				"Authorization": []string{"Bearer"},
			},
			expected:      "",
			expectedError: true,
		},
		{
			name: "Valid bearer token",
			headers: http.Header{
				"Authorization": []string{"Bearer abc123"},
			},
			expected:      "abc123",
			expectedError: false,
		},
		{
			name: "Extra spaces",
			headers: http.Header{
				"Authorization": []string{" Bearer xyz789   "},
			},
			expected:      "xyz789",
			expectedError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			token, err := GetBearerToken(test.headers)
			if test.expectedError {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect error, but got: %v", err)
				}
				if token != test.expected {
					t.Errorf("expected token %q, got %q", test.expected, token)
				}
			}
		})
	}
}
