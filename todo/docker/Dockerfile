FROM ubuntu:focal

# Services ports
EXPOSE 55000/tcp 

RUN mkdir -p /opt/

COPY ./bin /opt/bin

ENTRYPOINT [ "/opt/bin/todo" ]
