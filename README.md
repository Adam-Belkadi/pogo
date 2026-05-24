# Pogo 

Pogo is a blazingly fast port scanner written in Go, built for speed, simplicity, and clean terminal output.

## Current Status 😀

Right now, the only working feature is scanning the first 1000 ports of any domain.
With very strong concurency.

## Only Working Command

```bash
pogo.exe scan port --url scanme.nmap.org
```

## Example Output

```text
❯ pogo.exe scan port --url scanme.nmap.org

[#] Scanning Target: 2600:3c01::f03c:91ff:fe18:bb2f --- scanme.nmap.org

[+] Port 22 is open: SSH-2.0-OpenSSH_6.6.1p1 Ubuntu-2ubuntu2.13
[+] Port 80 is open: Unable to detect version

[#] Scan Complete: 2600:3c01::f03c:91ff:fe18:bb2f --- scanme.nmap.org
```

## Planned Features

- Full port range scanning
- Improved service/version detection
- UDP scanning
- IPv6 support
- JSON output support
- Configurable timeouts

## Installation

```bash
git clone https://github.com/Adam-Belkadi/pogo.git
cd pogo
go build
```

## Usage

```bash
pogo.exe scan port --url <target>
```

## Disclaimer

Pogo is currently experimental and under active development.
