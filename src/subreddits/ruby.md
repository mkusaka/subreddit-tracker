# ruby
## [1][What GUI library would you recommend for a simple to do ruby app?](https://www.reddit.com/r/ruby/comments/gg5ff5/what_gui_library_would_you_recommend_for_a_simple/)
- url: https://www.reddit.com/r/ruby/comments/gg5ff5/what_gui_library_would_you_recommend_for_a_simple/
---
I am not using Ruby on Rails as it will be just for pc
## [2][Is Ruby a good choice here?](https://www.reddit.com/r/ruby/comments/gg7774/is_ruby_a_good_choice_here/)
- url: https://www.reddit.com/r/ruby/comments/gg7774/is_ruby_a_good_choice_here/
---
I need to automate gathering certain data from Microsoft's reference pages and turn that into a neat C++ wrapper for parts of WinAPI. I could do it manually but at the current pace it would take me a better part of this year to finish it so I think it's time for me to learn how to work with the internet for once.

I've heard that Ruby is much better suited for dealing with this type of web stuff than Python. Is that true? If so, are there any tutorials from the "X for Y programmers" category? (I mean the ones that are at least serviceable, I tend to die a little bit inside every time someone explains to me what a variable is).
## [3][Looping thru Nested Arrays](https://www.reddit.com/r/ruby/comments/gg4lnf/looping_thru_nested_arrays/)
- url: https://www.reddit.com/r/ruby/comments/gg4lnf/looping_thru_nested_arrays/
---
Hi! Really new at this. About to start a bootcamp and need to complete some pre-work.

My current task is as follows: I have a nested array, all with integers, and **I must loop through** each array to iterate and return an array of the smallest numbers of each array.

I know of the map method, but have no idea how to use it. We haven't covered it yet.

Ideas?

**UPDATE:** solved! Thank you everyone for being so dope and helping me. I really appreciate it :)  
I ended up being stuck in an infinite loop. Had to do some screen sharing to figure it out. 
## [4][I'm supposed to find someone who has 10 years experience consistently using Chef - I'm wondering if that is that even possible? Do people even use it anymore and could they have used it that long?](https://www.reddit.com/r/ruby/comments/gfv7d3/im_supposed_to_find_someone_who_has_10_years/)
- url: https://www.reddit.com/r/ruby/comments/gfv7d3/im_supposed_to_find_someone_who_has_10_years/
---

## [5][Ruby Association -&gt; 2019 Grant Accomplishment Report](https://www.reddit.com/r/ruby/comments/gfo8bh/ruby_association_2019_grant_accomplishment_report/)
- url: https://www.ruby.or.jp/en/news/20200508
---

## [6][Interesting behavior of Hash.new](https://www.reddit.com/r/ruby/comments/gfoc59/interesting_behavior_of_hashnew/)
- url: https://www.reddit.com/r/ruby/comments/gfoc59/interesting_behavior_of_hashnew/
---
Hello,

&amp;#x200B;

I was building a script to format some hashes and I had something similar to this:

    data = { monday: { start: '08:00', end: '18:00' }, thuesday: { start: '08:00', end: '15:00' }}
    
    result = data.each_with_object(Hash.new([])) do |(day, schedule), result|
      result[schedule[:start]] &lt;&lt; day
    end

Nothing fancy, but I noticed that my result variable seemed empty, but if I tried to access a key, it was there.

See the following screenshot:

&amp;#x200B;

https://preview.redd.it/6kl2xkpgqhx41.png?width=472&amp;format=png&amp;auto=webp&amp;s=77d2f3a8c166fe4ef6c1e70b99834cac670e7275

