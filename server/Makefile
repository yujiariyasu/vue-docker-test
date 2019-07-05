.PHONY: deps clean build

deps:
		go get -u ./...

clean:
		rm -rf ./hello-world/hello-world

build:
		GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world
		GOOS=linux GOARCH=amd64 go build -o task-register/task-register ./task-register
		GOOS=linux GOARCH=amd64 go build -o task-fetcher/task-fetcher ./task-fetcher
# 複数のLambdaを用意する場合はここにビルド対象を追加していく

package:
		sam package --template-file template.yaml --output-template-file output-template.yaml --s3-bucket vue-go-test --profile yujiariyasu

deploy:
		sam deploy --template-file output-template.yaml --stack-name vue-go-test --capabilities CAPABILITY_IAM --profile yujiariyasu
