package cmd

import (
	"T-NNJES_CLI/pkg/cert"
	"T-NNJES_CLI/pkg/client"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	daid   string
	cid    int
	isJson bool
)

var rootCmd = &cobra.Command{
	Use:   "certInfo",
	Short: "Fetch and dispaly X.509 certificate information from IDeTrust",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		rawData, err := client.GetCertficate(ctx, daid, cid)
		if err != nil {
			return err
		}
		certDetails, err := cert.ParseCertFields(rawData, cid, daid)
		if err != nil {
			return err
		}
		if isJson {
			return json.NewDecoder(os.Stdout).Decode(certDetails)
		}
		fmt.Printf("\nCID: %d\nDAID: %s\nIssuer: %s\nSubject: %s\nNotAfter: %s\n, %s\nNotBefore:%s\n",
			certDetails.CID, certDetails.DAID, certDetails.Issuer, certDetails.Subject, certDetails.NotAfter.Format(time.RFC822), certDetails.NotBefore.Format(time.RFC822))
		return nil
	},
}

func Execute() {
	rootCmd.Flags().StringVar(&daid, "daid", "", "Developer Account ID(required)")
	rootCmd.Flags().IntVar(&cid, "cid", 0, " Credential ID(required)")
	rootCmd.Flags().BoolVar(&isJson, "json", false, "Deatils in JSON format")
	rootCmd.MarkFlagRequired("daid")
	rootCmd.MarkFlagRequired("cid")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
