# ruby
## [1][Better Git diff output for Ruby, Python, Elixir, Go and more](https://www.reddit.com/r/ruby/comments/je09u7/better_git_diff_output_for_ruby_python_elixir_go/)
- url: https://tekin.co.uk/2020/10/better-git-diff-output-for-ruby-python-elixir-and-more?src=twitter
---

## [2][Pass a struct with a string field to a native library from Ruby, with pre-initialized string field, via FFI](https://www.reddit.com/r/ruby/comments/jdwkmz/pass_a_struct_with_a_string_field_to_a_native/)
- url: https://www.reddit.com/r/ruby/comments/jdwkmz/pass_a_struct_with_a_string_field_to_a_native/
---
FFI gem.

&amp;#x200B;

I have ruby code which passes a struct to a native Rust library. A struct contains a string field. I need to be able to specify string on Ruby side.

            class MyStruct &lt; FFI::Struct
              layout :s1, :string,
                     :field2, :uint32,
                     # other fields

I've tried this:

             def initialize(ruby_str)
               self[:s1] = ruby_str
               # [.......]

And have gotten this:

        `[]=': Cannot set :string fields (ArgumentError)

And this blows segmentation fault:

        def initialize(ruby_str)
          p1 = FFI::MemoryPointer.from_string(ruby_str)
          self[:content] = p1
          self[:len] = ruby_str.length

&amp;#x200B;

I've read about this issue, but I've not found a clear solution, nor example. I need to be able to pass a string field **at least from Ruby (one way, that is)**.

&amp;#x200B;

And, additionally, also convert a string  received from a native Rust library (not NULL terminated, but with a length specified in a separate field) into a Ruby string.

&amp;#x200B;

How can it be done then?
## [3][Join our Ruby based Game Jam!](https://www.reddit.com/r/ruby/comments/jdj4zv/join_our_ruby_based_game_jam/)
- url: https://itch.io/jam/teenytiny-dragonruby-minigamejam-2020
---

## [4][Modern Ruby book recommendation?](https://www.reddit.com/r/ruby/comments/jdf374/modern_ruby_book_recommendation/)
- url: https://www.reddit.com/r/ruby/comments/jdf374/modern_ruby_book_recommendation/
---
Hello!

I was using Ruby quite extensively few years ago, I believe I got good grasp of it. I had some break because every company I worked for in the meantime was \*Pythonized\*. I would like to get back to this beautiful technology - and get to know all the good stuff that happened in between. I was thinking of reading a book on Ruby, cover to cover. Which would you recommend to me? I'd like it to cover at least Ruby 2.5. 

Caveat: I'm not interested in RoR at all, I just enjoy using Ruby as a scripting language over Python. I'm a C++ / Rust developer.
## [5][Introducing Boring Generators](https://www.reddit.com/r/ruby/comments/jdk8n7/introducing_boring_generators/)
- url: https://www.abhaynikam.me/posts/introducing_boring_generators/
---

## [6][Idiomatic Ruby through Leetcode](https://www.reddit.com/r/ruby/comments/jdhqy6/idiomatic_ruby_through_leetcode/)
- url: https://medium.com/@michaelviveros/idiomatic-ruby-through-leetcode-5ff395a0ef15
---

## [7][Set up ActionText and Trix Editor on Rails 6 Application](https://www.reddit.com/r/ruby/comments/jd9vuk/set_up_actiontext_and_trix_editor_on_rails_6/)
- url: https://www.youtube.com/watch?v=SMsqFRc25gw&amp;t=9s
---

## [8][How ERB astutely uses `binding` to preprocess .html.erb files 📁](https://www.reddit.com/r/ruby/comments/jdbqv4/how_erb_astutely_uses_binding_to_preprocess/)
- url: https://medium.com/rubycademy/context-binding-in-ruby-fa118ea62269
---

## [9][“How I generated incomes as a web developer without working as a freelancer”](https://www.reddit.com/r/ruby/comments/jcuube/how_i_generated_incomes_as_a_web_developer/)
- url: https://medium.com/the-developers-journey/my-3-revenue-streams-as-a-developer-without-freelance-work-c5135dfa515d
---

## [10][football-to-sqlite tool - load / read football.txt match datafiles into a sqlite database](https://www.reddit.com/r/ruby/comments/jcuhyl/footballtosqlite_tool_load_read_footballtxt_match/)
- url: https://github.com/sportdb/football.db/tree/master/football-to-sqlite
---

