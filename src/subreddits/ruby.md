# ruby
## [1][Modern Ruby book recommendation?](https://www.reddit.com/r/ruby/comments/jdf374/modern_ruby_book_recommendation/)
- url: https://www.reddit.com/r/ruby/comments/jdf374/modern_ruby_book_recommendation/
---
Hello!

I was using Ruby quite extensively few years ago, I believe I got good grasp of it. I had some break because every company I worked for in the meantime was \*Pythonized\*. I would like to get back to this beautiful technology - and get to know all the good stuff that happened in between. I was thinking of reading a book on Ruby, cover to cover. Which would you recommend to me? I'd like it to cover at least Ruby 2.5. 

Caveat: I'm not interested in RoR at all, I just enjoy using Ruby as a scripting language over Python. I'm a C++ / Rust developer.
## [2][Set up ActionText and Trix Editor on Rails 6 Application](https://www.reddit.com/r/ruby/comments/jd9vuk/set_up_actiontext_and_trix_editor_on_rails_6/)
- url: https://www.youtube.com/watch?v=SMsqFRc25gw&amp;t=9s
---

## [3][How to release a resource acquired in the contructor, automatically, whenever GC destroys an object?](https://www.reddit.com/r/ruby/comments/jddfmv/how_to_release_a_resource_acquired_in_the/)
- url: https://www.reddit.com/r/ruby/comments/jddfmv/how_to_release_a_resource_acquired_in_the/
---
In my ruby code I acquire a resource in the constructor. A resource is acquired via FFI:

        class MyFFI
          extend FFI::Library
        
          # [.........]
    
          attach_function :get_external_resource [], :uint32
          attach_function :destroy_external_resource, :uint32, :void
        end
        
        class MyRecource
          def initialize
            @res1 = MyFFI.get_external_resource
          end
    
          def my_destroy
            MyFFI.destroy_external_resource(@res1)
          end
        end

Then somewhere in code:

        rs1 = MyRecource.new

How can I ensure that **rs1.my\_destroy()** will be called automatically whenever an object **rs1** gets destroyed by Garbage collector?
## [4][How ERB astutely uses `binding` to preprocess .html.erb files 📁](https://www.reddit.com/r/ruby/comments/jdbqv4/how_erb_astutely_uses_binding_to_preprocess/)
- url: https://medium.com/rubycademy/context-binding-in-ruby-fa118ea62269
---

## [5][“How I generated incomes as a web developer without working as a freelancer”](https://www.reddit.com/r/ruby/comments/jcuube/how_i_generated_incomes_as_a_web_developer/)
- url: https://medium.com/the-developers-journey/my-3-revenue-streams-as-a-developer-without-freelance-work-c5135dfa515d
---

## [6][football-to-sqlite tool - load / read football.txt match datafiles into a sqlite database](https://www.reddit.com/r/ruby/comments/jcuhyl/footballtosqlite_tool_load_read_footballtxt_match/)
- url: https://github.com/sportdb/football.db/tree/master/football-to-sqlite
---

## [7][Scalable Concurrency for Ruby 3](https://www.reddit.com/r/ruby/comments/jcg8cu/scalable_concurrency_for_ruby_3/)
- url: https://www.youtube.com/watch?v=Y29SSOS4UOc
---

## [8][Error Handling in Ruby: Part I](https://www.reddit.com/r/ruby/comments/jcsa2q/error_handling_in_ruby_part_i/)
- url: https://medium.com/rubycademy/error-handling-in-ruby-part-i-557898185e2f
---

## [9][relational match pattern?](https://www.reddit.com/r/ruby/comments/jciwww/relational_match_pattern/)
- url: https://www.reddit.com/r/ruby/comments/jciwww/relational_match_pattern/
---
Hey so I just found out that C#9 is in the makings and I saw this snippet

https://preview.redd.it/gb2bcogj1jt51.png?width=381&amp;format=png&amp;auto=webp&amp;s=c83cd5d43ddbb206eb581e9ca784bbdc1c86d221

Is there any chance I'll get to write ruby like that? (Or is there a way to write code like that that I'm not aware of?)
## [10][Guide to Reactive Rails](https://www.reddit.com/r/ruby/comments/jc5k8g/guide_to_reactive_rails/)
- url: https://github.com/obie/guide-to-reactive-rails
---

