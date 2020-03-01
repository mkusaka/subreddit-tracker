# ruby
## [1][StackOverflow listed Ruby in the most commonly used programming languages](https://www.reddit.com/r/ruby/comments/fbsvvy/stackoverflow_listed_ruby_in_the_most_commonly/)
- url: https://learnworthy.net/stackoverflow-listed-the-most-commonly-used-programming-languages/
---

## [2][RubyKaigi 2020 postponed until 3–5 September due to COVID-19](https://www.reddit.com/r/ruby/comments/fbewtn/rubykaigi_2020_postponed_until_35_september_due/)
- url: https://esa-pages.io/p/sharing/68/posts/1006/b15a58c675f5a69d06e5.html
---

## [3][Can you download the fuel economy [dot] gov database?](https://www.reddit.com/r/ruby/comments/fblo87/can_you_download_the_fuel_economy_dot_gov_database/)
- url: https://www.reddit.com/r/ruby/comments/fblo87/can_you_download_the_fuel_economy_dot_gov_database/
---
I was thinking about making a headless watir build to scrape everything they've got and join it with current kbb data, but I thought to maybe ask here first if someone hasn't already done this or similar or know of a resource I'm not considering.
## [4][Reduce allocations for keyword argument hashes](https://www.reddit.com/r/ruby/comments/fb2e6u/reduce_allocations_for_keyword_argument_hashes/)
- url: https://github.com/ruby/ruby/pull/2945
---

## [5][Question using ruby2d Image Class making Tic-Tac-Toe](https://www.reddit.com/r/ruby/comments/fb8tnw/question_using_ruby2d_image_class_making_tictactoe/)
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
## [6][2.7's pattern matching official docs (recently merged)](https://www.reddit.com/r/ruby/comments/favshb/27s_pattern_matching_official_docs_recently_merged/)
- url: https://docs.ruby-lang.org/en/master/syntax/pattern_matching_rdoc.html
---

## [7][Benchmarking Ruby 2.7.0's Numbered Parameters](https://www.reddit.com/r/ruby/comments/fat90p/benchmarking_ruby_270s_numbered_parameters/)
- url: https://blog.schembri.me/post/benchmarking-ruby-2.7.0-numbered-parameters/
---

## [8][Faster Excel Parsing in Ruby](https://www.reddit.com/r/ruby/comments/faszl8/faster_excel_parsing_in_ruby/)
- url: https://blog.schembri.me/post/faster-excel-parsing-in-ruby/
---

## [9][I get nil for sort_by on an array of objects](https://www.reddit.com/r/ruby/comments/far4kd/i_get_nil_for_sort_by_on_an_array_of_objects/)
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
## [10][Has anyone used Kiba?](https://www.reddit.com/r/ruby/comments/fal2l2/has_anyone_used_kiba/)
- url: https://www.reddit.com/r/ruby/comments/fal2l2/has_anyone_used_kiba/
---
Hey! I've been looking for data engineering tools, found tons of them, to manage workflows, to generate ETL jobs, most of them use Python and I'd rather use Ruby. (i.e. amazon glue, airflow, astronomer, stitchdata, ...)

A few weeks ago, the Ruby Weekly featured Kiba v3 https://github.com/thbar/kiba 

I've been reading thorough their wiki and I feel very inclined to start using it, although I would also need to set up a server to run these ETL tasks, maybe a serverless instance or something similar. 

Just to give you some broad context, my main goal is to extract data from a tsv file, transform the data into a set of attributes I want, and lastly store the transformed data into a csv file so that later I can easily load it into redshift.

My question is, do you have any experiences to share regarding data engineering in Ruby? Or have you used Kiba previously? Any insights or comments? 

Anything would be appreciated!
