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
## [3][What are your thoughts on the front-end technology behind "Hey," the new email service by DHH?](https://www.reddit.com/r/rails/comments/hifsxg/what_are_your_thoughts_on_the_frontend_technology/)
- url: https://www.reddit.com/r/rails/comments/hifsxg/what_are_your_thoughts_on_the_frontend_technology/
---
I've been following the development of Hey very closely because of the supposed usage of new, revolutionary frontend paradigms (as touted by DHH), but after a few minutes of playing around with Hey it's pretty clear there are limitations.


There's no getting around the fact that the whole application *feels* like a Turbolinks + Stimulus application. Interactions don't feel snappy in the way that React/Angular/Svelte/Vue applications do as there's perceptible lag between interactions and page loads. This is in contrast to modern SPA style applications, where API interactions are optimistically triggered and non-API interactions are instantaneous.


Secondly, it seems like the user experience was built around the limitations of Turbolinks + Stimulus, in that anytime you want to perform a new workflow, you're directed to a new page. Want to write an email? New page. Check out some starred emails? New page. Read an email? New page. Sure, you can create a SPA style application using Turbolinks &amp; Stimulus, but given that these two libraries were never designed to be used as such, you'll quickly run into growing pains.


What are your thoughts?
## [4][Willing to pay $/hr to chat/seek help from someone who has previous experience deploying a Rails 5 app with Puma to AWS Elastic Beanstalk](https://www.reddit.com/r/rails/comments/hifmbd/willing_to_pay_hr_to_chatseek_help_from_someone/)
- url: https://www.reddit.com/r/rails/comments/hifmbd/willing_to_pay_hr_to_chatseek_help_from_someone/
---
**My details:**

Ruby version: 2.6.6p146

