FROM alpine
ADD html /html
ADD group_api-web /group_api-web
WORKDIR /
ENTRYPOINT [ "/group_api-web" ]