I know how to fix it, I replaced [`Hash.new`](https://Hash.new)`([])` with `Hash.new { |hash, key| hash[key] = [] }` and it works as expected.

&amp;#x200B;

But I was still curious why it happens and how come I can access the keys. Any thoughts?

&amp;#x200B;

Thanks
## [7][Rails team wants to know (really!) about frustrations and roadblocks you have with Rails](https://www.reddit.com/r/ruby/comments/gfdbxk/rails_team_wants_to_know_really_about/)
- url: https://weblog.rubyonrails.org/2020/5/7/A-May-of-WTFs/
---

## [8][Gems like laravel's live wire for rails?](https://www.reddit.com/r/ruby/comments/gfrjwx/gems_like_laravels_live_wire_for_rails/)
- url: /r/rubyonrails/comments/gfr7qh/any_alternatives_with_httpslaravellivewirecom/
---

## [9][Rodauth 2.0 and rodauth-rails released](https://www.reddit.com/r/ruby/comments/gf32hn/rodauth_20_and_rodauthrails_released/)
- url: https://www.reddit.com/r/ruby/comments/gf32hn/rodauth_20_and_rodauthrails_released/
---
Jeremy Evans has recently released Rodauth 2.0, which revamps the multifactor authentication flow, adds [active sessions](http://rodauth.jeremyevans.net/rdoc/files/doc/active_sessions_rdoc.html) and [audit logging](http://rodauth.jeremyevans.net/rdoc/files/doc/audit_logging_rdoc.html) features, and brings numerous other improvements. See the [release notes](http://rodauth.jeremyevans.net/rdoc/files/doc/release_notes/2_0_0_txt.html) for the full list of changes.

For those who are not familiar, **[Rodauth](https://github.com/jeremyevans/rodauth/)** is a full-featured Rack-based authentication framework built on top of Roda &amp; Sequel, but usable in any web framework. It's an alternative to Devise, Sorcery, Clearance, Authlogic etc, with the following features that make it stand out for me:

* amount of features built in, including multifactor authentication (TOTP, SMS codes, recovery codes, [WebAuthn](https://webauthn.io/)) and email authentication (aka "passwordless")
* JSON API support with [JWT](http://rodauth.jeremyevans.net/rdoc/files/doc/jwt_rdoc.html)
* [advanced security features](http://rodauth.jeremyevans.net/rdoc/files/README_rdoc.html#label-Security), such as ability to protect password hashes even in case of SQL injection (using 2 database accounts), and including account id in tokens allowing brute-force attempts only for a single account
* features are contained in a [single file](https://github.com/jeremyevans/rodauth/blob/master/lib/rodauth/features/remember.rb), instead of being [spread](https://github.com/heartcombo/devise/blob/master/lib/devise/models/rememberable.rb) [across](https://github.com/heartcombo/devise/blob/master/lib/devise/controllers/rememberable.rb) [many](https://github.com/heartcombo/devise/blob/master/lib/devise/strategies/rememberable.rb) [different](https://github.com/heartcombo/devise/blob/master/lib/devise/hooks/rememberable.rb) [places](https://github.com/heartcombo/devise/blob/master/lib/devise/hooks/forgetable.rb)
* authentication behaviour is configured in a one place (your Rodauth app), and each setting can be configured statically or dynamically based on request context and account record

In order to bring Rodauth closer to the Rails community, I've created **[rodauth-rails](https://github.com/janko/rodauth-rails)**, which provides the Rails glue I needed for my own Rails app at work (see the [demo app](https://github.com/janko/rodauth-demo-rails/)). It provides the following features:

* generators for Active Record migration, views and emails
* configures Sequel to [reuse Active Record's connection](https://github.com/janko/sequel-activerecord_connection)
* template rendering with Action Controller &amp; Action View
* email creation with Action Mailer
* integration with Rails' CSRF protection and flash messaging
* easier set of Rodauth defaults and other niceties

On a personal note, Rodauth is one of these projects that keep me genuinely interested in web development with Ruby. I'm thoroughly impressed by its design, and I feel like contributing to it has made me grow as a developer. I'm curious to hear your thoughts :)
## [10][Intro to JRuby article and interview with project co-lead](https://www.reddit.com/r/ruby/comments/gfd0ts/intro_to_jruby_article_and_interview_with_project/)
- url: https://www.hostingadvice.com/blog/charles-nutter-on-jruby/
---

