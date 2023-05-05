# seen-api

## Usage

```shell
# 株式会社フリー
$ curl -i "https://seen-api-akmjhvanuq-an.a.run.app/accounting/partners?name=%E6%A0%AA%E5%BC%8F%E4%BC%9A%E7%A4%BE%E3%83%95%E3%83%AA%E3%83%BC"
```

## Setup

```shell
$ go install golang.org/x/tools/cmd/goimports@latest
$ go install honnef.co/go/tools/cmd/staticcheck@latest

$ cp .env.sample .env
$ gcloud secrets versions access latest \
  --secret="freee-api-token" \
  --project colere-survey-stg > freee-secret.json

$ export GOOGLE_APPLICATION_CREDENTIALS="terraform/credentials/stg.json"

# access_token の値を .env にコピペした上で実行してください
$ make

$ curl -i "http://localhost:8080/accounting/partners?name=%E6%A0%AA%E5%BC%8F%E4%BC%9A%E7%A4%BE%E3%83%95%E3%83%AA%E3%83%BC"
```

## Deploy (adhoc)

設定

```shell
# stg
$ gcloud secrets versions access latest \
  --secret="colere-stg-write-pk-json" \
  --project colere-survey-stg > terraform/credentials/stg.json

$ source terraform/source_me
```

Docker Image の push

```shell
$ GCP_PROJECT_ID=colere-survey-stg \
  GCP_SERVICE_NAME=seen-api \
  GITHUB_REF_NAME=develop \
  IMAGE_NAME=gcr.io/${GCP_PROJECT_ID}/${GCP_SERVICE_NAME}:${GITHUB_REF_NAME}
$ docker build -t ${IMAGE_NAME} .
$ docker push ${IMAGE_NAME}
```
