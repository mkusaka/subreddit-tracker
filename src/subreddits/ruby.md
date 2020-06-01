# ruby
## [1][Is there an equivalent framework to Tornado (Python) in Ruby?](https://www.reddit.com/r/ruby/comments/guhanp/is_there_an_equivalent_framework_to_tornado/)
- url: https://www.reddit.com/r/ruby/comments/guhanp/is_there_an_equivalent_framework_to_tornado/
---
I have been spending a bit of time in Python since discovering the Tornado framework, but I still can't get to love Python the way I love Ruby.

I am looking for a Ruby framework that supports non-blocking HTTP requests and also WebSockets.

From looking around, it seems like you have to cobble something together from different libraries. Or it is possible with an existing framework but that they don't mention it explicitly because it's a given that it would support it.
## [2][gRPC concurrency](https://www.reddit.com/r/ruby/comments/guffav/grpc_concurrency/)
- url: https://www.reddit.com/r/ruby/comments/guffav/grpc_concurrency/
---
So, I have a lightweight ruby gRPC server running with docker and kubernetes. No threading or forking. I’ve used rails / rack apps for a long time, and I’m used to having some sense of concurrency via webrick, unicorn, puma, passenger. 

My question is around concurrency. Since this service has such a small footprint, of like 10m cpu and 10mb of ram, would it be best to scale up the pods and let the cluster handle the load balancing? My searches for “ruby grpc concurrency” were not fruitful, so there doesn’t seem to be anything out of the box for going the “traditional” way.
## [3][Help with if and else in a simple quiz program](https://www.reddit.com/r/ruby/comments/gudrah/help_with_if_and_else_in_a_simple_quiz_program/)
- url: https://www.reddit.com/r/ruby/comments/gudrah/help_with_if_and_else_in_a_simple_quiz_program/
---
Hello I am very new to Ruby I just started a udemy ruby course a few days ago. I am writing a simple quiz game program but need help.

when I type

`answer = gets`

`if answer == "yes"`

`p "correct"`

`else answer != "yes"`

`p incorrect`

`end`

&amp;#x200B;

&amp;#x200B;

I only ever get the output for if answer != "yes"

I have tried using elsif instead of else and doing

else answer == "no"

p "incorrect"

but it is the same result.
## [4][question on OOP and a module, not using `self` or @](https://www.reddit.com/r/ruby/comments/gui0lt/question_on_oop_and_a_module_not_using_self_or/)
- url: https://www.reddit.com/r/ruby/comments/gui0lt/question_on_oop_and_a_module_not_using_self_or/
---
Hi, so I was curious, I was using the below module in other classes of a program, and it has it's own `to_s` method. 

I was curious as to why simply `name`, `role` and `id` are allowed in the interpolated braces and not seen as for example local variables, and why `@name`, `@role` and `@id` are not needed.  Both scenarios work.  Even using `self.name` and so forth works in the below interpolation braces.  Why is simply `name` and so forth allowed in the below, when for example in a class itself, we need to - or do we? - have `self.` when accessing an instance variable / attribute of an instance: 


    module Info
      def to_s
        "Name: #{name}\n
         Role: #{role}\n
         ID: #{id}"
      end
    end
## [5][How do I deal with non-registered users on react-native app with Ruby on Rails Backend using Devise?](https://www.reddit.com/r/ruby/comments/gu4cqb/how_do_i_deal_with_nonregistered_users_on/)
- url: https://www.reddit.com/r/ruby/comments/gu4cqb/how_do_i_deal_with_nonregistered_users_on/
---
On my react-native app I have an option "Continue without account" because I think it will allow people to start using the application straight away.

My backend is complete, it's been running the web-application version for months. 

I use devise for user handling, however as far as I'm aware theres no options to allow for a user to be created with a blank email/password so how do I go about doing this?

I don't have much experience building mobile apps, so I don't know what the best possible solution would be.

