# ruby
## [1][99 Bottles of OOP - 2nd Edition Released](https://www.reddit.com/r/ruby/comments/hpzdxf/99_bottles_of_oop_2nd_edition_released/)
- url: https://sandimetz.com/99bottles
---

## [2][Looking for some kind of "quick" database implementation](https://www.reddit.com/r/ruby/comments/hqf3wb/looking_for_some_kind_of_quick_database/)
- url: https://www.reddit.com/r/ruby/comments/hqf3wb/looking_for_some_kind_of_quick_database/
---
Hey everyone. Some quick background: I lead a QA team full of ruby-ists where we build and maintain end to end tests in Webdriver and Capybara. We also have a set of "monitoring" scripts that are run via Cron.

These monitoring scripts are run each and every minute. Pieces of the script require a snapshot of the "state" of the system (basic things like true/false flags + statuses). So we thought it would be clever to just write them to text file.

That seems to work fine, however we've added so many flags that it's getting very messy to maintain, and the code is getting more and more difficult to understand.

I was thinking of writing a JSON representation of this information to a file, but it seems like what makes the most sense is to run a basic database implementation in order to store the 'state' of what our scripts are monitoring.

We're intimately familiar with MongoDB, but that seems really "heavy" for our needs. Is there something where I can just add a gem to our bundle and hit the ground running very quickly? Looking to avoid any new installations to our servers also.

I know that Rails has its own db implementation, but we have 0 familiarity with Rails and activerecord.
## [3][A Ruby Developer’s Adventures in Elixir](https://www.reddit.com/r/ruby/comments/hq5jkk/a_ruby_developers_adventures_in_elixir/)
- url: https://medium.com/@minghz42/a-ruby-developers-adventures-in-elixir-515380986bc3
---

## [4][What does the 'take' keyword do?](https://www.reddit.com/r/ruby/comments/hq7wlh/what_does_the_take_keyword_do/)
- url: https://www.reddit.com/r/ruby/comments/hq7wlh/what_does_the_take_keyword_do/
---
Edit: Solved - looks like this is a Shopify-specific keyword.

I'm trying to learn the basic Ruby and came across this line: 

     discounted_item = line_item.split(take: count)

`Count` is a variable that equates to 0 or 1.

Can anyone explain to me what the `take` parameter is called so I can Google it? I tried Googling a few things but can't seem to find what this is called or what it does.

Edit: In this context, `line_item` is one of the values in an array of items in a shopping cart, one line per item. Each line_item refers to a Shopify object, which looks like this when logged: 
`#&lt;LineItem:0x7febeed017a0 @grams=0, @original_line_price=#&lt;Money:0x7febeed01c20&gt;, @quantity=2, @discounts=[], @properties_was={}, @line_price=#&lt;Money:0x7febeed01d70&gt;, @adjustments=[], @variant=#&lt;Variant:0x7febeed022b0&gt;, @line_price_was=#&lt;Money:0x7febeed01d70&gt;, @source_indices={"32713989750828:24f88f9647288e16aec75d4feed7c66a"=&gt;2}, @properties={}&gt;`
## [5][Introducing Devpack - custom development stack for Bundler](https://www.reddit.com/r/ruby/comments/hpvwzn/introducing_devpack_custom_development_stack_for/)
- url: https://www.reddit.com/r/ruby/comments/hpvwzn/introducing_devpack_custom_development_stack_for/
---
Hi, Rubyists. I would like to introduce *Devpack*, a gem that allows developers to create a local `.devpack` file (either in the working directory or any parent directory) which specifies the developer's preferred development toolchain (i.e. a list of gems).

This allows a single gem (which has no external dependencies) to be added to a `Gemfile` instead of each developer adding their favourite debugging gems. This helps avoid conflicting functionality and simplifies dependency resolution while still allowing for flexibility and personal preference.

The gem is available here: [https://rubygems.org/gems/devpack](https://rubygems.org/gems/devpack)

And the source code is here: [https://github.com/bobf/devpack](https://github.com/bobf/devpack)

Feedback, questions, and bug reports are very welcome !

Enjoy the rest of the weekend.
## [6][Fun Facts about Ruby #10: Joining array elements without using Array#join](https://www.reddit.com/r/ruby/comments/hpsfzc/fun_facts_about_ruby_10_joining_array_elements/)
- url: https://i.redd.it/664sl5gtnea51.png
---

## [7][I am tired of hearing that Ruby is fine](https://www.reddit.com/r/ruby/comments/hpta1o/i_am_tired_of_hearing_that_ruby_is_fine/)
- url: https://www.reddit.com/r/ruby/comments/hpta1o/i_am_tired_of_hearing_that_ruby_is_fine/
---
Let me prefix this with words of gratitude. You are all wonderful people and a very welcoming community.

Ruby is not dying but its health is not great either.

Is Ruby and Ruby on Rails dead 2020?

No but just by the title alone you see how tightly coupled with RoR Ruby is. People no longer consider the happiness factor that Ruby brings relevant in the grand scheme of things.

1) Name some CLI applications written in Ruby that are pleasant to use.
I have recently tried Vagrant and oh boy is it a slow and dreadful thing.
More often than not the CLI/DevOps tooling is not really liked that much by the wider community of developers and users. It is tolerated but not liked.

2) No growth curve
This is simply a sad reality. Ruby no longer offers solutions to people's problems.

