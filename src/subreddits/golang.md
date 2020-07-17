# golang
## [1][Go 1.14.6 is out now](https://www.reddit.com/r/golang/comments/hsqdak/go_1146_is_out_now/)
- url: https://www.reddit.com/r/golang/comments/hsqdak/go_1146_is_out_now/
---
[https://golang.org/dl/](https://golang.org/dl/)  
[https://golang.org/doc/go1.14](https://golang.org/doc/go1.14)  
[https://golang.org/doc/devel/release.html#go1.14](https://golang.org/doc/devel/release.html#go1.14)
## [2][GoLand 2020.2 hits Beta with experimental support for generics a.k.a. type parameters](https://www.reddit.com/r/golang/comments/hsgx4r/goland_20202_hits_beta_with_experimental_support/)
- url: https://blog.jetbrains.com/2020/07/16/goland-2020-2-reaches-beta
---

## [3][InfraMap: Terraform infrastructure diagram showing only the resources that are relevant](https://www.reddit.com/r/golang/comments/hsuhe0/inframap_terraform_infrastructure_diagram_showing/)
- url: /r/devops/comments/hsavh3/inframap_terraform_infrastructure_diagram_showing/
---

## [4][I Automated the Official Go Install Guide, for $HOME (and %USERPROFILE%)](https://www.reddit.com/r/golang/comments/hsumy5/i_automated_the_official_go_install_guide_for/)
- url: https://www.reddit.com/r/golang/comments/hsumy5/i_automated_the_official_go_install_guide_for/
---
Pretty nifty, eh?

```
curl -fsS https://webinstall.dev/golang@1.14 | bash
```

And I'll explain just how this solves _many_ of your Go-install woes, but first:

# Wait Before Hating

I know this is going to get a lot of hate.

I believe most people are going to either **love this** (because we share similar pains), **or not understand** the problem it solves (which is fine - **we don't all have the same problems**).

But the small number of those who really hate this, really hate it... and I don't get why.

So if you've get an **emotional knee-jerk reaction**, or a have **moral argument**, such as

- _I_ would _never_ do it that way! (but why not?)
- You _shouldn't_ do that! (by what authority?)
- This is not _the way_! (says who!? why?)
- You've _offended_ my sensibilities! (do they deserve an upgrade?)

Please note that I'm trying to solve a technical problem that I have, that I know many other (but not all) people have.

There's literally nothing I know of to do that can solve the non-technical dislike for my solution.

As such, it's **not productive** - not for me, nor anyone else - to just say something to the effect of "This is such a crappy solution, it's so bad, ugh" - because such a sentence carries **NO ACTIONABLE INFORMATION** - I don't learn anything. No one else learns anything. None of us know how to make it better.

So before you **downvote** or **leave a hasty comment** I challenge you to **read** my explanation
and come up with a **valid logical argument** that identifies something I've done that violates a
technical principle (and if you do, I can probably fix it).

If you read this and it's apparent to you how I could better explain this in a way that is more appropriate, please let me know.

# The Official Go Install Guide

If you haven't Googled "install golang" and then read https://golang.org/doc/install, I'd suggest that you do.

The explanation is really good, but it's almost too much explanation, which can make it difficult to follow, maybe more so if you're not as strong in the Linux-fu and bash-fu.

Essentially, it says to do this:

- Download
- Unpack
- Add to PATH

That looks kinda like this:

```
# The hard part (sans GUI)
#OFFICIAL_RELEASE_API=https://golang.org/dl/?mode=json
#OFFICIAL_DOWNLOAD_URL=$(curl "$OFFICIAL_RELEASE_API" | jq magic magic magic | sort-by-latest-stable-acceptable-version)

# The easy part
curl "$OFFICIAL_DOWNLOAD_URL" -o golang-xxx.tar.gz
tar xvf golang-xxx.tar.gz
mv ./go ~/.local/opt/go
echo 'PATH="$HOME/.local/opt/go/bin:$HOME/go/bin:$PATH"' &gt; ~/.profile
export PATH="$HOME/.local/opt/go/bin:$HOME/go/bin:$PATH"
```

Obviously I kinda hand-waved away "the hard part"... EXCEPT THAT I DIDN'T (I did the work for it!):

- https://webinstall.dev/api/releases/golang@1.14.tab?os=macos&amp;arch=amd64&amp;ext=tar,zip,xz&amp;limit=100
    - `1.14.6	lts	stable	1970-01-01	macos	amd64	tar.gz	-	https://dl.google.com/go/go1.14.6.darwin-amd64.tar.gz`
- https://github.com/webinstall/packages/blob/master/golang/releases.js
    - (yeah yeah, it's node, I'm a traitor - or node.js is just better for quick-n-dirty JSON transforms: fight me)

# I automated it. And many other things.

This is my solution: https://webinstall.dev/golang

So we've got it Mac and Linux, which you saw before:

```
curl -fsS https://webinstall.dev/golang@1.14 | bash
```

But wait, there's more...

Please join me in welcoming (drum roll please) Windows 10 to the new age!!

```
curl.exe -fsSA "MS" https://webinstall.dev/golang@1.14 | powershell
```

(Windows 10 has come a long way. It has `curl.exe`, `tar.exe`, `ssh.exe`, and all sorts of goodies _baked in_ - **NO WSL** or msysgit, etc needed)

And one last thing:

You also get a _teensy_ tiny `bash` or `powershell` script called `webi` that lets you do things like this:

```
webi golang@stable
webi golang@beta
webi golang@1.14
```

And whatever you build with version _x_ stays with version _x_. No mismatched binaries or caches!

Version switching, without a package manager.

(basically it's just a wrapper around curl that sets the "User-Agent" header to the proper OS and CPU arch so that it can query the webi API to grab the correct build)

W00t!

## It. Is. Awesome!

Hate it if you wish, but it's awesome.

"Awesome" is one of those terrible moral / emotional arguments, so let me break that down technically:

- it's **memorable**
    - (no need to look up documentation - I'm a _dev_, I use _webinstall_, duh)
    - (and `go` redirects to `golang`, if you ever forget)
- it's **convenient**
    - (cute little website, cute little shell command, cute little cheat sheet, how cute!)
    - (and cross-platform to boot! hooray!)
- it's **FAST**
    - (just download what you NEED - not a whole freaking last 10 years and 600gb of brew commit history... guh!)
- it **WILL NOT** rely on or screw up your **system permissions**
    - (no more having `node`, `brew`, `apt` and `go` all fighting over who owns `/usr/local`!!)
- Look Ma, **NO ADMIN**
   - (sans `sudo`, sans root, sans... Windows UAC?)
   - (exception: `spctl` on macOS behave different between High Sierra and Catalina, so sometimes it prompts when it shouldn't, #WILLFIX)
- it gives you the **OFFICIAL BUILD**
    - (grrrr `apt`!! always giving weird patches and strange paths... hate 3rd party builds)
- it's **UP TO DATE**
    - (grrrrr `apt` again!! always 6 months to 6 years behind...)
    - (max of 15-minute cache lag behind golang.org's Release API)

Basically, anything that I use on a regular basis that is self-contained and dependency-free (as are `go`, `node`, `rust`, `jq`, `postgres` and just about **every other modern tool that I care about**), I'm adding to webinstall.dev. (and so could you!)

# Deflecting the Arguments

If you love what I did, that's all. Thumbs up if you've got 'em. :)

If you hate what I'm doing... this is really just for the benefit of others coming by, because I don't think any sort of logical argument will convince you that I actually did something good that will benefit the community at large.

## STOP! Read this!

Before you go any further I want you to do a quick search of your `.bash_history` and look at your last 10 or so of these:

```
go run xxx
```

```
npm install xxx
```

Before you _dare_ leave a nasty comment:

- look at all those **unverified** `curl | go` or `curl | node` that you did
- look at the list of _HUNDREDS_ or _THOUSANDS_ of completely random, often **anonymous**, authors that you "trust"

Just let that... hypocrisy... sink in for a good 5 seconds before you hit me with your best shot.

If you don't understand what _npm postinstall scripts_ are, or how anything you `go get` can `init()` or `import _ xxxx` to run with the same `sudo` permissions that you used to bind your app to ports 80 and 443... then you need to go understand that.

## NOT a replacement for `apt` / `brew` / `x` - it's just better

You want some 6 year-old package with 100,000 dependencies? Well... through some odd combination of `apt` and `npm` and I'm sure you can manage (maybe `sudo apt install nodejs; sudo npm install create-react-app`?).

Anyway, lots of software needs hard-core package managers because it's complicated software.

Go is NOT LIKE THAT.

- Go is simple!
- Go is self-contained!
- Go doesn't even need libstd! (`CGO_ENABLED=0`, ftw)
- Go programs can run in [Docker scratch](https://rollout.io/blog/building-minimal-docker-containers-for-go-applications/)! 

`webi` is for simple things. `go` and `webi` can be happy together.

## curl | bash is bad because bobby blogger said so...

False. And True. But mostly False.

Consider this:

&gt; Claim: Flying in a plane is dangerous. (true)
&gt; 
&gt; Counter-Claim: Flying in a plane is not _more dangerous_ than driving a car, riding a bike, or walking. (also true)

So essentially: `curl | bash` is less secure... than what? And since when?

Back in 1997, yeah, `curl | bash` had some issues. They were solved. A. Long. Time. Ago.

`brew`, `rust`, and many others install via `curl | bash`.

There's nothing inherently dangerous about `bash` or `curl`, _assuming_ that **HTTPS** (with the **S**) is used, and that you _trust_ the _build_ and the _vendor_.

## I don't _trust_ you! You're evil in disguise!

Installing anything **from a vendor that you don't trust** is dangerous - whether it's a `.exe`, `.pkg`, `go run`, or `npm install` or a `curl | bash` makes ZERO technical difference. That's why we've got all of this:

- https://webinstall.dev/about/
- https://webinstall.dev/legal/#privacy
- https://webinstall.dev/faq/#security
- https://github.com/webinstall/packages/blob/master/golang/

(again, how many authors in your `go.sum` or `package-lock.json` are _completely anonymous_ and definitely _NOT_ security experts?)

Also, I'm fairly public in general - podcasts and blogs and companies and such.

I'm no Rob Pike and not quite up there with mholt6 (who I chat with regularly) or indutny (who has fixed TLS bugs in node for me more than once) either, but I'm mostly likely mentioned if you grep your `node_modules` for `coolaj86`, or check the LICENSES of your AppleTV-enabled smart TV, I'm probably in there.

Ryan (co-creator of the project) is a bit more private, but very deliberate about putting his name to something.

The buddy who's been helping with the powershell stuff is literally the guy who analyses the kernel dumps that upload after a Windows 10 BSOD. He's the reason that Node.js works on Windows 10 on ARM (and he'll likely be involved in some of the upcoming Go porting as well).

## Fight me

But if you're going to fight me, do it honestly.

1. Do you even have the problem that this solves?
2. What technical problem do you see? What's my blind spot?

If you can't agree to anything in the "This. Is. Awesome." then either you're not being honest, or we just live in separate worlds. That's okay (the second part).

Can you admit to be doing something that results in a technical issue? If you can point it out, I can find and deploy a solution (and I'd be happy to have your help in doing so!)
## [5][Considering dropping GO386=387](https://www.reddit.com/r/golang/comments/hsnwto/considering_dropping_go386387/)
- url: https://groups.google.com/forum/#!topic/golang-nuts/Gl6bODRX1NE
---

## [6][Lightweight validation library that can export rules as JSON so browsers can apply the same rules](https://www.reddit.com/r/golang/comments/hsrn3o/lightweight_validation_library_that_can_export/)
- url: https://github.com/AgentCosmic/xvalid
---

## [7][Add semver, based on `git describe`, easy as 1, 2, 3 (and less than 150 LoC)](https://www.reddit.com/r/golang/comments/hso4zb/add_semver_based_on_git_describe_easy_as_1_2_3/)
- url: https://www.reddit.com/r/golang/comments/hso4zb/add_semver_based_on_git_describe_easy_as_1_2_3/
---
I created this a while back, but I'm not very social on reddit.
There may be some other tool that does similar that I didn't find before.

However, this is my solution and I think that it's simple enough that most people can appreciate it:

https://git.rootprojects.org/root/go-gitver

In short:

## 0) You already use `git tag` with SEMVER (in the form `vX.Y.Z`)

If you don't already use semver `git tag`s, then this post probably isn't something you care about
(but maybe you should - so go read up on those first).

## 1) You add **3 special variables** to your `main.go` (with optional `go generate`):

These are the fallback if `git` is not installed or unavailable.

```
//go:generate go run -mod=vendor git.rootprojects.org/root/go-gitver --fail

package main

var (
	GitRev       = "0000000"
	GitVersion   = "v0.0.0-pre0+0000000"
	GitTimestamp = "0000-00-00T00:00:00+0000"
)

func main() {
	if (len(os.Args) &gt; 1 &amp;&amp; "version" === os.Args[1]) {
		fmt.Printf("Foobar v%s (%s)", GitVersion, GitTimestamp)
	}
	// ...
}}
```

## 2) Generate `xversion.go`, which overwrites those vars:

```
go generate
```

or 

```
go run git.rootprojects.org/root/go-gitver
```

That will produce `xversion.go`, for example:

```
// Code generated by go generate; DO NOT EDIT.
package main

func init() {
	GitRev = "6dace8255b52e123297a44629bc32c015add310a"
	GitVersion = "v1.1.4-pre2+g6dace82"
	GitTimestamp = "2020-07-16T20:48:15-06:00"
}
```

If you do use `go generate`, you'll want to add `go-gitver` to your `tools/tools.go`
(so that `go mod tidy` can track it in `go.mod`)

```
// +build tools

// This is a dummy package for build tooling
package tools

import (
	_ "git.rootprojects.org/root/go-gitver"
)
```

## 3) Build your app as per usual:

For me that looks like this:

```
go mod tidy
go mod vendor
go generate -mod=vendor ./...
go build -mod=vendor .
```

## *) That's it.

There's some other options in the `README.md`, but I think this should hit the spot for most people.

Upvote if it helps. Downvote if it hurts. Comment with suggestions or questions.
## [8][Let's Go! is an outstanding book.](https://www.reddit.com/r/golang/comments/hsh39f/lets_go_is_an_outstanding_book/)
- url: https://www.reddit.com/r/golang/comments/hsh39f/lets_go_is_an_outstanding_book/
---
I recently finished "Let's Go" by Alex Edwards.  I took my time and coded the application along with the book.  My background is mostly in Rails development with a smaller amount of Go api development.

This book was very educational.  It makes me want to go back and re-develop the Go api application I built for my current job.  Alex does a fantastic job of covering not just the basics of web development using Go, but he also shows you how to properly structure it and how to architect it to make testing easier.  Along the way, he covers about everything you need to know to build a full-fledged web app.

The book doesn't skimp on testing either.  Also, he does this mostly using just Go, with a few small, focused third-party packages thrown in.  And he doesn't just tell you what to do, he fully explains why you would want to do it this way.  I learned a lot about things that are not Go-specific.

This book would work for someone who knows the basics of Go and wants to learn how to build a "real" web app and it also would work for a web developer coming from another language who wants to get up to speed quickly on transferring their skills to Go.
## [9][I start writing a basic framework for multi agent system (mas)](https://www.reddit.com/r/golang/comments/hstc39/i_start_writing_a_basic_framework_for_multi_agent/)
- url: https://github.com/JusticeN/chaos
---

## [10][How does everyone use the standard library efficiently?](https://www.reddit.com/r/golang/comments/hsuhn9/how_does_everyone_use_the_standard_library/)
- url: https://www.reddit.com/r/golang/comments/hsuhn9/how_does_everyone_use_the_standard_library/
---
I’ve been learning go and I’m curious how does everyone go about efficiently implementing the standard libraries ? I’ve found myself looking at the libraries for about 20min and ultimately going to stack overflow as they lack examples or don’t explain its usage very well.

For example, interfaces available to a method are not very well documented
