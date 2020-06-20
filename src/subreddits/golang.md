# golang
## [1][Worldwide - a GameBoyColor emulator written in Go (Ebiten)](https://www.reddit.com/r/golang/comments/hchnmh/worldwide_a_gameboycolor_emulator_written_in_go/)
- url: https://www.reddit.com/r/golang/comments/hchnmh/worldwide_a_gameboycolor_emulator_written_in_go/
---
[https://github.com/Akatsuki-py/Worldwide](https://github.com/Akatsuki-py/Worldwide)

https://i.redd.it/9znnudgpm0651.gif
## [2][OpenDiablo2 News: The open-source Diablo 2 engine written in Go](https://www.reddit.com/r/golang/comments/hc667q/opendiablo2_news_the_opensource_diablo_2_engine/)
- url: https://www.reddit.com/r/golang/comments/hc667q/opendiablo2_news_the_opensource_diablo_2_engine/
---
This is a repost from [r/OpenDiablo2](https://www.reddit.com/r/OpenDiablo2/comments/hc5w6f/news_multiplayer_mapgen_and_patreon/)

Just a quick status update:

* OpenDiablo2 now has a functional UDP client/server network implementation
* More data dictionary loaders for the txt files have been added
* Rudimentary map-generation work has begun, added the ability to stitch discrete map regions (pre-defined map sections). Character path-finding works between the regions.
* The game audio/input/rendering framework [Ebiten](https://ebiten.org/) has been updated and increased performance
* [A Patreon page has been set up](https://www.patreon.com/user?u=37261055&amp;fan_landing=true) \- We will be using the money to guide development efforts, offer bug bounties, and pay developers/designers for their work. Please, consider becoming a patron!

The project website can be found at [https://opendiablo2.com/](https://opendiablo2.com/)

The source can be found on Github at [https://github.com/OpenDiablo2/OpenDiablo2](https://github.com/OpenDiablo2/OpenDiablo2)

If you want to get involved, ask questions, suggest features, or just hang out you can find us over on our [Discord server](https://discord.gg/5jd4CwW)
## [3][People's efficiency expectations for generics in 'Go 2' and patterns of use](https://www.reddit.com/r/golang/comments/hcce4e/peoples_efficiency_expectations_for_generics_in/)
- url: https://utcc.utoronto.ca/~cks/space/blog/programming/Go2GenericsExpectedEfficiency
---

## [4][Integrate twitter login with golang application](https://www.reddit.com/r/golang/comments/hcm2ga/integrate_twitter_login_with_golang_application/)
- url: https://www.reddit.com/r/golang/comments/hcm2ga/integrate_twitter_login_with_golang_application/
---
I wrote a [blog](https://www.loginradius.com/engineering/blog/twitter-authentication-with-golang-and-goth/) on integrating twitter login with **golang** application and share it on [Github](https://github.com/LoginRadius/engineering-blog-samples/tree/master/GoLang/TwitterAuthenticationGoth) too. Please check this out and let me know if it helps uh in same
## [5][Go DataDog](https://www.reddit.com/r/golang/comments/hclrpj/go_datadog/)
- url: https://www.reddit.com/r/golang/comments/hclrpj/go_datadog/
---
Does anyone use Datadog and write integrations using Go? I know there is a python implemention([https://docs.datadoghq.com/developers/write\_agent\_check/?tab=agentv6v7](https://docs.datadoghq.com/developers/write_agent_check/?tab=agentv6v7)) but I have a preference to write my checks in Go, obviously.
## [6][How to change terminal/console output color in golang?](https://www.reddit.com/r/golang/comments/hcl9f1/how_to_change_terminalconsole_output_color_in/)
- url: https://www.reddit.com/r/golang/comments/hcl9f1/how_to_change_terminalconsole_output_color_in/
---

## [7][My Learning Path - GoLang](https://www.reddit.com/r/golang/comments/hcgctg/my_learning_path_golang/)
- url: https://www.reddit.com/r/golang/comments/hcgctg/my_learning_path_golang/
---
Hello guys,

I'm new to GoLang and came up with a small learning path for myself to get started with GoLang.Posting it here to get suggestions whether I'm going in right directions or not.

Also, I'm forcing myself to learn Go concepts/APIs first (as mentioned in **Directions to Follow**) from official docs and refer to other sources if I don't understand from official docs. Is this the right way to do it?

**Directions to Follow**

* Use official docs as primary source, this way I will learn how to read Go Packages and use it implement in my code.
* Use `other sources` if I’m really stuck on some tasks and not able to figure out from Official docs. (Because I can learn reading Docs and understanding APIs form Docs)

&amp;#x200B;

**Official Reference**

* Go Doc - [https://golang.org/doc/](https://golang.org/doc/)
* Go Packages - [https://golang.org/pkg/](https://golang.org/pkg/)
* Effective Go - [https://golang.org/doc/effective\_go.html](https://golang.org/doc/effective_go.html)

&amp;#x200B;

**Useful Links**

* [https://pmihaylov.com/](https://pmihaylov.com/)
* [https://medium.com/a-journey-with-go](https://medium.com/a-journey-with-go)
* [https://quii.gitbook.io/learn-go-with-tests/](https://quii.gitbook.io/learn-go-with-tests/)

&amp;#x200B;

**Basic**

* Get familiar with Go syntax - [https://gobyexample.com/](https://gobyexample.com/)
* Get familiar with Go project structure
* Build a basic App `Calculator-v1`
   * Ask user to select which arithmetic operation to do
   * Ask user to input 2 number to do operation on
   * Write test cases
   * Return result
* Build a basic App `Calculator-v2`
   * Think how to refactor `Calculator-v1`
   * Explore using command line arguments in Go
   * More test cases

&amp;#x200B;

**Intermediate**

* Understand Methods (Pointers vs Values) - [https://golang.org/doc/effective\_go.html#methods](https://golang.org/doc/effective_go.html#methods)
* Understand Interfaces - [https://golang.org/doc/effective\_go.html#interfaces\_and\_types](https://golang.org/doc/effective_go.html#interfaces_and_types)
* Understand Errors - [https://golang.org/doc/effective\_go.html#errors](https://golang.org/doc/effective_go.html#errors)
* Explore the net/http package - [https://golang.org/pkg/net/http/](https://golang.org/pkg/net/http/)
* Build an App `GithubTopRepos-v1`
   * Use Github APIs to get the top 10 repos for a particular programming language
* To be continued…

Feedback is much appreciated :)Thank you :)
## [8][Tested the new generics proposal and I think they will be a game changer](https://www.reddit.com/r/golang/comments/hck2sz/tested_the_new_generics_proposal_and_i_think_they/)
- url: https://medium.com/swlh/the-new-generics-proposal-tested-b9bd24f86547
---

## [9][Very basic doubt. Craving for some help :'(](https://www.reddit.com/r/golang/comments/hcijn3/very_basic_doubt_craving_for_some_help/)
- url: https://www.reddit.com/r/golang/comments/hcijn3/very_basic_doubt_craving_for_some_help/
---
--------------------
TLDR;
GoLand IDE not looking into $GOPATH/pkg/mod for the packages of my go project and hence, giving issues and red highlights on code where I am importing packages. GoLand seems to only look into $GOPATH/src for the packages where packages are not there coz packages got downloaded in pkg/mod where GoLand is not looking into.
----------------------
So, the thing is
I was going with open source dev with go...
: So, I forked and cloned a repo ....

did a "go get -d ./..."

And all the required modules and all got downloaded under the directory $GOPATH/pkg/mod
Now, I can "go build" the project ... It's all cool and all
BUT BUT BUT
BUT BUT BUT

when I am opening the project in GoLand IDE... It is giving issues and red highlights while importing the packages... And the reason seems that GoLand is not looking inside 
$GOPATH/pkg/mod

GoLand seems to look inside $GOPATH/src for any package and hence, provides auto suggestions accordingly but well, in my case the packages and all got downloaded in $GOPATH/pkg/mod.

Now, golang doesn't have a problem... Because while doing "go build", it happily looks into pkg/mod finds the required packages and builds the binary

"The problem is with GoLand not looking into pkg/mod and hence, not providing autosuggestion and just highlight the import statements as red  because it thinks the package does not exist"
I have been stuck at it since a long bloody time.... Rambled through soo many stackoverflow posts and all,... But nothing worked out :'(
## [10][Golang Vancouver Online Meetup on July 8](https://www.reddit.com/r/golang/comments/hcbzeb/golang_vancouver_online_meetup_on_july_8/)
- url: https://www.reddit.com/r/golang/comments/hcbzeb/golang_vancouver_online_meetup_on_july_8/
---
YouTube live steam at [https://www.youtube.com/watch?v=vjmoNQmXAeI](https://www.youtube.com/watch?v=vjmoNQmXAeI). RSVP at [https://www.meetup.com/golangvan/events/271386675](https://www.meetup.com/golangvan/events/271386675/)
