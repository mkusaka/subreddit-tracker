# ruby
## [1][Yukihiro Matsumoto: "Ruby is designed for humans, not machines"](https://www.reddit.com/r/ruby/comments/equpgz/yukihiro_matsumoto_ruby_is_designed_for_humans/)
- url: https://evrone.com/yukihiro-matsumoto-interview
---

## [2][how is to_h different from to_hash](https://www.reddit.com/r/ruby/comments/eqf2b3/how_is_to_h_different_from_to_hash/)
- url: https://www.reddit.com/r/ruby/comments/eqf2b3/how_is_to_h_different_from_to_hash/
---
I have a code like \[\["key", "value"\]\].to\_h that works fine but if I were to use to\_hash with it I am getting an error that says to\_hash isn't a method for Array. 

From here I know [https://zverok.github.io/blog/2016-01-18-implicit-vs-expicit.html](https://zverok.github.io/blog/2016-01-18-implicit-vs-expicit.html) that there are differences between to\_h and to\_hash but I'm not sure how I get an undefined method to\_hash for Array error.
## [3][Strange memory retention behaviour](https://www.reddit.com/r/ruby/comments/eqa37d/strange_memory_retention_behaviour/)
- url: https://www.reddit.com/r/ruby/comments/eqa37d/strange_memory_retention_behaviour/
---
This is a follow-up to my post three weeks ago: [https://www.reddit.com/r/ruby/comments/eheb9o/how\_to\_find\_all\_objects\_another\_object\_is/](https://www.reddit.com/r/ruby/comments/eheb9o/how_to_find_all_objects_another_object_is/)

If you follow [this link to the comment on Github](https://github.com/gettalong/hexapdf/issues/97#issuecomment-574455558), you will find code (also copied below) that retains memory but when changing some non-essential parts the memory retention goes way done.  The video linked in the comment explains the problem very well.

This is the code in question, install the latest HexaPDF and memory\_profiler gems to try it out, the font file is linked to in the Github comment:

    require 'memory_profiler'
    require 'hexapdf'
    
    report = MemoryProfiler.report(top: 5) do
      pdf = HexaPDF::Document.new·
      wrapper = pdf.fonts.add('NotoSansCJKtc-Regular.ttf')
      pdf.write('/tmp/gh97-2.pdf', validate: false)
      pdf.pages.each do |page|
      end
      pdf = nil
    end.pretty_print(scale_bytes: true, retained_strings: 0, allocated_strings: 0)·

Either changing `validate` to `true` (more work to do) or commenting out the `pdf.pages.each` loop (less work to do) will bring memory retention down. The thing is, however, that by the end of the outer MemoryProfiler block the whole `HexaPDF::Document` object is not referenced anymore and therefore should always be garbage collectible, including all the referenced objects.

Is there a reason for this very weird behaviour?

One more thing: Invoking `GC.start` at the end of the block changes nothing with respect to the weird behavour.
## [4][Ruby 2.7 speed with JIT?](https://www.reddit.com/r/ruby/comments/eq6dt0/ruby_27_speed_with_jit/)
- url: https://www.reddit.com/r/ruby/comments/eq6dt0/ruby_27_speed_with_jit/
---
Heya everyone! New to the thread, anyone able to tell me if Ruby 2.7 with JIT is beneficial for discord bots? I'm curious about it now, and I'm currently wondering is it worthwhile to use right now, thanks for reading!

EDIT: Some spelling :P
## [5][https://www.rubyonrails.ba/ let me share my small community like site for collecting ruby tech links](https://www.reddit.com/r/ruby/comments/eq66ct/httpswwwrubyonrailsba_let_me_share_my_small/)
- url: https://www.reddit.com/r/ruby/comments/eq66ct/httpswwwrubyonrailsba_let_me_share_my_small/
---
Yesterday we officially reached the **2 millionth** view on Ruby On Rails Bosnia. The site was officially published on November 20, 2016. It took us **24 months** to get first million views (November 11, 2018). 

&amp;#x200B;

https://preview.redd.it/jyhm27jcieb41.png?width=850&amp;format=png&amp;auto=webp&amp;s=a1519880438e2290471e2da119e0ea0d9e0be47b

For the next million views we need to wait next **14 months** (January 14, 2020).

In the last year, we have partnered with 2 very popular EU ruby conferences as media partners. In this way, we rewarded members of our site with few free conference tickets. We are thankful again to our partners [BalkanRuby](https://balkanruby.com/) and [RubyC](https://rubyc.eu/)

**Short site statistic preview:**

Currently, you can find over **2.900 links** on the site to various articles in the Ruby and Ruby On Rails technology stack area.

Website visits on a daily and monthly basis are constantly increasing. Currently, we reach about **2.600 views** on a daily basis. While on an average monthly basis we reach around **90.000 views**. If we consider that the site content is collections of short content previews and links to very specific technology tech stack I think that these results are very significant.

At the end we are inviting all those who are interested in advertising and cooperation to contact us and help us to reach next million views faster.

Thank you!

[https://www.rubyonrails.ba/](https://www.rubyonrails.ba/)
## [6][Extracting value from JSON via API in ruby](https://www.reddit.com/r/ruby/comments/eqebj5/extracting_value_from_json_via_api_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/eqebj5/extracting_value_from_json_via_api_in_ruby/
---
I receive a response from an API but I am unable to make the result as per below with location:  
`[`   
   `{`   
`"description": "...................",`   
`"pricePerUnit": "..........",`   
`"effectiveDate": "............."`   
`"location": "................"`   
   `}`   
`]`  
I have tried the below code:

`require 'httparty'`   
`class ApiTest`   
  `def self.required_data`       
   `url ="https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/AmazonCloudFront/current/index.json"`       
`response = HTTParty.get(url, :query =&gt; {},:headers =&gt; {} )`       
`parsed = JSON.parse(response.body)`   
`result = []`   
`parsed['terms']['OnDemand'].each do |_product_key, offers|`     
`offers.each do |_offer, details|`       
`effective_date = details['effectiveDate']`       
`details['priceDimensions'].each do |_name, dimension|`         
`result &lt;&lt; { 'description' =&gt; dimension['description'], 'effectiveDate' =&gt; effective_date }`  
`end`   
`end`   
`end`  
  `result`  


we have "Sku" under product struct in JSON and the same we have under the second "term" struct. How I can add *location* as well for each?
## [7][Ruby is #1 in NY :tada:](https://www.reddit.com/r/ruby/comments/epyngt/ruby_is_1_in_ny_tada/)
- url: https://hired.com/page/state-of-software-engineers/key-takeaways/
---

## [8][GitHub - e-wrks/edh: Đ (Edh) - The next-big-things ought to happen with Haskell not C/C++](https://www.reddit.com/r/ruby/comments/eqg9qv/github_ewrksedh_đ_edh_the_nextbigthings_ought_to/)
- url: https://github.com/e-wrks/edh
---

## [9][How to build a live, face-to-face video chat app in Ruby on Rails 6.0.2.1 ​](https://www.reddit.com/r/ruby/comments/epwkyk/how_to_build_a_live_facetoface_video_chat_app_in/)
- url: /r/rails/comments/eprjo9/how_to_build_a_live_facetoface_video_chat_app_in/
---

## [10][For those who use Ruby as their primary language, what are the main reasons you stick with it over other languages?](https://www.reddit.com/r/ruby/comments/epnlxh/for_those_who_use_ruby_as_their_primary_language/)
- url: https://www.reddit.com/r/ruby/comments/epnlxh/for_those_who_use_ruby_as_their_primary_language/
---
Ruby isn't as popular these days, but I'm currently learning it as I'm going through a full stack web developer course (app academy) that uses it for their backend teaching, but I'm curious for anyone that writes code day to day with Ruby and **does so out of choice**, why? Python is a common comparison, the advantage being the massive amount of libraries, or Haskell with purely functional programming goodness, Lisp has macros, Erlang or Ada for safety, C++ for performance, but where does Ruby fit in? I have no hatred for Ruby mind you, it seems pretty okay so far but nothing that stands out particularly.

My background is some bash and Common Lisp mainly, and both just for scripts and hobby projects, but I'm going the web developer route so I'm picking up Ruby as per the course. Should I stick with Ruby for the backend when I'm done? That's going to trigger biased answers of course but I'm still interested in what you have to say! Let me know what you think and thanks for reading.
