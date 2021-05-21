FROM google/cloud-sdk:latest

RUN curl -LO https://dl.google.com/go/go1.13.15.linux-amd64.tar.gz
RUN tar zxvf go1.13.15.linux-amd64.tar.gz -C /usr/local/
ENV PATH $PATH:/usr/local/go/bin

RUN mkdir -p /src

WORKDIR /src
