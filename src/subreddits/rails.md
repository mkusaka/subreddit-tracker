# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/j6qvuh/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/j6qvuh/personal_projects_show_off_your_own_project_andor/
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][Avo - Configuration-based, no-maintenance, extendable Ruby on Rails admin](https://www.reddit.com/r/rails/comments/jc85sc/avo_configurationbased_nomaintenance_extendable/)
- url: https://www.reddit.com/r/rails/comments/jc85sc/avo_configurationbased_nomaintenance_extendable/
---
Hi guys,

Today I'd like to show you [Avo](https://avohq.io), a beautiful next-generation framework that empowers you, the developer, to create fantastic admin panels for your Ruby on Rails apps with the flexibility to fit your needs as you grow.

Out of the box, it has an excellent CRUD interface, ordering, filters, and actions. It even knows how to handle your Active Record model relations. 

It's super easy to configure. There's one configuration file per model and one configuration line of code per field. You can add simple fields like text, textarea, dropdowns, and more complex ones like datetime, badges, loaders, currency, and others. There's even a cool one-liner single or multi-file Active Storage integration 🤯. 

**Avo's mission is to make the job of developers easier and help them and companies move faster.**

Try it in your app and let me know what you think.

Thank you,  
Adrian

[https://avohq.io](https://avohq.io)  
[https://twitter.com/avo\_hq](https://twitter.com/avo_hq)  
[https://github.com/avo-hq/avo](https://github.com/avo-hq/avo)   
[https://discord.gg/pkTF6y8](https://discord.gg/pkTF6y8)
## [4][Rails 6 with Bootstrap and Webpacker: Quick guide](https://www.reddit.com/r/rails/comments/jbv43i/rails_6_with_bootstrap_and_webpacker_quick_guide/)
- url: https://www.reddit.com/r/rails/comments/jbv43i/rails_6_with_bootstrap_and_webpacker_quick_guide/
---
I guess this exact tutorial topic is already a "mauvais ton", but after reading 10+ guides on "How to install bootstrap on Rails 6" I wrote a quick guide: [https://blog.corsego.com/2020/10/rails-6-with-bootstrap-and-webpacker.html](https://blog.corsego.com/2020/10/rails-6-with-bootstrap-and-webpacker.html)

I hope you'll find it useful!
## [5][Rails help](https://www.reddit.com/r/rails/comments/jc95ti/rails_help/)
- url: https://www.reddit.com/r/rails/comments/jc95ti/rails_help/
---
 hey guys, so i have this and i am trying to get each position(entertainment, src president) have its own page. lets say if i vote for the candidates of entertainment position it directs me to src president to vote and so on. how do i do this please? 

&amp;#x200B;

https://preview.redd.it/g0sljvp7dgt51.png?width=1366&amp;format=png&amp;auto=webp&amp;s=849b2c46173b6aa28f04e9b8279cce0aa042ebda

&gt;&lt;table&gt;  
&gt;  
&gt;  &lt;tbody&gt;  
&gt;  
&gt;&lt;%= form\_for u/election do |f| %&gt;  
&gt;  
&gt;  
&gt;  
&gt;&lt;% Position.includes(:candidates).order(:name).each do |position| %&gt;  
&gt;  
&gt;  
&gt;  
&gt;  &lt;tr&gt;  
&gt;  
&gt;&lt;th colspan="5"&gt;&lt;%= position.name %&gt;&lt;/th&gt;  
&gt;  
&gt;  &lt;/tr&gt;  
&gt;  
&gt;  &lt;% position.candidates.each do |candidate| %&gt;  
&gt;  
&gt;  &lt;tr&gt;  
&gt;  
&gt;&lt;td&gt;&lt;%= image\_tag(candidate.image, :size =&gt; '50x50') if candidate.image.attached? %&gt;&lt;/td&gt;  
&gt;  
&gt;&lt;td&gt;&lt;%= candidate.name %&gt;&lt;/td&gt;  
&gt;  
&gt;&lt;td&gt;&lt;%= candidate.info %&gt;&lt;/td&gt;  
&gt;  
&gt;&lt;td&gt;

&lt;%= f.check\_box :c\_votes %&gt;

&amp;#x200B;

&gt;&lt;/td&gt;  
&gt;  
&gt;  &lt;/tr&gt;  
&gt;  
&gt;  &lt;% end %&gt;  
&gt;  
&gt;  &lt;td&gt;  &lt;%= f.submit "Submit"  %&gt;&lt;/td&gt;  
&gt;  
&gt;&lt;% end %&gt;  
&gt;  
&gt;  &lt;/tbody&gt;  
&gt;  
&gt;&lt;/table&gt;  
&gt;  
&gt;  
&gt;  
&gt; &lt;% end %&gt;

here is the form please
## [6][Find unique places from Scraping website data](https://www.reddit.com/r/rails/comments/jc8hsd/find_unique_places_from_scraping_website_data/)
- url: https://www.reddit.com/r/rails/comments/jc8hsd/find_unique_places_from_scraping_website_data/
---
I have lists of restaurants in CSV file from 2 websites and also from google, but all files have some restaurants which are not present in other files.   
i want to combine them all and want a unique from them but all files have their own id, 2 places can have a same name And even lat , lng is different for a same place &amp; their address may have words not in a same pattern.   


Any senior can please help me ?
## [7][How to significantly archive and compress a given tmp Rails directory contents?](https://www.reddit.com/r/rails/comments/jc47cr/how_to_significantly_archive_and_compress_a_given/)
- url: https://www.reddit.com/r/rails/comments/jc47cr/how_to_significantly_archive_and_compress_a_given/
---
My workflow is as follows:

1. User requests to have the system send many PDF files (in the hundreds) to a given email address. I use Sidekiq to handle this job.
2. I generate these pdf documents and save them to a random-named directory in tmp folder (eg. tmp/pdfs/2e3b8f6f-81e6-40b1-9cdc-d1bb0a0ee519/)
3. I use the gem RubyZip and the example given to zip the `tmp/pdfs/2e3b8f6f-81e6-40b1-9cdc-d1bb0a0ee519/` directory into a zip file (eg. [thefiles.zip](https://thefiles.zip))
4. I send [thefiles.zip](https://thefiles.zip) to the user via email.

Problem: RubyZip does not significantly compress this archive. If I request that 24 PDF files equaling 12 MB be archived, the archive is 10MB for eg.

How can I do this?
## [8][I curated all the remote job openings from Hacker News who is hiring - October](https://www.reddit.com/r/rails/comments/jbv2j8/i_curated_all_the_remote_job_openings_from_hacker/)
- url: https://www.reddit.com/r/rails/comments/jbv2j8/i_curated_all_the_remote_job_openings_from_hacker/
---
Here I would like to share more than 480 remote jobs that I've curated from Hacker News Who is hiring thread. All these are 100% remote jobs not just allowed to work from home during COVID-19. These are 100% remote jobs and will continue to follow that after the covid.

https://remoteleaf.com/whoishiring.   

Note: Select "ruby" in the skills filter to view all Ruby/Rails jobs

✅ 100% remote full-time jobs.    
✅ Each and every job is manually curated and verified. Spent more than 14 hours for this.
## [9][Creating multiple user types in rails using devise?](https://www.reddit.com/r/rails/comments/jbp4tb/creating_multiple_user_types_in_rails_using_devise/)
- url: https://www.reddit.com/r/rails/comments/jbp4tb/creating_multiple_user_types_in_rails_using_devise/
---
Hi there, I want to have two different types of users (buyers, sellers) on my application. Should I create two types of users using devise to accomplish this? Or is there a more streamlined approach?
## [10][Upgrading a rails app from 4.2 to 5.0...i need help](https://www.reddit.com/r/rails/comments/jbya8p/upgrading_a_rails_app_from_42_to_50i_need_help/)
- url: https://www.reddit.com/r/rails/comments/jbya8p/upgrading_a_rails_app_from_42_to_50i_need_help/
---
O.k. so I finally got a rails server to load with Rails 5.0.0 and ruby 2.5.8 (I'm upgrading from ruby 2.2.4 and rails 4.2.0).

I don't have any sll issues.

I'm now at the "ActiveRecord::PendingMigrationError" level.

Basically I haven't created the tables in my db (the db is already created in postgres).

When I run the bin/rails db:migrate RAILS_ENV=development I get the following:

&gt;PG::DatatypeMismatch: ERROR:  column "root_comment" is of type boolean but default expression is of type integer
&gt;HINT:  You will need to rewrite or cast the expression.

I've read a ton, and it seems rails 5 changed both migration scripts and how to define columns in postgres.

Do I need to label my migration script with something like [4.2 or 5.0] or do I actually need to change the type (further screwing up my app)?

Thanks in advance.
## [11][Modal age confirmation? Or just a view for it?](https://www.reddit.com/r/rails/comments/jbtypu/modal_age_confirmation_or_just_a_view_for_it/)
- url: https://www.reddit.com/r/rails/comments/jbtypu/modal_age_confirmation_or_just_a_view_for_it/
---
Hi folks. I'm making a web app for my web dev portfolio and it has to do with the cannabis industry. So I'd like to add an age confirmation to the landing page. I was thinking it could just be a modal that pops up and just asks if you're 21 or over, or has you enter your DOB, and then lets you continue to the site and browse through it even if you're not logged in. Or should there just be a view dedicated to that? Either way it'll need a controller, right? 

Would this be a practical way in Rails to do this? The landing page will just have the log in or sign up links in the nav, and then there's an "enter" button in the middle of the page that would then prompt that modal. I'm just looking for advice or to be pointed in the right direction as I'm honestly not finding much in my research that would help me with this exact issue. A lot of what's coming up is deletion confirmation modals, or some other things that don't exactly help me. 

I'm still very new to Rails. I only just graduated boot camp a while ago so I'm building little apps for my portfolio and gaining that experience. Any help would be appreciated!
## [12][CanCanCan + rolify... is it necessary?](https://www.reddit.com/r/rails/comments/jbryal/cancancan_rolify_is_it_necessary/)
- url: https://www.reddit.com/r/rails/comments/jbryal/cancancan_rolify_is_it_necessary/
---
Hi guys, in my website now I have only two roles. User &amp; Staff.

But Staff is not a "real" role. In the DB I can set Staff true/false.

Now I want to improve my website and add 5 kind of rules: User, Verified User, Banned User, Mod and Staff.

Tecnically I can set their "power" in the `app/policies` area (using, for example  `user.try(:mod?)` )

A developer suggested me to use CanCanCan and rolify.

Is it really necessary? Just to add 4 roles? I don't want to add gems just to add one or two features...

p.s. I already use devise
