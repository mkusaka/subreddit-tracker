# golang
## [1][[Academic] Open Source Development Survey (Software Developers using GitHub)](https://www.reddit.com/r/golang/comments/j0ovtc/academic_open_source_development_survey_software/)
- url: /r/SampleSize/comments/iybrqr/academic_open_source_development_survey_software/
---

## [2][How to Structure the code?](https://www.reddit.com/r/golang/comments/j0m4ew/how_to_structure_the_code/)
- url: https://www.reddit.com/r/golang/comments/j0m4ew/how_to_structure_the_code/
---
Hello guys,

* I've written a small program to benchmark the query timing for an open source tool called [Apache Druid](https://druid.apache.org/).
* Apache Druid is a high performance real-time analytics database.
* Here is the code I wrote - [https://play.golang.org/p/5naCxt2NzoB](https://play.golang.org/p/5naCxt2NzoB)

&amp;#x200B;

What does the code do?

* Read a JSON file (Druid's query will  be in JSON format)
* Create `Transport`, `Client` and a `NewRequest`
* Execute the request.
* The purpose of the code is to execute the request and collect the elapsed time (for benchmarking purpose)

&amp;#x200B;

I'm relatively new to Go, and would like some feedback on code style, the way I execute requests (good or bad), how to separate the functions in their respective folders, also I've not created any custom `type` is this fine for this code?

&amp;#x200B;

End Goal

* I've 40 JSON files (each JSON file is a query which will be executed to Broker endpoint)
* Make all 40 queries run concurrently
* Aggregate and average out the timing for each query which runs several times

&amp;#x200B;

Any feedback would be much appreciated.

&amp;#x200B;

Thank you,

Jeet
## [3][What is var _ type = &amp;type{} used for?](https://www.reddit.com/r/golang/comments/j0b73d/what_is_var_type_type_used_for/)
- url: https://www.reddit.com/r/golang/comments/j0b73d/what_is_var_type_type_used_for/
---
I have seen this in a test file:

    import (
    
    "image"
    
    "image/draw"
    
    )
    
    var _ Node = &amp;BasicNode{}

What is the purpose of `var _ Node = &amp;BasicNode{}`?
## [4][Just released a little library, Peek-A-Buf](https://www.reddit.com/r/golang/comments/j0fv77/just_released_a_little_library_peekabuf/)
- url: https://www.reddit.com/r/golang/comments/j0fv77/just_released_a_little_library_peekabuf/
---
Hi,

I just released a little library I'd like to share here: Peek-A-Buf

Peek-A-Buf is a buffered reader with side effect free peeking capability.

Quite niche, very tiny, but maybe someone finds this a bit interesting. :)

[https://github.com/philiplb/peekabuf](https://github.com/philiplb/peekabuf)
## [5][Ideas on how to implement a timetracker](https://www.reddit.com/r/golang/comments/j0qn8z/ideas_on_how_to_implement_a_timetracker/)
- url: https://www.reddit.com/r/golang/comments/j0qn8z/ideas_on_how_to_implement_a_timetracker/
---
Hello everyone. Title is a bit messy i know.  


In my trip to learn golang i have created a personal tasklist app using react and a backend in go+gORM.

Everything is working well so i decided to try and create a time tracking capability.

This means that i want to go on my task, hit the button "track"/"start working on" or something similar and that will start tracking the time working on a task.

I know that i can track the time in the frontend and send the values on stop but if the browser dies, freezes, if i close the tab and maybe even change it(not sure what would happen then tbh) the time tracked will be lost.

I am guessing it's gonna need sockets to work this out?

Any known examples or some pointers to what to look into for this task?

TYVM.
## [6][Should I have started with some book other than "Go programming language"?](https://www.reddit.com/r/golang/comments/j0ee0o/should_i_have_started_with_some_book_other_than/)
- url: https://www.reddit.com/r/golang/comments/j0ee0o/should_i_have_started_with_some_book_other_than/
---
I started this book yesterday and just now got to point 2.

The book starts off by showcasing gifs, local webservers, and introduces things such as maps right off the bat. The excercises reflect stuff he glosses over earlier in the pages...

Am I missing something? It feels way too verbose for a beginner in a language, and this isn't even my first programming language (Know Python for 1 year)

Should I have went for another resource, or should I continue with this one?
## [7][Why does comparison between an error created with error.New() and an error assigned to a variable return false?](https://www.reddit.com/r/golang/comments/j0p15i/why_does_comparison_between_an_error_created_with/)
- url: https://www.reddit.com/r/golang/comments/j0p15i/why_does_comparison_between_an_error_created_with/
---
So if I have an error defined as

`var MyError = errors.New("my error")`

And then try to do the following 

`MyError == error.New("my error")`

This returns false? What mechanics are used for this comparison and why is it false?
## [8][ESP32 and ESP8266 support in TinyGo](https://www.reddit.com/r/golang/comments/j027qe/esp32_and_esp8266_support_in_tinygo/)
- url: https://aykevl.nl/2020/09/tinygo-esp32
---

## [9][Time.UTC working as intended?](https://www.reddit.com/r/golang/comments/j0gu2g/timeutc_working_as_intended/)
- url: https://www.reddit.com/r/golang/comments/j0gu2g/timeutc_working_as_intended/
---
I noticed in the example for [Time.Format](https://golang.org/pkg/time/#Time.Format), it outputs the following:

    Unix format: Wed Feb 25 11:06:39 PST 2015
    Same, in UTC: Wed Feb 25 11:06:39 UTC 2015

This was unexpected, why isn't it `19:06:39 UTC`?
## [10][Up to which struct size is passing by value quicker then passing by pointer?](https://www.reddit.com/r/golang/comments/j05usq/up_to_which_struct_size_is_passing_by_value/)
- url: https://www.reddit.com/r/golang/comments/j05usq/up_to_which_struct_size_is_passing_by_value/
---
I created some benchmarks which suggest that on my machine a struct consisting of up to 9 int 64 is faster to be passed as value. From 10 fields and more passing by pointer is faster.

[https://github.com/komkom/value-vs-ptr](https://github.com/komkom/value-vs-ptr)

Are these tests correct? And if they are correct why would it be that on both of my machines the breaking point is 9 fields?