Rails version:  [5.2.4.3](https://5.2.4.3)

Puma version: 4.3.5

AWS Beanstalk Platform: **Ruby 2.6 AL2 version 3.0.3** *64bit Amazon Linux 2 v3.0.3 running Ruby 2.6*

Any help would be greatly appreciated. Please let me know how much you would like to charge to help me deploy or understand my log file errors.
## [5][Best way to add a search/filter/sort to a list of records?](https://www.reddit.com/r/rails/comments/hinagw/best_way_to_add_a_searchfiltersort_to_a_list_of/)
- url: https://www.reddit.com/r/rails/comments/hinagw/best_way_to_add_a_searchfiltersort_to_a_list_of/
---
I'm looking for a package or gem that will let me add this functionality rather than having to code it from scratch.  What's the best solution for this?

Ideally, it would be nice to have it async so I don't have to refresh the page but that's not too important.
## [6][Sign into Rails app using Devise from PHP app using cURL, then redirect to a specified protected page in Rails](https://www.reddit.com/r/rails/comments/hi7uok/sign_into_rails_app_using_devise_from_php_app/)
- url: https://www.reddit.com/r/rails/comments/hi7uok/sign_into_rails_app_using_devise_from_php_app/
---
I'm a complete newbie to Rails and have been charged with solving a unique problem. I have a PHP app that needs to "SSO" a user into a Rails app using Devise that we inherited, then redirect them to a specific page within Rails app instead of the home page. Nobody in our shop has Rails experience, but it is the hand I've been dealt.

Ideally here's what I'd like to do:

1. User is authenticated in PHP app and visits protected page there, where they can access an external Rails app that uses Devise and MySQL.
2. Verify that a matching account exists in the Rails app, and add the matching account if they don't.
3. Use cURL to then authenticate the user with Devise, then upon successful authentication, redirect them to the page they selected, already logged in.

Right now, on redirect to the Rails app after the cURL request, it is stopping at the home page with a login prompt. When logging in, it just shows the home page rather than going to the link they chose. The Rails app tracks activities for each individual user, so using a generic login will not work; it has to be user specific.

If this is not possible, I would be satisfied with handling the authentication and custom redirect in Rails  / Devise. I suppose this would involve passing some user information as GET parameters in the redirect, along with the final destination page URL (there's 7 possible based on the specific activity button they click). Then, in Rails, it would extract the user info and destination URL, sign them in, then redirect to that destination.

The PHP end is pretty much done. But since I have no experience with Rails, I'm rather stuck on how to do this, or if there is a better way. Can anyone offer some guidance?
## [7][[Need Help] RoR with some python sprinkled on top.](https://www.reddit.com/r/rails/comments/hiep85/need_help_ror_with_some_python_sprinkled_on_top/)
- url: https://www.reddit.com/r/rails/comments/hiep85/need_help_ror_with_some_python_sprinkled_on_top/
---
Been working on a personal project lately and hit a wall. There's a feature I'd like to run, sort of a call to action that takes previous data and shows a "prediction" of what "might" occur should the trend continue without action. I'm seeing countless python tutorials on predicting stock prices (which seems like a step in the right direction with a different object), but no real resources for Ruby. So I'm wondering, have any of you ran some Python3 through a RoR app? If so, what would you say is the most efficient integration for it?
## [8][How do I approach setting up an outside API for existing rails app](https://www.reddit.com/r/rails/comments/hicirq/how_do_i_approach_setting_up_an_outside_api_for/)
- url: https://www.reddit.com/r/rails/comments/hicirq/how_do_i_approach_setting_up_an_outside_api_for/
---
Hey, sorry if my title wasn't clear enough basically I want to build a mobile app that uses data from my main web rails project through various API endpoints.

I'm just curious what I'd need to do to get there, specifically when it comes to security.
## [9][Unable to install Rails in Vagrant environment passed 4.2.10?](https://www.reddit.com/r/rails/comments/hi9o1g/unable_to_install_rails_in_vagrant_environment/)
- url: https://www.reddit.com/r/rails/comments/hi9o1g/unable_to_install_rails_in_vagrant_environment/
---
Hi there. I finally got my vagrant working after weeks of trouble. I was able to create my own Vagrantbox and install the latest version of ruby (2.7.1 as of now). Then I tried to install rails via "gem install rails" and I get an error message. It's like I'm getting a corrupted file directly from the gem website. What the heck?  Here's what happens: 

    vagrant@ubuntu-bionic:~$ gem install rails
    ERROR:  Error installing rails:
            invalid gem: package is corrupt, exception while verifying: undefined method `path' for "data.tar.gz":String (NoMethodError) in /home/vagrant/.rvm/gems/ruby-2.7.1/cache/mimemagic-0.3.5.gem

I installed Ruby using RVM. Installed it and then rvm install 2.7.1. The next step on the tutorial was to do the gem install rails command. But I ended up trying "sudo apt install rails" and it only installed version 4.2.10. Rails is at 6.0.3 I think now and I'd like that version. How else can I download rails in my vagrant environment?
## [10][Can you recommend convert text to image gems](https://www.reddit.com/r/rails/comments/hi14ou/can_you_recommend_convert_text_to_image_gems/)
- url: https://www.reddit.com/r/rails/comments/hi14ou/can_you_recommend_convert_text_to_image_gems/
---
I'm building a feature that would convert typed text to image. I'm not looking for "caption image" solution. I want to convert entire page converted to JPG/PNG with styling etc.



Best I can think of right now is to use [Prawn gem](https://github.com/prawnpdf/prawn) (gem for generating PDF)  to do that but intead of PDF generate Image [example](http://prawnpdf.org/docs/0.11.1/Prawn/Images.html)

It's just sound to me there must be a better way to do it by now in the community 

any advice ?

**update 30.07.2020 13:00**: 

I didn't specify well enough what will be the content of the generated image. It should be Poem or Story on multiple pages converted to multiple images (representing multiple pages). 

You could say it's like PDF but images as the platform is then providing view option where you cane read the content of the images (poems / stories) 

There will be only one style/template, so no custom templates.

There is a business reason why the end result  must be an image.
## [11][Issue with route](https://www.reddit.com/r/rails/comments/hhyrfd/issue_with_route/)
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
## [12][How to use .includes with a conditional?](https://www.reddit.com/r/rails/comments/hi3mfp/how_to_use_includes_with_a_conditional/)
- url: https://www.reddit.com/r/rails/comments/hi3mfp/how_to_use_includes_with_a_conditional/
---
I'm trying to search by a businesses state. Would this be using a conditional with the .includes statement? Should I be using .where here?   


Here is my current code

`@widget.inventory_units.includes(:business)`

&amp;#x200B;

I'm trying to get my code to do something like 

`@widget.inventory_units.includes(:business).where("business.state = "Alaska")`
