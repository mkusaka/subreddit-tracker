# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gvu083/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gvu083/personal_projects_show_off_your_own_project_andor/
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
## [2][Awesome-Ruby: A Curated Collection of 851 Ruby Links and Resources](https://www.reddit.com/r/rails/comments/gwv338/awesomeruby_a_curated_collection_of_851_ruby/)
- url: https://www.reddit.com/r/rails/comments/gwv338/awesomeruby_a_curated_collection_of_851_ruby/
---
Check this on github. It really helps. Awesome-Ruby: https://github.com/markets/awesome-ruby#readme
## [3][Stimulus course in context of Rails?](https://www.reddit.com/r/rails/comments/gx2ca6/stimulus_course_in_context_of_rails/)
- url: https://www.reddit.com/r/rails/comments/gx2ca6/stimulus_course_in_context_of_rails/
---
Are you aware of any online courses or resources that teach how to use Stimulus in the context of a Rails 6 site? I know Rails and JS in general, and I would like to pick up Stimulus, and learn best practices. Thank you!
## [4][Webpack struggles - need some help with an npm package](https://www.reddit.com/r/rails/comments/gx4geq/webpack_struggles_need_some_help_with_an_npm/)
- url: https://www.reddit.com/r/rails/comments/gx4geq/webpack_struggles_need_some_help_with_an_npm/
---
Hello all, I have a Rails 6 project where I recently started using Stimulus and Stimulus Reflex, which forced me to go down the webpack road a bit.  I am now starting to migrate everything else to webpack from the asset pipeline and running into a little bit of a wall...

I am using Bulma for this project and have successfully got that working as a webpack/yarn package (previously I was downloading the entire Bulma library and storing under the vendor folder of my project).

I was also using an extension called "Bulma Toast" which appears to have an npm package as well, but for the life of me I cannot get it working in the new webpack setup, and hoping someone can steer me in the right direction.  No matter what I say I get an error in the rails console complaining that bulma-toast cannot be found, but if I fix that then I get javascript console errors in my browser saying "bulmaToast is undefined"

&amp;#x200B;

Posting relevant files below:

**package.json**

    {
      "name": "asdf",
      "private": true,
      "dependencies": {
        "@creativebulma/bulma-tooltip": "^1.2.0",
        "@rails/webpacker": "4.2.2",
        "bulma": "^0.8.2",
        "bulma-toast": "^2.0.1",
        "stimulus": "^1.1.1",
        "stimulus_reflex": "^3.2.1",
        "url-loader": "^4.1.0"
      },
      "devDependencies": {
        "webpack-dev-server": "^3.10.0"
      }
    }

**app/javascript/packs/application.js**

    import "controllers"
    import "../application/stylesheets/application"
    import bulmaToast from 'bulma-toast';

&amp;#x200B;

**app/views/layouts/application.html.erb (excerpt)**

    &lt;head&gt;
        &lt;%= stylesheet_pack_tag 'application', media: 'all', 'data-turbolinks-track': 'reload' %&gt;
        &lt;%= javascript_pack_tag 'application', 'data-turbolinks-track': 'reload', async: true %&gt;
    &lt;/head&gt;
    ...
    &lt;main class="site-content"&gt;
        &lt;section class="section"&gt;
          &lt;div class="container"&gt;
            &lt;%= yield %&gt;
          &lt;/div&gt;
        
          &lt;% flash.each do |message_type, message| %&gt;
            &lt;%= javascript_tag "bulmaToast.toast({ message: '#{message}',
                                                  position: 'top-center',
                                                  type: 'is-#{message_type}',
                                                  animate: { in: 'fadeIn', out: 'fadeOut' },
                                                  duration: 2500 });" %&gt;
          &lt;% end %&gt;
        &lt;/section&gt;
    &lt;/main&gt;

&amp;#x200B;
## [5][Digest form input](https://www.reddit.com/r/rails/comments/gwwq00/digest_form_input/)
- url: https://www.reddit.com/r/rails/comments/gwwq00/digest_form_input/
---
So... I added username &amp; display name strings to my user model via rails g migration AddBlahToBlah username:string display\_name:string

