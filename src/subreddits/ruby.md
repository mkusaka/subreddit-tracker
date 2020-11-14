# ruby
## [1][SuperDiff for RSpec: Intelligently Display the Diff of Two Data Structures of Any Type.](https://www.reddit.com/r/ruby/comments/jtsz6u/superdiff_for_rspec_intelligently_display_the/)
- url: https://github.com/mcmire/super_diff
---

## [2][Drifting Ruby vs GoRails](https://www.reddit.com/r/ruby/comments/jttb9k/drifting_ruby_vs_gorails/)
- url: https://www.reddit.com/r/ruby/comments/jttb9k/drifting_ruby_vs_gorails/
---
Both of them are great, but I need to pick one. Given a limited budget, which one you'd choose between those two? 

I have experience in Ruby/Rails but I'm always keen to learn stuff that is not related to my daily work but scoped to the Ruby/Rails ecosystem.

Thanks in advance
## [3][How to test asyncronous methods, in Rspec?](https://www.reddit.com/r/ruby/comments/jtzpk5/how_to_test_asyncronous_methods_in_rspec/)
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
## [4][What is the best way to inherit the variables of other classes in Ruby? Rubocop complains on using a class var and I am wondering if there is another better way to do it?](https://www.reddit.com/r/ruby/comments/jtzab7/what_is_the_best_way_to_inherit_the_variables_of/)
- url: https://www.reddit.com/r/ruby/comments/jtzab7/what_is_the_best_way_to_inherit_the_variables_of/
---
For instance, I have a class named `base` and it contains all the basic information about superheroes. Now, I want to create subclasses that represent different type of superheroes and what they do. So, I have stored the basic information in `@@class variables` of the base class, so that I can inherit all of that information and create different kinds of superheroes using that info. 

  
Rubocop is complaining about using class vars, so how am I supposed to implement the above solution without using class vars? Any help is really appreciated.
## [5][Why the Release of Ruby 3 Will Be Monumental](https://www.reddit.com/r/ruby/comments/jtfzdp/why_the_release_of_ruby_3_will_be_monumental/)
- url: https://www.ruby3.dev/the-art-of-code/2020/11/12/ruby-3-monumental/
---

## [6][24/7 Lofi Hip Hop Midnight Radio - Live Stream beats to study/chill to](https://www.reddit.com/r/ruby/comments/jtxzxf/247_lofi_hip_hop_midnight_radio_live_stream_beats/)
- url: https://www.youtube.com/watch?v=fDZCOvWRxkU&amp;feature=share
---

## [7][6 Things to Do When Inheriting Legacy Rails Apps](https://www.reddit.com/r/ruby/comments/jthvt3/6_things_to_do_when_inheriting_legacy_rails_apps/)
- url: /r/rails/comments/jtgjla/6_things_to_do_when_inheriting_legacy_rails_apps/
---

## [8][DragonRuby Game Toolkit Sound Synthesis in Pure Ruby ^_^](https://www.reddit.com/r/ruby/comments/jtbtkb/dragonruby_game_toolkit_sound_synthesis_in_pure/)
- url: https://www.youtube.com/watch?v=zEzovM5jT-k&amp;feature=youtu.be&amp;ab_channel=AmirRajan
---

## [9][Custom exception on mailers deliver_later - question](https://www.reddit.com/r/ruby/comments/jtebo4/custom_exception_on_mailers_deliver_later_question/)
- url: https://www.reddit.com/r/ruby/comments/jtebo4/custom_exception_on_mailers_deliver_later_question/
---
Let's say we have a simple Mailer as:

    class PostmanMailer &lt; ApplicationMailer   
        rescue_from CustomError do |exception|
        ... do something ...
        end
         def invitation(user)
            mail(to: user.email,subject: user)
        end 
    end

And invoke that mailer with the line:

    PostmanMailer.invitation(user).deliver_later

ActionMailer is using delivery jobs and enqueues email delivery as a job through Active Job, so I can't wrap this with being/rescue to insert custom exception.

How would you handle this? Is there any way for not monkey patching?
## [10][footballdata-12xpert gem - download, convert &amp; import 22+ top football leagues from 25 seasons back to 1993/94 from Joseph Buchdahl (12Xpert)'s Football Data website](https://www.reddit.com/r/ruby/comments/jtee59/footballdata12xpert_gem_download_convert_import/)
- url: https://github.com/sportdb/sport.db.sources/tree/master/footballdata-12xpert
---

