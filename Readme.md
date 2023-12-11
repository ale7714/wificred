## Config on the device 

Followed [this tutorial](https://raspberrytips.com/access-point-setup-raspberry-pi/) enabled AP in raspberry PI. 

Add iptables redirect

```
sudo iptables -t nat -A PREROUTING -p tcp --dport 80 -j REDIRECT --to-port <Server PORT>
```

In /etc/dnsmasq.conf 

```sh
interface=wlan0

# Pool of IP addresses served via DHCP
dhcp-range=192.168.4.2,192.168.4.255,255.255.255.0,24h

# Redirect all domains (the #) to the address http://192.168.2.2/ (Pi)
address=/#/http://192.168.2.2/

dhcp-option-force=114,http://192.168.2.2/captive
```