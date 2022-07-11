package provider

import (
	"context"

	"github.com/ernestre/terraform-provider-teampasswordmanager/tpm"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	configHost       = "host"
	configPublicKey  = "public_key"
	configPrivateKey = "private_key"

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
	public_key := d.Get(configPublicKey).(string)
	private_key := d.Get(configPrivateKey).(string)

	if host == "" {
		return nil, diag.Errorf("host cannot be empty")
	}

	if public_key == "" {
		return nil, diag.Errorf("public key cannot be empty")
	}

	if private_key == "" {
		return nil, diag.Errorf("private key cannot be empty")
	}

	var diags diag.Diagnostics

	config := tpm.Config{
		Host:       host,
		PublicKey:  public_key,
		PrivateKey: private_key,
	}

	clients := clientRegistry{}
	clients[clientPassword] = tpm.NewPasswordClient(config)
	clients[clientProject] = tpm.NewProjectClient(config)

	return clients, diags
}
