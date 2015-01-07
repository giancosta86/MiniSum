FROM ubuntu:utopic

MAINTAINER Gianluca Costa <gianluca@gianlucacosta.info>

ENTRYPOINT ["/usr/sbin/SumClient"]

CMD ["portagent", "7070"]

ADD https://github.com/giancosta86/MiniSum/releases/download/v1.0/SumClient /usr/sbin/SumClient

RUN ["chmod", "500", "/usr/sbin/SumClient"]
