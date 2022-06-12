package provider

import (
	"context"

	"github.com/ernestre/terraform-provider-teampasswordmanager/tpm"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	config_host        = "host"
	config_public_key  = "public_key"
	config_private_key = "private_key"

	env_config_host        = "TPM_HOST"
	env_config_public_key  = "TPM_PUBLIC_KEY"
	env_config_private_key = "TPM_PRIVATE_KEY"

	client_password = "password_client"
	client_project  = "password_project"
)

type clientRegistry map[string]interface{}

func getProjectClient(m interface{}) tpm.ProjectClient {
	return m.(clientRegistry)[client_project].(tpm.ProjectClient)
}

func getPasswordClient(m interface{}) tpm.PasswordClient {
	return m.(clientRegistry)[client_password].(tpm.PasswordClient)
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			config_host: {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(env_config_host, nil),
				Description: "Host of the team password manager. (ie: http://localhost:8081)",
			},
			config_public_key: {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc(env_config_public_key, nil),
				Description: "Public key from http://{ host }/index.php/user_info/api_keys",
			},
			config_private_key: {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc(env_config_private_key, nil),
				Description: "Private key from http://{ host }/index.php/user_info/api_keys",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"tpmsync_password": resourcePassword(),
			"tpmsync_project":  resourceProject(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	host := d.Get(config_host).(string)
	public_key := d.Get(config_public_key).(string)
	private_key := d.Get(config_private_key).(string)

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
	clients[client_password] = tpm.NewPasswordClient(config)
	clients[client_project] = tpm.NewProjectClient(config)

	return clients, diags
}
