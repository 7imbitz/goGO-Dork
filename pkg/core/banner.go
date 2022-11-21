package core

import (
	"fmt"

	"github.com/projectdiscovery/gologger"
)

// Version is the current version of goGo-Dork
const Version = `1.0.0`

// Author of the tools
const Author = `7imbitz`

// goGo-Dork logo
var banner = fmt.Sprintf(`
                __________        ____             __  
   ____ _____  / ____/ __ \      / __ \____  _____/ /__
  / __ '/ __ \/ / __/ / / /_____/ / / / __ \/ ___/ //_/
 / /_/ / /_/ / /_/ / /_/ /_____/ /_/ / /_/ / /  / ,<   
 \__, /\____/\____/\____/     /_____/\____/_/  /_/|_|  
/____/                                
                                      %s - %s
`, Version, Author)

// showBanner is used to show the banner to the user
func ShowBanner() {
	gologger.Print().Msgf("%s", banner)
	gologger.Info().Msgf("goGO-Dork version %s", Version)
	gologger.Print().Label("INF").Msgf("This is a beta version")
}
