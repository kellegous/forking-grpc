FROM ubuntu:xenial
RUN apt-get update \
    && apt-get install -y php-dev php-pear phpunit vim unzip build-essential python3 git gdb \
    && pecl install protobuf \
    && curl -OL https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip \	
  	&& unzip protoc-3.6.1-linux-x86_64.zip -d /usr/local \	
  	&& rm protoc-3.6.1-linux-x86_64.zip \
	&& curl https://dl.google.com/go/go1.11.1.linux-amd64.tar.gz | tar xz -C /usr/local
COPY grpc.ini /etc/php/7.0/cli/conf.d/30-grpc.ini
ENV PATH="/usr/local/go/bin:${PATH}"
WORKDIR /app