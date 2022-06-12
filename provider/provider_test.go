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
		"tpmsync": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if err := os.Getenv(env_config_host); err == "" {
		t.Fatalf("%s must be set for acceptance tests", env_config_host)
	}
	if err := os.Getenv(env_config_public_key); err == "" {
		t.Fatalf("%s must be set for acceptance tests", config_public_key)
	}
	if err := os.Getenv(env_config_private_key); err == "" {
		t.Fatalf("%s must be set for acceptance tests", config_private_key)
	}
}
