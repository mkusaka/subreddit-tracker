# golang
## [1][Can gRPC be completely used in place of websockets?](https://www.reddit.com/r/golang/comments/j7wjlf/can_grpc_be_completely_used_in_place_of_websockets/)
- url: https://www.reddit.com/r/golang/comments/j7wjlf/can_grpc_be_completely_used_in_place_of_websockets/
---
I'm coming from nodejs and after implementing API in gRPC I found its more fun in building APIs with gRPC. I already knew golang for a while and was switching to golang. I am wondering if I want to implement a chat application or a server monitoring app with grpc instead of websockets, is it going to be a good idea?
## [2][Type-safe database client for Postgres, MySQL &amp; SQLite](https://www.reddit.com/r/golang/comments/j7fs49/typesafe_database_client_for_postgres_mysql_sqlite/)
- url: https://github.com/prisma/prisma-client-go
---

## [3][Why does a missing argument (or type mismatch) in fmt.Printf not result in a compiler error, but rather a runtime replacement?](https://www.reddit.com/r/golang/comments/j7ryfd/why_does_a_missing_argument_or_type_mismatch_in/)
- url: https://www.reddit.com/r/golang/comments/j7ryfd/why_does_a_missing_argument_or_type_mismatch_in/
---
I've written some non-trivial go code in the past, so I'm surprised this is the first time I ran into this behavior.

I can write a line of code as follows:

    fmt.Printf("hello %s world %s", "bright")

The output is: 

    hello bright world %!s(MISSING)

What was the language designers' rationale for this not being a compiler error instead? Is that theoretically impossible or was this a deliberate design decision?
## [4][variable naming conventions in go](https://www.reddit.com/r/golang/comments/j7o96q/variable_naming_conventions_in_go/)
- url: https://www.reddit.com/r/golang/comments/j7o96q/variable_naming_conventions_in_go/
---
I am mystified by the go community's adherence to terse naming conventions for variables, and I was wondering what reasonable opinions there are for this. I find full words, or combinations of words, to be the most readable.

`s := Server{}`

instead of

`server := Server()`

The go community seems to draw its naming conventions from the go designers themselves, although there could be other sources of inspiration.

On the one hand, the go libraries use long readable names, such as net.Dial, and x509.CreateCertificate, but on the other hand implementation code uses single letter variable names. The dichotomy suggests that public API's need to be readable, but implementation code does not need to be readable. I find this bizarre because when I am debugging code, whether I have written it or not, the last thing I need is to be stumped by the meaning behind a variable name. I have encountered many defenses of using terse variable names such as

1. "Its obvious what the variable is for / what it means"
2.  "I could rename it to 'server' but you wouldn't know what kind of server. It could be an http server or an ssh server, so clearly 'server' is not good enough either"
3. "I don't want to name my variables such as 'thisIsAnIntThatWillHoldAValue'"
4. Code is not english/natural language, and does not have to abide by the same rules
5. Its easier to type shorter names

While these arguments might seem like strawmen at the moment they are at least real reasons I have heard, so I will respond to these for now.

1. It is not always obvious what the name means. Even for things like 'w is for writer' there are other objects that can reasonably be named something starting with 'w', so therefore no, 'w' is not just for 'writer'. Language is hard enough when used in as clear a manner as humanly possible, and adding artificial barriers to understanding has no benefits.
2. This is somewhat true, and if someone looked at my code and said "I don't know what this variable name means" I would honestly consider their opinion and try to come up with a better name. It may be worthwhile naming something "httpServer". I don't think the need for a perfect name discounts the reason to have a somewhat readable name, though. "s" tells you close to nothing, while "server" tells you that this object is a server of some kind.
3. This is a false dichotomy, as if the only two choices are "single letter variable names" and "extremely long variable names". There is a middle ground that is usually a reasonable choice.
4. Code is not english, which is obviously true, but this statement implies a deeper meaning, which is that code is not meant for another human reader to consume but rather for the computer to execute. In that way code is different, but I think it is just as important for code to be readable by humans as regular language. Code that cannot be understood cannot be debugged, in which case it will be thrown out the moment it has a problem.
5. Get a better editor with autocompletion

I think the terse naming convention is doing harm to a generation of programmers. There are probably many programmers here with more than 10/15/20 years of experience and are unlikely to change their ways, but note that new programmers are looking at existing code to see what norms to follow. I have had something like the following conversation at work a handful of times already

  me: Can we use readable variable names in this go code?

  programmer: The go way is to use single letter names.

  me: Why? Isn't it harder to read?

  programmer: I don't know, thats just how it is.

&amp;#x200B;

