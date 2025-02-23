FROM hyperservice-service-image:latest

RUN mkdir -p /apps/service-y/bin
COPY /apps/service-y/bin/main /apps/service-y/bin
COPY /apps/service-y/moon.yml /apps/service-y

RUN mkdir .moon
COPY /.moon/* /.moon

COPY /.prototools /

ENTRYPOINT ["hy-dp"]
