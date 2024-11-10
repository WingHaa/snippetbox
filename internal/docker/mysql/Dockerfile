FROM mysql:9.1

ENV MYSQL_ALLOW_EMPTY_PASSWORD=true
ENV MYSQL_USER=docker
ENV MYSQL_PASSWORD=docker
ENV MYSQL_DATABASE=snippetbox

COPY ./config/my.cnf /etc/mysql/conf.d/my.cnf
