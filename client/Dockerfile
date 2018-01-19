FROM alpine:3.6

# 设置locale
ENV LANG en_US.UTF-8
ENV LANGUAGE en_US:en
ENV LC_ALL en_US.UTF-8
ENV TZ=Asia/Shanghai

RUN mkdir /app_home

WORKDIR /app_home

COPY client /app_home

RUN chmod +x /app_home/client

ENV CLIENT=/app_home/client
ENV PATH $CLIENT:$PATH