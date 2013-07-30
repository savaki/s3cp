#   Copyright 2013 Matt Ho
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

Vagrant.configure("2") do |config|
  config.vm.box       = "precise64"
  config.vm.box_url   = "http://files.vagrantup.com/precise64.box"

  config.vm.network :public_network

  config.vm.provision :shell, :path   => "scripts/install_go.sh"
  config.vm.provision :shell, :inline => "sudo apt-get install -y git bzr"
  config.vm.provision :shell, :inline => "sudo apt-get install -y build-essential"
  config.vm.provision :shell, :inline => "sudo apt-get install -y ruby1.9.3"
  config.vm.provision :shell, :inline => "sudo gem install fpm rake"
  config.vm.provision :shell, :inline => "cd /vagrant ; rake package"

  config.vm.provider :virtualbox do |vb|
    # enable the gui.  the default behavior is to start up headless
    vb.gui = true
  end
end

