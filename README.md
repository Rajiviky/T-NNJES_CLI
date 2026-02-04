## How to Build

We use a Makefile to automate the build process and ensure consistency across platforms.

- **To build for all required platforms (Linux/FreeBSD):**
  ```bash
  make all


- **To build for local machine:**
  ```bash
  make build

 - **Basic command:**
  ```bash
./certinfo --daid QC-DEMO --cid 3

$ ./certinfo --daid QC-DEMO --cid 3

CID: 3
DAID: QC-DEMO
Issuer: CN=QC DigSig Demo QC-DEMO https://www.idetrust.io 2026,O=QC DigSig Demo Inc.,L=Delmenhorst,C=DE
Subject: CN=https://idetrust.com/daid/QC%20DEMO/cid/3,O=IDeTRUST GmbH,C=DE
NotAfter: 20 Jan 27 05:12 UTC
NotBefore: 20 Jan 26 05:12 UTC
FingerPrint: 6C8276900B9E6447FE2571A7C391823FCF393CC33072C20328EAAACFA0225C83
```
 - **JSON output for scripts:**
  ```bash
./certinfo --daid QC-DEMO --cid 3 --json

$ ./certinfo --daid QC-DEMO --cid 3 --json
{
    "cid": 3,
    "daid": "QC-DEMO",
    "issuer": "CN=QC DigSig Demo QC-DEMO https://www.idetrust.io 2026,O=QC DigSig Demo Inc.,L=Delmenhorst,C=DE",
    "subject": "CN=https://idetrust.com/daid/QC%20DEMO/cid/3,O=IDeTRUST GmbH,C=DE",
    "notBefore": "2026-01-20T05:12:21Z",
    "notAfter": "2027-01-20T05:12:21Z",
    "fingerPrint": "6C8276900B9E6447FE2571A7C391823FCF393CC33072C20328EAAACFA0225C83"
}
```
 - **Filter expired certificate:**
  ```bash
./certinfo --daid QC-DEMO --cid 3 --onlyValid

$ ./certinfo --daid QC-DEMO --cid 3 --onlyValid
  
CID: 3
DAID: QC-DEMO
Issuer: CN=QC DigSig Demo QC-DEMO https://www.idetrust.io 2026,O=QC DigSig Demo Inc.,L=Delmenhorst,C=DE
Subject: CN=https://idetrust.com/daid/QC%20DEMO/cid/3,O=IDeTRUST GmbH,C=DE
NotAfter: 20 Jan 27 05:12 UTC
NotBefore: 20 Jan 26 05:12 UTC
FingerPrint: 6C8276900B9E6447FE2571A7C391823FCF393CC33072C20328EAAACFA0225C83
 ```