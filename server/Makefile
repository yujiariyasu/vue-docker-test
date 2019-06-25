.PHONY: deps clean build

deps:
		go get -u ./...

clean:
		rm -rf ./hello-world/hello-world

build:
		GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world
# 複数のLambdaを用意する場合はここにビルド対象を追加していく

packege:
		sam package --template-file server/template.yaml --output-template-file server/output-template.yaml --s3-bucket template-store --profile yujiariyasu

package:
		sam package --template-file template.yaml --output-template-file output-template.yaml --s3-bucket vue-go-test --profile yujiariyasu

deploy:
		sam deploy --template-file server/output-template.yaml --stack-name go-lambda --capabilities CAPABILITY_IAM --profile yujiariyasu

deploy:
		sam deploy --template-file output-template.yaml --stack-name vue-go-test --capabilities CAPABILITY_IAM --profile yujiariyasu
