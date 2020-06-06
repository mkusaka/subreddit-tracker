# ruby
## [1][Simple Downloads Folder Organiser Written In Ruby](https://www.reddit.com/r/ruby/comments/gxknzc/simple_downloads_folder_organiser_written_in_ruby/)
- url: https://projectlode.com/projects/simple-download-folder-organiser/guides/simple-downloads-folder-organiser-written-in-ruby
---

## [2][LocalJumpError given even if there is a block](https://www.reddit.com/r/ruby/comments/gxmrvh/localjumperror_given_even_if_there_is_a_block/)
- url: https://www.reddit.com/r/ruby/comments/gxmrvh/localjumperror_given_even_if_there_is_a_block/
---
I wrote this function:  


`def tail(times)`

  `return if times == 0`

  `yield`

  `tail(times - 1)`  
`end`  


then called it with a block:  


`tail(n) {|i| puts(i) }`  


and I got this:  


`tail.rb:6:in \`tail': no block given (yield) (LocalJumpError)`

  
Can someone explain how should I do this and why I get that exception raised?
## [3][A RSS reader in Ruby](https://www.reddit.com/r/ruby/comments/gx38l1/a_rss_reader_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/gx38l1/a_rss_reader_in_ruby/
---
Hello All,  I wrote a terminal based RSS reader [https://gitlab.com/mindaslab/rss\_reader](https://gitlab.com/mindaslab/rss_reader) , thought it might be useful for some one so am posting it here.
## [4][Can Ruby be compiled into high performance assembly?](https://www.reddit.com/r/ruby/comments/gx70k9/can_ruby_be_compiled_into_high_performance/)
- url: https://www.reddit.com/r/ruby/comments/gx70k9/can_ruby_be_compiled_into_high_performance/
---
I'm looking at C++ examples and then I'm looking back at Ruby and I don't see why Ruby can't be compiled to the same performance. Imagine if they made a Ruby compiler that had C++ performance, every other programming language would instantly die as Ruby would be the easiest and fastest. Can it be done honestly or is there some limitation I can't see?
## [5][TIL you can pass heredocs tokens as arguments and even stack them](https://www.reddit.com/r/ruby/comments/gwqllv/til_you_can_pass_heredocs_tokens_as_arguments_and/)
- url: https://www.reddit.com/r/ruby/comments/gwqllv/til_you_can_pass_heredocs_tokens_as_arguments_and/
---
I've been writing Ruby since before Rails came out and I had no idea you could do this. 

You can pass the token defining the heredoc as an argument, and the continue the line:

    call_method(&lt;&lt;DOC).do_something(5)
      SOME TEXT
    DOC

You can even stack the heredocs and call them as multiple arguments:

    call_method(&lt;&lt;-DOC1, &lt;&lt;~DOC2).do_something
      this is doc 1
    DOC1
      this is doc 2
    DOC2

Not something that will matter very often, but a cool discovery.

---

Edit - Other cool stuff from this thread:

/u/yxhuvud pointed out that it is possible to call methods on the heredoc too:

    foo(&lt;&lt;~SQL.strip)
      SELECT
        foo, bar
      FROM tables
    SQL

/u/riffraff pointed out the `DATA` and `__END__` syntax as well ([explanation link](https://www.honeybadger.io/blog/data-and-end-in-ruby/))

    call_method(DATA).do_something

    # some code

    __END__
    Evertyhing after the above line
    is what will go in DATA
## [6][Run Rails script as an ActiveJob job](https://www.reddit.com/r/ruby/comments/gx2jzd/run_rails_script_as_an_activejob_job/)
- url: https://blog.eq8.eu/til/run-script-as-a-activejob-sidekiq-job.html
---

## [7][session[:user_id] problem in a sinatra app](https://www.reddit.com/r/ruby/comments/gx0ize/sessionuser_id_problem_in_a_sinatra_app/)
- url: https://www.reddit.com/r/ruby/comments/gx0ize/sessionuser_id_problem_in_a_sinatra_app/
---
I wrote this, but I have problems with my helper "current\_user". I don't  know what I exactly should put there to get it into the work : [https://gist.github.com/prp-e/91ad2491ffa8926a4c5ff67268078b1f](https://slack-redir.net/link?url=https%3A%2F%2Fgist.github.com%2Fprp-e%2F91ad2491ffa8926a4c5ff67268078b1f&amp;v=3) 

(P.S : It's not a serious project, don't worry about it's security)
## [8][Yielding an array resulting faster than doing a for loop](https://www.reddit.com/r/ruby/comments/gx1kul/yielding_an_array_resulting_faster_than_doing_a/)
- url: https://www.reddit.com/r/ruby/comments/gx1kul/yielding_an_array_resulting_faster_than_doing_a/
---
Code (file iter.rb):

  
`require("benchmark")`

&amp;#x200B;

`array = (1..2000000).to_a()`

&amp;#x200B;

`class Array`

  `def iter()`

`yield self`

  `end`

`end`

&amp;#x200B;

&amp;#x200B;

[`Benchmark.bm`](https://Benchmark.bm)`() {|b|`

  [`b.report`](https://b.report)`("for: ") {|b|`

[`File.open`](https://File.open)`("for.txt", "w") {|f|`

`for i in array do`

`f.puts(i)`

`end`

`}` 

  `}`

  [`b.report`](https://b.report)`("iter:") {|b|`

[`File.open`](https://File.open)`("iter.txt", "w") {|f|`

`array.iter() {|i|`

`f.puts(i)`

`}`

`}`

  `}`

`}` 

&amp;#x200B;

On my Raspberry Pi 4 (`Linux Manjaro 4.19.122-1-MANJARO-ARM #1 SMP PREEMPT Mon May 11 14:20:27 CDT 2020 aarch64 GNU/Linux`), this produces the following:  
`user     system      total        real`

`for:   4.481485   0.152160   4.633645 (  4.665762)`

`iter:  2.933153   0.135692   3.068845 (  3.082631)`  


it's a second and a half faster, just by yielding the array instead of looping into it with a for loop.
## [9][Topaz still one of the fastest ruby implementation, even if it not developed some years, by this benchmarks](https://www.reddit.com/r/ruby/comments/gwt2xc/topaz_still_one_of_the_fastest_ruby/)
- url: https://github.com/kostya/jit-benchmarks
---

## [10][My thought on Ruby/Rails TIMTOWTDI](https://www.reddit.com/r/ruby/comments/gxd32w/my_thought_on_rubyrails_timtowtdi/)
- url: https://www.reddit.com/r/ruby/comments/gxd32w/my_thought_on_rubyrails_timtowtdi/
---
FYI, coming from Java/Python background. I've been using Ruby (on Rails) at work since last November, and I think I can understand why ruby is facing decline in popularity. A few points I feel annoyed:

* I'm mostly annoyed by TIMTOWTDI, no consistency enforced, which makes it difficult to maintain code and slows down the development.
* Also, even though there is more than one way to do, only few of them are documented properly and the rest are abandoned. Why not just move with one clear way?
*  I feel like Ruby/Rails' duck typing does not fit well. I once had to deal with error, and it was because of type of hashes, one with \`HashWithIndifferentAccess\`and another one with regular \`Hash\`.  Unless it's a language like js where \`{}\` are all object, this sucks.

My overall thought on this language&amp;framework so far is that someone literally put all different features from different languages in a bag and shook it up to create the language.
