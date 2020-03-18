# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/fgx7fz/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/fgx7fz/personal_projects_show_off_your_own_project_andor/
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
## [2][Taking Rails 6’s Action Mailbox for a Spin](https://www.reddit.com/r/rails/comments/fkmhsu/taking_rails_6s_action_mailbox_for_a_spin/)
- url: https://www.reddit.com/r/rails/comments/fkmhsu/taking_rails_6s_action_mailbox_for_a_spin/
---
In this article the author will take Action Mailbox for a spin and write a simple working Rails 6 app to demonstrate Action Mailbox in action.

Link to article: https://medium.com/merajulislam/taking-rails-6s-action-mailbox-for-a-spin-104d0e34d379
## [3][Google Oauth2 not applying image_size](https://www.reddit.com/r/rails/comments/fkombb/google_oauth2_not_applying_image_size/)
- url: https://www.reddit.com/r/rails/comments/fkombb/google_oauth2_not_applying_image_size/
---
I'm trying to apply image\_size to omniauth-google-oauth2, but everything I tried didn't work.

My devise.rb

    config.omniauth :google_oauth2, google_client_id, google_client_secret, {image_aspect_ratio: "square", image_size: 100}

I'm using my application\_helper to display the image with:

    module ApplicationHelper def avatar_url(user)     user.image   end end

