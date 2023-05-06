set -e
cd "$(dirname "$0")"

ENV=$1
if [ "${ENV}" != 'prod' ]; then
  ENV='stg'
fi

terraform fmt
terraform init -reconfigure "-backend-config=${ENV}.tfbackend"
terraform validate
terraform plan -var-file="${ENV}.tfvars"
terraform apply -var-file="${ENV}.tfvars" -auto-approve -replace=google_cloud_run_v2_service.default
