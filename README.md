# goGo-Dork

<p align="center">
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-_red.svg"></a>
<a href="https://github.com/7imbitz/goGO-Dork/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>
<a href="https://goreportcard.com/badge/github.com/7imbitz/goGO-Dork"><img src="https://goreportcard.com/badge/github.com/7imbitz/goGO-Dork"></a>

</p>

A Simple Google Dork Tools written in golang 

<h1 align="center">
  <img src="https://user-images.githubusercontent.com/26263598/202352303-0a49c358-53b7-4bc3-8932-9246ef8e1b1d.png" height="450">
  <br>
</h1>

### Installation Instructions

goGO-Dork requires **go1.17** to install successfully. Run the following command to get the repo - 

```bash
go install -v github.com/7imbitz/goGO-Dork/cmd/goGO-Dork@latest
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
   -domain string  Domain to scan
   -result int     Number of results per search (default 10)
```

### Current Capability
- Run through 11 dork
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
    
 # Notes

- goGO-Dork is still ongoing development (during my free time) , improvement and updates will be done in near future
- goGO-Dork was inspired by [dorks_hunter](https://github.com/six2dez/dorks_hunter) from [Six2dez](https://twitter.com/Six2dez)
- It is currently considered "automate" as user just need to provide domain and the tool does the work
- Next update will implement modular where user can choose to only dork a/several things rather than selecting all

-----

goGO-Dork is made to ease my work and also to study golang. Community contributions are welcomed. Any issues occur can be [reported](https://github.com/7imbitz/goGO-Dork/issues) 
