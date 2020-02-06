# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Do you add IDs to association tables?](https://www.reddit.com/r/rails/comments/ezo6rc/do_you_add_ids_to_association_tables/)
- url: https://www.reddit.com/r/rails/comments/ezo6rc/do_you_add_ids_to_association_tables/
---
The way I was taught database normalization, association tables do not need a primary key ID column. Their job is to point to the associated records, and may contain additional data related to the association.

In my case, a user

    has_many :memberships
    has_many :organizations, through: :memberships

and 

    create_table "memberships", id: false, force: :cascade do |t|
      t.bigint "user_id", null: false
      t.bigint "organization_id", null: false
      t.jsonb "profile"

In Rails however, this understandably confuses AR

    o = Organization.first
    u = User.first
    m = u.memberships.where(organization: o).first

    m.profile["favorite_color"] = "green"
    m.save

    # =&gt; ActiveRecord::UnknownPrimaryKey: Unknown primary key for table memberships in model Membership.
    # Cannot validate uniqueness for persisted record without primary key.

Of course I can still do

    updated_profile = m.profile.merge({"favorite_color" =&gt; "green"})
    Memberships.where(organization: o, user: u).update_attributes(profile: updated_profile)

but it feels considerably less Rails-y. Although it feels cleaner given the schema and its intentions.

So, what do you do in your apps and *why*? Any downside in adding primary keys? 

I started without them, because I wanted a clean schema, and now am considering adding primary keys to make life easier but am a bit torn.

Edit: clarity
## [3][Implementing STI. Routing forms is going absolutely crazy. Way out of my depth. Send help.](https://www.reddit.com/r/rails/comments/ezjy6h/implementing_sti_routing_forms_is_going/)
- url: https://www.reddit.com/r/rails/comments/ezjy6h/implementing_sti_routing_forms_is_going/
---
##Edit
For anyone with a similarly dumb problem in the future, /u/colonel_weasel pointed out to check my routes and make 100% sure I was referencing them correctly. I had site_item_tmp_group_path (singular) instead of site_item_tmp_groups_path (plural). "The Magic of using Raaaaaaaaaaiiiiils"

___

##Original Post

Hey all,

As the title says, I'm way out of my depth here and hoping someone can point out the... 3 or 4 or 20 things I'm doing wrong.

So I'm messing around with Single Table Inheritance and.. actually it's not going horribly! Yay! Right now I've got a few models, some of which inherit from others:

    class SiteItem &lt; ApplicationRecord
    belongs_to :site
    
    scope :tmp_groups, -&gt; { where(type: "TmpGroup") }
    scope :tmp_articles, -&gt; { where(type: "TmpArticle") }
    scope :tmp_galleries, -&gt; { where(type: "TmpGallery") }
    scope :tmp_cards, -&gt; { where(type: "TmpCard") }

As you can see, all items belong_to a single Site object, which acts as the owner of all the elements of my future website. The idea is to dynamically add, edit, and remove pages as I feel like it without having to re-code, re-commit, and re-deploy.

Two subclasses inherit from the **SiteItem** model: **Group** and **Page**, and several templates inherit from Page:

* tmp_gallery
* tmp_article
* tmp_card
* tmp_group

Most of these are just slightly different configurations of similar data. For instance, **TmpArticle** and **TmpCard** are nearly identical in the information they store, but are obviously displayed quite differently. TmpGallery is for displaying image galleries with little more than just that, and TmpGroup is used to group assortments of pages together.

Here is what my routes.rb looks like for this stuff btw...

    #SiteItems in Routes.rb
	resources :site_items do
		resources :tmp_groups, controller: :site_items, type: "TmpGroup"
		resources :tmp_articles, controller: :site_items, type: "TmpArticle"
		resources :tmp_galleries, controller: :site_items, type: "TmpGallery"
		resources :tmp_cards, controller: :site_items, type: "TmpCard"
	end

I have SUCCESSFULLY (yay) set things up so that site_items_controller#show and site_items_controller#index both work as expected (I added some test objects to my db in the console). However, creating a new item is giving me a massive headache! Mostly I'm getting a bunch of routing errors that I'm not experienced to figure out yet. They're caused by my trying to create a form for adding any new SiteItems. 

For instance, here's is my form for creating a new TmpGroup:

    &lt;%= form_for(:tmp_group, url: new_site_item_tmp_group_path) do |f| %&gt;
    ...
    &lt;% end %&gt;

This form renders as expected, but filling it out gives me the following:

