# ruby
## [1][Question using ruby2d Image Class making Tic-Tac-Toe](https://www.reddit.com/r/ruby/comments/fb8tnw/question_using_ruby2d_image_class_making_tictactoe/)
- url: https://www.reddit.com/r/ruby/comments/fb8tnw/question_using_ruby2d_image_class_making_tictactoe/
---
New to Ruby but not to programming and so I wanted to mess around and make something with ruby2d. The project I went for was tic-tac-toe but I ran into a situation handling the Image class. 

I made a similar project using python + pygame and in using that framework there is an Image class that you can load an image to and then draw that image with a specific function to do so at specified x and y values. In the example of tic-tac-toe, this allows me to store the board state in a way where I can do something like this:

    # board is a 1d array of symbols (x or o)
    # x_image is an Image object storing the image for the x symbol
    # o_image is an Image object storing the image for the o symbol
    
    # Using Ruby syntax for simplicity
    for i in (0...board.length) # board being a 1d array
        if board[i] == "x"
            draw(x_image, x, y) # x and y positions can be arbitrary
        elsif board[i] == "o"
            draw(o_image, x, y)
        end
    end

In ruby2d, drawing occurs on object creation and so I thought I would have to take a different approach. The main idea is to show a symbol (x or o depending on turn) when the user clicks a free square. In order to do this I created a 1d array of 9 Image objects, one for each square.

However there are issues with this approach. The first being that image paths seemed to be mandatory on Image creation, so I decided to set all the images to one path and then switch to the other path when the game switches turns. This lead me to my second issue which is that image paths cannot be changed once the Image object is created.

This all does make sense however it left me with my current solution:

    # On user clicking on the board
    if ... # free board square is available
        ... # Update board
        
        cell_images[square_number] = Image.new(
            IMAGE_PATHS[turn], # Image paths holds both paths, one for image_x and one
                               # for image_o
            x: ... , # Arbitrary x value
            y: ...   # Arbitrary y value
        )
    end

I really am not a fan of this implementation because on each mouse click on a free square a new Image object is created each time for each square, all holding the same image information (excluding x and y position).

For tic-tac-toe this issue isn't that bad but the solution doesn't seem scale-able for larger games (Go for example).

If there is a way to implement this where I have one x\_image and o\_image object that would be preferable. I imagine I'm missing some simple way to do this or I might be misunderstanding something obvious but any help and constructive criticism is appreciated.
## [2][Reduce allocations for keyword argument hashes](https://www.reddit.com/r/ruby/comments/fb2e6u/reduce_allocations_for_keyword_argument_hashes/)
- url: https://github.com/ruby/ruby/pull/2945
---

## [3][2.7's pattern matching official docs (recently merged)](https://www.reddit.com/r/ruby/comments/favshb/27s_pattern_matching_official_docs_recently_merged/)
- url: https://docs.ruby-lang.org/en/master/syntax/pattern_matching_rdoc.html
---

## [4][Benchmarking Ruby 2.7.0's Numbered Parameters](https://www.reddit.com/r/ruby/comments/fat90p/benchmarking_ruby_270s_numbered_parameters/)
- url: https://blog.schembri.me/post/benchmarking-ruby-2.7.0-numbered-parameters/
---

## [5][Faster Excel Parsing in Ruby](https://www.reddit.com/r/ruby/comments/faszl8/faster_excel_parsing_in_ruby/)
- url: https://blog.schembri.me/post/faster-excel-parsing-in-ruby/
---

## [6][I get nil for sort_by on an array of objects](https://www.reddit.com/r/ruby/comments/far4kd/i_get_nil_for_sort_by_on_an_array_of_objects/)
- url: https://www.reddit.com/r/ruby/comments/far4kd/i_get_nil_for_sort_by_on_an_array_of_objects/
---
    def sorted_objects
          @objects.sort_by do |l, r|
            if !l || !r
              @sort_results = 0
            else
              @sort_results = 1
            end
          end.to_a
          #p sort_results
         # p @objects
        end

so I have printed r and it shows nil every time

here's an example of what objects is storing  


    @objects &lt;&lt; Room.new(10 - x, y, tile.to_i, 220, 300, @max_x.to_i, @max_y.to_i,"floor",0)

no idea why r returns nil.
## [7][Has anyone used Kiba?](https://www.reddit.com/r/ruby/comments/fal2l2/has_anyone_used_kiba/)
- url: https://www.reddit.com/r/ruby/comments/fal2l2/has_anyone_used_kiba/
---
Hey! I've been looking for data engineering tools, found tons of them, to manage workflows, to generate ETL jobs, most of them use Python and I'd rather use Ruby. (i.e. amazon glue, airflow, astronomer, stitchdata, ...)

