require 'ipaddr'


net1 = IPAddr.new("98.210.237.192/26")
p net1.to_i.to_s(16)
net2 = IPAddr.new("98.210.237.193")
p net2.to_i.to_s(16)
net3 = IPAddr.new("98.210.237.75")
p net3.to_i.to_s(16)

p net1.include?(net2)
p net1.include?(net3)