#No route matches [POST] "/site_items/2/tmp_groups/new"

So my lizard brain just thinks.. Ok, I need to post to **site_item_tmp_group_path**, right? So I'll just remove the **new_** from my form url. But now when I go to load the create new group form page, rails throws a new error:

#ActionController::UrlGenerationError in SiteItems#new
##No route matches {:action=&gt;"show", :controller=&gt;"site_items", :site_item_id=&gt;"2", :type=&gt;"TmpGroup"}, missing required keys: [:id]

Anyway, if you made it this far you can probably tell I don't know what I'm doing. The good news is I badly want to, and would love some help getting there! Thanks for even reading!
## [4][Affordable transactional email sending options with high deliverability for a heroku hosted app?](https://www.reddit.com/r/rails/comments/eze069/affordable_transactional_email_sending_options/)
- url: https://www.reddit.com/r/rails/comments/eze069/affordable_transactional_email_sending_options/
---
Asking this because just had to bump my Sendgrid plan to include a dedicated IP. Without it, deliverability was not acceptable -- lots of blacklisted IPs a lot of the time. 

The dedicated IP plan is $90/month and that's a lot for a tiny operation running a beta MVP heavily relying on email messaging. 

Any suggestions for a more affordable but still reliable transactional email provider offering a dedicated IP and ideally offering inbound parsing webhooks (but also without that feature as it's a distant milestone)?

Ease of implementation also matters. I looked at Amazon SES but as a single dev I just can't prioritize going through that maze. If you have a good set of instructions for that, it would be helpful, too.

Apart from the host name, please include any other transactional email-related tips. 

Thank you!

Edit: based on feedback here, I'll go with Postmark at the end of the month. Thank you for your help.
## [5][How to make a relationship with Spree::User?](https://www.reddit.com/r/rails/comments/ezh6gv/how_to_make_a_relationship_with_spreeuser/)
- url: https://www.reddit.com/r/rails/comments/ezh6gv/how_to_make_a_relationship_with_spreeuser/
---
  am trying to build a friends feature on top of the Solidus framework but am having trouble establishing a many-many relationship with Spree::Users. I tried making a user\_decorator.rb file (in models/spree) but keep encountering the error: "expected user\_decorator.rb to define Spree::UserDecorator, but didn't". 

&amp;#x200B;

https://preview.redd.it/on1teif2f6f41.png?width=609&amp;format=png&amp;auto=webp&amp;s=2d612000b3f2e3190a9b887329b2d127fe52cf47

&amp;#x200B;

&amp;#x200B;

https://preview.redd.it/qasjwd64f6f41.png?width=739&amp;format=png&amp;auto=webp&amp;s=0dabff7e3c0852d1be94b4b4fa098ad75f9ca517

 

Error:

[https://i.stack.imgur.com/QOpEU.png](https://i.stack.imgur.com/QOpEU.png)
## [6][undefined method `friends_id' for #&lt;Spree::User:0x00007f46c89d1bb8&gt;](https://www.reddit.com/r/rails/comments/ezcg1h/undefined_method_friends_id_for/)
- url: https://www.reddit.com/r/rails/comments/ezcg1h/undefined_method_friends_id_for/
---
I am trying to implement a friends feature on top of the Solidus framework, but the ".friends" is not working on rails server, however it is working on rails console. Does anyone know how to fix this?

&amp;#x200B;

[spree\_users\_controller.rb code](https://preview.redd.it/1cxghj7gx4f41.png?width=706&amp;format=png&amp;auto=webp&amp;s=9a34af2f9fb9b1a02ce891a78eeedd61d8016b99)

&amp;#x200B;

[friendship.rb code](https://preview.redd.it/70xm4k4jx4f41.png?width=540&amp;format=png&amp;auto=webp&amp;s=5ce85ec8ec8903f5cfb3c1cd28f2950b2062c18c)

&amp;#x200B;

[spree\_user.rb code](https://preview.redd.it/05yl5xlox4f41.png?width=414&amp;format=png&amp;auto=webp&amp;s=047d3a3a1f7503fb06a4c7aafd7790ea3a7f5d72)
## [7][Persisting URL Parameters when Form Validation Errors are encountered](https://www.reddit.com/r/rails/comments/ez86ra/persisting_url_parameters_when_form_validation/)
- url: https://www.reddit.com/r/rails/comments/ez86ra/persisting_url_parameters_when_form_validation/
---
Hello, 

I have a link in which I pass a URL param when routing to the Create Form.  On the form I have a hidden field where I am storing the URL param.  If validation errors are raised, the form redirects to /stash\_entries instead of /stash\_entries/new and loses the URL param.  I am missing something here but not sure what, any ideas?

Link_to:
```
            &lt;%= link_to new_account_stash_stash_entry_path(stash_id: stash.id, stash_action: 'add'), class: "button is-small is-success" do %&gt;
              &lt;span class="icon"&gt;
                &lt;i class="fa fa-lg fa-plus"&gt;&lt;/i&gt;
              &lt;/span&gt;
            &lt;% end %&gt;
```

Form:
```
&lt;%= form_for([@stash_entry.stash.account, @stash_entry.stash, @stash_entry]) do |f| %&gt;
  &lt;% if @stash_entry.errors.any? %&gt;
    &lt;article id="error explanation" class="message is-danger"&gt;
      &lt;div class="message-header"&gt;
        Something went wrong!
      &lt;/div&gt;
      &lt;div class="message-body content"&gt;
        Please correct the following &lt;span class="has-text-weight-bold"&gt;&lt;%= pluralize(@stash_entry.errors.count, "error") %&gt;&lt;/span&gt;:
        &lt;ul&gt;
          &lt;% @stash_entry.errors.full_messages.each do |msg | %&gt;
            &lt;li&gt;
              &lt;%= msg %&gt;
            &lt;/li&gt;
          &lt;% end %&gt;
        &lt;/ul&gt;
      &lt;/div&gt;
    &lt;/article&gt;
  &lt;% end %&gt;

  &lt;div class="columns"&gt;
    
    &lt;!-- Fields --&gt;
    &lt;div class="column"&gt;

      &lt;!-- Action --&gt;
      &lt;div class="field"&gt;
          &lt;%= f.hidden_field :stash_action, value: params[:stash_action] %&gt;
      &lt;/div&gt;

      &lt;!-- Amount --&gt;
      &lt;div class="field"&gt;
        &lt;label class="label has-text-grey"&gt;Amount&lt;/label&gt;
        &lt;div class="control has-icons-left"&gt;
          &lt;%= f.text_field :amount, class: "input is-primary",
                                         placeholder: "Name" %&gt;
          &lt;span class="icon is-small is-left has-text-primary"&gt;
            &lt;i class="fa fa-file-text"&gt;&lt;/i&gt;
          &lt;/span&gt;
        &lt;/div&gt;
      &lt;/div&gt;
    &lt;/div&gt; &lt;!-- /column --&gt;

  &lt;/div&gt; &lt;!-- /columns --&gt;

  &lt;div class="field is-grouped"&gt;
    &lt;div class="control"&gt;
      &lt;%= link_to "Cancel", account_stashes_path, class: "button is-link is-outlined" %&gt;
    &lt;/div&gt;
    &lt;div class="control"&gt;
      &lt;%= f.button :stash_action,
          type: 'submit',
          class: "button is-link" %&gt;
    &lt;/div&gt;
  &lt;/div&gt;

&lt;% end %&gt;
```

Controller:
```
  # GET /stashes/new
  def new
    @stash = @account.stashes.build.decorate
  end

  # GET /stashes/1/edit
  def edit
  end

  # POST /stashes
  def create
    @stash = @account.stashes.build(stash_params).decorate

    if @stash.save
      redirect_to account_stashes_path, notice: 'Stash was successfully created.'
    else
      render :new
    end
  end
```

Routes:
```
  resources :accounts do
    resources :transactions
    resources :stashes do
      scope except: %i[index show edit update destroy] do
        resources :stash_entries
      end
    end

    member do
      get :deactivate
      get :activate
    end
  end
```

Thanks!
## [8][My first rails app deployed! (So excited, lol)](https://www.reddit.com/r/rails/comments/eyyffw/my_first_rails_app_deployed_so_excited_lol/)
- url: https://www.reddit.com/r/rails/comments/eyyffw/my_first_rails_app_deployed_so_excited_lol/
---
Omg i am so excited (dont hate me for this one)

I started with rails 2 monts ago  (did some javascript before and hate it now so much) and decided that i wanted to learn rails.

After one udemy course (dissecting ruby) by jordan (forgot his lastname) i decided to build something. This whole thing took me 4 days to build.

Functions:
User can signup and upload a video (for now it is a image, bc of heavy video resources)  and other users can comment on his form. It can be used for personal trainers to check out of their client is doing the excerice in the right form.

I created this webapp in 4 days, i know i know it is just a extended crud app.

- Users can comment on all posts
- User can only edit or delete own posts
- Only Admin can remove or edit all posts
- User can update bio, username, email
- User sees in his account all his own posts

Started out with sqlite and changed to postgress to be able to deploy to heroku

Check it out!

https://vormcheck.herokuapp.com

(Site is not mobile friendly)

Github https://github.com/sljmn/vormcheck

————
- Now i want to learn how to use helpers 
- How to use has_many through
- How to add more interactivity. This app is just a crud app. But i want to learn how to let users follow each other or like each others post
- How to consume an api and display it in the rails app

If you know a source, it would be appriciated
## [9][Early access program for Ruby on Rails + DevOps training](https://www.reddit.com/r/rails/comments/ezcpzr/early_access_program_for_ruby_on_rails_devops/)
- url: https://www.reddit.com/r/rails/comments/ezcpzr/early_access_program_for_ruby_on_rails_devops/
---
Hi,

we are a very early stage startup working towards making learning more feasible and integrated for working professionals who wants to keep up with the changing pace of technology and stand out among their peers. 

We believe that learning should be based on **practical relevance** at work, **emerging skill trends** in your field, **access to all the online resources** that are out there and much more. 

While we are working on a digital product that offers such a learning experience, we want to provide the same experience offline, personally. This is because we want to build a very user-centric product with the best user experience. To build such a digital product, we need understand the users’ needs and pain points a lot better and we want to gain that experience through you. We are offering this training for free so that you gain the skills you require and give us feedback during the process for us to improve.

With that said, we want to offer personalised training for you to advance your career in Ruby on Rails along with DevOps expertise. 

We will be in direct communication with you and provide you the resources that you will need to learn the skills you need to excel in these fields! You will have an expert available for your questions and a community of fellow learners to discuss on the topics. You will cover topics like Meta-programming, Redis, Money-patching, AWS, Terraform, CI/CD, Nginx, Kibana, SSL certificate management and much more.

Take a look at the offering and join us if you think you would benefit from these. 

[https://www.eduup.de/en/program?utm\_campaign=referral-reddit-dev&amp;utm\_source=referral&amp;utm\_content=dev&amp;utm\_medium=reddit](https://www.eduup.de/en/program?utm_campaign=referral-reddit-dev&amp;utm_source=referral&amp;utm_content=dev&amp;utm_medium=reddit)
## [10][To move 500 terabytes from S3 to Glacier, is Multipart Upload or ZIP better?](https://www.reddit.com/r/rails/comments/ez1czb/to_move_500_terabytes_from_s3_to_glacier_is/)
- url: https://www.reddit.com/r/rails/comments/ez1czb/to_move_500_terabytes_from_s3_to_glacier_is/
---
Hey Rubyists and RailsHeads,

We are going to save FIVE THOUSAND DOLLARS A MONTH by moving these 500 terabytes from S3 to Glacier. However, it seems kind of dumb for me to put all of the photos and videos in their own archives on glacier, because that means both more requests to archive them, and because you'd want the files to be logically grouped by user. I work for an unpopular social media site.

So, is it smarter to take all of the user's photos and video's and zip them before putting them on the archive? Seems like it would add a significant amount of time/computing power for the process of archiving the files.

Or is it better to use the multipart upload method in the aws ruby sdk v3, and dump them all into the a single archive? Seems... Just kind of complicated and annoying with the checksum and those kinds of uploads are unfamiliar to me.

The potential savings for grouping the S3 files in either of the above ways is... Like $300 for moving the files initially due to fewer requests, but I'm not sure if it will pay for itself in developer time or happiness lol. Also will make putting the files back on S3 a lot harder.
## [11][Rails Deployment Tutorial updated for Ubuntu 18.04 LTS / Debian 10.2](https://www.reddit.com/r/rails/comments/eyxc25/rails_deployment_tutorial_updated_for_ubuntu_1804/)
- url: https://www.reddit.com/r/rails/comments/eyxc25/rails_deployment_tutorial_updated_for_ubuntu_1804/
---
I just updated the Rails Deployment Tutorial for *Ubuntu 18.04 LTS* and *Debian 10.2:*

[https://www.ralfebert.de/tutorials/rails-deployment/](https://www.ralfebert.de/tutorials/rails-deployment/)

Have fun deploying your apps :)

If you encounter any issues using these steps, please let me know...
