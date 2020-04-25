# ruby
## [1][git_curate v1.0.0: peruse and delete git branches ergonomically](https://www.reddit.com/r/ruby/comments/g7s2um/git_curate_v100_peruse_and_delete_git_branches/)
- url: https://github.com/matt-harvey/git_curate
---

## [2][A bit of perspective... (link to source code in comments)](https://www.reddit.com/r/ruby/comments/g7c5ag/a_bit_of_perspective_link_to_source_code_in/)
- url: https://v.redd.it/0kcjy5y2ssu41
---

## [3][Ruby project suggestions.](https://www.reddit.com/r/ruby/comments/g7n6se/ruby_project_suggestions/)
- url: https://www.reddit.com/r/ruby/comments/g7n6se/ruby_project_suggestions/
---
Hello all,

I've been learning the basics for ruby and gosu for the past few weeks and I'd like to try and create some form of application that can be used. Does anyone have suggestions for apps that can be made for beginners?
## [4][The Ultimate Guide to Ruby Timeouts](https://www.reddit.com/r/ruby/comments/g76rcd/the_ultimate_guide_to_ruby_timeouts/)
- url: https://github.com/ankane/the-ultimate-guide-to-ruby-timeouts
---

## [5][A dumb noob question about syntactic sugar](https://www.reddit.com/r/ruby/comments/g7e2jc/a_dumb_noob_question_about_syntactic_sugar/)
- url: https://www.reddit.com/r/ruby/comments/g7e2jc/a_dumb_noob_question_about_syntactic_sugar/
---
Are there comprehensive guides to using Ruby in "raw" form? I may just be missing something very obvious, especially since this is just a hobby for me. Everything I seem to find is full of sugar, and I admittedly dislike using it. I tend to enjoy languages such as Ada, that don't use much of that sort of thing. I enjoy writing Ruby, I just also like my code to be written in "longhand" form so I fully understand what is going on.
## [6][What are the best online resources or books to learn Ruby and Rails?](https://www.reddit.com/r/ruby/comments/g75h4x/what_are_the_best_online_resources_or_books_to/)
- url: https://www.reddit.com/r/ruby/comments/g75h4x/what_are_the_best_online_resources_or_books_to/
---
I’m starting my Ruby on Rails focused coding bootcamp in London in July and I would like to get a head start.   

What are the best online resources or books to learn Ruby and Rails?
## [7][Suggestions for auto-formatting Ruby Hashes and Arrays in VIM?](https://www.reddit.com/r/ruby/comments/g76v8p/suggestions_for_autoformatting_ruby_hashes_and/)
- url: https://www.reddit.com/r/ruby/comments/g76v8p/suggestions_for_autoformatting_ruby_hashes_and/
---
Do you have any useful plugins or suggestions on how to do it? Ideally I would probably want to auto-format whole file, but the tools to do so are not quire there yet. At the moment pasting some nested structures from console to test files is quite a pain.
## [8][Configuring Jumpstart to send a welcome sequence w/ Heya](https://www.reddit.com/r/ruby/comments/g6xirx/configuring_jumpstart_to_send_a_welcome_sequence/)
- url: https://gist.github.com/joshuap/03347a54e34279a467d1b2b53cb61bdd
---

## [9][Implicit return](https://www.reddit.com/r/ruby/comments/g78la4/implicit_return/)
- url: https://www.reddit.com/r/ruby/comments/g78la4/implicit_return/
---
`return`s in setters are omitted by Ruby:

```ruby
class Coordinates
  def x=(value)
    return 42
  end
end

Coordinates.new.x = 84 # =&gt; 84
```

Feel free to visit this [blogpost](https://medium.com/rubycademy/the-return-keyword-in-ruby-df0a7f578fcb) to learn more about the `return` keyword in Ruby
## [10][Connecting to API doesn't work via Ruby? (feat. Stripe)](https://www.reddit.com/r/ruby/comments/g6y77h/connecting_to_api_doesnt_work_via_ruby_feat_stripe/)
- url: https://www.reddit.com/r/ruby/comments/g6y77h/connecting_to_api_doesnt_work_via_ruby_feat_stripe/
---
Hi Redditors, I expect this will be working as I connect to their API keys but there is no update received. I think api key hasn't been connected... but What have I done wrong?

&lt;code&gt;  


`require 'sinatra'`

`require 'stripe'`

&amp;#x200B;

`Stripe.api_key = ENV['rk_test_Hiia....']`

`endpoint_secret = ENV['sk_test_WUkA....']`

&amp;#x200B;

`Stripe::WebhookEndpoint.create({`

`url: '`[`https://example.com/test`](https://example.com/test)`',`

`enabled_events: [`

`'charge.succeeded',`

`],`

`})`

&amp;#x200B;

&lt;output on Atom&gt;  


\[2020-04-24 12:05:09\] INFO  WEBrick 1.4.2 \[2020-04-24 12:05:09\] INFO  ruby 2.6.6 (2020-03-31) \[x64-mingw32\] == Sinatra (v2.0.8.1) has taken the stage on 4567 for development with backup from WEBrick \[2020-04-24 12:05:09\] INFO  WEBrick::HTTPServer#start: pid=11856 port=4567

&amp;#x200B;

&lt;Stripe API keys on dashboard&gt;  


&amp;#x200B;

https://preview.redd.it/sqm1e5uhrnu41.png?width=1344&amp;format=png&amp;auto=webp&amp;s=5bb188cb7d12561007e4fe1ec573f2418e72e46e

&amp;#x200B;

&lt;Output on Stripe webpage&gt;  


No [`https://example.com/test`](https://example.com/test) !!

https://preview.redd.it/qu7kb5anrnu41.png?width=1379&amp;format=png&amp;auto=webp&amp;s=ed7b6b50dfb62b0c5eb2afad8d2cdf74e0f76000
