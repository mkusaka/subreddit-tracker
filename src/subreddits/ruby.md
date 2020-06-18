# ruby
## [1][An attempt at creating a "Heroku-like" Docker/Helm/GitHub Action setup for "classic-stack" Rails apps. Feedback wanted.](https://www.reddit.com/r/ruby/comments/haz3c0/an_attempt_at_creating_a_herokulike/)
- url: https://github.com/lewagon/rails-k8s-demo
---

## [2][I wanna make fancy(!) TUIs for Linux using Ruby. Are my only options the ancient and largely undocumented ncurses wrappers?](https://www.reddit.com/r/ruby/comments/hb4uqq/i_wanna_make_fancy_tuis_for_linux_using_ruby_are/)
- url: https://www.reddit.com/r/ruby/comments/hb4uqq/i_wanna_make_fancy_tuis_for_linux_using_ruby_are/
---
Every time I've tried to use one of the many ncurses wrappers, there's like one blog post and if you're lucky some docs written for C programmers.

I've tried reading the 'sup' source and it wasn't really helpful. Is the windowed TUI the stuff of wizards in 2020?
## [3][How to use a string array for case labels and call methods with same name](https://www.reddit.com/r/ruby/comments/hbc4uo/how_to_use_a_string_array_for_case_labels_and/)
- url: https://www.reddit.com/r/ruby/comments/hbc4uo/how_to_use_a_string_array_for_case_labels_and/
---
Hi All -

Can't quite figure this one out, hoping the /r/ruby brains trust can help me out.

I have code similar to this:

    cmd = get\_string\_from\_requestor
    case cmd
    when 'FOO' then foo # foo() is a method on the same class
    when 'BAR' then bar # as is bar()
    when 'BAZ' then bar # etc
    else
        # error
    end

And I was hoping to switch it to something a bit more dynamic:

    cmd = get\_string\_from\_requestor
    commands = %w(FOO BAR BAZ QUX QUUX)
    case cmd
    # magic required here: iterate(?) through commands[], match cmd
    # to the commands[i] and then call commands[i].downcase as a
    # method on current class...?
    else
        # error
    end

Anyone point me in the right direction?

Thanks!
## [4][Is there a full list of structs, unions and functions of the C API?](https://www.reddit.com/r/ruby/comments/hbdjd8/is_there_a_full_list_of_structs_unions_and/)
- url: https://www.reddit.com/r/ruby/comments/hbdjd8/is_there_a_full_list_of_structs_unions_and/
---
I am looking for a complete list of the \`struct\`s, \`typedef struct\`s, and functions in Ruby C API and it's body.  


I tried looking up but I found nothing.
## [5][Rails 6.1's ActiveModel Errors Revamp](https://www.reddit.com/r/ruby/comments/hap4u9/rails_61s_activemodel_errors_revamp/)
- url: https://code.lulalala.com/2020/0531-1013.html
---

## [6][Best way to parse shell style (VAR=value) configuration files](https://www.reddit.com/r/ruby/comments/hawtu4/best_way_to_parse_shell_style_varvalue/)
- url: https://www.reddit.com/r/ruby/comments/hawtu4/best_way_to_parse_shell_style_varvalue/
---
I have a string `GRUB_CMDLINE_LINUX="quiet rhgb crashkernel=auto"` which is pretty much shell style variable assignment. I need to process this string so I can extract string within the quotes  `quiet rhgb crashkernel=auto` and then split the options into an array.

I do it this way:

    line = 'GRUB_CMDLINE_LINUX="quiet rhgb crashkernel=auto"'
    kernel_opts = line.match(/"(.*)"/).to_s.delete('"').split(' ')

Is there a leaner way to do it?

Is there a way in Ruby to load shell style configuration files naively (without processing them as strings)?
## [7][Using Service Objects in Ruby on Rails](https://www.reddit.com/r/ruby/comments/haqv41/using_service_objects_in_ruby_on_rails/)
- url: https://blog.appsignal.com/2020/06/17/using-service-objects-in-ruby-on-rails.html
---

## [8][Interview with DHH: the Ruby on Rails creator, co-founder &amp; CTO of Basecamp and Le Mans class-winning race car driver](https://www.reddit.com/r/ruby/comments/hami35/interview_with_dhh_the_ruby_on_rails_creator/)
- url: https://evrone.com/dhh-interview
---

## [9][Encapsulate each validation error as an Error object](https://www.reddit.com/r/ruby/comments/hasv2r/encapsulate_each_validation_error_as_an_error/)
- url: https://blog.saeloun.com/2020/06/17/rails-active-model-errors
---

## [10][Ruby on rails cheatsheet from Michael Hartl tutorials](https://www.reddit.com/r/ruby/comments/ham5u8/ruby_on_rails_cheatsheet_from_michael_hartl/)
- url: https://www.reddit.com/r/ruby/comments/ham5u8/ruby_on_rails_cheatsheet_from_michael_hartl/
---
Hi folks,   

Hope everyone is staying safe!

I started learning Ruby-on-Rails framework two weeks before &amp; made a cheat sheet with minimal explanations of different concepts, referred from **Michael Hartl**’s "[Learn Web Development with Rails - Fourth Edition](https://www.pdfdrive.com/ruby-on-rails-tutorial-learn-web-development-with-rails-4th-edition-e184254589.html)" tutorials.  

The textbook is really awesome &amp; helped me to build a simple blogging website similar to the basic version of Twitter and the source code is available in github: [https://github.com/ddlogesh/rails-tutorial](https://github.com/ddlogesh/rails-tutorial) 

I would like to share these two cheat sheets, which may be a good start for preparing ruby-on-rails framework or recollect the concepts learned from the above-mentioned textbook.  

Rails: [https://www.notion.so/Ruby-on-Rails-Cheat-Sheet-61106a73031d46ec81c285daaebf1409](https://www.notion.so/Ruby-on-Rails-Cheat-Sheet-61106a73031d46ec81c285daaebf1409)

Ruby: [https://www.notion.so/Ruby-Cheat-Sheet-7c8aabc9268b4e65b5691245dd19068c](https://www.notion.so/Ruby-Cheat-Sheet-7c8aabc9268b4e65b5691245dd19068c)

If you are a beginner to Ruby, please watch this 4-hours tutorial from **freeCodeCamp**: [https://www.youtube.com/watch?v=t\_ispmWmdjY](https://www.youtube.com/watch?v=t_ispmWmdjY). Though it's quite a long, it's worth watching it!

Please do share these cheatsheets with your Ruby-on-Rails enthusiasts, who may be interested to learn Ruby-on-Rails.
