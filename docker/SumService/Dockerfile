FROM ubuntu:utopic

MAINTAINER Gianluca Costa <gianluca@gianlucacosta.info>

ENTRYPOINT ["/usr/sbin/SumService"]

CMD ["80"]

EXPOSE 80

ADD https://github.com/giancosta86/MiniSum/releases/download/v1.0/SumService /usr/sbin/SumService

RUN ["chmod", "500", "/usr/sbin/SumService"]
