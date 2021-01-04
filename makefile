PROJECT_ID=YOUR_GCP_PROJECT_NAME
CONTAINER_NAME=echo_scelton
VERSION=v0
REGION=asia-northeast1


CONTAINER_FULLNAME=${CONTAINER_NAME}:${VERSION}
GCR_CONTAINER_PATH=asia.gcr.io/${PROJECT_ID}/${CONTAINER_NAME}

testrun:
	PORT=8080 go run ./cmd/app

run:
	docker run -e PORT=8080 -p 8080:8080 ${CONTAINER_FULLNAME}

build:
	docker build -t ${CONTAINER_FULLNAME} ./

deploy:
	make build
	docker tag ${CONTAINER_FULLNAME} ${GCR_CONTAINER_PATH}
	docker push ${GCR_CONTAINER_PATH}
	gcloud run deploy ${CONTAINER_NAME} --image=${GCR_CONTAINER_PATH} --max-instances=1 --memory=64M --platform=managed --timeout=5s --allow-unauthenticated --region=${REGION} --project=${PROJECT_ID}

