package core

import (
	"cmd/cmd/pkg/args"
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/logrusorgru/aurora/v4"
	"github.com/projectdiscovery/gologger"
	googlesearch "github.com/rocketlaunchr/google-search"
)

var (
	ctx  = context.Background()
	masa = time.Now()
)

func ParseOptions(options *args.Options) {
	//Check domain argument (must be compulsory)
	if options.Domain != "" {

		//Color domain
		cDomain := aurora.Green(options.Domain)
		gologger.Info().Msgf("Analyzing domain %q\n", cDomain)
		fmt.Println("===========================================================")

		/*//subdomain
		findSubdomain(options)
		fmt.Println("===========================================================")
		time.Sleep(5 * time.Second)

		//Sub-subdomain
		findSubSubdomain(options)
		fmt.Println("===========================================================")
		time.Sleep(5 * time.Second)

		//login pages
		loginPages(options)
		fmt.Println("===========================================================")
		time.Sleep(5 * time.Second)

		//exposed document
		exposedDoc(options)
		fmt.Println("===========================================================")
		time.Sleep(5 * time.Second)

		//.git folder
		gitFolder(options)
		fmt.Println("===========================================================")
		time.Sleep(5 * time.Second)

		//wordpress file
		findWord(options)
		fmt.Println("===========================================================")
		time.Sleep(5 * time.Second)*/

		//php error
		//TODO: fix error google block
		findPHP(options)
		fmt.Println("===========================================================")
		time.Sleep(5 * time.Second)

		//sql error
		findSQL(options)
		fmt.Println("===========================================================")
		time.Sleep(5 * time.Second)

		//backup file
		findBak(options)
		fmt.Println("===========================================================")
		time.Sleep(5 * time.Second)

		//config file
		findConfig(options)
		fmt.Println("===========================================================")
		time.Sleep(5 * time.Second)

		//database file
		findDB(options)
		fmt.Println("===========================================================")
		time.Sleep(5 * time.Second)
		timeTaken()
	} else {
		gologger.Error().Msg("Please provide domain to be analyze")
	}
	//timeTaken()
}

//finding subdomain
func findSubdomain(options *args.Options) {
	gologger.Info().Msg("Subdomains")
	countSub := 0
	//dork for subdomain
	dork := "site:*." + options.Domain
	//googlesearch
	result, err := googlesearch.Search(ctx, dork, googlesearch.SearchOptions{Limit: options.Results})
	if len(result) == 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=site:*." + options.Domain + "")
		gologger.Error().Msgf("No subdomain found for domain %s\n", strings.ToLower(options.Domain))
	}
	//TODO : fix error when result > 0, shows google result but subdomain return 404
	if len(result) > 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=site:*." + options.Domain + "")
		gologger.Info().Msgf("Google result found for domain %s\n", strings.ToLower(options.Domain))
		for i := 0; i < len(result); i++ {
			invalidSub := regexp.MustCompile(options.Domain + `/.`)
			if !invalidSub.MatchString(result[i].URL) {
				masa := aurora.Cyan(masa.Format("[2006-01-02 15:04:05]"))
				fmt.Print(masa)
				fmt.Print(aurora.BrightYellow(" [Subdomain] "))
				fmt.Println(result[i].URL)
				countSub++
			}
		}
		if countSub == 0 {
			gologger.Error().Msgf("No subdomain found for domain %s\n", strings.ToLower(options.Domain))
		} else {
			gologger.Info().Msgf("Total subdomain found :%s\n", fmt.Sprint(countSub))
		}
	}
	if err != nil {
		gologger.Error().Msgf("Error found : %s\n", err)
	}
}

