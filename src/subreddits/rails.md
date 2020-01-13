# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/em8qtp/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/em8qtp/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][webpacker and svg vue components](https://www.reddit.com/r/rails/comments/enxxer/webpacker_and_svg_vue_components/)
- url: https://www.reddit.com/r/rails/comments/enxxer/webpacker_and_svg_vue_components/
---
Anyone have luck with vue, webpacker, and svg importing like components?

I'm getting this warning: Invalid Component definition /packs/media/images/icons/chevron-down-fa5a06b100afba790b995b0eb306a631.svg

I've tried the following:

1. removing \`- svg\` from extensions in \`webpacker.yml\`
2. adding vue-svg-loader
3. Adding a new loader in \`webpack/loaders/vue-svg-loader.js\`

&amp;#8203;

    module.exports = {
        test: /.svg$/, 
        use: [{ loader: 'vue-svg-loader' }] 
    }

I added the loader like I'm suppose too in the \`webpack/environment.js\`

&gt;const vueSvgLoader = require('./loaders/vue-svg-loader')environment.loaders.append('vue-svg-loader', vueSvgLoader)
## [3][Offline mode for mobile app with rails backend](https://www.reddit.com/r/rails/comments/eo3nrm/offline_mode_for_mobile_app_with_rails_backend/)
- url: https://www.reddit.com/r/rails/comments/eo3nrm/offline_mode_for_mobile_app_with_rails_backend/
---
Hey there!

I just started to work in a company developing a rails backend app with companion mobile app (ios). This app needs to have an offline mode available which goes down to downloading all the data from the backend when internet connection is available.

How it currently handled is: scheduled worker produces one huge json object with all the needed data (no pagination, no split to rest resources, simply everything in one big json file) and the second worker generates zip files with all the images/assets.

As it works relatively well, i have this irritating feeling it's not how it should be done and i'm pretty sure someone came up with something better than that. Producing such json is heavy and long, data is not updated live in cause of caching (it can't be simply based on caching and cache invalidation as producing that json takes up to 2 minutes, we can't handle that live during the request and cache it). Producing zip files with assets is also pretty bad as those zip files weights A LOT from mobile app perspective and downloading them takes tons of time. Overall, it all seems sketchy both from client and developer perspective.

Is there anything i should read more about in that matter? Any alternative approaches or patterns? Thanks in advance for any clues!
## [4][Update Redis in AWS Cloud9 IDE](https://www.reddit.com/r/rails/comments/enpzbl/update_redis_in_aws_cloud9_ide/)
- url: https://www.reddit.com/r/rails/comments/enpzbl/update_redis_in_aws_cloud9_ide/
---
`$ sudo yum install redis`

\*already did this - this gives me:

Package redis-3.2.12-2.el6.x86\_64 already installed and latest version

When I run Sidekiq: You are connecting to Redis v3.2.12, Sidekiq requires Redis v4.0.0 or greater

I can't figure out how to get this to update. $ bundle update/install shows Using redis 4.1.3 (as does the gemfile.lock)

I got it to update on heroku by eliminating redis-to-go and adding Heroku Redis add-on - still not right in development though.

**\*Solved** \- I ended up restarting my workspace with Ubuntu instead of Amazon Linux - It had redis at Redis server v=4.0.9 which got me what I needed to run in both dev and production.
## [5][Reality Series on Webdev](https://www.reddit.com/r/rails/comments/enzawz/reality_series_on_webdev/)
- url: https://www.reddit.com/r/rails/comments/enzawz/reality_series_on_webdev/
---
Hi

I'm doing a video series on creating a web application using Ruby on Rails. It's not meant to be a tutorial but more of a reality TV show on how one would go about creating an application. Mostly Ruby with some javascript stuff.

I was just curious if any of you know anything similar (webdev specifically) or anyone doing the same thing (coding each step of the way w/ mistakes). I'm looking for similar content where the mistakes are discovered on the fly and how one would go about fixing it.

Playlist link here:

[https://www.youtube.com/playlist?list=PL2-7U6BzddIZ35bJdCFx6RZ-QR8n\_JD82](https://www.youtube.com/playlist?list=PL2-7U6BzddIZ35bJdCFx6RZ-QR8n_JD82)

Videos ongoing. Trying to do a video once a week or if I have time. For this one, I'm trying to build a book keeping system.
## [6][Build a Blog Title Generator - Ruby Tutorial](https://www.reddit.com/r/rails/comments/ennekd/build_a_blog_title_generator_ruby_tutorial/)
- url: https://www.reddit.com/r/rails/comments/ennekd/build_a_blog_title_generator_ruby_tutorial/
---
Hi guys, I normally post a lot of Ruby on Rails builds on YouTube and I've shared a couple here in the past. This time I've decided to build something in pure Ruby. It's very beginner friendly but might be useful to those who dived straight into Rails without spending a ton of time learning the Ruby language. In the video I build a Ruby file that generates viral blog title ideas and returns them back to the user, the titles are based on the keyword that the user inputs via Terminal.

This could be easily adapted and used within a Rails app. Hopefully it's useful to some of you guys.

[https://www.youtube.com/watch?v=BC7VaxGkihY](https://www.youtube.com/watch?v=BC7VaxGkihY)

I'm planning to post some more Ruby content so would be open to suggestions about content ideas.
## [7][Do I need to deny access if the nginx root path is the public folder?](https://www.reddit.com/r/rails/comments/enlcp9/do_i_need_to_deny_access_if_the_nginx_root_path/)
- url: https://www.reddit.com/r/rails/comments/enlcp9/do_i_need_to_deny_access_if_the_nginx_root_path/
---
Do I need to configure anything to prevent people from accessing my ruby files in a rails app, if the nginx config root is the public folder? If I just set the root as the public folder path, there's no way the app source code can be accessed through nginx right?
## [8][Production Server setup — Vultr Tutorial](https://www.reddit.com/r/rails/comments/entvbl/production_server_setup_vultr_tutorial/)
- url: https://www.reddit.com/r/rails/comments/entvbl/production_server_setup_vultr_tutorial/
---

This series is a guide to deploy a Ruby on Rails application via Capistrano to a VPS (virtual private server):

[Part 1 : Deploy a server](https://itnext.io/deploy-rails-with-capistrano-tutorial-2020-part-1-fa732e0db144?source=friends_link&amp;sk=80430970822b0dcdaa1a472804da6711)

[Part 2 — Server setup](https://itnext.io/production-server-setup-part-2-vultr-tutorial-87a6e19f9a2f?source=friends_link&amp;sk=ac823106b2d3fd321ccfa4e4e2e3397c)
## [9][How do you handle image modals?](https://www.reddit.com/r/rails/comments/enjr2i/how_do_you_handle_image_modals/)
- url: https://www.reddit.com/r/rails/comments/enjr2i/how_do_you_handle_image_modals/
---
Hi rails community.

&amp;#x200B;

I am making a simple rails blog with the goal of it being a personal resume site. Posts in the blog are handled with a simple CRUD posts generator like so:

&amp;#x200B;

`&lt;div class="projectcolumn"&gt; &lt;% posts.each do |post| %&gt;`

`&lt;div class="blurbox"&gt; &lt;h1&gt;&lt;%= post.title %&gt;&lt;/h1&gt; &lt;p&gt;&lt;%= post.content %&gt;&lt;/p&gt;`

`&lt;div class="postimage" id="myImg"&gt; &lt;%= image_tag(post.image)%&gt; &lt;/div&gt; &lt;/div&gt; &lt;% end %&gt;`

I use active storage to store images with each post. I had the annoying idea that I want images on the site to be image modals, as so:

`&lt;script&gt;// create references to the modal...`

