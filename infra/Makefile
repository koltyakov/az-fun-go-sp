ifneq (,$(wildcard ./.env))
	include .env
	export
endif

init:
	terraform init \
		-backend-config="resource_group_name=Terraform" \
		-backend-config="storage_account_name=${TERRAFORM_BACKEND_STORAGE_ACCOUNT}" \
		-backend-config="container_name=tfstate" \
		-backend-config="key=az-fun-go-sp.dev.tfstate"

validate:
	terraform validate

plan:
	terraform plan

apply:
	terraform apply -auto-approve

terra: pack
	terraform apply

destroy:
	terraform destroy

pack:
	cd ../ && make pack