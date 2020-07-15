# ruby
## [1][System of a test: Proper browser testing in Ruby on Rails. Ditch Selenium for Ferrum and CDP, embrace Docker.](https://www.reddit.com/r/ruby/comments/hr2fyl/system_of_a_test_proper_browser_testing_in_ruby/)
- url: https://evilmartians.com/chronicles/system-of-a-test-setting-up-end-to-end-rails-testing
---

## [2][Experiences using slack-ruby-client?](https://www.reddit.com/r/ruby/comments/hra3zb/experiences_using_slackrubyclient/)
- url: https://www.reddit.com/r/ruby/comments/hra3zb/experiences_using_slackrubyclient/
---
Evening all,

I'm going to be building a Slack integration between my SaaS Rails app and our customer's Slack channels. Initially it's going to be basic slash command but if anticipate making it much more interactive over time.

In the past I've done a basic Slack posting without any special gems / library. I recently came across [**slack-ruby-client**](https://github.com/slack-ruby/slack-ruby-client) and it certainly looks fully featured and potentially useful.

I'm interested to hear if anyone has used this library and how was their experience? 

thanks!
## [3][Why do people think programming is so difficult?](https://www.reddit.com/r/ruby/comments/hrl7ow/why_do_people_think_programming_is_so_difficult/)
- url: https://www.reddit.com/r/ruby/comments/hrl7ow/why_do_people_think_programming_is_so_difficult/
---
For a presentation, I'm searching why the people think programming looks so difficult and I wonder why they think like that. Any opinions? thanks!
## [4][Rails 6.1 adds support for where with a comparison operator](https://www.reddit.com/r/ruby/comments/hqwq96/rails_61_adds_support_for_where_with_a_comparison/)
- url: https://blog.bigbinary.com/2020/07/14/rails-6-1-adds-support-for-where-with-comparison-operator.html
---

## [5][Value objects - a complete guide to Ruby code that is testable, readable and simple](https://www.reddit.com/r/ruby/comments/hqz2p0/value_objects_a_complete_guide_to_ruby_code_that/)
- url: https://pdabrowski.com/articles/rails-design-patterns-value-object
---

## [6][[notes only] ActiveStorage Direct Upload to Azure Blob when Ruby or Rails used as JSON API only](https://www.reddit.com/r/ruby/comments/hqxsk7/notes_only_activestorage_direct_upload_to_azure/)
- url: https://github.com/equivalent/scrapbook2/blob/master/active-storage_direct-upload_api-investigation.md
---

## [7][Need suggestions for CMS/Rails API](https://www.reddit.com/r/ruby/comments/hqp8ay/need_suggestions_for_cmsrails_api/)
- url: https://www.reddit.com/r/ruby/comments/hqp8ay/need_suggestions_for_cmsrails_api/
---
Hi all!

So I've got a little bit of experience with ruby now but never had to integrate CMS into it or anything like that. I'm trying to make a news website and was originally hoping to make a rails API, with a javascript/react front end that would allow an admins/owners of the website to just add new articles, photos, headlines, sports scores, etc to the API via the front end, in which the react front end would then fetch from the DB and just keep posting it in the template of the website.

I was wondering about any kinds of CMS for ruby such as Refinery that anyone has experience using. I only say refinery because that's the first one I read about a bit and it seems pretty straight forward but I want to know if I can still use it with a javascript/react front end or if I should keep the front end Rails? 

The website should be pretty straightforward, nothing too crazy just sorting news by categories and having a small editorial page. I'm making this for a few reasons, first it seems super interesting to work with something new like this for me when I'm used to just front end dev, and two I think its a great project experience for my resume as I'm a recent grad.

Any advice is much appreciated!
## [8][What should I use to get started with broadcasting services on the local network with .local domains? Like how Raspberry Pi uses a .local domain to advertise itself on the network.](https://www.reddit.com/r/ruby/comments/hqp21d/what_should_i_use_to_get_started_with/)
- url: https://www.reddit.com/r/ruby/comments/hqp21d/what_should_i_use_to_get_started_with/
---
This is the type of thing I'm talking about.

https://www.howtogeek.com/167190/how-and-why-to-assign-the-.local-domain-to-your-raspberry-pi/


It says it uses mDNS and so after looking into that I found the gem called dnssd, ( http://docs.seattlerb.org/dnssd/) but when trying to install it on my ubuntu system I get this error, https://pastebin.com/raw/vs5EVQYw (which I thought I solved already by installing ruby via rvm and not from the repos like I was told. So I'm stuck. What should I try now?
## [9][Loading Ruby Scripts from the JVM Classpath with JRuby](https://www.reddit.com/r/ruby/comments/hqjnst/loading_ruby_scripts_from_the_jvm_classpath_with/)
- url: https://github.com/jruby/jruby/wiki/Loading-Scripts-From-Classpath
---

## [10][Looking for some kind of "quick" database implementation](https://www.reddit.com/r/ruby/comments/hqf3wb/looking_for_some_kind_of_quick_database/)
- url: https://www.reddit.com/r/ruby/comments/hqf3wb/looking_for_some_kind_of_quick_database/
---
Hey everyone. Some quick background: I lead a QA team full of ruby-ists where we build and maintain end to end tests in Webdriver and Capybara. We also have a set of "monitoring" scripts that are run via Cron.

These monitoring scripts are run each and every minute. Pieces of the script require a snapshot of the "state" of the system (basic things like true/false flags + statuses). So we thought it would be clever to just write them to text file.

That seems to work fine, however we've added so many flags that it's getting very messy to maintain, and the code is getting more and more difficult to understand.

I was thinking of writing a JSON representation of this information to a file, but it seems like what makes the most sense is to run a basic database implementation in order to store the 'state' of what our scripts are monitoring.

We're intimately familiar with MongoDB, but that seems really "heavy" for our needs. Is there something where I can just add a gem to our bundle and hit the ground running very quickly? Looking to avoid any new installations to our servers also.

I know that Rails has its own db implementation, but we have 0 familiarity with Rails and activerecord.
