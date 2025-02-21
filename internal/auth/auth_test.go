package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		wantKey string
		wantErr error
	}{
		{
			name:    "normal case",
			headers: http.Header{"Authorization": []string{"ApiKey mysupersecretapikey123"}},
			wantKey: "mysupersecretapikey123",
			wantErr: nil,
		},
		{
			name:    "no header",
			headers: http.Header{},
			wantKey: "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:    "wrong format - no ApiKey",
			headers: http.Header{"Authorization": []string{"mysupersecretapikey123"}},
			wantKey: "",
			wantErr: errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("Error does not match: expected %v, got %v", tt.wantErr, err)
				}
			} else if err != nil {
				t.Errorf("Unexpected error: got %v", err)
			}
			if got != tt.wantKey {
				t.Errorf("Key does not match: expected %s, got %s", tt.wantKey, got)
			}

		})
	}

}

//func TestGetAPIKey(t *testing.T) {
//	type args struct {
//		headers http.Header
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := GetAPIKey(tt.args.headers)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("GetAPIKey() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
