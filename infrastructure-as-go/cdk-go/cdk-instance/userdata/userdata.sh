#!/bin/bash 
set -xe
yum update -y
yum install -y httpd
echo "Hello World" > /var/www/html/index.html
# Turn on web server 
chkconfig httpd on 
service httpd start


