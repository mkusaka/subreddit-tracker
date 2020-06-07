# ruby
## [1][Issue request on Github to rename Rubocop due to the word "cop"](https://www.reddit.com/r/ruby/comments/gy6cza/issue_request_on_github_to_rename_rubocop_due_to/)
- url: https://github.com/rubocop-hq/rubocop/issues/8091
---

## [2][Lets remove american cultural imperialism from the tech equation.](https://www.reddit.com/r/ruby/comments/gy8bop/lets_remove_american_cultural_imperialism_from/)
- url: https://www.reddit.com/r/ruby/comments/gy8bop/lets_remove_american_cultural_imperialism_from/
---
This seems to be a recurring problem where a very local minority of americans try to force their views and values onto the rest of the world through tech.

This.Needs.To.Stop.

Americans cant be the sole arbiters in the world to decide what's wrong and right and cant force their values on the rest of us. 
On top of that this is highly offensive,disrespectful and directly attacks the health and sanity of the individuals or organizations affected by your americansplaining.

What's americansplaining?

Patronizing explanation why an unnecessary change is required to meet cultural and societal standards of someone from the US.

Software is not political. You are political. You are making everything about politics and that's entirely your problem not mine.
Your issues are not my issues. Your believes are not my believes.
If something offends you that's your problem not mine. 

Keep american political views out of TECH. Better. Keep any political views out of tech.
## [3][Super basic question from a newbie](https://www.reddit.com/r/ruby/comments/gy5v15/super_basic_question_from_a_newbie/)
- url: https://www.reddit.com/r/ruby/comments/gy5v15/super_basic_question_from_a_newbie/
---
Hi, I am very new to ruby so I hope I'm in the right place for this. I was wondering what the easiest way to pass around variables is without using global variables.

eg.

I would like to have something like this:

def main

setting\_area

writing\_area

end

&amp;#x200B;

def setting\_area

var = 'johnny'

end

&amp;#x200B;

def writing\_area

puts var

end

&amp;#x200B;

main

&amp;#x200B;

But have it function like this:

&amp;#x200B;

def main

setting\_area

end

&amp;#x200B;

def setting\_area

var = 'johnny'

puts var

end

&amp;#x200B;

main

thankyou in advance :)
## [4][Rookie dev here... Quick question on non-rails coding.](https://www.reddit.com/r/ruby/comments/gy3ehy/rookie_dev_here_quick_question_on_nonrails_coding/)
- url: https://www.reddit.com/r/ruby/comments/gy3ehy/rookie_dev_here_quick_question_on_nonrails_coding/
---
I'm interested in coding some simple bots that can gather and share information to either a spreadsheet or website. Wondering if there is a solid guide on creating simple web-bots to handle some tedious tasks somewhere? I'm new to all this, so any advice is appreciated on the topic.
## [5][Return most occurring object in an array?](https://www.reddit.com/r/ruby/comments/gxzcqy/return_most_occurring_object_in_an_array/)
- url: https://www.reddit.com/r/ruby/comments/gxzcqy/return_most_occurring_object_in_an_array/
---
Hello! I am currently learning all to do with classes and how they relate to each other. Right now, I'm working on a has many: through relationships. The classes are: AirBnB listings, Trips, and Guests. the Listings can have many Trips, the Guests Trips, but the Trips only one Listing and Guest.

&amp;#x200B;

I'm writing methods in the Listing file. The method is self.most\_popular.

&amp;#x200B;

Here is my code so far:

`def self.most_popular`  
 `Trip.all.select {|trips| trips.listing}`  
 `end` 

This returns: 

`[#&lt;Trip:0x00007fe49e1c8498`

`u/guest=#&lt;Guest:0x00007fe49e1c86c8 u/name="ellen"&gt;,`

`u/listing=#&lt;Listing:0x00007fe49e1c8560 u/city="switzerland", u/listing_name="switzerland_humble"&gt;,`

`u/trip_name="copenhagen escapism"&gt;,`

`#&lt;Trip:0x00007fe49e1c8448`

`u/guest=#&lt;Guest:0x00007fe49e1c8718 u/name="rebecca"&gt;,`

`u/listing=#&lt;Listing:0x00007fe49e1c8560 u/city="switzerland", u/listing_name="switzerland_humble"&gt;,`

`u/trip_name="copenhagen escapism"&gt;,`

`#&lt;Trip:0x00007fe49e1c83f8`

`u/guest=#&lt;Guest:0x00007fe49e1c8650 u/name="robert"&gt;,`

`u/listing=#&lt;Listing:0x00007fe49e1c84e8 u/city="lithuania", u/listing_name="lithuania_shambles"&gt;,`

`u/trip_name="back to our roots"&gt;,`

`#&lt;Trip:0x00007fe49e1c83a8`

`u/guest=#&lt;Guest:0x00007fe49e1c86c8 u/name="ellen"&gt;,`

`u/listing=#&lt;Listing:0x00007fe49e1c84e8 u/city="lithuania", u/listing_name="lithuania_shambles"&gt;,`

`u/trip_name="back to our roots"&gt;,`

`#&lt;Trip:0x00007fe49e1c8358`

`u/guest=#&lt;Guest:0x00007fe49e1c8718 u/name="rebecca"&gt;,`

`u/listing=#&lt;Listing:0x00007fe49e1c84e8 u/city="lithuania", u/listing_name="lithuania_shambles"&gt;,`

`u/trip_name="back to our roots"&gt;]`

I would like it to return "lithuania\_shambles", which is the most occurring listing\_name.
## [6][Simple Downloads Folder Organiser Written In Ruby](https://www.reddit.com/r/ruby/comments/gxknzc/simple_downloads_folder_organiser_written_in_ruby/)
- url: https://projectlode.com/projects/simple-download-folder-organiser/guides/simple-downloads-folder-organiser-written-in-ruby
---

## [7][LocalJumpError given even if there is a block](https://www.reddit.com/r/ruby/comments/gxmrvh/localjumperror_given_even_if_there_is_a_block/)
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
## [8][A RSS reader in Ruby](https://www.reddit.com/r/ruby/comments/gx38l1/a_rss_reader_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/gx38l1/a_rss_reader_in_ruby/
---
Hello All,  I wrote a terminal based RSS reader [https://gitlab.com/mindaslab/rss\_reader](https://gitlab.com/mindaslab/rss_reader) , thought it might be useful for some one so am posting it here.
## [9][Can Ruby be compiled into high performance assembly?](https://www.reddit.com/r/ruby/comments/gx70k9/can_ruby_be_compiled_into_high_performance/)
- url: https://www.reddit.com/r/ruby/comments/gx70k9/can_ruby_be_compiled_into_high_performance/
---
I'm looking at C++ examples and then I'm looking back at Ruby and I don't see why Ruby can't be compiled to the same performance. Imagine if they made a Ruby compiler that had C++ performance, every other programming language would instantly die as Ruby would be the easiest and fastest. Can it be done honestly or is there some limitation I can't see?
## [10][TIL you can pass heredocs tokens as arguments and even stack them](https://www.reddit.com/r/ruby/comments/gwqllv/til_you_can_pass_heredocs_tokens_as_arguments_and/)
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
