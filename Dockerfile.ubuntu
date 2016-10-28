FROM ubuntu:16.04

COPY . /opt/teleirc/

RUN apt-get update && \
	apt-get install -y curl && \
	curl -sL https://deb.nodesource.com/setup_6.x \
	| bash - && \
	apt-get install -y \
		build-essential \
		git \
		nodejs && \
	cd /opt/teleirc && \
	npm install

COPY config.js.docker /opt/teleirc/config.js

WORKDIR /opt/teleirc

CMD ["node", "teleirc.js"]
