# ruby
## [1][How we migrated application servers from Unicorn to Puma (GitLab)](https://www.reddit.com/r/ruby/comments/htw2iy/how_we_migrated_application_servers_from_unicorn/)
- url: https://about.gitlab.com/blog/2020/07/08/migrating-to-puma-on-gitlab/
---

## [2][Using async/await syntax in ruby with async-await gem](https://www.reddit.com/r/ruby/comments/htpwuf/using_asyncawait_syntax_in_ruby_with_asyncawait/)
- url: https://www.reddit.com/r/ruby/comments/htpwuf/using_asyncawait_syntax_in_ruby_with_asyncawait/
---
Hello! My question is about such a popular concept in other programming languages - async/await. Have you tried out a gem [async-await](https://github.com/socketry/async-await) which is built on top of [async](https://github.com/socketry/async) that gives us async decorator for functions and await method (accepting a block)?
```
class AsyncClass
  include Async::Await
  
  async def foo
    # some sutff
  end
  async def bar
    # some bar stuff
  end
end

include Async::Await
a_cl = AsynClass.new
foo_result = await { a_cl.foo }
bar_result = await { a_cl.bar }
puts "Foo: #{foo_result}"
puts "Bar: #{bar_result}"
```
Please share your thoughts/experience about it.
## [3][How can I create class instance and run a class method when .txt file gets added to the directory in Ruby?](https://www.reddit.com/r/ruby/comments/htz846/how_can_i_create_class_instance_and_run_a_class/)
- url: https://www.reddit.com/r/ruby/comments/htz846/how_can_i_create_class_instance_and_run_a_class/
---

## [4][Intermediate/Advanced Ruby and Rails Resources](https://www.reddit.com/r/ruby/comments/hthhzz/intermediateadvanced_ruby_and_rails_resources/)
- url: https://www.reddit.com/r/ruby/comments/hthhzz/intermediateadvanced_ruby_and_rails_resources/
---
Hi everyone. I haven't ever held a job as a ruby developer, but I have developed a number of pretty large applications on my own with rails. Now I'm looking to land a position as a developer, but I'm learning there are a lot of more advanced ruby concepts that I have never been exposed to.

I first learned ruby and rails using online courses and by completing the Rails Tutorial. I think I have a really good grip on the basics, but I recently had an interview with a technical questionnaire that asked some fill in the blank questions about more advanced ruby and rails topics and I was pretty lost. Some of the topics I remember were polymorphic associations, ActiveSupport::Concern, Metaprogramming hook methods, and block vs proc vs lambda.

What resources would you recommend to get a deeper knowledge of ruby and rails which would expose me to more of these topics? I prefer hands on learning, which is why I loved the rails tutorial, but I don't know of anything that exists which dives deeper into the language.

Thanks in advance for your help!
## [5][How to control the enqueuing speed of Sidekiq jobs and their concurrency](https://www.reddit.com/r/ruby/comments/hthnne/how_to_control_the_enqueuing_speed_of_sidekiq/)
- url: https://minhajuddin.com/2020/07/13/how-to-control-enqueuing-speed-of-sidekiq-jobs-and-their-execution-concurrency/
---

## [6][Introduce renderer module - API only ruby on rails course (chapter 5)](https://www.reddit.com/r/ruby/comments/hth8hh/introduce_renderer_module_api_only_ruby_on_rails/)
- url: https://duetcode.io/rails-api-only-course/introduce-renderer-module
---

## [7][SAML 2.0 in Ruby](https://www.reddit.com/r/ruby/comments/ht6whl/saml_20_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/ht6whl/saml_20_in_ruby/
---
Hi guys. I am writing a post here appealing to the help of the community since my googling hasn't worked out at all.

Basically we want to integrate SAML 2.0 SSO with an enterprise. In the terminology, they are the Identity Providers and us the Service Providers.

So I found this gem that integrates with Identity Providers:  [https://github.com/onelogin/ruby-saml](https://github.com/onelogin/ruby-saml) being used on top of [https://github.com/apokalipto/devise\_saml\_authenticatable](https://github.com/apokalipto/devise_saml_authenticatable) 

That works fine. However, I am struggling immensely to make it work with the following features:

\- Encryption

\- Rotating Certificates

I am so confused about when to use idp\_cert\_multi or when to use certificate\_new, private\_key, and idp\_cert.

Has anyone here used this library that can help me? I tried to make a "fake" Identity Provider by using the  [https://github.com/saml-idp/saml\_idp](https://github.com/saml-idp/saml_idp)  gem, which is a fork of the original gem and it seems to be lacking of proper documentation. Whenever I want to use encryption, it breaks.

So I have one question here:

1. If encryption is required on both ends, that means that both the Sdp and Idp generate a private key which they will use for encryption and a public certificate that the other party will use for decryption right?
2. What's the signature for then? Cannot we skip it? If we are using HTTPS already, why I am concerned about the signature and why the Idp cert is necessary?
## [8][Generating Melodies With Markov Chains In Ruby](https://www.reddit.com/r/ruby/comments/ht238k/generating_melodies_with_markov_chains_in_ruby/)
- url: https://mattbettinson.com/2020/07/13/markov-melody-generation-with-ruby.html
---

## [9][Metaprogramming Ruby 2 Notes](https://www.reddit.com/r/ruby/comments/hswppm/metaprogramming_ruby_2_notes/)
- url: https://www.reddit.com/r/ruby/comments/hswppm/metaprogramming_ruby_2_notes/
---
Hello Rubyists -

I read Paolo Perrotta's Metaprogramming Ruby 2 a few years ago and it instantly became my favorite Ruby book.

After being away from Ruby for awhile now (PHP job), I've begun re-reading this book and taking notes on some of the chapters.

I thought I'd share my chapter by chapter notes here for anyone interested.

I highly recommend you pick up a copy and work through it as well.

[https://github.com/Jberczel/metaprogramming\_ruby\_2\_notes](https://github.com/Jberczel/metaprogramming_ruby_2_notes)
## [10][Sidekiq/ActiveJob style guide](https://www.reddit.com/r/ruby/comments/hswx6b/sidekiqactivejob_style_guide/)
- url: https://www.reddit.com/r/ruby/comments/hswx6b/sidekiqactivejob_style_guide/
---
Finally, [the guide on how to painlessly work with Sidekiq and  ActiveJob](https://github.com/toptal/active-job-style-guide) I've been working on for so long is out. I'm extremely happy to share it with you.

It's based on:

  - Sidekiq's wiki
  - ActiveJob documentation
  - many background jobs related code reviews
  - known and rare pitfalls experienced in practice during past years

The publication of this guide is a big achievement for me, the biggest on the open-source front I can think of.

Hope you'll find it useful.
As for me, if the company I've worked for had this guide before starting DelayedJob to Sidekiq migration, we could have avoided many major headaches.

Some guidelines are unique in this guide, you won't find them in any other source.

A common belief is that ActiveJob is redundant when working with Sidekiq, and bare Sidekiq is preferable.
It's hard to argue with that.

Do not let the very first guidelines to repel you, glance over the rest of the guide.

The guide covers both topics, Sidekiq and Active Job, but Sidekiq part prevails.

Read between the lines and you'll realize the unknown unknowns there are in background job processing.
Yet, still, background job processing keeps surprising as you dive deeper and deeper.

I've barely mentioned monitoring, but it's an essential part.
Think Tetris, but three-dimensional given an extra queue dimension.
Feed your workers in an optimal way.
Otherwise you'll experience saturation, lags, and perceived slowdown.

The guide is not nearly complete.
There's [a ticket which I used as a todo list](https://github.com/toptal/active-job-style-guide/issues/1) for future guidelines.
You can help here, too.

Pull requests, additions to the todo list and any feedback are kindly appreciated.
