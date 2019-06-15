# iron/go is the alpine image with only ca-certificates added
FROM iron/go
WORKDIR /app

# add the binary
ADD miniflux-linux-amd64 /app/
# add Google Service Account credential file
ADD google-service-account.json /app/

# set Env parameters
# Using ARG will allow us to overrides the value during build or runtime.
ARG DATABASE_URL=postgres://user:password@dbhost/miniflux?sslmode=disable
ENV DATABASE_URL="${DATABASE_URL}"

ARG LISTEN_ADDR=0.0.0.0:8080
ENV LISTEN_ADDR="${LISTEN_ADDR}"

# polling freq. (in minute)
ARG POLLING_FREQUENCY=60
ENV POLLING_FREQUENCY="${POLLING_FREQUENCY}"

ARG GCP_PROJECT_ID=gatrabali
ENV GCP_PROJECT_ID="${GCP_PROJECT_ID}"

ARG GCP_PUBSUB_TOPIC=SyncData
ENV GCP_PUBSUB_TOPIC="${GCP_PUBSUB_TOPIC}"

ENV GOOGLE_APPLICATION_CREDENTIALS=/app/google-service-account.json

ENTRYPOINT ["./miniflux-linux-amd64"]