package provider

import "github.com/ernestre/terraform-provider-teampasswordmanager/tpm"

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
