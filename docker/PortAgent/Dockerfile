FROM ubuntu:utopic

MAINTAINER Gianluca Costa <gianluca@gianlucacosta.info>

ENTRYPOINT ["/usr/sbin/PortAgent"]

CMD ["7070"]

ADD https://github.com/giancosta86/MiniSum/releases/download/v1.0/PortAgent /usr/sbin/PortAgent

RUN ["chmod", "500", "/usr/sbin/PortAgent"]
