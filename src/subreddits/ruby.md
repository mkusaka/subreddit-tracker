# ruby
## [1][Build a Twitter clone in 10 minutes with Rails, CableReady, and StimulusReflex](https://www.reddit.com/r/ruby/comments/gaaefh/build_a_twitter_clone_in_10_minutes_with_rails/)
- url: https://www.youtube.com/watch?v=F5hA79vKE_E
---

## [2][Why Can't I Manually Resize or Drag My GUI Window? (Shoes 3.3)](https://www.reddit.com/r/ruby/comments/gaulro/why_cant_i_manually_resize_or_drag_my_gui_window/)
- url: https://www.reddit.com/r/ruby/comments/gaulro/why_cant_i_manually_resize_or_drag_my_gui_window/
---
I just started learning how to code GUIs using Shoes, and I have a problem.  For some reason, I cannot drag my window or resize it manually.  I am using Windows 10, and the code is posted below:

&amp;#x200B;

&gt;\#filename = ask\_open\_file  
&gt;  
&gt;[Shoes.app](https://Shoes.app) :width =&gt; 300, :height =&gt; 165 do  
&gt;  
&gt;	\#para [File.read](https://File.read)(filename)  
&gt;  
&gt;	button "Open File to be Read"  
&gt;  
&gt;end
## [3][Railsconf 2020 Couch Edition on Tuesday 5/5](https://www.reddit.com/r/ruby/comments/gajxmi/railsconf_2020_couch_edition_on_tuesday_55/)
- url: http://railsconf.com
---

## [4][Implementation for proper multithreading programming?](https://www.reddit.com/r/ruby/comments/gaqwm5/implementation_for_proper_multithreading/)
- url: https://www.reddit.com/r/ruby/comments/gaqwm5/implementation_for_proper_multithreading/
---
Since CRuby (aka MRI) has a GIL (global interpreter lock) multithreading programming is impossible and that would not properly be multithreading, what implementation can I use?

Ik JRuby, and I have it installed because of curiosity, but it's heavy, so, what else?
## [5][Ruby Toolbox gets an API for fetching project data](https://www.reddit.com/r/ruby/comments/gag6lc/ruby_toolbox_gets_an_api_for_fetching_project_data/)
- url: https://www.ruby-toolbox.com/blog/2020-04-29/api-for-project-data
---

## [6][Undefined Local Variable Error?](https://www.reddit.com/r/ruby/comments/gan7h4/undefined_local_variable_error/)
- url: https://www.reddit.com/r/ruby/comments/gan7h4/undefined_local_variable_error/
---
Hi - I am in the middle of a coding challenge and am running up against an error of an undefined local variable or method - Here is what I am trying to do followed by my code:

\# Asking the user how many tickets they sold. There's a bulk pricing option for groups of 4 tickets. I am trying to store the $ amount (calculated based on user input) in an array. The error seems to reference ticket\_income. Essentially, I want the calculated number to be defined as ticket\_income and that value stored in my array. 

Any help would be greatly appreciated! 

&amp;#x200B;

    sales = []
    bulk = 112.00
    individ = 35.00
    
    puts "How many tickets did they purchase? Please enter as a number (ex. 1,2,3)"
      tickets = gets.chomp.to_i
      if tickets % 4 == 0
        tickets/4*bulk = ticket_income
      elsif tickets &gt; 4 &amp;&amp; tickets % 4 != 0
        tickets/4*bulk + tickets % 4*individ = ticket_income
      else tickets*individ = ticket_income
        sales.push ticket_income
      end
## [7][Deploy a new Rails app to AWS Lambda using Lamby (my notes)](https://www.reddit.com/r/ruby/comments/gaku6w/deploy_a_new_rails_app_to_aws_lambda_using_lamby/)
- url: /r/rails/comments/gafh0n/deploy_a_new_rails_app_to_aws_lambda_using_lamby/
---

## [8][Active Rails - Learn how to build production-quality Rails apps from the ground up](https://www.reddit.com/r/ruby/comments/gajstf/active_rails_learn_how_to_build_productionquality/)
- url: https://activerailsbook.com/
---

## [9][Process SEPA transactions with Ruby](https://www.reddit.com/r/ruby/comments/gabnk8/process_sepa_transactions_with_ruby/)
- url: https://www.reddit.com/r/ruby/comments/gabnk8/process_sepa_transactions_with_ruby/
---
Explanations on how to process SEPA transactions in house with Ruby and which tools may get us there faster.

&amp;#x200B;

[https://www.ramblingcode.dev/posts/process\_sepa\_with\_ruby/](https://www.ramblingcode.dev/posts/process_sepa_with_ruby/)
## [10][Looking for an online IDE for ruby](https://www.reddit.com/r/ruby/comments/ga7pr8/looking_for_an_online_ide_for_ruby/)
- url: https://www.reddit.com/r/ruby/comments/ga7pr8/looking_for_an_online_ide_for_ruby/
---
I'm a hobbyist programmer and I usually use Ruby for automating tasks, scarping, etc..

Sometimes I work on my desktop that has a dev environment setup, but sometimes I'm on my Chromebook.

I tried https://repl.it but I kept having problems. Do you guys have any recommendations?

Edit: I ended up using https://paiza.cloud/ . It let me do the things I needed, install gems, write files, etc... and it was really easy.
