# ruby
## [1][How Heroku Sped Up Time-Related Syscalls on Dynos](https://www.reddit.com/r/ruby/comments/hsv9bc/how_heroku_sped_up_timerelated_syscalls_on_dynos/)
- url: https://blog.heroku.com/clocksource-tuning
---

## [2][How to use a Transaction Script(aka Service Objects) in Ruby on Rails. Simple example](https://www.reddit.com/r/ruby/comments/hssfqu/how_to_use_a_transaction_scriptaka_service/)
- url: https://www.reddit.com/r/ruby/comments/hssfqu/how_to_use_a_transaction_scriptaka_service/
---
The logic of small applications can be present as a series of transactions. Using the Transaction Scripts pattern, we get an application that is easier to maintain, to cover with tests, and to scale.

In the [tutorial](https://jtway.co/how-to-use-a-transaction-script-aka-service-objects-in-ruby-on-rails-simple-example-161b7e228942) we will develop an [application](https://github.com/dgorodnichy/transaction-script-example) that has Post, User, and Like models. Users should be able to like posts. The first version of the controller will contain extra code, which we will extract into a separate Transaction Script. We also describe when we need to use the Transaction Scripts and the pros of the transaction script usage.

Full tutorial: [How to use a Transaction Script (aka Service Objects) in Ruby on Rails. Simple example](https://jtway.co/how-to-use-a-transaction-script-aka-service-objects-in-ruby-on-rails-simple-example-161b7e228942)
## [3][First gem, interested in feedback](https://www.reddit.com/r/ruby/comments/hsbta7/first_gem_interested_in_feedback/)
- url: https://www.reddit.com/r/ruby/comments/hsbta7/first_gem_interested_in_feedback/
---
Gem name: Backpedal

Repo: [https://github.com/Greyoxide/backpedal](https://github.com/Greyoxide/backpedal)

RubyGems: [https://rubygems.org/gems/backpedal/versions/0.1.2](https://rubygems.org/gems/backpedal/versions/0.1.2)

The goal with this gem is to provide developers with an easy way to move users back on their respective navigational path.  I found myself arriving at views from different angles and I wanted to develop a utilitarian method for moving users back along their respective navigational path.

For example: /orders/22 -&gt; /products/141 In this case the back\_link on the product view would take the user back to /orders/22.  Using this same example: /product\_lines/41 -&gt; /products/141 This controller action renders the same template as before, however the back button would take the user back to /product\_lines/41

This is my first gem, and I'd love some feedback.
## [4][Question regarding using end in flow control blocks](https://www.reddit.com/r/ruby/comments/hsgyf3/question_regarding_using_end_in_flow_control/)
- url: https://www.reddit.com/r/ruby/comments/hsgyf3/question_regarding_using_end_in_flow_control/
---
Hello everyone.

I have recently started learning ruby and currently I am trying to write the Tic Tac Toe game.

However, I am really confused by the usage of end in Ruby. I have this method and I keep getting errors saying expecting end.

        def play
          has_won = false
          turn = 1
          winner = nil
          until !has_won &amp;&amp; turn &lt; 10 do
            puts print_board
            if turn.odd?
              selection = 0
              until valid_move?(selection) do
                puts "#{players[0].name} select where you want to put your symbol (O): "
                selection = gets.chomp.to_i
                if valid_move?(selection)
                  @board[selection] = player[0].symbol
                  has_won = won?(player[0])
                  winner = player[0] if has_won
                else
                  puts "Wrong selection. Try again."
                end
              end
            else 
              selection = 0
              until valid_move?(selection) do
                puts "#{players[1].name} select where you want to put your symbol (O): "
                selection = gets.chomp.to_i
                if valid_move?(selection)
                  @board[selection] = player[1].symbol
                  has_won = won?(player[1])
                  winner = player[1] if has_won
                else
                  puts "Wrong selection. Try again."
                end
              end
            end
    
            turn++
          end
          
    
          if has_won ?
            puts("#{winner.name} has won the game. Congratulations! "):
            puts("The game ended in a draw!")
            
        end

My question is when exactly do I need to use end? What is wrong in this method? Do I need to use end in such cases as `winner = player[1] if has_won`
## [5][RubyMine vs IntelliJ + RubyMine plugin](https://www.reddit.com/r/ruby/comments/hsbcg5/rubymine_vs_intellij_rubymine_plugin/)
- url: https://www.reddit.com/r/ruby/comments/hsbcg5/rubymine_vs_intellij_rubymine_plugin/
---
Is there any difference between RubyMine and IntelliJ + RubyMine plugin? I'd rather use one IDE (IntelliJ) for everything rather than separate ones for each language.
## [6][Defining gem dependencies in single-file Ruby scripts](https://www.reddit.com/r/ruby/comments/hrzy5d/defining_gem_dependencies_in_singlefile_ruby/)
- url: https://nithinbekal.com/posts/bundler-inline-gemfile/
---

## [7][Nexss Programmer - Open Source tool for the programmers - Ruby included](https://www.reddit.com/r/ruby/comments/hsi0rj/nexss_programmer_open_source_tool_for_the/)
- url: https://www.reddit.com/r/ruby/comments/hsi0rj/nexss_programmer_open_source_tool_for_the/
---
Hi guys,

We have made Open Source Tool for the Programmers - Nexss Programmer. 

50 different programming languages together... 

Ruby is included so you can combine it with other languages!

If you guys want to check it out here is the link for the video presentation: [https://www.youtube.com/watch?v=7WbnYyEnBNk](https://www.youtube.com/watch?v=7WbnYyEnBNk)

It would be great if we can get more feedback on it.

Have a nice day!
## [8][Ruby VS Python](https://www.reddit.com/r/ruby/comments/hs2nsk/ruby_vs_python/)
- url: https://www.reddit.com/r/ruby/comments/hs2nsk/ruby_vs_python/
---
More than a versus (it's more striking) I would really like to know what things they have in common and how they differ, what their main advantages are compared to the other and of course their opinion of them. One point to take into account is because they believe that python is more used in AI and Machine learning developments.
## [9][System of a test: Proper browser testing in Ruby on Rails](https://www.reddit.com/r/ruby/comments/hs54c6/system_of_a_test_proper_browser_testing_in_ruby/)
- url: https://evilmartians.com/chronicles/system-of-a-test-setting-up-end-to-end-rails-testing
---

## [10][Patching Ruby? How to handle mixing Module#prepend with method aliasing](https://www.reddit.com/r/ruby/comments/hs20zh/patching_ruby_how_to_handle_mixing_moduleprepend/)
- url: https://www.mayerdan.com/ruby/2020/07/15/ruby-gems-patching-std-lib
---

