FROM golang:1.24.0-bookworm

RUN echo "Install base dependencies" && \
    apt-get update && apt-get install -y \
    curl nano lsof openssh-client collectd git


# Install Air
RUN go install github.com/air-verse/air@latest

# Install hyperservice Dataplane
RUN curl -sL https://raw.githubusercontent.com/tiagovinicius/hyperservice/main/get-hy-dp.sh | bash && \
    chmod +x hy-dp && mv hy-dp /usr/local/bin/

# Install Kuma
RUN curl -L https://kuma.io/installer.sh | VERSION=2.9.3 sh - && \
    mv kuma-2.9.3/bin/* /usr/local/bin/ && \
    rm -rf kuma-2.9.3

# Install Moonrepo CLI
RUN curl -fsSL https://moonrepo.dev/install/moon.sh | bash && \
    mv "$HOME/.moon/bin/moon" /usr/local/bin/

RUN mkdir -p /apps/service-y/bin
COPY ./apps/service-y/bin/main /apps/service-y/bin
COPY ./apps/service-y/moon.yml /apps/service-y

RUN mkdir /.moon
COPY ./.moon/* /.moon

COPY ./.prototools /

ENTRYPOINT ["hy-dp"]
