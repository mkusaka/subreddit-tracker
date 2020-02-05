# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/evmx0w/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/evmx0w/personal_projects_show_off_your_own_project_andor/
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
## [3][Persisting URL Parameters when Form Validation Errors are encountered](https://www.reddit.com/r/rails/comments/ez86ra/persisting_url_parameters_when_form_validation/)
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
## [4][My first rails app deployed! (So excited, lol)](https://www.reddit.com/r/rails/comments/eyyffw/my_first_rails_app_deployed_so_excited_lol/)
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
## [5][To move 500 terabytes from S3 to Glacier, is Multipart Upload or ZIP better?](https://www.reddit.com/r/rails/comments/ez1czb/to_move_500_terabytes_from_s3_to_glacier_is/)
- url: https://www.reddit.com/r/rails/comments/ez1czb/to_move_500_terabytes_from_s3_to_glacier_is/
---
Hey Rubyists and RailsHeads,

We are going to save FIVE THOUSAND DOLLARS A MONTH by moving these 500 terabytes from S3 to Glacier. However, it seems kind of dumb for me to put all of the photos and videos in their own archives on glacier, because that means both more requests to archive them, and because you'd want the files to be logically grouped by user. I work for an unpopular social media site.

So, is it smarter to take all of the user's photos and video's and zip them before putting them on the archive? Seems like it would add a significant amount of time/computing power for the process of archiving the files.

Or is it better to use the multipart upload method in the aws ruby sdk v3, and dump them all into the a single archive? Seems... Just kind of complicated and annoying with the checksum and those kinds of uploads are unfamiliar to me.

The potential savings for grouping the S3 files in either of the above ways is... Like $300 for moving the files initially due to fewer requests, but I'm not sure if it will pay for itself in developer time or happiness lol. Also will make putting the files back on S3 a lot harder.
## [6][Rails Deployment Tutorial updated for Ubuntu 18.04 LTS / Debian 10.2](https://www.reddit.com/r/rails/comments/eyxc25/rails_deployment_tutorial_updated_for_ubuntu_1804/)
- url: https://www.reddit.com/r/rails/comments/eyxc25/rails_deployment_tutorial_updated_for_ubuntu_1804/
---
I just updated the Rails Deployment Tutorial for *Ubuntu 18.04 LTS* and *Debian 10.2:*

[https://www.ralfebert.de/tutorials/rails-deployment/](https://www.ralfebert.de/tutorials/rails-deployment/)

Have fun deploying your apps :)

If you encounter any issues using these steps, please let me know...
## [7][Rails has_many through association: query or scope condition on join table?](https://www.reddit.com/r/rails/comments/eyywgz/rails_has_many_through_association_query_or_scope/)
- url: https://www.reddit.com/r/rails/comments/eyywgz/rails_has_many_through_association_query_or_scope/
---
I have three tables using a has\_many through set up like: User--&gt;Membership&lt;--Group, so users can be in many groups, and groups can have many users.

In the membership join table I've added a "type" column that will hold static membership types like "owner", "admin", "standard" for example. What I would like to do is set up associations or scopes so that I could get all group.standard\_members and it would return all users in the group that have membership type == "standard". Then I would also like to be able to use group.owner to return the one user in the group with membership type == "owner, as well as a number of related queries from the user model side of things.

Basically it's set up just like the example in the rails guides for a has\_many through association if that makes is easier to provide an example for: [https://guides.rubyonrails.org/association\_basics.html#the-has-many-through-association](https://guides.rubyonrails.org/association_basics.html#the-has-many-through-association)

Note the "appointment\_date" field in the join table "appointments", that is very similar to my "type" field in my "memberships" join table.