//exposed documents
func exposedDoc(options *args.Options) {
	gologger.Info().Msg("Exposed Documents")
	countDoc := 0
	//dork for exposed document
	dork := "site:" + options.Domain + "+ext:doc+| ext:docx | ext:odt | ext:pdf | ext:rtf | ext:sxw | ext:psw | ext:ppt | ext:pptx | ext:pps | ext:csv)"
	//googlesearch
	result, err := googlesearch.Search(ctx, dork, googlesearch.SearchOptions{Limit: options.Results})
	if len(result) == 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=site:" + options.Domain + "%20ext%3Adoc%20%7C%20ext%3Adocx%20%7C%20ext%3Aodt%20%7C%20ext%3Apdf%20%7C%20ext%3Artf%20%7C%20ext%3Asxw%20%7C%20ext%3Apsw%20%7C%20ext%3Appt%20%7C%20ext%3Apptx%20%7C%20ext%3Apps%20%7C%20ext%3Acsv)")
		gologger.Error().Msgf("No exposed document found for domain %s\n", strings.ToLower(options.Domain))
	}
	if len(result) > 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=site:" + options.Domain + "%20ext%3Adoc%20%7C%20ext%3Adocx%20%7C%20ext%3Aodt%20%7C%20ext%3Apdf%20%7C%20ext%3Artf%20%7C%20ext%3Asxw%20%7C%20ext%3Apsw%20%7C%20ext%3Appt%20%7C%20ext%3Apptx%20%7C%20ext%3Apps%20%7C%20ext%3Acsv)")
		gologger.Info().Msgf("Google result found for domain %s\n", strings.ToLower(options.Domain))
		for i := 0; i < len(result); i++ {
			masa := aurora.Cyan(masa.Format("[2006-01-02 15:04:05]"))
			fmt.Print(masa)
			fmt.Print(aurora.BrightYellow(" [Exposed Document] "))
			fmt.Println(result[i].URL)
			countDoc++
		}
		gologger.Info().Msgf("Total exposed document found :%s\n", fmt.Sprint(countDoc))
	}
	if err != nil {
		gologger.Error().Msgf("Error found : %s\n", err)
	}
}

//login/register/signup pages
func loginPages(options *args.Options) {
	gologger.Info().Msg("Login | Register | Signup pages")
	countPages := 0
	//dork for login/register/signup pages
	dork := "site:" + options.Domain + " inurl:signup | inurl:register | intitle:Signup"
	//googlesearch
	result, err := googlesearch.Search(ctx, dork, googlesearch.SearchOptions{Limit: options.Results})
	if len(result) == 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=site:" + options.Domain + "%20inurl%3Asignup%20%7C%20inurl%3Aregister%20%7C%20intitle%3ASignup")
		gologger.Error().Msgf("No login / register / signup pages found for domain %s\n", strings.ToLower(options.Domain))
	}
	if len(result) > 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=site:" + options.Domain + "%20inurl%3Asignup%20%7C%20inurl%3Aregister%20%7C%20intitle%3ASignup")
		gologger.Info().Msgf("Google result found for domain %s\n", strings.ToLower(options.Domain))
		for i := 0; i < len(result); i++ {
			masa := aurora.Cyan(masa.Format("[2006-01-02 15:04:05]"))
			fmt.Print(masa)
			fmt.Print(aurora.BrightYellow(" [Login Pages] "))
			fmt.Println(result[i].URL)
			countPages++
		}
		gologger.Info().Msgf("Total pages found :%s\n", fmt.Sprint(countPages))
	}
	if err != nil {
		gologger.Error().Msgf("Error found : %s\n", err)
	}

}

