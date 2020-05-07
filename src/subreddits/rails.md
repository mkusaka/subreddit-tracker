# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/g616hm/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/g616hm/personal_projects_show_off_your_own_project_andor/
---
In this thread you can showcase your personal pet project to other redditors.

Need help with a specific problem or just wanna have some extra eyeballs on your code? Ask away!

A suggested format to get you started:

1. **Name of your project**
2. **A short description**
3. **Application stack**
4. **Link to Live app**
5. **Link to GitHub**
6. **You experience level**
7. **Other information or areas that you would like advice on**

 

^(Many thanks to Kritnc for getting the ball rolling.)
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/
---
In this thread you can showcase your personal pet project to other redditors.

Need help with a specific problem or just wanna have some extra eyeballs on your code? Ask away!

A suggested format to get you started:

1. **Name of your project**
2. **A short description**
3. **Application stack**
4. **Link to Live app**
5. **Link to GitHub**
6. **You experience level**
7. **Other information or areas that you would like advice on**

 

^(Many thanks to Kritnc for getting the ball rolling.)
## [3][Notes from DHH's RailsConf keynote interview](https://www.reddit.com/r/rails/comments/geyr61/notes_from_dhhs_railsconf_keynote_interview/)
- url: https://www.reddit.com/r/rails/comments/geyr61/notes_from_dhhs_railsconf_keynote_interview/
---
For those of you who haven't had the chance to watch yet:

https://www.joshuawood.net/notes/2020-railsconf-dhh-keynote

One of the parts I liked best was this comment on repeating history as an industry:

&gt; The amnesia is partially caused by so many new people entering the industry. They haven’t experienced it any other way.

There are a ton of new people who haven’t been exposed to the joy of Rails because JavaScript is mainstream and/or someone told them not to bother.

The “Rails Way” is completely foreign to many of the new JavaScript developers I’ve met. Don’t assume people get it just because it’s been around forever. Spread the word:

&gt; Ruby is great because you can know a little JavaScript and then jump to Ruby, understand it, and own the full stack.

What do y'all think about the state of Ruby and Rails in 2020?
## [4][Please help. Running out of hairs to pull.](https://www.reddit.com/r/rails/comments/geza8v/please_help_running_out_of_hairs_to_pull/)
- url: https://www.reddit.com/r/rails/comments/geza8v/please_help_running_out_of_hairs_to_pull/
---
This is likely more of an OS issue than it is a dev/coding  issue, but I trust you guys' judgement. I am constantly having to reboot my terminal (running on zsh since apple forced it) to run the devise:install and all other gem variants. 

Here's what happens: 

I do something, whatever on the terminal. Then, I run one of the install operators. It works fine. Next time I try it - just finished running the simple\_forms install and now doing devise views - it just freezes up and does nothing. Opening a new tab and running it there doesn't work. Need to cmd + Q out of that b\*\*\*\* and reopen the terimnal, then cd back into the app, then it works fine. I've done this now 6 times in the last 45 minutes. I will be bald by the end of the night if I can't fix this nuance.
## [5][Where should decorator methods that don't belong to a model go?](https://www.reddit.com/r/rails/comments/gf5r1v/where_should_decorator_methods_that_dont_belong/)
- url: https://www.reddit.com/r/rails/comments/gf5r1v/where_should_decorator_methods_that_dont_belong/
---
Hi all,

I've got pretty familiar with using the draper gem now to store decorator methods for my models in order to clean up my views.  


But now I have the following method in my home controller:

    request.location.city ? @location = request.location.city : @location = "..."

And this is how I'm calling it in my home page:

    placeholder: "Search for results in #{@location}",

The controller seems like the wrong place to put logic purely related to the view.  


I can't see anything in the draper documentation about decorator methods for views. It's all about decorator methods for models.  


So I'm wondering, where should I be putting a method like the one just mentioned?  


Thanks
## [6][Is Everyday Rails Testing with RSpec a Good Read for a Mid-Level Developer?](https://www.reddit.com/r/rails/comments/gf4gqo/is_everyday_rails_testing_with_rspec_a_good_read/)
- url: https://www.reddit.com/r/rails/comments/gf4gqo/is_everyday_rails_testing_with_rspec_a_good_read/
---
As it says on the tin, I'd be keen to hear from developers with, say, 2+ years of experience on whether they would recommend the book?

