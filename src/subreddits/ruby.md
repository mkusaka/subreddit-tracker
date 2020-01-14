# ruby
## [1][Rails tutorial - Why you should use Sentry logger in your project and how to configure it](https://www.reddit.com/r/ruby/comments/eol6qd/rails_tutorial_why_you_should_use_sentry_logger/)
- url: https://hixonrails.com/ruby-on-rails-tutorials/ruby-on-rails-sentry-logger-installation-and-configuration/
---

## [2][www.viky.ai a Natural Language Processing platform (Ruby On Rails app + C NLP component)](https://www.reddit.com/r/ruby/comments/eojgtk/wwwvikyai_a_natural_language_processing_platform/)
- url: https://www.reddit.com/r/ruby/comments/eojgtk/wwwvikyai_a_natural_language_processing_platform/
---
Hi,

[viky.ai](https://www.viky.ai/) is a **Natural Language Processing platform**. It allows you to **extract information from unstructured text contents**.

The technical component *nlp*, written in C, allows the extraction of structured information, this extraction is defined within agents. The agents are multilingual assistants to find relevant data. The *nlp* component takes as input a set of agents in JSON format and unstructured textual content in order to provide a JSON stream of structured data as output.

The second technical component webapp is a **Rails application** that allows you to work collaboratively to set up agents by offering dedicated interfaces. It also provides the interpret API in order to allow integration into a third-party system.

**viky.ai** **is released under open-source license** (MIT), code is available on [Github](https://github.com/viky-ai/viky-ai).
## [3][Use cases for ruby aside from web](https://www.reddit.com/r/ruby/comments/eoieg8/use_cases_for_ruby_aside_from_web/)
- url: https://www.reddit.com/r/ruby/comments/eoieg8/use_cases_for_ruby_aside_from_web/
---
So yesterday I had asked about some use cases for the ruby language in data science, and the general idea was that Ruby isn't good for data science &amp; engineering. Which led me to wonder , "aside from rails, where else does ruby shine?"
## [4][Cache Crispies - Fast, Flexible Rails Serializer](https://www.reddit.com/r/ruby/comments/eoegih/cache_crispies_fast_flexible_rails_serializer/)
- url: https://www.reddit.com/r/ruby/comments/eoegih/cache_crispies_fast_flexible_rails_serializer/
---
Picking a method of doing JSON serialization in Rails has not been an easy decision as of late. Especially if you're not able to break your APIs by moving to a JSON API structure. And trying to mix in a caching strategy certainly doesn't help. That's the problem the new Cache Crispies gem was written to fix.  


[https://codenoble.com/blog/introducing-cache-crispies/](https://codenoble.com/blog/introducing-cache-crispies/)  
[https://github.com/codenoble/cache-crispies](https://github.com/codenoble/cache-crispies)
## [5][Business software projects ?](https://www.reddit.com/r/ruby/comments/eol0b0/business_software_projects/)
- url: https://www.reddit.com/r/ruby/comments/eol0b0/business_software_projects/
---
Sometime this semester, I will have a project to build a system in an area of my choosing. Among other things, I want to see if I can hack building software with business applications. Does anyone on this sub have an idea on what I can build?
## [6][Ruby for Data Science.](https://www.reddit.com/r/ruby/comments/eo6ox3/ruby_for_data_science/)
- url: https://www.reddit.com/r/ruby/comments/eo6ox3/ruby_for_data_science/
---
This semester I will be doing a scientific computing unit which I have come to learn is about data science (correct me if I'm wrong)  
We will be using python, but I wanted to know if any experienced people on this sub have used Ruby for data science projects, and if so, would you mind sharing your experience with regards to the learning curve associated with it and any other relevant information?
## [7][Problem with gems not bundling](https://www.reddit.com/r/ruby/comments/eo8je4/problem_with_gems_not_bundling/)
- url: https://www.reddit.com/r/ruby/comments/eo8je4/problem_with_gems_not_bundling/
---
Hey all, I’m brand new to software engineering. I’m in the process of building a CLI gem using Ruby. I’m running in to a snag getting a few gems to bundle, namely nokogiri and pry. I’m running macOS 10.15.2

If anyone can help or point me in the right direction I’d greatly appreciate it!
## [8][Trying to understand advantages of dry-monads](https://www.reddit.com/r/ruby/comments/eo4829/trying_to_understand_advantages_of_drymonads/)
- url: https://www.reddit.com/r/ruby/comments/eo4829/trying_to_understand_advantages_of_drymonads/
---
Hey there!

Today i decided to try dry-monads for the first time but i don't really get if it's advantage is purely around visual aspects of the code or is there more to it? I'm basing my opinion on the only real-life examples i was able to find on the internet which was always about a following refactoring or similar:

Original code:

    if current_user &amp;&amp; current_user.info &amp;&amp; current_user.info.bio
      return current_user.info.bio 
    end

Refactored code (i'm pretty sure that's not the actual dry-monads code, just a pseudo-code to visualize what i'm talking about:

    return Maybe(current_user).map { |user| user.info }
                              .map { |info| info.bio }
                              .or_else { 'No bio set' }

And that makes me think...what's the profit? Is that better than writing it as below?

    return current_user&amp;.info&amp;.bio || 'No bio set'

Pardon my ignorance, it's probably because i can't find any real-life examples of some more complex refactors using dry family. No matter what - i have a strong feeling i'm not getting something in here fully. Thanks in advance for any clues!
## [9][Best guide/tutorial/course for Ruby?](https://www.reddit.com/r/ruby/comments/enyhnr/best_guidetutorialcourse_for_ruby/)
- url: https://www.reddit.com/r/ruby/comments/enyhnr/best_guidetutorialcourse_for_ruby/
---
Hi, I'm a frontend developer, and this year I want to learn backend development, I choosed Ruby, with Ruby on Rails

But for now I want to learn Ruby, can you please recommend the best Ruby guide/tutorial/course for you?

Thanks.
## [10][Reality Show on Ruby Programming](https://www.reddit.com/r/ruby/comments/enywfx/reality_show_on_ruby_programming/)
- url: https://www.reddit.com/r/ruby/comments/enywfx/reality_show_on_ruby_programming/
---
Hi

I'm doing a video series on creating a web application using Ruby on Rails. It's not meant to be a tutorial but more of a reality TV show on how one would go about creating an application.

I was just curious if any of you know anything similar (preferably in Ruby) or anyone doing the same thing (coding each step of the way w/ mistakes). 

Playlist link here:

[https://www.youtube.com/playlist?list=PL2-7U6BzddIZ35bJdCFx6RZ-QR8n\_JD82](https://www.youtube.com/playlist?list=PL2-7U6BzddIZ35bJdCFx6RZ-QR8n_JD82)

Videos ongoing. Trying to do a video once a week or if I have time. For this one, I'm trying to build a book keeping system.
