# deletes the generated openapi server folder and regenerates it
openapi-generate-server:
	rm -rf backend/openapi && openapi-generator generate -i openapi/v1/ticketing.yaml -g go-server -o backend --package-name openapi --additional-properties sourceFolder=openapi,serverPort=3000