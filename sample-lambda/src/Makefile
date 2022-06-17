lambdaname=sample-lambda
localstackendpoint=http://localhost:4566
zippedlambdaname=function.zip
compiledfilename=main
define LAMBDAENVVARS
{
  "Variables": {
    "FOO" : "HELLO WORLD!"
  }
}
endef
export LAMBDAENVVARS

build:
	@echo ":: Generating Deployment Package ::"
	cd .; GOOS=linux go build -o ../bin/${compiledfilename}
	cd ../deployment; zip -rj ${zippedlambdaname} ../bin/${compiledfilename}
	@echo ":: Done! ::"

deploy: build
	-aws --endpoint-url=${localstackendpoint} lambda delete-function --function-name ${lambdaname} 
	@echo ":: Creating lambda ::"
	cd ../deployment; aws --endpoint-url=${localstackendpoint} lambda create-function \
--function-name ${lambdaname} \
--environment "$$LAMBDAENVVARS" \
--handler ${compiledfilename} \
--runtime go1.x \
--role anyrole \
--timeout 20 \
--zip-file fileb://${zippedlambdaname} 
	@echo ":: Done! ::"
test: 
	aws --endpoint-url=${localstackendpoint} lambda invoke \
--function-name ${lambdaname} /dev/stdout
