/*
@Time : 2024/8/4 下午4:37
@Author : ljn
@File : template
@Software: GoLand
*/

package main

var DockerFileTemplate = `FROM ubuntu:20.04
WORKDIR /root/
COPY deploy ./deploy
RUN chmod +x ./deploy/*
{{EXPOSE_PORTS}}
# Update package sources and install required packages
RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list && \
    sed -i 's@//ports.ubuntu.com@//mirrors.ustc.edu.cn@g' /etc/apt/sources.list && \
    apt-get update && \
    apt-get install -y curl unzip wget && \
    apt-get install -y redis-server && \
    sed -i 's/^bind 127.0.0.1 ::1/bind 0.0.0.0/' /etc/redis/redis.conf && \
    sed -i 's/^protected-mode yes/protected-mode no/' /etc/redis/redis.conf && \
    redis-server /etc/redis/redis.conf && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y mysql-server mysql-client && \
    service mysql start && \
    sleep 5 && \
    mysql -u root -e "INSTALL PLUGIN validate_password SONAME 'validate_password.so'; \
    SET GLOBAL validate_password_policy=0; \
    SET GLOBAL validate_password_mixed_case_count=0; \
    SET GLOBAL validate_password_number_count=0; \
    SET GLOBAL validate_password_special_char_count=0; \
    SET GLOBAL validate_password_length=6; \
    SHOW VARIABLES LIKE 'validate_password%'; \
    ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '123456'; \
    UPDATE mysql.user SET Host='%' WHERE User='root' AND Host='localhost'; \
    FLUSH PRIVILEGES;" && \
    mysql -u root -p123456 -e "CREATE DATABASE IF NOT EXISTS {{DB_NAME}} CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;" && \
    mysql -u root -p123456 {{DB_NAME}} < ./deploy/sql/{{DB_NAME}}.sql && \
    sed -i 's/^bind-address\s*=\s*127.0.0.1/bind-address = 0.0.0.0/' /etc/mysql/mysql.conf.d/mysqld.cnf && \
    service mysql restart && \
    apt install -y nginx && \
#    {{INSTALL_JAVA}}

RUN cp -f ./deploy/nginx.conf /etc/nginx/nginx.conf
WORKDIR /root/deploy/back
CMD redis-server /etc/redis/redis.conf && service mysql restart && nginx && {{START_CMD}}
`

var NginxTemplate = `user  root;
worker_processes  1;
error_log  error.log;
error_log  error.log  notice;
error_log  error.log  info;
events {
    worker_connections  1024;
}
http {
    include       mime.types;
    default_type  application/octet-stream;
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';
    access_log  access.log  main;
    sendfile        on;
    keepalive_timeout  65;
    server {
        listen      8081;
        server_name  localhost;
        root	/root/deploy/front/dist;
        location / {
            try_files $uri $uri/ @router;
            index index.html;
        }
        location /api {
            proxy_pass {{PROXY_URL}};
        }
        location @router{
            rewrite ^.*$ /index.html last;
        }
    }
    include servers/*;
}
`
