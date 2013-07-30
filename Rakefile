#   Copyright 2013 Matt Ho
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#     http:#www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

task :default => :test

DIST="dist"

task :prepare do 
  system "go get launchpad.net/goamz/aws ; go get launchpad.net/goamz/s3"
end

desc "go test"
task :test => :prepare do
  system "go test"
end

desc "build the installer package"
task :build => :prepare do 
  system "go build"
end

desc "clean"
task :clean do
  system "rm -rf dist"
  system "rm -f s3"
  system "rm -f *.deb"
end

desc "create the debian content directory"
task :contents => :build do 
  system "mkdir -p #{DIST}/usr/local/bin"
  system "cp s3 #{DIST}/usr/local/bin"
end

desc "create a debian package"
task :package => :contents do
  system <<EOF
  fpm \
    --force \
    --deb-user 0 \
    --deb-group 0 \
    --url http://github.com/tmtt/s3 \
    --name s3 \
    --version 1 \
    --vendor "The Marketing Tool Tool" \
    -s dir \
    -t deb \
    -C #{DIST} \
    usr
EOF
end

