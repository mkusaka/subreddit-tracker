# golang
## [1][Some golang spotted in a movie I'm about to watch - Guns Akimbo](https://www.reddit.com/r/golang/comments/fbeb9t/some_golang_spotted_in_a_movie_im_about_to_watch/)
- url: https://i.redd.it/2fz3icbkvvj41.jpg
---

## [2][Seeking testers for gotop](https://www.reddit.com/r/golang/comments/fbtyl0/seeking_testers_for_gotop/)
- url: https://www.reddit.com/r/golang/comments/fbtyl0/seeking_testers_for_gotop/
---
I'm looking for some volunteers to help testing the next release of gotop.

The original author of gotop decided to work on another project and flagged gotop as "unmaintained."  I've been working with him to assume maintainership, and in the past month have released a number of versions that merge long-outstanding pull requests and add new features and bug fixes.  I have a new major version ready to go out, but as I'm starting to get distribution packagers on board with the switch, testing becomes more of a concern for me.  I'm appealing to the community for some light assistance in smoke testing releases for other architectures and OSes that I'm unable to do myself, and for assistance with distribution packages.

- OS/Arch testing. I need testers for FreeBSD, Darwin, and Windows, and for architectures other than AMD64.
- Homebrew testing.  I'd appreciate help testing the Homebrew tap, and the instructions that are in the README.
- Device testing.  One of the features being added is instrumentation support for GPUs.  NVIDIA support is pending, and AMD and Intel are in the queue.
- Distribution packages.  The AUR packages have already been transferred, but I can only test that the packages work for my CPU architecture. I'm both unfamiliar and unsure how to start with Debian and Redhat derivatives; do they have AUR analogues?  What about Slackware and Alpine?
- Performance testing. Mostly, I'd like feedback about behaviors and would especially like to have folks out there watching for changes that negatively impact.  I think I'm pretty good about this, but I can only really watch how it performs on my machine.

Also, if any tuning gurus are willing to help with profiling, I'd love to get some tickets in about areas to investigate.  I'm capable of doing this myself, but will probably not get to it for a while as I continue to ramp up on the maintainer aspects.

Finally, I'm always happy to look at pull requests, and will try to prioritize those.

I'm getting some help already from a few folks who have been active in helping me with the transition, but I still have gaps (no ARM testers, for instance, and FreeBSD is a big worry).

To smooth the transition, I've [kept the project in github](https://github.com/xxxserxxx/gotop).

In addition to bug fixes, the features released (and currently tested) are:

- Instrumentation data export via HTTP.
- A battery gauge
- Support for GPUs
- Customizable widget layout via config file
- Error log rotation
- Config file support
- Support for multiple network interfaces (incl. filtering)
- Searching in the proc widget

Once again, the project URL is https://github.com/xxxserxxx/gotop .  An call for testers is in the issue tracker; if you're willing to help, please post to that ticket.  I'm only seeking sanity checks for things I can't test myself, so if someone has already volunteered to, e.g., test FreeBSD, don't feel obliged to also volunteer.

Thanks for reading the long post.
## [3][Go Security Policy](https://www.reddit.com/r/golang/comments/fbtk68/go_security_policy/)
- url: https://golang.org/security
---

## [4][How to create an overlayfs mount in golang?](https://www.reddit.com/r/golang/comments/fbrutx/how_to_create_an_overlayfs_mount_in_golang/)
- url: https://www.reddit.com/r/golang/comments/fbrutx/how_to_create_an_overlayfs_mount_in_golang/
---
I'm new to go. I need to create an overlay mount. What is the correct way to do this? Say, the lower dir is /A ( read only) and upper dir is /B (read write).
## [5][Early Impressions of Go from a Rust Programmer](https://www.reddit.com/r/golang/comments/fb7n67/early_impressions_of_go_from_a_rust_programmer/)
- url: https://pingcap.com/blog/early-impressions-of-go-from-a-rust-programmer/
---

