#!/bin/bash

# copy rsa key from host to /root/.ssh
cp /secret/github_rsa /root/.ssh

# rename rsa to avoid ssh agent
mv /root/.ssh/github_rsa /root/.ssh/id_rsa

# chenge file permissions
chmod 600 /root/.ssh/id_rsa