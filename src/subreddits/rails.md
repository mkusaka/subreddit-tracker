# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/)
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
## [3][MVP : start with Devise (or equivalent) or auth0 (or equivalent) ?](https://www.reddit.com/r/rails/comments/jhqo9r/mvp_start_with_devise_or_equivalent_or_auth0_or/)
- url: https://www.reddit.com/r/rails/comments/jhqo9r/mvp_start_with_devise_or_equivalent_or_auth0_or/
---
If you had to build a MVP for a new startup, will you choose an authentication provider (auth0 or whatever  equivalent), or will you integrate it inside your application (Clearance, Devise, Sorcery...). I have this choice to do and I find it hard to make the right decision. In theory less software is better, but I feel I loose control over what is happening, and the main interest of Rails is precisely ability to integrate anything. Thanks for your thoughts!
## [4][Help: Good way to handle mail templates stored in database](https://www.reddit.com/r/rails/comments/jhomuk/help_good_way_to_handle_mail_templates_stored_in/)
- url: https://www.reddit.com/r/rails/comments/jhomuk/help_good_way_to_handle_mail_templates_stored_in/
---
Hi all, 

I have an administration system for customers, where they can control the product we are providing. We need to send emails on different actions, to confirm then changes has been done. 

Sometimes it need to sent an email with different content to 3 receptions. 

* The user who did the action
* The administrator(s) for the users customer
* Our administrators (or helpdesk shared email)

The emails contains variables, with data from the execution. 

Currently  we are handing it directly in rails. But now we have partners which also have customers, and might want to make changes in the templates. 

Therefore, we need someway smart to handle all the mail templates in a user friendly interface together with the rest of the application. Maybe some markdown support in a simple wysiwyg editor, however that is more a frontend question. First we need the backend to support til mail templates. 

To spice it, we also need to handle multiple languages, however currently we only support two. 

We are using Rails 6 api only as backend and React as frontend.
## [5][What are some of the most beautiful rails made websites you've visited?](https://www.reddit.com/r/rails/comments/jhf166/what_are_some_of_the_most_beautiful_rails_made/)
- url: https://www.reddit.com/r/rails/comments/jhf166/what_are_some_of_the_most_beautiful_rails_made/
---

## [6][Automatically check whether the current user is allowed to modify a record](https://www.reddit.com/r/rails/comments/jhfm2h/automatically_check_whether_the_current_user_is/)
- url: https://www.reddit.com/r/rails/comments/jhfm2h/automatically_check_whether_the_current_user_is/
---
Solution: I'm disappointed at myself. It was right in [the docs](https://guides.rubyonrails.org/action_controller_overview.html#filters):

```rb
class BoardsController &lt; ApplicationController
    before_action :authenticate_user!
    before_action :must_be_board_owner
    skip_before_action :must_be_board_owner, only: [:new, :create, :index]

    # ...

    def must_be_board_owner
        if current_user != Board.find_by(slug: params[:slug]).user
            redirect_to root_path
        end
    end
end
```

---

Hi. I currently have the following problem:
On my controllers, I have this before_action filter:

```rb
# as you can see, I'm using Devise
before_action :authenticate_user!
```

This checks whether the user is signed in or not.
However, I also need to check whether the current user is allowed to see/modify the current record.
Right now, I have this (horrible) method:

```rb
    private def redirect_unless_board_owner! owner
        redirect_to root_path unless current_user == owner
    end
```

Using it in an action looks like this:

```rb
        @board = Board.find_by slug: params[:board_slug]
        redirect_unless_board_owner! @board.user

        @task = Task.find_by slug: params[:slug]
```

