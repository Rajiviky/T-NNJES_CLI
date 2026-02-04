package cert_test

import (
	"T-NNJES_CLI/pkg/cert"
	"strings"
	"testing"
)

func TestParseCertFields(t *testing.T) {
	mockResBody := []byte(`NotBefore: 2026-01-20T05:12:21Z
NotAfter: 2027-01-20T05:12:21Z
Subject: CN=https://idetrust.com/daid/QC%20DEMO/cid/3,O=IDeTRUST GmbH,C=DE
-----BEGIN CERTIFICATE-----
MIIHYzCCBUugAwIBAgIQFjpLAfvNoP2KH7GI1C3wWjANBgkqhkiG9w0BAQsFADB/
MRwwGgYDVQQKDBNRQyBEaWdTaWcgRGVtbyBJbmMuMTwwOgYDVQQDDDNRQyBEaWdT
aWcgRGVtbyBRQy1ERU1PIGh0dHBzOi8vd3d3LmlkZXRydXN0LmlvIDIwMjYxCzAJ
BgNVBAYTAkRFMRQwEgYDVQQHDAtEZWxtZW5ob3JzdDAeFw0yNjAxMjAwNTEyMjFa
Fw0yNzAxMjAwNTEyMjFaMFkxCzAJBgNVBAYTAkRFMRYwFAYDVQQKEw1JRGVUUlVT
VCBHbWJIMTIwMAYDVQQDDClodHRwczovL2lkZXRydXN0LmNvbS9kYWlkL1FDJTIw
REVNTy9jaWQvMzAQMAsGCWCGSAFlAwQCAQMBAKOCBBMwggQPMIICmAYEKIGeGASC
Ao57ImRpZ3NpZ2luZm8iOnsic3BlY2lmaWNhdGlvbnZlcnNpb24iOiJJU08vSUVD
IDIwMjQ4OjIwMjIiLCJkYXVyaSI6Imh0dHBzOi8vaWRldHJ1c3QuY29tIiwiZGFp
ZCI6IlFDIERFTU8iLCJjaWQiOjMsInZlcmlmaWNhdGlvbnVyaSI6WyJodHRwczov
L2lkZXRydXN0LmNvbSJdLCJyZXZvY2F0aW9udXJpIjpbImh0dHBzOi8vaWRldHJ1
c3QuY29tIl0sInByZXZlcmlmeSI6eyJlbiI6IiJ9LCJhY2NlcHR2ZXJpZnkiOnsi
ZW4iOiIifSwicmVqZWN0dmVyaWZ5Ijp7ImVuIjoiIn0sInBvc3R2ZXJpZnkiOnsi
ZW4iOiIifSwic3RydWN0dXJlZGRvY3VyaSI6eyJlbiI6IiJ9fSwiZGF0YWZpZWxk
cyI6W3siZmllbGRpZCI6InNwZWNpZmljYXRpb252ZXJzaW9uIn0seyJmaWVsZGlk
IjoiZGF1cmkifSx7ImZpZWxkaWQiOiJkYWlkIn0seyJmaWVsZGlkIjoiY2lkIn0s
eyJmaWVsZGlkIjoic2lnbmF0dXJlIn0seyJmaWVsZGlkIjoidGltZXN0YW1wIiwi
cmFuZ2UiOiJbMjAyNi0wMS0yMC4uMjAyNy0wMS0yMF0ifSx7ImZpZWxkaWQiOiJl
eGFtcGxlIiwidHlwZSI6InN0cmluZyIsImZpZWxkbmFtZSI6eyJlbiI6IkV4YW1w
bGUgZmllbGQiLCJ0aCI6IuC4leC4seC4p+C4reC4ouC5iOC4suC4hyJ9LCJic2ln
biI6dHJ1ZSwicmFuZ2UiOiJbYS16QS1aMC05XSJ9XX0wgasGA1UdIASBozCBoDCB
nQYJKwYBBAGD1CIBMIGPMGQGCCsGAQUFBwICMFgMVklTTy9JRUMgMjAyNDggRGln
U2lnIGNlcnRpZmljYXRlLiBTZWUgdGhlIFYzIGZpZWxkIDEuMC4yMDI0OCBmb3Ig
dGhlIGRhdGEgZGVzY3JpcHRpb24uMCcGCCsGAQUFBwIBFhtVUkk6aHR0cDovL2lk
ZXRydXN0LmNvbS9jcHMwKwYDVR0QBCQwIoAPMjAyNjAxMjAwMDAwMDBagQ8yMDI3
MDEyMDIzNTk1OVowNQYDVR0fBC4wLDAqoCigJoYkVVJJOmh0dHBzOi8vaWRldHJ1
c3QuY29tL2lkZWRzY2EuY3JsMA8GA1UdJQQIMAYGBCiBnhgwHQYDVR0OBBYEFNo5
o+5ea0sNMlW/75VgGJCv2AcJMB8GA1UdIwQYMBaAFLPpammJiVndzunNJWuNtxgk
egEjMA4GA1UdDwEB/wQEAwIHgDANBgkqhkiG9w0BAQsFAAOCAgEAtnnngrzNjccF
JL1Ye6Dppi61atUj7ALpTmpRqTq/uY9mO/xFTmr3JrDHxfZCKLv8fvHTYg8wv1lV
0j+rBiDDL49WiBQ+eg6Ogb/FqR687437DQr961yxFPtJKp6WQnDjkHu5O0aSauUu
IZdG1r+vyKfI4Q25MN0IwA3hrhHUxYLBv5tAChJqhjttiWfh/7nARdp7F2CnD/KN
hwTxxcJ3bo5qUKeVd+vxb4oLFLjnGeuoMSQz2etjTTN9Zx+xby1EtEtKIZQX7+PD
XbTk6Xw0Ix5QEmslXhHvEmi1OKxlx0942OOFz7tOrWlsX0u/WA3r4u0uuizDqnU6
k57cfE6uuNBfml1CXBGOl4MMFB2L+Fp0z0Kk5TFoLxjwK/MNYyn/mtmx/rTNmtma
RqpdcSyR1gKiKZ8r5Jc6v/aGtG2VVVS1R4btg1YhjFhlq9b+0QZkAxipFottFX5h
IYVY+UMj/Z1s6MJ9JT70aaiCsI2K7twzlVUJbCSdUOdTmL5s79m7SacUB0tZ7yJH
AenFLK8z4BfmHVUMnhuuAeryUGRpQwCCVk85YJF5d+Gi/ZsclfB4/Ykqs+OD++Fm
nbRFQ1M9EYZ5/bDE4z1EgtDKopC6guSf9S5AVv7TOPCewfR56WL6puEU9dbIo7AE
r75GZlqA6ocOjNBX6tkMSONcySO8USQ=
-----END CERTIFICATE-----`)

	tests := []struct {
		name    string
		resBody []byte
		cid     int
		daid    string
		wantErr bool
	}{
		{
			name:    "Successful Parse with Metadata Noise",
			resBody: mockResBody,
			cid:     3,
			daid:    "QC-DEMO",
			wantErr: false,
		},
		{
			name:    "Invalid Data",
			resBody: []byte("invalid data"),
			cid:     1,
			daid:    "FAIL",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cert.ParseCertFields(tt.resBody, tt.cid, tt.daid)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCertFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if got.CID != tt.cid || got.DAID != tt.daid {
					t.Errorf("Mapping failed: got CID %d, DAID %s", got.CID, got.DAID)
				}
				if !strings.Contains(got.Subject, "idetrust.com") {
					t.Errorf("Subject parsing error, got: %s", got.Subject)
				}
				if len(got.FingerPrint) != 64 {
					t.Errorf("Invalid Fingerprint length: %d", len(got.FingerPrint))
				}
			}
		})
	}
}
