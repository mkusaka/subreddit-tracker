# golang
## [1][[2005.11710] Featherweight Go](https://www.reddit.com/r/golang/comments/gtxflh/200511710_featherweight_go/)
- url: https://arxiv.org/abs/2005.11710
---

## [2][Go RabbitMQ Beginners Tutorial](https://www.reddit.com/r/golang/comments/gtx8d6/go_rabbitmq_beginners_tutorial/)
- url: https://youtu.be/pAXp6o-zWS4
---

## [3][Thank you to this community](https://www.reddit.com/r/golang/comments/gtjkss/thank_you_to_this_community/)
- url: https://www.reddit.com/r/golang/comments/gtjkss/thank_you_to_this_community/
---
I haven't posted here before, but I've been lurking for over a year.  Sorry that my first post isn't going to be technical, but I want to share my story.  Mods, sorry if this isn't what you want on this sub.

I got started with programming pretty late in life compared to most people by going to a web dev bootcamp at 26.  Yes, I know the reputation some of those have, but it worked for me.  I learned Java, and have been working with that professionally for the last 4 years.  I discovered Go around a year and a half ago, and fell in love with it fairly quickly.  I never had an opportunity to use it professionally though, just home projects and tinkering.

I lost my last job in March, right before my city went on lockdown.  It made job hunting pretty difficult, with no dev meetups or socializing.  I don't live in a particularly tech heavy city.  Every job site I looked on returned between 0 and 1 Go job.  I knew I wanted a chance to work with Go professionally, but I wasn't really holding my breath for it.  Still made sure to have it on my resume and job hunt profiles, of course, but I honestly looked at it with the same odds one would play a scratch off lottery ticket.  I had pretty much resigned myself to taking another Java job with one of my city's mega-corporations, pretty much the only people hiring at the moment.

Then, a week and a half ago, a recruiter reached out to me because I had Go listed on my linkedin profile.  She wanted me to do an assessment test, then have an interview for a Go position.  The interview was probably the toughest grilling I've ever received in a professional setting before, and I can understand why.  Comparatively little experience with programming in general, and no professional experience with the language I was aiming for.  I'm lucky it was a video interview from home, because I was sweating bullets by the end.

Apparently, I aced it.  They made me an offer same day.  I'm going to be starting my first professional Go position in two weeks, and I want to thank this community for that.  The posts here are educational, insightful, and challenging.  Keeping up with the more experienced devs who post to this sub has forced me to learn so much, and think in ways I haven't had to before, and I credit that for my success.  I couldn't have done this without you all.

For anyone currently struggling due to covid and the events of the world, hang in there.  The world will stabilize, and though it may not look the same after as it did before, knowledge you gain will carry through.  For those who want to use Go professionally, but are having trouble getting it introduced at your current job or can't find a new employer who is already using it, keep pushing.  Go was build by devs, for devs, and it will be up to us to push for more adoption, but we're getting there.  We see posts all the time on this sub about Go's growing popularity, and that is being driven by us.  From a late starter who is the walking embodiment of imposter syndrome, I promise you can do it.
## [4][HTTP Adaptive Streaming video player written in golang](https://www.reddit.com/r/golang/comments/gtxovr/http_adaptive_streaming_video_player_written_in/)
- url: https://www.reddit.com/r/golang/comments/gtxovr/http_adaptive_streaming_video_player_written_in/
---
We've created a headless video player and a testbed framework for the evaluation of HTTP Adaptive Streaming (HAS) algorithms, protocols and so much more...

We've had a lot of fun writing this HAS video player, over the last year.  Hope you enjoy it

Player - [godash](https://github.com/uccmisl/godash)
Testbed- [godashbed](https://github.com/uccmisl/godashbed)
## [5][lithdew/casso: A fast Go implementation of the Cassowary constraint solver for laying out UIs.](https://www.reddit.com/r/golang/comments/gtg6oz/lithdewcasso_a_fast_go_implementation_of_the/)
- url: https://github.com/lithdew/casso
---

## [6][Rethinking of CGI as a selfhosted lambda server](https://www.reddit.com/r/golang/comments/gtwrkb/rethinking_of_cgi_as_a_selfhosted_lambda_server/)
- url: https://trusted-cgi.reddec.net/
---

## [7][SPIL - small functional programming language written in Go](https://www.reddit.com/r/golang/comments/gtyzle/spil_small_functional_programming_language/)
- url: https://www.reddit.com/r/golang/comments/gtyzle/spil_small_functional_programming_language/
---
During the self-isolation I wrote small functional programing language using Go.

It has LISP-like syntax and some features like:

\- lazy evaluation (lazy lists)

\- tail call optimization

\- optional function memoization

\- static type checking

\- generic functions and generic type support (work in progress).

You can find some documentation here: [avoronkov.github.io/spil-docs](https://avoronkov.github.io/spil-docs/) and sources here: [github.com/avoronkov/spil](https://github.com/avoronkov/spil/) .

P.S. a bit of SPIL code:

    (use std)
    
    ; check if list l contains element x
    (def contains (x:a '()) :bool 'F) 
    (def contains (x:a l:list[a]) :bool
         (if (= (head l) x)
           'T  
           (contains x (tail l))))
    
    ; test if number is prime:
    (def prime? (1) 'F) 
    (def prime? (n) (prime? n 2)) 
    (def prime? (n i)
        (if (&gt; (* i i) n)
          'T  
          (if (= (mod n i) 0)
            'F  
            (prime? n (inc i)))))
    
    ; lazy list of positive integers:
    (set ints (gen inc 0) :list[int]) 
    
    ; check if 17 is among first 20 prime numbers:
    (print (contains 17 (take 20 (filter prime? ints)))) 
    ; output: true
## [8][funcv is a new super awesome command-line parser for Go](https://www.reddit.com/r/golang/comments/gtxcrx/funcv_is_a_new_super_awesome_commandline_parser/)
- url: https://github.com/moshenahmias/funcv
---

## [9][Writing an API using Golang](https://www.reddit.com/r/golang/comments/gt84cm/writing_an_api_using_golang/)
- url: https://www.reddit.com/r/golang/comments/gt84cm/writing_an_api_using_golang/
---
I am extremely new to Golang and I have been using a Udemy course to learn about the language. I am starting a new job as a back-end developer and my first task is to create an API using Golang. If anyone has any good guides/youtubes to throw my way that would be great. Or if anyone has any pointers or best practices please give me all the knowledge. Thanks!
## [10][Mainflux 0.11 — Digital Twin, MQTT Proxy And More](https://www.reddit.com/r/golang/comments/gti1tc/mainflux_011_digital_twin_mqtt_proxy_and_more/)
- url: https://medium.com/mainflux-iot-platform/mainflux-0-11-digital-twin-mqtt-proxy-and-more-46bde98635fe
---

