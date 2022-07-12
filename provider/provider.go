package provider

import (
	"context"
	"fmt"

	"github.com/ernestre/terraform-provider-teampasswordmanager/tpm"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	configHost       = "host"
	configPublicKey  = "public_key"
	configPrivateKey = "private_key"
	configAPIVersion = "api_version"

	envConfigHost       = "TPM_HOST"
	envConfigPublicKey  = "TPM_PUBLIC_KEY"
	envConfigPrivateKey = "TPM_PRIVATE_KEY"

	clientPassword = "password_client"
	clientProject  = "password_project"
)

type clientRegistry map[string]interface{}

func getProjectClient(m interface{}) tpm.ProjectClient {
	return m.(clientRegistry)[clientProject].(tpm.ProjectClient)
}

func getPasswordClient(m interface{}) tpm.PasswordClient {
	return m.(clientRegistry)[clientPassword].(tpm.PasswordClient)
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			configHost: {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(envConfigHost, nil),
				Description: "Host of the team password manager. (ie: http://localhost:8081)",
			},
			configPublicKey: {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc(envConfigPublicKey, nil),
				Description: "Public key from http://{ host }/index.php/user_info/api_keys",
			},
			configPrivateKey: {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc(envConfigPrivateKey, nil),
				Description: "Private key from http://{ host }/index.php/user_info/api_keys",
			},
			configAPIVersion: {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: "This attribute was added only for v4 support and will be removed in the future releases. Please upgrade your TeamPasswordManager to the latest version.",
				Default:    tpm.DefaultApiVersion,
				Description: fmt.Sprintf(
					"Api version to use (defaults to %s). Lower versions than v4 might not work correctly or at all. For more information https://teampasswordmanager.com/docs",
					tpm.DefaultApiVersion,
				),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"teampasswordmanager_password": resourcePassword(),
			"teampasswordmanager_project":  resourceProject(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"teampasswordmanager_password": dataSourcePassword(),
			"teampasswordmanager_project":  dataSourceProject(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	host := d.Get(configHost).(string)
	publicKey := d.Get(configPublicKey).(string)
	privateKey := d.Get(configPrivateKey).(string)
	apiVersion := d.Get(configAPIVersion).(string)

	if host == "" {
		return nil, diag.Errorf("%s cannot be empty", configHost)
	}

	if publicKey == "" {
		return nil, diag.Errorf("%s key cannot be empty", configPublicKey)
	}

	if privateKey == "" {
		return nil, diag.Errorf("%s key cannot be empty", configPrivateKey)
	}

	if apiVersion == "" {
		return nil, diag.Errorf("%s key cannot be empty", configAPIVersion)
	}

	var diags diag.Diagnostics

	config := tpm.Config{
		Host:       host,
		PublicKey:  publicKey,
		PrivateKey: privateKey,
		ApiVersion: apiVersion,
	}

	clients := clientRegistry{}
	clients[clientPassword] = tpm.NewPasswordClient(config)
	clients[clientProject] = tpm.NewProjectClient(config)

	return clients, diags
}
