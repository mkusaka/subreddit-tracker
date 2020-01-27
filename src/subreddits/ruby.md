# ruby
## [1][Evaluating Ruby in Ruby](https://www.reddit.com/r/ruby/comments/euasxs/evaluating_ruby_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/euasxs/evaluating_ruby_in_ruby/
---
[https://ilyabylich.svbtle.com/evaluating-ruby-in-ruby](https://ilyabylich.svbtle.com/evaluating-ruby-in-ruby)
## [2][tty-box ignoring everything past a certain point.](https://www.reddit.com/r/ruby/comments/euh1dt/ttybox_ignoring_everything_past_a_certain_point/)
- url: https://www.reddit.com/r/ruby/comments/euh1dt/ttybox_ignoring_everything_past_a_certain_point/
---
Anyone work with tty-box? It seems to be ignoring everything past the title: line. It prints the box to the console, shows the title in both area's, but ignores everything past the first costomzation. 

Is it the formatting I'm using? There's no good examples in the documentation that has multple comtomizations applied. 

    def title_box
      main_title = TTY::Box.frame(
        width: 30, height: 10,
        title: {top_left: 'Title', bottom_right: 'v0.01',
        border: {
        top_left: :cross,
        top_right: :cross,
        bottom_left: :cross,
        bottom_right: :cross,
        style: {
      fg: :bright_yellow,
      bg: :blue,
      border: {
        fg: :bright_yellow,
        bg: :blue
        }}}
      }
    )
    
        print main_title
          main
      end
## [3][Ruby variable or method](https://www.reddit.com/r/ruby/comments/eua27z/ruby_variable_or_method/)
- url: https://medium.com/@igor04/ruby-variable-or-method-ac6814b4e8b1
---

## [4][planet.rb quick starter script updated (v1.0) - add articles, blogs to your static (jekyll &amp; friends) website via feeds (and planet pluto)](https://www.reddit.com/r/ruby/comments/euch1c/planetrb_quick_starter_script_updated_v10_add/)
- url: https://github.com/feedreader/planet.rb
---

## [5][Why the Sorbet typechecker is fast](https://www.reddit.com/r/ruby/comments/eu34es/why_the_sorbet_typechecker_is_fast/)
- url: https://blog.nelhage.com/post/why-sorbet-is-fast/
---

## [6][A Template for Starting New Projects on Ruby on Rails](https://www.reddit.com/r/ruby/comments/eu683v/a_template_for_starting_new_projects_on_ruby_on/)
- url: https://hackernoon.com/a-template-for-starting-new-projects-on-ruby-on-rails-t66c36iz
---

## [7][ActiveInteractor v1.0.0 Release](https://www.reddit.com/r/ruby/comments/eu52ao/activeinteractor_v100_release/)
- url: /r/rails/comments/eu4z1b/activeinteractor_v100_release/
---

## [8][Interactive Shell Examples?](https://www.reddit.com/r/ruby/comments/etwsrt/interactive_shell_examples/)
- url: https://www.reddit.com/r/ruby/comments/etwsrt/interactive_shell_examples/
---
Are there any examples of creating an interactive shell in Ruby? I'm wanting to create somethig like a command prompt, only I want to build it from the ground up so I can learn along the way. I've been searching for hours and finding only tutorials on the IRB, which is not what I'm after. 

I had started, however I lost all my code, books marks, etc, little as there was. Good oppurtunity to start fresh.
## [9][IRB breaking on Greek Characters.](https://www.reddit.com/r/ruby/comments/etz8bd/irb_breaking_on_greek_characters/)
- url: https://www.reddit.com/r/ruby/comments/etz8bd/irb_breaking_on_greek_characters/
---
Hello guys.

I have some scripts that are working normally in Ubuntu machines. But when I try to run them in a certain Windows machine the script just goes bananas.

The line that hangs is this

browser.input(:value =&gt; "Εκτύπωση").click and in the irb is showing up like thisbrowser.input(:value =&gt; "

It seems that any Greek character i place in irb it breaks everything.

Thanks

When i type any Greek Character in IRB i get  
\[process exited with code 1\]
## [10][Ruby Quiz - Challenge #16 - Build the Manuscripts Book Manifest for Documents in Markdown](https://www.reddit.com/r/ruby/comments/etehqg/ruby_quiz_challenge_16_build_the_manuscripts_book/)
- url: https://github.com/planetruby/quiz/tree/master/016
---

