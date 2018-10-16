FROM ubuntu:18.04

RUN apt-get update && \
    apt-get install -y \
    curl \
    ipset \
    autoconf \
    automake \
    make \
    iproute2 \
    lsof \
    gcc \
    kmod \
    iputils-ping \
    inetutils-traceroute \
    unzip

ENV IPRANGE_VERSION 1.0.3

RUN curl -L https://github.com/firehol/iprange/releases/download/v$IPRANGE_VERSION/iprange-$IPRANGE_VERSION.tar.gz | tar zvx -C /tmp && \
    cd /tmp/iprange-$IPRANGE_VERSION && \
    ./configure --prefix= --disable-man && \
    make && \
    make install && \
    cd && \
    rm -rf /tmp/iprange-$IPRANGE_VERSION

ENV FIREHOL_VERSION 3.1.3

RUN curl -L https://github.com/firehol/firehol/releases/download/v$FIREHOL_VERSION/firehol-$FIREHOL_VERSION.tar.gz | tar zvx -C /tmp && \
    cd /tmp/firehol-$FIREHOL_VERSION && \
    ./autogen.sh && \
    ./configure --prefix= --disable-doc --disable-man && \
    make && \
    make install && \
    cp contrib/ipset-apply.sh /bin/ipset-apply && \
    cd && \
    rm -rf /tmp/firehol-$FIREHOL_VERSION