Then I added the digest to the registrations controller.

Then I added the fields to the form/update view (using devise).

Looks kinda like this:  
&lt;%= f.label :username, "Username" %&gt;

&lt;%= f.text\_field :username, placeholder: "Username" %&gt; 

With some styling in between.   


I loaded up the edit account view and everything worked perfectly. However, none of the changes are actually saving. 

So I went into the rails console to confirm this, and sure enough the values are still nil. Something is happening in the database, but none of the inputs are going through. What do?
## [6][Bootstrap-filestyle](https://www.reddit.com/r/rails/comments/gwwdzv/bootstrapfilestyle/)
- url: https://www.reddit.com/r/rails/comments/gwwdzv/bootstrapfilestyle/
---
Has anyone used this before?

[bootstrap-filestyle](https://markusslima.github.io/bootstrap-filestyle/)

I can kind of get it to work.. I added the classes to the file_field in my view file and when I refresh the page my file input is styled.  However if I leave and come back to the page it’s gone, but if I refresh it comes back. I’m on mobile at my daughters dance so I can’t easily copy my code at this moment but it was on my mind and thought I’d see if this community had any experience with it.
## [7][Recommendations for courses/books/tutorials](https://www.reddit.com/r/rails/comments/gwmgsu/recommendations_for_coursesbookstutorials/)
- url: https://www.reddit.com/r/rails/comments/gwmgsu/recommendations_for_coursesbookstutorials/
---
I am trying to determine the best path to salvage a legacy RAILS website that has been offline since the host dropped RAILS support. One thought is to containerize it and host it on AWS. But, I haven't done web development in years, and little experience with RAILS. I went through the free part of Michael Hartl's Ruby on RAILS tutorial and it seems good, although I hesitate to do a subscription service as opposed to a pay once and download e-book. Can anyone either provide a recommendation for Michael Hartl's tutorial or suggest other resources for coming up to speed on RAILS? Also, if anyone has dockerized a RAILS website, I would appreciate any perspectives on that as well.

Cheers!
## [8][First Rails project in Ubuntu terminal - Having Troubles with Server Startup](https://www.reddit.com/r/rails/comments/gwncxb/first_rails_project_in_ubuntu_terminal_having/)
- url: https://www.reddit.com/r/rails/comments/gwncxb/first_rails_project_in_ubuntu_terminal_having/
---
Hello there. I have come to seek help here again regarding version control for my ruby apps. I found a tutorial on how to set up a rails project in Ubuntu starting with installing RVM and the latest Ruby/Rails versions. I installed the bundler and stuff. But when trying to start the server, it's like I missed something. The webpack install? So it tells me to run "rails webpack:install" and I"m getting this error below: 

`rails aborted!`

`ArgumentError: Malformed version number string 0.32+git`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/webpacker-4.2.2/lib/tasks/webpacker/check_yarn.rake:12:in \`block (2 levels) in &lt;main&gt;'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/railties-6.0.3.1/lib/rails/commands/rake/rake_command.rb:23:in \`block in perform'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/railties-6.0.3.1/lib/rails/commands/rake/rake_command.rb:20:in \`perform'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/railties-6.0.3.1/lib/rails/command.rb:48:in \`invoke'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/railties-6.0.3.1/lib/rails/commands.rb:18:in \`&lt;main&gt;'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:23:in \`require'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:23:in \`block in require_with_bootsnap_lfi'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/loaded_features_index.rb:92:in \`register'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:22:in \`require_with_bootsnap_lfi'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/bootsnap-1.4.6/lib/bootsnap/load_path_cache/core_ext/kernel_require.rb:31:in \`require'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/activesupport-6.0.3.1/lib/active_support/dependencies.rb:324:in \`block in require'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/activesupport-6.0.3.1/lib/active_support/dependencies.rb:291:in \`load_dependency'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/activesupport-6.0.3.1/lib/active_support/dependencies.rb:324:in \`require'`

`/mnt/c/dev/src/my-portfolio/bin/rails:9:in \`&lt;top (required)&gt;'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/spring-2.1.0/lib/spring/client/rails.rb:28:in \`load'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/spring-2.1.0/lib/spring/client/rails.rb:28:in \`call'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/spring-2.1.0/lib/spring/client/command.rb:7:in \`call'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/spring-2.1.0/lib/spring/client.rb:30:in \`run'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/spring-2.1.0/bin/spring:49:in \`&lt;top (required)&gt;'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/spring-2.1.0/lib/spring/binstub.rb:11:in \`load'`

`/home/spartan-626/.rvm/gems/ruby-2.7.1/gems/spring-2.1.0/lib/spring/binstub.rb:11:in \`&lt;top (required)&gt;'`

`/mnt/c/dev/src/my-portfolio/bin/spring:15:in \`require'`

`/mnt/c/dev/src/my-portfolio/bin/spring:15:in \`&lt;top (required)&gt;'`

`bin/rails:3:in \`load'`

`bin/rails:3:in \`&lt;main&gt;'`

`Tasks: TOP =&gt; webpacker:install =&gt; webpacker:check_yarn`

`(See full trace by running task with --trace)`

&amp;#x200B;

I'm not sure if this is because the rails server should be started already? If I try to start a rails server I get an even longer output of an error, but I'm not sure if this is what should fixed first before being able to start it in the first place. 

I'm still quite a noob on a lot of this type of stuf. My coding bootcamp took us mainly through front end stuff and a lot of Rails specifically using vagrant. I've wanted to try something else other than vagrant now that I'm on my own.  I'm trying to get my portfolio finally finished so I can launch it to the domain I purchased ASAP, so any help getting my environment set up all silky smooth would be much appreciated!
## [9][Updating legacy storage](https://www.reddit.com/r/rails/comments/gwho9k/updating_legacy_storage/)
- url: https://www.reddit.com/r/rails/comments/gwho9k/updating_legacy_storage/
---
Our application stores all images and files in an S3 bucket. Uploads are handled by Paperclip in various places. We're trying to find the best solution to update. 

Initially, we were looking at Active Storage because it natively supports one-to-many relations, but don't like the lack of control (mostly regarding filenames and [paths](https://stackoverflow.com/questions/48389782/how-to-specify-a-prefix-when-uploading-to-s3-using-activestorages-direct-upload)). 

We've also been looking into Shrine, which does less out of the box, but seems to allow for more customization.

What do you use and why? Are there pitfalls to lookout for?
## [10][Should I make a POST request with a PNG blob as a parameter in Rails 5?](https://www.reddit.com/r/rails/comments/gwhoct/should_i_make_a_post_request_with_a_png_blob_as_a/)
- url: https://www.reddit.com/r/rails/comments/gwhoct/should_i_make_a_post_request_with_a_png_blob_as_a/
---
Hey! I am developing a product customization tool on my Rails 5 webapp, which requires me to send the final design's blob string to a different web-service via the webapp's server (in order to generate a composite preview).

Is it scalable if I am making a POST request to my server with blob string as a parameter? I fear that can bloat the memory. Am I right?

The blob size will be around 40-50 Kb tops.

Thank you in advance.
## [11][100 days of code to learn rails](https://www.reddit.com/r/rails/comments/gw2y2l/100_days_of_code_to_learn_rails/)
- url: https://www.reddit.com/r/rails/comments/gw2y2l/100_days_of_code_to_learn_rails/
---
just wanted to share I'm starting my web dev journey and choosing to go through Hartl's rails 6 tutorial. a lot of negatives on rails hype dying, but really looking forward to get a foundation of being able to build MVP's for some of my ideas I have. I've been stalking this sub a bit and it's really cool that you guys are so practical and helpful. I hope it's okay to share here and probably will be posting here for some help along the way :)

starting a substack if anyone is interested in following a newb: [https://joelchoi.substack.com/p/day-1](https://joelchoi.substack.com/p/day-1)
