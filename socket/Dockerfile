FROM golang:1.19-alpine

ENV PROJECT_DIR=/socket \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPRIVATE=github.com/jjh930301 

ARG GITACCOUNT
ARG GITTOKEN

WORKDIR /socket
RUN mkdir "/build"
COPY . .
# install git
RUN apk add --no-cache git 
# private repo config
RUN git config --global url."https://${GITACCOUNT}:${GITTOKEN}@github.com/jjh930301/needsss_common".insteadOf "https://github.com/jjh930301/needsss_common"
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon@latest
ENTRYPOINT CompileDaemon -build="go build -o /build/socket" -command="/build/socket"