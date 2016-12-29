Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/trusty64"
  config.vm.hostname = "walker-vm"
  config.vm.network "private_network", ip: "10.0.0.10"

  config.vm.provider :virtualbox do |vb|
    vb.customize [
      "modifyvm", :id,
      "--cpuexecutioncap", "25",
      "--memory", "256",
    ]
  end
end
