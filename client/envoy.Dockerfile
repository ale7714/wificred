FROM thegrandpkizzle/envoy:1.26.1
COPY ./envoy.yaml /etc/envoy/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml

sudo iptables -t nat -A PREROUTING -d 0/0 -p tcp –dport 80 -j DNAT –to 192.168.2.2:8081
sudo iptables -t nat -A PREROUTING -d 0/0 -p tcp –dport 443 -j DNAT –to 192.168.2.2:8081

sudo iptables -t nat -A PREROUTING -d 0/0 -p tcp --dport 80 -j DNAT --to-destination 192.168.2.2:8081
sudo iptables -t nat -A PREROUTING -d 0/0 -p tcp --dport 80 -j DNAT --to-destination 192.168.2.2:8081