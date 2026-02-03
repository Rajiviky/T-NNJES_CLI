package cert

import (
	"encoding/pem"
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

	certInfo, _ := pem.Decode()
}
