# ruby
## [1][Release JRuby 9.2.12.0 Released · jruby/jruby · GitHub](https://www.reddit.com/r/ruby/comments/hjh3vu/release_jruby_92120_released_jrubyjruby_github/)
- url: https://github.com/jruby/jruby/releases/tag/9.2.12.0
---

## [2][AnyCable 1.0: Four years of real-time web with Ruby and Go](https://www.reddit.com/r/ruby/comments/hjb2td/anycable_10_four_years_of_realtime_web_with_ruby/)
- url: https://evilmartians.com/chronicles/anycable-1-0-four-years-of-real-time-web-with-ruby-and-go
---

## [3][Gem with complex parameters](https://www.reddit.com/r/ruby/comments/hjs5x7/gem_with_complex_parameters/)
- url: https://www.reddit.com/r/ruby/comments/hjs5x7/gem_with_complex_parameters/
---
Hi there.

Imagine that I'm creating a gem that takes care of some heavy and complex calculations.

    def compute_stuff(...)

One of the inputs that gem needs is also non trivial. The gem needs to receive a collection of objects, each one of those objects would have a few fields itself. Something like this:

    [
        {"name": "John", "location": "US", "order": {...} },
        {"name": "Jane", "location": "UK", "order": {...} }
        ...
    ]

Which strategy would you user to pass that kind of inputs to the gem's method?

1. A simple hash: just send an hash to the gem
2. Create a class mapping each object and "make" the caller use that class?
3. Other option?

&amp;#x200B;

I have a past writing c# code and I really like interfaces to be as explicit as possible, how can we try to achieve that for a gem like this one?
## [4][digest-crc 0.6.0 released. ~40x performance improvement!](https://www.reddit.com/r/ruby/comments/hjndvt/digestcrc_060_released_40x_performance_improvement/)
- url: https://www.reddit.com/r/ruby/comments/hjndvt/digestcrc_060_released_40x_performance_improvement/
---
[digest-crc](https://github.com/postmodern/digest-crc#readme) [0.6.0](https://github.com/postmodern/digest-crc/blob/master/ChangeLog.md#060--2020-07-01) has been released. This release introduces _optional_ C extensions, that when built on CRuby will override the pure-Ruby CRC methods with C equivalents, providing a ~40x performance improvement. If the C extensions cannot be built (no compiler installed, no `ruby.h` headers, not using CRuby) digest-crc will fallback to using the pure-Ruby CRC methods.

**Note:** If you use [ruby-kafka](https://github.com/zendesk/ruby-kafka#readme), run `bundle update digest-crc` to take advantage of the improved performance (see: https://github.com/zendesk/ruby-kafka/issues/620).
## [5][Provide a way for methods to omit their return value (rev.2) by shyouhei](https://www.reddit.com/r/ruby/comments/hjg5hr/provide_a_way_for_methods_to_omit_their_return/)
- url: https://github.com/ruby/ruby/pull/3271
---

## [6][Rails Architects Conference is happening this week](https://www.reddit.com/r/ruby/comments/hjf9yt/rails_architects_conference_is_happening_this_week/)
- url: https://www.reddit.com/r/ruby/comments/hjf9yt/rails_architects_conference_is_happening_this_week/
---
Hello,

Yesterday we have started an online event called Rails Architects Conference. It's a 7 days even consisting of 6 talks, related to big Rails apps.

It started yesterday with my talk titled "This time it will be different - how to properly start your next Rails app", where I have presented an event-driven approach to starting new Rails apps.

The recording (already close to 3k YT views) and the agenda are available on [https://railsarchitects.com/conference/](https://railsarchitects.com/conference/)

Today (Wed) is the Tomasz Wróbel talk on Rails multitenant apps based on PostgreSQL schemas.

Disclaimer1: The conference is fully free, but after the talks we present the offer to join our 13-weeks online course called the Rails Architect Masterclass 4.0 - [https://railsarchitects.com](https://railsarchitects.com) (price: $799)

Disclaimer2: Subscribing to the conference means subscribing to the Arkency Ruby Newsletter, feel free to unsubscribe after the conference is over. 

The current agenda:

[**This time it will be different - how to properly start your next Rails app**](https://www.youtube.com/watch?v=kURp3CE-FvM)  
Tuesday, 30 June 17:00 UTC by [Andrzej Krzywda](https://twitter.com/andrzejkrzywda)  
**Multitenancy in Rails apps: PostgreSQL schemas**  
Wednesday, 1 July 20:00 UTC by [Tomasz Wróbel](https://twitter.com/tomasz_wro)  
**Painless Rails upgrades**  
Thursday, 2 July 20:00 UTC by [Szymon Fiedler](https://twitter.com/szymonfiedler)  
**Simplify and speed up your Rails views**  
Friday, 3 July 11:00 UTC by [Mirosław Pragłowski](https://twitter.com/mpraglowski)  
**Your Ruby code is pretty but often has low cohesion and high coupling** \- live code review,   
Saturday, 4 July 10:00 UTC by Tomasz Stolarczyk and Andrzej Krzywda  
**TBA**  
Monday, 6 July 18:00 UTC by [Andrzej Krzywda](https://twitter.com/andrzejkrzywda)
## [7][I built an API for generating PDF documents from HTML or templates](https://www.reddit.com/r/ruby/comments/hji5wv/i_built_an_api_for_generating_pdf_documents_from/)
- url: https://docamatic.com
---

## [8][An important factor in choosing the name Ruby](https://www.reddit.com/r/ruby/comments/hjcqg1/an_important_factor_in_choosing_the_name_ruby/)
- url: https://twitter.com/RubyCademy/status/1278349786224250880
---

## [9][Data Analyst new to Ruby - our devs at work need me to use Sidekiq APIs to provide visibility on the number of retries at any given point, as well as breakdown on the type of retries.](https://www.reddit.com/r/ruby/comments/hioh4x/data_analyst_new_to_ruby_our_devs_at_work_need_me/)
- url: https://www.reddit.com/r/ruby/comments/hioh4x/data_analyst_new_to_ruby_our_devs_at_work_need_me/
---
Hi Ruby community!

Virtually all of our company's products are built on rails. I understand that Ruby uses a cron scheduler called sidekiq that handles tries and retries. 

This is all unfamiliar territory for me because I'm not familiar with Ruby. I'm familiar however with other tools like Python. I'm curious as to what the best way of doing this would be. Is there an API I can query that would serve this exact info?
## [10][Rails 6.1 tracks Active Storage variant in the database](https://www.reddit.com/r/ruby/comments/hii694/rails_61_tracks_active_storage_variant_in_the/)
- url: https://blog.bigbinary.com/2020/06/30/rails-6-1-tracks-active-storage-variant-in-the-database.html
---

