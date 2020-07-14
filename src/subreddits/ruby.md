# ruby
## [1][Rails 6.1 adds support for where with a comparison operator](https://www.reddit.com/r/ruby/comments/hqwq96/rails_61_adds_support_for_where_with_a_comparison/)
- url: https://blog.bigbinary.com/2020/07/14/rails-6-1-adds-support-for-where-with-comparison-operator.html
---

## [2][Handing infinite loop](https://www.reddit.com/r/ruby/comments/hr0aip/handing_infinite_loop/)
- url: https://www.reddit.com/r/ruby/comments/hr0aip/handing_infinite_loop/
---
Hello! 

I am writing a simple jumping game - the player jumps within a row of 
wooden posts on a lake. Each post has a number on it that indicates
where the player should jump to: the number itself describes the number of posts to jump over, the sign indicates the direction. On positive numbers the player jumps to the right, on negative numbers to the left. The number of posts is limited, thus it may happen to the player that he jumps into the water.

If I have an input like this (post A: 1, post B: 1, post C: -1), player would jump between the post B and post C infinitely. 

How should I handle this? I was thinking about tracking the visited posts somehow and if it has been repeated for some times, the program would stop.

Is there any other better way of handling cases like this? 

(Also, before someone mentions it, I am aware of the Halting problem :) )
## [3][Value objects - a complete guide to Ruby code that is testable, readable and simple](https://www.reddit.com/r/ruby/comments/hqz2p0/value_objects_a_complete_guide_to_ruby_code_that/)
- url: https://pdabrowski.com/articles/rails-design-patterns-value-object
---

## [4][Need suggestions for CMS/Rails API](https://www.reddit.com/r/ruby/comments/hqp8ay/need_suggestions_for_cmsrails_api/)
- url: https://www.reddit.com/r/ruby/comments/hqp8ay/need_suggestions_for_cmsrails_api/
---
Hi all!

So I've got a little bit of experience with ruby now but never had to integrate CMS into it or anything like that. I'm trying to make a news website and was originally hoping to make a rails API, with a javascript/react front end that would allow an admins/owners of the website to just add new articles, photos, headlines, sports scores, etc to the API via the front end, in which the react front end would then fetch from the DB and just keep posting it in the template of the website.

I was wondering about any kinds of CMS for ruby such as Refinery that anyone has experience using. I only say refinery because that's the first one I read about a bit and it seems pretty straight forward but I want to know if I can still use it with a javascript/react front end or if I should keep the front end Rails? 

The website should be pretty straightforward, nothing too crazy just sorting news by categories and having a small editorial page. I'm making this for a few reasons, first it seems super interesting to work with something new like this for me when I'm used to just front end dev, and two I think its a great project experience for my resume as I'm a recent grad.

Any advice is much appreciated!
## [5][Loading Ruby Scripts from the JVM Classpath with JRuby](https://www.reddit.com/r/ruby/comments/hqjnst/loading_ruby_scripts_from_the_jvm_classpath_with/)
- url: https://github.com/jruby/jruby/wiki/Loading-Scripts-From-Classpath
---

## [6][What should I use to get started with broadcasting services on the local network with .local domains? Like how Raspberry Pi uses a .local domain to advertise itself on the network.](https://www.reddit.com/r/ruby/comments/hqp21d/what_should_i_use_to_get_started_with/)
- url: https://www.reddit.com/r/ruby/comments/hqp21d/what_should_i_use_to_get_started_with/
---
This is the type of thing I'm talking about.

https://www.howtogeek.com/167190/how-and-why-to-assign-the-.local-domain-to-your-raspberry-pi/


It says it uses mDNS and so after looking into that I found the gem called dnssd, ( http://docs.seattlerb.org/dnssd/) but when trying to install it on my ubuntu system I get this error, https://pastebin.com/raw/vs5EVQYw (which I thought I solved already by installing ruby via rvm and not from the repos like I was told. So I'm stuck. What should I try now?
## [7][[notes only] ActiveStorage Direct Upload to Azure Blob when Ruby or Rails used as JSON API only](https://www.reddit.com/r/ruby/comments/hqxsk7/notes_only_activestorage_direct_upload_to_azure/)
- url: https://github.com/equivalent/scrapbook2/blob/master/active-storage_direct-upload_api-investigation.md
---

## [8][Looking for some kind of "quick" database implementation](https://www.reddit.com/r/ruby/comments/hqf3wb/looking_for_some_kind_of_quick_database/)
- url: https://www.reddit.com/r/ruby/comments/hqf3wb/looking_for_some_kind_of_quick_database/
---
Hey everyone. Some quick background: I lead a QA team full of ruby-ists where we build and maintain end to end tests in Webdriver and Capybara. We also have a set of "monitoring" scripts that are run via Cron.

These monitoring scripts are run each and every minute. Pieces of the script require a snapshot of the "state" of the system (basic things like true/false flags + statuses). So we thought it would be clever to just write them to text file.

That seems to work fine, however we've added so many flags that it's getting very messy to maintain, and the code is getting more and more difficult to understand.

I was thinking of writing a JSON representation of this information to a file, but it seems like what makes the most sense is to run a basic database implementation in order to store the 'state' of what our scripts are monitoring.

We're intimately familiar with MongoDB, but that seems really "heavy" for our needs. Is there something where I can just add a gem to our bundle and hit the ground running very quickly? Looking to avoid any new installations to our servers also.

I know that Rails has its own db implementation, but we have 0 familiarity with Rails and activerecord.
## [9][Explain the following syntax](https://www.reddit.com/r/ruby/comments/hql0w9/explain_the_following_syntax/)
- url: https://www.reddit.com/r/ruby/comments/hql0w9/explain_the_following_syntax/
---
Hi y'all,

I'm new to Ruby and just stumbled upon the following syntax being used to define a route in a Rails project:

```ruby
Rails.application.routes.draw do
    # ...
    get 'about' =&gt; 'pages#about' # &lt;= THIS
    # ...
end
```

What would that look like in, say, PHP, C or JavaScript?
## [10][99 Bottles of OOP - 2nd Edition Released](https://www.reddit.com/r/ruby/comments/hpzdxf/99_bottles_of_oop_2nd_edition_released/)
- url: https://sandimetz.com/99bottles
---

