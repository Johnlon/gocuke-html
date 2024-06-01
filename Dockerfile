FROM ubuntu:20.04
SHELL ["/bin/bash", "-c"]
LABEL maintainer="Rodrigo Odhin"

ARG DEBIAN_FRONTEND=noninteractive

# Install libs
RUN apt-get update --fix-missing
RUN apt-get install -f -y openssh-server curl git sudo
RUN apt-get install -y apache2 apache2-utils php php-zip
RUN apt-get install -y nodejs npm
RUN a2enmod rewrite

## Install golang
RUN wget https://go.dev/dl/go1.20.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz
RUN rm -rf go1.20.linux-amd64.tar.gz
RUN apt-get clean

# Change PHP config
ENV PHP_UPLOAD_MAX_FILESIZE 50000000
ENV PHP_POST_MAX_SIZE 1000000000
ENV PHP_MEM_LIMIT 1000000000
RUN echo "post_max_size = 1024M" >> /etc/php/7.4/apache2/php.ini
RUN echo "upload_max_filesize = 50M" >> /etc/php/7.4/apache2/php.ini
RUN echo "max_file_uploads = 999999" >> /etc/php/7.4/apache2/php.ini

# Change Apache config
RUN echo "ServerName localhost" >> /etc/apache2/apache2.conf

# Install Docsify
RUN npm i docsify-cli -g

# Copy Website
RUN cd /var/www/html && rm -f index*
COPY website/html /var/www/html
RUN chown -R www-data:www-data /var/www/html

# Copy Docs
COPY docs /docs

# Set GO path
RUN echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc

# Copy API public files
COPY api/public /api/public

# Copy GOCURE binaries
COPY bin/gocure /
COPY bin/gocureAPI /

# Expose ports
EXPOSE 80
EXPOSE 7087
EXPOSE 3000

# Copy scripts
RUN mkdir /scripts
COPY scripts/entrypoint.sh /scripts/entrypoint.sh
RUN ["chmod", "777", "/scripts/entrypoint.sh"]
ENTRYPOINT ["/scripts/entrypoint.sh"]
