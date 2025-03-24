package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	noApiKeyHeader := http.Header{}
	noApiKeyHeader.Set("Content-Type", "application/json")

	apiKey, err := GetAPIKey(noApiKeyHeader)
	if err.Error() != "no authorization header included" {
		t.Errorf("header should not contain api key: %s", apiKey)
	}

	apiKeyHeader := http.Header{}
	apiKeyHeader.Set("Authorization", "ApiKey some-api-key")

	apiKey, err = GetAPIKey(apiKeyHeader)
	if err != nil {
		t.Errorf("header should contain api key: %v", err)
	}
	if apiKey != "some-api-key" {
		t.Errorf("api key's do not match --> %s != %s <--", apiKey, "some-api-key")
	}

	apiKeyMalformed := http.Header{}
	apiKeyMalformed.Set("Authorization", "APIkey some-api-key")

	_, err = GetAPIKey(apiKeyMalformed)
	if err.Error() != "malformed authorization header" {
		t.Errorf("api key should be malformed")
	}
}