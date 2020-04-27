# ruby
## [1][Just released sequel-activerecord-adapter gem](https://www.reddit.com/r/ruby/comments/g8z5iv/just_released_sequelactiverecordadapter_gem/)
- url: https://www.reddit.com/r/ruby/comments/g8z5iv/just_released_sequelactiverecordadapter_gem/
---
As some of you might know, I'm a big advocate for the [Sequel](https://github.com/jeremyevans/sequel) database library. However, it can be difficult to use it an app that's already using ActiveRecord, because Sequel requires its own database connection.

So I've created the [sequel-activerecord-adapter](https://github.com/janko/sequel-activerecord-adapter) gem that allows Sequel to reuse the existing ActiveRecord connection for its database interaction.

    require "sequel-activerecord-adapter"

    ActiveRecord::Base.establish_connection("postgresql:///mydb")

    DB = Sequel.activerecord # reuses the ActiveRecord connection
   
    DB.tables #=&gt; [:users, :articles, ...]

    DB[:articles].insert(title: "Sequel ActiveRecord Adapter")
    DB[:articles].all #=&gt; [{ title: "Sequel ActiveRecord Adapter" }]
    DB[:articles].update(author_id: 2)

This means you can now easily use Sequel in parts of your app that might be more performance-sensitive or require more advanced database queries that ActiveRecord doesn't support.

It also means you can use libraries like [Rodauth](https://github.com/jeremyevans/rodauth) that use Sequel under-the-hood without increased performance and mental overhead that would come with separate database connections.
## [2][Idiomatic way to check if object responds to method, otherwise return blank string](https://www.reddit.com/r/ruby/comments/g8zydx/idiomatic_way_to_check_if_object_responds_to/)
- url: https://www.reddit.com/r/ruby/comments/g8zydx/idiomatic_way_to_check_if_object_responds_to/
---
I have a web scraper that's parses rows, and sometimes it cannot find elements and throws an exception.  I've added a check to see if the query returns nill object, and if so return a blank string instead.

I was curious if there's a better way to do this in general or more idiomatic way in Ruby to write this.

&amp;#x200B;

    def parse_single_post(row)
    title     =  row.at_css('.topictitle')
    replies   =  row.at_css('.replies')
    views     =  row.at_css('.posts')
    last_post =  row.at_css('.lastpost')
    author    =  row.at_css('.username')
    link      =  row.at_css('a.topictitle')

    posts &lt;&lt; OpenStruct.new(
      title:     title ? title.text : "",
      link:      link ? link['href'] : "",
      author:    author ? author.text : "",
      last_post: last_post ? last_post.text : "",
      replies:   replies ? replies.text : "",
      views:     views ? views.text : ""
    )
 end
## [3][Proper objects for CSV headers and rows, convert column values, filter columns and rows, small(-ish) perfomance overhead, no dependencies other than Ruby stdlib.](https://www.reddit.com/r/ruby/comments/g8o50r/proper_objects_for_csv_headers_and_rows_convert/)
- url: https://github.com/buren/honey_format
---

## [4][Looking for a Ruby content writer](https://www.reddit.com/r/ruby/comments/g8xi30/looking_for_a_ruby_content_writer/)
- url: https://www.reddit.com/r/ruby/comments/g8xi30/looking_for_a_ruby_content_writer/
---
Hey everyone,   


I'm looking for a freelance Ruby content writer to write one or more blog posts about the Ruby web scraping ecosystem for my company blog. 

Ideally, I'd like the article to be about to be an introduction to web scraping with Ruby, starting with "low-level" things like HTTP clients and HTML parsing gems, up to Selenium and frameworks that handle many things. 

\- HTTP clients like HTTParty / Faraday

\- Nokogiri

\- Selenium

\- "Scrapy-like" frameworks such as Kimurai

Do you have recommendations about where I can find someone for this?   


Cheers
## [5][Alt::BrightonRuby — A slightly odd, quasi-conference for strange times](https://www.reddit.com/r/ruby/comments/g8ey7w/altbrightonruby_a_slightly_odd_quasiconference/)
- url: https://www.reddit.com/r/ruby/comments/g8ey7w/altbrightonruby_a_slightly_odd_quasiconference/
---
I've just launched the online version of my normally in-person conference. I'll miss the 'regular' version, but I think this might be fun.

[https://alt.brightonruby.com](https://alt.brightonruby.com)

There's a video here that gives a bit more detail of what it is.

[https://www.youtube.com/watch?v=oXpfMGNFFYs](https://www.youtube.com/watch?v=oXpfMGNFFYs)

If you've got any questions, happy to answer them.
## [6][Is rubygems download real?](https://www.reddit.com/r/ruby/comments/g8l5vn/is_rubygems_download_real/)
- url: https://www.reddit.com/r/ruby/comments/g8l5vn/is_rubygems_download_real/
---
Hello guys I hope you are doing fine.

In the past week I had some of my projects wrapped in Gem and sent them to [rubygems.org](https://rubygems.org).

The odd is that this gems are very specific. The functions that I created are for a very specific job and are only applied to a specific geographical region.  The first gem got instantly 107 downloads so i thought that this might be some crawlers or other repositories pulling from rubygem but the second one that I made has now double the downloads of the first. 

So is the rubygems download counter valid?
## [7][PSA: net-sftp v3.0.0.rc1 is available for testing](https://www.reddit.com/r/ruby/comments/g8aj5m/psa_netsftp_v300rc1_is_available_for_testing/)
- url: https://www.reddit.com/r/ruby/comments/g8aj5m/psa_netsftp_v300rc1_is_available_for_testing/
---
"net-sftp" last official release was in 2013. The current maintainer (Miklós Fazekas) has issued a release candidate today.

I believe it is important for a language like Ruby to have a well maintained library for this protocol (which is widely used in mid/large companies for low-tech data exchange).

If you have some use of it, please go test it and provide feedback!

[https://github.com/net-ssh/net-sftp/issues/92](https://github.com/net-ssh/net-sftp/issues/92)

Have a safe day.
## [8][Shouldn't it show nothing because I did not print anything?](https://www.reddit.com/r/ruby/comments/g8e2rq/shouldnt_it_show_nothing_because_i_did_not_print/)
- url: https://i.redd.it/qusc8o5yv5v41.jpg
---

## [9][RoR JWT Encryption from Scratch](https://www.reddit.com/r/ruby/comments/g8iwpq/ror_jwt_encryption_from_scratch/)
- url: https://link.medium.com/i4VbElle05
---

## [10][Gmail API. Error 400: redirect_uri_mismatch](https://www.reddit.com/r/ruby/comments/g8hiz6/gmail_api_error_400_redirect_uri_mismatch/)
- url: https://www.reddit.com/r/ruby/comments/g8hiz6/gmail_api_error_400_redirect_uri_mismatch/
---
Hey,

So I was following this guide for Gmail API, but keep on getting 'Error 400: redirect\_uri\_mismatch'  when trying to authorize. Was unable to find any info googling about this. any thoughts? :)

https://imgur.com/D9JCsxq
