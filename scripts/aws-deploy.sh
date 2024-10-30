REGION=us-east-1
ACCOUNT_ID=442042544151
ECR_REPO=greasy-balls/go-web-service
DOCKER_IMAGE_NAME=greasy-balls/go-web-service
CLUSTER_NAME=NodeCluster
SERVICE_NAME=GoWebService
BUCKET_NAME=sb-greasy-balls-site

aws s3 rm s3://${ BUCKET_NAME }/assets --recursive

aws s3 cp ./site/vite-site/dist/ s3://${ BUCKET_NAME }/ --recursive

aws ecr get-login-password --region ${REGION} --profile sb-iamadmin-general | docker login --username AWS --password-stdin ${ACCOUNT_ID}.dkr.ecr.${REGION}.amazonaws.com

cd api/web-service

docker build \
--platform=linux/amd64 \
-t ${DOCKER_IMAGE_NAME} .

docker tag ${DOCKER_IMAGE_NAME}:latest ${ACCOUNT_ID}.dkr.ecr.${REGION}.amazonaws.com/${DOCKER_IMAGE_NAME}:latest
docker push ${ACCOUNT_ID}.dkr.ecr.${REGION}.amazonaws.com/${ECR_REPO}:latest

aws ecs update-service \
  --region "${REGION}" \
  --cluster "${CLUSTER_NAME}" \
  --service "${SERVICE_NAME}" \
  --force-new-deployment \
