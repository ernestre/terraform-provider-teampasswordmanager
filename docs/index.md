---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "teampasswordmanager Provider"
subcategory: ""
description: |-
  
---

# teampasswordmanager Provider



## Example Usage

```terraform
provider "teampasswordmanager" {
  host        = "http://localhost:8081"
  public_key  = "1356a192b7913b04c54574d18c28d46e6395428ab44f2ef0cabc9347835b9ea5"
  private_key = "5c005bc16db8b0e9f407c6747d4656fc48bbf0d6773e681f47fd86e1e7d6009b"

  // optional, default version is v5
  api_version = "v5" // "v4"

  // optional, skip TLS certificate verification?
  tls_skip_verify = false

  # Or you can provide these values via env variables: TPM_HOST, TPM_PUBLIC_KEY, TPM_PRIVATE_KEY and TPM_API_VERSION
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `host` (String) Host of the team password manager. (ie: http://localhost:8081)
- `private_key` (String, Sensitive) Private key from http://{ host }/index.php/user_info/api_keys
- `public_key` (String, Sensitive) Public key from http://{ host }/index.php/user_info/api_keys

### Optional

- `api_version` (String, Deprecated) Api version to use (defaults to v5). Lower versions than v4 might not work correctly or at all. For more information https://teampasswordmanager.com/docs
- `tls_skip_verify` (Boolean) Whether the TLS certificate verification should be skipped (defaults to false).