//exposed .git folder
//TODO filter result not containing .git
func gitFolder(options *args.Options) {
	gologger.Info().Msg("Exposed .git page")
	countGit := 0
	//dork for .git folder
	//dork := "intitle:index of /.git/hooks \"" + options.Domain + "\""
	dork := "intext:index of /.git parent directory " + options.Domain
	//googlesearch
	result, err := googlesearch.Search(ctx, dork, googlesearch.SearchOptions{Limit: options.Results})
	if len(result) == 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=intext:index%20of%20/.git%20parent%20directory%22" + options.Domain + "%22")
		gologger.Error().Msgf("No .git folder found for domain %s\n", strings.ToLower(options.Domain))
	}
	if len(result) > 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=intext:index%20of%20/.git%20parent%20directory%22" + options.Domain + "%22")
		gologger.Info().Msgf("Google result found for domain %s\n", strings.ToLower(options.Domain))
		for i := 0; i < len(result); i++ {
			masa := aurora.Cyan(masa.Format("[2006-01-02 15:04:05]"))
			fmt.Print(masa)
			fmt.Print(aurora.BrightYellow(" [Git] "))
			fmt.Println(result[i].URL)
			countGit++
		}
		gologger.Info().Msgf("Total .git found :%s\n", fmt.Sprint(countGit))
	}
	if err != nil {
		gologger.Error().Msgf("Error found : %s\n", err)
	}
}

//wordpress file
func findWord(options *args.Options) {
	gologger.Info().Msg("Wordpress file")
	countWord := 0
	//dork for wordpress file
	dork := "site:" + options.Domain + " inurl:wp-content | inurl:wp-includes"
	//googlesearch
	result, err := googlesearch.Search(ctx, dork, googlesearch.SearchOptions{Limit: options.Results})
	if len(result) == 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=site:" + options.Domain + "%20inurl:wp-content%20|%20inurl:wp-includes")
		gologger.Error().Msgf("No wordpress file found for domain %s\n", strings.ToLower(options.Domain))
	}
	if len(result) > 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=site:" + options.Domain + "%20inurl:wp-content%20|%20inurl:wp-includes")
		gologger.Info().Msgf("Google result found for domain %s\n", strings.ToLower(options.Domain))
		for i := 0; i < len(result); i++ {
			masa := aurora.Cyan(masa.Format("[2006-01-02 15:04:05]"))
			fmt.Print(masa)
			fmt.Print(aurora.BrightYellow(" [Wordpress] "))
			fmt.Println(result[i].URL)
			countWord++
		}
		gologger.Info().Msgf("Total wordpress found :%s\n", fmt.Sprint(countWord))
	}
	if err != nil {
		gologger.Error().Msgf("Error found : %s\n", err)
	}
}

//PHP Error
func findPHP(options *args.Options) {
	gologger.Info().Msg("PHP Error")
	countPhp := 0
	//dork for php error
	dork := "site:" + options.Domain + " \"PHP Parse error\" | \"PHP Warning\" | \"PHP Error\""
	//googlesearch
	result, err := googlesearch.Search(ctx, dork, googlesearch.SearchOptions{Limit: options.Results})
	if len(result) == 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=site:" + options.Domain + "%20%22PHP%20Parse%20error%22%20|%20%22PHP%20Warning%22%20|%20%22PHP%20Error%22")
		gologger.Error().Msgf("No php error found for domain %s\n", strings.ToLower(options.Domain))
	}
	if len(result) > 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=site:" + options.Domain + "%20%22PHP%20Parse%20error%22%20|%20%22PHP%20Warning%22%20|%20%22PHP%20Error%22")
		gologger.Info().Msgf("Google result found for domain %s\n", strings.ToLower(options.Domain))
		for i := 0; i < len(result); i++ {
			masa := aurora.Cyan(masa.Format("[2006-01-02 15:04:05]"))
			fmt.Print(masa)
			fmt.Print(aurora.BrightYellow(" [PHP Error] "))
			fmt.Println(result[i].URL)
			countPhp++
		}
		gologger.Info().Msgf("Total php error found :%s\n", fmt.Sprint(countPhp))
	}
	if err != nil {
		gologger.Error().Msgf("Error found : %s\n", err)
	}
}

