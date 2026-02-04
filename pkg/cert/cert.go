package cert

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"
)

type CertInfo struct {
	CID         int       `json:"cid"`
	DAID        string    `json:"daid"`
	Issuer      string    `json:"issuer"`
	Subject     string    `json:"subject"`
	NotBefore   time.Time `json:"notBefore"`
	NotAfter    time.Time `json:"notAfter"`
	FingerPrint string    `json:"fingerPrint"`
}

func ParseCertFields(resBody []byte, cid int, daid string) (*CertInfo, error) {

	certInfo, _ := pem.Decode(resBody)
	if certInfo == nil {
		return nil, fmt.Errorf("no valid data found in the output")
	}

	c, err := x509.ParseCertificate(certInfo.Bytes)
	if err != nil {
		return nil, err
	}

	return &CertInfo{
		CID:         cid,
		DAID:        daid,
		Issuer:      c.Issuer.String(),
		Subject:     c.Subject.String(),
		NotBefore:   c.NotBefore,
		NotAfter:    c.NotAfter,
		FingerPrint: fmt.Sprintf("%X", sha256.Sum256(c.Raw)),
	}, nil
}
