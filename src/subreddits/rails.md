# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/)
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
## [2][Friendly-id and a bug with the sequence of a title with two words](https://www.reddit.com/r/rails/comments/gmlktq/friendlyid_and_a_bug_with_the_sequence_of_a_title/)
- url: https://www.reddit.com/r/rails/comments/gmlktq/friendlyid_and_a_bug_with_the_sequence_of_a_title/
---
Following the tips of the author's gem (soruce: [https://github.com/norman/friendly\_id/issues/480](https://github.com/norman/friendly_id/issues/480)), I added a title\_and\_sequence like this

     extend FriendlyId   
     friendly_id :slug_candidates, use: [:slugged, :finders]
    
     def slug_candidates
         [:title, :title_and_sequence] end def title_and_sequence
         slug = title.to_param
         sequence = Book.where("slug LIKE '#{slug}--%'").count + 2
     "#{slug}--#{sequence}"
     end

but I noted that there is a bug. If I use as title "one", it works correctly

    /books/one 
    /books/one-2 
    /books/one-3 
    /books/one-4 etc.

But if I use a title with two words, for example "one two" I have a "bug" after the second result.

    /books/one-two 
    /books/one-two-2 
    /books/one-two-5j12-123j-afu4-jasdk 
    /books/one-two-9as6-k273-ewu1-87srt 
    etc.

**Do you have any idea about how to solve it?**
## [3][Being a "one-man army" with Rails](https://www.reddit.com/r/rails/comments/gm4a0y/being_a_oneman_army_with_rails/)
- url: https://www.reddit.com/r/rails/comments/gm4a0y/being_a_oneman_army_with_rails/
---
When I started learning Rails, the immediate thing that occurred to me was that, for the first time, I'm "competently" doing full-stack development.

And by "competently" I mean that the idea of convention over configuration frees up a lot of your mental resources to focus on actual business needs.  As a Rails developer you are building the database - if it doesn't already exist - and you are creating models that connect to and modify the database through migrations.  You're also building the API using controllers, and building out the views and probably even digging into UI efforts like JS frameworks and organizing CSS/Sass/whatever-framework-or-library.

I know of at least one person who maintains a Rails app that he developed over 10 years ago, and it has been his primary source of income this entire time.

So my question is, how many of you folks out there are doing this?  How many "one-man armies" are out there?  And if you aren't one, do you know of others who are pulling this off?

Would be nice to know, if for no other reason, because it's inspiring.
## [4][belongs_to &amp; has_many - do you really need both?](https://www.reddit.com/r/rails/comments/gmnmt6/belongs_to_has_many_do_you_really_need_both/)
- url: https://www.reddit.com/r/rails/comments/gmnmt6/belongs_to_has_many_do_you_really_need_both/
---
I am following a training example with microposts and users, and it says you need belongs\_to and has\_many. Is that not duplication? What happens if you only have one and not the other?
## [5][Rails 6 Wysiwyg integration.](https://www.reddit.com/r/rails/comments/gmbw65/rails_6_wysiwyg_integration/)
- url: https://www.reddit.com/r/rails/comments/gmbw65/rails_6_wysiwyg_integration/
---
I've followed a few instructional videos using both Wysiwyg rails and trix editor, but neither seem to be working on my app. 

Not sure if I'm missing anything, but is Webpack messing up the flow of it all? Does Rails 6 not support them? If you have any workaround or functional source for this, I'd very much appreciate it.
## [6][Any good tutorials for someone with experience in Node/Express?](https://www.reddit.com/r/rails/comments/gmanc2/any_good_tutorials_for_someone_with_experience_in/)
- url: https://www.reddit.com/r/rails/comments/gmanc2/any_good_tutorials_for_someone_with_experience_in/
---
Hi, I've just starting learning rails and was wondering if anyone had any good resources for someone with more exp working with Node/Express. Any help is appreciated, thanks!
## [7][How would you integrate this?](https://www.reddit.com/r/rails/comments/gmb3dk/how_would_you_integrate_this/)
- url: https://www.reddit.com/r/rails/comments/gmb3dk/how_would_you_integrate_this/
---
Hey Rails superstars!

I'm trying to utilize this library ([https://github.com/caroso1222/notyf](https://github.com/caroso1222/notyf)) for flash notifications in my rails app. So far, this is what I've been able to do without success:

* Application controller: I've added custom flash types of 'error' and 'success' as follows:  


&amp;#8203;

    add_flash_types :error, :success

&amp;#x200B;

* Application JS: My code looks like

&amp;#x200B;

    //= require rails-ujs 
    //= require turbolinks 
    //= require_tree . 
    
    var notyf = new Notyf();

&amp;#x200B;

* Application SCSS includes this line:

