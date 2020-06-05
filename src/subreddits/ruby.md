# ruby
## [1][A RSS reader in Ruby](https://www.reddit.com/r/ruby/comments/gx38l1/a_rss_reader_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/gx38l1/a_rss_reader_in_ruby/
---
Hello All,  I wrote a terminal based RSS reader [https://gitlab.com/mindaslab/rss\_reader](https://gitlab.com/mindaslab/rss_reader) , thought it might be useful for some one so am posting it here.
## [2][TIL you can pass heredocs tokens as arguments and even stack them](https://www.reddit.com/r/ruby/comments/gwqllv/til_you_can_pass_heredocs_tokens_as_arguments_and/)
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
## [3][Run Rails script as an ActiveJob job](https://www.reddit.com/r/ruby/comments/gx2jzd/run_rails_script_as_an_activejob_job/)
- url: https://blog.eq8.eu/til/run-script-as-a-activejob-sidekiq-job.html
---

## [4][session[:user_id] problem in a sinatra app](https://www.reddit.com/r/ruby/comments/gx0ize/sessionuser_id_problem_in_a_sinatra_app/)
- url: https://www.reddit.com/r/ruby/comments/gx0ize/sessionuser_id_problem_in_a_sinatra_app/
---
I wrote this, but I have problems with my helper "current\_user". I don't  know what I exactly should put there to get it into the work : [https://gist.github.com/prp-e/91ad2491ffa8926a4c5ff67268078b1f](https://slack-redir.net/link?url=https%3A%2F%2Fgist.github.com%2Fprp-e%2F91ad2491ffa8926a4c5ff67268078b1f&amp;v=3) 

(P.S : It's not a serious project, don't worry about it's security)
## [5][Yielding an array resulting faster than doing a for loop](https://www.reddit.com/r/ruby/comments/gx1kul/yielding_an_array_resulting_faster_than_doing_a/)
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
## [6][Topaz still one of the fastest ruby implementation, even if it not developed some years, by this benchmarks](https://www.reddit.com/r/ruby/comments/gwt2xc/topaz_still_one_of_the_fastest_ruby/)
- url: https://github.com/kostya/jit-benchmarks
---

## [7][The Ruby Blend Episode 15: Rails Testing Tools and Best Practices with Jason Swett](https://www.reddit.com/r/ruby/comments/gwn67u/the_ruby_blend_episode_15_rails_testing_tools_and/)
- url: https://fireside.fm/s/ouBAUjGy+C2WN3swa
---

## [8][Dynamic OmniAuth strategy utilizing OpenId Connect?](https://www.reddit.com/r/ruby/comments/gw91oa/dynamic_omniauth_strategy_utilizing_openid_connect/)
- url: https://www.reddit.com/r/ruby/comments/gw91oa/dynamic_omniauth_strategy_utilizing_openid_connect/
---
I’m using OmniAuth right now and things are working fine using the available strategies like ‘omniauth-google-oauth2’. 

I’d like to write on generic strategy, where I can pass in the ‘client_id’, ‘client_secret’, a provider url, etc. It’s my understanding that a generic strategy will only work if the provider is OpenId Connect compliant for standard naming, endpoints, etc. 

I have bit and pieces of it figured out, but struggling to put it all together.  There’s a couple gems that look like generic OpenId Connect strategies, but I’ve run into issues dynamically updating them in the initializer following a pattern like this [dynamic provider ](https://github.com/omniauth/omniauth/wiki/Dynamic-Providers)


Has anyone implemented something similar?
## [9][Shipped print copies of Why's Poignant Guide this Monday as part of Alt::BrightonRuby](https://www.reddit.com/r/ruby/comments/gvu5r1/shipped_print_copies_of_whys_poignant_guide_this/)
- url: https://www.reddit.com/r/ruby/comments/gvu5r1/shipped_print_copies_of_whys_poignant_guide_this/
---
Hey there,

I've read and re-read this for typos and layout errors and am getting ready to send to the printers on Monday (8th). We'll ship anywhere for the price of the £29 Alt::BrightonRuby "ticket" as well as access to the talks (15th June) and the ability to ask questions that I'll put to the speakers on a podcast (late June/early July).

Tickets and further information at https://alt.brightonruby.com

[The cover](https://preview.redd.it/y19ghclw7p251.png?width=1586&amp;format=png&amp;auto=webp&amp;s=c9be747a6a499d76b7ffa862c0eb87da91b27ce2)

[It's a real book.](https://preview.redd.it/6ooqvdsa8p251.jpg?width=3024&amp;format=pjpg&amp;auto=webp&amp;s=41cf241a1748acca77866d60ff9bb95387babde2)
## [10][Learning from others...](https://www.reddit.com/r/ruby/comments/gw3bgv/learning_from_others/)
- url: https://www.reddit.com/r/ruby/comments/gw3bgv/learning_from_others/
---
Noob here. I’ve always found reading another spoken language once you have a handle on the pronunciation helps to cement some of the nuances of the languages just by repetitive review of structures of the language, even if I don’t understand what I’m reading right away the brain seems to work out the semantics. 

Seems like that works with programming as well so I was just curious if anyone has any suggestions for a well-documented/organized small to medium size open-sourced Ruby projects on Github that a beginner would benefit from looking through to see the file structure, documentation, etc.? Thanks for any suggestions in advance.

Edit:  Thanks for all the ideas. Much appreciated.
