# golang
## [1][V1.5 of sqlc released: Compile SQL to type-safe Go](https://www.reddit.com/r/golang/comments/i4iv8o/v15_of_sqlc_released_compile_sql_to_typesafe_go/)
- url: https://sqlc.dev/posts/2020/08/05/sqlc-one-point-five-released.html
---

## [2][Interface Definition in implementing package](https://www.reddit.com/r/golang/comments/i4p1xi/interface_definition_in_implementing_package/)
- url: https://www.reddit.com/r/golang/comments/i4p1xi/interface_definition_in_implementing_package/
---
In [CodeReviewComments: Interfaces](https://github.com/golang/go/wiki/CodeReviewComments#interfaces), it is stated that an interface definition should not be done in the package that implements that interface, but in the package that uses implementations of that interface instead:

&amp;#x200B;

&gt;Go interfaces generally belong in the package that uses values of the interface type, not the package that implements those values. The implementing package should return concrete (usually pointer or struct) types

&amp;#x200B;

There's even an example that shows how *not* to do it:

&amp;#x200B;

&gt;// DO NOT DO IT!!!  
&gt;  
&gt;package producer  
&gt;  
&gt;type Thinger interface { Thing() bool }  
&gt;  
&gt;type defaultThinger struct{ … }  
&gt;  
&gt;func (t defaultThinger) Thing() bool { … }  
&gt;  
&gt;func NewThinger() Thinger { return defaultThinger{ … } }

&amp;#x200B;

However, I'm asking myself: Why? I see this being done many, many times, especially in libraries. And I actually like the idea that a package provides an interface as a clean public API and provides an invisible implementation itself.

Is it really that bad? Can anyone explain what's wrong with the example?
## [3][Google Photos API client](https://www.reddit.com/r/golang/comments/i4oj4q/google_photos_api_client/)
- url: https://www.reddit.com/r/golang/comments/i4oj4q/google_photos_api_client/
---
If anyone is interested here's my take on Google Photos API client. 

[https://github.com/duffpl/google-photos-api-client](https://github.com/duffpl/google-photos-api-client)

Google removed some time ago Photos from list of auto generated clients. I've been using ones already existing (based on mirrored copy of the generated client) but I didn't like them very much so I've decided to give it a shot and write one from scratch. Library API shouldn't change much in future.

What's implemented:

\- Most of endpoints/methods  
\- Basic uploading functionality (no "fancy" stuff like resumable uploads)  
\- Additional wrapper methods that deal with paging automatically (async with channels and sync versions spewing out slices)  
\- Basic error handling

Todo:  
\- Tests (units and functional to keep checking if API is still responding as expected)  
\- Implement sharedAlbums endpoints  
\- Better error handling

Constructive criticism always welcomed :)  
Enjoy
## [4][An aesthetically pleasing video on SEARCHING ALGORITHMS.](https://www.reddit.com/r/golang/comments/i4c9kg/an_aesthetically_pleasing_video_on_searching/)
- url: https://youtu.be/FBJKwjTwNTo
---

## [5][Blank-Xu/sqlx-adapter: Sqlx Adapter for Casbin V2](https://www.reddit.com/r/golang/comments/i4oqrk/blankxusqlxadapter_sqlx_adapter_for_casbin_v2/)
- url: https://github.com/Blank-Xu/sqlx-adapter
---

## [6][How do you deal with Go-Java or Go-Python environments](https://www.reddit.com/r/golang/comments/i45rpr/how_do_you_deal_with_gojava_or_gopython/)
- url: https://www.reddit.com/r/golang/comments/i45rpr/how_do_you_deal_with_gojava_or_gopython/
---
  I'm not sure if this is the right place to ask this.  But I've been coding in Go for the past 4 years.   Some of my earlier Go projects consisted of working with some former Googlers.  So they beat into me the Go way of doing things.   Go was a match made in heaven for me.  Its all about simple and explicit above all else.

  I love working with Go, and it's really the only language I love working with these days.  With that said I've had a few projects where I'm seeing Go written in some very odd and non-idiomatic ways.   I recently joined a position where I constantly see Java devs trying to make Go into Java.  There is a lot of magic and confusing abstractions all over the place.   

  What is worse is that there are many devs here who still actively develop in Java since we still have a lot of code in Java.   And some of our code bases are several years old written in a Java style.

   Have you seen this?  How have you been able to promote Go in your workplace and make sure people are keeping things simple and idiomatic.  How do you make sure people don't carry over mindsets taught in other languages?

  Again if this is the wrong forum to discuss this, I do apologize.  I do see myself encountering even more Go-Java code bases in the future.  Currently in my workplace peoole hate Go.  Mostly because its harder for their confusing abstractions that they write in their Java code.  I'd argue many of these abstractions are unnecessary, but who am really?   Again, any advice?
## [7][Question about Golang performance](https://www.reddit.com/r/golang/comments/i4kz5r/question_about_golang_performance/)
- url: https://www.reddit.com/r/golang/comments/i4kz5r/question_about_golang_performance/
---
First of all i would like to clarify that i love the go. I love the go-ways to do things in contrast to any language out there, i love the standard library and i love the simplicity. This is not a bashing post, it's just that the performance is a really big seller for me.

As i understand it, go is really fast. This lead me to think that, for example, it would way more performant than say 'C#' in anything. However, just today i decided to google some performance comparisons and got surprised to find a lot of resources that indicate go under performs next to C#. These benchmarks/comparisons are not from long ago,

[Resource 1](https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/go-csharpcore.html)

[Resource 2](https://www.reddit.com/r/golang/comments/a88vww/why_is_go_without_generics_is_slower_than_c_with/ec926b7?utm_source=share&amp;utm_medium=web2x)

[Resource 3](https://news.ycombinator.com/item?id=20947286)

[Resource 4](https://www.reddit.com/r/rust/comments/akluxx/rust_now_on_average_outperforms_c_in_the/ef63zdi?utm_source=share&amp;utm_medium=web2x)

[Resource 5](https://www.reddit.com/r/rust/comments/akluxx/rust_now_on_average_outperforms_c_in_the/ef5vuyp?utm_source=share&amp;utm_medium=web2x)

[Resource 6](https://medium.com/@alexyakunin/go-vs-c-part-1-goroutines-vs-async-await-ac909c651c11)

[Resource 7](https://www.reddit.com/r/rust/comments/akluxx/rust_now_on_average_outperforms_c_in_the/)

[Resource 8](https://www.reddit.com/r/rust/comments/akluxx/rust_now_on_average_outperforms_c_in_the/ef5y2ub/?utm_source=share&amp;utm_medium=web2x)

I know that at the end of the day, they are just that: benchmarks. They don't really say 'anything'  outside the context of that particular benchmark. It's just that i was a religious believer on Golang performance and thought it was WAY faster than C#/Java in pretty much anything.  I also want to clarify that i'm extremely naive to these low level things and what these benchmarks really mean.  

So, my intention with this post is to find out if Golang is still what i think it is, and if these benchmarks mean nothing at all and i should take it with a grain of salt. I look forward to any piece wisdom/clarification or even criticism you could have. Thanks
## [8][Learn Go Programming Online with these 12 Best Golang Tutorials &amp; Courses](https://www.reddit.com/r/golang/comments/i4p47f/learn_go_programming_online_with_these_12_best/)
- url: https://www.reddit.com/r/golang/comments/i4p47f/learn_go_programming_online_with_these_12_best/
---
Learn or improve your [Golang](https://blog.coursesity.com/best-golang-tutorials?utm_source=reddit&amp;utm_medium=social&amp;utm_campaign=redditPost&amp;utm_term=best-go) skills online with these curated online tutorials and courses for beginners
## [9][Cloud Automation and DevOps Consulting Services](https://www.reddit.com/r/golang/comments/i4pyo1/cloud_automation_and_devops_consulting_services/)
- url: http://selleo-devops.icu
---

## [10][read image from/write to clipboard(supported windows,mac,linux)](https://www.reddit.com/r/golang/comments/i4mb94/read_image_fromwrite_to_clipboardsupported/)
- url: https://www.reddit.com/r/golang/comments/i4mb94/read_image_fromwrite_to_clipboardsupported/
---
[https://github.com/skanehira/clipboard-image](https://github.com/skanehira/clipboard-image)
