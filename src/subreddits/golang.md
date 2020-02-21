# golang
## [1][I have published Goxygen - a tool that helps generate full stack projects in Go and React. I would love to hear your feedback](https://www.reddit.com/r/golang/comments/f7aco4/i_have_published_goxygen_a_tool_that_helps/)
- url: https://github.com/Shpota/goxygen
---

## [2][What does this mean in Go?](https://www.reddit.com/r/golang/comments/f727uy/what_does_this_mean_in_go/)
- url: https://www.reddit.com/r/golang/comments/f727uy/what_does_this_mean_in_go/
---
I've been perusing the standard lib and found in the `mutex.go` source code this syntax. Can someone explain to me what the IF statement is saying? Particularly the &amp; and () with the types in there?

    const (
    	mutexLocked = 1 &lt;&lt; iota // mutex is locked
    	mutexWoken
    	mutexStarving
    	mutexWaiterShift = iota
    )
    
    ...
    
    if old&amp;(mutexLocked|mutexStarving) == mutexLocked {
        ...
    }
## [3][SAML 2.0 SSO in Go](https://www.reddit.com/r/golang/comments/f723yk/saml_20_sso_in_go/)
- url: https://www.reddit.com/r/golang/comments/f723yk/saml_20_sso_in_go/
---
I need to implement SAML 2.0 SSO in our go backend. I've only used OAUTH2 and done it for each provider (Google, FB, Linkedin, etc), and not generically like this before. I'm having a bit of trouble understanding the million acronyms and terms used.  I'm trying to pretend to be an IdP somewhere so I can write the code for us to act as a service provider.

I see this library: [https://github.com/crewjam/saml](https://github.com/crewjam/saml)

And I set myself up to be an "identity provider" in Onelogin, as admin of my own account. I set up the OneLogin SAML Test (IdP). I get a screen with information like:

    ISSUER_URL  
    https://app.onelogin.com/saml/metadata/b238b8ae-1657-4f1c-893c-c0be41c8fcc2 
    SAML 2.0 Endpoint (HTTP) 
    https://mysubdomain.onelogin.com/trust/saml2/http-post/sso/b238b8ae-1657-4f1c-893c-c0be41c8fcc2 
    SLO Endpoint (HTTP) 
    https://quoter.onelogin.com/trust/saml2/http-redirect/slo/1082688

&amp;#x200B;

And on another "Configuration" screen where it wants me to provide info : SAML Consumer URL, ACS URL Validator, etc.

On another screen I get a X509 cert and a fingerprint for that.

So far I gather I have to take a POST to my app, acting as a SP, this info provided by an Identity Provider, so I can persist it to the database for that account.

      {
    	"issuer_url": "https://app.onelogin.com/saml/metadata/b238b8ae-1657-4f1c-893c-c0be41c8fcc2",
    	"login_endpoint_url": "https://mysubdomain.onelogin.com/trust/saml2/http-post/sso/b238b8ae-1657-4f1c-893c-c0be41c8fcc2",
    	"logout_endpoint_url": "https://mysubdomain.onelogin.com/trust/saml2/http-redirect/slo/1082688",
    	"fingerprint": "97:D8:43:05:blah:blah",
    	"certificate":"theCertificateBase64Encoded"	
    }

So I have a route that does that.

&amp;#x200B;

Now: whats my next step, using a library like the crewjam one specified above? What values need to go in to the "SAML Consumer URL, ACS URL Validator" fields? What routes do I need to make to handle these? What url is the \`idp metadata url\` wanted by the crewjam/saml library? If I had to take a guess it would be the \`issuer\_url\` since it has metadata in the link? When do I use it?

Since OneLogin only provides me the x509 cert in .pem, where do I get the private key needed by the crewjam/saml library?

So far I can see:

account owner posts to route auth/saml in my app with their account info request data -&gt;  

i save it to the db -&gt;  

a user under that account posts to login/saml -&gt;

i do a lookup on the account, yes saml\_sso is enabled, redirect them to the login\_endpoint\_url -&gt;

they should enter their username/pw there im assuming and get redirected back to ???? (is this the consumer URL i want them hit? its really a callback url?) -&gt;

i validate they got logged in (using the certificate?) -&gt;

i issue a JWT -&gt;

they are logged in.

I don't really understand the [samltest.id](https://samltest.id) stuff in the crewjam library since it seems to operate so differently to how an actual idp would operate (as far as I can tell).
## [4][Standard library appreciation post](https://www.reddit.com/r/golang/comments/f7alwj/standard_library_appreciation_post/)
- url: https://www.reddit.com/r/golang/comments/f7alwj/standard_library_appreciation_post/
---
Hi all,

I've been deep-diving in Rust lately after working almost exclusively in Go for several years because learning new things is good, and I just needed to resurface and leave a note of appreciation for just how comprehensive and powerful the Go Standard Library is. 

Don't get me wrong, for what Rust is targeted at, it's very good, but I've spent WAY too much time searching standard library documentation for functionality that doesn't exist there.

Just one piece of the frictionless developer experience Go aims for.
## [5][Using Golang (Or Any Language) in Alexa Skills](https://www.reddit.com/r/golang/comments/f79xgw/using_golang_or_any_language_in_alexa_skills/)
- url: https://phillipsoft.com/alexa/using-golang-in-alexa-skills.html
---

## [6][gqlgen with the clean architecture](https://www.reddit.com/r/golang/comments/f7594r/gqlgen_with_the_clean_architecture/)
- url: https://www.reddit.com/r/golang/comments/f7594r/gqlgen_with_the_clean_architecture/
---
I am trying to figure out what project structure I should use with gqlgen v0.11 alongside a clean architecture. Should I just keep my gqlgen related files under the graph directory and have my use cases and repository files elsewhere? What are some of your suggestions?
## [7][gokita - A go-kit(https://gokit.io) boiler plate](https://www.reddit.com/r/golang/comments/f79cla/gokita_a_gokithttpsgokitio_boiler_plate/)
- url: https://github.com/SeamPay/gokita
---

## [8][How good are sites like goreportcard.com at actually finding issues in your code?](https://www.reddit.com/r/golang/comments/f703o6/how_good_are_sites_like_goreportcardcom_at/)
- url: https://www.reddit.com/r/golang/comments/f703o6/how_good_are_sites_like_goreportcardcom_at/
---
I've recently started learning go and I've written a discord bot that lets users chose their roles by adding reactions to a message. Just for fun I've tried looking at the review on  [goreportcard.com](https://goreportcard.com) for my code, because I've seen it's reports in some repos and it reported A+ which is just way to good for my code. 

So how reliable are those sites at finding any issues in code and are there better tools?

[Github Repo](https://github.com/deinernstjetzt/rolebot) for reference
## [9][GitHub - twitchtv/twirp: A simple RPC framework with protobuf service definitions](https://www.reddit.com/r/golang/comments/f6rpz3/github_twitchtvtwirp_a_simple_rpc_framework_with/)
- url: https://github.com/twitchtv/twirp
---

## [10][Hugo v0.65.0 released: Draft, expire, resource bundling, and fine grained publishing control for any page](https://www.reddit.com/r/golang/comments/f6tilj/hugo_v0650_released_draft_expire_resource/)
- url: https://gohugo.io/news/0.65.0-relnotes/
---

