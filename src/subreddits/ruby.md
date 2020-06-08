# ruby
## [1][The RuboCop Name Drama Redux](https://www.reddit.com/r/ruby/comments/gywwmr/the_rubocop_name_drama_redux/)
- url: https://metaredux.com/posts/2020/06/08/the-rubocop-name-drama-redux.html
---

## [2][Galaaz: a system for tightly coupling (Truffle)Ruby and R](https://www.reddit.com/r/ruby/comments/gyolnv/galaaz_a_system_for_tightly_coupling_truffleruby/)
- url: https://github.com/rbotafogo/galaaz
---

## [3][Tracking Changes on Action Text](https://www.reddit.com/r/ruby/comments/gyxurg/tracking_changes_on_action_text/)
- url: https://www.driftingruby.com/episodes/tracking-changes-on-action-text?utm_medium=social&amp;utm_campaign=weekly_episode&amp;utm_source=reddit
---

## [4][Issue request on Github to rename Rubocop due to the word "cop"](https://www.reddit.com/r/ruby/comments/gy6cza/issue_request_on_github_to_rename_rubocop_due_to/)
- url: https://github.com/rubocop-hq/rubocop/issues/8091
---

## [5][Super basic question from a newbie](https://www.reddit.com/r/ruby/comments/gy5v15/super_basic_question_from_a_newbie/)
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
## [6][Rookie dev here... Quick question on non-rails coding.](https://www.reddit.com/r/ruby/comments/gy3ehy/rookie_dev_here_quick_question_on_nonrails_coding/)
- url: https://www.reddit.com/r/ruby/comments/gy3ehy/rookie_dev_here_quick_question_on_nonrails_coding/
---
I'm interested in coding some simple bots that can gather and share information to either a spreadsheet or website. Wondering if there is a solid guide on creating simple web-bots to handle some tedious tasks somewhere? I'm new to all this, so any advice is appreciated on the topic.
## [7][Return most occurring object in an array?](https://www.reddit.com/r/ruby/comments/gxzcqy/return_most_occurring_object_in_an_array/)
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
## [8][ROM-rb + Rails in production](https://www.reddit.com/r/ruby/comments/gxv3y0/romrb_rails_in_production/)
- url: https://www.reddit.com/r/ruby/comments/gxv3y0/romrb_rails_in_production/
---
Hi all,

O would like to know about prodution experiences in ROM-rb + Rails. Its robust enough?
## [9][Simple Downloads Folder Organiser Written In Ruby](https://www.reddit.com/r/ruby/comments/gxknzc/simple_downloads_folder_organiser_written_in_ruby/)
- url: https://projectlode.com/projects/simple-download-folder-organiser/guides/simple-downloads-folder-organiser-written-in-ruby
---

## [10][LocalJumpError given even if there is a block](https://www.reddit.com/r/ruby/comments/gxmrvh/localjumperror_given_even_if_there_is_a_block/)
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
