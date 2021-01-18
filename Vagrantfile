Vagrant.configure("2") do |config|
  config.vm.box = "debian/buster64"
  config.vm.synced_folder "./", "/home/vagrant/nolabase"

  config.vm.provider "virtualbox" do |v|
    v.name = "nolabase"
    v.memory = 1024
    v.cpus = 2
  end
end


