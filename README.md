oscer
====
simple command-line OSC sending tool

##FEATURES
- send OSC packet via UDP
- support int and float parameters

##Usage
```
 oscer host port /osc/address [args ...]
```

##Example
```
 oscer localhost 10000 /hello
 oscer 192.168.1.100 11000 /1/push1 1
 oscer 192.168.1.101 12000 /accxyz 0.5 0.2 1.0
```

##WEBSITE
http://github.com/aike/oscer

##CREDIT
oscer program is licenced under MIT License.  
Contact: twitter @aike1000
