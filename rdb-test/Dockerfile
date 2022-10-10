FROM postgres:13.4

COPY initdb.sh /docker-entrypoint-initdb.d

RUN chmod +x /docker-entrypoint-initdb.d/initdb.sh
