provider "teampasswordmanager" {
  host        = "http://localhost:8081"
  public_key  = "1356a192b7913b04c54574d18c28d46e6395428ab44f2ef0cabc9347835b9ea5"
  private_key = "5c005bc16db8b0e9f407c6747d4656fc48bbf0d6773e681f47fd86e1e7d6009b"

  // optional, default version is v5
  api_version = "v5" // "v4"

  // optional, should the certificate be verified?
  tls_verify = true

  # Or you can provide these values via env variables: TPM_HOST, TPM_PUBLIC_KEY, TPM_PRIVATE_KEY and TPM_API_VERSION
}
