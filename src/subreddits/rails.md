# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/em8qtp/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/em8qtp/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][How did you turn your Rails app into an Android and iOS app?](https://www.reddit.com/r/rails/comments/eoiq8q/how_did_you_turn_your_rails_app_into_an_android/)
- url: https://www.reddit.com/r/rails/comments/eoiq8q/how_did_you_turn_your_rails_app_into_an_android/
---
Hello!

How did you turn your Rails app into an Android and iOS app? What were the pros/cons? Was it easy to get accepted into App Store and Google Play?

Im very new to the game but from what I understand, for a Rails app to be turned into mobile apps it would have to use PhoneGap to act as a browser basically? And then set up an API to communicate with the user's mobile phone and the Rails server?

Thanks!
## [3][Anyone willing to help an old man troubleshoot my rails 6 app?](https://www.reddit.com/r/rails/comments/eobzel/anyone_willing_to_help_an_old_man_troubleshoot_my/)
- url: https://www.reddit.com/r/rails/comments/eobzel/anyone_willing_to_help_an_old_man_troubleshoot_my/
---
I'm 99 per cent sure the issue is with how I set up webpacker, despite my best efforts. User can't sign in or log out. Dropdown menus don't function and modal won't fire.

Not sure what I did! I'm really an old man (77), volunteer webmaster for my church website. Is someone willing to clone it and give it a spin? [github.com/ThomasConnolly/paul](https://github.com/ThomasConnolly/paul)

The error message I get when trying to log out is "ActiveRecord::RecordNotFound - Couldn't find User with 'id'=sign\_out:"

The dropdowns &amp; modal just don't filre; no error message.
## [4][What to keep in mind if I change an existing deployed Rails app to start using Docker?](https://www.reddit.com/r/rails/comments/eojyk1/what_to_keep_in_mind_if_i_change_an_existing/)
- url: https://www.reddit.com/r/rails/comments/eojyk1/what_to_keep_in_mind_if_i_change_an_existing/
---
I've got an existing Rails app that's been running for a year or so. It's running Rails 5.2 and I deploy it to AWS using Elastic Beanstalk, with the database living in RDS. I was learning Rails (and a lot of other web dev stuff) as I made it, so now that I look back, I probably could have gotten away with a cheaper DigitalOcean instance or something similar. As is the way of things, there's enough of the app hooked into the AWS ecosystem that pulling it out would be pretty daunting.

I was attempting to upgrade the app to Rails 6, but my existing instance was running Ruby 2.3, so I tried to make a new instance and do a blue-green deployment, but after quite a lot of trial-and-error, I was never able to get Rails 6 to work on Elastic Beanstalk. I tried searching around and trying arcane EB configuration tweaks, but I couldn't ever deploy it successfully. I'm guessing that webpacker and/or yarn weren't installing correctly, and those are two things I have only the most basic knowledge of.

It occurred to me that having to tweak Elastic Beanstalk this much was sort of defeating the point of using a deployment system that's supposed to be simple, so I thought it was time I bit the bullet and start experimenting with Docker.

I've been able to get development and test environments running on my local systems with docker-compose using a very basic [Dockerfile](https://github.com/HeadBeeGuy/alt-activities/blob/rails-6-docker/Dockerfile) and [docker-compose](https://github.com/HeadBeeGuy/alt-activities/blob/rails-6-docker/docker-compose.yml). I'm realizing that I can't just push these up to a webserver and expect it to take the place of the old instance, though. At the very least, I'd have to tell it that it's running in production, tell it how to talk to the production database, and make sure it's using a web server appropriate for production instead of development. I've got a couple newbie questions that I wanted to ask anyone who uses Docker in production:

* Do I specify separate Dockerfiles or docker-compose.yml files for production? Is there a way to test these before deploying them?
* Do I need the "db" image for production? Since the database lives in RDS, I imagine Rails will just talk to it and it doesn't need to be handled by Docker?
* From my research so far, it looks like I can use Elastic Beanstalk, ECS, or Fargate to deploy Rails apps. Is there one that's generally preferable?
* Should I use Puma? Passenger? nginx? Several in combination?

The app has been able to handle its traffic thus far using only one instance, so for now I don't think I need to worry about auto-scaling. Raw performance hasn't been as important as being to iterate on features rapidly.

Sorry this is scattered! Any advice would be welcome!
## [5][Cache Crispies - Fast, Flexible Rails Serializer](https://www.reddit.com/r/rails/comments/eoes6h/cache_crispies_fast_flexible_rails_serializer/)
- url: https://www.reddit.com/r/rails/comments/eoes6h/cache_crispies_fast_flexible_rails_serializer/
---
Picking a method of doing JSON serialization in Rails has not been an easy decision as of late. Especially if you're not able to break your APIs by moving to a JSON API structure. And trying to mix in a caching strategy, certainly doesn't help. That's the problem the new Cache Crispies gem was written to fix.  


