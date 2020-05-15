# ruby
## [1][Why Fullstaq Ruby?](https://www.reddit.com/r/ruby/comments/gk5i19/why_fullstaq_ruby/)
- url: https://www.joyfulbikeshedding.com/blog/2020-05-15-why-fullstaq-ruby.html
---

## [2][A heavily tested (2k lines) and commented classic Red Black Tree implementation in Python and Ruby. Great for learning the material.](https://www.reddit.com/r/ruby/comments/gk51nf/a_heavily_tested_2k_lines_and_commented_classic/)
- url: https://www.reddit.com/r/ruby/comments/gk51nf/a_heavily_tested_2k_lines_and_commented_classic/
---
Back when I was trying to implement the structure, I could not find any open source implementations that were well written and commented. I did not manage to find any implementation that had any significant amount of tests and as such was not sure if it even worked correctly.

I tried my best to describe the different operations needed thoroughly and have written a lot of tests (functional too) covering all operations, with drawn out trees in comments.

https://github.com/stanislavkozlovski/Red-Black-Tree

Any feedback is greatly appreciated :)
## [3][Foot Traffic: pure Ruby DSL for Chrome scripting based on Ferrum](https://www.reddit.com/r/ruby/comments/gjqaji/foot_traffic_pure_ruby_dsl_for_chrome_scripting/)
- url: https://github.com/lewagon/foot_traffic
---

## [4][LeetCode Challenges Contains Duplicate | Jewels and Stones](https://www.reddit.com/r/ruby/comments/gk2iby/leetcode_challenges_contains_duplicate_jewels_and/)
- url: https://youtu.be/3g7BGGzGw5I
---

## [5][Aaron Patterson (tenderlove) and Zamith are live streaming an hour long conversation on Ruby](https://www.reddit.com/r/ruby/comments/gjqexq/aaron_patterson_tenderlove_and_zamith_are_live/)
- url: https://twitch.tv/wearesubvisual
---

## [6][Matz is calling for feedback on Ruby 2.7/3.0 keyword argument pain](https://www.reddit.com/r/ruby/comments/gjjayp/matz_is_calling_for_feedback_on_ruby_2730_keyword/)
- url: https://discuss.rubyonrails.org/t/new-2-7-3-0-keyword-argument-pain-point/74980
---

## [7][How to Fix Slow Code in Ruby](https://www.reddit.com/r/ruby/comments/gjqwh9/how_to_fix_slow_code_in_ruby/)
- url: https://engineering.shopify.com/blogs/engineering/how-fix-slow-code-ruby
---

## [8][Thread scheduler for light weight concurrency. by ioquatix · Pull Request #3032 · was merged. :heart:](https://www.reddit.com/r/ruby/comments/gjjtla/thread_scheduler_for_light_weight_concurrency_by/)
- url: https://github.com/ruby/ruby/pull/3032
---

## [9][Binpacking SQS batch requests for fun and profit](https://www.reddit.com/r/ruby/comments/gjte73/binpacking_sqs_batch_requests_for_fun_and_profit/)
- url: http://www.wjwh.eu/posts/2020-05-14-binpacking-sqs.html
---

## [10][Help porting a ruby script?](https://www.reddit.com/r/ruby/comments/gjupij/help_porting_a_ruby_script/)
- url: https://www.reddit.com/r/ruby/comments/gjupij/help_porting_a_ruby_script/
---
First the script:

    require 'mechanize'
    require 'nokogiri'
    require 'open-uri'

    variables here
    
    m=Mechanize.new

    def client_login(a)
        a.get($GW+"/users/sign_in") do |page|
            login=page.forms.first
            login['user[email]']=$GWUSER
            login['user[password]']=$GWPASS
            login.submit
        end
    end

    def system_login(a,url)
        a.get(url+"/login") do |page|
            login=page.forms.first
            login['login'] = $USER;
            login['password'] = $PASS;
            login.submit
        end
    end

    def request_report(a,url,path,paramaters=[])
        a.get(url+path,paramaters)
    end

    client_login(m)
    puts "System login:"
    pp system_login(m, $HOST)
    puts "Report:"
    request_report(m,$HOST,$REPORT).save_as $REPORTNAME

This was given to me to handle some automated reporting.  While this works fine as proof of concept, it needs to be built out a decent amount (looping through several sites, several reports, etc) and I'd really rather not add yet another language to my desperately needs to be standardized work environment.

Unfortunately i'm basically a complete beginner when it comes to web concepts/http requests outside of some light stuff with postman, but i've spent the last few days trying to port this python, with 0 success, failing instantly at the client_login, as i can't seem to get it to parse.

The requests/beautiful soup/pymechanize attempts didn't work as it seems to just give me back some junk in the header (i think images being returned as 64b encodings), and I can't figure out any way to either ignore that, parse it, or skip it like the ruby script seems to be doing.

Hitting thesite with a postman get works fine, but I really don't want to do it there if I can avoid it, as I haven't figured out how to 'log in' through that and continues to clutter how many languages/tools i'm using.

From some reading i figured it might have to do with JS(altough i'm not sure), so i was going to automate it with selenium, which seems to be the goto, but this introduces new problems.  The original ruby script is very fast, since as far as i can tell it's not using a browser, but selenium crawls and is horrifically inconsistent (in part due to our lousy connection), so something that the ruby script can do in moments the selenium "half answer" i mocked up was taking at least half a minute, even headless.  This is unacceptably slow.

I realize that asking on the ruby subreddit for help getting away from ruby might not be the best place, but all the python help has been a total deadend so far, and I've been digging through documentation and and stack with no luck.

If someone could even just point me  in the right direction, something like "oh yeah mechanize is doing X here and that's why it can read the html so you'll want to look for this in the documentation" that would help a ton.  I'm feeling very lost and while my final fallback will be just maintain this one script in ruby, i honestly already using too many languages in my work setup.
