# ruby
## [1][Rodauth 2.0 and rodauth-rails released](https://www.reddit.com/r/ruby/comments/gf32hn/rodauth_20_and_rodauthrails_released/)
- url: https://www.reddit.com/r/ruby/comments/gf32hn/rodauth_20_and_rodauthrails_released/
---
Jeremy Evans has recently released Rodauth 2.0, which revamps the multifactor authentication flow, adds [active sessions](http://rodauth.jeremyevans.net/rdoc/files/doc/active_sessions_rdoc.html) and [audit logging](http://rodauth.jeremyevans.net/rdoc/files/doc/audit_logging_rdoc.html) features, and brings numerous other improvements. See the [release notes](http://rodauth.jeremyevans.net/rdoc/files/doc/release_notes/2_0_0_txt.html) for the full list of changes.

For those who are not familiar, **[Rodauth](https://github.com/jeremyevans/rodauth/)** is a full-featured Rack-based authentication framework built on top of Roda &amp; Sequel, but usable in any web framework. It's an alternative to Devise, Sorcery, Clearance, Authlogic etc, with the following features that make it stand out for me:

* amount of features built in, including multifactor authentication (TOTP, SMS codes, recovery codes, [WebAuthn](https://webauthn.io/)) and email authentication (aka "passwordless")
* JSON API support with [JWT](http://rodauth.jeremyevans.net/rdoc/files/doc/jwt_rdoc.html)
* [advanced security features](http://rodauth.jeremyevans.net/rdoc/files/README_rdoc.html#label-Security), such as ability to protect password hashes even in case of SQL injection (using 2 database accounts), and including account id in tokens allowing brute-force attempts only for a single account
* features are contained in a [single file](https://github.com/jeremyevans/rodauth/blob/master/lib/rodauth/features/remember.rb), instead of being [spread](https://github.com/heartcombo/devise/blob/master/lib/devise/models/rememberable.rb) [across](https://github.com/heartcombo/devise/blob/master/lib/devise/controllers/rememberable.rb) [many](https://github.com/heartcombo/devise/blob/master/lib/devise/strategies/rememberable.rb) [different](https://github.com/heartcombo/devise/blob/master/lib/devise/hooks/rememberable.rb) [places](https://github.com/heartcombo/devise/blob/master/lib/devise/hooks/forgetable.rb)
* authentication behaviour is configured in a one place (your Rodauth app), and each setting can be configured statically or dynamically based on request context and account record

In order to bring Rodauth closer to the Rails community, I've created the **[rodauth-rails](https://github.com/janko/rodauth-rails)** gem, which provides the Rails glue that I needed for my own Rails app at work (see the [demo app](https://github.com/janko/rodauth-demo-rails/)). It brings the following features:

* generators for Active Record migration, views and emails
* configures Sequel to [reuse Active Record's connection](https://github.com/janko/sequel-activerecord_connection)
* template rendering with Action Controller &amp; Action View
* email creation with Action Mailer
* integration with Rails' CSRF protection and flash messaging
* easier set of Rodauth defaults and other niceties

On a personal note, Rodauth is one of these projects that keep me genuinely interested in web development with Ruby. I'm thoroughly impressed by its design, and I feel like contributing to it has made me grow as a developer. I'm curious to hear your thoughts :)
## [2][My notes from Aaron Patterson's RailsConf 2020 keynote](https://www.reddit.com/r/ruby/comments/gexsq2/my_notes_from_aaron_pattersons_railsconf_2020/)
- url: https://www.reddit.com/r/ruby/comments/gexsq2/my_notes_from_aaron_pattersons_railsconf_2020/
---
Aaron's talk was very Aaron: 20 minutes of jokes, 40 minutes of brain bending Rails performance show-and-tell. His goal: to teach you how to profile your Rails app. If you haven't watched his keynote yet, here's the gist:

https://www.joshuawood.net/notes/2020-railsconf-aaron-patterson-keynote

A few takeaways:
- Not all code is equal
- Similar behaviors should have similar performance
- **Poor performance is a bug**

Miss Aaron's talk? In THIS economy?

👌
## [3][Understanding complex Ruby application with profiling tools](https://www.reddit.com/r/ruby/comments/genzmb/understanding_complex_ruby_application_with/)
- url: https://katafrakt.me/2020/05/03/understanding-ruby-app-with-profiling/
---

## [4][Using Interactors in Rails](https://www.reddit.com/r/ruby/comments/gf24fk/using_interactors_in_rails/)
- url: https://blog.saeloun.com/2020/05/06/rails-using-interactor-gem
---

## [5][For people like me who are into listening music while coding... [off]](https://www.reddit.com/r/ruby/comments/geqz8q/for_people_like_me_who_are_into_listening_music/)
- url: https://www.reddit.com/r/ruby/comments/geqz8q/for_people_like_me_who_are_into_listening_music/
---
... check out the [list](https://spoti.fi/3cPiDAs) I use when I’m coding: 8+ hours of retro synth music inspired by Stranger Things.

Also curious to know what kind of stuff do you listen to, if any.
## [6][Custom scalar in GraphQL-ruby](https://www.reddit.com/r/ruby/comments/geqjx5/custom_scalar_in_graphqlruby/)
- url: https://www.abhaynikam.me/posts/custom-scalar-in-graphql-ruby/
---

## [7][Anybody else dislike the idea of "endless" methods?](https://www.reddit.com/r/ruby/comments/geku74/anybody_else_dislike_the_idea_of_endless_methods/)
- url: https://www.reddit.com/r/ruby/comments/geku74/anybody_else_dislike_the_idea_of_endless_methods/
---
I'll be perfectly up front here in my opinion that we Rubyists are far too obsessed with keeping our code as "minimalist" as possible.  Far too much reliance on implied operators and a focus bordering on obsession on using method chains pack a ton of functionality into a small space.  As someone who has studied psychology extensively I find that a lot of what is in the style guides makes code harder to parse by a human and thus harder to work with.

And now there is a proposal to add "endless" methods into Ruby, so you can stick single line method literally onto one or two lines and dispense with the "end" keyword.  I'm sorry, but from my perspective this is a bad idea.  Is requiring a clear demarcation of the end of a method really such a bad thing?  These statements make it very clear where something terminates, and the next person reading the code doesn't have to spend mental energy figuring it out.  And yeah, for one liners it's rather trivial, but will this be where it stops?  What about multi-line methods down the road?  No disrespect to Python users, but we're not writing Python here and delineating code blocks using white space isn't what we signed up for.  Coffeescript used that philosophy and I think we're all glad CS is going away.  Do we really need another method definition syntax to consider?

Understand I'm not going to tell anyone else that they're right or wrong here, this is just my opinion on this change.  If there are benefits here I'm not seeing I'm happy to consider them, but from what I see so far this looks like trying to fix what ain't broken.  If you're going to up or downvote that's fine but let me know your reasons why, I'd like to hear other peoples thoughts, and thank you for listening to mine.

Edit:  From reading the comments and perspective of others I'm more comfortable with this as it seems that it won't be that big of a change.  Thank you to all for your input.
## [8][I'm so new and bad at this it's not even funny.](https://www.reddit.com/r/ruby/comments/geyitf/im_so_new_and_bad_at_this_its_not_even_funny/)
- url: https://www.reddit.com/r/ruby/comments/geyitf/im_so_new_and_bad_at_this_its_not_even_funny/
---
Hello, I am a student studying to become a networking technician and have fallen upon Programming and a huge issue where I am at the end of my studies, and because of complications have to create a pong game that uses gosu, yaml comfigs and UDPSocket for x,y cooridnates data. This is of course in ruby and I'm looking around and it doesn't seem that complicated all though I am having minor panic attacks..

Due to my inexperience in programming I need to re-learn and remember how even the basics works and I'm looking at: [https://docs.ruby-lang.org/en/2.2.0/UDPSocket.html](https://docs.ruby-lang.org/en/2.2.0/UDPSocket.html)  


And want to use it, test it out in like irb. But I honestly don't know how to even get it to start working since gem install doesn't work? I'm sorry I've got three weeks to make something out of this in an OOP manner and I've forgotten even the basics of how to use stuff like this and I don't know where to ask or check.
## [9][Has anyone here used DRb in the last few years? What was your experience?](https://www.reddit.com/r/ruby/comments/geihqa/has_anyone_here_used_drb_in_the_last_few_years/)
- url: https://www.reddit.com/r/ruby/comments/geihqa/has_anyone_here_used_drb_in_the_last_few_years/
---
I ended up browsing DRb's docs today. There isn't much written about it. Hence the curiosity whether anybody is using it.

I need a kind of micro-service that's part of the Rails app. A micro-service in the light that all puma/web-server/sidekiq instances talk to the same Class. I can imagine something like Redis but being part of Rails that I don't have to provision as a separate service. That would be straight-forward in Elixir/Phoenix, but there isn't a clear path in Ruby/Rails.

That's how I ended up reading about DRb... this obscure ruby tech.
## [10][Crystal yay or nay?](https://www.reddit.com/r/ruby/comments/ge4h03/crystal_yay_or_nay/)
- url: https://www.reddit.com/r/ruby/comments/ge4h03/crystal_yay_or_nay/
---
Just wondering what my fellow rubyists think of Crystal.
Do you see some potential in Crystal or you very much prefer the way Ruby works?

Crystal is young and full of flaws but that does not mean those cant be fixed.
