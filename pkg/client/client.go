package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func GetCertficate(ctx context.Context, daid string, cid int) ([]byte, error) {

	// in case we use daid with - in the flag
	cleanDaid := strings.ReplaceAll(daid, "_", "")

	url := fmt.Sprintf("https://idetrust.com/daid/%s/cid/%d", cleanDaid, cid)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "T-NNJES-CertInfo/1.0")

	httpClient := &http.Client{Timeout: 10 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned status: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)

}
