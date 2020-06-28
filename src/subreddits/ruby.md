# ruby
## [1][Just started to publish API-only ruby on rails course. We’re building an Instapaper clone from scratch.](https://www.reddit.com/r/ruby/comments/hhb1ib/just_started_to_publish_apionly_ruby_on_rails/)
- url: https://www.reddit.com/r/ruby/comments/hhb1ib/just_started_to_publish_apionly_ruby_on_rails/
---
I'm happy that I just published the first two chapters of **API-only** ruby on rails course that I [mentioned](https://www.reddit.com/r/ruby/comments/gpebe0/im_planning_to_build_a_rails_course_would_you/) almost one month ago on this subreddit. 🎉🎉

In this course, we will build a functional clone of Instapaper from scratch. Also, it targets the developers already familiar with rails tech stack and has a written format. It's an ongoing course, and I'm planning to share new chapters weekly basis.

You can reach the content from [https://duetcode.io/rails-api-only-course](https://duetcode.io/rails-api-only-course). Even though it's just two chapters, I wanted to share as early as possible to get some insights and feedback from the audience.

Please, let me know if you have any thoughts/feedback about the content or the website in general.
## [2][tapping_device - A new tool that helps you debug with ease](https://www.reddit.com/r/ruby/comments/hh93ss/tapping_device_a_new_tool_that_helps_you_debug/)
- url: https://www.reddit.com/r/ruby/comments/hh93ss/tapping_device_a_new_tool_that_helps_you_debug/
---
Knowing what your objects "exactly" do can be hard when debugging a ruby program. Either using \`puts\` or debuggers requires you to jump between files and insert debugging code everywhere.

That's why I created a gem called [tapping\_device](https://github.com/st0012/tapping_device), which can track objects and print all the traces in a nice format with just 1 line of code.  


https://preview.redd.it/9m2n2eib7l751.png?width=1200&amp;format=png&amp;auto=webp&amp;s=a69b519a2f8ba78a2e576e58ffe3c5f732570a40
## [3][rubocop and docker in vscode](https://www.reddit.com/r/ruby/comments/hhalj1/rubocop_and_docker_in_vscode/)
- url: https://www.reddit.com/r/ruby/comments/hhalj1/rubocop_and_docker_in_vscode/
---
Hello Folks, any idea how i can include rubocop in rails app with docker ?

update Problem Solved with : 
https://code.visualstudio.com/docs/remote/containers
https://stelligent.com/2020/03/20/development-acceleration-through-vs-code-remote-containers-an-introduction/
## [4][Gladiator (Glimmer Editor) - Ugliest Text Editor Ever!](https://www.reddit.com/r/ruby/comments/hgve8k/gladiator_glimmer_editor_ugliest_text_editor_ever/)
- url: https://www.reddit.com/r/ruby/comments/hgve8k/gladiator_glimmer_editor_ugliest_text_editor_ever/
---
&amp;#x200B;

[Gladiator \(Glimmer Editor\)](https://preview.redd.it/9kstmde20h751.png?width=1233&amp;format=png&amp;auto=webp&amp;s=90b3e5edebba2916b4d220145ed3959374bae4ce)

[https://github.com/AndyObtiva/glimmer-cs-gladiator](https://github.com/AndyObtiva/glimmer-cs-gladiator)
## [5][Deploy Ruby On Rails On Google Cloud](https://www.reddit.com/r/ruby/comments/hgtok7/deploy_ruby_on_rails_on_google_cloud/)
- url: https://youtu.be/3d7xBvmu6Z4
---

## [6][Fun Facts about Ruby #6](https://www.reddit.com/r/ruby/comments/hgokpp/fun_facts_about_ruby_6/)
- url: https://i.redd.it/r9rmstmyde751.png
---

## [7][How to return the output of a If/Else in ruby](https://www.reddit.com/r/ruby/comments/hgqxln/how_to_return_the_output_of_a_ifelse_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/hgqxln/how_to_return_the_output_of_a_ifelse_in_ruby/
---
I have a method that returns one this if it is true and another thing if it is false.

    def has_perm?(item)
      return 
          if(true)
            item.this?
          else 
            item.that?
          end
    end

I wanted to write like this but it doesn't work. What is some clean way I can write code for this? The below looks fine but seems unnecessary assignment

    def has_perm?(item)
      result =  
          if(true)
            item.this?
          else 
            item.that?
          end
      result
    end
## [8][If you're using GraphQL with ruby these generators will save you many keystrokes](https://www.reddit.com/r/ruby/comments/hgb82l/if_youre_using_graphql_with_ruby_these_generators/)
- url: https://www.reddit.com/r/ruby/comments/hgb82l/if_youre_using_graphql_with_ruby_these_generators/
---
I released this library of graphql rails generators a few months ago after doing an ungodly amount of typing on a data-heavy application backed by a graphql API. The generators look at your Rails models and generate graphql types, mutations and input types from your database scheme (currently only activerecord supported, but other databases would be trivial to add). 

Hope people find some use from this and saves you a few thousand keystrokes. Cheers.

[https://github.com/ajsharp/graphql-rails-generators](https://github.com/ajsharp/graphql-rails-generators)
## [9][Add NEWS entries about JIT optimizations · ruby/ruby@200c5f4 · GitHub](https://www.reddit.com/r/ruby/comments/hg8787/add_news_entries_about_jit_optimizations/)
- url: https://github.com/ruby/ruby/commit/200c5f4075cb1d179c2eba5b30b5b0a500870f67
---

## [10][Acts As Tracked: Selectively track changes made on your ActiveRecord Models](https://www.reddit.com/r/ruby/comments/hg8evn/acts_as_tracked_selectively_track_changes_made_on/)
- url: https://www.reddit.com/r/ruby/comments/hg8evn/acts_as_tracked_selectively_track_changes_made_on/
---
Hi, i've made a gem to selectively track changes on AR models, where audited gem would be an overkill. ActsAsTracked can be plugged into ActiveRecord model, and then used whenever you need a history of changes and actors made on the record.

You can find docs in the [repository](https://github.com/ramblingcode/acts-as-tracked)

Blog post [here](https://www.ramblingcode.dev/posts/track_changes_on_your_activerecord_models/)

Example usage in this [project](https://github.com/ramblingcode/rails6-acts-as-tracked-usage)

Hope you find it useful.
