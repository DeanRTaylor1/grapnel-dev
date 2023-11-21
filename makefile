run-docker:
	docker build -t sysd .; docker run -p 8080:8080 sysd
build-css:
	./tailwindcss -i ./handlers/templates/input.css -o ./handlers/styles/output.css
test:
	go test -cover ./...
