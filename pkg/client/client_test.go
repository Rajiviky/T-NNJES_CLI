package client_test

import (
	"T-NNJES_CLI/pkg/client"
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCertficate(t *testing.T) {
	mockResponse := []byte(`NotBefore: 2026-01-20T05:12:21Z
NotAfter: 2027-01-20T05:12:21Z
Subject: CN=https://idetrust.com/daid/QC%20DEMO/cid/3,O=IDeTRUST GmbH,C=DE
-----BEGIN CERTIFICATE-----
MIIHYzCCBUugAwIBAgIQFjpLAfvNoP2KH7GI1C3wWjANBgkqhkiG9w0BAQsFADB/
... (shortened for brevity) ...
-----END CERTIFICATE-----
NotBefore: 2025-11-13T00:00:00Z
NotAfter: 2046-12-31T23:59:59Z
Subject: CN=QC DigSig Demo QC-DEMO https://www.idetrust.io 2026...
-----BEGIN CERTIFICATE-----
... (intermediate cert) ...
-----END CERTIFICATE-----`)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/daid/QCDEMO/cid/3" {
			w.WriteHeader(http.StatusOK)
			w.Write(mockResponse)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	// 2. Point client to mock server
	originalURL := client.BaseURL
	client.BaseURL = server.URL
	defer func() { client.BaseURL = originalURL }()

	tests := []struct {
		name    string
		daid    string
		cid     int
		wantErr bool
	}{
		{
			name:    "Successful  Fetch",
			daid:    "QC-DEMO",
			cid:     3,
			wantErr: false,
		},
		{
			name:    "Invalid CID",
			daid:    "QC-DEMO",
			cid:     999,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetCertficate(context.Background(), tt.daid, tt.cid)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetCertficate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if !bytes.Contains(got, []byte("-----BEGIN CERTIFICATE-----")) {
					t.Error("Response body missing PEM block")
				}
				if !bytes.Contains(got, []byte("NotBefore:")) {
					t.Error("Response body missing metadata labels")
				}
			}
		})
	}
}
