package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey my-secret-key")
	got, _ := GetAPIKey(header)
	want := "my-secret-key"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
	header.Set("Authorization", "ApiKey ")
	got, _ = GetAPIKey(header)
	want = " "
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
	header.Set("Authorization", "ApiKey")
	_, gotError := GetAPIKey(header)
	wantError := errors.New("malformed authorization header")
	if gotError == nil || gotError.Error() != wantError.Error() {
		t.Fatalf("expected: %v, got: %v", wantError, gotError)
	}

	header.Set("Authorization", "")
	_, gotError = GetAPIKey(header)
	wantError = errors.New("no authorization header included")
	if gotError == nil || gotError.Error() != wantError.Error() {
		t.Fatalf("expected: %v, got: %v", wantError, gotError)
	}
}
