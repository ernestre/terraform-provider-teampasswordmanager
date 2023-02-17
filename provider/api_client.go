package provider

import (
	"os"

	"github.com/ernestre/terraform-provider-teampasswordmanager/tpm"
)

type clientRegistry map[string]interface{}

func getProjectClient(m interface{}) tpm.ProjectClient {
	return m.(clientRegistry)[clientProject].(tpm.ProjectClient)
}

func getPasswordClient(m interface{}) tpm.PasswordClient {
	return m.(clientRegistry)[clientPassword].(tpm.PasswordClient)
}

func getGroupClient(m interface{}) tpm.GroupClient {
	return m.(clientRegistry)[clientGroup].(tpm.GroupClient)
}

func newTestClientConfig() tpm.Config {
	return tpm.Config{
		Host:       os.Getenv(envConfigHost),
		PublicKey:  os.Getenv(envConfigPublicKey),
		PrivateKey: os.Getenv(envConfigPrivateKey),
	}
}

func newTestGroupClient() tpm.GroupClient {
	return tpm.NewGroupClient(newTestClientConfig())
}

func newTestProjectClient() tpm.ProjectClient {
	return tpm.NewProjectClient(newTestClientConfig())
}

func newTestPasswordClient() tpm.PasswordClient {
	return tpm.NewPasswordClient(newTestClientConfig())
}

func newTestUserClient() tpm.UserClient {
	return tpm.NewUserClient(newTestClientConfig())
}
