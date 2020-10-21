# ruby
## [1][RuboCop 1.0](https://www.reddit.com/r/ruby/comments/jfaxse/rubocop_10/)
- url: https://metaredux.com/posts/2020/10/21/rubocop-1-0.html
---

## [2][Explaining magic behind popular Ruby code](https://www.reddit.com/r/ruby/comments/jekiry/explaining_magic_behind_popular_ruby_code/)
- url: https://longliveruby.com/articles/the-magic-behind-ruby-code
---

## [3][How to ship a ruby gem/library with an external custom native library as a dependency?](https://www.reddit.com/r/ruby/comments/jeou5f/how_to_ship_a_ruby_gemlibrary_with_an_external/)
- url: https://www.reddit.com/r/ruby/comments/jeou5f/how_to_ship_a_ruby_gemlibrary_with_an_external/
---
I'm building a library/gem in Ruby that calls some functions from a native **custom** library written in C++. A C++ library, in a pre-compiled form, is available for Windows, Linux and MacOS, and, of course, they're 3 different files.

&amp;#x200B;

**Questions**:

how should I go about shipping my gem with with 1 or all 3 at once C++ libraries? 

Should I put them into a subdirectory of my gem? Then, how would I allow a user to download a gem with a C++ library that's been pre-compiled for his OS rather than including all 3 at once?

&amp;#x200B;

Or require a user to download it on its own? Then how would get a path to that library, how would a user pass it into my gem?
## [4][How to create a wrapper for a function that returns a result in a callback (block or proc)?](https://www.reddit.com/r/ruby/comments/jel38y/how_to_create_a_wrapper_for_a_function_that/)
- url: https://www.reddit.com/r/ruby/comments/jel38y/how_to_create_a_wrapper_for_a_function_that/
---
I have a third-party library, in which there's a function that returns \`void\` directly, but has a block that's used as a callback:

        def my_get_data1(arg1)
          Thread.new do
            ExternalLib.get_data_by_block("arg1") do |a, b, c|
              # [ result of computation and some other data comes here ]
            end
          end
        end

&amp;#x200B;

And there's the same kind of function that accepts a \`proc\` instead, and they both work identically:

&amp;#x200B;

        MyProc = Proc.new do |a, b, c|
            # [ result of computation and some other data comes here ]
        end
    
        def my_get_data2(arg1)
          Thread.new do
            ExternalLib.get_data_by_proc(arg1, MyProc)
          end
        end

&amp;#x200B;

