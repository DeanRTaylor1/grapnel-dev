run-docker:
	docker build -t sysd .; docker run -p 8080:8080 sysd