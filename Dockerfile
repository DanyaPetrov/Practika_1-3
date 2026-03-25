FROM ubuntu:latest

RUN apt-get update && apt-get install -y dpkg

COPY myprogram.deb /tmp/

RUN dpkg -i /tmp/myprogram.deb

RUN apt-get install -f -y

ENTRYPOINT ["/usr/bin/lab"]
