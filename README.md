# seen-api

## Usage

```shell
curl -i https://seen-api-akmjhvanuq-an.a.run.app/accounting/partners
```

## Setup

```shell
$ go install golang.org/x/tools/cmd/goimports@latest
$ go install honnef.co/go/tools/cmd/staticcheck@latest

$ cp .env.sample .env
$ make

$ curl -i http://localhost:8080/accounting/partners
```

## Deploy (adhoc)

```shell
# stg
$ gcloud secrets versions access latest \
  --secret="colere-stg-write-pk-json" \
  --project colere-survey-stg > terraform/credentials/stg.json

$ source terraform/source_me
```
