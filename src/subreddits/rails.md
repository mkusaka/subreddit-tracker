# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/fx6je4/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/fx6je4/personal_projects_show_off_your_own_project_andor/
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
## [3][Beginner Friendly Guided Project - Building a Link Shortener with Rails 6](https://www.reddit.com/r/rails/comments/g0uvsg/beginner_friendly_guided_project_building_a_link/)
- url: https://www.reddit.com/r/rails/comments/g0uvsg/beginner_friendly_guided_project_building_a_link/
---
Hi there. 

I've been building Rails apps for a long time, but I never fully embraced TDD. Now that I've got a bit more time on my hands thanks to all this COVID business, I decided to give it more of a go.

I'm recording a series on Youtube of me building a simple link shortener from scratch with TDD. It's pretty beginner friendly, and if you go through the whole thing with me, you'll have a real project to show for it at the end.

Just thought I'd drop this here in case it helps someone: [https://www.youtube.com/playlist?list=PLjItgYqIzJ9WGy9WMf-44DXHB\_7aWwJMc](https://www.youtube.com/playlist?list=PLjItgYqIzJ9WGy9WMf-44DXHB_7aWwJMc)

If you watch the playlist and have any questions, just ask in the comments of the video and I'll give you the best answer I can.

Hope everyone is well.

Peace.
## [4][Rails 6 CSS organization](https://www.reddit.com/r/rails/comments/g14b3n/rails_6_css_organization/)
- url: https://www.reddit.com/r/rails/comments/g14b3n/rails_6_css_organization/
---
I'm finally toying around Ruby-on-Rails 6 (better later than never), JavaScript has finally evolved with webpack, which is a good thing, but I'm quite surprised to see that "The CSS Rails way" didn't changed. It's \*\*still\*\* old-fashioned CSS by default (where is SASS ? BEM ?), using the "=require" directive. This is a bit annoying when setting up new rails project with the "rails new" command. Moreover, this is not consistent with using webpacker, this makes 2 front-end builds (webpack + sprockets) running side-by-side. Finally, it do not encourage "componentization of the front-end", unless you choose a "full JS Framework" like React or Vue that you may not need anyway... So my question is : why is CSS default organization \*\*still\*\* so old-fashioned after all these years ?
## [5][Cleaning up my messy gem environment](https://www.reddit.com/r/rails/comments/g13r2s/cleaning_up_my_messy_gem_environment/)
- url: https://www.reddit.com/r/rails/comments/g13r2s/cleaning_up_my_messy_gem_environment/
---
I am half way through a udemy course on Rails, and in that course, they recommended the installing of rvm, and a whole bunch of gems. When I run rails routes, it now looks a mess of different options! Also, when I run `rails generate controller Articles` I get this error:

&gt;The dependency tzinfo-data (&gt;= 0) will be unused by any of the platforms Bundler is installing for

And Terminal just freezes.

