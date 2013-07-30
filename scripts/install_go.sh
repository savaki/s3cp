#!/bin/bash

cat <<EOF | sudo tee -a /etc/apt/sources.list
deb http://archive.mattho.com/ precise universe
EOF

sudo apt-get update
sudo apt-get install -y --force-yes go