[https://codenoble.com/blog/introducing-cache-crispies/](https://codenoble.com/blog/introducing-cache-crispies/)  
[https://github.com/codenoble/cache-crispies](https://github.com/codenoble/cache-crispies)
## [6][The best way to pass join table has_many ids to frontend?](https://www.reddit.com/r/rails/comments/eocqod/the_best_way_to_pass_join_table_has_many_ids_to/)
- url: https://www.reddit.com/r/rails/comments/eocqod/the_best_way_to_pass_join_table_has_many_ids_to/
---
I'm writing a Rails/React app and need to have a: 

    class Outage &lt; ApplicationRecord 
        has_many :service_outages
        has_many :services, through: :service_outages
    end
    
    class Service &lt; ApplicationRecord 
        has_many :service_outages
        has_many :outages, through: :services_outages
    end
    
    class ServiceOutages &lt; ApplicationRecord
        belongs_to :service 
        belongs_to :outage
    end

with the \`ServiceOutage\` join table having \`service\_id\` and \`outage\_id\` as attributes. 

&amp;#x200B;

What is the best way to pass this data to my frontend so I can find appropriate relationships without having to create a separate controller and routes?
## [7][Offline mode for mobile app with rails backend](https://www.reddit.com/r/rails/comments/eo3nrm/offline_mode_for_mobile_app_with_rails_backend/)
- url: https://www.reddit.com/r/rails/comments/eo3nrm/offline_mode_for_mobile_app_with_rails_backend/
---
Hey there!

I just started to work in a company developing a rails backend app with companion mobile app (ios). This app needs to have an offline mode available which goes down to downloading all the data from the backend when internet connection is available.

How it currently handled is: scheduled worker produces one huge json object with all the needed data (no pagination, no split to rest resources, simply everything in one big json file) and the second worker generates zip files with all the images/assets.

As it works relatively well, i have this irritating feeling it's not how it should be done and i'm pretty sure someone came up with something better than that. Producing such json is heavy and long, data is not updated live in cause of caching (it can't be simply based on caching and cache invalidation as producing that json takes up to 2 minutes, we can't handle that live during the request and cache it). Producing zip files with assets is also pretty bad as those zip files weights A LOT from mobile app perspective and downloading them takes tons of time. Overall, it all seems sketchy both from client and developer perspective.

Is there anything i should read more about in that matter? Any alternative approaches or patterns? Thanks in advance for any clues!
## [8][How should I include CSS and JS without Sprockets](https://www.reddit.com/r/rails/comments/eoambe/how_should_i_include_css_and_js_without_sprockets/)
- url: https://www.reddit.com/r/rails/comments/eoambe/how_should_i_include_css_and_js_without_sprockets/
---
Hi, I created a new app with the --skip-sprockets flag. When I run the app, it shows a resource not found error in the JavaScript console for application.css. Since i'm not using sprockets, I should put all my CSS and JS in the public folder? The stylesheets folder in the app folder is not used at all now?
## [9][Opensource Project using liquid gem/templates?](https://www.reddit.com/r/rails/comments/eo86y3/opensource_project_using_liquid_gemtemplates/)
- url: https://www.reddit.com/r/rails/comments/eo86y3/opensource_project_using_liquid_gemtemplates/
---
Does anyone know of an open source project that uses liquid gem for creating liquid templates like Shopify does for its themes?
## [10][webpacker and svg vue components](https://www.reddit.com/r/rails/comments/enxxer/webpacker_and_svg_vue_components/)
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
## [11][Boston Startup Hiring Software Engineers with Ember/Rails focus [$70-140k DOE] [NO REMOTE]](https://www.reddit.com/r/rails/comments/eo7ygh/boston_startup_hiring_software_engineers_with/)
- url: https://www.reddit.com/r/rails/comments/eo7ygh/boston_startup_hiring_software_engineers_with/
---
[Mapdwell](https://mapdwell.com/) is hiring for a few engineering positions (Frontend Engineer, Full-Stack Engineer, and Senior Full-Stack Engineer), with a focus on Ember and Ruby on Rails.

The positions involve working across our entire stack and product suite, with a focus on building beautiful, rich front-end web applications. There will also be involvement with algorithm design, machine learning models, data processing, geospatial databases, back-end services, APIs and deployment.

**About Mapdwell**  
Mapdwell, an MIT spinoff, is a software and data intelligence startup that provides advanced tools for energy decision-making. We level the field for consumers by providing intuitive and actionable tools and leverage user engagement into powerful data intelligence for industry. Mapdwell uses hard science and data intelligence to uncover opportunities that align good business and best practices with a more resilient, healthier and more sustainable future by driving demand and empowering industry. Learn more at [mapdwell.com](http://mapdwell.com/).

See all openings and apply here: [http://jobs.mapdwell.com/](http://jobs.mapdwell.com/)
