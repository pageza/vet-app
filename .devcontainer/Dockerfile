# Use the official VS Code dev container image for Go
FROM mcr.microsoft.com/vscode/devcontainers/go:0-1.16

# [Optional] Install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends postgresql-client redis


WORKDIR /workspace