&amp;#x200B;

    @import url('https://cdn.jsdelivr.net/npm/notyf@3/notyf.min.css');

&amp;#x200B;

* ..then in my application layout file, I've done included this:

&amp;#x200B;

    &lt;% unless flash.empty? %&gt;
     &lt;script type="text/javascript"&gt;
 &lt;% flash.each do |f| %&gt;
      &lt;% type = f[0].to_s %&gt;
 notyf.&lt;%= type %&gt;('&lt;%= f[1] %&gt;');
 &lt;% end %&gt;
 &lt;/script&gt;

    &lt;% end %&gt;

&amp;#x200B;

..and when I run this, I see an error in the console "Uncaught reference error...notyf is not defined"...what am I missing?

Thanks...and stay safe!
## [8][Rails Inflections acronym conflicts between libraries](https://www.reddit.com/r/rails/comments/gm1m3m/rails_inflections_acronym_conflicts_between/)
- url: https://www.reddit.com/r/rails/comments/gm1m3m/rails_inflections_acronym_conflicts_between/
---
I'm about to start a Rails 6 project. I'm intending to use Grape gem together with ActivityNotification gem. Since Rails 6, to make Grape works, according to the documentation, we have to include `inflect.acronym 'API'` in `inflections.rb` file, which will turn all 'api', 'Api' into 'API'. This will cause problems with ActivityNotification gem's class names. Have anyone encountered this? Do you have any solutions?

If worse comes to worst, I think I would have to either use another lib or fork ActivityNotification to change all class names to use the acronym API or roll my own solution...
## [9][[Advice] Stripe Tutorials?](https://www.reddit.com/r/rails/comments/gm1r1m/advice_stripe_tutorials/)
- url: https://www.reddit.com/r/rails/comments/gm1r1m/advice_stripe_tutorials/
---
I am trying to use stripe-ruby gem for the first time and I have multiple plans according to the date and year set from the dashboard. I was wondering if there are some tutorials or a particular way i can follow to give me a heads up on what to do and where can i put stuff so that I don't mess up the payments and have all the required plans and subscriptions information as well. Any tutorials that you guys suggest?
## [10][Shareable Rails app templates](https://www.reddit.com/r/rails/comments/glljbs/shareable_rails_app_templates/)
- url: https://www.reddit.com/r/rails/comments/glljbs/shareable_rails_app_templates/
---
There's been some discussion lately about Rails app templates and how they could be a lot more useful. After having spent a ton of time building Rails app templates for GoRails and Jumpstart, I can definitely feel the pain point. Nate Berkopec [tweeted](https://twitter.com/nateberkopec/status/1259634098353537024) about templates as well recently.

One of the things that is definitely frustrating is installing something like Bootstrap. You've got a bunch of different files to configure and it's easy to miss a step. We've packaged all that into a [Rails app template](https://railsbytes.com/public/templates/x9Qsqx) so all you have to do is run a single command to install it.

These app templates are definitely trickier to build because they'll need to verify Rails version and other dependencies (like is webpacker installed, etc), but I think if done well they can save Rails devs a ton of time. Plus, app templates can be run against existing Rails apps, not just brand new ones. 

So here you have it: [railsbytes.com](https://railsbytes.com)

It took u/king601 and I about 8 hours to build this first version so we know there's a lot to improve, but we hope it can serve as a foundation to help make adding common gems, features, and other things in Rails much easier.

Some interesting examples:

Installing Devise is interactive. It can ask you what model you want and what extra fields so you don't have to remember any of the commands: [https://railsbytes.com/public/templates/X8Bsjx](https://railsbytes.com/public/templates/X8Bsjx) 

Same thing applies to gems like HoneyBadger that need to ask for an API key: [https://railsbytes.com/public/templates/zNPsmV](https://railsbytes.com/public/templates/zNPsmV)

You can have a Rails app template trigger chain other Rails app templates together like this one that installs Rspec, Factorybot, and Standardrb: [https://railsbytes.com/public/templates/V4YsyX](https://railsbytes.com/public/templates/V4YsyX) 

Wanted to share this with you guys and we're curious to hear what you guys think. It seems like it's got a ton of potential to do all sorts of interesting things.
## [11][ISeries Rails Environment and development](https://www.reddit.com/r/rails/comments/glv47u/iseries_rails_environment_and_development/)
- url: https://www.reddit.com/r/rails/comments/glv47u/iseries_rails_environment_and_development/
---
Happy Quarantine everyone. We are looking to start running some web apps on top of some of our clients iSeries systems (System i, AS/400). We got the default rails environment up and running on one of the servers. Does anyone have any recommendations on resources getting started guides or something like that for rails on the iSeries? 

Thanks in advance
