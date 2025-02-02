FROM tsl0922/ttyd:alpine AS ttyd
FROM alpine:latest AS fontcollector

# Install Fonts
RUN apk add --no-cache \
    --repository=http://dl-cdn.alpinelinux.org/alpine/edge/main \
    --repository=http://dl-cdn.alpinelinux.org/alpine/edge/community \
    --repository=http://dl-cdn.alpinelinux.org/alpine/edge/testing \
    font-adobe-source-code-pro font-source-code-pro-nerd \
    font-dejavu font-dejavu-sans-mono-nerd \
    font-fira-code font-fira-code-nerd \
    font-hack font-hack-nerd \
    font-ibm-plex-mono-nerd \
    font-inconsolata font-inconsolata-nerd \
    font-jetbrains-mono font-jetbrains-mono-nerd \
    font-liberation font-liberation-mono-nerd \
    font-noto \
    font-roboto-mono \
    font-ubuntu font-ubuntu-mono-nerd \
    font-noto-emoji

FROM golang:1.23.4

RUN apt-get update

# Add fonts
COPY --from=fontcollector /usr/share/fonts/ /usr/share/fonts

# Install latest ttyd
COPY --from=ttyd /usr/bin/ttyd /usr/bin/ttyd

# Expose port
EXPOSE 1976

# Install Dependencies
RUN apt-get -y install ffmpeg chromium bash

# Install VHS
RUN apt-get install -y wget
RUN wget https://github.com/charmbracelet/vhs/releases/download/v0.9.0/vhs_0.9.0_arm64.deb
RUN dpkg -i vhs_0.9.0_arm64.deb

# Mimic alpine default color option
RUN echo 'alias ls="ls --color"' >> ~/.bashrc

# Install
# COPY vhs /usr/bin/

ENV VHS_PORT="1976"
ENV VHS_HOST="0.0.0.0"
ENV VHS_GID="1976"
ENV VHS_UID="1976"
ENV VHS_KEY_PATH="/vhs/vhs"
ENV VHS_AUTHORIZED_KEYS_PATH=""
ENV VHS_NO_SANDBOX="true"

# ENTRYPOINT ["bash"]
CMD ["bash"]
