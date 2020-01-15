# ruby
## [1][Creating my first ruby gems](https://www.reddit.com/r/ruby/comments/ep0lhp/creating_my_first_ruby_gems/)
- url: https://www.reddit.com/r/ruby/comments/ep0lhp/creating_my_first_ruby_gems/
---
Hello

I've just made my first gem! It's a wrapper for an instant search engine API (MeiliSearch). It means indexes, documents, and search handling.  
[https://github.com/meilisearch/meilisearch-ruby](https://github.com/meilisearch/meilisearch-ruby)

I love ruby language, it's literally my favorite one, but I use it more with Rails than in this kind of situation.

So, I need your feedback to correct and improve it!

Thanks for your time.

And for the curious ones, here is the repository of our search engine MeiliSearch. We are open-source and we really need your support! ⭐️ Each star and feedback are important for us 🙂[https://github.com/meilisearch/MeiliSearch](https://github.com/meilisearch/MeiliSearch)
## [2][Developer with 10 years experience of Ruby / Rails (12 in programming) is looking for juniors/mids to teach/help/mentor - anyone interested?](https://www.reddit.com/r/ruby/comments/eolm8x/developer_with_10_years_experience_of_ruby_rails/)
- url: https://www.reddit.com/r/ruby/comments/eolm8x/developer_with_10_years_experience_of_ruby_rails/
---
Hello everyone, I hope you are well. I'm not sure if it's the best place for such announcements but I was sharing with you my articles for the past 2 years so I decided to write here as I appreciate this community and the value it gives.

I have just finished one of my projects and have more free time now. I thought that I may work with programmers that have fewer experience than me to help them improve the quality of their code or careers as Ruby developers. Does anyone need a consultation, code review or advice? It will be free, maybe in the future, I will think about some kind of compensation in case of a longer partnership but for now, I just wanted to check if there is anyone interested that I could share my knowledge and experience with and see what type of problems do you have.

Short about me:

Experience: 10 years with Ruby / Rails (12 in programming)

Github: [https://github.com/pdabrowski6](https://github.com/pdabrowski6)

Website: [https://pdabrowski.com/](https://pdabrowski.com/)

Podcasts: 

* [https://devchat.tv/ruby-rogues/rr-366-build-your-own-rspec-a-gentle-metaprogramming-intro-with-pawel-dabrowski/](https://devchat.tv/ruby-rogues/rr-366-build-your-own-rspec-a-gentle-metaprogramming-intro-with-pawel-dabrowski/)
* https://devchat.tv/my-ruby-story/mrs-069-pawel-dabrowski/

LinkedIn: [https://www.linkedin.com/in/pdabrowski6/](https://www.linkedin.com/in/pdabrowski6/)

Feel free to contact me here or via Linkedin if you are interested in the partnership. I would love to help other developers and learn a lot by doing this. Have a nice day!
## [3][BigDecimal Equality Strangeness](https://www.reddit.com/r/ruby/comments/eos2gf/bigdecimal_equality_strangeness/)
- url: https://www.reddit.com/r/ruby/comments/eos2gf/bigdecimal_equality_strangeness/
---
I ran into a BigDecimal value which has confused the heck out of me.  This is on Ruby 2.6.3

    2.6.3 :007 &gt; BigDecimal("0.08442931")
     =&gt; 0.8442931e-1
    2.6.3 :008 &gt; BigDecimal("0.08442931") - 0.8442931e-1
     =&gt; 0.1e-16

Take note I copied the value from the first line and pasted it to the second.  I otherwise get

    2.6.3 :002 &gt;  BigDecimal("0.08442931") -  BigDecimal("0.08442931")
     =&gt; 0.0

What is going on here?  It doesn't happen often with other values, for example if I change the last digit

    2.6.3 :004 &gt;  BigDecimal("0.08442932") - 0.8442932e-1
     =&gt; 0.0

Any idea what's happening?
## [4][Building a Hangman Game in Ruby in 25 minutes](https://www.reddit.com/r/ruby/comments/eopvss/building_a_hangman_game_in_ruby_in_25_minutes/)
- url: https://www.reddit.com/r/ruby/comments/eopvss/building_a_hangman_game_in_ruby_in_25_minutes/
---
Hey guys, I've launched a new Ruby build on YouTube. This is a pretty interesting one as it involves a lot of interaction with the user and handling the user's input from the Terminal. Guess the wrong letter too many times and it's game over. Was a lot of fun writing this code for the video, hopefully it's useful to anyone getting into Ruby.

[https://www.youtube.com/watch?v=uBwGfswwRL4](https://www.youtube.com/watch?v=uBwGfswwRL4)

As always, I'm planning to post some more Ruby content to my channel to promote the Ruby language, so would be open to suggestions about content ideas. Thanks :)
## [5][A Migration Path to Bundler 2+](https://www.reddit.com/r/ruby/comments/eolrut/a_migration_path_to_bundler_2/)
- url: http://eregon.me/blog/2020/01/13/a-migration-path-to-bundler2.html
---

## [6][Rails tutorial - Why you should use Sentry logger in your project and how to configure it](https://www.reddit.com/r/ruby/comments/eol6qd/rails_tutorial_why_you_should_use_sentry_logger/)
- url: https://hixonrails.com/ruby-on-rails-tutorials/ruby-on-rails-sentry-logger-installation-and-configuration/
---

## [7][can someone explain the difference between rangeinclude and rangecover like I am 5 years old?](https://www.reddit.com/r/ruby/comments/eoor5w/can_someone_explain_the_difference_between/)
- url: https://www.reddit.com/r/ruby/comments/eoor5w/can_someone_explain_the_difference_between/
---
I've search the same question but I still don't get why...
## [8][How can you do proper async with ruby?](https://www.reddit.com/r/ruby/comments/eoo8sp/how_can_you_do_proper_async_with_ruby/)
- url: https://www.reddit.com/r/ruby/comments/eoo8sp/how_can_you_do_proper_async_with_ruby/
---
I want to be able to run an method, and while it runs run another method not waiting for the first to end/exit. It can easily be done in JS with stuff like promises but I couldn't figure out something similar for ruby.

Something like:

    def method
        some_method_that_doesnt_end
        other_method
    end

&amp;#x200B;

While `some_method_that_doesnt_end` can run indefinitely, I want to run several of those at the same time.
## [9][Faster Ruby APIs with Postgres](https://www.reddit.com/r/ruby/comments/eom0cc/faster_ruby_apis_with_postgres/)
- url: https://goiabada.blog/faster-ruby-apis-with-postgres-238c2f4a272c
---

## [10][www.viky.ai a Natural Language Processing platform (Ruby On Rails app + C NLP component)](https://www.reddit.com/r/ruby/comments/eojgtk/wwwvikyai_a_natural_language_processing_platform/)
- url: https://www.reddit.com/r/ruby/comments/eojgtk/wwwvikyai_a_natural_language_processing_platform/
---
Hi,

[viky.ai](https://www.viky.ai/) is a **Natural Language Processing platform**. It allows you to **extract information from unstructured text contents**.

The technical component *nlp*, written in C, allows the extraction of structured information, this extraction is defined within agents. The agents are multilingual assistants to find relevant data. The *nlp* component takes as input a set of agents in JSON format and unstructured textual content in order to provide a JSON stream of structured data as output.

The second technical component webapp is a **Rails application** that allows you to work collaboratively to set up agents by offering dedicated interfaces. It also provides the interpret API in order to allow integration into a third-party system.

**viky.ai** **is released under open-source license** (MIT), code is available on [Github](https://github.com/viky-ai/viky-ai).
