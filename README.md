# goGo-Dork

<p align="center">
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-_red.svg"></a>
<a href="https://github.com/7imbitz/goGO-Dork/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>
<a href="https://goreportcard.com/badge/github.com/7imbitz/goGO-Dork"><img src="https://goreportcard.com/badge/github.com/7imbitz/goGO-Dork"></a>

</p>

A Simple Google Dork Tools written in golang 

<h1 align="center">

  <img alt="goGO-Dork" src="https://user-images.githubusercontent.com/26263598/207503339-b35fd676-d343-40ef-af12-e6eb0c68bf0d.png" height="650">
  <br>
</h1>

### Installation Instructions
_for the meantime, this command does not install v`1.1.0` hence please use the [command](https://github.com/7imbitz/goGO-Dork#workaround-installation) below_

goGO-Dork requires **go1.17** to install successfully. Run the following command to get the repo - 

```bash
go install -v github.com/7imbitz/goGO-Dork/cmd/goGO-Dork@latest
```

### Workaround Installation

```bash
git clone https://github.com/7imbitz/goGO-Dork.git
cd goGO-Dork
go build -o goGO-Dork cmd/goGO-Dork/main.go
./goGO-Dork -h
```
_You can move the binary in the `$GOPATH` directory_
```bash
mv goGO-Dork $GOPATH/bin
```

### Flags
```bash
goGO-Dork -h
```

This will display the flags/argument needed for the tool

```bash
Simple Google Dork Search

Usage:
  goGO-Dork [flags]

Flags:
   -d, -domain string  Domain to scan
   -r, -result int     Number of results per search (default 10)
```

### Current Capability
- Run through 14 dork
    - Subdomain
    - Sub-subdomain
    - Login pages
    - Exposed document
    - .git folder
    - Wordpress file
    - PHP error
    - SQL error
    - Config file
    - Database file
    - Directory listing
    - Github info
    - Laravel Debug mode
    
 # Notes

- goGO-Dork is still ongoing development (during my free time) , improvement and updates will be done in near future
- goGO-Dork was inspired by [dorks_hunter](https://github.com/six2dez/dorks_hunter) from [Six2dez](https://twitter.com/Six2dez1)
- It is currently considered "automate" as user just need to provide domain and the tool does the work
- Next update will implement modular where user can choose to only dork a/several things rather than selecting all

-----

goGO-Dork is made to ease my work and also to study golang. Community contributions are welcomed. Any issues occur can be [reported](https://github.com/7imbitz/goGO-Dork/issues) 
