# ruby
## [1][Using Rails's Attributes API to serialize Value Objects](https://www.reddit.com/r/ruby/comments/f2o2j0/using_railss_attributes_api_to_serialize_value/)
- url: https://aliou.me/posts/2019/10/attributes-api-and-value-objects/
---

## [2][Getting Started With System Tests in Rails With Minitest](https://www.reddit.com/r/ruby/comments/f2qmkh/getting_started_with_system_tests_in_rails_with/)
- url: https://blog.appsignal.com/2020/02/12/getting-started-with-system-tests-in-ruby-with-minitest.html
---

## [3][How to pitch GraphQL to your team?](https://www.reddit.com/r/ruby/comments/f2py8b/how_to_pitch_graphql_to_your_team/)
- url: https://blog.graphqleditor.com/how-to-convince-your-team-to-use-graphql/
---

## [4][[X-Post from r/emacs] How to debug Ruby in Emacs?](https://www.reddit.com/r/ruby/comments/f2p534/xpost_from_remacs_how_to_debug_ruby_in_emacs/)
- url: https://old.reddit.com/r/emacs/comments/f2ocgc/how_to_debug_ruby_in_emacs/
---

## [5][What ruby talk topics would you want to hear?](https://www.reddit.com/r/ruby/comments/f2c5hc/what_ruby_talk_topics_would_you_want_to_hear/)
- url: https://www.reddit.com/r/ruby/comments/f2c5hc/what_ruby_talk_topics_would_you_want_to_hear/
---
I host a local ruby meetup. Most of the time people come to me with talk ideas, but every once in a while, I ask someone known in the local ruby community to speak, and the ask me what they should speak about.

What sort of talks do you like? What topics? Anything new and interesting come out in the ruby/rails space that you would attend a talk for? What other questions should I be asking?

I want to keep this local ruby group valuable but I am just a baby in my career and I don't know what I don't know.
## [6][A Developer's Notebook - Repurposing an old Android phone as a Ruby web server](https://www.reddit.com/r/ruby/comments/f25fmm/a_developers_notebook_repurposing_an_old_android/)
- url: https://lbrito1.github.io/blog/2020/02/repurposing-android.html
---

## [7][Why is .nil? not working as intended here?](https://www.reddit.com/r/ruby/comments/f2b33s/why_is_nil_not_working_as_intended_here/)
- url: https://www.reddit.com/r/ruby/comments/f2b33s/why_is_nil_not_working_as_intended_here/
---
I have a method that takes in a parameter, letter. If the parameter is empty, not a valid letter, or nil, an exception will be thrown. The exception is thrown if letter is empty or not a valid letter, but the test case for checking if letter is nil never seems to pass. I have provided the code, test case, and the error that gets thrown below.

Code:

`def guess(letter)`    

`letter = letter.downcase`

`if letter.empty? || /[A-Za-z]/ !~ letter || letter.nil?`

`raise ArgumentError`

`end`

&amp;#x200B;

Test case:

`context 'invalid' do`      `before :each dou/game = HangpersonGame.new('foobar')end`      `it 'throws an error when empty' do`        `expect { u/game.guess('') }.to raise_error(ArgumentError)end`      `it 'throws an error when not a letter' do`        `expect { u/game.guess('%') }.to raise_error(ArgumentError)end`      `it 'throws an error when nil' do`        `expect { u/game.guess(nil) }.to raise_error(ArgumentError)endend`

&amp;#x200B;

Error:

`Failure/Error: expect { u/game.guess(nil) }.to raise_error(ArgumentError)`

`expected ArgumentError, got #&lt;NoMethodError: undefined method \downcase' for nil:NilClass&gt; with backtrace:\``
## [8][Why is my if else statement faster than my case switch statement?](https://www.reddit.com/r/ruby/comments/f26g8h/why_is_my_if_else_statement_faster_than_my_case/)
- url: https://www.reddit.com/r/ruby/comments/f26g8h/why_is_my_if_else_statement_faster_than_my_case/
---
I was under the impression that switch statements were faster than, or at least as fast as if statements, yet when I did a benchmark the if statement was 15-20% faster. Just trying to learn when it's appropriate to use switch statements, if it has some quirks that slow it down or something.    
  
Benchmark code: https://pastebin.com/MgVaZXbS
## [9][Is Ruby worth learning in 2020?](https://www.reddit.com/r/ruby/comments/f1sx71/is_ruby_worth_learning_in_2020/)
- url: https://www.reddit.com/r/ruby/comments/f1sx71/is_ruby_worth_learning_in_2020/
---
I've been looking for a new language to use on side-projects, and I recently watched a few conference talks about Ruby, and I'm really impressed with the ergonomics of the language. I was thinking about learning the language, but I was wondering what the community and ecosystem looks like these days. I know that Ruby has fallen away a bit in the "new hotness" sphere, but I'm hoping that this is due to stability and that I'll be able to pick up a language that has staying power.  


Why is Ruby worth learning in 2020?
## [10][Manage Redis by Ruby On Rails](https://www.reddit.com/r/ruby/comments/f1x37r/manage_redis_by_ruby_on_rails/)
- url: https://www.reddit.com/r/ruby/comments/f1x37r/manage_redis_by_ruby_on_rails/
---
I come to present a tool that I have developed with 2 others developers [https://github.com/OpenGems/redis\_web\_manager](https://github.com/OpenGems/redis_web_manager)

This  tool allows you to manage Redis, thanks to a web interface you  can  easily manage your Redis instance (see your keys, the memory used,   clients connected, etc …).

Enjoy !
