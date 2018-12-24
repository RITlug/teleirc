FROM fedora:29

# Cache layers which will not change
RUN groupadd teleirc -g 65532 \
    && useradd -u 65532 -g teleirc -s /bin/bash -M -d /opt/teleirc teleirc

COPY . /opt/teleirc/
WORKDIR /opt/teleirc

RUN dnf -y upgrade --setopt=deltarpm=false \
    && dnf -y install nodejs nodejs-yarn libicu-devel python gcc-c++ make --setopt=deltarpm=false \
    && nodejs-yarn \
    && dnf -y remove libicu-devel gcc-c++ \
    && dnf clean all \
    && chown -R teleirc:teleirc /opt/teleirc

CMD ["node", "teleirc.js"]
