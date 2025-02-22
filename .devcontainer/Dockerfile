FROM golang:1.24.0-bookworm

RUN echo "Install base dependencies" && \
    apt-get update && apt-get install -y \
    curl nano lsof openssh-client collectd

# Install Air
RUN go install github.com/air-verse/air@latest

# Install Moonrepo CLI
RUN curl -fsSL https://moonrepo.dev/install/moon.sh | bash && \
    mv "$HOME/.moon/bin/moon" /usr/local/bin/

ENTRYPOINT ["bash", "true"]
 