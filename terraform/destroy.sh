set -e
cd "$(dirname "$0")"

ENV='stg'
terraform destroy -var-file="${ENV}.tfvars" -auto-approve