//Possible SQLi
func findSQL(options *args.Options) {
	gologger.Info().Msg("SQL Error")
	countSQL := 0
	//dork for php error
	dork := "inurl:\".php?id=\" \"You have an error in your SQL syntax\"" + options.Domain
	//googlesearch
	result, err := googlesearch.Search(ctx, dork, googlesearch.SearchOptions{Limit: options.Results})
	if len(result) == 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=inurl:%22.php?id=%22%20%22You%20have%20an%20error%20in%20your%20SQL%20syntax%22%20" + options.Domain)
		gologger.Error().Msgf("No sql error found for domain %s\n", strings.ToLower(options.Domain))
	}
	if len(result) > 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=inurl:%22.php?id=%22%20%22You%20have%20an%20error%20in%20your%20SQL%20syntax%22%20" + options.Domain)
		gologger.Info().Msgf("Google result found for domain %s\n", strings.ToLower(options.Domain))
		for i := 0; i < len(result); i++ {
			masa := aurora.Cyan(masa.Format("[2006-01-02 15:04:05]"))
			fmt.Print(masa)
			fmt.Print(aurora.BrightYellow(" [SQL Error] "))
			fmt.Println(result[i].URL)
			countSQL++
		}
		gologger.Info().Msgf("Total sql error found :%s\n", fmt.Sprint(countSQL))
	}
	if err != nil {
		gologger.Error().Msgf("Error found : %s\n", err)
	}
}

//Backup file
func findBak(options *args.Options) {
	gologger.Info().Msg("Backup File")
	countBak := 0
	//dork for php error
	dork := "site:" + options.Domain + " ext:bkf | ext:bkp | ext:bak | ext:old | ext:backup"
	//googlesearch
	result, err := googlesearch.Search(ctx, dork, googlesearch.SearchOptions{Limit: options.Results})
	if len(result) == 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=site:" + options.Domain + "%20ext:bkf%20|%20ext:bkp%20|%20ext:bak%20|%20ext:old%20|%20ext:backup")
		gologger.Error().Msgf("No backup file found for domain %s\n", strings.ToLower(options.Domain))
	}
	if len(result) > 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=site:" + options.Domain + "%20ext:bkf%20|%20ext:bkp%20|%20ext:bak%20|%20ext:old%20|%20ext:backup")
		gologger.Info().Msgf("Google result found for domain %s\n", strings.ToLower(options.Domain))
		for i := 0; i < len(result); i++ {
			masa := aurora.Cyan(masa.Format("[2006-01-02 15:04:05]"))
			fmt.Print(masa)
			fmt.Print(aurora.BrightYellow(" [Backup File] "))
			fmt.Println(result[i].URL)
			countBak++
		}
		gologger.Info().Msgf("Total backup file found :%s\n", fmt.Sprint(countBak))
	}
	if err != nil {
		gologger.Error().Msgf("Error found : %s\n", err)
	}
}

//config file
func findConfig(options *args.Options) {
	gologger.Info().Msg("Config Files")
	countConfig := 0
	//dork for config file
	dork := "site:" + options.Domain + " ext:xml | ext:conf | ext:cnf | ext:reg | ext:inf | ext:rdp | ext:cfg | ext:txt | ext:ora | ext:env | ext:ini"
	//googlesearch
	result, err := googlesearch.Search(ctx, dork, googlesearch.SearchOptions{Limit: options.Results})
	if len(result) == 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=site:" + options.Domain + "%20ext:xml%20|%20ext:conf%20|%20ext:cnf%20|%20ext:reg%20|%20ext:inf%20|%20ext:rdp%20|%20ext:cfg%20|%20ext:txt%20|%20ext:ora%20|%20ext:env%20|%20ext:ini")
		gologger.Error().Msgf("No config file found for domain %s\n", strings.ToLower(options.Domain))
	}
	if len(result) > 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=site:" + options.Domain + "%20ext:xml%20|%20ext:conf%20|%20ext:cnf%20|%20ext:reg%20|%20ext:inf%20|%20ext:rdp%20|%20ext:cfg%20|%20ext:txt%20|%20ext:ora%20|%20ext:env%20|%20ext:ini")
		gologger.Info().Msgf("Config file found for domain %s\n", strings.ToLower(options.Domain))
		for i := 0; i < len(result); i++ {
			masa := aurora.Cyan(masa.Format("[2006-01-02 15:04:05]"))
			fmt.Print(masa)
			fmt.Print(aurora.BrightYellow(" [Config File] "))
			fmt.Println(result[i].URL)
			countConfig++
		}
		gologger.Info().Msgf("Total config found :%s\n", fmt.Sprint(countConfig))
	}
	if err != nil {
		gologger.Error().Msgf("Error found :%s\n", err)
	}
}

