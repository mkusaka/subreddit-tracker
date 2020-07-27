# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/hwehh6/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/hwehh6/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Cant access rails web console on Macos Catalina](https://www.reddit.com/r/rails/comments/hypwj0/cant_access_rails_web_console_on_macos_catalina/)
- url: https://www.reddit.com/r/rails/comments/hypwj0/cant_access_rails_web_console_on_macos_catalina/
---
Yesterday i just realized i cant access rails' web console and even better\_errors gem's exception page on my mac.

But the point is everything s cool on my linux machine. Any ideas why can it happen?And its also not only on my mac, my co-workers also cant see the live shell on their macs.

It's the same code, same browser, same exception but different exception pages.

Macos;

[There's not a live shell on macos.](https://preview.redd.it/h5k3w18xndd51.png?width=1072&amp;format=png&amp;auto=webp&amp;s=5a49937d43e6355bdf7eb384b0a7032f653493a2)

Linux;

[better\_errors gem is working and there's a live shell on linux.  ](https://preview.redd.it/bg9ojc8rndd51.png?width=1237&amp;format=png&amp;auto=webp&amp;s=2fba51ce926b6626c88b3fa892f99393b53ffb07)

&amp;#x200B;
## [3][The more I look take a look at different apps, I see that they're just CRUD applications.](https://www.reddit.com/r/rails/comments/hybezw/the_more_i_look_take_a_look_at_different_apps_i/)
- url: https://www.reddit.com/r/rails/comments/hybezw/the_more_i_look_take_a_look_at_different_apps_i/
---
I'm not sure if I'm thinking about this correctly, but I'm starting to see that the applications that I'm looking at are all perform the same basic functions.

To elaborate, I see that most apps read and display data from a database. For example, an exercise web app shows users the exercise name, target muscles, and a video demonstration. Another example, a restaurant booking app, the app looks at how many time slots are not reserved and shows the data to the person checking for available times. Another example, Mint the personal finance app. It connects with your bank and it shows the user spending and alerts.

There's a lot of times when I visited a website and said to myself "Oh this app is just here to organize and regurgitate my data."

The question I have is, is this considered "business logic", or is there another layer to developing applications?
## [4][Webpack Environment Config File](https://www.reddit.com/r/rails/comments/hygoa4/webpack_environment_config_file/)
- url: https://www.reddit.com/r/rails/comments/hygoa4/webpack_environment_config_file/
---
Hi folks,

In my Rails 6 project, I have the following `config/webpack/environment.js`:

```javascript
const { environment } = require('@rails/webpacker')
const webpack = require('webpack')

environment.plugins.prepend('Provide',
  new webpack.ProvidePlugin({
    $: 'jquery/src/jquery',
    jQuery: 'jquery/src/jquery',
    Popper: ['popper.js', 'default'],
  })
)

module.exports = environment
```

My understanding is that this configuration is setting global variables `$`, `jQuery`, and `Popper`. Any other JavaScript that I write can use these variables, as they are globally available. 

I have confirmed that this is the case for `$` and for `jQuery`, but when I try to reference `Popper` in the browser's console, I get an uncaught reference error. What am I missing here?
## [5][Dockerizing a super old rails app, looking for advice](https://www.reddit.com/r/rails/comments/hyd2q1/dockerizing_a_super_old_rails_app_looking_for/)
- url: https://www.reddit.com/r/rails/comments/hyd2q1/dockerizing_a_super_old_rails_app_looking_for/
---
Have any of you had to dockerize an old rails app at some point? I'm talking Ruby on Rails 2.3, Ruby 1.8.6, MySQL 2.7. That's actually what I'm trying to do but I'm having a lot of trouble.

I don't think the official ruby image on Docker Hub even goes back far enough to 1.8, and all of the Dockerfiles I'm finding on Github don't work, or they use SQLite instead of MySQL.

Any advice or resources are super appreciated

## Update
So I think I finally got it! However there is one small problem, and I cannot find the solution online.

How can I import a database dump (.sql file) into a rails database? In rails 4+, you could just use `rails db &lt; ./path/to/db_dump.sql`, but I have no idea how to do that in rails 2.3 haha. I don't think i'm able to import the sql file using the mysql cli since the mysql container and app container are separate. I think... I learned docker over the weekend so i'm extremely new to all this haha.
## [6][Where did you learn Stimulus and Stimulus Reflex?](https://www.reddit.com/r/rails/comments/hy3dwd/where_did_you_learn_stimulus_and_stimulus_reflex/)
- url: https://www.reddit.com/r/rails/comments/hy3dwd/where_did_you_learn_stimulus_and_stimulus_reflex/
---

## [7][Platform to learn Ruby on rails](https://www.reddit.com/r/rails/comments/hy6iuz/platform_to_learn_ruby_on_rails/)
- url: https://www.reddit.com/r/rails/comments/hy6iuz/platform_to_learn_ruby_on_rails/
---
Hello everyone ,

I am returning to RoR world again , I  want to know if thee a platform  for RoR  like symfonycast or laracast for Symfony and Laravel .

Thank you in advance
## [8][How to build a book tracker in rails!](https://www.reddit.com/r/rails/comments/hxx6s6/how_to_build_a_book_tracker_in_rails/)
- url: https://www.reddit.com/r/rails/comments/hxx6s6/how_to_build_a_book_tracker_in_rails/
---
Hey guys , this series helps you understand the basics of rails through an interesting project. It takes nothing for granted, teaching you all the nitty gritty details. 

Enjoy.

 [https://www.youtube.com/watch?v=uEwu7D5G-hU&amp;list=PLB4RncStK2LUbl9VWLQAHznLJrYz2YMB4](https://www.youtube.com/watch?v=uEwu7D5G-hU&amp;list=PLB4RncStK2LUbl9VWLQAHznLJrYz2YMB4)
## [9][Starting a monthly newsletter for cool stuff I encounter on Rails, React and Graphql.](https://www.reddit.com/r/rails/comments/hxp9nm/starting_a_monthly_newsletter_for_cool_stuff_i/)
- url: https://www.reddit.com/r/rails/comments/hxp9nm/starting_a_monthly_newsletter_for_cool_stuff_i/
---
Starting a monthly newsletter on Ruby on Rails, React and GraphQL developers to read about some cool stuff happening in the community.

In a month we come across a lot cool stuff happening in the community, learn new things and contribute to the open-source world and newsletters is the best way to share some of those finding/blogs/videos.

This newsletter would consist of the following things:

* Dev Joke
* My Blogs.
* Reading/watching list.
* Open-source libraries/contributions.
* Rails changelog.
* Watching/Reading/Reviews.

If you'd like to read my monthly newsletter. Please do subscribe here [https://buttondown.email/abhaynikam](https://buttondown.email/abhaynikam)
## [10][How to use Rails in the server side to make a Web Game](https://www.reddit.com/r/rails/comments/hxuy62/how_to_use_rails_in_the_server_side_to_make_a_web/)
- url: https://www.reddit.com/r/rails/comments/hxuy62/how_to_use_rails_in_the_server_side_to_make_a_web/
---
I'm trying to make a webgame just for fun and start to learn how to make websites, i understand the basic things: HTML is the website, CSS is for make it very cool and Javascript is for animations, process data, etc. Well a simple summary, let me know if i have some bad concept please.

Now i don't understand how to run code(a simple bucle) in the server side, maybe i need a extended library but i would like to use RoR only. I would like if someone guide me for a good simple way to do what i need.

I use RoR because programming is my hobby and i've only done things in Ruby, i have used socket for Ruby to make a simple MMORPG and works fine, that's my experience.

Something to send messages to clients from server side(gameloop) and using javascript to apply the changes in client side.
## [11][Rails equivalent of Django F() Objects?](https://www.reddit.com/r/rails/comments/hxx6ee/rails_equivalent_of_django_f_objects/)
- url: https://www.reddit.com/r/rails/comments/hxx6ee/rails_equivalent_of_django_f_objects/
---
I’m wondering if there’s a library for rails or a rails native implementation of https://docs.djangoproject.com/en/3.0/ref/models/expressions/#f-expressions

EDIT:
Essentially I’m wanting to increment the number value of a field/value by 1 on multiple model instances.
