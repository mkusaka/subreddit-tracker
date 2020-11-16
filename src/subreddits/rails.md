# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/jsthk8/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/jsthk8/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Building a user onboarding | Ruby on Rails Livestreaming](https://www.reddit.com/r/rails/comments/jv67po/building_a_user_onboarding_ruby_on_rails/)
- url: https://www.reddit.com/r/rails/comments/jv67po/building_a_user_onboarding_ruby_on_rails/
---
Hi everyone! Every week I live-stream some Ruby on Rails and once in a while add updates on here in case anyone would like to check them out.

Over the last few weeks, we've been working on building a step-by-step user-onboarding flow for the app, here are links to the relevant streams:

Let's do some design/planning! - [https://youtu.be/YDS_DKsCXyo](https://youtu.be/YDS_DKsCXyo) 
Some of the initial assumptions I'd made about the app changed after looking at some other examples, so drew out the changes and how they would affect the onboarding flow for users.

Finishing our user onboarding - [https://youtu.be/d-4DkTq3e2s](https://youtu.be/d-4DkTq3e2s)
Built out some of the onboarding flow, we wanted to ask the user for their details last, after they'd already entered the details to create their site.

Tidying up our user onboarding - [https://youtu.be/SqrjDCyY9Nw](https://youtu.be/SqrjDCyY9Nw)
Went back and sorted out stuff that I'd missed along the way.

I will go back and add some chapter markers to those videos when I get a some time this evening.
## [3][What are good open-source apps using RoR to learn from?](https://www.reddit.com/r/rails/comments/jumg91/what_are_good_opensource_apps_using_ror_to_learn/)
- url: https://www.reddit.com/r/rails/comments/jumg91/what_are_good_opensource_apps_using_ror_to_learn/
---
Hi! I'm a Ruby on Rails beginner. Do you have any recommendations about well written open-source apps in Rails?
## [4][Preventing flaky tests due to ambiguous IDs](https://www.reddit.com/r/rails/comments/juz0uv/preventing_flaky_tests_due_to_ambiguous_ids/)
- url: https://www.reddit.com/r/rails/comments/juz0uv/preventing_flaky_tests_due_to_ambiguous_ids/
---
Hi Rails! I have a bit of a funny situation so I'll try to explain my problem as best I can and see if you have any comments on my approach to solve it.

# Scenario

I work on a medium sized Rails app and test with Rspec, FactoryBot, and DatabaseCleaner. We have around 30 people submitting PRs to the code + test suite. We're on Postgres.

&amp;#x200B;

# Problem

Sometimes we make mistakes in our specs - we pass in an ID of one record when we meant to pass the ID from another. These specs pass in certain situations (more below), and fail in others. Flaky specs suck.

# Example

```ruby
describe SomeController do
  let(:cat) { create(:cat) }
  let(:dog) { create(:dog) }

  subject { Cat.find(dog.id) }

  it { is_expected.to ba_a Cat}
end
```

In this (extremely contrived) example, the person clearly meant to pass `cat.id` to the `Cat.find` call.  In any case, this leads to two scenarios where the spec either passes or fails:

**PASS** \- If the spec starts the run with a clean database (all ID sequences starting from 1), then the spec will pass because `cat` and `dog` both have an `ID:1`. As far as Cat's finder is concerned, we passed the correct ID.

**FAIL** \- If other specs in the suite before this one have run, the spec will fail. This is expected since we're passing Dog's ID into Cat's finder, and they are most likely different.

&amp;#x200B;

# Solution?

I want to create a global sequence for FactoryBot, starting at some random number and incrementing for each record that uses it.

This ID will be used in every factory, like so


```ruby
# Global, shared sequence
FactoryBot.define do
  sequence(:random_id, rand(1000000))
end

...

FactoryBot.define do
  factory :cat do
    id { generate(:random_id) }
    eyes { 'blue' }
  end
end
```

If we look back at our example, the test would now consistently fail until the author noticed the mistake and replaced the call with [`cat.id`](https://cat.id)

&amp;#x200B;

**I'm wondering if anyone else has dealt with cases like this - and if there's something obvious that I'm missing as a better approach (besides, of course, not making mistakes ;) )**
## [5][Moving everything to another hard drive?](https://www.reddit.com/r/rails/comments/juv0ij/moving_everything_to_another_hard_drive/)
- url: https://www.reddit.com/r/rails/comments/juv0ij/moving_everything_to_another_hard_drive/
---
Hello! I just was wondering what the process is in moving my whole environment over to my 2nd hard drive. I have a 250gb ssd for my C drive and that's what my files are on as well as what all the dependencies are installed on. So I'd have to install all of those onto the hard drive as well if I were to just drag and drop my entire "Code" file with all my projects into it right? I just don't want to miss a step and have all my projects be useless on this machine. I'm just trying to free up space on my C drive for just OS necessities and a few other programs. Thanks in advance!
## [6][jQuery + Device/Login](https://www.reddit.com/r/rails/comments/jun8t7/jquery_devicelogin/)
- url: https://www.reddit.com/r/rails/comments/jun8t7/jquery_devicelogin/
---
Hey guys,

I have this jQuery script to make a live count animation

    jQuery(document).ready(function() {
    	function count($this){
    		var current = parseInt($this.html(), 10);
    		current = current + 1; /* Where 1 is increment */
    		$this.html(++current);
    		if(current &gt; $this.data('count')){
    			$this.html($this.data('count'));
    		} else {
    			setTimeout(function(){count($this)}, 50);
    		}
    	}
    
    	jQuery(".stat-count").each(function() {
    		jQuery(this).data('count', parseInt(jQuery(this).html(), 10));
    		jQuery(this).html('0');
    		count(jQuery(this));
    	});
    });

&amp;#x200B;

and in the view page I have this

    &lt;font class="stat-count"&gt;&lt;%= current_user.score %&gt;&lt;/font&gt;

But I want to execute it ONLY after the login. Just one time (i am using the device gem).  
**How to do?**

Adding it as count.js in \\app\\assets\\javascripts it is repeated every time in every page.
## [7][Karachi circular railway](https://www.reddit.com/r/rails/comments/jut6ed/karachi_circular_railway/)
- url: https://youtu.be/oZuXf1sHTPw
---

## [8][Help debugging &lt;Errno::ENOENT: No such file or directory - getcwd&gt;](https://www.reddit.com/r/rails/comments/juntg1/help_debugging_errnoenoent_no_such_file_or/)
- url: https://www.reddit.com/r/rails/comments/juntg1/help_debugging_errnoenoent_no_such_file_or/
---
When I run tests in Jenkins as part of CI, the tests fail with 500 randomly while it doesn't fail locally on my machine.   
It fails with this reason 

    &lt;Errno::ENOENT: No such file or directory - getcwd&gt;

And this is the backtrace I could find. 

    
    ruby-2.4.7/gems/activesupport-4.2.11.3/lib/active_support/dependencies.rb:274:in `require'
    ruby-2.4.7/gems/activesupport-4.2.11.3/lib/active_support/dependencies.rb:274:in `block in require'
    ruby-2.4.7/gems/activesupport-4.2.11.3/lib/active_support/dependencies.rb:240:in `load_dependency'
    ruby-2.4.7/gems/activesupport-4.2.11.3/lib/active_support/dependencies.rb:274:in `require'
    ruby-2.4.7/lib/ruby/2.4.0/open-uri.rb:268:in `open_http'
    ruby-2.4.7/lib/ruby/2.4.0/open-uri.rb:741:in `buffer_open'
    ruby-2.4.7/lib/ruby/2.4.0/open-uri.rb:212:in `block in open_loop'
    ruby-2.4.7/lib/ruby/2.4.0/open-uri.rb:210:in `catch'
    ruby-2.4.7/lib/ruby/2.4.0/open-uri.rb:210:in `open_loop'
    ruby-2.4.7/lib/ruby/2.4.0/open-uri.rb:151:in `open_uri'
    ruby-2.4.7/lib/ruby/2.4.0/open-uri.rb:721:in `open'
    ruby-2.4.7/lib/ruby/2.4.0/open-uri.rb:35:in `open'

I not sure how and what file it tries to open since this doesn't happen when I run the tests locally. Even in CI this happens only randomly.   


Can someone help me understand how to debug this?
## [9][Should models use direct associations or through?](https://www.reddit.com/r/rails/comments/juhtou/should_models_use_direct_associations_or_through/)
- url: https://www.reddit.com/r/rails/comments/juhtou/should_models_use_direct_associations_or_through/
---
I am wondering when it comes to scaling, is using through associations a better way to model a database? I am assuming the fewer columns, the better?

&amp;#x200B;

Assuming we have the following models

&amp;#x200B;

**Order**

has\_many line\_items

**Product**

*#product on the website*

has\_many line\_items, optional: true

belongs\_to vendor

**LineItem**

*#this is the products in the cart*

belongs\_to product

**Vendor**

*#vendor of the product*

has\_many products

has\_many line\_items through products

&amp;#x200B;

When an order comes in, the order will have many line items.  Is there any reason to have a vendor\_id association in the LineItem model? 

Or is \`through:\` the best case scenario?

I assume through is the best way to go, especially for database bloat when scaling on unneeded space being ocupied but I am wondering is there is any benefit to not using a through association
## [10][Best chatbot gems/libraries to use for rails app?](https://www.reddit.com/r/rails/comments/jucos4/best_chatbot_gemslibraries_to_use_for_rails_app/)
- url: https://www.reddit.com/r/rails/comments/jucos4/best_chatbot_gemslibraries_to_use_for_rails_app/
---
Just need a basic text chatbot to save information to a model &amp; wondering what you guys would recommend for gems/tutorials to do this with a rails app.
## [11][a bootstrap theme implementation with rails !!!](https://www.reddit.com/r/rails/comments/ju9s0n/a_bootstrap_theme_implementation_with_rails/)
- url: https://www.reddit.com/r/rails/comments/ju9s0n/a_bootstrap_theme_implementation_with_rails/
---
Hello Folks, spent 3 days trying to implement an html template to new project with rails the problem is  the theme based on Webpack and laravel-mix which compiles Sass, ES6 JavaScript, handles production builds, watchers, multiple CSS themes and more. This entire workflow is contained into an installable package named theme-mix. any idea what that means please help ?

wtf how i can use laravel-mix with rails webpacker ? any idea i would appreciate that .

&amp;#x200B;

https://preview.redd.it/5azi6qegx9z51.png?width=2056&amp;format=png&amp;auto=webp&amp;s=4ee0251ef67326c1c4617b9246716ad3d4121446

https://preview.redd.it/n586mregx9z51.png?width=2560&amp;format=png&amp;auto=webp&amp;s=30d1783293e80e5e20fb6cd26bb8d260f660407d

https://preview.redd.it/w1ngsf4xx9z51.png?width=2560&amp;format=png&amp;auto=webp&amp;s=e3e2417e7a21226c59a7385e7e4d5683e0c3dca3

https://preview.redd.it/lawkcregx9z51.png?width=2338&amp;format=png&amp;auto=webp&amp;s=112fbf66ca9ab1a5f629da46272c9278438fb75a

https://preview.redd.it/wb52dregx9z51.png?width=2416&amp;format=png&amp;auto=webp&amp;s=65e9debd41049ca5d06949c8c93ed193fffeb9e5

https://preview.redd.it/z034m8uwx9z51.png?width=2518&amp;format=png&amp;auto=webp&amp;s=f08e76ccdb00d3011fdb117e85c499c0b9f67b56