Is there any way that I can make this check a before_action filter? This just looks so ugly in my code, ugh. Thanks.
## [7][Can I use Pundit to shadow ban?](https://www.reddit.com/r/rails/comments/jhozhz/can_i_use_pundit_to_shadow_ban/)
- url: https://www.reddit.com/r/rails/comments/jhozhz/can_i_use_pundit_to_shadow_ban/
---
Can you make me an example to hide the posts only for the shadow banned user using pundit gem?
## [8][Some columns are not created after going to MySQL](https://www.reddit.com/r/rails/comments/jh9oaz/some_columns_are_not_created_after_going_to_mysql/)
- url: https://www.reddit.com/r/rails/comments/jh9oaz/some_columns_are_not_created_after_going_to_mysql/
---
I made a bunch of migrations, and this is my schema :

    # This file is auto-generated from the current state of the database. Instead
    # of editing this file, please use the migrations feature of Active Record to
    # incrementally modify your database, and then regenerate this schema definition.
    #
    # This file is the source Rails uses to define your schema when running `rails
    # db:schema:load`. When creating a new database, `rails db:schema:load` tends to
    # be faster and is potentially less error prone than running all of your
    # migrations from scratch. Old migrations may fail to apply correctly if those
    # migrations use external dependencies or application code.
    #
    # It's strongly recommended that you check this file into your version control system.
    
    ActiveRecord::Schema.define(version: 2020_10_14_075849) do
    
      create_table "comments", options: "ENGINE=InnoDB DEFAULT CHARSET=utf8", force: :cascade do |t|
        t.text "body"
        t.datetime "created_at", precision: 6, null: false
        t.datetime "updated_at", precision: 6, null: false
        t.integer "post_id"
        t.integer "user_id"
        t.index ["post_id"], name: "index_comments_on_post_id"
        t.index ["user_id"], name: "index_comments_on_user_id"
      end
    
      create_table "favorites", options: "ENGINE=InnoDB DEFAULT CHARSET=utf8", force: :cascade do |t|
        t.string "favoritable_type", null: false
        t.bigint "favoritable_id", null: false
        t.string "favoritor_type", null: false
        t.bigint "favoritor_id", null: false
        t.string "scope", default: "favorite", null: false
        t.boolean "blocked", default: false, null: false
        t.datetime "created_at", precision: 6, null: false
        t.datetime "updated_at", precision: 6, null: false
        t.index ["blocked"], name: "index_favorites_on_blocked"
        t.index ["favoritable_id", "favoritable_type"], name: "fk_favoritables"
        t.index ["favoritable_type", "favoritable_id"], name: "index_favorites_on_favoritable_type_and_favoritable_id"
        t.index ["favoritor_id", "favoritor_type"], name: "fk_favorites"
        t.index ["favoritor_type", "favoritor_id"], name: "index_favorites_on_favoritor_type_and_favoritor_id"
        t.index ["scope"], name: "index_favorites_on_scope"
      end
    
      create_table "likes", options: "ENGINE=InnoDB DEFAULT CHARSET=utf8", force: :cascade do |t|
        t.bigint "post_id", null: false
        t.bigint "user_id", null: false
        t.datetime "created_at", precision: 6, null: false
        t.datetime "updated_at", precision: 6, null: false
        t.index ["post_id"], name: "index_likes_on_post_id"
        t.index ["user_id"], name: "index_likes_on_user_id"
      end
    
      create_table "posts", options: "ENGINE=InnoDB DEFAULT CHARSET=utf8", force: :cascade do |t|
        t.string "image"
        t.text "caption"
        t.datetime "created_at", precision: 6, null: false
        t.datetime "updated_at", precision: 6, null: false
        t.integer "user_id"
        t.index ["user_id"], name: "index_posts_on_user_id"
      end
    
      create_table "users", options: "ENGINE=InnoDB DEFAULT CHARSET=utf8", force: :cascade do |t|
        t.string "email", default: "", null: false
        t.string "encrypted_password", default: "", null: false
        t.string "reset_password_token"
        t.datetime "reset_password_sent_at"
        t.datetime "remember_created_at"
        t.datetime "created_at", precision: 6, null: false
        t.datetime "updated_at", precision: 6, null: false
        t.string "username"
        t.index ["email"], name: "index_users_on_email", unique: true
        t.index ["reset_password_token"], name: "index_users_on_reset_password_token", unique: true
        t.index ["username"], name: "index_users_on_username", unique: true
      end
    
      add_foreign_key "likes", "posts"
      add_foreign_key "likes", "users"
    end

