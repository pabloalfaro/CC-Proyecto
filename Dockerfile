FROM golang:alpine

RUN apk update \
	&& apk add --no-cache git \
	&& apk add --no-cache make build-base

# Creación del usuario
ENV USER=appuser
ENV UID=12345

# Creo un nuevo usuario siguiendo estos pasos: https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \
    --disabled-password \
    --gecos "" \
    --shell "/sbin/nologin" \
    --uid "${UID}" \
    "$USER"
 
RUN mkdir -p app/test

USER "$USER"

WORKDIR /app/test

CMD ["make", "test"]
