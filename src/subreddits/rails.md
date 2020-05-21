# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gnbebg/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gnbebg/personal_projects_show_off_your_own_project_andor/
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
## [3][Any free Rails hosting options that don't use Postgres?](https://www.reddit.com/r/rails/comments/gnwlmf/any_free_rails_hosting_options_that_dont_use/)
- url: https://www.reddit.com/r/rails/comments/gnwlmf/any_free_rails_hosting_options_that_dont_use/
---
I simply cannot get Postgres to play ball on Windows 10, and I need a free hosting option for pitching/testing purposes. I know it's far from ideal, but are there any free hosting services that will accept Rails with Sqlite?
## [4][ActiveJob::SerializationError on create model for uploading](https://www.reddit.com/r/rails/comments/gnwl80/activejobserializationerror_on_create_model_for/)
- url: https://www.reddit.com/r/rails/comments/gnwl80/activejobserializationerror_on_create_model_for/
---
i am trying to do background job for uploading pics ,

the job something like this:

```
  def perform(model, new_image, file)
    model.meta_gallery.attach(
          io: new_image, 
          filename: file.original_filename, 
          content_type:   file.content_type
     )
  end
```

what I got is `ActiveJob::SerializationError (Unsupported argument type: Tempfile):`

I called that job in my controller like this:
```
 params[:gallery][:files].each do |file|
      new_image = helpers.upload_resized_image(file.tempfile, 400, 400)
      UploadJob(@model, new_image, file)
 end

```
I tried to change: file.tempfile.path as well but doesn't work
does anyone has same issue before like this?
## [5][Authenticate wordpress users using Devise](https://www.reddit.com/r/rails/comments/gnvfiu/authenticate_wordpress_users_using_devise/)
- url: https://www.reddit.com/r/rails/comments/gnvfiu/authenticate_wordpress_users_using_devise/
---
I have a Rails API and I'm currently building a wordpress site.  When users sign in on the wordpress site, I'd like to authenticate their credentials using Devise and my Rails API.  Are there any solutions for this type of authentication?  Wordpress offers plugins that utilize [OAuth single sign on](https://wordpress.org/plugins/miniorange-login-with-eve-online-google-facebook/) but I'm not sure how to integrate that with Devise.  

Any ideas on how to solve this?
## [6][Give your Rails app a fast GraphQL API without writing any code.](https://www.reddit.com/r/rails/comments/gni988/give_your_rails_app_a_fast_graphql_api_without/)
- url: https://www.reddit.com/r/rails/comments/gni988/give_your_rails_app_a_fast_graphql_api_without/
---
I run a popular Rails App (on App Engine) it gets significant traffic (movnorth.com). Recently I started introducing React into the frontend and also decided to start using GraphQL in the backend for queries. This started me down a rabbit hole and here is the result.

I built Super Graph an automatic GraphQL to SQL translator (in GO). Just run it along side your Rails app it will learn your database schema, relationships and allow you to query your database using just simple GraphQL.

Super Graph understands Rails cookies and works with session stores to get the current authenticated users ID. 

Github:
https://github.com/dosco/super-graph

Documentation:
https://supergraph.dev/docs/home
## [7][You must use Bundler 2 or greater with this lockfile](https://www.reddit.com/r/rails/comments/gnr5nk/you_must_use_bundler_2_or_greater_with_this/)
- url: https://www.reddit.com/r/rails/comments/gnr5nk/you_must_use_bundler_2_or_greater_with_this/
---
I am trying to deploy an application on Linux Debian 9 and I am getting an error  
`You must use Bundler 2 or greater with this lockfile`  
`Ruby -v returns ruby 2.6.3p62`  
`rails -v returns Rails` [`6.0.3.1`](https://6.0.3.1)  
`which bundle returns /usr/local/rvm/gems/ruby-2.6.3/bin/bundle`  
`bundle version returns Bundler version 2.1.4 (2020-01-05 commit 32a4159325)`

So, I have tried various things such as

1. Updating the bundler gem(gem install bundler -v 2.1.4)
2. Updating the rubygems Package manager( gem update --system)
3. Reinstalling bundler gem

but they dont seem to work and  everytime I try to work it throws the same error.  
Anyone who has been facing the same error and could help?
## [8][Image upload.](https://www.reddit.com/r/rails/comments/gnkn0e/image_upload/)
- url: https://www.reddit.com/r/rails/comments/gnkn0e/image_upload/
---
Hit a block here. I've handled image uploading before via action text, but not as an element in a db that I can call back. Wondering, would I just add the params/permit and migration into the post table? If so, what would be the syntax for the migration? And as far as views go, it's pretty straight forward via simple\_for &amp; a framework right?

It's for a post/index page which will have a title/description linking to a full post. Hoping I can get a db-driven image since I don't want them all to have a static image. Will be 1 square image per post for the card.
## [9][Rolify and many-to-many relationships](https://www.reddit.com/r/rails/comments/gnov33/rolify_and_manytomany_relationships/)
- url: https://www.reddit.com/r/rails/comments/gnov33/rolify_and_manytomany_relationships/
---
Let's say I'm building an app that models court cases. I have a `Case` model and a `User` model for authentication. I then use Rolify to assign different roles to `Users`, `Plaintiff`, `Defendant`, `Judge`, `Counsel` , and so on. These roles determine what CRUD actions users can take on `Cases`.

What is the best way to get all `Cases` for a `Judge`? I don't think I can use Rolify to set up those many-to-many relationships.

Thanks in advance!
## [10][Question: How do you organize code in your Rails app?](https://www.reddit.com/r/rails/comments/gnbukt/question_how_do_you_organize_code_in_your_rails/)
- url: https://www.reddit.com/r/rails/comments/gnbukt/question_how_do_you_organize_code_in_your_rails/
---
One thing I constantly stuggle is the problem where my code should live and how the thing that holds it should be named. Rails out of the box doesn't help you much with that so patterns emerge that help to organize things. 

Here's some obvious patterns that most Rails apps have:

**Service Objects** You initialize service object and it has a single public method like `run` or `execute`. So you send it some input and it just does things. It may change the state of the database, send emails, etc. Generally returns true/false. Example:

    class FooService
      def initialize(foo:)
        @foo = foo
      end
  
      def run
        do_things_to(@foo)
      end
    end

**Presenters** Like a view helper, but used specifically to present objects in the views. Example:

    class FooPresenter
      def initialize(object)
        @foo = object
      end
  
      def avatar
        image_tag(@foo.avatar)
      end
    end

**Decorators** I think the idea is that methods from this class are mixed into object that is passed in. Honestly, I've been treating decorators like pseudo presenters that are not used in the views. For example: `FooDecorator.new(foo).some_method_inside_decorator_class`. It's clearly not a decorator pattern but it's not really a presenter in a sense that it will provide outputs to be rendered on the screen. What should it be really called?

**Form Objects** To deal with complicated logic about handling certain forms. 99% of the time you'd just handle record validation and saving from the controller, but sometimes things get hairy so you'd do something like this:

    class FooForm
      include ActiveModel::Model
      attr_accessor :bar
  
      def save
        if valid?
          save_something
        else
          add_some_errors
          false
        end
      end
    end

and in the controller you'd have

    @foo_form = FooForm.new(account_claim_form_params)
    if @foo_form.save
      flash[:success] = "Yey"
      redirect_to path
    else
      flash[:error] = "Boo!"
      render :new
    end

I also have things like `pdf_generators`, `csv_exporters`. Just organizing code based on what it does. 

Now, there are some things I have no idea how to classify. For example:

- There's a class that digs through yaml file. It initializes with some inputs and then has a pile of public methods you can call. It's not a service and not really a model since it has absolutely nothing to do with database.
- There's a class that figures out "next work day". You feed it a date and with `increment_by(n)` if will give you next work date. What is this thing? I lives in our services right now, but it's clearly doesn't belong there. How would you classify this?
- I have a ton of view helpers that have methods like `project_category_options_for_select`. It doesn't feel right to have them in global view helpers. Is this a presenter/decorator... or something else?
  
I'm working with Rails apps for ages now, but organizing code is literally the hardest thing that I ahve to deal with.
## [11][Find the first submitted form record](https://www.reddit.com/r/rails/comments/gng8sg/find_the_first_submitted_form_record/)
- url: https://www.reddit.com/r/rails/comments/gng8sg/find_the_first_submitted_form_record/
---
Hi all, I have a db table that stores records of my form when the status changes.  I am currently getting data like this in a service

\[    {form\_id: 1, status: 0, created\_at: date form was created},   {form\_id: 1, status: 4, created\_at: first submitted date}, //record i want   {form\_id: 1, status: 2, created\_at: date form was returned},   {form\_id: 1, status: 4, created\_at: date form was resubmitted},   {form\_id: 2, status: 0, created\_at: date form was created},   {form\_id: 2, status: 4, created\_at: first submitted date}, //record i want   {form\_id: 2, status: 2, created\_at: date form was returned},   {form\_id: 2, status: 4, created\_at: date form was resubmitted},   {form\_id: 3, status: 0, created\_at: date form was created},   {form\_id: 3, status: 4, created\_at: first submitted date}, //record i want   {form\_id: 3, status: 2, created\_at: date form was returned},   {form\_id: 3, status: 4, created\_at: date form was resubmitted},   \]

&amp;#x200B;

&amp;#x200B;

I've commented the records I want to grab, the first submitted record where status = 0 and the oldest date

&amp;#x200B;

edit:

FormHistoryRecord.where(status: 4).order(:form\_id, :created\_at).uniq{|a| a.form\_id }  seems to work
## [12][Integrating Ckeditor per user](https://www.reddit.com/r/rails/comments/gnhz9y/integrating_ckeditor_per_user/)
- url: https://www.reddit.com/r/rails/comments/gnhz9y/integrating_ckeditor_per_user/
---
Can anyone help me out getting Ckeditor setup so that each user that uploads photos will only see their photos and not every photo uploaded by everyone?

I am using this gem  [https://github.com/galetahub/ckeditor](https://github.com/galetahub/ckeditor)