I appreciate there will be many elements that are for juniors but if it also caters beyond that then I'll be purchasing a copy.

Thanks!
## [7][Should I use multi-tenancy in my project ?](https://www.reddit.com/r/rails/comments/geo7h6/should_i_use_multitenancy_in_my_project/)
- url: https://www.reddit.com/r/rails/comments/geo7h6/should_i_use_multitenancy_in_my_project/
---
So, I was taking this course in which  you code a SaaS Management Project. This app uses milia gem for multi tenenacy. There is two types of user. Company and Employee. A company can create a project and then invite employees to be part of this project. And in the projects you can upload files, texts., etc...

Unfortunately I couldn’t get through this course's section due the very outdated content such as milia gem and rails 4.

I was looking forward to code this app because I could imagine my undergraduate thesis project being built using the very bases of this SaaS management project app. 

I’m trying to code a simple video sharing plataform where a professor could have their own private “YouTube channel” where their could invite their students to participate and watch the lectures. 

So Insted of Employee and Company, my models would be Professors and Students.

So, I’m looking for this “privacy” characteristic. That you only would be able to see the content if you were invited or something like that. 

The multi tenancy for my project is the correct approach?
## [8][How do you allow users to select between multiple currencies?](https://www.reddit.com/r/rails/comments/gepybq/how_do_you_allow_users_to_select_between_multiple/)
- url: https://www.reddit.com/r/rails/comments/gepybq/how_do_you_allow_users_to_select_between_multiple/
---
We have an app with ~5 languages but we would like to support more currencies.

So someone in the Netherlands could choose the Euro as the currency and have the app in English etc.

No need to convert between currencies, basically just the symbol + the correct formatting.

We are looking for a way to add multiple languages into a `select` field, and still be able to use formatting.

This is how we control which languages are available

```ruby
# config/application.rb
config.i18n.available_locales = [:en, :de, :is, :fr, :es, :'en-GB']
```

`&lt;%= number_to_currency(123456, locale: 'en') %&gt;`

`&lt;%= Money.from_amount(123456, 'USD').format %&gt;`

I thought this was a trivial thing but searching for hours gave me nothing concrete, I must be missing something, therefor I am asking here.

Do we need gems like Money or `currency_select` for this?
## [9][New to Rails and MVC, looking for understanding about where certain logic should live?](https://www.reddit.com/r/rails/comments/gep01i/new_to_rails_and_mvc_looking_for_understanding/)
- url: https://www.reddit.com/r/rails/comments/gep01i/new_to_rails_and_mvc_looking_for_understanding/
---
I'm a newb dev...most of the work I've done has been in PHP (I used to work at a WordPress company).  


When I tried to create my first app in Rails, it was incredibly confusing.  I ended up using Sinatra instead at the suggestion of a colleague since it's more forgiving of bad practices, and that helped me understand how controllers, routes, and views work (I don't have a DB so I've not used models yet).  


Now I'm trying in Rails again and even though I know roughly how things are working now, I still don't really understand best practices.  What/How much logic should live in a view vs a model?  


I guess I understand roughly what Rails is doing technically, but not best practices/theory.  Where might I go to better understand what should live where?
## [10][RSpec and Minitest](https://www.reddit.com/r/rails/comments/gedoal/rspec_and_minitest/)
- url: https://www.reddit.com/r/rails/comments/gedoal/rspec_and_minitest/
---
Do you create RSpec and Minitest in your app? or you choose only one?
## [11][Database design for rails application](https://www.reddit.com/r/rails/comments/gefv84/database_design_for_rails_application/)
- url: https://www.reddit.com/r/rails/comments/gefv84/database_design_for_rails_application/
---
This is my [database design](https://dbdiagram.io/d/5ea1b31e39d18f5553fe19e6). I am not sure if my associations are correct. Should I use composite keys or foreign keys? Can you guys tell me what is wrong with my database design. Thankss
## [12][Creating model to controller to view or create another model and make associations](https://www.reddit.com/r/rails/comments/geatl7/creating_model_to_controller_to_view_or_create/)
- url: https://www.reddit.com/r/rails/comments/geatl7/creating_model_to_controller_to_view_or_create/
---
Hello there, I am done creating my unit testing for my first created model. What to do next? Do I need to create the next model and associations of each or I will continue build my model until controller and view?
