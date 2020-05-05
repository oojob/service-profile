#!/bin/bash
# inspiration: https://github.com/ShelterTechSF/askdarcel-web/blob/master/tools/docker-build.sh
set -ex

SERVICE_NAME=company
REPO="docker.pkg.github.com/oojob/service-$SERVICE_NAME/$SERVICE_NAME"

COMMIT=$CODEBUILD_RESOLVED_SOURCE_VERSION
if [[ -z "$COMMIT" ]]; then
  COMMIT=$(git log -1 --format=%H)
fi
COMMIT=${COMMIT::8}

DOCKER_HOST="docker.pkg.github.com"

if [[ "$ACCOUNT_SVC_PROFILE" == "development" ]]; then
  TAG="dev"
elif [[ "$ACCOUNT_SVC_PROFILE" == "testing" ]]; then
  TAG="test"
elif [[ "$ACCOUNT_SVC_PROFILE" == "production" ]]; then
  TAG="prod"
elif [[ "$ACCOUNT_SVC_PROFILE" == "staging" ]]; then
  TAG="stag"
fi

echo $TAG
echo $COMMIT

echo "Creating version.json..."
echo "{
  \"commit\": \"$COMMIT\",
  \"image\": \"$TAG\"
}" > ./scripts/version.json

echo "Building docker image..."
docker build -f scripts/docker/Dockerfile -t $REPO:$COMMIT .
docker tag $REPO:$COMMIT $REPO:$TAG

echo "Pushing docker image..."
docker push $REPO:$TAG

echo "Writing image definitions file..."
echo "[
  {
    \"name\":\"AccountService\",
    \"imageUri\":\"$REPO:$TAG\"
  }
]" > ./scripts/imagedefinitions.json

echo "generating linkerd proxies"
cat scripts/kubernetes/$SERVICE_NAME-service.yml | linkerd inject  - > scripts/kubernetes/$SERVICE_NAME-linkerd.yml
