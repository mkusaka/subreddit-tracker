# golang
## [1][Atari2600 Emulator - Gopher2600](https://www.reddit.com/r/golang/comments/fmi0ok/atari2600_emulator_gopher2600/)
- url: https://www.reddit.com/r/golang/comments/fmi0ok/atari2600_emulator_gopher2600/
---
Hello,

A couple of years ago I started to build an Atari 2600 emulator as a way of learning about that old games console. To heap complication on top of confusion I also decided to use the project as an opportunity to learn a new language, Go.

As it happens, Go has been really fun to use and I'm glad I chose it.

Today, I've completed what I consider to be version 0.1 of the project. This means it's not finished, I still have ideas I want to implement, but it's good enough to use as an emulator to play games. It also has a useful debugger which should be useful when developing new ROMs.

The emulation is very accurate. I won't say *completely* accurate just yet but it's close. In particular, there are a couple of cartridge formats that it doesn't yet handle.

From a Go perspective, it only uses a handful of third party packages. Notably it uses inkyblackness/imgui-go, a port of the excellent Dear Imgui. (Although currently, gopher2600 pulls from my fork of the project because a couple of required additions haven't yet been merged.)

The project can be found on Github and I'd love it if you can have a look and make comments on my Go programming decisions. In particular, I'm looking for ideas on how to improve performance. I can get a maximum of 70fps  (headless emulation) on my machine which is good but it could do with more headroom for better performance in the debugger.

Enjoy

https://github.com/JetSetIlly/Gopher2600
## [2][Open Source Identity and Access Management](https://www.reddit.com/r/golang/comments/fmofd2/open_source_identity_and_access_management/)
- url: https://www.reddit.com/r/golang/comments/fmofd2/open_source_identity_and_access_management/
---
Hi all   

I am looking for Open Source Identity and Access Management like https://www.keycloak.org/ but for modern application. What I do not like about Keycloak is, that it is not flexible, for example changing the login page is cumbersome.    

Does it exist Open Source Identity and Access Management that is written in Go?   

Thanks
## [3][First Project in Go: Creep - a CLI tool for downloading random images.](https://www.reddit.com/r/golang/comments/fmic2e/first_project_in_go_creep_a_cli_tool_for/)
- url: https://www.reddit.com/r/golang/comments/fmic2e/first_project_in_go_creep_a_cli_tool_for/
---
I'm learning Go and recently built a CLI tool, [creep](https://github.com/splode/creep). It's used to download images from URLs that provide raw images, such as https://thispersondoesnotexist.com. It's a bit of a silly utility, but I use when seeding application databases and I'm using it as a learning exercise.

This is my first project in Go, so I'd be happy to hear any feedback.

Thanks!

https://preview.redd.it/b3xk498832o41.png?width=1200&amp;format=png&amp;auto=webp&amp;s=3389d636befd3b2996d1098abf7f979652a5d626
## [4][Any open source projects that need help (personal projects or larger-scale e.g. K8s)?](https://www.reddit.com/r/golang/comments/fmp54i/any_open_source_projects_that_need_help_personal/)
- url: https://www.reddit.com/r/golang/comments/fmp54i/any_open_source_projects_that_need_help_personal/
---
I've been using Golang at work for about a year and a half now, and have fallen in love with the language and ecosystem. However I feel like I've fallen into a comfort zone within my repo at work, and would like to branch out and explore/contribute to other Golang projects.

Does anyone have a personal project they need help on, or can recommend a fairly welcoming open-source repo to start learning and making PRs to? Thanks!
## [5][[help] Unable to setup go modules with private Bitbucket with go module (normal go get works)](https://www.reddit.com/r/golang/comments/fmrmbq/help_unable_to_setup_go_modules_with_private/)
- url: https://www.reddit.com/r/golang/comments/fmrmbq/help_unable_to_setup_go_modules_with_private/
---
We have a private bitbucket accessible at: stash.mycompany.com/project/repo. Its only available to access over ssh. Till now our projects are pre go modules so a git config like:

    [url "ssh://git@stash.mycompany.com"]
	 insteadOf = https://stash.mycompany.com/scm

Unfortunately, we are trying to move to go modules and when we try to get the above, it seems it goes to sum.golang.org which then looks up HTTPS version. Doing a go get comes up as:

https://gist.github.com/psykidellic/ef900a49c53b5683847dbd8cc2397c27

If you look at the gist, it even gets the correct version v0.0.9 

Since the thing is behind corporate firewall (over VPN) and has no https access, that obviously does not work.

How can I set this up?
## [6][What do employers look in GitHub repo of golang services?](https://www.reddit.com/r/golang/comments/fma0dz/what_do_employers_look_in_github_repo_of_golang/)
- url: https://www.reddit.com/r/golang/comments/fma0dz/what_do_employers_look_in_github_repo_of_golang/
---
Apologies if already asked, I have started learning golang and am writing a few backend services. What are the things that employers specifically seek in such personal projects that could help me stand out?
## [7][How to distribute Go executables](https://www.reddit.com/r/golang/comments/fmn0jn/how_to_distribute_go_executables/)
- url: https://www.reddit.com/r/golang/comments/fmn0jn/how_to_distribute_go_executables/
---
Sorry if this question is simplistic, as I am new to golang. I've just finished my first go project, and would like to send the executable to a friend, without them needing to compile my code from source. I've tested this by sending myself the executable (via dropbox), but when I try and open it, my computer only recognizes it as a text file, and will not let me execute it. How does one properly and elegantly distribute go binaries without running into this problem?
## [8][golang-dev &amp; the next few months](https://www.reddit.com/r/golang/comments/flzoc1/golangdev_the_next_few_months/)
- url: https://groups.google.com/forum/#!topic/golang-dev/UxvN1W2B-zg
---

## [9][Go-Corona - A Golang client library for accessing global coronavirus (COVID-19, SARS-CoV-2) outbreak data.](https://www.reddit.com/r/golang/comments/fmitic/gocorona_a_golang_client_library_for_accessing/)
- url: https://www.reddit.com/r/golang/comments/fmitic/gocorona_a_golang_client_library_for_accessing/
---
[Go-Corona](https://preview.redd.it/4e33ftitb2o41.jpg?width=1200&amp;format=pjpg&amp;auto=webp&amp;s=7d4ef186f98b0efc17ce991a0f0408f7846de6f5)

Coronavirus has been spreading rapidly across the world, affecting more than 160 countries. So, I completed an **#opensource** project for accessing global coronavirus (COVID-19, SARS-CoV-2) outbreak data using **#golang** so that you could use it in your apps and spread more awareness around the world.

**Check out:** [**https://github.com/itsksaurabh/go-corona**](https://github.com/itsksaurabh/go-corona)

It fetches data from the data repository operated by the The Johns Hopkins University Center for Systems Science and Engineering (JHU CSSE). More data sources to be added later.

Feel free to share, discuss and contribute.

**#corona** **#coronavirus** **#covid19** **#go** **#golang** **#opensource** **#gocorona**
## [10][Design pattern to avoid hitting the rate limit on an external API resource](https://www.reddit.com/r/golang/comments/flyvzg/design_pattern_to_avoid_hitting_the_rate_limit_on/)
- url: https://www.reddit.com/r/golang/comments/flyvzg/design_pattern_to_avoid_hitting_the_rate_limit_on/
---
I have to get the PR reviews from hundreds of repos, lets say I have these in batches of 20 PRs for ever repo

like this 

    {RepoA: #1, #2, ..........#20}
    {RepoA: #21, #22, ..........#40}
    .
    .
    {RepoZ: #1, #2, ..........#20}

Something like this, and I want to run goroutines for all these different batches, but the Github API limit is 5000 requests per hour

Is there a design pattern that would apply to this problem to solve it in an elegant way,

also should I launch the goroutines for a specific number of batches and use the rest later or should I run all the goroutines and pause them when I hit the rate limit and resume them later on once a certain amount of time has passed?

I know for server-side APIs there are patterns like token bucket and leaky bucket and all but I am a client of these APIs is there something similar for throttling/regulating the frequency of requests from my application?
