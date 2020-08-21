# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/id911w/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/id911w/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][Best stack for a new SaaS side project?](https://www.reddit.com/r/rails/comments/ido300/best_stack_for_a_new_saas_side_project/)
- url: https://www.reddit.com/r/rails/comments/ido300/best_stack_for_a_new_saas_side_project/
---
I'm currently planning out a SaaS side project and I'm trying to figure out the best stack for a modern Rails app. I've done small projects with both a monolith and Rails APIs, so I'm really fine with either approach. One goal here is to up my frontend game. The options I'm considering are as follows:

* 1. Rails API with a separate React frontend
   * Pros: I've done this sort of stack before; API is relatively quick to set up and deploying a separate frontend to Netlify can be convenient.
   * Cons: There are some security concerns with a JWT auth solution, which is what I usually use for APIs. Additionally, I'll be implementing multi-tenancy/admin roles and views, so I'm not exactly sure how well that sort of thing would work with a decoupled frontend/backend.
* 2. Monolith Rails + React
   * Pros: Security would concern me less with a full-on Devise/sessions set up. No security solution is perfect, but I'd feel better about my ability to set up auth in a monolith vs putting together a JWT solution.
   * Cons: I'm not sure I'm 100% sold on the idea of using React within Rails; it would require a bit of set up but I know people do it all the time, so maybe I'm overthinking it. Using a robust FE library like that within Rails would just be a new thing, I guess.
* 3. Monolith Rails + StimulusReflex
   * Pros: I'm super interesting in learning StimulusReflex. Maybe it's just HEY-related hype, but I like the way HEY's app works and it seems like a great way to get a reactive frontend without React or \[insert other framework/library here\].
   * Cons: TBH, it's kind of hard to find stuff about SR beyond basic examples since it's relatively new. I'm also concerned about using CableReady, since my understanding is that disconnections can cause issues(?). I also haven't seen that many posts about SR on Reddit, so I'd be curious to get thoughts on the notion of using SR in a possible production environment.
* 4. Pure Monolith Rails with good ol' ERB.
   * Eh, it works. I've done it before. Gets the job done.

I'd like to take my Rails skills further and try something new, so I'm really leaning towards Rails + StimulusReflex; that seems to be a bit of a wave in terms of the future direction of Rails FE development, but I'd love to get some thoughts from other folks! I'm a relatively new Rails developer (been using Rails for about a year) and don't currently use Rails in the context of my job, so I'm sure there's stuff I'm missing/not taking into consideration.

TL;DR: If you were starting a brand new SaaS side project in Rails and weren't tied down stack-wise, which direction would you go for a modern Rails app?

Edit: some formatting/fixing a sentence or two
## [4][JSON API specification for rails api](https://www.reddit.com/r/rails/comments/idppfn/json_api_specification_for_rails_api/)
- url: https://www.reddit.com/r/rails/comments/idppfn/json_api_specification_for_rails_api/
---
Which one do you guys use or prefer using for rails api response in json api specification.

