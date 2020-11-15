# ruby
## [1][HexaPDF 0.13.0 released - Cross-reference Table Reconstruction](https://www.reddit.com/r/ruby/comments/jui5et/hexapdf_0130_released_crossreference_table/)
- url: https://hexapdf.gettalong.org/news/2020/hexapdf-0-13-0.html
---

## [2][football-cat gem - concatenate football.csv datafiles - make out of many, one (for easy (re)use or imports)](https://www.reddit.com/r/ruby/comments/jujvc5/footballcat_gem_concatenate_footballcsv_datafiles/)
- url: https://github.com/sportdb/football.db/tree/master/football-cat
---

## [3][Beginner's Guide to Rails Testing (PDF)](https://www.reddit.com/r/ruby/comments/ju2kw4/beginners_guide_to_rails_testing_pdf/)
- url: https://www.reddit.com/r/ruby/comments/ju2kw4/beginners_guide_to_rails_testing_pdf/
---
I want to share something with you guys that I've been working on for the last few weeks.

I've been teaching Rails testing for close to three years now. In that time there are certain questions I've seen come up over and over. Recently I decided to write a series of blog posts answering these questions.

The questions are:

* What are the different kinds of Rails tests?
* What are all the Rails testing tools and how do I use them?
* Which testing framework should I use (RSpec or Minitest)?
* How can I make testing a habitual part of my development work?
* What level of test coverage should I shoot for?
* How do I add tests to an existing Rails project?
* How do I set up a new Rails project for testing?
* Should I be doing test-driven development?

Recently I aggregated all my answers into a small "e-book", a 45-page PDF. I figured it would make a handy guide where all the answers are in one place.

[http://www.codewithjason.com/wp-content/uploads/2020/11/beginners\_guide\_to\_rails\_testing.pdf](http://www.codewithjason.com/wp-content/uploads/2020/11/beginners_guide_to_rails_testing.pdf)

If you're someone who's new to testing (or trying to get better at testing), I hope it helps you.

And if there are any questions my guide does NOT answer for you, please feel free to ask your question in a comment and I'll do my best to answer.
## [4][SuperDiff for RSpec: Intelligently Display the Diff of Two Data Structures of Any Type.](https://www.reddit.com/r/ruby/comments/jtsz6u/superdiff_for_rspec_intelligently_display_the/)
- url: https://github.com/mcmire/super_diff
---

## [5][What is the best way to inherit the variables of other classes in Ruby? Rubocop complains on using a class var and I am wondering if there is another better way to do it?](https://www.reddit.com/r/ruby/comments/jtzab7/what_is_the_best_way_to_inherit_the_variables_of/)
- url: https://www.reddit.com/r/ruby/comments/jtzab7/what_is_the_best_way_to_inherit_the_variables_of/
---
For instance, I have a class named `base` and it contains all the basic information about superheroes. Now, I want to create subclasses that represent different type of superheroes and what they do. So, I have stored the basic information in `@@class variables` of the base class, so that I can inherit all of that information and create different kinds of superheroes using that info. 

  
Rubocop is complaining about using class vars, so how am I supposed to implement the above solution without using class vars? Any help is really appreciated.
## [6][How to test asyncronous methods, in Rspec?](https://www.reddit.com/r/ruby/comments/jtzpk5/how_to_test_asyncronous_methods_in_rspec/)
- url: https://www.reddit.com/r/ruby/comments/jtzpk5/how_to_test_asyncronous_methods_in_rspec/
---
I have several Ruby methods that call methods from a library written in C++, via \`FFI\` gem. Some of the C++ methods involve spawning new threads inside a native library, others not.

&amp;#x200B;

The ruby methods look, basically, like this:

&amp;#x200B;

           my_params = create_params_for_method1()
           MyModule::method1(context, my_params) do |result|
             puts "result is: #{result}"
           end
    

All works fine. 

&amp;#x200B;

A difficulty I have is that how to test these asyncronous methods with Rspec? Some methods return value almost immediately and there's no issue with testing them:

        expect { |b| MyModule.method1(my_context, my_params, &amp;b) }.to yield_control
        MyModule.method1(my_context, my_params) { |r| u/result = r }
    
        expect(@result).to eq "some_result123" # OK
    

&amp;#x200B;

But others incurr a delay or keep sending data in a loop and therefore can't be tested in this way:

    expect { |b| MyModule.method1(my_context, my_params, &amp;b) }.to yield_control # error
    MyModule.method1(my_context, my_params) { |r| @result = r }                 # error
    
    expect(@result).to eq "some_result123"        # @result is still nil
    
    

&amp;#x200B;

&amp;#x200B;

I could insert **sleep N** in a rspec test, but this be a sensible approach? 

        # expect { |b| MyModule.method1(my_context, my_params, &amp;b) }.to yield_control
        MyModule.method1(my_context, my_params) { |r| @result = r }                
    
        sleep 10
    
        expect(@result).to eq "some_result123"            # not nil, but...
    

How long in seconds should I N be so that it'll always pass? The more the better, but how long? How would I guess?

&amp;#x200B;

Also, it's a hack, isn't it?

Also, in some cases it'll **keep** sending data from a native library multiple times. How would I make sure that I've received it all, in a test?

&amp;#x200B;

And  a test isn't supposed to take a long, undetermenistic time to test a method, is it?

&amp;#x200B;

What's the approach in such a case?
## [7][Drifting Ruby vs GoRails](https://www.reddit.com/r/ruby/comments/jttb9k/drifting_ruby_vs_gorails/)
- url: https://www.reddit.com/r/ruby/comments/jttb9k/drifting_ruby_vs_gorails/
---
Both of them are great, but I need to pick one. Given a limited budget, which one you'd choose between those two? 

I have experience in Ruby/Rails but I'm always keen to learn stuff that is not related to my daily work but scoped to the Ruby/Rails ecosystem.

Thanks in advance
## [8][Why the Release of Ruby 3 Will Be Monumental](https://www.reddit.com/r/ruby/comments/jtfzdp/why_the_release_of_ruby_3_will_be_monumental/)
- url: https://www.ruby3.dev/the-art-of-code/2020/11/12/ruby-3-monumental/
---

## [9][6 Things to Do When Inheriting Legacy Rails Apps](https://www.reddit.com/r/ruby/comments/jthvt3/6_things_to_do_when_inheriting_legacy_rails_apps/)
- url: /r/rails/comments/jtgjla/6_things_to_do_when_inheriting_legacy_rails_apps/
---

## [10][DragonRuby Game Toolkit Sound Synthesis in Pure Ruby ^_^](https://www.reddit.com/r/ruby/comments/jtbtkb/dragonruby_game_toolkit_sound_synthesis_in_pure/)
- url: https://www.youtube.com/watch?v=zEzovM5jT-k&amp;feature=youtu.be&amp;ab_channel=AmirRajan
---

