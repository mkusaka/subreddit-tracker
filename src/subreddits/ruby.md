# ruby
## [1][`insert_all` and `upsert_all` ActiveRecord methods (a short recap)](https://www.reddit.com/r/ruby/comments/fd9iim/insert_all_and_upsert_all_activerecord_methods_a/)
- url: https://frontdeveloper.pl/2020/03/insert_all-and-upsert_all-activerecord-methods/
---

## [2][Hanami::API on Amazon AWS Lambda](https://www.reddit.com/r/ruby/comments/fdctli/hanamiapi_on_amazon_aws_lambda/)
- url: https://lucaguidi.com/2020/03/04/hanamiapi-on-amazon-aws-lambda/
---

## [3][puma 3.12.3 and 4.3.2 released then yanked](https://www.reddit.com/r/ruby/comments/fd14qi/puma_3123_and_432_released_then_yanked/)
- url: https://www.reddit.com/r/ruby/comments/fd14qi/puma_3123_and_432_released_then_yanked/
---
Puma 3.12.3 and 4.3.2 were released to fix [CVE-2020-5247](https://nvd.nist.gov/vuln/detail/CVE-2020-5247), a vulnerability where "an attacker can use newline characters to end the header and inject malicious content, such as additional headers or an entirely new response body". Which actually seems potentially serious to me, so good to update. 

Github's dependabot may have sent you an automatic PR to do so. So helpful! Maybe you accepted it. 

And then started finding today that you can't deploy/install your project, with some pretty confusing error messages that are confusing you, wait what?

Because... puma [3.12.3](https://rubygems.org/gems/puma/versions/3.12.3) and [4.3.2](https://rubygems.org/gems/puma/versions/4.3.2) were yanked. 

I think because they maybe didn't successfully patch the vulnerability? Maybe added new bugs too? Hard to say from the puma [History.md](https://github.com/puma/puma/blob/master/History.md#433-and-3124--2020-02-28). This seems to be the [relevant PR](https://github.com/puma/puma/pull/2136), referncing [this issue](https://github.com/puma/puma/issues/2132) I don't entirely understand what happened; maybe 3.12.3/4.3.2 failed to patch the vuln, and introduced a fairly severe new bug?

So, anyway, just update to puma 3.12.4 or 4.3.3, and you're good. 

Yanking a gem is really disruptive and confusing, but sometimes has to be done. Perhaps this one of those times.

It's interesting how dependabot in the mix can make things even more confusing.
## [4][David Bryant Copeland, Author of Sustainable Web Development with Ruby on Rails - The Rails with Jason Podcast](https://www.reddit.com/r/ruby/comments/fd16pk/david_bryant_copeland_author_of_sustainable_web/)
- url: https://www.codewithjason.com/rails-with-jason-podcast/david-bryant-copeland/
---

## [5][Ruby 2.7 introduces numbered parameters as default block parameters](https://www.reddit.com/r/ruby/comments/fcwvgg/ruby_27_introduces_numbered_parameters_as_default/)
- url: https://blog.bigbinary.com/2020/03/03/ruby-2-7-introduces-numbered-parameters-as-default-block-parameters.html
---

## [6][Ruby websocket server + event loop](https://www.reddit.com/r/ruby/comments/fd0co2/ruby_websocket_server_event_loop/)
- url: https://www.reddit.com/r/ruby/comments/fd0co2/ruby_websocket_server_event_loop/
---
Hello my ruby friends!  
We need to build a simple websocket server for our internal use and started experimenting with [https://github.com/igrigorik/em-websocket](https://github.com/igrigorik/em-websocket) which has a lot of stars but last release was on Apr 23, 2014 which made me worried. 

My questions are:

1. If you want to build websocket server on ruby what's the gem to go now? (I am aware about performance issues, it's not the case for us). falcon + async-websocket is a good option?
2. For some reason when I run our prototype with EventMachine + em-websocket in docker locally it stuck on [`EM.run`](https://EM.run) line but works fine if I pass `-t` option to allocate virtual pseudo terminal `docker run -t`. Since I want to run it with AWS ECS I can't pass that option there, is there any work around?
## [7][SpreadsheetArchitect v4.0.0 is now released!](https://www.reddit.com/r/ruby/comments/fcyoue/spreadsheetarchitect_v400_is_now_released/)
- url: https://www.reddit.com/r/ruby/comments/fcyoue/spreadsheetarchitect_v400_is_now_released/
---
Spreadsheet Architect v4.0.0 has now been released. Changes include switching to Caxlsx gem, Ruby 2.3+ required, and XLSX freeze support. Check the changelog for more details. [https://github.com/westonganger/spreadsheet\_architect](https://github.com/westonganger/spreadsheet_architect)
## [8][Optimizing full-text search with Postgres materialized view in Rails](https://www.reddit.com/r/ruby/comments/fctdqz/optimizing_fulltext_search_with_postgres/)
- url: https://caspg.com/blog/optimizing-full-text-search-with-postgres-materialized-view-in-rails
---

## [9][geo-info - A dead simple reverse geocoding API](https://www.reddit.com/r/ruby/comments/fd1dmn/geoinfo_a_dead_simple_reverse_geocoding_api/)
- url: https://geo-info.co
---

## [10][Tip: A great way to learn about Ruby - join a local meetup - what's your local meetup?](https://www.reddit.com/r/ruby/comments/fd0j46/tip_a_great_way_to_learn_about_ruby_join_a_local/)
- url: https://www.reddit.com/r/ruby/comments/fd0j46/tip_a_great_way_to_learn_about_ruby_join_a_local/
---
Hello, a great way to keep learning about ruby is joining a local meetup. What's your local? Any meetups around? Any comments welcome. 

PS: Greetings from Vienna, Austria. We host an informal meetup / stammtisch once a month in a coffee house, see &lt;https://viennarb.github.io/&gt;. What's your format? What's your preference - more talks, less coding or just burger &amp; beer?
