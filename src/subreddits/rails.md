# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hrnm2o/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/hrnm2o/personal_projects_show_off_your_own_project_andor/
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/hwehh6/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/hwehh6/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][Discovered SimpleCov, what can I do with it ?](https://www.reddit.com/r/rails/comments/hwf6um/discovered_simplecov_what_can_i_do_with_it/)
- url: https://www.reddit.com/r/rails/comments/hwf6um/discovered_simplecov_what_can_i_do_with_it/
---
Hi, I just discovered [SimpleCov](https://github.com/colszowka/simplecov) and looove the html output and the conclusions we can draw with it.

However, I can't really find a way to use it effectively. Thus I wonder, how do you use it (if you do of course) and what for ? 

&amp;#x200B;

I learnt about SimpleCov looking around Skunk that uses it at some point; this is IMO a great usage of SimpleCov.

&amp;#x200B;

I'd love a regular check of it on one of our projects that needs to be really well covered, maybe looking around Github action ? Anyway, just wanted to know if anyone uses it and what for. 

&amp;#x200B;

Thanks ! 

      _      _      _
    &gt;(.)__ &lt;(.)__ =(.)__
     (___/  (___/  (___/
## [4][Trouble with Elasticsearch DSL in a Ruby on Rails Project](https://www.reddit.com/r/rails/comments/hw5p69/trouble_with_elasticsearch_dsl_in_a_ruby_on_rails/)
- url: https://www.reddit.com/r/rails/comments/hw5p69/trouble_with_elasticsearch_dsl_in_a_ruby_on_rails/
---
Hey All, I'm trying to implement fuzzy search with Elasticsearch. I'm able to return results when I type in the query correctly, but now I am having trouble with fuzziness. I'm not having any luck with any of the online guides, does anyone have suggestions for a rails article with elasticsearch and a simple fuzzy search set up?  


`def search`  
 `@elasticsearchresults = Car.__elasticsearch__.search(`  
`{`  
 `query: {`  
 `match: {`  
 `message: {`  
 `query: params[:q],`  
 `fuzziness: 1`  
`}`  
`}`  
`}`  
`}).records`  
`end`
## [5][Having trouble with Capybara not waiting for remote form](https://www.reddit.com/r/rails/comments/hwdxnr/having_trouble_with_capybara_not_waiting_for/)
- url: https://www.reddit.com/r/rails/comments/hwdxnr/having_trouble_with_capybara_not_waiting_for/
---
I have a view where customers can change the delivery date of their shipments by selecting a date in [Flatpickr](https://flatpickr.js.org/). The flatpickr instance is connected to a remote `form_with`input field which submits the form on change via stimulus. 

This is all working manually in the browser but I obviously also want to test it automatically via system tests (I'm using Rails' system tests with rspec, selenium and capybara). However, it looks like Capybara doesn't wait long enough for the remote form to submit. I currently have the following test:

&gt; delivery_date = find("span", text: "27")

&gt; expect {

&gt; &gt;delivery_date.click

&gt; }.to change { shipment.reload.delivery_date }

which fails. However, if I change it to sleep for a second it works:

&gt; delivery_date = find("span", text: "27")

&gt; expect {

&gt; &gt; delivery_date.click

&gt; &gt; sleep 1

&gt; }.to change { shipment.reload.delivery_date }

Is there something that I'm doing wrong here like some syntax that I'm not following? The only thing that I manually added to rspec regarding system tests is the following in my `rails_helper.rb`:

&gt;  config.before(:each, type: :system) do

&gt;    driven_by :rack_test

&gt;  end

&gt;  config.before(:each, type: :system, js: true) do

&gt;    driven_by :selenium_chrome_headless

&gt;  end

which just changes the driver based on whether or not the test needs js (which I have enabled for the above test). I also tried changing the `Capybara.default_max_wait_time` to 10 in a before action without luck.
## [6][Security issues with generating temporary passwords](https://www.reddit.com/r/rails/comments/hw2ile/security_issues_with_generating_temporary/)
- url: https://www.reddit.com/r/rails/comments/hw2ile/security_issues_with_generating_temporary/
---
Hello. So I am developing an app with an application system. The user applies to my platform and then an admin or somebody with high enough permissions decides whether to accept the application or deny it. 

Let's say he accepts it. When he accepts it, the user is automatically created. I am wondering if it is safe to generate a temporary password with something like [Passgen gem](https://github.com/cryptice/Passgen) and then store it as the user's password until he changes it.

I would also need to send a confirmation email to his email address with a success message and a temporary password. Is it safe to send a password like that?

I am also wondering. If I hash that random password, how do I send it in plain text for a user to understand in the email?

Or is there any other, better way to handle this situation. Thanks!
## [7][How to insert font-awesome icons into action mailer views?](https://www.reddit.com/r/rails/comments/hw7ppm/how_to_insert_fontawesome_icons_into_action/)
- url: https://www.reddit.com/r/rails/comments/hw7ppm/how_to_insert_fontawesome_icons_into_action/
---
No problem embedding images, but haven't had luck leveraging FA-icons for this.
## [8][BCrypt::Errors::InvalidHash (invalid hash)](https://www.reddit.com/r/rails/comments/hvy13w/bcrypterrorsinvalidhash_invalid_hash/)
- url: https://www.reddit.com/r/rails/comments/hvy13w/bcrypterrorsinvalidhash_invalid_hash/
---
As of yesterday afternoon I am getting the error `BCrypt::Errors::InvalidHash (invalid hash)` anytime I attempt to create a User, or a log a user into my application. If I run `BCrypt::Password.create('password')` in the console I get the same error.

Happy to post any code people think is relevant, but from the searching I have done it appears to be a BCrypt issue that may be related to my OS, which is Ubuntu 20.04.

Any body have any idea's or suggestions?

Update: I ran an older application I have using virtually the same User login system and it works without issue to authenticate and create a user.
## [9][RoR + TailwindCSS Configuration](https://www.reddit.com/r/rails/comments/hvxur8/ror_tailwindcss_configuration/)
- url: https://www.reddit.com/r/rails/comments/hvxur8/ror_tailwindcss_configuration/
---
Been playing with the RailsBytes templates (awesome btw) lately, but I'm at a loss here for the Tailwind templates. 

I see 5 different templates for what strikes me as the same thing:

1) Tailwind (with config &amp; tailwindUI)

2) Tailwind (with config &amp; tailwindui &amp; PurgeCSS)

3) Tailwind CSS

4) Tailwind

5) Lucky - TailwindCSS

Can someone explain the differences between these? I don't know what TailwindUI is or how it differs from TailwindCSS, I know PurgeCSS is sort of what inspired Tailwind,  but not sure exactly how the two come together in here.
## [10][Grouping resources within the same table](https://www.reddit.com/r/rails/comments/hw8uch/grouping_resources_within_the_same_table/)
- url: https://www.reddit.com/r/rails/comments/hw8uch/grouping_resources_within_the_same_table/
---
Hello, I'm new to both Rails and web development.  
  
I'm working on an app that creates holiday schedules. The fall will have to related schedules, one for minor holidays and one for major holidays, but the summer will only have one.   

How would you guys go about linking the related schedules? I was thinking of adding an incrementing group_id attribute that would the multiple tables together. 
Any better ideas?
## [11][Rails.configuration.MYCONFIG versus App::MYCONFIG ??](https://www.reddit.com/r/rails/comments/hvvfu2/railsconfigurationmyconfig_versus_appmyconfig/)
- url: https://www.reddit.com/r/rails/comments/hvvfu2/railsconfigurationmyconfig_versus_appmyconfig/
---
Just pondering where best to put some app level constants in Rails 6. For brevity I'd prefer the later method (\`App::MYCONFIG\`) but I'm wondering what the difference is compared to \`Rails.configuration.MYCONFIG\` and why Rails makes the configuration property at all since there are clearly [lots of other ways to do it](https://stackoverflow.com/questions/4110866/ruby-on-rails-where-to-define-global-constants)...?

&amp;#x200B;

**Clarifications:**

The  actually configs are various validator params that are common to  several models and are also used in several views. For instance:   DISPLAY\_NAME\_REGEX.  Its used in the input field for client side  validation, then used in the model validator. And it's also shared  amongst several, not-entirely-related models.

It's  one of those things, I start off with it in the model but then realise  it needs to be shared amongst a few models and then used in the views....then I get confused on  where best to put it...  Custom class maybe??

&amp;#x200B;
## [12][Is DidYouMean supposed to fetch records? (also rails seemingly ignores `RUBYOPT="--disable-did_you_mean"`)](https://www.reddit.com/r/rails/comments/hvwcul/is_didyoumean_supposed_to_fetch_records_also/)
- url: https://www.reddit.com/r/rails/comments/hvwcul/is_didyoumean_supposed_to_fetch_records_also/
---
Sorry long post :)

## Background

I implemented a `failsafe_render` method that will show a generic error box if the partial fails to render for whatever reason (it will still notify the exception though). This is for some partials in our back office that employees can play around with on their own - Ruby is amazingly easy to grasp for non-programmers for some simple statistical aggregations, etc. These play fields we don't code check much (edge cases, etc.).

If the partial breaks it displays the exception message and 3 lines of backtrace. This is because whilst I get the exception in my error tracker, they might want to fix the edge case in their playfield on their own.

## Problems

While testing I noticed that some errors causes the site to render very slow (at least in development). Especially if the error happens a lot of times (like a row partial that repeats and breaks 30 times).

Consider a page with 30 failing partials caused by different issues and their respective load time:

* `# no error` Completed 200 OK in 2282ms (Views: 1580.7ms | ActiveRecord: 145.0ms | Allocations: 1863991)
* `&lt;%= 1/0 %&gt;` Completed 200 OK in 619ms (Views: 537.9ms | ActiveRecord: 52.2ms | Allocations: 630146)
* `&lt;%= order.idx %&gt;` Completed 200 OK in 729ms (Views: 648.7ms | ActiveRecord: 50.3ms | Allocations: 786558)
* `&lt;%= orderx.id %&gt;` Completed 200 OK in 33064ms (Views: 32856.5ms | ActiveRecord: 162.6ms | Allocations: 33043056)


Now `order.idx` will cause

    NoMethodError: undefined method `idx' for #&lt;Order:0x00007fc676087408&gt;
    Did you mean?  id
                   id=
                   id?

And `orderx.id` will cause

    NameError: undefined local variable or method `orderx' for #&lt;#&lt;Class:0x00007fc677bde648&gt;:0x00007fc677bdca78&gt;
    Did you mean?  order
                   @orders

## Causes

Watching the console I can see that calling the exception's `#message` method causes DB queries and since the message is displayed and being called by the exception tracker it does the whole thing twice.

    ↳ app/views/shared/_failsafe_erb_error.html.erb:8:in `message'
    User Load (0.6ms)  SELECT `users`.* FROM `users` WHERE `users`.`id` IN (...)
    ↳ app/views/shared/_failsafe_erb_error.html.erb:8:in `message'
    OrderItem Load (0.6ms)  SELECT `order_items`.* FROM `order_items` WHERE `order_items`.`order_id` IN (...)
    ↳ app/views/shared/_failsafe_erb_error.html.erb:8:in `message'
    AffiliateTag Load (0.4ms)  SELECT `affiliate_tags`.* FROM `affiliate_tags` WHERE `affiliate_tags`.`id` IN (...)
    ↳ app/views/shared/_failsafe_erb_error.html.erb:8:in `message'
    Voucher Load (0.6ms)  SELECT `vouchers`.* FROM `vouchers` WHERE `vouchers`.`id` IN (...)

As you can see `ex.message` causes "random" queries. These come from eager loading but when running without any errors or with a division by zero it will - for example - select order items twice but with the orderx error it is being called 362 times (4 of which are not cached).  
Why 4 uncached now, who knows but why does it call everything so often, in fact 2 are normal and we have 30 failing partials, so (362-2)/30=12 times per suggestion/exception?  
Just to figure out those 2 suggestions. The thing is the queries (since cached) are not the issue but it's very CPU intensive apparently to "Did you mean?" something on the view context.

## Solutions

Any ideas? Disabling the feature alltogether seems fine by me but I am seemingly unable to disable DidYouMean with Rails for some reason?

    $ pry
    [1] pry(main)&gt; class Foo; def aa; end; end
    [2] pry(main)&gt; Foo.new.a
    NoMethodError: undefined method `a' for #&lt;Foo:0x00007fbae441d3f8&gt;
    Did you mean?  aa

    $ RUBYOPT="--disable-did_you_mean" pry
    [1] pry(main)&gt; class Foo; def aa; end; end
    [2] pry(main)&gt; Foo.new.a
    NoMethodError: undefined method `a' for #&lt;Foo:0x00007f85da3bf0f8&gt;

    $ RUBYOPT="--disable-did_you_mean" DISABLE_SPRING=1 rails c
    Loading development environment (Rails 6.0.3.2)
    [1] pry(main)&gt; class Foo; def aa; end; end
    [2] pry(main)&gt; Foo.new.a
    NoMethodError: undefined method `a' for #&lt;Foo:0x00007ff7497bf4d8&gt;
    Did you mean?  ai
                   aa
