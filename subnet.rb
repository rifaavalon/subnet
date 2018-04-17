require 'ipaddr'


net1 = IPAddr.new("98.210.237.192/26")
p net1.to_i.to_s(16)
ip2 = IPAddr.new("98.210.237.193")
p ip2.to_i.to_s(16)
ip3 = IPAddr.new("98.210.237.75")
p ip3.to_i.to_s(16)

p net1.include?(ip2)
p net1.include?(ip3)
