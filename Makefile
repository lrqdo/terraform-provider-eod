reset:
	go build
	mv ./terraform-provider-eod ~/.terraform.d/plugins/hashicorp.com/edu/eod/0.2/linux_amd64
	rm -Rf .terraform*
	terraform init
	terraform apply -auto-approve