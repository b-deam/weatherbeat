FROM raspbian/stretch
RUN mkdir -p /usr/local/bin/weatherbeat /etc/weatherbeat \
    && apt-get update
COPY weatherbeat /usr/local/bin/weatherbeat
COPY weatherbeat.yml /etc/weatherbeat

CMD ./usr/local/bin/weatherbeat/weatherbeat -e -c /etc/weatherbeat/weatherbeat.yml
