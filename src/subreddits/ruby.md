# ruby
## [1][Promises in Ruby? Particularly in "ruby-concurrency"?](https://www.reddit.com/r/ruby/comments/jpmb10/promises_in_ruby_particularly_in_rubyconcurrency/)
- url: https://www.reddit.com/r/ruby/comments/jpmb10/promises_in_ruby_particularly_in_rubyconcurrency/
---
Gem [https://github.com/ruby-concurrency](https://github.com/ruby-concurrency)

There're some examples in the documentation, but there're so many methods and functions in **Promise**, that it's confusing how to use it at all, and which functions to use and in what cases.

&amp;#x200B;

Are there simple real examples of how to use Promises, how to chain them, wait for them to resolve and so on?

Or what gem would you recommend instead? A library that's used by many people, yet simple and decent. I don't need a gem that's abandoned or use by few people.
## [2]["What are the different kinds of Rails tests and when should I use each?"](https://www.reddit.com/r/ruby/comments/jpd20o/what_are_the_different_kinds_of_rails_tests_and/)
- url: https://www.reddit.com/r/ruby/comments/jpd20o/what_are_the_different_kinds_of_rails_tests_and/
---
I recently conducted a survey of Rails developers to ask what their biggest testing questions are.

Out of the 8 questions people voted on, "What are the different kinds of Rails tests and when should I use each?" was solidly the top question.

My answer is that you can meet 98% of your needs with just system specs and model specs. I don't really use request specs, helper specs, view specs, routing specs, mailer specs or jobs specs, for various different reasons.

I explain all my reasoning in this post:

[What are the different kinds of Rails tests and when should I use each?](https://www.codewithjason.com/different-kinds-rails-tests-use/)

I would welcome any feedback anyone has. From experienced testers, I'd especially like to know if you disagree, and if so why. From people newer to testing, I'm curious to what extent this answers your questions and if this leaves you with any questions unanswered.

Hope you find the post useful.
## [3][Rails autoloading — how it works!](https://www.reddit.com/r/ruby/comments/jpd7ya/rails_autoloading_how_it_works/)
- url: https://www.urbanautomaton.com/blog/2020/11/04/rails-autoloading-heaven/
---

## [4][Why Shopify's engineering team stuck with Ruby when rewriting their storefront](https://www.reddit.com/r/ruby/comments/jounah/why_shopifys_engineering_team_stuck_with_ruby/)
- url: https://www.youtube.com/watch?v=_Blz2mFjWis
---

## [5][Turkish language support for ruby](https://www.reddit.com/r/ruby/comments/jp1kno/turkish_language_support_for_ruby/)
- url: https://github.com/sbagdat/turkish_support
---

## [6][mltype - Typing practice for Ruby and other languages](https://www.reddit.com/r/ruby/comments/jon8bl/mltype_typing_practice_for_ruby_and_other/)
- url: https://www.reddit.com/r/ruby/comments/jon8bl/mltype_typing_practice_for_ruby_and_other/
---
**What is it?**

Command line tool that uses a character-level LSTM model to generate text that resembles a real language (including programming languages). One can both train a network from scratch or download a pretrained one (Ruby, Java, Go, JavaScript, C, C++, Python,..).

**Motivation**

I recently switched to touch typing and I realized that there is basically no way to practise typing of programming languages (other than actually programming). Additionally, I revisited the famous blog post [http://karpathy.github.io/2015/05/21/rnn-effectiveness/](http://karpathy.github.io/2015/05/21/rnn-effectiveness/) and thought it would be cool to use a model like this to generate infinite amount of custom text to type.

&amp;#x200B;

&amp;#x200B;

[Demo](https://i.redd.it/usz1sh7okgx51.gif)

**Links**

* docs: [https://mltype.readthedocs.io/en/latest/](https://mltype.readthedocs.io/en/latest/)
* github: [https://github.com/jankrepl/mltype](https://github.com/jankrepl/mltype)

PS: I trained the model on Jekyll since that is the only Ruby project I ever used!
## [7][GVL in MRI -- Is that really needed?](https://www.reddit.com/r/ruby/comments/joq6wg/gvl_in_mri_is_that_really_needed/)
- url: https://www.reddit.com/r/ruby/comments/joq6wg/gvl_in_mri_is_that_really_needed/
---
I am a compiler engineer, and I know first versions (0.95..1.8.7) of MRI are just tree walk interpreters, they build up an AST and then interpret it, meanwhile nowadays YARV builds up an AST, compile it into bytecode and then interpret it.

(MJIT produces an AST, builds up a C source file, compile with either GCC or Clang and then run it).

Now, I am wondering why GVL is needed in MRI, since by adding a GVL you limit the current process' threads to run *asynchronously*, but not in *parallel*, but this at what pro?
## [8][A Q&amp;A with Nick Sutterer: Creator of the Trailblazer framework](https://www.reddit.com/r/ruby/comments/johxuo/a_qa_with_nick_sutterer_creator_of_the/)
- url: https://superhighway.dev/nick-sutterer-interview
---

## [9][Multiple Databases in Rails](https://www.reddit.com/r/ruby/comments/joht2y/multiple_databases_in_rails/)
- url: https://www.reddit.com/r/ruby/comments/joht2y/multiple_databases_in_rails/
---
Rails 6 supports multiple primary databases, replicas and sharding (coming in 6.1), and this is great. However, I'm wondering if there's any way to connect to independent databases at runtime, in a way that is not defined a priori in the database.yml, but instead defined by the application.

Take for example a "bring your own database" cloud service, where the "host" db of the app contains account information and other metadata, but all customer specific data is stored elsewhere on a database server controlled by the customer.

The logic can be implemented at the app level, or perhaps a thin HTTP API can be installed on the remote db server, but would there be a way to have the client run just a standard Postgres database, give auth to the app, and have the benefits of Activerecord while communicating with the remote db?
## [10][Registered for Rubyconf 2020 yet?](https://www.reddit.com/r/ruby/comments/joaf2g/registered_for_rubyconf_2020_yet/)
- url: https://www.reddit.com/r/ruby/comments/joaf2g/registered_for_rubyconf_2020_yet/
---
It’s online and $150. No hotel or flight necessary but lots of great speakers! Let’s support Ruby Central, the Ruby community and make this event awesome.

http://rubyconf.org/
