# ruby
## [1][Rails 6 adds rails db:system:change command](https://www.reddit.com/r/ruby/comments/ez5gxl/rails_6_adds_rails_dbsystemchange_command/)
- url: https://blog.saeloun.com/2020/02/04/rails-6-adds-rails-db-system-change-command
---

## [2][news.rb - news quick starter script - build your own newsfeed in 1-2-3 steps in minutes](https://www.reddit.com/r/ruby/comments/ez6irg/newsrb_news_quick_starter_script_build_your_own/)
- url: https://github.com/feedreader/news.starter
---

## [3][Exploring Big-O Notation With Ruby](https://www.reddit.com/r/ruby/comments/eywrym/exploring_bigo_notation_with_ruby/)
- url: https://www.reddit.com/r/ruby/comments/eywrym/exploring_bigo_notation_with_ruby/
---
You know Big-O is important - not only for acing your next job interview but for knowing how code works at scale. But have you ever taken the time to go beyond a superficial understanding of the subject? In this article, Julie Kent uses equal parts math and Ruby to reveal the beating heart of Big-O and show us how it ticks.

[https://www.honeybadger.io/blog/big-o-notation-ruby/](https://www.honeybadger.io/blog/big-o-notation-ruby/)
## [4][Handling Ruby 2.7 deprecations warnings](https://www.reddit.com/r/ruby/comments/eywvfw/handling_ruby_27_deprecations_warnings/)
- url: https://kukicola.io/posts/handling-ruby-27-deprecations-warnings
---

## [5][Rails 6.1 introduces class_names helper](https://www.reddit.com/r/ruby/comments/ez3y6c/rails_61_introduces_class_names_helper/)
- url: https://blog.bigbinary.com/2020/02/04/rails-6-1-introduces-class_names-helper.html
---

## [6][Help With error](https://www.reddit.com/r/ruby/comments/eyzljh/help_with_error/)
- url: https://www.reddit.com/r/ruby/comments/eyzljh/help_with_error/
---
I'm trying some simple exercise to capitalize book titles, with some rules.

When I test most of the of it it's ok, but one case has made me go nuts.

Here is the logic of the Class:

    class Book
      # write your code here
    
      def title
        @title
      end
    
      def title=title_str
        @title = title_str
    
        if is_there_space?
          new_str = ""
          split_words = @title.split(" ") #array made
    
          p split_words
    
          split_words.each do |word|
            word = word.to_s
    
            if word == "i"
              # p "i converted"
              new_str += word.capitalize! + " " #capitalize 'I'
    
            elsif word == "and" or word  == "a" or word == "an" or word == "in" or word == "the" or word == "of"
              # p "articles/prepositions converted"
              new_str += word + " "
    
            elsif word.is_a?(Integer)
              p "elsif integer"
    
            else
              # p word + " last else"
              new_str += word.capitalize! + " "
            end
    
            p new_str
          end
    
          @title = new_str[0].capitalize!.to_s + new_str.slice(0, new_str.length).to_s
          @title.strip!
    
        else # just one word capitalize
          @title.capitalize!
    
        end
      end
    
      def is_there_space?
        @title.include?(" ")
        # puts "True"
      end
    
    end
    
    
    # tests
    @book = Book.new
    # @book.title = "the man in the iron mask" #ok
    # @book.title = "stuart little" #ok
    @book.title = "what i wish i knew when i was 20" # ERROR
    # @book.title = "love in the time of cholera" #ok
    # @book.title = "to eat an apple a day" #ok
    
    
    
    p @book.title
    

When I test against :

    @book.title = "love in the time of cholera"

I returns as intended, but when there's a number in the final of the string, like:

    @book.title = "what i wish i knew when i was 20"

I got:

    Traceback (most recent call last):
    	3: from book.rb:58:in `&lt;main&gt;'
    	2: from book.rb:17:in `title='
    	1: from book.rb:17:in `each'
    
    book.rb:33:in `block in title=': undefined method `+' for nil:NilClass (NoMethodError)

Can anyone point me what I'm missing?

Thanks
## [7][I'd like help putting together a game plan for learning Ruby](https://www.reddit.com/r/ruby/comments/ez0gvg/id_like_help_putting_together_a_game_plan_for/)
- url: https://www.reddit.com/r/ruby/comments/ez0gvg/id_like_help_putting_together_a_game_plan_for/
---
Experience: I've taken 2 1/2 courses of beginner programming in C++ in college. And I completed half a book on beginner Ruby a few years back and forgot most of it.

I'm about to give codecademy a try. I know nothing about rails, and I've seen people suggest that you "branch out" from codecademy. I'm a stickler for feeling confident that I'm learning things efficiency through the best resources. I've recently learned about codewars and I think that will be fantastic for learning with as well. In what order should I use what resources to learn Ruby and RoR? 

I've recently dropped out of College with my Wife, she got a paid internship with a startup (ACH transfer platform) through a friend but won't be programming for another few months...SO, I want to get a head start learning Ruby so I can help my wife learn more quickly while the leeches at her company play tug-o-war with her time. Any tips/resources are welcome!
## [8][Bundler default version issue](https://www.reddit.com/r/ruby/comments/eyuj75/bundler_default_version_issue/)
- url: https://www.reddit.com/r/ruby/comments/eyuj75/bundler_default_version_issue/
---
Edit: Seems like what I'm referring to is "expected behavior" (thanks to @hsbt from the ruby/core team): https://github.com/rubygems/bundler/issues/7601

Thanks for your help, guys. I guess I need to better understand Ruby tooling!

---

Hey r/ruby!

I'm running into an error with the default bundle version after installing ruby 2.7.0 via rbenv.

The problem is that, when I run `gem install bundler`, it installs the latest version 2.1.4, but still defaults to 2.1.2:

    gem list
    
    *** LOCAL GEMS ***
    
    bundler (2.1.4, default: 2.1.2)
    .
    .
    .

I've tried to run a `gem install bundler:2.1.4 --default`, but it just sets both 2.1.4 AND 2.1.2 as default, which results in not being able to uninstall either version.

Any insight on how to fix this problem?
## [9][pluto-news gem - newsreader for easy (re)use - build your own facebook newsfeed in 5 minutes](https://www.reddit.com/r/ruby/comments/eyrb81/plutonews_gem_newsreader_for_easy_reuse_build/)
- url: https://github.com/feedreader/pluto/tree/master/pluto-news
---

## [10][Ruby by the Bay](https://www.reddit.com/r/ruby/comments/eyisk4/ruby_by_the_bay/)
- url: https://www.reddit.com/r/ruby/comments/eyisk4/ruby_by_the_bay/
---
Hi all,

I wanted to let everyone know about [Ruby by the Bay](https://rubybythebay.org/), the west coast version of Ruby for Good that will be happening in the SF area April 3 - April 6. Put simply, we gather a bunch of rubyists and javascripters (along with anyone else interested in doing good) and we get together for a long weekend (Fri - Mon) of using our awesome skills to build apps for social and civic good organizations. Last year we built applications for diaper banks, Refuge Restrooms, animal shelters, and many other great places. We're focusing on projects this year that affect our environment -- organizations that cater to those doing the important work of saving our planet.

The event is all inclusive, meaning that we stay in dorms, all the food is provided and we have pretty decent internet too. Our days are filled with hacking and our evenings are full of excellent camaraderie -- playing boardgames, video games, karaoke, werewolf, cards and lots of excellent socializing!

We are excited to announce that we're going to be holding the event at [Naturebridge](https://naturebridge.org/locations/golden-gate) in Marin County.

You can find more information about this years' event here: [https://rubybythebay.org/attend](https://rubybythebay.org/attend) Here is a list of the most frequently asked questions here: [https://rubybythebay.org/faq](https://rubybythebay.org/faq)

I'd be delighted to answer any questions.

Happiness,

Sean (and the rest of the Ruby by the Bay team)

PS - We'd love any sponsorship if anyone knows of anyone who may be interested sponsoring our event this year!
