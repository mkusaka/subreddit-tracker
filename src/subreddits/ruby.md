# ruby
## [1][SAML 2.0 in Ruby](https://www.reddit.com/r/ruby/comments/ht6whl/saml_20_in_ruby/)
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
## [2][Generating Melodies With Markov Chains In Ruby](https://www.reddit.com/r/ruby/comments/ht238k/generating_melodies_with_markov_chains_in_ruby/)
- url: https://mattbettinson.com/2020/07/13/markov-melody-generation-with-ruby.html
---

## [3][Metaprogramming Ruby 2 Notes](https://www.reddit.com/r/ruby/comments/hswppm/metaprogramming_ruby_2_notes/)
- url: https://www.reddit.com/r/ruby/comments/hswppm/metaprogramming_ruby_2_notes/
---
Hello Rubyists -

I read Paolo Perrotta's Metaprogramming Ruby 2 a few years ago and it instantly became my favorite Ruby book.

After being away from Ruby for awhile now (PHP job), I've begun re-reading this book and taking notes on some of the chapters.

I thought I'd share my chapter by chapter notes here for anyone interested.

I highly recommend you pick up a copy and work through it as well.

[https://github.com/Jberczel/metaprogramming\_ruby\_2\_notes](https://github.com/Jberczel/metaprogramming_ruby_2_notes)
## [4][Sidekiq/ActiveJob style guide](https://www.reddit.com/r/ruby/comments/hswx6b/sidekiqactivejob_style_guide/)
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
## [5][Using self](https://www.reddit.com/r/ruby/comments/ht4dqu/using_self/)
- url: https://www.reddit.com/r/ruby/comments/ht4dqu/using_self/
---
Could someone give me an explanation of self as if you were explaining to a 6 year old? Every time I think I get, I see it applied differently and get pretty confused. 

Thank all!
## [6][Recursion in Ruby](https://www.reddit.com/r/ruby/comments/ht7hnu/recursion_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/ht7hnu/recursion_in_ruby/
---
Hey folks, I'm having a problem that for the life of my I can figure out. Practicing some recursion for some interview prep and I'm getting some very strange output.

    require 'byebug'
    module Solution
        def self.route_between_nodes(graph)
            paths = []
            graph[0].each do |child_array_index|
               arr = follow_index(graph, child_array_index, [0])
               puts "The returned arr is #{arr}"
               puts ""
               paths.append(arr)
            end
            paths
        end
    
        private
        def self.follow_index(graph, index, path)
            puts "The path before processing #{path}"
            if (graph[index].length == 0)
                puts "The value of the array is #{graph[index]}"
                puts "The current index value is #{index}"
                puts ""
                path.append(index)
                puts "The current path is #{path}"
                puts ""
                return path
            end
    
            graph[index].each do |new_index|
                path.append(index)
                puts "Adding the current index of #{index}"
                puts "The new index is #{new_index}"
                puts "The path is #{path}"
                puts ""
                follow_index(graph, new_index, path)
            end
        end
    end
    
    require_relative 'solution'
    RSpec.describe Solution do
        it "should return the number of paths" do
            input = [[1, 2], [3], [3], []]
            expect(Solution.route_between_nodes(input)).to eq([[0,1,3], [0,2,3]])
        end
    end

Above is the code and the test case. It should return \[\[0,1,3\], \[0,2,3\]\] but when I run the code I get \[\[3\],\[3\]\]. The strange part is that with all the print lines i've added it seems that the code should return what I expect but alas the return array is always the same. Below is an example of a run of the test case.

&amp;#x200B;

    The path before processing [0]
    Adding the current index of 1
    The new index is 3
    The path is [0, 1]
    
    The path before processing [0, 1]
    The value of the array is []
    The current index value is 3
    
    The current path is [0, 1, 3]
    
    The returned arr is [3]
    
    The path before processing [0]
    Adding the current index of 2
    The new index is 3
    The path is [0, 2]
    
    The path before processing [0, 2]
    The value of the array is []
    The current index value is 3
    
    The current path is [0, 2, 3]
    
    The returned arr is [3]

As you can see in the above, the code does successful generate the proper array but it returns \[3\]. I'm hoping that it is something very obvious I am missing but it is really driving me insane.
## [7][How Heroku Sped Up Time-Related Syscalls on Dynos](https://www.reddit.com/r/ruby/comments/hsv9bc/how_heroku_sped_up_timerelated_syscalls_on_dynos/)
- url: https://blog.heroku.com/clocksource-tuning
---

## [8][The Ruby Blend #18: Interviewing](https://www.reddit.com/r/ruby/comments/ht0mm6/the_ruby_blend_18_interviewing/)
- url: https://fireside.fm/s/ouBAUjGy+v8e3cfms
---

## [9][How to use a Transaction Script(aka Service Objects) in Ruby on Rails. Simple example](https://www.reddit.com/r/ruby/comments/hssfqu/how_to_use_a_transaction_scriptaka_service/)
- url: https://www.reddit.com/r/ruby/comments/hssfqu/how_to_use_a_transaction_scriptaka_service/
---
The logic of small applications can be present as a series of transactions. Using the Transaction Scripts pattern, we get an application that is easier to maintain, to cover with tests, and to scale.

In the [tutorial](https://jtway.co/how-to-use-a-transaction-script-aka-service-objects-in-ruby-on-rails-simple-example-161b7e228942) we will develop an [application](https://github.com/dgorodnichy/transaction-script-example) that has Post, User, and Like models. Users should be able to like posts. The first version of the controller will contain extra code, which we will extract into a separate Transaction Script. We also describe when we need to use the Transaction Scripts and the pros of the transaction script usage.

Full tutorial: [How to use a Transaction Script (aka Service Objects) in Ruby on Rails. Simple example](https://jtway.co/how-to-use-a-transaction-script-aka-service-objects-in-ruby-on-rails-simple-example-161b7e228942)
## [10][First gem, interested in feedback](https://www.reddit.com/r/ruby/comments/hsbta7/first_gem_interested_in_feedback/)
- url: https://www.reddit.com/r/ruby/comments/hsbta7/first_gem_interested_in_feedback/
---
Gem name: Backpedal

Repo: [https://github.com/Greyoxide/backpedal](https://github.com/Greyoxide/backpedal)

RubyGems: [https://rubygems.org/gems/backpedal/versions/0.1.2](https://rubygems.org/gems/backpedal/versions/0.1.2)

The goal with this gem is to provide developers with an easy way to move users back on their respective navigational path.  I found myself arriving at views from different angles and I wanted to develop a utilitarian method for moving users back along their respective navigational path.

For example: /orders/22 -&gt; /products/141 In this case the back\_link on the product view would take the user back to /orders/22.  Using this same example: /product\_lines/41 -&gt; /products/141 This controller action renders the same template as before, however the back button would take the user back to /product\_lines/41

This is my first gem, and I'd love some feedback.
