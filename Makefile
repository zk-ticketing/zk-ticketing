# deletes the generated openapi server folder and regenerates it
openapi-generate-server:
	rm -rf backend/openapi && openapi-generator generate -i openapi/v1/ticketing.yaml -g go-server -o backend --package-name openapi --additional-properties sourceFolder=openapi,serverPort=3000

.PHONY: start-swagger-ui
start-swagger-ui:
	docker run -p 8081:8080 -e SWAGGER_JSON=/openapi/v1/ticketing.yaml -v $(PWD)/openapi/v1:/openapi/v1 swaggerapi/swagger-ui