So, with this in mind, I was thinking it would be nice to have a cleaner environment where I just have rails installed, as I repeatedly go through the Getting Started with Rails tutorial on [rubyonrails.org](https://rubyonrails.org).

With this goal in mind, what is the best way to go about it? I don't want to uninstall all the gems I added, because I want to finish the course at udemy. But while going through the [rubonrails.org](https://rubonrails.org) tutorials, I want to keep things a bit cleaner.

What would you recommend?
## [6][turbolinks + early hints + localstorage = pseudo SPA?](https://www.reddit.com/r/rails/comments/g0xz7v/turbolinks_early_hints_localstorage_pseudo_spa/)
- url: https://www.reddit.com/r/rails/comments/g0xz7v/turbolinks_early_hints_localstorage_pseudo_spa/
---
Just wondering if anyones managed to set up something like this for their reactive front ends? I'm mainly thinking of this for a standard react-rails and normal rails doing the server rendering a la `&lt;%= react_component "thing" %&gt;`. Also with some front end cache invalidation/refetching?
## [7][Docker - Failed to open TCP connection to localhost:35729 (Cannot assign requested address](https://www.reddit.com/r/rails/comments/g0x7ax/docker_failed_to_open_tcp_connection_to/)
- url: https://www.reddit.com/r/rails/comments/g0x7ax/docker_failed_to_open_tcp_connection_to/
---
Trying to get my Rails 5.2 app running on a Docker host on Windows10 and I feel that I'm 90% of the way there:

The app container error that I get when visiting `localhost:3000`

    Started GET "/" for 192.168.240.1 at 2020-04-14 02:29:21 +0000
    app_1  | Cannot render console from 192.168.240.1! Allowed networks: 127.0.0.1, ::1, 127.0.0.0/127.255.255.255
    app_1  |
    app_1  | PG::ConnectionBad (could not connect to server: Connection refused
    app_1  |        Is the server running on host "127.0.0.1" and accepting
    app_1  |        TCP/IP connections on port 5432?
    app_1  | ):
    app_1  |

My Dockerfile:

    FROM ruby:2.6.1
    RUN apt-get update -qq &amp;&amp; apt-get install -y nodejs postgresql-client
    RUN mkdir /usr/src/app
    WORKDIR /usr/src/app
    COPY . /usr/src/app
    RUN gem update --system &amp;&amp; gem install bundler &amp;&amp; bundler update --bundler
    RUN bundle install
    
    # Add a script to be executed every time the container starts.
    COPY docker_entrypoint.sh /usr/bin/
    RUN chmod +x /usr/bin/docker_entrypoint.sh
    ENTRYPOINT ["docker_entrypoint.sh"]
    EXPOSE 3000
    
    # Start the main process.
    CMD ["rails", "server", "-b", "0.0.0.0"]

My docker-compose.yml:

    version: '3'
    services:
      app:
        build: .
        command: bash -c "rm -f tmp/pids/server.pid &amp;&amp; bundle exec rails s -p 3000 -b '0.0.0.0'"
        environment:
          - POSTGRES_HOST=db
          - POSTGRES_USER=postgres
          - POSTGRES_PASSWORD=password
        volumes:
          - app:/usr/src/app
        ports:
          - 127.0.0.1:3000:3000
        links:
          - db
        depends_on:
          - db
      db:
        image: postgres
        volumes:
          - postgres:/var/lib/postgresql/data
        environment:
          - POSTGRES_USER=postgres
          - POSTGRES_PASSWORD=password
    volumes:
      postgres:
      app:

My understanding is that `app` and `db` should be open binded together with `links:` but doesn't appear be so. A similar error requires me to have named volumes instead of relative as the two containers cannot communicate with each other (a windows specific permissions issue).

Am I close? Any suggestions? Thank you
## [8][Issue with deleting active storage image attachment](https://www.reddit.com/r/rails/comments/g0n04i/issue_with_deleting_active_storage_image/)
- url: https://www.reddit.com/r/rails/comments/g0n04i/issue_with_deleting_active_storage_image/
---
 

I have a products page in the admin namespace. I have image attachments for each product via active storage. I am trying to have a link to remove the images but I can't seem to get it to work.

link:

    &lt;% admin_product.images.each do |image| %&gt;
    
    &lt;%= image_tag url_for(image), :class =&gt; 'img-fluid' %&gt;
    
    &lt;%= link_to 'Remove', delete_image_attachment_admin_product(image.id), method: :delete, data: { confirm: 'Are you sure?' } %&gt;
    
    &lt;% end %&gt;

Admin/products controller method:

    def delete_image_attachment
      @image = ActiveStorage::Attachment.find(params[:id]) 
      @image.purge     
      redirect_back(fallback_location: request.referer) 
    end

I keep getting errors saying there is no method in products/edit. I'm sure I have the URL wrong to hit the delete\_image\_attachment method but can't seem to get it straight.

Any suggestions are appreciated!
## [9][How Do You Organize Your Stylesheets?](https://www.reddit.com/r/rails/comments/g0oohy/how_do_you_organize_your_stylesheets/)
- url: https://www.reddit.com/r/rails/comments/g0oohy/how_do_you_organize_your_stylesheets/
---
My stylesheets have grown to one big monolith and have incurred a lot of code debt. I am going to reorganize everything. How do you guys organize and structure your stylesheets?
## [10][RVM/Bundler : gem uninstall -aIX fails. -- Why are there two different directories for gems that are being used?](https://www.reddit.com/r/rails/comments/g0pubd/rvmbundler_gem_uninstall_aix_fails_why_are_there/)
- url: https://www.reddit.com/r/rails/comments/g0pubd/rvmbundler_gem_uninstall_aix_fails_why_are_there/
---
Hi everyone,

I'm gonna start with telling you what is the goal.

**Goal:**

1. Uninstalling everything related to rvm, bundler, and gems. Literally making things gone.
2. Then re-installing everything back again as if this is a new fresh server.
3. Installing gems after becoming a correct system user through bundler to avoid any permission related errors on any ruby on rails app.
4. This should work on any system regardless of how f\*\*\*ed up the rvm/ruby/gem configuration.

**The problem:**

Hi, gem uninstall -aIX fails (Goal #1) because turns out the gems are installed in different directories etc.

And the error log tells me to go to a different directory and run the commands over there.

$GEM\_PATH: /usr/local/rvm/gems/ruby-2.6.3:/usr/local/rvm/gems/ruby-2.6.3@global

$GEM\_HOME: /usr/local/rvm/gems/ruby-2.6.3

**Error log:**

**$ gem uninstall -aIX**

ERROR: While executing gem ... (Gem::InstallError)

bundler-unload is not installed in GEM\_HOME, try:

gem uninstall -i /usr/local/rvm/rubies/ruby-2.6.3/lib/ruby/gems/2.6.0 bundler-unload

**The question:**

Why the second path exists for some of the gems? (and why the second directory ends with 2.6.0)

\- /usr/local/rvm/gems/ruby-2.6.3/gems/

\- /usr/local/rvm/rubies/ruby-2.6.3/lib/ruby/gems/2.6.0
## [11][Using Atom IDE with Rails - package removal and additions](https://www.reddit.com/r/rails/comments/g0jz0a/using_atom_ide_with_rails_package_removal_and/)
- url: https://www.reddit.com/r/rails/comments/g0jz0a/using_atom_ide_with_rails_package_removal_and/
---
I am using Atom as my IDE for Rails, and it says I have 82 installed packages. That seems rather a lot! I think most came as default. For example, it has C# language support that I never use.

Do you remove any of these?

I am thinking that when you get the popup list of options when you type your code, many won't be relevant if you have package installs for Go Language etc. Or have I got that wrong?

What else do you add if you are using HTML, CDD, Ruby and Rails?
## [12][Is @ option or not in link_to?](https://www.reddit.com/r/rails/comments/g0hgo8/is_option_or_not_in_link_to/)
- url: https://www.reddit.com/r/rails/comments/g0hgo8/is_option_or_not_in_link_to/
---
I am a little confused on link\_to method. Let's say you have link\_to 'article', article\_path(@article). Do I have to use the @? Or is the @ only used in earlier Ruby versions?
