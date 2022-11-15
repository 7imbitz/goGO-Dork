# goGo-Dork Architecture Document

A brief overview of goGo-Dork architecture

## cmd
Contain main.go and pkg

## main
Contain graceful exit, main function to call banner from [[## banner]] and to return parser into [[## query]]

### cmd/pkg
Contain package args and core

## cmd/pkg/args
Contain args

## arg
Contain argument that need to be parse in main
- Domain to be analyze
- Result (Number of result to be return, default 10)

## cmd/pkg/core
Contain banner and query

## banner
Contain logo, version, author

## query
Contain parser to process (Domain, result)
1. Pass domain, dorker won't run if no domain provided

```go
if domain == "" {
    //dork
} else {
    gologger.Error().Msg("Please provide domain to be analyze")
}
```

2. If domain provided, push into each function of dorking google
    Subdomain           - https://www.google.com/search?q=site:*.(domain)
    Exposed Document    - https://www.google.com/search?q=site:(domain) ext:doc | ext:docx | ext:odt | ext:pdf | ext:rtf | ext:sxw | ext:psw | ext:ppt | ext:pptx | ext:pps | ext:csv)
    Login Page          - https://www.google.com/search?q=site:(domain) inurl:signup | inurl:register | intitle:Signup
    Exposed .git        - https://www.google.com/search?q=intitle:index of /.git/hooks "(domain)"
    Wordpress Files     - https://www.google.com/search?q=site:(domain) inurl:wp-content | inurl:wp-includes
    PHP Error           - https://www.google.com/search?q=site:(domain) "PHP Parse error" | "PHP Warning" | "PHP Error"
    Backup File         - https://www.google.com/search?q=site:(domain) ext:bkf | ext:bkp | ext:bak | ext:old | ext:backup
    Config File         - https://www.google.com/search?q=site:(domain) ext:xml | ext:conf | ext:cnf | ext:reg | ext:inf | ext:rdp | ext:cfg | ext:txt | ext:ora | ext:env | ext:ini
    Database File       - https://www.google.com/search?q=site:(domain) ext:sql | ext:dbf | ext:mdb | ext:db
    Sub-subdomain       - https://www.google.com/search?q=site:*.*.(domain)

<p class="callout warning">This tools is still in development stage, there will be error here and there.</p>


