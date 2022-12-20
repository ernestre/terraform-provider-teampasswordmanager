package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	testAccProviders map[string]*schema.Provider
	testAccProvider  *schema.Provider
)

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"teampasswordmanager": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if err := os.Getenv(envConfigHost); err == "" {
		t.Fatalf("%s must be set for acceptance tests", envConfigHost)
	}

	if err := os.Getenv(envConfigPublicKey); err == "" {
		t.Fatalf("%s must be set for acceptance tests", envConfigPublicKey)
	}

	if err := os.Getenv(envConfigPrivateKey); err == "" {
		t.Fatalf("%s must be set for acceptance tests", envConfigPrivateKey)
	}

	if err := os.Getenv(envConfigAPIVersion); err == "" {
		t.Fatalf("%s must be set for acceptance tests", envConfigAPIVersion)
	}
}
