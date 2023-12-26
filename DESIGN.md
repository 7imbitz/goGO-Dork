# goGo-Dork Architecture Document

A brief overview of goGo-Dork architecture

## cmd
Contain goGO-Dork.go and pkg

## goGO-Dork
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
- exposed git 		: intext:index of /.git parent directory (domain)
- sql error 			: inurl:".php?id=" "You have an error in your SQL syntax" (domain)
- php error 			: site:(domain) "PHP Parse error" | "PHP Warning" | "PHP Error""
- backup file 		: site:(domain) ext:bkf | ext:bkp | ext:bak | ext:old | ext:backup
- database file 		: site:(domain) ext:sql | ext:dbf | ext:mdb | ext:db"
- config file 		: site:(domain) ext:xml | ext:conf | ext:cnf | ext:reg | ext:inf | ext:rdp | ext:cfg | ext:txt | ext:ora | ext:env | ext:ini
- laravel debug mode 	: site:(domain) intitle:"Whoops! There was an error" intext:"Environment Variables"
- directory listing 	: site:(domain) intitle:index.of
- login pages 		: site:(domain) inurl:signup | inurl:register | intitle:Signup
- wordpress 			: site:(domain) inurl:wp-content | inurl:wp-includes
- exposed document 	: site:(domain) ext:doc| ext:docx | ext:odt | ext:pdf | ext:rtf | ext:sxw | ext:psw | ext:ppt | ext:pptx | ext:pps | ext:csv)
- subdomain 			: site:*.(domain)
- sub-subdomain 		: site:*.*.(domain)
- github information 	: site:github.com | site:gitlab.com (domain)

> [!WARNING]
> This tool is still in development stage, there will be error here and there.


