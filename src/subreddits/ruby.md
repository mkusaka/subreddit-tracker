# ruby
## [1][Scalable Concurrency for Ruby 3](https://www.reddit.com/r/ruby/comments/jcg8cu/scalable_concurrency_for_ruby_3/)
- url: https://www.youtube.com/watch?v=Y29SSOS4UOc
---

## [2][“How I generated incomes as a web developer without working as a freelancer”](https://www.reddit.com/r/ruby/comments/jcuube/how_i_generated_incomes_as_a_web_developer/)
- url: https://medium.com/the-developers-journey/my-3-revenue-streams-as-a-developer-without-freelance-work-c5135dfa515d
---

## [3][football-to-sqlite tool - load / read football.txt match datafiles into a sqlite database](https://www.reddit.com/r/ruby/comments/jcuhyl/footballtosqlite_tool_load_read_footballtxt_match/)
- url: https://github.com/sportdb/football.db/tree/master/football-to-sqlite
---

## [4][Error Handling in Ruby: Part I](https://www.reddit.com/r/ruby/comments/jcsa2q/error_handling_in_ruby_part_i/)
- url: https://medium.com/rubycademy/error-handling-in-ruby-part-i-557898185e2f
---

## [5][relational match pattern?](https://www.reddit.com/r/ruby/comments/jciwww/relational_match_pattern/)
- url: https://www.reddit.com/r/ruby/comments/jciwww/relational_match_pattern/
---
Hey so I just found out that C#9 is in the makings and I saw this snippet

https://preview.redd.it/gb2bcogj1jt51.png?width=381&amp;format=png&amp;auto=webp&amp;s=c83cd5d43ddbb206eb581e9ca784bbdc1c86d221

Is there any chance I'll get to write ruby like that? (Or is there a way to write code like that that I'm not aware of?)
## [6][Guide to Reactive Rails](https://www.reddit.com/r/ruby/comments/jc5k8g/guide_to_reactive_rails/)
- url: https://github.com/obie/guide-to-reactive-rails
---

## [7][How to work with strings in Ruby to send to / receive from a Rust library, via FFI?](https://www.reddit.com/r/ruby/comments/jcbx8z/how_to_work_with_strings_in_ruby_to_send_to/)
- url: https://www.reddit.com/r/ruby/comments/jcbx8z/how_to_work_with_strings_in_ruby_to_send_to/
---
I have a library in Rust which I want to use in Ruby via \`ffi\` gem. A Rust library has got some functions that operate with pure Rust strings.

&amp;#x200B;

In C++ a Rust string, if it was exported from Rust library, would be represented as follows:

    
        typedef struct {
          const char* content;
          uint32_t len;
        } rust_string_t;

In Ruby I do this:

          class RustString &lt; FFI::Struct
            layout
              :content, :string,
              :len, :uint32
          end

&amp;#x200B;

But I also will need to convert normal Ruby strings to \`RustString\` back and forth. How would I do it? 

&amp;#x200B;

        def send_str_to_rust_func1(s1)
          rs1 = RustString.new
          rs1.content = s1
          rs1.len = s1.length # should I additional 1 to the  length ?
    
          # [......]
        end

and

        # returns string
        def get_str_from_rust_func2
          r_str = get_str_data_from_rust()
          # should I return r_str as is?
        end

&amp;#x200B;

Also, do I have to add 1 to the length of a string when converting to, or from, if at all?
## [8][Seed_pic the gem to make seeding images easy](https://www.reddit.com/r/ruby/comments/jcaayn/seed_pic_the_gem_to_make_seeding_images_easy/)
- url: https://www.reddit.com/r/ruby/comments/jcaayn/seed_pic_the_gem_to_make_seeding_images_easy/
---
I made my first gem this week and wanted to share it with you. I made a gem that helps seed image URLs into your database for you can use them in image tags to make your life easier. [seed_pic](https://rubygems.org/gems/seed_pic)
## [9][github analytics with the hubba-reports gem - What are your most used languages (in char / bytes count)?](https://www.reddit.com/r/ruby/comments/jc7zvy/github_analytics_with_the_hubbareports_gem_what/)
- url: https://www.reddit.com/r/ruby/comments/jc7zvy/github_analytics_with_the_hubbareports_gem_what/
---
Hello,  I have split the hubba github analytics gem into two, that is, [hubba](https://github.com/rubycoco/git/tree/master/hubba) and [hubba-reports](https://github.com/rubycoco/git/tree/master/hubba-reports) for easier (re)use and split the data gathering / collecting via github api calls and the report generation. Anyways, I have added a [new language report](https://github.com/rubycoco/git/blob/master/hubba-reports/lib/hubba/reports/reports/languages.rb) that lists all your languages used by char / bytes count and by number of repos. See [LANGUAGES.md](https://github.com/yorobot/backup/blob/master/LANGUAGES.md) as an example. Happy data crunching with ruby. Cheers. Prost. 

PS: Know any other alternative github scripts / gems, please tell.
## [10][Looking for feedback on a gem I have been working on: active_snapshot](https://www.reddit.com/r/ruby/comments/jbzqju/looking_for_feedback_on_a_gem_i_have_been_working/)
- url: https://www.reddit.com/r/ruby/comments/jbzqju/looking_for_feedback_on_a_gem_i_have_been_working/
---
I am looking for feedback on the design of a new "gem" that I have been working on called [`active_snapshot`](https://github.com/westonganger/active_snapshot)

**Gem:** [`active_snapshot`](https://github.com/westonganger/active_snapshot)

**Description:** Simplified snapshots and restoration for ActiveRecord models and associations with a transparent white-box implementation.

I appreciate your feedback and help.