## [6][I am confused](https://www.reddit.com/r/golang/comments/fbqs84/i_am_confused/)
- url: https://www.reddit.com/r/golang/comments/fbqs84/i_am_confused/
---
Hello community,
If I want to deploy a php webapp, I usually need any kind of webserver. I am a little bit confused, how to do this in golang. Afak there is no need for any webserver, because its a kind of built in, but how does the server manage the Go app? Has the application to be registered / started as a service?

Kind regards
## [7][Can someone explain me what bufio is &amp; does?](https://www.reddit.com/r/golang/comments/fbqqa7/can_someone_explain_me_what_bufio_is_does/)
- url: https://www.reddit.com/r/golang/comments/fbqqa7/can_someone_explain_me_what_bufio_is_does/
---
New to go lang and I went through [golangbot.com](https://golangbot.com) and was able to get it's topics. I can get to know what bufio does.
## [8][tarry: A simple command line tool written in Go for waiting until a specific time](https://www.reddit.com/r/golang/comments/fbe0qk/tarry_a_simple_command_line_tool_written_in_go/)
- url: https://www.reddit.com/r/golang/comments/fbe0qk/tarry_a_simple_command_line_tool_written_in_go/
---
[https://github.com/metaphyze/tarry](https://github.com/metaphyze/tarry)

A simple command line tool for waiting until a specific time. This is not the same as "sleep" which will wait for a duration of time. This is useful if you want to execute something at a specific time or more likely execute several things at exactly the same time such as testing if a server can handle multiple *very* simultaneous requests. You could use it like this with "&amp;&amp;" on Linux, Mac, or Windows:

      tarry -until=16:03:04 &amp;&amp; someOtherCommand 

This would wait until 4:03:04 PM and then execute someOtherCommand. Here's a Linux/Mac example of how to run multiple requests all scheduled to start at the same time:

      for request in 1 2 3 4 5 6 7 8 9 10    
      do        
           tarry -until=16:03:04 &amp;&amp; date &gt; results.$request &amp;    
      done
## [9][List of Go Resources for Anybody Who Wants to Learn Golang](https://www.reddit.com/r/golang/comments/fbds4f/list_of_go_resources_for_anybody_who_wants_to/)
- url: https://www.reddit.com/r/golang/comments/fbds4f/list_of_go_resources_for_anybody_who_wants_to/
---
Hello everyone, I have curated a collection of links to different resources for anybody who wants to learn Golang. I hope it will be helpful. 

I sorted them by their level so it has something for people from each level but I know there are some veterans out there and I am not sure about the relevancy of advanced resources. If you have any suggestions, as a resource or feedback, please leave a comment below. I would be happy to add new links to this list.

Here's the List;

[**Golang Module| Jooseph**](https://www.jooseph.com/go-lang)

This list was curated for our platform Jooseph which is basically playlists for learning. You can curate and follow collection of links to different resources for learning a subject. You can also share yours :)

Thanks in advance
## [10][I Created a Coronavirus Tracker with Golang to help you stay updated](https://www.reddit.com/r/golang/comments/fbacj2/i_created_a_coronavirus_tracker_with_golang_to/)
- url: https://www.reddit.com/r/golang/comments/fbacj2/i_created_a_coronavirus_tracker_with_golang_to/
---
I'm in China and all of a sudden the coronavirus hit, things just went down. I couldn't go out, quarantine was applied, every store was closed, the city turned lifeless, I'm stuck in my apartment 24/7.

But I'm a programmer &amp; I work on my project every single day, so I thought why not create a coronavirus tracker with and make a video about it, since I'm making videos anyway?

So I did, and here's the site I've built Golang: [http://pandemicalert.xyz/](http://pandemicalert.xyz/)

And here's the video I've made about it: [https://youtu.be/vC52Cm8hHpI](https://youtu.be/vC52Cm8hHpI)

I'm making more videos about programming &amp; startups, and I write all my backend with Golang, please subscribe if you guys are interested. 😉

(admin: please remove if violates policies, this is a self promotion but I don't think it's a shameless self-promo, trying to provide value here, thanks).