If I check the image\_url it still shows [https://lh3.googleusercontent.com/a-/AOh14Gg9KNrUy\_97f3wDhevAJGE52nT8Y6ho48Uo9vJwHw](https://lh3.googleusercontent.com/a-/AOh14Gg9KNrUy_97f3wDhevAJGE52nT8Y6ho48Uo9vJwHw) without any change in size

I've also tried auth.info.image.sub('?sz=32','?sz=100'), but it didn't work.
## [4][ActionText uploading a gif is being shown as an image](https://www.reddit.com/r/rails/comments/fk8anx/actiontext_uploading_a_gif_is_being_shown_as_an/)
- url: https://www.reddit.com/r/rails/comments/fk8anx/actiontext_uploading_a_gif_is_being_shown_as_an/
---
Trying to upload a gif through action text and it gets uploaded correctly but on update or creation something happens and an image variant is created and is used to display the post  

would appreiciate any help, here is the SO link with a bit more detail

[https://stackoverflow.com/questions/60721539/actiontext-gif-attachment-converts-to-photo-when-submitting-rails-form](https://stackoverflow.com/questions/60721539/actiontext-gif-attachment-converts-to-photo-when-submitting-rails-form)
## [5][Example of Search Form with Google Places Autocomplete in Rails 6?](https://www.reddit.com/r/rails/comments/fkch08/example_of_search_form_with_google_places/)
- url: https://www.reddit.com/r/rails/comments/fkch08/example_of_search_form_with_google_places/
---
Does anyone have a tutorial of an example of Google Places Autocomplete being used in a search form in Rails 6? I am trying to add it to my app, but it looks like most of the resources available are for Rails 3/4.
## [6][Aliasing polymorphic relationships](https://www.reddit.com/r/rails/comments/fkdccv/aliasing_polymorphic_relationships/)
- url: https://www.reddit.com/r/rails/comments/fkdccv/aliasing_polymorphic_relationships/
---
Hi, I have an address model. It has an belongs\_to :addressable relationship which is polymorphic. One of the models that can be an addressable is "property".

I'm aware there's a syntax that allows me to somewhat alias the addressable relationship with an specific relation.

Somewhat like

    belongs_to :addressable, polymorphic: true
    belongs_to :property, as: :addressable, -&gt; {where(addressable_type: 'Property')}

&amp;#x200B;

What's the correct syntax for it? I have done this before but I can't find anything online or figure out where I've done it.
## [7][is rack on the same layer as fastCGI?](https://www.reddit.com/r/rails/comments/fjwgtu/is_rack_on_the_same_layer_as_fastcgi/)
- url: https://www.reddit.com/r/rails/comments/fjwgtu/is_rack_on_the_same_layer_as_fastcgi/
---
Are rack and fastCGI on the same layer solving the same problem and are alternatives to each other, or is rack built on top where fastCGI is?
## [8][Anybody having readline issues with macos](https://www.reddit.com/r/rails/comments/fjux86/anybody_having_readline_issues_with_macos/)
- url: https://www.reddit.com/r/rails/comments/fjux86/anybody_having_readline_issues_with_macos/
---
https://preview.redd.it/nkm2w8kie4n41.png?width=1206&amp;format=png&amp;auto=webp&amp;s=738055f6456cde40f478ed3768d10e925e9f9533

Whats expected is for the number 2 to be printed,  
instead the rails repl prints some garbage after the number 2.
## [9][Can't modify frozen String on form tags after upgrade to Rails 6.0.0](https://www.reddit.com/r/rails/comments/fjsp5w/cant_modify_frozen_string_on_form_tags_after/)
- url: https://www.reddit.com/r/rails/comments/fjsp5w/cant_modify_frozen_string_on_form_tags_after/
---
I'm in the middle of upgrading a monolith from Rails 5.2 to 6.0.0. This upgrade so far has gone pretty easy. I was able to knock out the deprecations and a co-worker of mine was able to fix the test suite (for the most part. Still having an issue with Capybara on 2 tests).

I ran `rails app:update` and manually updated the config files that were updated in Rails 6.0.0. One line I forgot to change was `config.load_defaults '6.0'`, in fact this was still on `config.load_defaults '5.1.'` even though we were on 5.2.

After updating this line to `config.load_defaults '6.0'` and fixing any zeitwerk loading issues (by running `rails zeitwerk:check`) we're getting nearly 200 errors in our test suite.

`ActionView::Template::Error: can't modify frozen String`

Every one of these I've checked (so far) are pointing to a form tag. The main one pointing to a `form_tag` call and a few others pointing to a `semantic_form_for` tag in an older part of our app (we've since upgrade to simple_form).

I've added the configs we were using in the old `new_framework_defaults_5_x` files to application.rb hoping it was a config in one of those but no dice...

Not really sure what to do from here. Any ideas?
## [10][Date default value](https://www.reddit.com/r/rails/comments/fjk0mu/date_default_value/)
- url: https://www.reddit.com/r/rails/comments/fjk0mu/date_default_value/
---
So I have a model `Task` which has an attribute `deadline` of type `datetime`. I also have at some point in my views the following input:

    &lt;input type="date" id="start" name="task[deadline]" min="2020-01-01" max="2025-12-31" value="&lt;%= @task.deadline.strftime("%Y/%m/%d - %H:%M %p") %&gt;"&gt;

Well, apparently strftime does not work, as it shows `mm/dd/yyyy` instead of the corresponding date. I don't know how to fix this, but probably is a simple thing that right now I don't see.
## [11][Standing out as a true Pro Rails Developer](https://www.reddit.com/r/rails/comments/fjedn3/standing_out_as_a_true_pro_rails_developer/)
- url: https://www.reddit.com/r/rails/comments/fjedn3/standing_out_as_a_true_pro_rails_developer/
---
Sorry for the long windedness, but here goes...

In my day job, I work as a JavaScript developer at a DotNet shop.  However, I've loved Rails - and Ruby - ever since I started playing around with them about 8 years ago.

People have always told me "don't waste your time on that, stick with dotnet (or Java), because that's where the money is."

But I honestly don't care about working in large enterprise organizations and I want to do the thing that I enjoy doing in programming.  Over the years I've done lots of Rails tutorials and made small projects.  Followed YouTube videos about various Rails subjects, and gone through books - like Sandi Metz excellent "Practical Object Oriented Design in Ruby".

Over the next few months I would like to transition into my first job as a Rails developer.  But I just have this hunch, some sneaky feeling that "real world" Rails development is much messier and different than what I'm seeing in the tutorials and books that I'm working with.

So I have a few questions for you Rails pros out there who have been in the trenches for a long time.

**What knowledge would I need to know to make me stand out from other Rails newbies?**  What do the tutorials **not** prepare you for?

For instance, in the DotNet world, we've had new developers who had experience working with Entity Framework - the dotnet equivalent of ActiveRecord.

But in most companies I've worked at, nobody actually uses Entity Framework, or any ORM for that matter.  The databases have generally been around for a long time and are super complicated and not very well designed.  To get the data you need, there are vast oceans of hand-made stored procedures and custom designed SQL - a lot of them pointed at OLAP cubes.

So a lot of newbies feel like there's this whole world they weren't prepared for.

In "real world" Rails development, are you really working with a SQLite database in Dev, updating migration files, and just running those against a PostGres or MySql database?  Or is it much more complicated than that?

Just looking for tips of what I need to study at a more advanced level that maybe the tutorials and books aren't preparing me for.
