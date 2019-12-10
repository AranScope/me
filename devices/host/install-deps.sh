#!/bin/bash

# Install docker
curl -sSL https://get.docker.com | sh

# Add user to the docker group
usermod -aG docker <your_user>

# Install pip
curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py && sudo python3 get-pip.py

# Install docker compose
sudo pip3 install docker-compose

# Install vim
sudo apt-get update && sudo apt-get install vim -y

# Allow containers to access host servers LAN
sudo sysctl net.ipv4.conf.all.forwarding=1
sudo iptables -P FORWARD ACCEPT
