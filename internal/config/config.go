package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Host, Port                    string
	Certificate, Key, CAChain     string
	Email, AppleCertificateID     string
	PoolSize, CleanupDelaySeconds int
}

func Load() (Config, error) {
	fileValues, err := readEnvFile(".env")
	if err != nil {
		return Config{}, err
	}
	get := func(name, fallback string) string {
		if value := strings.TrimSpace(os.Getenv(name)); value != "" {
			return value
		}
		if value := strings.TrimSpace(fileValues[name]); value != "" {
			return value
		}
		return fallback
	}
	poolSize, err := strconv.Atoi(get("PASSKIT_POOL_SIZE", "4"))
	if err != nil || poolSize < 1 {
		return Config{}, fmt.Errorf("PASSKIT_POOL_SIZE must be a positive integer")
	}
	cleanupDelay, err := strconv.Atoi(get("PASSKIT_CLEANUP_DELAY_SECONDS", "0"))
	if err != nil || cleanupDelay < 0 {
		return Config{}, fmt.Errorf("PASSKIT_CLEANUP_DELAY_SECONDS must be a non-negative integer")
	}
	return Config{
		Host: get("PASSKIT_GRPC_HOST", "grpc.pub1.passkit.io"), Port: get("PASSKIT_GRPC_PORT", "443"),
		Certificate: get("PASSKIT_CERTIFICATE", "certs/certificate.pem"), Key: get("PASSKIT_KEY", "certs/key.pem"),
		CAChain: get("PASSKIT_CA_CHAIN", "certs/ca-chain.pem"), Email: get("PASSKIT_EMAIL", ""),
		AppleCertificateID: get("PASSKIT_APPLE_CERTIFICATE_ID", ""), PoolSize: poolSize,
		CleanupDelaySeconds: cleanupDelay,
	}, nil
}

func (c Config) Validate(workflow string) error {
	for label, path := range map[string]string{"certificate": c.Certificate, "private key": c.Key, "CA chain": c.CAChain} {
		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("%s file %q is not accessible: %w", label, path, err)
		}
	}
	if c.Email == "" {
		return fmt.Errorf("PASSKIT_EMAIL is required")
	}
	if workflow == "flights" && c.AppleCertificateID == "" {
		return fmt.Errorf("PASSKIT_APPLE_CERTIFICATE_ID is required for flights")
	}
	return nil
}

func readEnvFile(path string) (map[string]string, error) {
	values := map[string]string{}
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		return values, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, found := strings.Cut(line, "=")
		if !found {
			continue
		}
		values[strings.TrimSpace(key)] = strings.Trim(strings.TrimSpace(value), "\"'")
	}
	return values, scanner.Err()
}
