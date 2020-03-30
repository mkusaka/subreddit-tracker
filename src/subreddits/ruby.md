# ruby
## [1][The first version of dry-rails was released!](https://www.reddit.com/r/ruby/comments/frofpd/the_first_version_of_dryrails_was_released/)
- url: https://www.reddit.com/r/ruby/comments/frofpd/the_first_version_of_dryrails_was_released/
---
If you're using Rails and you'd like to try out [dry-rb](https://dry-rb.org/) then I have great news for you - the first version of dry-rails was released today and it provides built-in support for dry-system that works OOTB with both Rails 5 and 6 (works with Zeitwerk enabled too). It also gives you additional features that are integrated with the controller layer:

* `safe_params` - an alternative to strong-parameters, powered by dry-schema
* `application_contract` - an alternative to ActiveModel/Record validations powered by dry-validation
* `controller_helpers` - a couple of convenient methods that make it easier to access your application container

## Links

* [0.1.0 release on GitHub](https://github.com/dry-rb/dry-rails/releases/tag/v0.1.0) and feel
* [API docs](https://rubydoc.info/gems/dry-rails)
* [User docs](https://dry-rb.org/gems/dry-rails)

Check it out, provide feedback and feel free to ask any questions!
## [2][Understanding The Ruby Method Lookup Chain](https://www.reddit.com/r/ruby/comments/frmbz5/understanding_the_ruby_method_lookup_chain/)
- url: https://medium.com/better-programming/understanding-the-ruby-method-lookup-chain-d6f9d7997849
---

## [3][How to avoid procrastination NOW!](https://www.reddit.com/r/ruby/comments/frm84k/how_to_avoid_procrastination_now/)
- url: https://www.reddit.com/r/ruby/comments/frm84k/how_to_avoid_procrastination_now/
---
Now, we have plenty of time due to Corona. We have time to learn that new programming language, that new tool, to work on that project. However, we tend to procrastinate.

As a software engineer, I have been there myself. Netflix, videogames, YouTube seem to be more interesting. But now that you have the time, work on yourself, become an expert in your field.

These are the quick tips: 

- Use a timer technique (such as Pomodoro) 

- Write down the small tasks you will do in a piece of paper 

- Work in each small task during the day 

- Do not multitask 

- Leave your phone in another room 

- After working, give yourself a break watching Netflix or videogames 

Comment below if you are also procrastinating and not getting things done these days.
## [4][Endpoint_flux is another gem to deal with endpoints](https://www.reddit.com/r/ruby/comments/frbwj5/endpoint_flux_is_another_gem_to_deal_with/)
- url: https://www.reddit.com/r/ruby/comments/frbwj5/endpoint_flux_is_another_gem_to_deal_with/
---
Hi! Just wanted to share my first open source project. It's forked from original endpoint_flux and I like it so much that decided to contribute myself. Currently I'm looking for any feedback anyone can leave about the project. And ofc you're welcome to ask any questions you have. 

Short descriptions and ideas:
1) Organize endpoints in a better way, split business logic into different layers
2) Run everywhere including non Rails projects

https://github.com/pavelkvasnikov/endpoint-flux
## [5][Starter kit for test automation engineers (Ruby stack)](https://www.reddit.com/r/ruby/comments/fro7vq/starter_kit_for_test_automation_engineers_ruby/)
- url: https://gitlab.com/Rukomoynikov/yesterday-e2e
---

## [6][Auto Rubocop](https://www.reddit.com/r/ruby/comments/frn3xi/auto_rubocop/)
- url: https://www.reddit.com/r/ruby/comments/frn3xi/auto_rubocop/
---
Automatically applies rubocop safe auto correct to modified ruby files in your project, thus increases your code Karma. [https://gitlab.com/mindaslab/auto-rubocop](https://gitlab.com/mindaslab/auto-rubocop)
## [7][Why is Rails boot so slow on macOS? - rubyonrails-talk](https://www.reddit.com/r/ruby/comments/fr0hog/why_is_rails_boot_so_slow_on_macos_rubyonrailstalk/)
- url: https://discuss.rubyonrails.org/t/why-is-rails-boot-so-slow-on-macos/74021
---

## [8][What are the best, customizable, open source Ruby/Rails based shopping systems? Can you recommend one?](https://www.reddit.com/r/ruby/comments/fr1c9h/what_are_the_best_customizable_open_source/)
- url: https://www.reddit.com/r/ruby/comments/fr1c9h/what_are_the_best_customizable_open_source/
---

## [9][Color in Ruby Curses not Working](https://www.reddit.com/r/ruby/comments/fqq821/color_in_ruby_curses_not_working/)
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
## [10][How to PIMP your ruby objects with Enumerable module](https://www.reddit.com/r/ruby/comments/fqbyun/how_to_pimp_your_ruby_objects_with_enumerable/)
- url: https://medium.com/railstips/how-to-pimp-your-ruby-objects-with-enumerable-module-206f2dd5f10b
---

