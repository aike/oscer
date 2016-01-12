oscer
====
simple command-line OSC sending tool

##FEATURES
- send OSC packet over UDP
- support int32 and float32 parameters
- support IPv4 and IPv6 protocol
- cross compiled executable binaries are available (MacOSX, Windows, Linux, RaspberryPi)

##Usage
```
 oscer host port /osc/address [args ...]
```

##Example
```
 oscer localhost 10000 /hello
 oscer fe80::1%lo0 11000 /world
 oscer 192.168.1.100 12000 /1/push1 1
 oscer 192.168.1.101 13000 /accxyz 0.5 0.2 1.0
```

##WEBSITE
http://github.com/aike/oscer

##CREDIT
oscer program is licenced under MIT License.  
Contact: twitter @aike1000