`var modal = document.getElementById('myModal');`

`var images = document.getElementsByClassName('postimage');`

`var modalImg = document.getElementById("img01");`

`var captionText = document.getElementById("caption");`

`// Go through all of the images with our custom class`

`for (var i = 0; i &lt; images.length; i++) {`

`var img = images[i];`

`img.onclick = function(evt)`

`{ modal.style.display = "block"; modalImg.src = this.src; }`

`}`

`var span = document.getElementsByClassName("close")[0]; span.onclick = function() { modal.style.display = "none"; }`

`&lt;/script&gt;`

The modal appears, however the image contents returns undefined. However testing with "static" images linked by URL works as expected. I thought that by adding the javascript at the bottom of the html, it would load after the ruby code had executed. Is there a problem in this line of thinking?

Is there a "rails way" of doing image modals that you prefer?
## [10][Why does Active Record Migrations not set default value for created_at in the migration files?](https://www.reddit.com/r/rails/comments/enkkww/why_does_active_record_migrations_not_set_default/)
- url: https://www.reddit.com/r/rails/comments/enkkww/why_does_active_record_migrations_not_set_default/
---
I just came across this issue about the \`insert\_all\` method throwing ActiveRecord::NotNullViolation on created\_at and wondered why it is not set in the migration files by default as I think this issue didn't come up if it had been set.

The issue I'm referring: \[insert\_all throws \`ActiveRecord::NotNullViolation\` for timestamps\]([https://github.com/rails/rails/issues/35493](https://github.com/rails/rails/issues/35493))

A merged PR to support RDB specific functions like NOW() as a column default value: \[Add \`:expression\` option support on the schema default\]([https://github.com/rails/rails/pull/20005](https://github.com/rails/rails/pull/20005))
## [11][Web Based Accounting System One to Many UI Flow](https://www.reddit.com/r/rails/comments/enfhf5/web_based_accounting_system_one_to_many_ui_flow/)
- url: https://www.reddit.com/r/rails/comments/enfhf5/web_based_accounting_system_one_to_many_ui_flow/
---
Hi

I've been doing a video series on how I would go about in building a web based accounting system. It's not a tutorial type but more of a learning process in exploring how to improve as a developer. I got to the part with a typical one to many relationship (to mode the T-Account entries in accounting using models AccountingEntry has\_many JournalEntry). Looking for advice as to how to deal with modifying / managing a one to many instance in a single module. Thinking of several aspects such as managing the main model (AccountingEntry) and having submodules to add in the has\_many records via an API layer similar to the current one in the video (see link below). I know we can have them as separate modules / crud sets but was looking into a richer user experience instead of having to load multiple pages. Anyway, just looking for opinions on it. If the structure in the current video is helpful to others, hoping to look into improving it for people who'd encounter a similar usecase.

Regards,

Video here:

[https://www.youtube.com/watch?v=ITrhSwV4Gqs&amp;feature=youtu.be](https://www.youtube.com/watch?v=ITrhSwV4Gqs&amp;feature=youtu.be)