**my\_get\_data1()** and **my\_get\_data2()** serve the purpose of an interface or wrapper of \`ExternalLib.get\_data()\` for the end user, this is what I need.

&amp;#x200B;

Q: How can I actually allow a user to get access to the result of the computation -- to what's inside of the \`block\` or \`proc\`?

&amp;#x200B;

I want it to be something like this:

        p1 = Proc.new do |a|
          # [ user will access result here, via a]
        end
    
        my_get_data1(arg1, p1)
    
    
        my_get_data2(arg1) do do |a|
          # [ user will access result here, via a]
        end

&amp;#x200B;

P.S.

I've found out that that I have to use **Thread.new** because otherwise, it'll hang when I run it as "ruby my\_script1.rb", probably due to GIL. In \`irb\` it won't, though.
## [5][What are your guys thoughts on Crystal and Elixir?](https://www.reddit.com/r/ruby/comments/je90yt/what_are_your_guys_thoughts_on_crystal_and_elixir/)
- url: https://www.reddit.com/r/ruby/comments/je90yt/what_are_your_guys_thoughts_on_crystal_and_elixir/
---
Hi everyone, hoping to get some insights into what you think of the two languages that seem to be popular with Rubyists.

No judgement just interested to see if anyone has experience with them and what they think about them.

For me I really like Crystal but is in a very new state and not even 1.0 so is a long time away from will pick up any major steam in terms of libs (complete personal opinion however!). 

Elixir on the other hand goes into the functional world which we are all told is better but I still have issues not explaining concepts in terms of objects and state not being held closely to those objects. I love [Brian Goetz's talk on OOP vs FP](https://www.youtube.com/watch?v=HSk5fdKbd3o) where he thinks we should be object orientated on the outside and functional within.

What do you guys think? Anyone have experience with them?
## [6][Ruby code hangs when run in the terminal via "ruby script.rb", whereas in Irb it doesn't hang](https://www.reddit.com/r/ruby/comments/jenlag/ruby_code_hangs_when_run_in_the_terminal_via_ruby/)
- url: https://www.reddit.com/r/ruby/comments/jenlag/ruby_code_hangs_when_run_in_the_terminal_via_ruby/
---
A simplied version of the code I have is this:

        def my_get_data1(arg1)
          Thread.new do
            ExternalLib.get_data_by_block("arg1") do |a, b, c|
              # [ result of computation and some other data comes here ]
            end
          end
        end
    

I behaves differently in the terminal and in Irb. When run as a ruby script in the terminal - \`ruby ./script1.rb\` it hangs. When I copy the code of \`script1.rb\` into Irb and thus run it, it doesn't hang and returns a result.

&amp;#x200B;

I've heard that ruby code can behave different in such a case.

&amp;#x200B;

What can be an issue or cause?
## [7][Better Git diff output for Ruby, Python, Elixir, Go and more](https://www.reddit.com/r/ruby/comments/je09u7/better_git_diff_output_for_ruby_python_elixir_go/)
- url: https://tekin.co.uk/2020/10/better-git-diff-output-for-ruby-python-elixir-and-more?src=twitter
---

## [8][How to Estimate The Size of a Rails Application](https://www.reddit.com/r/ruby/comments/je718a/how_to_estimate_the_size_of_a_rails_application/)
- url: https://www.fastruby.io/blog/rails/code-quality/how-we-estimate-rails-application-size.html
---

## [9][tutor/mentor for ruby](https://www.reddit.com/r/ruby/comments/jedvcc/tutormentor_for_ruby/)
- url: https://www.reddit.com/r/ruby/comments/jedvcc/tutormentor_for_ruby/
---
Hi,

Looking for someone who is an expert in Ruby who would like to be my tutor/mentor. Currently working a cli/api project, and reading "Practical Object-Oriented Design" by Sandi Metz. In need of someone to talk to about everything and get some help. Thanks :)
## [10][OpenAPI generator in Ruby, client code](https://www.reddit.com/r/ruby/comments/jebsyw/openapi_generator_in_ruby_client_code/)
- url: https://www.reddit.com/r/ruby/comments/jebsyw/openapi_generator_in_ruby_client_code/
---
Hello Reddit crowd!

I couldn't find explanation about this, so I try my luck here!

I'm trying the `openapi-generator` (4.3.1). It generates models from the API description file. 

For instance, it generates this:

    module OpenapiClient
      class Thing
        attr_accessor :id
        attr_accessor :description
    
        # Attribute mapping from ruby-style variable name to JSON key.
        def self.attribute_map
          {
            :'id' =&gt; :'id',
            :'description' =&gt; :'description'
          }
        end
    
        # Attribute type mapping.
        def self.openapi_types
          {
            :'id' =&gt; :'Integer',
            :'description' =&gt; :'String'
          }
        end
    
        # ...
      end
    end

This behaves as expected:

    # All this is expected:
    t = OpenapiClient::Thing.new(id: 1)
    t.valid?
    =&gt; false
    t.list_invalid_properties
    =&gt; ["invalid value for \"description\", description cannot be nil."]

But **this does not behave as expected**:

    t = OpenapiClient::Thing.new(id: "abc", description: 3)
    t.valid?
    =&gt; true

`id` should be int and `description` a string, so it fails to check the type. Is there anything I'm doing wrong?

🙏
