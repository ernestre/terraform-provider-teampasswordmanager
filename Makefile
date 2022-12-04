test:
	TF_ACC=1 \
	TPM_HOST=http://localhost:8081 \
	TPM_PUBLIC_KEY=1356a192b7913b04c54574d18c28d46e6395428ab44f2ef0cabc9347835b9ea5 \
	TPM_PRIVATE_KEY=5c005bc16db8b0e9f407c6747d4656fc48bbf0d6773e681f47fd86e1e7d6009b \
	go test --race ./...

test-unit:
	go test --race ./...

build-tf:
	go build -o terraform-provider-teampasswordmanager
	mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/teampasswordmanager/0.2/linux_amd64
	mv terraform-provider-teampasswordmanager ~/.terraform.d/plugins/hashicorp.com/edu/teampasswordmanager/0.2/linux_amd64/

tf-run: build-tf
	cd ./examples/ && rm -rf .terraform* && terraform init && terraform apply --auto-approve

generate:
	go generate ./...
