# golang
## [1][We are the Go Time podcast. Ask us anything. (AMA)](https://www.reddit.com/r/golang/comments/io94yi/we_are_the_go_time_podcast_ask_us_anything_ama/)
- url: https://www.reddit.com/r/golang/comments/io94yi/we_are_the_go_time_podcast_ask_us_anything_ama/
---
Hi everyone! I'm Jon Calhoun, one of the panelists on the Go Time podcast. For those of you unfamiliar - it is a Go podcast that we record live every Tuesday at 3pm ET. We usually have a guest or two on each episode, and discuss a specific topic (defer, testing, databases, infra, etc). You can check it out here: &lt;https://changelog.com/gotime&gt;

This coming episode we want to try something a little different - we want to make a Q&amp;A episode. There are two reasons for doing this:

1. We are hoping this inspires topics for future episodes.
2. We want a venue to discuss questions that don't really fit into an entire episode on their own.

To make this happen we would like everyone here to post any Go-related questions that you would like us to discuss on air, and we will try to get to as many as possible. I'll also try to type up answers here while we discuss them on the episode.

We will be answering questions live tomorrow, Tuesday, Sep 8. We will repeat questions on air, and since we record live you can join in on the Gophers Slack to ask follow-up questions or to elaborate on questions.
## [2][What books are best for learning Golang?](https://www.reddit.com/r/golang/comments/ir8t5l/what_books_are_best_for_learning_golang/)
- url: https://www.reddit.com/r/golang/comments/ir8t5l/what_books_are_best_for_learning_golang/
---
The speed of C++ and the ease of developement like python Ive decided I want to learn Go! What resources can I use to get started. Do you know of any excellent books I can use that are on the same level  of something like Automate the boring stuff or better?
## [3][Parse RSS, Atom and JSON feeds in Go](https://www.reddit.com/r/golang/comments/ir42b5/parse_rss_atom_and_json_feeds_in_go/)
- url: https://golangweekly.com/link/95041/5616be026c
---

## [4][New to microservices](https://www.reddit.com/r/golang/comments/irbifl/new_to_microservices/)
- url: https://www.reddit.com/r/golang/comments/irbifl/new_to_microservices/
---
Hey guys, I'm new to the world of microservices, can you highlight books that would help me develop and can you share some that you have? Thanks
## [5][What’s everyone’s take on ORM or plain SQL?](https://www.reddit.com/r/golang/comments/ircaxk/whats_everyones_take_on_orm_or_plain_sql/)
- url: https://www.reddit.com/r/golang/comments/ircaxk/whats_everyones_take_on_orm_or_plain_sql/
---
I’ve started learning go and after I built my CRUD application I realized that there was a neat ORM library called GORM.

Which method do you prefer and why?
## [6][LRU cache code generator](https://www.reddit.com/r/golang/comments/ir9t71/lru_cache_code_generator/)
- url: https://github.com/maxim2266/go-LRU-cache
---

## [7][Project introduces breaking changes to library - bumps minor versioning instead of major version because go module's v2+ handling is so quirky](https://www.reddit.com/r/golang/comments/iqoiok/project_introduces_breaking_changes_to_library/)
- url: https://www.reddit.com/r/golang/comments/iqoiok/project_introduces_breaking_changes_to_library/
---
I'm very unhappy with go modules and this project shying away from bumping the major version in favor of bumping the minor version while introducing huge breaking changes just speaks for itself.

[https://github.com/gofiber/fiber/issues/736#issuecomment-690750255](https://github.com/gofiber/fiber/issues/736#issuecomment-690750255)

What do you think?

I just can't stop thinking that in other ecosystems people are not scared to stick to semantic versioning, so this might be a "smell" that Go has, that is not present in other languages.

Oh and please please *don't you all jump into that Github issue and tell the maintainers how they should run their project*. I don't want to cause them any trouble by posting here!
## [8][Small tool to update DNS record on Namecheap web hosting provider](https://www.reddit.com/r/golang/comments/ir7k0b/small_tool_to_update_dns_record_on_namecheap_web/)
- url: https://www.reddit.com/r/golang/comments/ir7k0b/small_tool_to_update_dns_record_on_namecheap_web/
---
[https://github.com/Kulak/namecheap-ddns](https://github.com/Kulak/namecheap-ddns)

100% GO language.
## [9][Using Gonum Graphs to Solve Word Ladder Puzzles | Gonum](https://www.reddit.com/r/golang/comments/iqt2iu/using_gonum_graphs_to_solve_word_ladder_puzzles/)
- url: https://www.gonum.org/post/word_ladder/
---

## [10][Final Year Project](https://www.reddit.com/r/golang/comments/iqt4gr/final_year_project/)
- url: https://www.reddit.com/r/golang/comments/iqt4gr/final_year_project/
---
Let me start off with the fact that I'm not such a smart man .I've slogged through each semester and found that I'm not fascinated about anything particular until I started learning Go. Though I've fallen in love with it , I just know the basic stuff.

The thing is all my classmates have teamed up for the final year project and left me to fend for myself .I'm really lost as to how I'll manage this situation.  I really need your help , I beg you, to please give me  an idea for a project that can span for a period of 2-3 months .

I really appreciate Go and this community  for making me realize how much passionate I am about this language.
## [11][io.Writer mutex question](https://www.reddit.com/r/golang/comments/iqt34v/iowriter_mutex_question/)
- url: https://www.reddit.com/r/golang/comments/iqt34v/iowriter_mutex_question/
---
    var wg sync.WaitGroup
    for i := 0; i &lt; 200000; i++ {
    	go func(id int) {
    		wg.Add(1)
    		os.Stdout.Write(append(bytes.Repeat([]byte("01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29"), 500), byte('\n')))
    		wg.Done()
    	}(i)
    }
    wg.Wait()

I'm trying to check cases without mutex, but got something unexpected.. This maybe due to I'm new to Go..

    ~$ ./myProg &gt; out.txt

This creates about 8GB text file. My expectation was lines will be messed up as it will print large enough in a goroutine. Below are what I expected from at least some line in the 8GB output. (messed up order)

    01 02 01 03 02 04 ...

However, when I run it, I didn't see any of those cases.. This was same even when I output to `*os.File` instead of `os.Stdout`. Is this normal?? I always thought I had to use mutex, but if this is true, probably I don't have to on any io.Writer?

&amp;#x200B;

# SOLVED: 0xjnml answered, "POSIX mandates atomic writes."
