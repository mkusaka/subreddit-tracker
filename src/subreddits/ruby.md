# ruby
## [1][Why is Rails boot so slow on macOS? - rubyonrails-talk](https://www.reddit.com/r/ruby/comments/fr0hog/why_is_rails_boot_so_slow_on_macos_rubyonrailstalk/)
- url: https://discuss.rubyonrails.org/t/why-is-rails-boot-so-slow-on-macos/74021
---

## [2][What are the best, customizable, open source Ruby/Rails based shopping systems? Can you recommend one?](https://www.reddit.com/r/ruby/comments/fr1c9h/what_are_the_best_customizable_open_source/)
- url: https://www.reddit.com/r/ruby/comments/fr1c9h/what_are_the_best_customizable_open_source/
---

## [3][Color in Ruby Curses not Working](https://www.reddit.com/r/ruby/comments/fqq821/color_in_ruby_curses_not_working/)
- url: https://www.reddit.com/r/ruby/comments/fqq821/color_in_ruby_curses_not_working/
---
Recently I have found [this tutorial](https://www.2n.pl/blog/basics-of-curses-library-in-ruby-make-awesome-terminal-apps) which got me interested in learning the curses gem for Ruby. I have the following program:

    require "curses"
    include Curses
    
    puts has_colors?
    
    begin
        init_screen
        start_color
        noecho
        curs_set(0)
        init_pair(1, 1, 0)
    
        loop do
            attron(color_pair(1)) {stdscr &lt;&lt; "Hello, World!"}
            attroff(color_pair(1))
            refresh
            
            c = getch()
            case c
            when "q"
                break
            end
        end
    ensure
        close_screen
    end

The `init_pair(1, 1, 0)` function should setup a foreground/background pair of colors (I believe Red/Black) and then `attron(color_pair(1))` should pick the color I had set and use it for the printed message.

I've tried googling the issue as well as looking through the [docs](https://ruby-doc.org/stdlib-2.1.0/libdoc/curses/rdoc/Curses.html) to try to find anything on the issue and I couldn't find anything. I tried using another window instead of the `stdscr` and set the color attribute that way and still got nothing.

I feel like I'm missing something incredibly simple but I just cannot find it. If anyone has used curses and would like to help I greatly appreciate it. If it matters I'm using windows and command prompt as the terminal. `has_colors?` returns true so it seems that I can use colors but I'm not sure how.

Any help would be appreciated.
## [4][How to PIMP your ruby objects with Enumerable module](https://www.reddit.com/r/ruby/comments/fqbyun/how_to_pimp_your_ruby_objects_with_enumerable/)
- url: https://medium.com/railstips/how-to-pimp-your-ruby-objects-with-enumerable-module-206f2dd5f10b
---

## [5][Adjusting RuboCop’s Defaults](https://www.reddit.com/r/ruby/comments/fpuk4q/adjusting_rubocops_defaults/)
- url: https://metaredux.com/posts/2020/03/27/adjusting-rubocop-s-defaults.html
---

## [6][Hey Guys, anyone is looking for a Senior RoR position in Amsterdam?](https://www.reddit.com/r/ruby/comments/fpxfb6/hey_guys_anyone_is_looking_for_a_senior_ror/)
- url: https://www.reddit.com/r/ruby/comments/fpxfb6/hey_guys_anyone_is_looking_for_a_senior_ror/
---
For a domain marketplace company! Salary 50-80k
## [7][Ruby container types other than stdlib?](https://www.reddit.com/r/ruby/comments/fpwq6j/ruby_container_types_other_than_stdlib/)
- url: https://www.reddit.com/r/ruby/comments/fpwq6j/ruby_container_types_other_than_stdlib/
---
I've been digging into Rust for an embedded project and found that there is a lot more container types available.  I see implementations for SwissTable and other containers, lots to choice from based on if you want constant-time lookups vs quick insertion.  I'm having a hard time finding different implementations of containers in Ruby. Do they exist?  Is it even possible?  I'd love to be able to pickup a Gem that implements different container types and choice the correct implementation based on my needs.
## [8][Ruby for programmers(video tutorial)](https://www.reddit.com/r/ruby/comments/fpla5v/ruby_for_programmersvideo_tutorial/)
- url: https://www.reddit.com/r/ruby/comments/fpla5v/ruby_for_programmersvideo_tutorial/
---
Hey guys, i created a single-video tutorial aiming for programmers to learn Ruby. 

This tutorial goes through implementing an actual program (sorting numbers with quicksort).

It would be nice to get some feedbacks, i have plans to create similar videos for other languages.

https://youtu.be/WYSHbSzyonk
## [9][cURL wrapper around RubyGems.org API](https://www.reddit.com/r/ruby/comments/fpdugp/curl_wrapper_around_rubygemsorg_api/)
- url: https://github.com/davidesantangelo/curl-gems
---

## [10][Webpacker 5.0 released](https://www.reddit.com/r/ruby/comments/fp6r12/webpacker_50_released/)
- url: https://prathamesh.tech/2020/03/25/webpacker-5-0-released/
---

