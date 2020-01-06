FROM ubuntu:trusty
EXPOSE 80
EXPOSE 443
RUN apt-get update
RUN apt-get install -y curl
RUN apt-get install -y git
RUN curl https://getcaddy.com/ | bash -s personal http.prometheus
WORKDIR /var/www
CMD ["caddy"]
