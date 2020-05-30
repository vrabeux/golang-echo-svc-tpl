VERSION=0.0.1

SERVICE_DOCKER_IMAGE_NAME=svc-replace-mes:latest
GITHUB_DOCKER_REPO_ROOT=docker.pkg.github.com/bellesdespins/svc-replace-mes

docker: $(SOURCES)
	@docker build -t ${SERVICE_DOCKER_IMAGE_NAME} -f ./Dockerfile .

docker-push:
	docker tag ${SERVICE_DOCKER_IMAGE_NAME} ${GITHUB_DOCKER_REPO_ROOT}/${SERVICE_DOCKER_IMAGE_NAME}
	docker push ${GITHUB_DOCKER_REPO_ROOT}/${SERVICE_DOCKER_IMAGE_NAME}

