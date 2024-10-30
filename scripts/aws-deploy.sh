REGION=us-east-1
ACCOUNT_ID=442042544151
ECR_REPO=greasy-balls/go-web-service
DOCKER_IMAGE_NAME=greasy-balls/go-web-service
CLUSTER_NAME=NodeCluster
SERVICE_NAME=GoWebService
BUCKET_NAME=sb-greasy-balls-site

aws s3 rm s3://${BUCKET_NAME}/assets --recursive

aws s3 cp ./site/vite-site/dist/ s3://${BUCKET_NAME}/ --recursive