A few weeks ago, the Ruby Weekly featured Kiba v3 https://github.com/thbar/kiba 

I've been reading thorough their wiki and I feel very inclined to start using it, although I would also need to set up a server to run these ETL tasks, maybe a serverless instance or something similar. 

Just to give you some broad context, my main goal is to extract data from a tsv file, transform the data into a set of attributes I want, and lastly store the transformed data into a csv file so that later I can easily load it into redshift.

My question is, do you have any experiences to share regarding data engineering in Ruby? Or have you used Kiba previously? Any insights or comments? 

Anything would be appreciated!
## [8][The state of ruby blogs and news (2020/2 Edition) - 36 channels, 1464 items](https://www.reddit.com/r/ruby/comments/fahwbr/the_state_of_ruby_blogs_and_news_20202_edition_36/)
- url: https://www.reddit.com/r/ruby/comments/fahwbr/the_state_of_ruby_blogs_and_news_20202_edition_36/
---
Hello,

   I've put together a little survey about the state of the ruby feed-iverse that includes personal blogs, ruby project news and more.

  See the [planet.ini](https://github.com/planetruby/planet/blob/master/planet.ini) for all feeds included in the survey.

Q: What feed formats are in use?

    Formats  (n=36)
    ---------------------------------
      atom        23 (63%) | *************************************
      rss 2.0     13 (36%) | *********************


Q: What servers are in use?

    Servers  (n=36)
    ---------------------------------
      GitHub.com  12 (33%) | *******************
      nginx        8 (22%) | *************
      Cowboy       5 (13%) | *******
      Apache       4 (11%) | ******
      cloudflare   3 ( 8%) | ****
      GSE          2 ( 5%) | ***
      Netlify      2 ( 5%) | ***


Q: What (web site) publishing tools are in use?

    Generators  (n=36)
    ---------------------------------
      ?           22 (61%) | ************************************
      Jekyll       9 (25%) | ***************
      wordpress    3 ( 8%) | ****
      webgen       1 ( 2%) | *
      Ghost        1 ( 2%) | *


Q: What's the update frequency of posts?

    Update Frequency  (n=35)
    ---------------------------------
      &lt;=   7 days   3 ( 8%) | ****
      &lt;=  14 days   4 (11%) | ******
      &lt;=  30 days   7 (20%) | ************
      &lt;=  90 days  12 (34%) | ********************
      &lt;= 180 days   5 (14%) | ********
      &lt;= 365 days   4 (11%) | ******

That's all for now. Comments and ideas on how to improve
the stats are more than welcome. Cheers. Prost.
## [9][Building a reddit bot using Ruby](https://www.reddit.com/r/ruby/comments/fam6w7/building_a_reddit_bot_using_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fam6w7/building_a_reddit_bot_using_ruby/
---
I apologize if this has been asked recently, I had trouble finding anything semi-current on the subject. If there's recent threads on this that I didn't find, please let me know as I'd love to see them.

I'm a newish Ruby developer and I have been toying with the idea of making a simple reddit bot with Ruby, mainly to help me learn various skills, but also to be functional and serve a sub I mod. I want it to fetch info from one outside site and post the data it finds there when requested. Simple enough I hope (not simple for me lol).

I'm having trouble finding much info on how to implement a reddit bot using Ruby. I know Python is what's used most often, and I do have some experience with it but I'm focusing on Ruby atm and I don't want to go off on a Python tangent. I found the [Reddit\_Bot](https://rubygems.org/gems/reddit_bot) gem which seems to be popular and recently updated. I'm going through the docs which are helpful and have a lot of examples, but there's basic info missing that is probably common knowledge to more experienced devs but a vast abyssal mystery to newbies like me.

So to be specific:

1. Is Ruby okay to use for this, or should I set aside cash now for when I inevitably smash my computer in frustration?
2. If (1) is yes, are there any videos, tutorials, anything you know of that are Ruby-specific that might be helpful for this project? (I did already watch several youtube vids on Python bots which were helpful, but since I'm still a baby dev I have a hard time translating what I learn there to Ruby context)
3. Any advice, tips, or whatever you have for this would be appreciated. I have mainly worked in a Rails framework for general web development and have only a little experience working with API's, so pretty much every part of this process is going to be a learning experience for me.

Thanks!
## [10][Ruby Hash#transform_keys now accepts a hash that maps existing keys to new keys](https://www.reddit.com/r/ruby/comments/fae2jf/ruby_hashtransform_keys_now_accepts_a_hash_that/)
- url: https://blog.saeloun.com/2020/02/26/ruby-hash-transform_keys-now-accepts-a-hash-that-maps-existing-keys-to-new-keys
---