Then, I added this to the `database.yml` file to make it work with MySQL :

    default: &amp;default
      adapter: mysql2
      encoding: utf8
      pool: &lt;%= ENV.fetch("RAILS_MAX_THREADS") { 5 } %&gt;
      username: railstagram
      password: 
      socket: /tmp/mysql.sock

First, I'm a bit confused why it didn't ask for port or host, but I assume it's handled somewhere else.

By the way, when I run `rails db:create` and `rails db:migrate` everything seems just fine, but when I create a user, it seems it doesn't care for the *username* field.

This is a [link](https://github.com/prp-e/railstagram) to my project. I need to deploy it tonight and I have this problem with MySQL. My development environment is macOS Catalina.
## [9][Advice on Application Architecture](https://www.reddit.com/r/rails/comments/jh42dn/advice_on_application_architecture/)
- url: https://www.reddit.com/r/rails/comments/jh42dn/advice_on_application_architecture/
---
Hey everyone

I'm building a social-network-like platform with a Rails API and React all on AWS EBS. This isn't a hobby project, so I need to be careful about my design choices and would therefore appreciate some advice.

What's worrying me is how I will introduce realtime functionality. Like most social networks, my app will have a feed, likes, and maybe comments--all of which need to be updated in (near) realtime when something happens in the database. There is also a 1-to-1 chat component.

Because I'm kind of on a ticking clock, I'm thinking about taking the easy way out with the chat and using Firebase for that (chat is not really relevant to my biz logic). From what I'm reading, realtime chat seems to be finicky with Rails and ActionCable. I hypothesize that ActionCable will do the job for all the other realtime stuff. Would you agree?

I've seen alternatives like Pusher and Pubnub and they look great, but they're also quite pricey. Anyone run into a similar situation?
## [10][Moving a method: :delete in another page](https://www.reddit.com/r/rails/comments/jh7teq/moving_a_method_delete_in_another_page/)
- url: https://www.reddit.com/r/rails/comments/jh7teq/moving_a_method_delete_in_another_page/
---
Hi guys,

in the page in view app/view/books/edit.html.erb I have this to delete the page

    &lt;%= link_to "Destroy", book_path(id: @book.slug), method: :delete, data: { confirm: "are you sure?" } %&gt;

I want to move it in a custom popup. So I edited it in this way

    &lt;%= link_to "Destroy", warning_path(:type =&gt; 'delete_book'), remote: true, data: { load_popup: true } %&gt;

and in app/views/warning/index.html.erb I added this

    &lt;% if params[:type] == 'delete_book' %&gt;
    	&lt;%= link_to "Destroy", book_path(id: @book.slug), method: :delete, data: { confirm: "are you sure?" } %&gt;
    &lt;% end %&gt;

But I have this error: `Undefined method 'slug' for nil:NilClass`. How to fix?
## [11][[Help] Rails Server gives error "ERR_CONNECTION_REFUSED" in browser](https://www.reddit.com/r/rails/comments/jgx048/help_rails_server_gives_error_err_connection/)
- url: https://www.reddit.com/r/rails/comments/jgx048/help_rails_server_gives_error_err_connection/
---
EDIT: Solved changing port 8000 seemed to do the trick thanks for your help

I'm new to rails development and only have a couple of simple projects. All of my projects are stored in a wsl2 Ubuntu 18.04LTS vm. I have always been able to run \`rails s\` and view my project on localhost:3000. Even last night I was able to view 2 separate projects this way. Now, today when I start rails server everything looks like it works as expected within the terminal (starting on [127.0.0.1:3000](https://127.0.0.1:3000)) but when I try to load in my browser I get \`ERR\_CONNECTION\_REFUSED\`. Other than the message in the browser the server in the terminal looks just like normal and shows no error messages. Any help would be greatly appreciated.
## [12][Security : Is sending back an invalid password is a good or bad practice ?](https://www.reddit.com/r/rails/comments/jgop0h/security_is_sending_back_an_invalid_password_is_a/)
- url: https://www.reddit.com/r/rails/comments/jgop0h/security_is_sending_back_an_invalid_password_is_a/
---
If user wants to register to my app and enter a too short password, is it bad practice to fill again this password in the rendered template ? (Or for a JSON request, show this "bad" password in the JSON response).
