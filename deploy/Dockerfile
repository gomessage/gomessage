#############################################################################
#
#   docker run -d -p 7077:7077 gomessage/gomessage
#
#############################################################################

FROM centos:7.9.2009

RUN mkdir -p /opt/gomessage && rm -rf /opt/gomessage/*

WORKDIR /opt/gomessage

ADD ./gomessage.tar.gz /opt/gomessage/

EXPOSE 7077

ENTRYPOINT ["/opt/gomessage/gomessage"]
