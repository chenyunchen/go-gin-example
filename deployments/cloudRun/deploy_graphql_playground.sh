. ./deployments/cloudRun/env/base.env

SERVICE_NAME=mockserver

gcloud run deploy "${SERVICE_NAME}" \
      --image "${IMAGE}" \
      --allow-unauthenticated \
      --cpu 1000m \
      --memory 128Mi \
      --max-instances 1 \
      --platform managed \
      --region asia-east1 \
      --project "${PROJECT_NAME}" \
      --set-env-vars ENV="${ENV}"

