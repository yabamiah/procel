# procel
Is a tiny network CLI made in Go with urfave/cli package, which can quickly look up IPs addresses, hosts, DNS NS and MX. 
Mainly made to train the Go language and the implementation of your packages and project organization.
<p align="center"><img src="Captura de tela_2022-10-07_22-18-30.png" width="700"></p>

---

## Installation

```bash
go get github.com/yabamiah/procel
```

## Usage

#### Get a IP addresses by a host link

```bash
myapp ips --host wikipedia.org
```

#### Get a Host name by IP addresses
```bash
myapp host --ip 5.255.255.50
```

#### Get a DNS NS by a host link
```bash
myapp ns --host wikipedia.org
```
#### and MX
```bash
myapp mx --host wikipedia.org
```

## Quick Starting
```bash
git clone https://github.com/yabamiah/procel
cd procel
./build/myapp
```

## Contributing
I'm still new to Go, and I appreciate any kind of help!
