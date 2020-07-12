# ruby
## [1][Fun Facts about Ruby #10: Joining array elements without using Array#join](https://www.reddit.com/r/ruby/comments/hpsfzc/fun_facts_about_ruby_10_joining_array_elements/)
- url: https://i.redd.it/664sl5gtnea51.png
---

## [2][I am tired of hearing that Ruby is fine](https://www.reddit.com/r/ruby/comments/hpta1o/i_am_tired_of_hearing_that_ruby_is_fine/)
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
## [3][I am tired of hearing that Ruby is dead](https://www.reddit.com/r/ruby/comments/hp3yar/i_am_tired_of_hearing_that_ruby_is_dead/)
- url: https://www.reddit.com/r/ruby/comments/hp3yar/i_am_tired_of_hearing_that_ruby_is_dead/
---
Ruby has always been my favorite development language, which is why seeing all the posts about its death kind of sucks. I wrote an article on whether [Ruby is not dead in 2020](https://syndicode.com/2020/07/08/why-is-ruby-still-our-choice-in-2020-2/), but developers on my team didn't seem to agree on all the points. What do you think?
## [4][How should I best convert Article.last into a URL path?](https://www.reddit.com/r/ruby/comments/hpbqn9/how_should_i_best_convert_articlelast_into_a_url/)
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
## [5][Testing views in Sinatra application?](https://www.reddit.com/r/ruby/comments/hph1rj/testing_views_in_sinatra_application/)
- url: https://www.reddit.com/r/ruby/comments/hph1rj/testing_views_in_sinatra_application/
---
We have some legacy code that keeps breaking because the code in the views are not checking for empty values in the response from an API. Here is an example from the SLIM file that was causing an error:  


     - @response.keys.each do |account|
        - if !@params.empty?
          - next unless @params.include? account
        - if @cache[account][0] == 200

While we have gone back and added checks, we still might be missing some. 

I know how to do TDD with MiniTest and Ruby for unit testing, but I've never implemented testing with SLIM views. What testing framework should I use to mock up different responses with missing data to make sure the view logic in the SLIM files are good? Cucumber?
## [6][Integrate activemodel serializers - API only ruby on rails course (chapter 4)](https://www.reddit.com/r/ruby/comments/hp7dai/integrate_activemodel_serializers_api_only_ruby/)
- url: https://duetcode.io/rails-api-only-course/integrate-activemodel-serializers
---

## [7][Replacing a range of objects within an array](https://www.reddit.com/r/ruby/comments/hp9che/replacing_a_range_of_objects_within_an_array/)
- url: https://www.reddit.com/r/ruby/comments/hp9che/replacing_a_range_of_objects_within_an_array/
---
I'm a bit of a newb so forgive my perhaps simple question.

Is there an easy way to replace multiple objects within an array with another object?

I want to start with an array that will return 0-9 and then replace 4, 5 &amp; 6 with say "apple" so that the array would return 0, 1, 2, 3, "apple", "apple", "apple",  7, 8, 9

My code:

array1 = Array.new(10) { | i | i.to_i }

p array1.map { | i | i == 4 ? "apple" : i }

Will return 0-9 but replace 4 with "apple" but when I want to replace a range...

p array1.map { | i | i == 4..6 ? "apple" : i }

It returns "warning: integer literal in conditional range"
and 0-3 but the 5th through to 10th index are all "apple"

Which I don't understand why it does that.

Or do I just have to replace index's one at a time?🤔
## [8][Ruby-on-Rails beginner question: How do I get the value of :id?](https://www.reddit.com/r/ruby/comments/hp2bvz/rubyonrails_beginner_question_how_do_i_get_the/)
- url: https://www.reddit.com/r/ruby/comments/hp2bvz/rubyonrails_beginner_question_how_do_i_get_the/
---
I'm trying to learn Ruby on Rails to broaden my horizon a bit and become a generally better programmer. My native tongue is Python and I usually make websites in Flask. Right now I'm trying to put together a simple REST api to host some files in directories off my disk. I have this route from rails:

post GET    /posts/:id(.:format)                                                                     posts#show

which then maps into my controller:

&amp;#x200B;

    class PostsController &lt; ApplicationController
    	def index
    		@posts = Array.new
    		@directories = Dir["/home/narco/Code/blog/*-*-*"]
    		for directory in @directories
    			@post = {}
    			@post["date"] = directory[22..32]
    			info = JSON.load File.open directory+"/info.json"
    			@post['title'] = info['title']
    			@post['cover'] = info['cover']
    			@post['author'] = info['author']
    			@posts.push @post
    		end
    		render :json =&gt; @posts
    	end
    	
    	def show
    		info = JSON.load File.open "/home/narco/Code/blog/" + :id + "/info.json"
    		render :json =&gt; info
    	end
    end

The show method is where I have run into my problems. In flask I would pass in the url as a variable, in this case I assumed ":id" and then I could do this string concatenation like I am doing here. I have come to understand that symbols are a very different thing, there are plenty of guides to their differences online so I'm reading those right now. But I feel like even if I understand what a symbol is I am approach this problem wrong. I'm approaching it like a python programmer not a Ruby programmer. Can some one give me some guidance on how I am supposed to go about this specific task?
## [9][How to create native encrypted attributes for Rails with ActiveModel attributes](https://www.reddit.com/r/ruby/comments/hosxlt/how_to_create_native_encrypted_attributes_for/)
- url: https://alexc.link/blog/native-encrypted-attributes-for-rails-active-model
---

## [10][All You Need to Know About Comments in Ruby - Tarka Labs](https://www.reddit.com/r/ruby/comments/horh49/all_you_need_to_know_about_comments_in_ruby_tarka/)
- url: https://medium.com/tarkalabs/all-you-need-to-know-about-comments-in-ruby-97d991714cf3
---

