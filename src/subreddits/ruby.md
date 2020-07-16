# ruby
## [1][Defining gem dependencies in single-file Ruby scripts](https://www.reddit.com/r/ruby/comments/hrzy5d/defining_gem_dependencies_in_singlefile_ruby/)
- url: https://nithinbekal.com/posts/bundler-inline-gemfile/
---

## [2][Ruby VS Python](https://www.reddit.com/r/ruby/comments/hs2nsk/ruby_vs_python/)
- url: https://www.reddit.com/r/ruby/comments/hs2nsk/ruby_vs_python/
---
More than a versus (it's more striking) I would really like to know what things they have in common and how they differ, what their main advantages are compared to the other and of course their opinion of them. One point to take into account is because they believe that python is more used in AI and Machine learning developments.
## [3][System of a test: Proper browser testing in Ruby on Rails](https://www.reddit.com/r/ruby/comments/hs54c6/system_of_a_test_proper_browser_testing_in_ruby/)
- url: https://evilmartians.com/chronicles/system-of-a-test-setting-up-end-to-end-rails-testing
---

## [4][Patching Ruby? How to handle mixing Module#prepend with method aliasing](https://www.reddit.com/r/ruby/comments/hs20zh/patching_ruby_how_to_handle_mixing_moduleprepend/)
- url: https://www.mayerdan.com/ruby/2020/07/15/ruby-gems-patching-std-lib
---

## [5][.times method - how to describe what it does?](https://www.reddit.com/r/ruby/comments/hrtcc7/times_method_how_to_describe_what_it_does/)
- url: https://www.reddit.com/r/ruby/comments/hrtcc7/times_method_how_to_describe_what_it_does/
---
Hi all, Ruby beginner here. I am preparing for an assessment where I need to describe what happens on each line of code and got a bit stuck with the .times method. Specifically the following code: 

`a = 5` 

`3.times do |n|` 

  `a = 3` 

`end` 

`puts a`

Here is my description: 

On line 1 variable `a` is initialized in the outer scope which references an Integer Object with the value `5`.

**On line 2 we call the .times method on Integer 3 and pass the do..end block to it as an argument. We pass |n| as an argument as well.** The times method iterates the given block 3 times, passing the values from 0 to 2 and returns self (3 in this case).

On line 3 the variable `a` is reassigned to Integer `3`, so it points to a different object.

On line 5 method `puts` is called passing variable `a` as an argument.

The method outputs `3` and returns `nil`. The variable `a` is available to the inner scope created by `3.times do..end` which allowed the code to re-assign the value of `a`.

This example demonstrates local variable scoping Rules in Ruby - variables initialized in the outer scope are accessible inside the block.

\--- 

**My question is - is it correct to say that we call the times method on integer 3 and pass block do..end and |n| as an argument?**
## [6][Looking for feedback on new gem: ops_team](https://www.reddit.com/r/ruby/comments/hrqjpi/looking_for_feedback_on_new_gem_ops_team/)
- url: https://www.reddit.com/r/ruby/comments/hrqjpi/looking_for_feedback_on_new_gem_ops_team/
---
Hey all,

I've been working on a gem for a while that aims to make it easy to provide simple and discoverable automation for your projects, without complexities like DSLs or namespaces. It's only config file is a single, plain `YAML` file in the root of your repo. The tagline is "The operations team for your project".

[https://github.com/nickthecook/ops](https://github.com/nickthecook/ops)

    actions:
      connect-to-db:
        command: psql ssl_mode=verify-full -h $DB_HOST -U $DB_USER -d $DB_DATABASE
      alias: db
      load_secrets: true

It's written in Ruby, but works for projects written in any language. I started it because I was struggling to automate a `terraform` project properly (`terraform` doesn't have tools for loading environment-specific config or secrets like Rails does). Also, a company I worked at had a tool like this, and it was a huge help to have a single tool to work with thousands of repos. Sadly, that tool was private.

We've been using \`ops\_team\` at my new company for a few months now, but I'd love feedback from a broader audience. Please have a look and lmk what you think.
## [7][Custom directory for Webpacker in Rails apps](https://www.reddit.com/r/ruby/comments/hrpvxe/custom_directory_for_webpacker_in_rails_apps/)
- url: https://prathamesh.tech/2020/07/15/custom-directory-for-webpacker-in-rails-apps/
---

## [8][Is it possible to make a nested loop more efficient than O(N^2)](https://www.reddit.com/r/ruby/comments/hrxzbw/is_it_possible_to_make_a_nested_loop_more/)
- url: https://www.reddit.com/r/ruby/comments/hrxzbw/is_it_possible_to_make_a_nested_loop_more/
---
Consider a 'game', where the player jumps from a wooden post to a wooden post. Each post has a number on it that that describes the number of posts to jump over. On positive numbers the player jumps to the right, on negative numbers to the left.  

The number of posts is limited, thus it may happen that the player jumps into the water. 

Input: `input_array = [[1,1,1,1,-2], [1,1,1,1], [1,1,-1], [1,0,1,0]]`

So, my current solution is:

- loop over the input array
  - run a nested while loop until there can be made successful jump: check each integer and determine whether the player stays dry or jumps into the water


```
    input_array.each do |posts|
    
    # create a new round 
    	round = Round.new(posts)
    	# initialize new iterator
    	iterator = round.iterator
    	
    	while iterator.has_next?
    		# add current index to visited posts list
    		iterator.visited &lt;&lt; iterator.ind
    		# update next index
    		iterator.update_next_index
    		# set the index for the next cycle
    		iterator.ind = iterator.next_index
    		# update the current element
    		iterator.element = iterator.array[iterator.ind]
    	end
    end

```



    class WoodenPostsIterator
    	include IteratorHelpers
    
    	attr_reader :array, :visited
    	attr_accessor :ind, :next_index, :element
    
      def initialize(array)
        @array = array
        @ind = 0
        @visited = []
        @element = array[@ind]
        @next_index = nil
      end
    
      # @return [boolean] indicating whether there is a next post
      def has_next?
      	# player stops jumping if they come across number '0'
    		if post_zero?(element)
    			puts "The player stays dry due to innactivity"
    			return false
    		end
    
    		# check if post already visited: if yes, stop the program
    		if post_is_visited?(ind)
    			puts "The player stays dry."
    			return false
    		end
    
    		next_index = calculate_next_index
    		# halt the game if the player jumped into the water
    		if water_jump(array, next_index)
    			puts "The player gets wet!"
    			return false 
    		end
    
    		true
      end
      
      # @return [nil] 
      def update_next_index
      	# update the next index
    		self.next_index = calculate_next_index
      end
    
      private
      # @return [Integer] indicating next index
      def calculate_next_index
      	# determine the sign
    		sign = define_operator(element)
    		# calculate next index (next wooden post)
    		ind.public_send(sign, element.abs)
      end
    
      # @param [Integer] post for checking
    	# @return [boolean] indicating whether the passed element is zero
     	def post_zero?(element)
     		# halt the program if the element is zero
    		element == 0
    	end
    
    	# @param [Integer] index for checking
    	# @return [Boolean]  indicating whether 'visited' list contains passed index 
    	def post_is_visited?(ind)
    		# if node (post) has been visited, halt the program
    		# reason: The player stays dry due to the infinite jumping
    		# NOTE: this method can be made to be more efficient by removing visited 
    		# posts from the array. 
    		visited.include?(ind)
    	end
    
    	# @param [Array] posts list
    	# @param [Integer] next index
    	# @return [Boolean] indicating if the player is jumping in the water
    	def water_jump(arr, next_index)
    		# returns true if player jumps into the 
    		# water (try to access an element out of the range)
    		true if next_index + 1 &gt; arr.length || next_index &lt; 0
    	end
    end

This solution is `O(N^2)`. How can I make it more efficient?
## [9][System of a test: Proper browser testing in Ruby on Rails. Ditch Selenium for Ferrum and CDP, embrace Docker.](https://www.reddit.com/r/ruby/comments/hr2fyl/system_of_a_test_proper_browser_testing_in_ruby/)
- url: https://evilmartians.com/chronicles/system-of-a-test-setting-up-end-to-end-rails-testing
---

## [10][Experiences using slack-ruby-client?](https://www.reddit.com/r/ruby/comments/hra3zb/experiences_using_slackrubyclient/)
- url: https://www.reddit.com/r/ruby/comments/hra3zb/experiences_using_slackrubyclient/
---
Evening all,

I'm going to be building a Slack integration between my SaaS Rails app and our customer's Slack channels. Initially it's going to be basic slash command but if anticipate making it much more interactive over time.

In the past I've done a basic Slack posting without any special gems / library. I recently came across [**slack-ruby-client**](https://github.com/slack-ruby/slack-ruby-client) and it certainly looks fully featured and potentially useful.

I'm interested to hear if anyone has used this library and how was their experience? 

thanks!
