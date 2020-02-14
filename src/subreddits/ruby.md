# ruby
## [1][Ruby Memory Model](https://www.reddit.com/r/ruby/comments/f3r46i/ruby_memory_model/)
- url: https://www.reddit.com/r/ruby/comments/f3r46i/ruby_memory_model/
---
I found it interesting and good documentation on the ruby memory model, by concurrent ruby.  
Sharing for those who forgot to visit.  
[https://docs.google.com/document/d/1pVzU8w\_QF44YzUCCab990Q\_WZOdhpKolCIHaiXG-sPw/](https://docs.google.com/document/d/1pVzU8w_QF44YzUCCab990Q_WZOdhpKolCIHaiXG-sPw/edit#)
## [2][New ruby user here with questions about loops](https://www.reddit.com/r/ruby/comments/f3o6l7/new_ruby_user_here_with_questions_about_loops/)
- url: https://www.reddit.com/r/ruby/comments/f3o6l7/new_ruby_user_here_with_questions_about_loops/
---
I was wondering if someone can tutor me a bit on loops in ruby through private messages.
## [3][Can't open .rb file on IRB using the ruby command. Getting errors!](https://www.reddit.com/r/ruby/comments/f3o5lx/cant_open_rb_file_on_irb_using_the_ruby_command/)
- url: https://www.reddit.com/r/ruby/comments/f3o5lx/cant_open_rb_file_on_irb_using_the_ruby_command/
---
I've been trying to open files on IRB by using the ruby command (entering something like ruby text.rb to the terminal) but it only gives me errors. I'm trying it with a really basic new file (just one puts "hello") to be sure there are no typos, errors, or hidden characters, but it still doesn't work. I get this error:

ruby text.rb

SyntaxError: (irb):2: syntax error, unexpected tIDENTIFIER, expecting keyword\_do or '{' or '('

ruby text.rb

\^

from /usr/bin/irb:11:in \`&lt;main&gt;'

I made sure to cd into the location where the .rb file is at. No code errors in the file. I have no idea what's going on. It used to work fine months ago, but not anymore. I can load files using the load command, but I'm following some courses that require me to open files using the ruby command to be able to enter command line arguments. 

Does anyone know what's going on?
## [4][Rails adds ActiveRecord API for switching multiple database connections](https://www.reddit.com/r/ruby/comments/f3s13y/rails_adds_activerecord_api_for_switching/)
- url: https://blog.saeloun.com/2020/02/14/rails-6-multiple-database-support
---

## [5][Sandi Metz Talk in Toronto - Sat Feb 22/2020](https://www.reddit.com/r/ruby/comments/f3c2x0/sandi_metz_talk_in_toronto_sat_feb_222020/)
- url: https://www.eventbrite.com/e/cover-presents-an-evening-with-sandi-metz-registration-94178429217
---

## [6][Sinatra Web App Parameter not being passed intermittently](https://www.reddit.com/r/ruby/comments/f3mim2/sinatra_web_app_parameter_not_being_passed/)
- url: https://www.reddit.com/r/ruby/comments/f3mim2/sinatra_web_app_parameter_not_being_passed/
---
 I am creating my first application using the Sinatra framework and I am running into an issue where my content parameter displays its contents intermittently. My website is a simple movie reviews model where users will sign up or login create a review by inputting a title, and then the content for the review. Currently I see the title and author of the review being passed in without issue but the content appears empty. It initially passed in fine but I am not sure whats changed. Below is the review controller method the view for the page and the output when it did work and now when it hasn't. 

 

`post '/review_entries' do`         

`redirect_if_not_logged_in`         

`if params[:content] != ""` 

  `@review_info = Review.create(content: params[:content], user_id: current_user.id, title: params[:title])` 

  `#binding.pry`             

  `redirect "review_entries/#{@review_info.id}"` 

`else`             

  `redirect 'review_entries/new'` 

 `end` 

`end`

My view is below

`VIEWS:`

`&lt;h3&gt;&lt;%= @review_info.title %&gt;&lt;/h3&gt;`

&amp;#x200B;

`&lt;p&gt;&lt;%= @review_info.content %&gt;&lt;/p&gt;`

&amp;#x200B;

&amp;#x200B;

`&lt;p&gt;Created by: &lt;%= @review_info.user.username %&gt;&lt;/p&gt;`

&amp;#x200B;

&amp;#x200B;

[The bold is the title the text below is the content value it only worked for 2 of the 6 reviews.](https://preview.redd.it/k1qjoou2ctg41.png?width=345&amp;format=png&amp;auto=webp&amp;s=045dbd28a8e8cf2d19851b227472f7b380e494e2)
## [7][content security policy using sinatra and heroku](https://www.reddit.com/r/ruby/comments/f3dt0x/content_security_policy_using_sinatra_and_heroku/)
- url: https://www.reddit.com/r/ruby/comments/f3dt0x/content_security_policy_using_sinatra_and_heroku/
---
i've asked this question on stack overflow yesterday so it's just as easy to link to: 

https://stackoverflow.com/questions/60197579/heroku-app-on-firefox-wont-load-javascripts-due-to-content-security-policy

There is one answer but it relates to rails - I thought I could find something for sinatra/rack, but have been unable to. I tried [this gem](https://github.com/grempe/rack-content_security_policy) but no matter what i did i got errors - and it's been archived for some time so generally feels like not the way to go. 

tbh, i've never had this problem before, and I'm not doing anything fancy. I've used sinatra many times, it's been sometime since I've used heroku for an app with jquery in it, but...still, it's pretty basic stuff. 

Anyone else encountered this and have a suggestion?
## [8][Installing MSYS2 Packages with Gem](https://www.reddit.com/r/ruby/comments/f3eb72/installing_msys2_packages_with_gem/)
- url: https://www.reddit.com/r/ruby/comments/f3eb72/installing_msys2_packages_with_gem/
---
I don't have a Windows machine or VM to test with at the moment, but I have a two part question regarding installing MSYS2 packages as gem dependencies.

As far as I understand it, the MSYS2 installation that comes with a typical Ruby installation on Windows for building C extensions, and is quite capable of also installing packages from their own repos via `pacman`. I am currently working on a C extension that has a dependency that might be problematic distributing on Windows, and even worse to try and automate the build process. It does have a PKGBUILD for MSYS2 that would allow for easily installation, so...

1. I am not misunderstanding anything, this is entirely possible, and I can then link with the binary compiled by MINGW within the MSYS2 environment?
2. If possible, is this a bad practice, as I have not been able to find a single example of this being done.
## [9][Simple ruby error I can not figure out.](https://www.reddit.com/r/ruby/comments/f3a7af/simple_ruby_error_i_can_not_figure_out/)
- url: /r/learnprogramming/comments/f39z2z/simple_ruby_error_i_can_not_figure_out/
---

## [10][Understanding Rails secrets/credentials](https://www.reddit.com/r/ruby/comments/f33dby/understanding_rails_secretscredentials/)
- url: https://www.codewithjason.com/understanding-rails-secrets-credentials/
---