First thing that comes to mind is creating a GuestUser model that doesn't have these restrictions on it, has all the same relations that the normal User model has; and when a GuestUser wants to register their account, all the records that once belonged to the GuestUser are copied across, the user-token on the mobile side is updated and the GuestUser account is deleted.


What do you guys think?
## [6][tty-pager - adds support for streaming content to a terminal pager](https://www.reddit.com/r/ruby/comments/gtewm3/ttypager_adds_support_for_streaming_content_to_a/)
- url: https://github.com/piotrmurach/tty-pager
---

## [7][What to use for web API + small single page web app?](https://www.reddit.com/r/ruby/comments/gtec37/what_to_use_for_web_api_small_single_page_web_app/)
- url: https://www.reddit.com/r/ruby/comments/gtec37/what_to_use_for_web_api_small_single_page_web_app/
---
Hi.  I’m looking to build a small web service which is basically mainly a DB + lots of custom business logic. It need to provide an API for integration with different parts of the company. It will also have a small web UI part which is basically a status dashboard, and some minor editing functionality for some of the entities.

It’s been a while since I’ve done anything like full stack webdev work. I was somewhat familiar with rails 3 ... but looking at rails now, it seems to have gotten a lot larger and more complex.

So, what’s your current favorite web framework for “mostly backend” stuff like this?
Sinatra? I’ve had a quick look at Hanami, and it seems to fit, but it also seems to have slowed down its development a lot in the last ~6 months. 
I’m most comfortable in Ruby, and this is also the wrong subreddit to ask ... but if anyone has experience with Python’s flask, I would be happy to hear how it compares.

Some form of data mapper would be nice to have.
It as long as I don’t have to escape the SQL myself, I’m happy.

And if there any tips for matching front end things that work well together, that would be appreciated as well. Although the easiest will probably be to just use a template engine and vanilla HTML/CSS/JS.
## [8][Test-Driving a Decision Engine](https://www.reddit.com/r/ruby/comments/gt77s1/testdriving_a_decision_engine/)
- url: https://medium.com/one-medical-technology/building-a-decision-engine-54f6640dd3d?source=friends_link&amp;sk=9dcb2d0618c8a4a65e4969b62a0bb001
---

## [9][From a Newbie to another . Do you want collaborate with a project to Learn by coding?](https://www.reddit.com/r/ruby/comments/gtimtw/from_a_newbie_to_another_do_you_want_collaborate/)
- url: https://www.reddit.com/r/ruby/comments/gtimtw/from_a_newbie_to_another_do_you_want_collaborate/
---
Hi There 

I am a newbie in Ruby and decide to embrace a project to help me learn more ... several people believe this is a very good way ... learn by code and better if you have a REAL PROJECT.

I will bring my Shopify Liquid knowledge as the project is develop an integration app between shopify and a Shipping platform here in Brasil 

Maybe the code when finished and working can be listed into Shopp App Store and generate a recurring revenue that I agree to split with collaborator ...

Why Ruby ?  Well  in fact my interest came up when I discover that was the language used to build shopify ... but ... I think could also be Python  .. maybe I will post this on Python community too  

&amp;#x200B;

Best Regards
## [10][Ruby gem to work with SEPA Reason Codes](https://www.reddit.com/r/ruby/comments/gszrsn/ruby_gem_to_work_with_sepa_reason_codes/)
- url: https://www.reddit.com/r/ruby/comments/gszrsn/ruby_gem_to_work_with_sepa_reason_codes/
---
I wrote a super minimal gem to work with SEPA Reason Codes.   


When implementing payments with SEPA, I had a pain point of not finding information regarding reason codes. So, I found and combined them into a gem to make working with them easier.

&amp;#x200B;

[https://www.ramblingcode.dev/posts/sepa\_reason\_codes\_ruby\_gem/](https://www.ramblingcode.dev/posts/sepa_reason_codes_ruby_gem/)
