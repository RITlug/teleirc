FROM node:10-alpine

COPY . /opt/teleirc
WORKDIR /opt/teleirc

RUN apk add --no-cache --update bash build-base python \
    && addgroup -g 65532 teleirc \
    && adduser -s /bin/bash -h /opt/teleirc -D -H teleirc -u 65532 -G teleirc \
    && yarn \
    && chown -R teleirc:teleirc /opt/teleirc


USER teleirc
CMD ["node", "teleirc.js"]
