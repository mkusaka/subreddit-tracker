# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/j6qvuh/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/j6qvuh/personal_projects_show_off_your_own_project_andor/
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
## [3][Generating multi-page pdf reports consisting of headers and tables](https://www.reddit.com/r/rails/comments/j8b6fb/generating_multipage_pdf_reports_consisting_of/)
- url: https://www.reddit.com/r/rails/comments/j8b6fb/generating_multipage_pdf_reports_consisting_of/
---
What do you guys use to accomplish this? I looked into wicked\_pdf but generating 30 pages took way too long. Am I doing it wrong because I was forced to use combine\_pdf gem to combine 30 files to one. Should I be using Prawn gem instead? Is there another method?
## [4][How to hook standalone graphiql client to Rails routes?](https://www.reddit.com/r/rails/comments/j8gvig/how_to_hook_standalone_graphiql_client_to_rails/)
- url: https://www.reddit.com/r/rails/comments/j8gvig/how_to_hook_standalone_graphiql_client_to_rails/
---
Graphiql-Rails hasn't been updated for almost two years and it's missing quite some features (like request header). 

I downloaded the standalone graphiql client, but having trouble how to hook it up to the routes. Do I need to define some additional routes in rails for it to work? Currently I only have 

`post "/graphql", to: "graphql#execute"`
## [5][How Do I Serve GZip CSS and JS with Rails 6 App Using Cloudfront?](https://www.reddit.com/r/rails/comments/j860xd/how_do_i_serve_gzip_css_and_js_with_rails_6_app/)
- url: https://www.reddit.com/r/rails/comments/j860xd/how_do_i_serve_gzip_css_and_js_with_rails_6_app/
---
Hello everyone!  

I'm trying to serve my CSS and JS compressed as a gzip from Cloudfront with my Rails 6 app, but I can't get it to work.

