# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hrnm2o/personal_projects_show_off_your_own_project_andor/)
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
## [3][How do you think votes or scores are handled in gargantuanly gigantic web apps? Help me guess how they do it...](https://www.reddit.com/r/rails/comments/hude4e/how_do_you_think_votes_or_scores_are_handled_in/)
- url: https://www.reddit.com/r/rails/comments/hude4e/how_do_you_think_votes_or_scores_are_handled_in/
---
I want to make a [reddit-like web app](https://www.reddit.com/r/RedditAlternatives/comments/hi97fz/list_of_active_reddit_alternatives_v5/) using Rails, but I've always wondered how these platforms handle so much traffic, like votes, post submissions, posts/comments sorting (hot, rising, etc.) in Reddit or likes/reactions in Facebook, there must be millions of these events within a couple of hours. Don't tell me Rails is not efficient for these kinds of apps, GitHub was built on Rails.

I come up with very vague ideas, like:

* When making a comment or upvoting/downvoting, run a background job to save that record (`comment` or `vote`) in the database. I guess to calculate the number of votes or comments we shouldn't run a `count` query, but the best option would be to have columns like `votes_count`, `comments_count`, etc., which should be updated by the same background jobs (obviously we would have to use a huge pool of mutexes to atomically increase the values of these columns)
* I guess that there are processes that calculates and updates the score of each post submission and comment and saves them (in columns like `score_hot`, `score_rising`, etc.) from time to time. This may be done with an ultra optimized Postgres (or another DBMS) function and for the sort-paginate (for each subreddit or Facebook group) we could use some Redis low-level caching, that way we don't do a query for each visit, only once every few minutes.
* etc.

But what happens if for some reason several instances of Redis die and I can't cache for a couple of seconds, the databases would explode doing queries, and how are the logs handled? I'm full of questions, how do you think these apps handle all these stuff that seem unmanageable? What kind of databases are most useful in these cases? Or is it enough with postgres for most tasks?

I invite you to think about these issues or if someone with more experience knows the answers, please enlighten us.
## [4][Roadmap Ruby and Ruby On Rails](https://www.reddit.com/r/rails/comments/huf6mq/roadmap_ruby_and_ruby_on_rails/)
- url: https://www.reddit.com/r/rails/comments/huf6mq/roadmap_ruby_and_ruby_on_rails/
---
You know of a roadmap (similar [https://roadmap.sh/roadmaps](https://roadmap.sh/roadmaps) [https://www.quora.com/What-is-Ruby-On-Rails-framework-learning-road-map](https://www.quora.com/What-is-Ruby-On-Rails-framework-learning-road-map)) but for Ruby or Ruby On Rails, I'd like to create one for Ruby and another for Ruby On Rails, but first I'd like to know if you know of one that already exists and thus extend it or develop it as a learning path that can help novice programmers in their learning. Thanks for your input!
## [5][Rookie here, wanting to start hitting front-end a little stronger since I'm somewhat confident on building a functional back-end.](https://www.reddit.com/r/rails/comments/hucftn/rookie_here_wanting_to_start_hitting_frontend_a/)
- url: https://www.reddit.com/r/rails/comments/hucftn/rookie_here_wanting_to_start_hitting_frontend_a/
---
As  the title says, I have a grand total of maybe 5 months of  self-studying. I'm comfortable with building a functional back-end and  integrating multiple libraries using Ruby on Rails. I'm also pretty  comfortable using HTML &amp; CSS, along with some frameworks such as:  Bulma, SemanticUI, Bootstrap, MaterializeCSS, and (my personal favorite)  TailwindCSS. But the current state of the internet leaves these wanting  a lot, particularly in the field of JS. Which I have absolutely no clue  about. It looks like something straight out of hell to read through. So  I have a mild plan to learn it, it's a simple 3 step familiarization  process.

3 resources: W3schools, FreeCodeCamp, and the TypeScript Deep Dive book.

3 weeks: each covering the aforementioned in their respective order.

3  goals: Understand JS, replicate/personalize JS, build at least 3  features without vicariously googling every other line of code as if my  life depended on it.

Thoughts/recommendations?

Bonus  question: Would this be handled as usual  (via WebPacker:Install:Typescript and just type away in the Apps &gt;  Javascripts &gt; Packs &gt; MyFile.js/ts/tsx)? Or is there some super  annoying and buggy method to incorporating TS front end to the app?

TL;DR

Wanting to improve front-end skills. Unsure of what the best way to do it is. Am I headed the right direction or getting lost in the infinite stream of information?

EDIT:

Fixed wordiness.
## [6][Noob question: How to use module namespaced classes placed under /app](https://www.reddit.com/r/rails/comments/huhxnu/noob_question_how_to_use_module_namespaced/)
- url: https://www.reddit.com/r/rails/comments/huhxnu/noob_question_how_to_use_module_namespaced/
---
Hi guys

I apologize if this is a silly question but so far I haven't found a clear answer yet.

I know that Rails auto loads the folders stored under /app and sure enough, when I make a folder and place a file directly in that folder, it does load that class automatically.

However, I would love to be able to namespace those classes under a module so I can call them in my rails code like so Module::Class.  


Now, when I try to wrap my class in a module or use subfolders in the folder placed under /app, I'm getting errors regarding uninitialized constants.   


I tried adding the path under auto\_load\_paths in application.rb but to no avail.

&amp;#x200B;

Is there anyone that can help me out?  


Thanks in advance!
## [7][Getting Rubocop Lint to work on VS Code for Windows](https://www.reddit.com/r/rails/comments/hufc93/getting_rubocop_lint_to_work_on_vs_code_for/)
- url: https://www.reddit.com/r/rails/comments/hufc93/getting_rubocop_lint_to_work_on_vs_code_for/
---
Hey guys,

Finally got around to trying out WSL for Rails dev on my Windows machine and seems pretty great so far, but I am struggling to get Rubocop/Linter to work with VSCode.

I've read quite a few gists and posts that range from creating a bat that points to the shim your Rubocop is using, to creating a \`.rubocop.yaml\` file and pre-populating a config, to assigning the profile variables as the directions tell you but I'm still having issues.

&amp;#x200B;

At this point I have the bat - Which occasionally VSCode will tell me "it is not executable".

The bats content is:

`u/echo off`

`wsl "/home/user/.rbenv/versions/2.7.1/bin/rubocop $^(echo '%*' ^| sed -e 's^|\\^|/^|g' -e 's^|\^([A-Za-z]\^)\:/\^(.*\^)^|/mnt/\L\1\E/\2^|g'^)`

I then have settings.json (VSCode) set with

&amp;#x200B;

`{"ruby.rubocop.executePath": "/mnt/c/bin/",// "ruby.rubocop.configFilePath": "/mnt/d/Misc Projects (Coding)/Rails/weatherWalk/.rubocop.yml"}`

I copied out the rubocop.yml because it didn't do anything. Right now when I hit Cntrl-Shift-F to format it tells me "An error occurred during auto-correction." which diving into just takes me to the Rubocop extension page.

&amp;#x200B;

One side-note, and I forget where I even read this but 

`which rubocop`

returns 

`/home/user/.rbenv/shims/rubocop`

Which is slightly different than I have in my .bat file.

Is anyone able to have gotten it to work, or know of a VScode alternative for WSL?
## [8][Rails yield and content_for weird behavior, `yield :filter` only work if placed after default yield](https://www.reddit.com/r/rails/comments/hu1bb1/rails_yield_and_content_for_weird_behavior_yield/)
- url: https://www.reddit.com/r/rails/comments/hu1bb1/rails_yield_and_content_for_weird_behavior_yield/
---
# SOLVED

Lets say i have this partial i am trying to render

    #layouts/_subheader.html.erb 
    &lt;div class="subheader"&gt; 
       &lt;%= yield %&gt; 
       &lt;%= yield :filters %&gt; 
    &lt;/div&gt;

when i use this partial in a view like this

    &lt;%= render 'layouts/sub_header' do %&gt;   
       &lt;h2&gt; Content For Yield &lt;/h2&gt; 
       &lt;% content_for :filters do %&gt;      
         &lt;h2&gt; Content for Filters &lt;/h2&gt; 
        &lt;% end %&gt; 
    &lt;% end %&gt;

i am getting the HTML output as

    &lt;div class="subheader"&gt; 
       &lt;h2&gt; Content For Yield &lt;/h2&gt; 
       &lt;h2&gt; Content for Filters &lt;/h2&gt; 
    &lt;/div&gt;

this works as expected, but the problem arises when i change the order of the yieldtags in the partial

instead of the above, if i rewrite the partial as

    #layouts/_subheader.html.erb 
    &lt;div class="subheader"&gt; 
       &lt;%= yield :filters %&gt; 
       &lt;%= yield %&gt; 
    &lt;/div&gt;

i am getting output as

    &lt;div class="subheader"&gt; 
      &lt;h2&gt; Content For Yield &lt;/h2&gt; 
    &lt;/div&gt;

the content\_for :filtersis not being rendered.

what an i doing wrong here ? is this the correct behavior or am i doing something wrong ?

what should i do if i have to make the content of the yield :filtersappear before the plain yield

&amp;#x200B;

EDIT

&amp;#x200B;

thanks to u/[module85](https://www.reddit.com/user/module85/), u/[pacMakaveli90](https://www.reddit.com/user/pacMakaveli90/), u/[Flimsy\_Pomelo](https://www.reddit.com/user/Flimsy_Pomelo/)

ITS NOW WORKING.

&amp;#x200B;

this is what i did.

    #index.html.erb 
    &lt;% content_for :filters do %&gt;         
    &lt;h2&gt; Content for Filters &lt;/h2&gt;  
    &lt;% end %&gt;   
    &lt;%= render 'layouts/sub_header' do %&gt;      
    &lt;h2&gt; Content For Yield &lt;/h2&gt;  
    &lt;% end %&gt;   
    
    
    #layouts/_subheader.html.erb  
    &lt;div class="subheader"&gt;         
       &lt;%= content_for :filters %&gt;         
       &lt;%= yield %&gt; 
    &lt;/div&gt; 

&amp;#x200B;
## [9][Hello guys, trying to build a Rails API but having issues with postgress gem](https://www.reddit.com/r/rails/comments/hu4bba/hello_guys_trying_to_build_a_rails_api_but_having/)
- url: https://www.reddit.com/r/rails/comments/hu4bba/hello_guys_trying_to_build_a_rails_api_but_having/
---
Hello everyone, I want to use postgresql database for my rails API but i am unable to bundle install or even install the pg gem on its own. I know I have postgres downloaded because I see the little elephant icon on the top right hand of my imac screen. 

    sehrishbaloch@Sehrishs-iMac ~ % gem install pg
    Building native extensions. This could take a while...
    ERROR:  Error installing pg:
    ERROR: Failed to build gem native extension.

    current directory: /Users/sehrishbaloch/.rvm/gems/ruby-2.6.1/gems/pg-1.2.3/ext
    /Users/sehrishbaloch/.rvm/rubies/ruby-2.6.1/bin/ruby -I 
    /Users/sehrishbaloch/.rvm/rubies/ruby-2.6.1/lib/ruby/site_ruby/2.6.0 -r ./siteconf20200719- 
    1496-1qwk47i.rb extconf.rb
    checking for pg_config... no
       No pg_config... trying anyway. If building fails, please try again with
       --with-pg-config=/path/to/pg_config
       checking for libpq-fe.h... no
      Can't find the 'libpq-fe.h header
    *** extconf.rb failed ***
    Could not create Makefile due to some reason, probably lack of necessary
     l.   ibraries and/or headers.  Check the mkmf.log file for more details.  You may
    need configuration options.

    Provided configuration options:
	--with-opt-dir
	--with-opt-include
	--without-opt-include=${opt-dir}/include
	--with-opt-lib
	--without-opt-lib=${opt-dir}/lib
	--with-make-prog
	--without-make-prog
	--srcdir=.
	--curdir
	--ruby=/Users/sehrishbaloch/.rvm/rubies/ruby-2.6.1/bin/$(RUBY_BASE_NAME)
	--with-pg
	--without-pg
	--enable-windows-cross
	--disable-windows-cross
	--with-pg-config
	--without-pg-config
	--with-pg_config
	--without-pg_config
	--with-pg-dir
	--without-pg-dir
	--with-pg-include
	--without-pg-include=${pg-dir}/include
	--with-pg-lib
	--without-pg-lib=${pg-dir}/lib

    To see why this extension failed to compile, please check the mkmf.log which can be found here:

      /Users/sehrishbaloch/.rvm/gems/ruby-2.6.1/extensions/x86_64-darwin-18/2.6.0/pg-1.2.3/mkmf.log

    extconf failed, exit code 1

    Gem files will remain installed in /Users/sehrishbaloch/.rvm/gems/ruby-2.6.1/gems/pg-1.2.3 for inspection.
    Results logged to /Users/sehrishbaloch/.rvm/gems/ruby-2.6.1/extensions/x86_64-darwin-18/2.6.0/pg-1.2.3/gem_make.out
    sehrishbaloch@Sehrishs-iMac ~ % 


Thats the error I get
## [10][[video] Integration Testing Best Practices](https://www.reddit.com/r/rails/comments/hu2rl3/video_integration_testing_best_practices/)
- url: https://www.reddit.com/r/rails/comments/hu2rl3/video_integration_testing_best_practices/
---
Watch here: [https://www.semicolonandsons.com/episode/integration-testing-best-practices](https://www.semicolonandsons.com/episode/integration-testing-best-practices)
## [11][Any way to set up emails on a custom domain with Heroku for free?](https://www.reddit.com/r/rails/comments/hu47jt/any_way_to_set_up_emails_on_a_custom_domain_with/)
- url: https://www.reddit.com/r/rails/comments/hu47jt/any_way_to_set_up_emails_on_a_custom_domain_with/
---
I'm looking at options like Mailgun and CloudMailIn and they all need you to be on a paid plan if you want custom domain/email.

Is there any way of setting up my email for free via Gmail or similar?
## [12][SLIM HTML Help With Flexbox](https://www.reddit.com/r/rails/comments/hu1nea/slim_html_help_with_flexbox/)
- url: https://www.reddit.com/r/rails/comments/hu1nea/slim_html_help_with_flexbox/
---
I'm enduring my whole weekend with this, I want to let my current screens ([https://postimg.cc/t1gHLSyF](https://postimg.cc/t1gHLSyF))  like this: [https://postimg.cc/T56ykV4N](https://postimg.cc/T56ykV4N) \- it's no the same model, but I will adapt them all to the same pattern: 

[https://github.com/LeoFragozo/sistema\_loja/blob/master/app/views/categories/index.html.slim](https://github.com/LeoFragozo/sistema_loja/blob/master/app/views/categories/index.html.slim)

can someone help me?