a) Can I build a desktop app with Ruby? Yes I can but it's not pleasant and people are not really doing it anymore anyway.

b) Can I build a mobile app with Ruby? Yes but it's an entirely different can of worms and not even the biggest Ruby shops are working with that technology. Looking at you RubyMotion. It's simply because the experience and community around it is super small.

c) Building CLI tools and devops tooling? Yes people do this but the user experience is usually bad and nobody really likes to work with those tools written in Ruby.

d) frontend? Almost non existent.

e) backend? Mostly RoR and projects we never hear about 

3) Naming the big companies. Github,Stripe,Shopify.

Yes they use Ruby but we hear more and more microservices that are written in Go and other faster languages doing the heavy lifting. People are not interested in writing C extensions and then gluing it with Ruby code. It sucks. The experience is bad.

4) Ruby is regional.
While Ruby is healthy and fine in places like the US,London,Berlin it's rapidly disappearing from Singapore, Asia in general or the rest of the Western world.
Even if you like Ruby and want to write Ruby you just simply can't in some parts of the world.

5) Ruby will not get significantly faster and if it will so will the other languages.
Yes Ruby is not getting faster anytime soon. At least not in a sense that would be significantly noticeable by most of people.
If Ruby gets faster with GraalVM, well so will Python and whatever GraalVM supports.

6) Ruby community is dived, fractured and its presence on the internet is weak.

Just look at this sub. This sub has fewer visitors than some irrelevant niche subs.
I have not been able to find a decent Discord or IRC. They all have very few visitors.
This signifies that Ruby is no longer interesting to newbies.

7) All the low hanging fruits are gone and significant improvements are hard to get right
All the 3x3, JIT and other perf improvements are really appreciated but nowhere near as important to most of real world production applications in RoR.

8) Lack of vision
We keep hearing about all those great additions to Ruby like pipe operator,anonymous struct literals etc...I think they make the language less readable and ergonomic. They bring more problems than solutions to be honest.
Also they are mostly "cosmetic" changes which could be the main problem.


To me Ruby as a project is complete. I have lowered my expectations for Ruby and realized I will never be able to use it in areas where I would use other programming languages. I have realized that the constant tech/context switching is not worth it anymore. 

Things like JavaScript or C# are not as nice to work with but they allow me to do stuff I could not really do with Ruby.
Ruby to me is like Swift. Really nice to work with if you can (iOS apps) but absolutely disappointing for anything else (android,windows).
## [8][I am tired of hearing that Ruby is dead](https://www.reddit.com/r/ruby/comments/hp3yar/i_am_tired_of_hearing_that_ruby_is_dead/)
- url: https://www.reddit.com/r/ruby/comments/hp3yar/i_am_tired_of_hearing_that_ruby_is_dead/
---
Ruby has always been my favorite development language, which is why seeing all the posts about its death kind of sucks. I wrote an article on whether [Ruby is not dead in 2020](https://syndicode.com/2020/07/08/why-is-ruby-still-our-choice-in-2020-2/), but developers on my team didn't seem to agree on all the points. What do you think?
## [9][How should I best convert Article.last into a URL path?](https://www.reddit.com/r/ruby/comments/hpbqn9/how_should_i_best_convert_articlelast_into_a_url/)
- url: https://www.reddit.com/r/ruby/comments/hpbqn9/how_should_i_best_convert_articlelast_into_a_url/
---
I'm going to apologise in advance, if this is an exceedingly basic question that I should be able to answer myself. I've been learning ruby/rails for a hot minute and this is the first time I'm trying to go off-script from the course.

For a blog site I have set up 4 cards on the home page for the 4 latest articles, each showing a title and a truncated section of the text. I can display the title and text easily enough for each of the articles I want, e.g.:

    &lt;%= Article.last.title %&gt;
    &lt;%= simple_format(Article.last.description.truncate(100)) %&gt;

And then similar code for the other 3 cards - i.e. second_to_last, third_to_last, Article.order(id: :desc).fourth

**HOWEVER**, my problem starts when I then try to create a link for a show action, so you can click through and read the article. I've been trying to google-fu and RTFM my way out of it but I'm a little lost on where to actually look. I am still searching, but am hoping this speeds up the process.

Is there any chance anyone could send me to a specific page of documentation or guide on exactly how to do this? If you could explain it then that would be amazing, but a boot in the right direction is great.

I guess a bonus question is: should I actually be referencing database entries like that directly on a page, or should I be putting it somewhere more out of the way?
## [10][Testing views in Sinatra application?](https://www.reddit.com/r/ruby/comments/hph1rj/testing_views_in_sinatra_application/)
- url: https://www.reddit.com/r/ruby/comments/hph1rj/testing_views_in_sinatra_application/
---
We have some legacy code that keeps breaking because the code in the views are not checking for empty values in the response from an API. Here is an example from the SLIM file that was causing an error:  


     - @response.keys.each do |account|
        - if !@params.empty?
          - next unless @params.include? account
        - if @cache[account][0] == 200

While we have gone back and added checks, we still might be missing some. 

I know how to do TDD with MiniTest and Ruby for unit testing, but I've never implemented testing with SLIM views. What testing framework should I use to mock up different responses with missing data to make sure the view logic in the SLIM files are good? Cucumber?
