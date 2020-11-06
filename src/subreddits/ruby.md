# ruby
## [1][Why Shopify's engineering team stuck with Ruby when rewriting their storefront](https://www.reddit.com/r/ruby/comments/jounah/why_shopifys_engineering_team_stuck_with_ruby/)
- url: https://www.youtube.com/watch?v=_Blz2mFjWis
---

## [2][mltype - Typing practice for Ruby and other languages](https://www.reddit.com/r/ruby/comments/jon8bl/mltype_typing_practice_for_ruby_and_other/)
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
## [3][Turkish language support for ruby](https://www.reddit.com/r/ruby/comments/jp1kno/turkish_language_support_for_ruby/)
- url: https://github.com/sbagdat/turkish_support
---

## [4][GVL in MRI -- Is that really needed?](https://www.reddit.com/r/ruby/comments/joq6wg/gvl_in_mri_is_that_really_needed/)
- url: https://www.reddit.com/r/ruby/comments/joq6wg/gvl_in_mri_is_that_really_needed/
---
I am a compiler engineer, and I know first versions (0.95..1.8.7) of MRI are just tree walk interpreters, they build up an AST and then interpret it, meanwhile nowadays YARV builds up an AST, compile it into bytecode and then interpret it.

(MJIT produces an AST, builds up a C source file, compile with either GCC or Clang and then run it).

Now, I am wondering why GVL is needed in MRI, since by adding a GVL you limit the current process' threads to run *asynchronously*, but not in *parallel*, but this at what pro?
## [5][A Q&amp;A with Nick Sutterer: Creator of the Trailblazer framework](https://www.reddit.com/r/ruby/comments/johxuo/a_qa_with_nick_sutterer_creator_of_the/)
- url: https://superhighway.dev/nick-sutterer-interview
---

## [6][Multiple Databases in Rails](https://www.reddit.com/r/ruby/comments/joht2y/multiple_databases_in_rails/)
- url: https://www.reddit.com/r/ruby/comments/joht2y/multiple_databases_in_rails/
---
Rails 6 supports multiple primary databases, replicas and sharding (coming in 6.1), and this is great. However, I'm wondering if there's any way to connect to independent databases at runtime, in a way that is not defined a priori in the database.yml, but instead defined by the application.

Take for example a "bring your own database" cloud service, where the "host" db of the app contains account information and other metadata, but all customer specific data is stored elsewhere on a database server controlled by the customer.

The logic can be implemented at the app level, or perhaps a thin HTTP API can be installed on the remote db server, but would there be a way to have the client run just a standard Postgres database, give auth to the app, and have the benefits of Activerecord while communicating with the remote db?
## [7][Registered for Rubyconf 2020 yet?](https://www.reddit.com/r/ruby/comments/joaf2g/registered_for_rubyconf_2020_yet/)
- url: https://www.reddit.com/r/ruby/comments/joaf2g/registered_for_rubyconf_2020_yet/
---
It’s online and $150. No hotel or flight necessary but lots of great speakers! Let’s support Ruby Central, the Ruby community and make this event awesome.

http://rubyconf.org/
## [8][Ruby 3.0.0 preview in 13 benchmarks](https://www.reddit.com/r/ruby/comments/jof4nr/ruby_300_preview_in_13_benchmarks/)
- url: https://github.com/kostya/jit-benchmarks
---

## [9][Termpot: A ruby gem for visualising streaming data in your terminal](https://www.reddit.com/r/ruby/comments/jnwhps/termpot_a_ruby_gem_for_visualising_streaming_data/)
- url: https://www.reddit.com/r/ruby/comments/jnwhps/termpot_a_ruby_gem_for_visualising_streaming_data/
---
I built a tool to visualise streaming data in the terminal. You can plot any command that can be piped into stdin. Its super handy for quickly visualising trends in the terminal.

Check out the repo [here](https://github.com/Martin-Nyaga/termplot)!

[Example output](https://preview.redd.it/huaw6yxo18x51.png?width=737&amp;format=png&amp;auto=webp&amp;s=2bb5b692ec52678dde7f495b9ad18df9d758df58)

Edit: Title should be \*Termplot :)
## [10][The Complete Guide to Rails Shims](https://www.reddit.com/r/ruby/comments/jo03ni/the_complete_guide_to_rails_shims/)
- url: https://www.fastruby.io/blog/rails/upgrades/rails-upgrade-shims.html
---

