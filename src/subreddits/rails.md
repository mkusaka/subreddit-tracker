# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/hwehh6/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/hwehh6/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Best open source/community maintained project templates?](https://www.reddit.com/r/rails/comments/hz8p72/best_open_sourcecommunity_maintained_project/)
- url: https://www.reddit.com/r/rails/comments/hz8p72/best_open_sourcecommunity_maintained_project/
---
Recently I read an article about the Rails templates (sorry, I didn't saved it, but here is the [docs](https://guides.rubyonrails.org/rails_application_templates.html)). I thought that there might be open source or community maintained rails templates, and I found:

* [This thoughtbot gem that builds rails app with certain configurations](https://github.com/thoughtbot/suspenders)
* [This template](https://github.com/dao42/rails-template)
* [Also this template](https://github.com/mattbrictson/rails-template)
* [And some paid templates](https://hixonrails.com/)

&amp;#x200B;

Then I come here to r/rails to see if I could see more, but I didn't found any post about this, that's why I wanted to ask if you use a template for new projects or have any favorite that you want to share. Thanks! :)
## [3][Would someone be able to explain Redis clients/connections with actioncable to me like I'm a dummy?](https://www.reddit.com/r/rails/comments/hz14wg/would_someone_be_able_to_explain_redis/)
- url: https://www.reddit.com/r/rails/comments/hz14wg/would_someone_be_able_to_explain_redis/
---
So let's pretend for a minute that I am a Rails dev that knows just enough about how the actual servers work to be dangerous.  So when I am connecting to actioncable through an api with redis on heroku, what constitutes a client?  I am asking, first because I don't understand it and I want to learn things.  But second, because I see on heroku there is a 20 client limit on the hobby version heroku redis add-on.  Does that mean that only 20 people can connect to actioncable at a time?  Or is that some other kind of client/connection?

EDIT:  So based on this https://www.ably.io/blog/rails-actioncable-the-good-and-the-bad, what I'm getting is that action cable runs as it's own process, that has one connection to redis.  So every use would not be a client, but the action cable server/process would be one connection/client.  At least that's the way I'm reading it.
## [4][convert rails api to rails app](https://www.reddit.com/r/rails/comments/hzemhx/convert_rails_api_to_rails_app/)
- url: https://www.reddit.com/r/rails/comments/hzemhx/convert_rails_api_to_rails_app/
---
I have a Rails Api that has a Vue front end (in it's own directory called client). Now we're wanting to add admin screen but the client has indicated they'd prefer those just be rails pages like RailsAdmin. Is there a good guide for converting a Rails Api to a full Rails App while not blowing away the things you've already done? Most guides seem to indicate running rails-new again and carefully selecting things not to overwrite but that doesn't really tell you if you're missing anything once it's done.
## [5][Why is the Rails object not defined in Javascript?](https://www.reddit.com/r/rails/comments/hz9v3y/why_is_the_rails_object_not_defined_in_javascript/)
- url: https://www.reddit.com/r/rails/comments/hz9v3y/why_is_the_rails_object_not_defined_in_javascript/
---
My understanding was that an object called Rails should be available globally in Javascript. See this link:

 [https://www.rubyguides.com/2019/03/rails-ajax/](https://www.rubyguides.com/2019/03/rails-ajax/) 

&amp;#x200B;

However, opening the Chrome console and typing in Rails gives me:

    VM514:1 Uncaught ReferenceError: Rails is not defined
        at &lt;anonymous&gt;:1:1

I found [this issue on Github](https://github.com/rails/rails/issues/36686), but it's unclear what the conclusion is. Another user reports that the Rails object is undefined, someone provides a specific sequence of statements to import the object, and then the issue is closed. But why isn't the object available by default? 

If I add the following code to `application.js`, then the Rails object is defined in the Chrome javascript console:

    import Rails from "@rails/ujs";
    window.Rails = Rails;
    if(Rails.fire(document, "rails:attachBindings")) {
        Rails.start();
    }

Is it normal to have to add this for the Rails object to be defined?
## [6][What are the dangers of this code?](https://www.reddit.com/r/rails/comments/hz8xvp/what_are_the_dangers_of_this_code/)
- url: https://www.reddit.com/r/rails/comments/hz8xvp/what_are_the_dangers_of_this_code/
---
Hello! 

I have a situation where I have two models connected with an association. 

I want to allow users to copy a set of images from a current set. Sets are just a bunch of associations between the set and the individual images. 

I know that convention says that I should move all this logic from the controller to the model, and I will be doing that next. But I started working on this problem in the controller and wanted to finish my thoughts there to get it to work first. I named the sets something else because they were a reserved word, but using it here for clarity. 

    def copy         
      set = Set.find(params[:id])
        set.line_items.each do |item|
          @line_item = LineItem.new
          @line_item = item.dup  
          @line_item.bin_id = current_bin.id
          @line_item.set_id = nil
          @line_item.save
        end
        redirect_to :controller =&gt; 'exemplars', :action =&gt; 'index'
    end

I have a link from the set page to the copy action, and everything works. I wanted to create a new @line_item (the association) for each existing line item in the set we are copying from. Then dup it, then set the bin.id to the current bin which is a temporary holding model, and then set the @line_item.set_id to nil. Then it saves each one. 

I know there are probably better ways to do this, but my real question is what happens if something interrupts this copy process? How common would that be, and is there a way to safeguard against that?
## [7][Can't log in after upgrading 4.0 -&gt; 4.2, advice please](https://www.reddit.com/r/rails/comments/hyvhjd/cant_log_in_after_upgrading_40_42_advice_please/)
- url: https://www.reddit.com/r/rails/comments/hyvhjd/cant_log_in_after_upgrading_40_42_advice_please/
---
Hey everyone.

Title describes the situation.  I used the rake app:upgrade tool to get the various configs updated.  The site launches and appears to run just fine, except login doesn't work.

I'm switching over to other work but I'll be looking at this again later on in the week and I thought I would throw out the question to everyone here.  Do you have any advice on what I should be looking for?  I'm really not super familiar with rails (used it back in the 1.x/2.x days only), so it's not clear to me what the problem could be.  My initial thought is to check if the way passwords are salted/hashed has changed.

Thanks everyone!
## [8][Importing a database dump into a rails project on Docker](https://www.reddit.com/r/rails/comments/hyupi7/importing_a_database_dump_into_a_rails_project_on/)
- url: https://www.reddit.com/r/rails/comments/hyupi7/importing_a_database_dump_into_a_rails_project_on/
---
Let me preface this post with this: I started learning docker over the weekend, so if you see something here that is bad practice, please please call me out on it.

I know the command in Rails 4+, it's

    rails db &lt; ./path/to/db_dump.sql

but I don't know how to do it with Rails 2.3. I know it's old, but it's all I have to work with.

I cannot use the MySQL CLI. I think have to use a Rake task because the application is dockerized, so my MySQL container doesn't "know" about my rails_app_development database.

Does anyone know the Rails 2.3 command for this, if there is one?

## Things I've tried:

**1. This one i feel like has gotten me the closest**

```ruby
namespace :db do
  task :import_from_source =&gt; :environment do
    Encoding.default_external = Encoding::UTF_8
    Encoding.default_internal = Encoding::UTF_8
    ActiveRecord::Base.connection.execute(IO.read("./file.sql"))
  end
end
```

So i basically just made a rake task. Then in docker I can run: 

```
$ docker-compose run app rake db:import_from_source
```

And it appears that everything worked. In fact, I see this in my logs:

```
app_1  | -- Dump completed on 2020-07-24 6:20:02
app_1  |
app_1  |   SQL (0.7ms)   SET SQL_AUTO_IS_NULL=0
```

However when I refresh the page, I still get a "mysql table dosen't exist" error

**2. This doesn't work and I'm not sure why**

First I run:

```
docker-compose run --rm app rake db:create
```

and that works, but when I go into my db container using `docker-compose run db bash` and run

```
mysql -u root -p rails_db_development &lt; /db_dump/db_dump.sql
```

I get this error:

```
ERROR 1049 (42000): Unknown database 'rails_db_development'
```

Even though I created it in the app container.
## [9][Cant access rails web console on Macos Catalina](https://www.reddit.com/r/rails/comments/hypwj0/cant_access_rails_web_console_on_macos_catalina/)
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
## [10][The more I look take a look at different apps, I see that they're just CRUD applications.](https://www.reddit.com/r/rails/comments/hybezw/the_more_i_look_take_a_look_at_different_apps_i/)
- url: https://www.reddit.com/r/rails/comments/hybezw/the_more_i_look_take_a_look_at_different_apps_i/
---
I'm not sure if I'm thinking about this correctly, but I'm starting to see that the applications that I'm looking at are all perform the same basic functions.

To elaborate, I see that most apps read and display data from a database. For example, an exercise web app shows users the exercise name, target muscles, and a video demonstration. Another example, a restaurant booking app, the app looks at how many time slots are not reserved and shows the data to the person checking for available times. Another example, Mint the personal finance app. It connects with your bank and it shows the user spending and alerts.

There's a lot of times when I visited a website and said to myself "Oh this app is just here to organize and regurgitate my data."

The question I have is, is this considered "business logic", or is there another layer to developing applications?
## [11][Webpack Environment Config File](https://www.reddit.com/r/rails/comments/hygoa4/webpack_environment_config_file/)
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
