reset:
	go build
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/lrqdo/eod/0.0.9/linux_amd64
	mv ./terraform-provider-eod ~/.terraform.d/plugins/registry.terraform.io/lrqdo/eod/0.0.9/linux_amd64/
	rm -Rf .terraform*
	terraform init
	terraform apply -auto-approve