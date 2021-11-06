#!/bin/bash
yum update -y
yum install httpd -y
systemctl start httpd
systemctl enable httpd
cd /var/www/html
echo "This is a test page running on Apache on EC2 in the AWS Cloud" > index.html

# Test with this command:
curl http://169.254.169.254/latest/user-data