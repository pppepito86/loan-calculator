#!/bin/bash

set -e -x

add-apt-repository ppa:ubuntu-lxc/lxd-stable -y
apt-get update

# git & stuff
apt-get install -y curl git gcc make python-dev vim-nox jq cgroup-lite silversearcher-ag

# nodejs
curl -sL https://deb.nodesource.com/setup_7.x | sudo -E bash -
sudo apt-get install -y nodejs
sudo apt-get install -y build-essential

# angular
npm install -g angular-cli

#clone app
git clone https://github.com/pppepito86/loan_calculator.git /app/judge

#install go
apt-get install golang -y

#set GOPATH
echo "GOPATH=/app/judge" >> /etc/environment
echo "LC_ALL=en_US.UTF-8" >> /etc/environment
echo "LC_CTYPE=en_US.UTF-8" >> /etc/environment
echo "LANG=en_US.UTF-8" >> /etc/environment
echo "LANGUAGE=en_US.UTF-8" >> /etc/environment
source /etc/environment