1. [fast_jsonapi](https://github.com/Netflix/fast_jsonapi)    now forked and mainted as      [jsonapi-serializer](https://github.com/jsonapi-serializer/jsonapi-serializer) 

2. [jsonapi-rb](https://github.com/jsonapi-rb/jsonapi-rb) (latest activity 3 years ago)

3. Others( would be great if you specify :) )
## [5][Follower/Following system using devise](https://www.reddit.com/r/rails/comments/idremq/followerfollowing_system_using_devise/)
- url: https://www.reddit.com/r/rails/comments/idremq/followerfollowing_system_using_devise/
---
I searched on the internet, I probably found 10 or more articles, saying the same thing and most of them are incomplete. 

Follower/Following system is important in most cases nowadays, a lot of employers think their website might look much fancier with this ability. Anyway, I'm not here to discuss how employers annoy me, I'm just curious, isn't there any easy and understandable way to implement follower/following system in a rails app?
## [6][Javascript Library loading twice in Rails using Turbolinks](https://www.reddit.com/r/rails/comments/idfohv/javascript_library_loading_twice_in_rails_using/)
- url: https://www.reddit.com/r/rails/comments/idfohv/javascript_library_loading_twice_in_rails_using/
---
My javascript library is loading twice. I have consulted..

 [https://stackoverflow.com/questions/34388869/javascript-library-loads-twice-in-rails](https://stackoverflow.com/questions/34388869/javascript-library-loads-twice-in-rails) (taking include\_tag out of application.html.erb)

 [https://github.com/turbolinks/turbolinks/issues/403](https://github.com/turbolinks/turbolinks/issues/403) (changing  &lt;%= javascript\_include\_tag 'application', 'data-turbolinks-track': 'reload' %&gt; in my application.html.erb file to  &lt;%= javascript\_include\_tag 'application', 'data-turbolinks-track': 'false' %&gt; )

&amp;#x200B;

Has anyone come along this issue, and have a solution?
## [7][Redirect after sign in using devise?](https://www.reddit.com/r/rails/comments/idcr94/redirect_after_sign_in_using_devise/)
- url: https://www.reddit.com/r/rails/comments/idcr94/redirect_after_sign_in_using_devise/
---
Hi Folks,
Building an ecommerce application and when someone attempts to add something to their cart I want to force them to sign in/sign up. Using devise for this, but currently they get routed back to the main home page. Would prefer that they get routed back to the item they were looking at before (I'm storing the URL in session variable). Any thoughts on how I should modify the basic devise controllers to accomplish this?
## [8][Active Storage is not saving attached files - why?](https://www.reddit.com/r/rails/comments/idgp2a/active_storage_is_not_saving_attached_files_why/)
- url: https://www.reddit.com/r/rails/comments/idgp2a/active_storage_is_not_saving_attached_files_why/
---
Hi!

Encountered a problem where Active Storage won't for the sake of it save uploaded files.

Controller:

```ruby
  def create
    @post = Post.new(permit_post())

    if @post.save
      flash[:success] = 'Post was uploaded!'
      redirect_to post_path(@post)
    else
      flash[:error] = @post.errors.full_messages
      redirect_to new_post_path
    end
  end

  private def permit_post
    params.permit(:post).permit :image, :description
  end
```

Model:

```ruby
class Post &lt; ApplicationRecord
  has_one_attached :image
end
```

View:

```erb
&lt;%= form_for @post, html: { multipart: true } do |f| %&gt;
&lt;%= f.text_area :description, placeholder: 'Some informational description' %&gt;&lt;br&gt;
&lt;%= f.file_field :image %&gt;&lt;br&gt;
&lt;%= f.submit %&gt;
&lt;% end %&gt;
```

Maybe I missed something? It does create the individual posts, but doesn't upload the files to my storage directory (all default) and doesn't show any attached files in the database.

Ran `rails activestorage:install` and migrated already. Thank you very, very much in advance!
## [9][Handling Multiple Points of Relation](https://www.reddit.com/r/rails/comments/idn1zs/handling_multiple_points_of_relation/)
- url: https://www.reddit.com/r/rails/comments/idn1zs/handling_multiple_points_of_relation/
---
Hey everyone!

I would appreciate your help with this. Not really sure what to look for on StackOverflow. My Googling hasn't led me far. 

&amp;#x200B;

*Example:*

**I have the following tables:** 

1. Records. Like vinyl records. 
2. Storage Locations. Think: A shelf or a cabinet. 
3. Boxes: A box that a vinyl record may or may not be in.

**Some parameters:** 

A box must have a storage location. 

A vinyl record may or may not be in a box. If it is not in a box, it is just sitting directly in the storage location, e.g. directly in the cabinet. 

&amp;#x200B;

**So . . .**

This means that I have two possible points of relation b/w records &amp; storage locations. In some cases they relate directly, and in others, they relate through the box. 

&amp;#x200B;

**Are there best practices for handling this type of relationship?** Right now, I've just been doing a lot of "*if* [record.box](https://record.box) *then* record.box.storage\_location *else* record.storage\_location" statements. 

Thoughts?

Any articles or best practices would be super helpful! Thanks :-D
## [10][Rails pagination with in_groups_of using Activerecord query](https://www.reddit.com/r/rails/comments/idmovf/rails_pagination_with_in_groups_of_using/)
- url: https://www.reddit.com/r/rails/comments/idmovf/rails_pagination_with_in_groups_of_using/
---
I'm hoping to use [in\_groups\_of](https://apidock.com/rails/v5.2.3/Array/in_groups_of) for my ActiveRecord query, but this doesn't seem to work well with Kaminari pagination.

I want to do something like the below where I split all the records equally (e.g 3), and view only one part of the records (e.g first third) along with pagination:

    @document_texts = Text.where(active: true).in_groups_of(3)[0].page(params[:page])

EDIT:

It looks like in\_groups\_by isn't what I need. What I want to do is split 300 records say by 3 (or 4 or 5) and be able to access the first 100, second 100, or third 100 with pagination.
## [11][Javascript library loading twice](https://www.reddit.com/r/rails/comments/idfbm9/javascript_library_loading_twice/)
- url: https://www.reddit.com/r/rails/comments/idfbm9/javascript_library_loading_twice/
---
I am using Turbolinks and my javascript library is loading twice. I have consulted..

 [https://stackoverflow.com/questions/34388869/javascript-library-loads-twice-in-rails](https://stackoverflow.com/questions/34388869/javascript-library-loads-twice-in-rails) 

&amp;#x200B;

I have also tried changing  'data-turbolinks-track': 'false' as suggested here.  [https://github.com/turbolinks/turbolinks/issues/403](https://github.com/turbolinks/turbolinks/issues/403) 

&amp;#x200B;

Has anyone run into this and have a solution?
## [12][Rails on Heroku: Guide to how many dynos and which size](https://www.reddit.com/r/rails/comments/icnl4q/rails_on_heroku_guide_to_how_many_dynos_and_which/)
- url: https://www.reddit.com/r/rails/comments/icnl4q/rails_on_heroku_guide_to_how_many_dynos_and_which/
---
I just published an [exhaustive and opinionated guide](https://railsautoscale.com/how-many-dynos/) to dynos on Heroku. It answers the questions I've been hearing over and over for years:

* How many dynos should you be running?
* Which dyno type is right for your app?

I hope you find it helpful!
