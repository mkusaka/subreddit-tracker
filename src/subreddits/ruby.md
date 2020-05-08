# ruby
## [1][Rails team wants to know (really!) about frustrations and roadblocks you have with Rails](https://www.reddit.com/r/ruby/comments/gfdbxk/rails_team_wants_to_know_really_about/)
- url: https://weblog.rubyonrails.org/2020/5/7/A-May-of-WTFs/
---

## [2][Interesting behavior of Hash.new](https://www.reddit.com/r/ruby/comments/gfoc59/interesting_behavior_of_hashnew/)
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
## [3][Ruby Association -&gt; 2019 Grant Accomplishment Report](https://www.reddit.com/r/ruby/comments/gfo8bh/ruby_association_2019_grant_accomplishment_report/)
- url: https://www.ruby.or.jp/en/news/20200508
---

## [4][Gems like laravel's live wire for rails?](https://www.reddit.com/r/ruby/comments/gfrjwx/gems_like_laravels_live_wire_for_rails/)
- url: /r/rubyonrails/comments/gfr7qh/any_alternatives_with_httpslaravellivewirecom/
---

## [5][Rodauth 2.0 and rodauth-rails released](https://www.reddit.com/r/ruby/comments/gf32hn/rodauth_20_and_rodauthrails_released/)
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
## [6][Intro to JRuby article and interview with project co-lead](https://www.reddit.com/r/ruby/comments/gfd0ts/intro_to_jruby_article_and_interview_with_project/)
- url: https://www.hostingadvice.com/blog/charles-nutter-on-jruby/
---

## [7][Anyone using Best Buy APIs?](https://www.reddit.com/r/ruby/comments/gfb8ez/anyone_using_best_buy_apis/)
- url: https://github.com/rootstrap/best_buy_ruby/
---

## [8][Ruby Gem for Speech recognition](https://www.reddit.com/r/ruby/comments/gfa38l/ruby_gem_for_speech_recognition/)
- url: https://www.reddit.com/r/ruby/comments/gfa38l/ruby_gem_for_speech_recognition/
---
I have a ruby script that runs on the CLI.

Accepts several commands via Standard Input.

I want to make some of the commands executable via voice commands.

Which speech recognition gem do you recommend?

&amp;#x200B;

It looks like there are a few but they are all pretty dated.
## [9][My notes from Aaron Patterson's RailsConf 2020 keynote](https://www.reddit.com/r/ruby/comments/gexsq2/my_notes_from_aaron_pattersons_railsconf_2020/)
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
## [10][Improve my mental model of metaclasses](https://www.reddit.com/r/ruby/comments/gf7nbm/improve_my_mental_model_of_metaclasses/)
- url: https://www.reddit.com/r/ruby/comments/gf7nbm/improve_my_mental_model_of_metaclasses/
---
Here’s my mental model.

It’s probably not completely correct. (i’m still kinda confused)

These are all equivalent:

    class A
      def greet
        puts 'Hi'
      end
    end
    
    class A
      self.define_method(:greet) do
        puts 'Hi'
      end
    end
    
    class A
      A.define_method(:greet) do
        puts 'Hi'
      end
    end
    
    class A;end
    A.define_method(:greet) do
      puts 'Hi'
    end

`A` is an object

`A` has some internal blueprint for making new copies (instances) of itself

`define_method` says: add this method to the blueprint `A` is keeping track of

&amp;#x200B;

this special syntax of `def`

    def self.greet
      puts 'Hello'
    end

translates to

    self's_meta_class.define_method(:greet) do
      puts 'Hello'
    end

an example:

    class A
      def self.greet
        'Hello'
      end
    end

is a shorthand for:

    class A
      a_metaclass = class &lt;&lt; self; self; end
    
      a_metaclass.define_method(:greet) do
        'Hello'
      end
    end

`A's_metaclass` is an object

we added `greet` to it’s internal *make new copies of myself* blueprint

**but**, a metaclass can’t make new instances of itself

\--despite this, `A` the object *is* an instance of `A's_metaclass`

if you do `A.ancestors` you get:

    =&gt; [A, Object, Kernel, BasicObject]

but i think the metaclass sits magically in front

    =&gt; [#&lt;Class:A&gt;, A, Object, Kernel, BasicObject]

&amp;#x200B;

before I couldn't figure out why this wouldn't work

    class Dog;end
    
    def Dog.bark
      puts 'Bark!'
    end
    
    
    dog_metaclass = class &lt;&lt; Dog; self; end
    
    dog_metaclass.bark # exception

If you take what I said earlier, then when we did

    def Dog.bark
      puts 'Bark!'
    end

that was acting on the metaclass.

effectively:

    dog_metaclass.define_method(:bark) do
      puts 'Bark!'
    end

but! remember -- `define_method` says, add this to your new instance blueprint

so `dog_metaclass.bark` doesn't work

but `dog_metaclass.new.bark` would work (if a metaclass could do new)

&amp;#x200B;

so `bark` *is* on the `dog_metaclass` object, it's just in `dog_metaclass`'s new instance blueprint

    dog_metaclass.instance_methods.first # :bark

and, because the object `Dog` is effectively an instance created by `dog_metaclass`, `Dog.bark` works
