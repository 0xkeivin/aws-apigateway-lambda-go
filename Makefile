build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o build/bin/app lambda-go/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/bin/app lambda-go/main.go
init:
	terraform init terraform_infra

plan:
	terraform plan terraform_infra

apply:
	terraform apply --auto-approve terraform_infra

destroy:
	terraform destroy --auto-approve terraform_infra