FROM hyperservice-service-image:latest

RUN mkdir -p /apps/service-x/bin
COPY /apps/service-x/bin/main /apps/service-x/bin
COPY /apps/service-x/moon.yml /apps/service-x

RUN mkdir .moon
COPY /.moon/* /.moon

COPY /.prototools /

ENTRYPOINT ["hy-dp"]