So, if using the rails guide example how could I set up a scope or model association so that I could return all patients with an appointment\_date that is today(where appointment\_date: [Time.new](https://Time.new) kind of thing) using active record like \`@physician.todays\_patients\`, also I would like to be able to do the reverse, where I could get a list of all physicians a patient is seeing today using something like \`@patient.todays\_physicians\`.

I'm not concerned about the query conditions so much, as I'm really trying to figure out how to setup scopes or associations that query related records(in a has\_many through association) but make the association by setting conditions on the join tables data.

Does that make sense? Any help or just a push in the right direction would be much appreciated as my searches aren't turning up much relevant info.

Here's my three models, and my latest feeble attempt at one of the associations I was trying to make work \`@group.owner\` and \`@user.owned\_groups\` :

    class User &lt; ApplicationRecord
      # Include default devise modules. Others available are:
      # :confirmable, :lockable, :timeoutable, :trackable and :omniauthable
      devise :database_authenticatable, :registerable,
             :recoverable, :rememberable, :validatable
    
      # user can be a member of many groups through memberships
      has_many :memberships
      has_many :groups, through: :memberships
      
      # User can have many owned groups
      # has_many :owned_groups, class_name: "Group", -&gt; { joins(:memberships).where( membership: { type: "owner" } ) }
    
    ...
    end
    
    ###############
    
    class Membership &lt; ApplicationRecord
      belongs_to :user, inverse_of: :memberships
      belongs_to :group, inverse_of: :memberships
    end
    
    ###############
    
    class Group &lt; ApplicationRecord
      
      # Group can have many members, through memberships for when extra membership data is needed
      has_many :memberships 
      has_many :members, through: :memberships, source: :user # instead of group.users use group.members
        
      # Group will only have ONE owner
      # belongs_to :owner, -&gt; { joins(:memberships).where( {memberships: {type: "owner} } ) }
      
    ...
    end
## [8][How to get the batch iteration number with find_each?](https://www.reddit.com/r/rails/comments/eyusso/how_to_get_the_batch_iteration_number_with_find/)
- url: https://www.reddit.com/r/rails/comments/eyusso/how_to_get_the_batch_iteration_number_with_find/
---
I'm using [find\_each](https://api.rubyonrails.org/classes/ActiveRecord/Batches.html#method-i-find_each) to do a batch request. I want to sleep for some time for each batch iteration (the size of the batch is 500 so every 500 record iterations). Any idea how to get the current batch iteration number?

`Author.find_each(batch_size: 500).with_index do |author, index|`  
 `# sleep for each batch iteration`  
 `end`
## [9][Installing Apartment gem on an existing app - migrating existing data to a new schema](https://www.reddit.com/r/rails/comments/ez02aa/installing_apartment_gem_on_an_existing_app/)
- url: https://www.reddit.com/r/rails/comments/ez02aa/installing_apartment_gem_on_an_existing_app/
---
Hello!
I'm a group guitar teacher and I created this app that (in short) allows me to track peoples sessions (if they've paid up front, how many sessions they have left)

I've been using it for a while, and I showed it to a friend who asked if they could try it out for their classes.

When I originally made it, I made it only for me without any intention of sharing it, and now allowing people to have their own students and sessions seems like a great little project.

I've never used Apartment before, and I've been messing around with it for a week, and I feel like I'm close, I just can't quite get what I want to work. 

**Problem:** My existing data. There are about 6 tables or so with quite a bit of data and they are all on the 'public' schema. I want to create a new schema 'seshna' for me (seshna.myapp.com), make all of those tables a part of that schema, except for the 'users' table, which I've excluded so that it can be accessed from all apartments.

How do I move that existing data from those tables to a new schema? I'll keep trying different approaches, if I figure it out I'll share it, but some help would really be appreciated!
## [10][Update on Building an Accounting (Suggestions)](https://www.reddit.com/r/rails/comments/ez1jhz/update_on_building_an_accounting_suggestions/)
- url: https://www.reddit.com/r/rails/comments/ez1jhz/update_on_building_an_accounting_suggestions/
---
Hi

I recently posted here with regards to a video series in developing a Rails based application from scratch (an accounting system). It's not a tutorial type series as it is more of a reality show where I attempt to capture everything from mistakes to looking up stuff on the web in case I get stuck so it's really quite boring unless you want to see how developers suffer when running into a problem (there's even a video session where I play Eve Online while trying to work out some css lol). But in any case, I already got as far as creating a Trial Balance report and looking into doing another report, probably General Ledger. Looking for comments and recommendations for what else I can put in. From the top of my head, thinking of doing things in the domain of:

* Multi-company segregation of data
* Materalized Views
* ActionCable
* ActiveStorage (where applicable in accouting)

For those interested:

[Video Playlist] (https://www.youtube.com/playlist?list=PL2-7U6BzddIZ35bJdCFx6RZ-QR8n_JD82)

[Source Code] (https://github.com/ralampay/bookkeeper)

Regards and happy coding!
## [11][Debugging story: Mysteriously truncated timestamps with ActiveRecord](https://www.reddit.com/r/rails/comments/eypj02/debugging_story_mysteriously_truncated_timestamps/)
- url: https://www.reddit.com/r/rails/comments/eypj02/debugging_story_mysteriously_truncated_timestamps/
---
Saving [\#Ruby](https://www.linkedin.com/feed/hashtag/?highlightedUpdateUrns=urn%3Ali%3Aactivity%3A6628670995878289408&amp;keywords=%23Ruby&amp;originTrackingId=FZCWQu%2BAcSd9sZ%2F%2BpT1cog%3D%3D) objects in [\#PostgreSQL](https://www.linkedin.com/feed/hashtag/?highlightedUpdateUrns=urn%3Ali%3Aactivity%3A6628670995878289408&amp;keywords=%23PostgreSQL&amp;originTrackingId=FZCWQu%2BAcSd9sZ%2F%2BpT1cog%3D%3D)  is not a rocket science, and there is little that could surprise me. Have you ever thought that? Well, we had too until we tried to debug a flaky test.     


Read the whole story: [https://www.toptal.com/ruby-on-rails/timestamp-truncation-rails-activerecord-tale](https://www.toptal.com/ruby-on-rails/timestamp-truncation-rails-activerecord-tale) to learn why saving timestamps may cause headache :)
## [12][How would you go about using insert_all while ensuring params are white-listed?](https://www.reddit.com/r/rails/comments/eyrrxs/how_would_you_go_about_using_insert_all_while/)
- url: https://www.reddit.com/r/rails/comments/eyrrxs/how_would_you_go_about_using_insert_all_while/
---
I want to save multiple new objects in one call. I recently came across the class method [insert\_all](https://apidock.com/rails/v6.0.0/ActiveRecord/Persistence/ClassMethods/insert_all) that seems to accomplishes this:

    result = Article.insert_all([{...}, {...}])

I was wondering if there was a way to white-list the possible params in the same way you'd do a single save?
