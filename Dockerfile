FROM ubuntu:latest
MAINTAINER Jason Waldrip <jason@waldrip.net>

# Install and setup go
RUN DEBIAN_FRONTEND=noninteractive apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y golang
ENV GOPATH /usr/local
ENV PATH $GOPATH/bin:$PATH

# Compile docker-ruby
ADD . /usr/local/src/github.com/passfail/docker-ruby
RUN cd /usr/local/src/github.com/passfail/docker-ruby && go install

# Cleanup
RUN apt-get purge -y golang
RUN rm -rf $GOPATH/src

# ENV
ENV RUBIES_DIR /rubies

# TEST
RUN docker-ruby version
