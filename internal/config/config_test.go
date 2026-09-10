package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidateRequiresFilesAndFlightCertificate(t *testing.T) {
	dir := t.TempDir()
	paths := []string{"certificate.pem", "key.pem", "ca-chain.pem"}
	for _, name := range paths {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("test"), 0600); err != nil {
			t.Fatal(err)
		}
	}
	cfg := Config{Certificate: filepath.Join(dir, paths[0]), Key: filepath.Join(dir, paths[1]),
		CAChain: filepath.Join(dir, paths[2]), Email: "developer@example.com", PoolSize: 1}
	if err := cfg.Validate("membership"); err != nil {
		t.Fatalf("membership config should be valid: %v", err)
	}
	if err := cfg.Validate("flights"); err == nil {
		t.Fatal("flights should require an Apple certificate ID")
	}
	cfg.AppleCertificateID = "pass.com.example"
	if err := cfg.Validate("flights"); err != nil {
		t.Fatalf("flight config should be valid: %v", err)
	}
}
