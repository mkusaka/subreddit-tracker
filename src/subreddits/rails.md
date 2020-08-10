# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/i00rha/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/i00rha/personal_projects_show_off_your_own_project_andor/
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
## [2][Does anyone know what's going on with apidock?](https://www.reddit.com/r/rails/comments/i6tngy/does_anyone_know_whats_going_on_with_apidock/)
- url: https://www.reddit.com/r/rails/comments/i6tngy/does_anyone_know_whats_going_on_with_apidock/
---
Whenever I try to visit apidock.com, it takes forever to load and then i just get a server error.

Is anyone else experiencing this?
## [3][Plyr video player and Rails 6](https://www.reddit.com/r/rails/comments/i74o1y/plyr_video_player_and_rails_6/)
- url: https://www.reddit.com/r/rails/comments/i74o1y/plyr_video_player_and_rails_6/
---
[https://plyr.io/](https://plyr.io/) is a javascript library to play videos with support for YouTube and Vimeo. I'm trying to implement it within a Rails 6 site. What are the steps to make this happen?

Alternatively, what is your choice of video player in Rails 6 that supports MP4, Vimeo, and YouTube?
## [4][The best Ruby/Rails profiler?](https://www.reddit.com/r/rails/comments/i74jms/the_best_rubyrails_profiler/)
- url: https://www.reddit.com/r/rails/comments/i74jms/the_best_rubyrails_profiler/
---
What is the best profiling tool for Ruby on Rails? I'm looking for a tool which could return executed functions stack with function names, number of calls and time elapsed on each function. The goal is to find slow methods which are causing performance issues.
## [5][Turbolinks doesn't update the page to display model validation errors](https://www.reddit.com/r/rails/comments/i6k4qy/turbolinks_doesnt_update_the_page_to_display/)
- url: https://www.reddit.com/r/rails/comments/i6k4qy/turbolinks_doesnt_update_the_page_to_display/
---
Hi!

just a quick question, really stuck here:
I'm new to Rails and I want to display form validation errors to my users.
I have the following in my posts controller:

```ruby
class Blog::PostsController &lt; ApplicationController
  def new
    @post = Post.new
  end

  def create
    @post = Post.new post_params
    if @post.save
      flash[:notice] = 'Post was created successfully!'
      redirect_to :controller =&gt; 'posts', :action =&gt; 'new'
    else
      render 'new'
    end
  end

  private def post_params
    params.require(:post).permit(:name, :slug, :tags, :content)
  end
end
```

Then, I have a view like this:

```erb
&lt;h1&gt;Create a Blog Post&lt;/h1&gt;

&lt;% @post.errors.full_messages.each do |msg|%&gt;
  &lt;li&gt;&lt;%= msg %&gt;&lt;/li&gt;
&lt;% end %&gt;

&lt;%= form_with model: @post, url: blog_post_path do |f| %&gt;
  &lt;%= f.label :name %&gt;
  &lt;%= f.text_field :name %&gt;&lt;br&gt;

  &lt;%= f.label :slug %&gt;
  &lt;%= f.text_field :slug %&gt;&lt;br&gt;

  &lt;%= f.label :tags %&gt;
  &lt;%= f.text_field :tags %&gt;&lt;br&gt;

  &lt;%= f.label :content %&gt;
  &lt;%= f.text_area :content %&gt;&lt;br&gt;

  &lt;%= f.submit %&gt;
&lt;% end %&gt;
```

As you can see right at the top, I try to show the validation errors. But for some reason, it never shows them on the page (Turbolinks doesn't add them!). If I take a look into the JS console, I can see the response the server sent back, including the validation errors. Still, Turbolinks doesn't want to show them. How do I fix this?
## [6][Text field: cannot set default value](https://www.reddit.com/r/rails/comments/i6mqzh/text_field_cannot_set_default_value/)
- url: https://www.reddit.com/r/rails/comments/i6mqzh/text_field_cannot_set_default_value/
---
Hi. I'm trying to set a text field's default value (I'm creating a form to edit existing records).

This is my view:

```erb
&lt;%= form_with model: @post, url: blog_post_path, local: true do |f| %&gt;
  &lt;%= f.label :name %&gt;
  &lt;%= f.text_field :name, value: @post_current['name'] %&gt;&lt;br&gt;

  &lt;%= f.label :slug %&gt;
  &lt;%= f.text_field :slug, value: @post_current['slug'] %&gt;&lt;br&gt;

  &lt;%= f.label :tags %&gt;
  &lt;%= f.text_field :tags, value: @post_current['tags'] %&gt;&lt;br&gt;

  &lt;%= f.label :content %&gt;
  &lt;%= f.text_area :content, value: @post_current['content'] %&gt;&lt;br&gt;

  &lt;%= f.submit %&gt;
&lt;% end %&gt;
```

This is my controller action:

```ruby
  def edit
    if Post.where(slug: params[:slug]).exists?
      @post = Post.new
      @post_current = Post.where slug: params[:slug]
    else
      raise ActiveRecord::RecordNotFound
    end
  end
```

When I try to load the page, it just gives me an error page saying "no implicit conversion of String into Integer" on the first line where I try to set a value (`&lt;%= f.text_field :name, value: @post_current['name'] %&gt;&lt;br&gt;`). Why does it do that and how can I set a value?

Still new to Rails, thank you.
## [7][How to return extra association attributes on a collection](https://www.reddit.com/r/rails/comments/i6xfhn/how_to_return_extra_association_attributes_on_a/)
- url: https://www.reddit.com/r/rails/comments/i6xfhn/how_to_return_extra_association_attributes_on_a/
---
So here is my setup. 

2 models:
* events
* showings

events have many showings and a showing belongs to an event.

users can rsvp to showings, so when I write `current_user.showings`, it will return all the showings that any user has rsvped to.

Now, when I run `@showings = current_user.showings` and **when I loop through `@showings`** I get all the attributes that are in the showings table of the database. This is perfect and expected. Any `showing` record has attributes like this:

`{ start_time, end_time, name, etc... }`

Okay, but **in each showing that I loop through, I also want to have access to the `event.name`**. I.e, I want to see this:

`{ start_time, end_time, name, event_name, etc... }` 

or something like this. As long as it's included in the hash. I know you might be thinking I can just use `showing.event.name` but I can't because that's ruby code. I'm turning the collection of showings (i.e., `@showings`) into json and sending it to javascript for front end processing. 

This means that the event name for each showing has to be included in each showing "hash" so that when I convert it to json, all the data is there.

I know i can just loop through them in the controller, and basically build an array of hashes that are duplicates of the showings, and then include any other relevant info about the event in the hashes, but i'm just wondering if there is any built in rails method for doing this instead of doing it manually.

### this is my manual (seemingly not very elegant) solution

```ruby
showings = Array.new
current_user.showings.each do |showing|
  showings &lt;&lt; {id: showing.id, name: showing.name, start_time: showing.start_time, event_name: showing.event.name}
end
@showings_json = showings.to_json
```
## [8][Rails 6 and the Secret Keys](https://www.reddit.com/r/rails/comments/i68uij/rails_6_and_the_secret_keys/)
- url: https://www.reddit.com/r/rails/comments/i68uij/rails_6_and_the_secret_keys/
---
I am in the process of attempting my first deployment of a new app in Rails 6 and when it goes to do the first db migration I get `ArgumentError: Missing \`secret_key_base\` for 'production' environment, set this string with \`rails credentials:edit\``

I am all for setting this up, but I at this point I am confused as to if this is all something I set up production space or development space. Is there a good walk-through of how and where to set this up? I tend to learn best for seeing something set up in front of me. Any help will be great!
## [9][Deploying a Rails app on Apache2](https://www.reddit.com/r/rails/comments/i689yt/deploying_a_rails_app_on_apache2/)
- url: https://www.reddit.com/r/rails/comments/i689yt/deploying_a_rails_app_on_apache2/
---
Hi 👋

I want to deploy a Rails app on my Apache2 web server. Coming from a PHP background, I'm used to the CGI-like approach of mod_php.

After reading through resources on this topic, I wanted to ask you whether it's a good or at least okay idea to do the following:

1.	Use Puma on the Rails app to start a local server
2.	Reverse proxy that through Apache2

I also read a lot about this Passenger thing, but it seems kinda weird to deploy. Any thoughts?
## [10][Design Patterns and Anti-Patterns in Rails?](https://www.reddit.com/r/rails/comments/i5vg9b/design_patterns_and_antipatterns_in_rails/)
- url: https://www.reddit.com/r/rails/comments/i5vg9b/design_patterns_and_antipatterns_in_rails/
---
OK, it's more like a *software engineering* topic than a rails related one. But I asked one of my friends about deleting a table manually and re-do the migration (the project is written in Django and not rails) and he told me "This is an Anti-Pattern in Django". 

I knew possible dangers of the idea and I suggested it with the knowledge, but I jokingly answered him "You call everything you don't understand an anti-pattern". 

Now, I'm curious, is there a set of patterns and anti-patterns SPECIFICALLY for rails?
## [11][New Video: Devise Profile by User Role](https://www.reddit.com/r/rails/comments/i60iis/new_video_devise_profile_by_user_role/)
- url: https://www.reddit.com/r/rails/comments/i60iis/new_video_devise_profile_by_user_role/
---
In this tutorial video I walk through how to create dynamic profile pages by user role.  We will use a technique I learned from my favorite Ruby on Rails tutorial creators, GoRails, to dynamically display the correct partial by current user role.

I received permission before posting from Chris at GoRails to demonstrate this technique which I have seen a few times, but never in this way.  I thought it may be an interesting application of the technique and also answered a subscriber request.  

I have only created 25 or so videos, so still very new to the process, all feedback is welcome as I just want to improve.  

[https://youtu.be/wbRDqZCchs0](https://youtu.be/wbRDqZCchs0)
