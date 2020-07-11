# ruby
## [1][I am tired of hearing that Ruby is dead](https://www.reddit.com/r/ruby/comments/hp3yar/i_am_tired_of_hearing_that_ruby_is_dead/)
- url: https://www.reddit.com/r/ruby/comments/hp3yar/i_am_tired_of_hearing_that_ruby_is_dead/
---
Ruby has always been my favorite development language, which is why seeing all the posts about its death kind of sucks. I wrote an article on whether [Ruby is not dead in 2020](https://syndicode.com/2020/07/08/why-is-ruby-still-our-choice-in-2020-2/), but developers on my team didn't seem to agree on all the points. What do you think?
## [2][Integrate activemodel serializers - API only ruby on rails course (chapter 4)](https://www.reddit.com/r/ruby/comments/hp7dai/integrate_activemodel_serializers_api_only_ruby/)
- url: https://duetcode.io/rails-api-only-course/integrate-activemodel-serializers
---

## [3][Ruby-on-Rails beginner question: How do I get the value of :id?](https://www.reddit.com/r/ruby/comments/hp2bvz/rubyonrails_beginner_question_how_do_i_get_the/)
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
## [4][Replacing a range of objects within an array](https://www.reddit.com/r/ruby/comments/hp9che/replacing_a_range_of_objects_within_an_array/)
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
## [5][How to create native encrypted attributes for Rails with ActiveModel attributes](https://www.reddit.com/r/ruby/comments/hosxlt/how_to_create_native_encrypted_attributes_for/)
- url: https://alexc.link/blog/native-encrypted-attributes-for-rails-active-model
---

## [6][All You Need to Know About Comments in Ruby - Tarka Labs](https://www.reddit.com/r/ruby/comments/horh49/all_you_need_to_know_about_comments_in_ruby_tarka/)
- url: https://medium.com/tarkalabs/all-you-need-to-know-about-comments-in-ruby-97d991714cf3
---

## [7][How to "kill" a `rake task` within `binding.pry` instance?](https://www.reddit.com/r/ruby/comments/hovnmm/how_to_kill_a_rake_task_within_bindingpry_instance/)
- url: https://www.reddit.com/r/ruby/comments/hovnmm/how_to_kill_a_rake_task_within_bindingpry_instance/
---
So I'm trying my hand at writing my first custom rake task. I put in a `binding.pry` call to help me test things.

I'm using `CTRL-D` to advance past the `binding.pry` stop.

HOWEVER, there are over 4,000 items I'm having it run through. 

How can I "kill/stop" the entire task from progressing further? Otherwise, I've gotta sit here and hold down `CTRL-D` for a good amount of time :p

Thanks for pointers!
## [8][Ruby Chat Help](https://www.reddit.com/r/ruby/comments/hoz0pe/ruby_chat_help/)
- url: https://www.reddit.com/r/ruby/comments/hoz0pe/ruby_chat_help/
---
I'm having a heck of a time getting back into Ruby, and thought hey...maybe there are people online who would have time and interest in helping via real time chat!
Anyone who's interested, perhaps a discord server could be started.
## [9][Scraping a URL Fragment with Nokogiri?](https://www.reddit.com/r/ruby/comments/hoj7ce/scraping_a_url_fragment_with_nokogiri/)
- url: https://www.reddit.com/r/ruby/comments/hoj7ce/scraping_a_url_fragment_with_nokogiri/
---
Hi, I have a project due so this is time sensitive, I'm a coding student.

Nokogiri scraping data:

If I want to select to scrape by css, but for like, just a filtered list that appears with a url fragment tag, how do I format that?

example:[https://www.domain.com/phones/mobile#smartphone](https://www.domain.com/phones/mobile#smartphone)

how would i format that for nokogiri using doc.css?  I see the filtered list is under its own div and has a "class=selected" but the info I need that is display is in a different div.  Any help appreciated, I'm really exhausted and stayed up until 5am last night despite having work today, and I really am stuck.  I don't want to fail out of my class.

&amp;#x200B;

I'm trying to get around the following error 

\`NoMethodError: undefined method \`\[\]' for nil:NilClass\` - it's failing because of the pagination of the filtered lists, so I just wanted to select one of the url fragment filters and then my program will work like it should and I can go to bed.
## [10][Open Source Status Update – April-June (dry-schema, dry-validation, dry-rails, rom-rb 6.0 and more)](https://www.reddit.com/r/ruby/comments/ho1k88/open_source_status_update_apriljune_dryschema/)
- url: https://solnic.codes/2020/07/09/open-source-status-update-april-june-2020/
---

