FROM ubuntu:18.04

WORKDIR /
ENV GIN_MODE=release \
    USE_GIN_LOGGER=0 \
    ZERO_LOG_LEVEL=info \
    MYSQL_CLOUD_HOST=mysql.nas.tenbay \
    ENABLE_SHOWSQL=0
COPY ./main /usr/sbin/homepeer
CMD ["homepeer"]