Following [Thoughtbot's tutorial](https://thoughtbot.com/blog/content-compression-with-rack-deflater), I added this to application.rb:

    config.middleware.use Rack::Deflater 

I have also configured Cloudfront to automatically set [Compress Objects Automatically](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/ServingCompressedFiles.html) to 'yes' in Cloudfront. What am I doing wrong?

&amp;#x200B;

https://preview.redd.it/rheow3c4j4s51.png?width=676&amp;format=png&amp;auto=webp&amp;s=aee0401e385caa5e76a968aa3bbc80d9bcaae133

&amp;#x200B;

https://preview.redd.it/d71jnp56j4s51.png?width=912&amp;format=png&amp;auto=webp&amp;s=e26abe0d3be450b7e72adc5fd6800734808115b1
## [6][Active Storage -- IOError closed stream](https://www.reddit.com/r/rails/comments/j81jyh/active_storage_ioerror_closed_stream/)
- url: https://www.reddit.com/r/rails/comments/j81jyh/active_storage_ioerror_closed_stream/
---
Hi. I don't know what the hell is going on here. I try to mirror my Cloudinary Active Storage uploads to my local disk. This is my storage.yml:

```yml
test:
  service: Disk
  root: &lt;%= Rails.root.join("tmp/storage") %&gt;

local:
  service: Disk
  root: &lt;%= Rails.root.join("storage") %&gt;

cloudinary:
  service: Cloudinary
  api_key: &lt;%= ENV.fetch 'CL_API_KEY' %&gt;
  api_secret: &lt;%= ENV.fetch 'CL_API_SECRET' %&gt;
  cloud_name: &lt;%= ENV.fetch 'CL_CLOUD_NAME' %&gt;
  upload_preset: &lt;%= ENV.fetch 'CL_UPLOAD_PRESET' %&gt;

backup:
  service: Mirror
  primary: cloudinary
  mirrors:
    - local
```

In my development and production env files, I have configured `backup` to be the Active Storage "driver":

```rb
config.active_storage.service = :backup
```

Everytime I make an upload, it goes through just fine to Cloudinary, but craps out saving it locally with said error (IOError closed stream).

The model that I use for uploads:

```rb
class Post &lt; ApplicationRecord
    has_many :comments, dependent: :destroy
    has_many :taggings, dependent: :destroy
    has_many :tags, through: :taggings

    has_many_attached :attachments

    validates :title, presence: true, length: { maximum: 255 }
    validates :slug, presence: true, length: { maximum: 255 }, uniqueness: true
    validates :content, presence: true

    attr_accessor :remove_attachments

    def to_param
        self.slug
    end

    def all_tags= names
        self.tags = names.split(',').map do |name|
            Tag.where(name: name).first_or_create!
        end
    end

    def all_tags
        tags.map(&amp;:name).join ', '
    end

    def self.tagged_with name
        Tag.find_by!(name: name).posts
    end

    after_save :delete_attachments!, if: :remove_attachments
    private def delete_attachments!
        attachments.purge_later
    end
end
```

The controller's `create` and `post_params` methods:

```rb
    def create
        @post = Post.new post_params

        if @post.save
            redirect_to @post
        else
            render 'new'
        end
    end

    private def post_params
        params.require(:post).permit :title, :slug, :content, :is_draft, :is_archived, :all_tags, :remove_attachments, attachments: []
    end
```

I really don't know what's causing this and I couldn't find anything online. I'm going on vacation tomorrow and I'd love to fix it today, but I can't figure out how. Thanks a lot!
## [7][Moving to Rails from the Front End](https://www.reddit.com/r/rails/comments/j7llb1/moving_to_rails_from_the_front_end/)
- url: https://www.reddit.com/r/rails/comments/j7llb1/moving_to_rails_from_the_front_end/
---
Hello friends! I'm currently a senior front end developer. I've been using Rails in personal projects for a little over 3 years now and I'm reaching out to you all for some advice. I'm finding that I have greatly enjoyed working on the server side and have fallen in love with Ruby and Rails and I'd like to attempt a pivot in my career and move to a position that is more focused on those technologies. I have worked on the server side professionally with node but it's been in a very limited capacity so it's hard to point to that experience as some kind of proof.

I understand that the competition is fierce and that these domains can be pretty different but I'm hoping someone can help me think a little creatively about moving into this awesome community.
## [8][React/Rails + devise question: how the F*** do I render current_user data in my React front-end?](https://www.reddit.com/r/rails/comments/j7plfk/reactrails_devise_question_how_the_f_do_i_render/)
- url: https://www.reddit.com/r/rails/comments/j7plfk/reactrails_devise_question_how_the_f_do_i_render/
---
Need some help! 

I am trying to render my current\_user data on my React front-end. 

The JSON renders fine directly in the browser at localhost:3000/current\_user (a custom route I wrote), but when I try to fetch that data in react, I get a null response. I am using devise+google\_oauth2 with a Rails API. The authentication is working, user\_signed\_in? returns true, current user checks out, all of that.

 I have been googling for hours and haven't found anything that works. I tried installing the devise-token-authentication gem and overwriting the standard devise current\_user with this method in my application controller like so:

`def current_user`  
  `token = request.headers["Authorization"].to_s`  
  `User.find_for_database_authentication(authentication_token: token)`  
`end` 

but that didn't work either (stack overflow solution). I am using the `before_action :authenticate_user!` in my custom controller, which seems to be doing something as I get a 401 response without it. No matter how I try to tackle this, the JSON response from my fetch returns null. If anyone can help. I would greatly appreciate it! I'm assuming it has something to do with the request from my front-end not being authenticated but I can't figure this one out for the life of me...
## [9][Storing user settings in rails 5 API](https://www.reddit.com/r/rails/comments/j79doh/storing_user_settings_in_rails_5_api/)
- url: https://www.reddit.com/r/rails/comments/j79doh/storing_user_settings_in_rails_5_api/
---
I'm wondering what the best way to store user preferences / settings in a rails 5 api application. Most of my search results use deprecated gems, or are behind a paywall. (looking at you gorails). Any info on the topic would be greatly appreciated, thanks
## [10][link_to generates path like /controller?id=1234, I need /controller/1234](https://www.reddit.com/r/rails/comments/j711dn/link_to_generates_path_like_controllerid1234_i/)
- url: https://www.reddit.com/r/rails/comments/j711dn/link_to_generates_path_like_controllerid1234_i/
---
I am having some trouble rendering my show.html.erb view. I have the following line to create a button on the page:

`&lt;%= link_to('Click Here', books_path(id: book.id), :class =&gt; 'button') %&gt;`

which generated a path like /books?id=1234. This does not render my show view or hit my show action in the controller. However, if I manually type /books/1234 into the address bar it seems to work and render my view.

in routes.rb I have:

`resources :books`

Can anyone shed any light on what I'm doing wrong here?
## [11][Scoped associations in Rails](https://www.reddit.com/r/rails/comments/j6lxl8/scoped_associations_in_rails/)
- url: https://www.reddit.com/r/rails/comments/j6lxl8/scoped_associations_in_rails/
---
Howdy folks! 

I've never thought that ActiveRecord associations could be used as scopes until last week. I couldn't find a step-by-step explanation, so [I did a quick write up about how one can scope associations in Rails](https://remimercier.com/scoped-active-record-associations/). 

Some key take-aways: 
👀  Make your code much more readable (stylish, even)
✂  DRY things up
⚡️  Preload your scopes and remove your N+1 queries

Hope you'll like it!
## [12][RailsAdmin: How to disable edit action?](https://www.reddit.com/r/rails/comments/j6qhyq/railsadmin_how_to_disable_edit_action/)
- url: https://www.reddit.com/r/rails/comments/j6qhyq/railsadmin_how_to_disable_edit_action/
---
Hi there,

I'm working on a rails project with [rails\_admin](https://github.com/sferik/rails_admin) and multiple models. There are several people working on the backend and I want to remove the ability to edit some of the records which have a `imported` boolean set to `true`. This records should just be readable in rails\_admin.

Is it possible to remove the pencil-button in the list view of a record an make this records read only with the condition that the flag is `true`?

The documentation [shows up](https://github.com/sferik/rails_admin/wiki/Base-action#only) some similar behavior to disable action in general but there is no way to bind this configuration on a condition (I guess).

I just found out that you can define a method called `readonly?` on the model but that won't work while I have to update or change attributes while processing this record in my programming logic.

Many thanks!
