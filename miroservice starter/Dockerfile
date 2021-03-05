FROM ubuntu:hirsute-20201119

RUN adduser --disabled-password --home /home/knoxservice --shell /bin/sh knoxservice

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update

RUN apt-get install -y apt-transport-https ca-certificates

WORKDIR /home/knoxservice

ENV SERVICE_HOME=/home/knoxservice

RUN mkdir bin scripts conf docs

COPY conf/app.yaml conf/app.yaml

COPY bin/server bin/server

COPY scripts/entrypoint.sh scripts/entrypoint.sh

ENTRYPOINT ["sh","scripts/entrypoint.sh"]
