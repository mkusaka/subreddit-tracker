# golang
## [1][[Q&amp;A] io/fs draft design](https://www.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/)
- url: https://www.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/
---
I posted a draft design today for new file system interfaces for Go.

Video: https://golang.org/s/draft-iofs-video

Design: https://golang.org/s/draft-iofs-design

Let's do the Q&amp;A about the design here in Reddit. My hope is that the threading support will help keep questions and answers matched.

Please start a new top-level comment for each new question.

See also the related [Q&amp;A for the //go:embed draft design](https://golang.org/s/draft-embed-reddit).
## [2][Design Draft: First Class Fuzzing](https://www.reddit.com/r/golang/comments/hvpr96/design_draft_first_class_fuzzing/)
- url: https://go.googlesource.com/proposal/+/refs/heads/master/design/40307-fuzzing.md
---

## [3][Bill Kennedy on value vs pointer semantics](https://www.reddit.com/r/golang/comments/i1dot5/bill_kennedy_on_value_vs_pointer_semantics/)
- url: https://youtu.be/i5nyPaAwM3s
---

## [4][Go: How to Reduce Lock Contention with the Atomic Package](https://www.reddit.com/r/golang/comments/i1ryn9/go_how_to_reduce_lock_contention_with_the_atomic/)
- url: https://medium.com/a-journey-with-go/go-how-to-reduce-lock-contention-with-the-atomic-package-ba3b2664b549
---

## [5][GLab: A Gitlab CLI tool](https://www.reddit.com/r/golang/comments/i1lnqo/glab_a_gitlab_cli_tool/)
- url: https://www.reddit.com/r/golang/comments/i1lnqo/glab_a_gitlab_cli_tool/
---
Glab is a simple and elegant GitLab CLI tool written in Go (golang) for creating and managing issues, merge requests, and a lot more... Do [give it a try](https://github.com/profclems/glab). Your feedback is much appreciated.

 [https://github.com/profclems/glab](https://github.com/profclems/glab) 

&amp;#x200B;

https://preview.redd.it/cfcjpttwabe51.png?width=1024&amp;format=png&amp;auto=webp&amp;s=99d159203ce82dbddae57f732becea64b64c00d2
## [6][Kyubin: a simple queue based bin in pure go + storm db](https://www.reddit.com/r/golang/comments/i1q4lv/kyubin_a_simple_queue_based_bin_in_pure_go_storm/)
- url: https://stackb.in
---

## [7][hashicorp nomad](https://www.reddit.com/r/golang/comments/i16vm2/hashicorp_nomad/)
- url: https://www.reddit.com/r/golang/comments/i16vm2/hashicorp_nomad/
---
Does anyone here use Nomad? I like the fact its written in Go and very easy to deploy. I use it for Legacy Applications (C++ and Java). I don't need anything fancy like kubernetes. So far I am impressed with it.. Any other good alternatives (hopefully written in Go)?
## [8][GoFlow is a Golang based high performance, scalable and distributed workflow framework](https://www.reddit.com/r/golang/comments/i1a2y0/goflow_is_a_golang_based_high_performance/)
- url: https://www.reddit.com/r/golang/comments/i1a2y0/goflow_is_a_golang_based_high_performance/
---
It allows to programmatically author distributed workflows as Directed Acyclic Graphs (DAGs) of tasks (nodes). Goflow executes your tasks on an array of goflow workers by uniformly distributing the load

LINK: [https://github.com/faasflow/goflow](https://github.com/faasflow/goflow)

[give it a look](https://preview.redd.it/xqt3f3cut7e51.jpg?width=1007&amp;format=pjpg&amp;auto=webp&amp;s=4f5a9b68363caf2c037fdb46d660a1113b26f479)
## [9][Comparing Crystal’s Concurrency with that of Go’s (Part I)](https://www.reddit.com/r/golang/comments/i1j287/comparing_crystals_concurrency_with_that_of_gos/)
- url: https://link.medium.com/782maOnbA8
---

## [10][I just landed a massive refactor of my open source project Super Graph over 3K lines changed. Ask me anything.](https://www.reddit.com/r/golang/comments/i1mql8/i_just_landed_a_massive_refactor_of_my_open/)
- url: https://www.reddit.com/r/golang/comments/i1mql8/i_just_landed_a_massive_refactor_of_my_open/
---
Super Graph is a GraphQL to SQL compiler in Go that I've been working on for a while now. Overtime it got a lot of useful features and as the single core contributor some parts of the codebase got away from me. A couple weeks ago I came across a bug in the part of code dealing with nested inserts and updated where we generate SQL to update or insert multiple related tables with a single SQL query.  The fix for this was simple enough but somehow I couldn't get myself to make it since it felt like I'm forcing code where it didn't belong. When you work on the same codebase for a long while you naturally develop a fine tuned sense about it and code smell is very evident to you and I felt this was not just a smell it stunk. 

I didn't make the fix instead I shut my laptop and went on a long walk as it helps me think. Two packages `psql` and `qcode` made up the core of the compiler, `qcode` took the GraphQL and parsed it into an AST called `QCode` which was then passed to `psql` to generate SQL from. The issue I found was that `QCode` was a not rich enough, it didn't carry with it information about database tables and relationships. This meant that `psql` had to do a lot more work and instead of just focusing on generating SQL it had to do the additional work of pulling in database schema information and this made the code more complicated.

The one thing I knew from experience was that generating the SQL was hard enough why make it do more. I decided it was time for a refactor and this required a design rethink. The new design I felt needed to focus on improving readability and simplify complex functions. The best way to simplify code in Go is to extract out concerns into new packages and add more information to structs that are passed around so you don't have to pull in this information all over the place. 

In the new design the core GraphQL parser was extracted out from `qcode` into it's own package. Next up the database schema and relationship discovery code was extracted out from `psql` into it's own package and finally the `QCode` AST was made richer. The final result was a codebase much nicer and easier to work with and tons of bugs fixing themselves. Major new features going ahead like support for MySQL seems much easier to reason about.

Refactoring Go code is a breeze the fast compiler makes running tests while making changes almost realtime. I'm just happy with how quickly I could do this and the higher quality codebase that it resulted in. Ask me anything.

https://github.com/dosco/super-graph

The Commit:
https://github.com/dosco/super-graph/commit/992c78ee947e6497a0e7b03059067e18fa9a22e2
## [11][On choosing a buffer size](https://www.reddit.com/r/golang/comments/i1cro6/on_choosing_a_buffer_size/)
- url: https://www.reddit.com/r/golang/comments/i1cro6/on_choosing_a_buffer_size/
---
This has me stumped.

As a learning experience, I wrote an implementation of the cat command and tested it by using another program I wrote to write random bytes to stdout.  In developing both programs though, I noticed that a buffer size of 4096 offered less performance than 65536, which on my machine seems like a sweet spot since anything above that further degrades throughput.  My question is, what factors might be influencing this?  I thought 4096 was a common buffer size for file and pipe I/O, even 8192 at times (though that was even worse in my tests).  I don't know if it's related, but the page size on my system is 4096.


The random byte generator is called 'fon', and the testing command chain is as follows:

fon -s 65536 | cat | pv -a &gt; /dev/null


So 'fon' is writing 64GiB of nonsense (no, that's not a typo) into cat, where 'pv -a' gets the average throughput from cat to /dev/null.  Testing against FreeBSD's cat, I got 425MiB/s, whereas my cat would get identical speed with a buffer of 4096, but switching to 65536 brings it up to 572MiB/s.  Likewise, 'fon' performs best at that buffer size with about 600MiB/s average output speed, but a bigger or smaller buffer reduces the speed.  

Just for fun, I even wrote a version of 'yes' and compared it to FreeBSD's version.  Mine peaked at 6.7GiB/s average whereas FreeBSD peaked at about 4.1GiB/s.  Guess what buffer size I used?


So to sum it up, does anyone have an explanation for why 64KiB buffers are outperforming 4KiB, 8KiB, even 128KiB or 1MiB buffers?  I know very little about the inner workings of this sort of thing, and I don't want to fall into a trap of assuming all machines will behave the same way (hence determining buffer size at runtime).  If it helps, I'm running FreeBSD 12.1 Release-p6 on a Thinkpad T430 with 8GB RAM and an Intel i5-2520M, Go version 1.14.6.


UPDATE

Having posted the CPU I'm using, it got me thinking that maybe the CPU cache sizes had something to do with this.  Maybe it's just the universe laughing at me, but the L1 cache size is 64KiB instruction and 64KiB data.  Seeing how fon is writing 64KiB at a time using read(buff) from math/rand, I can't help but wonder if this is meaningless or if I'm onto something.
## [12][[PROJECT] Larder CLI - Bookmarking for Developers](https://www.reddit.com/r/golang/comments/i1bfap/project_larder_cli_bookmarking_for_developers/)
- url: https://www.reddit.com/r/golang/comments/i1bfap/project_larder_cli_bookmarking_for_developers/
---
TLDR; Larder is an extension for bookmarking things on the web you'll need again. I've written a [Larder command line interface](https://github.com/theycallmemac/larder) in Go for all us people who don't like browser extensions :D

I worked on this as relief from my final year project in college (consulted this  subreddit a lot during that time!) I decided to make a robust, easy to use command line interface for Larder. I did this mainly because I have a 2012 Macbook Pro with 4GB of RAM that really cannot handle any more browser extensions. I also did this because I know there’s a lot of folk who live in their terminals and detest browser extensions.

[Larder.io](https://larder.io) is a bookmarking tool for developers, making  the process of organizing bookmarks by categorizing and tagging them. Categorizing is done by placing bookmarks in an organized set of folders, allow you to annotate each bookmark with specific search terms. Later on when you need it, you can search your bookmarks by tags you’ve used, it’s title, or it’s original URL.

This tool also supports synchronizing your GitHub stars too. So, you can track updates from projects that you starred on GitHub. If you want to bookmark something automatically, Larder extension is available for browsers, Android, and there’s also an API.

What really makes this new Larder CLI stand out is the ability to refresh your access token. Tokens expire in a month and can be refreshed for a new access token at any time, invalidating the original access and refresh tokens. This built in functionality does not exist in alternatives to this Larder CLI. This process is automated so the user can continuously use the Larder CLI without having to manually remove and add new access tokens. Who doesn’t love a little bit of automation :D

You can check out Larder [on Github](https://github.com/theycallmemac/larder).
