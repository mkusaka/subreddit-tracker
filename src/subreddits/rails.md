# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/har6r7/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/har6r7/personal_projects_show_off_your_own_project_andor/
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/hfkxk4/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/hfkxk4/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][Issue with route](https://www.reddit.com/r/rails/comments/hhyrfd/issue_with_route/)
- url: https://www.reddit.com/r/rails/comments/hhyrfd/issue_with_route/
---
I have a list of product variants that belong\_to a product. I am trying to click the show link on the variants index and have it hit the url like product/1/variant/8. I keep getting an error stating that it is a no method in variants error. What am I doing wrong?

Rake Routes:

    product_variant GET    /products/:product_id/variants/:id(.:format)                                             variants#show

Link\_to:

      &lt;% @variants.each do |variant| %&gt;
        &lt;%= content_tag :tr, id: dom_id(variant), class: dom_class(variant) do %&gt;
    
    
                      &lt;td&gt;&lt;%= variant.sku %&gt;&lt;/td&gt;
                      &lt;td&gt;&lt;%= variant.price_cents %&gt;&lt;/td&gt;
                      &lt;td&gt;&lt;%= variant.price_currency %&gt;&lt;/td&gt;
                      &lt;td&gt;&lt;%= variant.is_master %&gt;&lt;/td&gt;
                      &lt;td&gt;&lt;%= variant.product.name %&gt;&lt;/td&gt;
    
    
            &lt;td&gt;&lt;%= link_to 'Show', product_variant_path(variant) %&gt;&lt;/td&gt;
    
        &lt;% end %&gt;
      &lt;% end %&gt;

Error with \_path added:

    No route matches {:action=&gt;"show", :controller=&gt;"variants", :product_id=&gt;#&lt;Variant id: 37, sku: "SK010", price_cents: 20, price_currency: "USD", is_master: nil, product_id: 6, created_at: "2020-06-29 03:07:14", updated_at: "2020-06-29 03:07:14", weight: nil, count_on_hand: 0, position: nil&gt;}, missing required keys: [:id]

&amp;#x200B;

routes:

    resources :products do
        resources :option_values
        resources :option_value_variants
        resources :variants do
            collection do
              post :update_positions
            end
          end
      end
## [4][Resources for learning RSpec for Rails?](https://www.reddit.com/r/rails/comments/hhggoh/resources_for_learning_rspec_for_rails/)
- url: https://www.reddit.com/r/rails/comments/hhggoh/resources_for_learning_rspec_for_rails/
---
Hey All. I am learning RSpec for work. I've taken a couple of courses covering the basics of testing Ruby with RSpec, and now I want to apply that knowledge to a Rails application.   


Any suggestions on a course/exercises that I can follow? I want to be able to look up answers/explanations when I get stuck (which is often)   


Thanks
## [5][Basecamp/Dhh prefers mysql over postgres? did i get that right](https://www.reddit.com/r/rails/comments/hhimlt/basecampdhh_prefers_mysql_over_postgres_did_i_get/)
- url: https://www.reddit.com/r/rails/comments/hhimlt/basecampdhh_prefers_mysql_over_postgres_did_i_get/
---
[https://gist.github.com/dhh/782fb925b57450da28c1e15656779556](https://gist.github.com/dhh/782fb925b57450da28c1e15656779556)
## [6][Opinions on various AWS hosting options - ElasticBeanstalk vs ECS/Fargate vs Chef](https://www.reddit.com/r/rails/comments/hgwy07/opinions_on_various_aws_hosting_options/)
- url: https://www.reddit.com/r/rails/comments/hgwy07/opinions_on_various_aws_hosting_options/
---
Hello everyone,

I've been tasked with building a RESTful API.

Initially, the prototype was various Lambda functions backed by AWS API Gateway

We are pivoting to Rails because of all the amazing bells and whistles that come with a framework.

Right now, I have the Rails API running on localhost. I am also running Sidekiq in a separate terminal session. We are leveraging PostgreSQL as our database.

We are ready to start deploying the application to our AWS environment. The team is highly-skilled in AWS and so we open to different types of deployment and hosting options. 

I'm hoping to receive some feedback on the Pros/Cons of hosting a Rails app via the following AWS Services

# Option #1

## Bootstrapped via Chef

This is our current design specification. We wrote some Terraform IaC to build the following resources

- Application Load Balancer

- Auto Scaling Group

- PostgreSQL RDS Instance

The Chef cookbook, bootstraps the instances in the ASG by performing the following commands

- Clones the source code repo

- Runs bundle install, db:migrate, etc

- Runs redis-server in daemon mode, sidekiq, and rails server

- Nginx fronted in each app server to perform reverse proxy



__Pros__ - Design and code is already done, we can begin testing deployment

__Cons__ - Coming from a Heroku background, it does not seem intuitive enough to perform rolling updates/blue green deployments. We will have to write some automation to perform that. And possible rollbacks seem difficult. And that is why I started to research Option #2


## Option #2 ElasticBeanstalk

ElasticBeanstalk looks like it will give us the closest usability like Heroku. We can simply run \`eb deploy\` at first, and then incorporate a CI/CD pipeline to perform the deployment for us. I like the fact that application versions are stored so that we can easily rollback if needed.

__Pros__

- Most "Heroku"-like service in AWS. 

- Leverage the eb cli for deployment

- Application Versions allow for easy rollback

__Cons__

- There will be development work needed to figure out .ebextensions, e.g. how to run sidekiq. A simple default Rails 6 app does not simply "deploy" to ElasticBeanstalk successfully out of the box. Theres problems with Node.js versions, webpacker, etc. 


## Option #3 ECS/Fargate

I'm all for containerizing our workloads to help mitigate against differences in envrionment e.g. localhost, dev, test, prod, etc. So thats why I started to look into ECS/Fargate. AWS has released a ruby gem for provisioning a cluster. But unfortunately, its in preview mode/experimental. I spent a few hours last night playing around with it. I believe it will take some development time to massage the CDK output to match our infrastructure requirements, e.g. naming conventions, tagging, instance sizes etc. It seems like every change I make is overwritten via `aws-rails-provisioner build`.

[https://github.com/awslabs/aws-rails-provisioner](https://github.com/awslabs/aws-rails-provisioner)

__Pros__
- Containerizing our application can mitigate differences between environments.
- ECS supports various deployment options
- Closets "single click" deployment option I've found

__Cons__
- The `aws-rails-provisioner` rubygem is in experimental stage. 
- The infrastructure output from the CDK does not match our requirements and will need to be updated and ensured that it is not overwritten via `aws-rails-provisioner build`


It's hard to look through a crystal ball to understand the long-term tradeoffs for each option. I am hoping to find the most "Heroku-like" experience so that we can iterate quickly and efficiently without having to write a bunch of automation code. 

Thank you
## [7][Rails &amp; Amazon SES for emails. Best practices and throttling?](https://www.reddit.com/r/rails/comments/hgrr27/rails_amazon_ses_for_emails_best_practices_and/)
- url: https://www.reddit.com/r/rails/comments/hgrr27/rails_amazon_ses_for_emails_best_practices_and/
---
Hi all,

For a Rails site I'd like to switch to Amazon SES for sending emails because we're having some trouble delivering to the big providers from our own servers.

Now, SES has a pretty low throttling threshold by default (1/sec) which is a problem due to the nature of my application. My users send out invoices once a month, in bursts of thousands. Once I exceed Amazon's rate, emails get rejected. I need to do some throttling / queuing *somewhere*.

The way I see it, I can do that in 2 different ways:

1. Have the Rails app talk directly to the SES SMTP and do the throttling in-app, perhaps through something like DelayedJob?
2. Have the Rails app talk to a local MTA like Exim, which sets up a relay to SES, including throttling and retries.

Options 2 seems to be more robust, but it requires yet another service to be configured and maintained.

&amp;#x200B;

Secondly, anyone know any good guides how to set up SES with SNS on AWS to get insight into the individual bounces, rejects, etc.? I found their original documentation a bit lacking.
## [8][Exploring Rails View Components](https://www.reddit.com/r/rails/comments/hgbe2n/exploring_rails_view_components/)
- url: https://www.reddit.com/r/rails/comments/hgbe2n/exploring_rails_view_components/
---
I recorded a short exploration of the view\_component gem from Github. I have to say I'm pretty excited about it. There's some good stuff going on in the Stimulus Reflex community around ViewComponentReflex([https://github.com/joshleblanc/view\_component\_reflex](https://github.com/joshleblanc/view_component_reflex)), so I figured I'd start here before jumping into that: [https://youtu.be/oKqJmEAn-X0](https://youtu.be/oKqJmEAn-X0)
## [9][Acts as Tracked gem: selectively track changes made on your AR models](https://www.reddit.com/r/rails/comments/hg8fe0/acts_as_tracked_gem_selectively_track_changes/)
- url: https://www.reddit.com/r/rails/comments/hg8fe0/acts_as_tracked_gem_selectively_track_changes/
---
Hi, i've made a gem to selectively track changes on AR models, where audited gem would be an overkill. ActsAsTracked can be plugged into ActiveRecord model, and then used whenever you need a history of changes and actors made on the record.

You can find docs in the [repository](https://github.com/ramblingcode/acts-as-tracked)

Blog post [here](https://www.ramblingcode.dev/posts/track_changes_on_your_activerecord_models/)

Example usage in this [project](https://github.com/ramblingcode/rails6-acts-as-tracked-usage)

Hope you find it useful.
## [10][How to access a variable from a ruby controller action in JS?](https://www.reddit.com/r/rails/comments/hgcijg/how_to_access_a_variable_from_a_ruby_controller/)
- url: https://www.reddit.com/r/rails/comments/hgcijg/how_to_access_a_variable_from_a_ruby_controller/
---
I have a variable created in an action within one of my ruby controllers. I can access it from the corresponding view, but how can I pass it so that my JS can grab it? Does it need to be embedded in the view as a hidden tag/variable?
## [11][Glimmer DSL for Opal v0.0.8 (Rails Web GUI Adapter for Desktop Apps)](https://www.reddit.com/r/rails/comments/hgb3q5/glimmer_dsl_for_opal_v008_rails_web_gui_adapter/)
- url: https://www.reddit.com/r/rails/comments/hgb3q5/glimmer_dsl_for_opal_v008_rails_web_gui_adapter/
---
Glimmer DSL for Opal v0.0.8 is out with table data-binding support, including selection, sorting, and editing (can be easily styled with standard CSS). Glimmer DSL for Opal is an experimental web GUI adaptor for webifying [Glimmer](https://github.com/AndyObtiva/glimmer) desktop apps (i.e. apps built with [Glimmer DSL for SWT](https://github.com/AndyObtiva/glimmer-dsl-swt)) via [Rails](https://rubyonrails.org/) and [Opal](https://opalrb.com/) without changing a line of code. Apps may then be custom-styled for the web via standard CSS.

[https://github.com/AndyObtiva/glimmer-dsl-opal](https://github.com/AndyObtiva/glimmer-dsl-opal)

&amp;#x200B;

[Glimmer DSL for SWT \(Desktop App\)](https://preview.redd.it/3bapevj4aa751.png?width=393&amp;format=png&amp;auto=webp&amp;s=a7d788bbe54350ef0378d507d72d3ed4c26be914)

[Glimmer DSL for Opal \(Rails Web App\)](https://preview.redd.it/asyhz831aa751.png?width=612&amp;format=png&amp;auto=webp&amp;s=df6751b0a364bf836524992f3c979a864e7604e2)
## [12][Help choosing model associations](https://www.reddit.com/r/rails/comments/hgb12a/help_choosing_model_associations/)
- url: https://www.reddit.com/r/rails/comments/hgb12a/help_choosing_model_associations/
---
So i have a multi user blog platform, it's a simple crud app but my issue is with the association because i have many types of users (i have simple user that can write blog post and it show his name) and i have a company account (which is like the user but have other information like company website instead of the user username)


my issue is how to handle the association (i already implemented the user model using device) which have_many blog 

so every blog have the belong to associated to the user

how to add another user type (Company) which have the the same many blog 

i searched and found the STI but it's not what i need as there is differences between the users types ( i may add another type later) 



should i use polymorphic association or are there another method to this relationship    





TL;DR i have user, CompanyUser, blablaUser which all of them can make blog post how to make the association ?