Which I think is real tragedy. To be fair, this isn't strictly a problem with Go, but more than 90% of Go code I read online uses single letter variable names (or abbreviated names, which are somehow worse). I don't know this for a fact, but it seems like this problem really started from the beginnings of programming in the 50s and 60s when using longer variable names had real practical problems: namely the compilers were extremely memory limited, and there were no editors with autocompletion. I can sympathize with people from that era, but today we no longer have those problems, and can cast off the restrictions they applied to themselves.
## [5][time.ParseDuration format in other languages](https://www.reddit.com/r/golang/comments/j7un0z/timeparseduration_format_in_other_languages/)
- url: https://www.reddit.com/r/golang/comments/j7un0z/timeparseduration_format_in_other_languages/
---
I'm writing a task runner service, and I would like for the client of this service to be able to specify preferred duration for a task. So my initial implementation was for "task start" requests to have duration as a string that could be parsed by time.ParseDuration.

But then I started thinking. This would be part of a public interface. Not all clients would be written in golang and the duration format would not be standard in those languages. 

So do you know of any libraries in other languages that construct time duration in the golang format? 

It's not hard code to write from scratch, but I would like to provide some references if someone doesn't want to or can't write a duration parser. I expect most of the clients of this service to be in JS or python.
## [6][Get data from data frame](https://www.reddit.com/r/golang/comments/j7rw5s/get_data_from_data_frame/)
- url: https://www.reddit.com/r/golang/comments/j7rw5s/get_data_from_data_frame/
---
I'm using qframe ([https://github.com/tobgu/qframe](https://github.com/tobgu/qframe)) to load a large csv file into a data frame. 

I can get the column names out of the data frame, but is there a way I can get the data out of this frame? e.g. getting each row of the data frame as an array, which can then be search through.

I'm new to this type of data formatting (I think that's the right term) so I don't really know what I'm doing. I found the documentation confusing
## [7][Anyone has experience with logrus SetReportCaller()?](https://www.reddit.com/r/golang/comments/j7s7xb/anyone_has_experience_with_logrus_setreportcaller/)
- url: https://www.reddit.com/r/golang/comments/j7s7xb/anyone_has_experience_with_logrus_setreportcaller/
---
So I'm following the examples on 

[https://github.com/sirupsen/logrus/issues/63](https://github.com/sirupsen/logrus/issues/63)

where I just do:

log.SetReportCaller(true)

&amp;#x200B;

But all this is logging is:  


func="[github.com/sirupsen/logrus.(\*Logger](https://github.com/sirupsen/logrus.(*Logger)).Log" file="C:&lt;my go location&gt;/go/pkg/mod/github.com/sirupsen/logrus@v1.4.2/logger.go:192"  


This happens no matter what, whether I log it in main(), or other functions. How do I get it to show the function that called the logger? Ideally, I was hoping for something like:  


func="myfunc" file="myFile.go"

&amp;#x200B;

Just me, or did someone run into this before as well?
## [8][How to do memory profiling](https://www.reddit.com/r/golang/comments/j7xus8/how_to_do_memory_profiling/)
- url: https://www.reddit.com/r/golang/comments/j7xus8/how_to_do_memory_profiling/
---
I am trying to find a memory leak and I struggle to get basic profiling going.

All the tutorial I found seem to be outdated or at least I don't get the expected results by following them.

I managed to write out a heap dump using pprof.WriteHeapProfile().

The resulting file is compressed with gzip(), but when I uncompress it it's a proprietary binary format.

I tried to read it with "go tool pprof heap.dmp", but I get "Possible precedence issue with control flow operator at /home/user/tools/go3/pkg/tool/linux\_amd64/pprof line 3008."  


When I try "go tool pprof a.out heap.dmp" I get "Possible precedence issue with control flow operator at /home/user/tools/go3/pkg/tool/linux\_amd64/pprof line 3008. heap.dmp: header size &gt;= 2\*\*16"

Can you guys help me? 

I'd be happy with just a CSV of my allocations.

Thanks!
## [9][Assembly-optimised positional popcount routines—now with up to 66 GB/s!](https://www.reddit.com/r/golang/comments/j7dqxb/assemblyoptimised_positional_popcount_routinesnow/)
- url: https://github.com/clausecker/pospop
---

## [10][Web Safety - Golang Backend](https://www.reddit.com/r/golang/comments/j7n8wn/web_safety_golang_backend/)
- url: https://www.reddit.com/r/golang/comments/j7n8wn/web_safety_golang_backend/
---
Looking for any recommended articles on general Web Security for a simple Golang backend.

Thanks!
