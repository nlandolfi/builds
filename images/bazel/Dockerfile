FROM ubuntu:xenial

# Bucket used to store already built binaries
# ARG TOOLS_BUCKET

# Installing necessary packages
RUN rm -rf /var/lib/apt/lists/* \
    && apt-get update --fix-missing -qq \
    && apt-get install -qqy git iptables procps sudo xz-utils \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Adding sudo group user no password access.
# This is used by bootstrap user to start docker service
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

# Installing
ADD scripts /tmp/istio_tmp/scripts
RUN chmod +x /tmp/istio_tmp/scripts/linux-install-software
#RUN /tmp/istio_tmp/scripts/linux-install-software \
#      -b "${TOOLS_BUCKET}" \
#    && rm -rf /tmp/istio_tmp
RUN /tmp/istio_tmp/scripts/linux-install-software \
    && rm -rf /tmp/istio_tmp

ENV PATH /usr/local/go/bin:/usr/lib/google-cloud-sdk/bin:${PATH}

ENV HOME /home/bootstrap
RUN useradd -c "Bootstrap user" -d ${HOME} -G docker,sudo -m bootstrap -s /bin/bash
ENV TEST_TMPDIR /home/bootstrap/.cache/bazel

#ADD github_helper /usr/local/bin/github_helper
ADD entrypoint /usr/local/bin/entrypoint
RUN chmod +rx /usr/local/bin/entrypoint

USER bootstrap
WORKDIR ${HOME}

RUN mkdir -p /home/bootstrap/go/src
ENV GOPATH=/home/bootstrap/go
ENV PATH ${GOPATH}/bin:${PATH}

ENTRYPOINT ["entrypoint"]
