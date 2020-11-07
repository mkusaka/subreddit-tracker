# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/jnwqje/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/jnwqje/personal_projects_show_off_your_own_project_andor/
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
## [3][Bootstrap-table-rails](https://www.reddit.com/r/rails/comments/jpnu5h/bootstraptablerails/)
- url: https://www.reddit.com/r/rails/comments/jpnu5h/bootstraptablerails/
---
hi, I'm trying to install this gem: https://github.com/bjevanchiu/bootstrap-table-rails in rails 6. The problem is that when I put it in the gemfile and run the bundle install command, it doesn't appear in node_modules. With linux find I looked for the bootstrap-table-rails directory but I can't find it anywhere. How can I integrate Bootstrap-table into rails? Thanks so much for the future help
## [4][Securing Rails API access - what method?](https://www.reddit.com/r/rails/comments/jpgg3d/securing_rails_api_access_what_method/)
- url: https://www.reddit.com/r/rails/comments/jpgg3d/securing_rails_api_access_what_method/
---
Imagine your typical Rails API scenario: you're building an API but you want to restrict the access to it (also to have opportunity to see how many requests a certain consumer of the API made for statistical purposes).

One way which seems simple enough and is often used in tutorials on the net is securing it using API keys which are tied to the other Rails apps consuming this new fancy API. 

My question is: where do you put the list of generated API keys? Inside the API itself in some table (e.g. approved_clients)? Or do you use some other solution?

I'm trying to brainstorm here which approach I should take when building an API shared between several Rails-based apps.

Using oAuth makes no sense here to me because the said API would not be user-centric but Rails app (client) centric: e.g. you'd ask the shared geo API to return you a list of cities close to given geo coordinates. So we don't need to login individual users here to do something on them, just to have a shared services that's generic enough no to involve oAuth or other more complicated flows. 

Or could JWTs somehow fit here?

Thanks for your ideas.
## [5][paragraphs not showed in my markdown redcarpet](https://www.reddit.com/r/rails/comments/jpq7mt/paragraphs_not_showed_in_my_markdown_redcarpet/)
- url: https://www.reddit.com/r/rails/comments/jpq7mt/paragraphs_not_showed_in_my_markdown_redcarpet/
---
Hi guys. I'm editing the website made by another developer.

He used markdown redcarpet gem but I'm noticing that it doesn't show the paragraphs (tag &lt;p&gt;).

What to check to enable them?
## [6][Confused about pulling changes without overwriting local env](https://www.reddit.com/r/rails/comments/jpgc5y/confused_about_pulling_changes_without/)
- url: https://www.reddit.com/r/rails/comments/jpgc5y/confused_about_pulling_changes_without/
---
Can somebody explain with a TL;DR on this? I'm trying to figure out the following scenario:

\- I clone a repo and get it up and running with bundle install, setup local db, create env file, compile webpacker, and so on. This goes fine.

\- A dev makes changes to that remote repo which I want to get on my local machine. What do I do now? How do I get those changes onto my machine without overwriting my env file or any other setup I've done locally?
## [7][Best React resource for Rails developers](https://www.reddit.com/r/rails/comments/jpb17u/best_react_resource_for_rails_developers/)
- url: https://www.reddit.com/r/rails/comments/jpb17u/best_react_resource_for_rails_developers/
---
hey friends, I'm a purely backend developer and I barely played with React so far. What are the best resources that you recommend for me to go over to get up to speed when it comes to React on top of Rails.
## [8][I created an Xbox backward compatible game catalogue with Rails](https://www.reddit.com/r/rails/comments/jpc1tb/i_created_an_xbox_backward_compatible_game/)
- url: https://www.reddit.com/r/rails/comments/jpc1tb/i_created_an_xbox_backward_compatible_game/
---
Hi there,

I've recently purchased an Xbox Series S and excited about the games from the previous generation as I didn't have much time to dive in. I created [this small web app](https://backwardible.com/) with Rails and thought it might help other people too.

Oh, I tend to write blog posts about the side-projects I do, so if anyone's interested in terms of what I used and why I used, the concerning post is [here](https://arkan.me/backwardible/) as well!

I would love to hear your feedback.
## [9][How to scaffold a preconfigured Rails base app ?](https://www.reddit.com/r/rails/comments/jp9dj3/how_to_scaffold_a_preconfigured_rails_base_app/)
- url: https://www.reddit.com/r/rails/comments/jp9dj3/how_to_scaffold_a_preconfigured_rails_base_app/
---
I try new ideas frequently and I am building Rails apps from scratch by copying migrations, Gemfile/initializers, front-end libs (bootstrap, jquery etc.) then customize the additional stuff (new models, controllers etc.).

A great part of the apps is the same, e.g. devise/pundit for auth, bullet/robocop/letter\_opener for development etc. 

I want to build a \_base\_ for my apps and I want to be able to scaffold a new one easily with just a command. Creating an app and pushing to a git repo could do the trick, but I wonder If there are any other ways, for example building an engine with all these and just include it in the Gemfile. 

What do you guys do in similar cases?
## [10][How to clone record with active storage attachments?](https://www.reddit.com/r/rails/comments/jpd2zv/how_to_clone_record_with_active_storage/)
- url: https://www.reddit.com/r/rails/comments/jpd2zv/how_to_clone_record_with_active_storage/
---
I have a function to clone records in a rails application. In addition to the record data I would like to copy/attach any active storage file uploads that are attached to the source object to the new object. The files are stored on local storage. I would like to create new files to attach and not just point the cloned record at the same files. Any ideas on how to do this? The dup record works great. Just need to clone the attachments. Here is my action and model:

Trying to do something like this:

      def copy
       u/source = Compitem.find(params[:id])
       u/compitem = u/source.dup
       if u/source.uploads.any?
          u/source.uploads.each do |upload|
            cloned_upload= upload.dup
            u/compitem.uploads &lt;&lt; cloned_upload
          end
        end
       render 'new'
      end

Model:

    class Compitem &lt; ApplicationRecord
      belongs_to :user
      has_many_attached :uploads, dependent: :destroy
    end
## [11][Gem for intercepting server errors for report purposes?](https://www.reddit.com/r/rails/comments/jp9887/gem_for_intercepting_server_errors_for_report/)
- url: https://www.reddit.com/r/rails/comments/jp9887/gem_for_intercepting_server_errors_for_report/
---
Is there any Gems/middleware that allow you to intercept any server errors and send the stacktrace via email or make a custom API call? 

Basically have a few servers that are in beta, they're going to have errors but I'd like to either email or do some kind of handling if a 500 error occurs.
## [12][teaspoon or blade runner to browser test the js code you write in a rails application?](https://www.reddit.com/r/rails/comments/jp7qvi/teaspoon_or_blade_runner_to_browser_test_the_js/)
- url: https://www.reddit.com/r/rails/comments/jp7qvi/teaspoon_or_blade_runner_to_browser_test_the_js/
---
  


https://preview.redd.it/dmoi45958nx51.png?width=2400&amp;format=png&amp;auto=webp&amp;s=fe5ed733b0bef76a2e2f78aa9b5aefec8222f7e3

&amp;#x200B;

which one do you recommend/prefer
