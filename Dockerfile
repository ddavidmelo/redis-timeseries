FROM node:lts-alpine as build-ui-stage

ENV PROJECT_PATH=/redis-timeseries
ENV MODE=prod

RUN mkdir -p $PROJECT_PATH
WORKDIR $PROJECT_PATH

COPY ./ui/package*.json $PROJECT_PATH/
RUN npm install

COPY ./ui $PROJECT_PATH
RUN npm run build -- --mode $MODE

FROM golang:1.19.1 as build-go-stage

ENV PROJECT_PATH=/redis-timeseries

RUN mkdir -p $PROJECT_PATH
COPY . $PROJECT_PATH
WORKDIR $PROJECT_PATH

RUN go mod download
RUN go build -o redis-timeseries .
RUN chmod +x redis-timeseries


# Final stage
FROM debian:buster-slim as production

ENV PROJECT_PATH=/redis-timeseries

RUN apt-get update && \
		apt-get install -y \
		ca-certificates \
		iputils-ping \
		curl \
		procps \
		&& rm -rf /var/lib/apt/lists/*

RUN mkdir -p $PROJECT_PATH
COPY . $PROJECT_PATH
COPY --from=build-ui-stage $PROJECT_PATH/dist $PROJECT_PATH/ui/dist
COPY --from=build-go-stage $PROJECT_PATH/redis-timeseries $PROJECT_PATH
WORKDIR $PROJECT_PATH

RUN chmod +x redis-timeseries

USER nobody:nogroup

ENTRYPOINT ["./redis-timeseries"]