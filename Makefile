HOST=https://teampasswordmanager.localhost
PUBLIC_KEY=1356a192b7913b04c54574d18c28d46e6395428ab3e412034b8b325de8791d25
PRIVATE_KEY=bcc17645f12660f6bbf00801e48830429c66b52775bd8c39fca3f2d6d16c6578
TPM_TLS_SKIP_VERIFY=true

test: test-v5 test-v4

test-v4:
	TF_ACC=1 \
	TPM_HOST=${HOST} \
	TPM_PUBLIC_KEY=${PUBLIC_KEY} \
	TPM_PRIVATE_KEY=${PRIVATE_KEY} \
	TPM_TLS_SKIP_VERIFY=${TPM_TLS_SKIP_VERIFY} \
	TPM_API_VERSION=v4 \
	go test --race ./...

test-v5:
	TF_ACC=1 \
	TPM_HOST=${HOST} \
	TPM_PUBLIC_KEY=${PUBLIC_KEY} \
	TPM_PRIVATE_KEY=${PRIVATE_KEY} \
	TPM_TLS_SKIP_VERIFY=${TPM_TLS_SKIP_VERIFY} \
	TPM_API_VERSION=v5 \
	go test --race ./...

test-unit:
	TPM_HOST=${HOST} \
	TPM_PUBLIC_KEY=${PUBLIC_KEY} \
	TPM_PRIVATE_KEY=${PRIVATE_KEY} \
	go test --race ./...

build-tf:
	go build -o terraform-provider-teampasswordmanager
	mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/teampasswordmanager/0.2/linux_amd64
	mv terraform-provider-teampasswordmanager ~/.terraform.d/plugins/hashicorp.com/edu/teampasswordmanager/0.2/linux_amd64/

tf-run: build-tf
	cd ./examples/ && rm -rf .terraform* && terraform init && terraform apply --auto-approve

generate:
	go generate ./...