//find database file
func findDB(options *args.Options) {
	gologger.Info().Msg("Database files")
	countDb := 0
	//dork for database file
	dork := "site:" + options.Domain + " ext:sql | ext:dbf | ext:mdb | ext:db"
	//googlesearch
	result, err := googlesearch.Search(ctx, dork, googlesearch.SearchOptions{Limit: options.Results})
	if len(result) == 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=site:" + options.Domain + "%20ext:sql%20|%20ext:dbf%20|%20ext:mdb%20|%20ext:db")
		gologger.Error().Msgf("No database file found for domain %s\n", strings.ToLower(options.Domain))
	}
	if len(result) > 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=site:" + options.Domain + "%20ext:sql%20|%20ext:dbf%20|%20ext:mdb%20|%20ext:db")
		gologger.Info().Msgf("Database file found for domain %s\n", strings.ToLower(options.Domain))
		for i := 0; i < len(result); i++ {
			masa := aurora.Cyan(masa.Format("[2006-01-02 15:04:05]"))
			fmt.Print(masa)
			fmt.Print(aurora.BrightYellow(" [Database File] "))
			fmt.Println(result[i].URL)
			countDb++
		}
		gologger.Info().Msgf("Total database file found :%s\n", fmt.Sprint(countDb))
	}
	if err != nil {
		gologger.Error().Msgf("Error found :%s\n", err)
	}
}

//finding Sub-subdomain
func findSubSubdomain(options *args.Options) {
	gologger.Info().Msg("Sub-Subdomains")
	countSubs := 0
	//dork for subdomain
	dork := "site:*.*." + options.Domain
	//googlesearch
	result, err := googlesearch.Search(ctx, dork, googlesearch.SearchOptions{Limit: options.Results})
	if len(result) == 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=site:*.*." + options.Domain + "")
		gologger.Error().Msgf("No sub-subdomain found for domain %s\n", strings.ToLower(options.Domain))
	}
	//TODO : fix error when result > 0, shows google result but subdomain return 404
	if len(result) > 0 {
		printDork()
		gologger.Print().Msg(" Dorking https://www.google.com/search?q=site:*.*." + options.Domain + "")
		gologger.Info().Msgf("Google result found for domain %s\n", strings.ToLower(options.Domain))
		for i := 0; i < len(result); i++ {
			invalidSub := regexp.MustCompile(options.Domain + `/.`)
			if !invalidSub.MatchString(result[i].URL) {
				masa := aurora.Cyan(masa.Format("[2006-01-02 15:04:05]"))
				fmt.Print(masa)
				fmt.Print(aurora.BrightYellow(" [Sub-subdomain] "))
				fmt.Println(result[i].URL)
				countSubs++
			}
		}
		if countSubs == 0 {
			gologger.Error().Msgf("No sub-subdomain found for domain %s\n", strings.ToLower(options.Domain))
		} else {
			gologger.Info().Msgf("Total sub-subdomain found :%s\n", fmt.Sprint(countSubs))
		}
	}
	if err != nil {
		gologger.Error().Msgf("Error found : %s\n", err)
	}
}

//print DORK in color
func printDork() {
	fmt.Print("[")
	fmt.Print(aurora.Magenta("DRK"))
	fmt.Print("]")
}

//scan time taken
func timeTaken() {
	now := time.Now()
	defer func() {
		gologger.Info().Msgf("Scan execution took %s", time.Since(now))
	}()
}
