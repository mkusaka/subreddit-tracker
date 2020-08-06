# ruby
## [1][Record a Call in Ruby with Vonage Voice API WebSockets](https://www.reddit.com/r/ruby/comments/i4nt0z/record_a_call_in_ruby_with_vonage_voice_api/)
- url: https://www.nexmo.com/blog/2020/08/04/record-a-call-in-ruby-with-vonage-voice-api-websockets
---

## [2][Could you please help me with this JRuby question?](https://www.reddit.com/r/ruby/comments/i4qj12/could_you_please_help_me_with_this_jruby_question/)
- url: https://stackoverflow.com/q/63251312
---

## [3][Trouble With Deploying Rails App with Elastic Beanstalk](https://www.reddit.com/r/ruby/comments/i4nq6a/trouble_with_deploying_rails_app_with_elastic/)
- url: https://www.reddit.com/r/ruby/comments/i4nq6a/trouble_with_deploying_rails_app_with_elastic/
---
I've been trying to figure out how to deploy a simple rails app to elastic beanstalk using the cli on my macOS, but every time I get to `eb create` I always get the error at the bottom:


    2020/08/06 07:22:56.626563 [INFO] Executing instruction: StageApplication
    2020/08/06 07:22:56.626674 [INFO] extracting /opt/elasticbeanstalk/deployment/app_source_bundle to /var/app/staging/
    2020/08/06 07:22:56.626695 [INFO] Running command /bin/sh -c /usr/bin/unzip -q -o /opt/elasticbeanstalk/deployment/app_source_bundle -d /var/app/staging/
    2020/08/06 07:22:56.638657 [INFO] finished extracting /opt/elasticbeanstalk/deployment/app_source_bundle to /var/app/staging/ successfully
    2020/08/06 07:22:56.640331 [INFO] Executing instruction: RunAppDeployPreBuildHooks
    2020/08/06 07:22:56.640351 [INFO] The dir .platform/hooks/prebuild/ does not exist in the application. Skipping this step...
    2020/08/06 07:22:56.640356 [INFO] Executing instruction: stage ruby application
    2020/08/06 07:22:56.640360 [INFO] stage ruby application ....
    2020/08/06 07:22:56.640386 [INFO] Running command /bin/sh -c bundle config set --local deployment true
    2020/08/06 07:22:56.669550 [ERROR] An error occurred during execution of command [app-deploy] - [stage ruby application]. Stop running the command. Error: install dependencies in Gemfile failed with error Command /bin/sh -c bundle config set --local deployment true failed with error exit status 1. Stderr:rbenv: version `ruby-2.7.0' is not installed (set by /var/app/staging/.ruby-version)


After that I run `eb terminate` so that I don't incur any costs. But I don't understand the solution. Previously, the error said that 'ruby-2.6.3' is not installed' and I was unable to update to the latest version using rbenv so I removed it with homebrew and installed Ruby 2.7.0 using rvm, then tried to remove rails with `gem uninstall rails`. 

Then I tried installing rails again using `\curl -sSL https://get.rvm.io | bash -s stable --rails` which I got from the [RVM install guide](https://rvm.io/rvm/install), but then the above error happened. 

I imagine the answer would be to somehow update Ruby using rbenv? It was pretty difficult for me to figure that out, I believe since there is a version on Mac that is difficult to upgrade since the os uses it for legacy apps? Does anybody have a resource to help with that? I'm sure there's lots of flaws in my understanding of all this so any direction would be appreciated.
## [4][Introduction to Ruby on Rails Patterns and Anti-patterns](https://www.reddit.com/r/ruby/comments/i45k0j/introduction_to_ruby_on_rails_patterns_and/)
- url: https://blog.appsignal.com/2020/08/05/introduction-to-ruby-on-rails-patterns-and-anti-patterns.html
---

## [5][Rails 6.1 automatically generates an abstract class when using multiple databases](https://www.reddit.com/r/ruby/comments/i44blm/rails_61_automatically_generates_an_abstract/)
- url: https://blog.bigbinary.com/2020/08/04/rails-6-1-automatically-generates-abstract-class-when-using-multiple-databases.html
---

## [6][gRPC implementation's float precision](https://www.reddit.com/r/ruby/comments/i4b39z/grpc_implementations_float_precision/)
- url: https://www.reddit.com/r/ruby/comments/i4b39z/grpc_implementations_float_precision/
---
Hi everyone.

I'm writing a ruby client for a gRPC service (a service I've also written but in Kotlin) and one of the endpoints returns an object which has properties at are floats. The response I'm getting in ruby client is a float with wildly varying precision despite the server providing values truncated to 3 decimal places. I even hard coded a response to `0.26` and this is what I get when making a request to the server with the ruby client `&lt;Google::Protobuf::FloatValue: value: 0.25999999046325684&gt;`

When I make the same request using a Kotlin gRPC client I get the response that I expect.

Does anyone have any thoughts on what's going on here and how I can fix it. Thanks.
## [7][Beginner trouble](https://www.reddit.com/r/ruby/comments/i45ijz/beginner_trouble/)
- url: https://www.reddit.com/r/ruby/comments/i45ijz/beginner_trouble/
---
I'm a complete beginner trying to teach myself how to program, starting with Ruby. I had started off by trying to learn the language in a "training" website, but have been told that I would probably learn better by just downloading it and giving myself projects to work on.I'm downloading it now, and I'm just a little confused. I'm using this:  
 [https://www.dummies.com/programming/ruby/how-to-install-and-run-ruby-on-windows/](https://www.dummies.com/programming/ruby/how-to-install-and-run-ruby-on-windows/)    
for guidance, but it seems like it's out of date? I downloaded the option that said it had the dev kit included, but I can't find it to do the rest of the steps. Would I not find it? Is there another guide I should be following.
## [8][I wonder if all should adopt Ruby 2.7's new JavaScript-style object destructuring?](https://www.reddit.com/r/ruby/comments/i3javx/i_wonder_if_all_should_adopt_ruby_27s_new/)
- url: https://idiosyncratic-ruby.com/68-assignments-in-style.html
---

## [9][Intermediate learning resources?](https://www.reddit.com/r/ruby/comments/i3s03r/intermediate_learning_resources/)
- url: https://www.reddit.com/r/ruby/comments/i3s03r/intermediate_learning_resources/
---
Anyone know some direction for online intermediate level learning?? 

I’ve done the intro books, built simple text based games, little tutorial programs ... how do I level up now?!
Just need some starting points for the next level. 
Thanks all.
## [10][Open source status update, July 2020 (completing initial Hanami 2.0 action/view integration)](https://www.reddit.com/r/ruby/comments/i3gulr/open_source_status_update_july_2020_completing/)
- url: https://timriley.info/writing/2020/08/03/open-source-status-update-july-2020/
---

