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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/)
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
## [3][bjt(bundle jump to), a quick navigation tool for bundle packages like njt](https://www.reddit.com/r/rails/comments/i9hyhx/bjtbundle_jump_to_a_quick_navigation_tool_for/)
- url: https://www.reddit.com/r/rails/comments/i9hyhx/bjtbundle_jump_to_a_quick_navigation_tool_for/
---
hi i write rubygem: [bjt](https://github.com/superiorlu/bjt), is a quick navigation for rubygems, if you like it,  star it. :smile:
## [4][Rendering HTML in text fields](https://www.reddit.com/r/rails/comments/i9i52q/rendering_html_in_text_fields/)
- url: https://www.reddit.com/r/rails/comments/i9i52q/rendering_html_in_text_fields/
---
Although I think it's not that much of a rails related question, but I ask it here anyway :P

I have created a very simple model :

    rails g scaffold Blog title:string text:text

In `views` it includes a form, and of course it has two inputs. When I type something like:

    &lt;p&gt; Hello! &lt;/p&gt;
    &lt;p&gt; This is my today's post &lt;/p&gt;

it doesn't render it to a proper HTML file and no surprise. I'm just curious, is there any ways to force the HTML rendering to textareas in ERB?
## [5][Breaking Down Test-Driven Development (TDD) Free Lunch and Learn Webinar](https://www.reddit.com/r/rails/comments/i991hs/breaking_down_testdriven_development_tdd_free/)
- url: https://www.reddit.com/r/rails/comments/i991hs/breaking_down_testdriven_development_tdd_free/
---
Hi Rubyists! 

Def Method is hosting a free lunch and learn webinar on TDD! Test driving software leads to cleaner code, better design, and fewer bugs. But when your stakeholders push for faster optimization, why is it so important to continue forward with TDD? In our opinion, testing is paramount for the success of a product. 

In this panel discussion featuring Def Method Engineers, panelists will discuss topics including: the proper amount of test coverage, how to know if your unit tests are producing value, the importance of TDD when building a proof of concept, working on legacy software with TDD, and useful practices while test-driving. 

Featuring panelists: Senior Software Engineers Keith Hickman-Perfetti, Jeff Jia, Ben Gross, Mark Simpson. Moderated by: Yisselda Rhoc

RSVP here: [https://www.eventbrite.com/e/breaking-down-test-driven-development-tdd-panel-discussion-tickets-116644287203?aff=Reddit](https://www.eventbrite.com/e/breaking-down-test-driven-development-tdd-panel-discussion-tickets-116644287203?aff=Reddit)
## [6][50K users](https://www.reddit.com/r/rails/comments/i92ace/50k_users/)
- url: https://www.reddit.com/r/rails/comments/i92ace/50k_users/
---
Hi guys, quick question. For a rails project that is expecting 50k active users (all at the same time), what are your recommendations for a server infrastructure that can widthstand this kind of load? I have already investigated SOA infrastructure. Currently the app is monolithic.

EDIT: Basically it's for a Virtual Fair that a client of ours is asking for. It would have many requests to the database, mainly fetching information for stands like videos, jobs listings, webchat, CV/resume upload, downloading brochures. We had a similar virtual fair early this year and our server started crashing at around 700 active users. We tried vertical scaling but it didn't help.
## [7][How to create lots of web projects that are copies of each other (90%) but also have some unique elements (10%) in each? And that also support updates](https://www.reddit.com/r/rails/comments/i9fs43/how_to_create_lots_of_web_projects_that_are/)
- url: https://www.reddit.com/r/rails/comments/i9fs43/how_to_create_lots_of_web_projects_that_are/
---
I have a web project in Rails of which I plan to be creating copies and running them all. Let's say, it's a blog. Each blog will be deployed on its own domain and they basically, 90%, will be identical. And around 10% will be unique, custom on each blog.

&amp;#x200B;

Whenever I update the code of 'Blog1' locally, I'll push an update, then I'll log in to the servers of other blogs and do "git pull" to retrieve an update, and rebuild them. Yet, an update won't overwrite the 10% of custom stuff that each blog have.

&amp;#x200B;

Custom stuff will include: js, css, images files, some parts of the pages, names, texts, etc...

&amp;#x200B;

**(1)** How will I organise all that? The amount of js, css, images, etc can very on each blog, therefore I can't simply add "custom\_style.css" to .gitignore

&amp;#x200B;

Saving css, js, image in a database isn't an option, because retrieving them on each request will cost too much.

&amp;#x200B;

Or perhaps, do it somehow statically and \*only once\* at booting a blog, when they'll be retrieved from Sqlite3.

&amp;#x200B;

**(2)** And how would I go about rendering unique parts of html pages? Such as footers, top nav bars. They may or may not, or may be partially different.

&amp;#x200B;

**(3)** All the code itself will identical, at this point, probably. Otherwise, how would I go about code as well? Namely, some blogs would have peaces or whole files of code unique to them. Something similar to plug-in system, I figure.

&amp;#x200B;

P.S. I'm aware of tenant architecture, that's not completely it.
## [8][Is there a better way for nested relationships?](https://www.reddit.com/r/rails/comments/i93oqh/is_there_a_better_way_for_nested_relationships/)
- url: https://www.reddit.com/r/rails/comments/i93oqh/is_there_a_better_way_for_nested_relationships/
---
Hi, I've done a few tutorials on rails but wanted to jump in as fast as possible just wondering if the following is the rails way or if it can be improved upon?

Pokemon.rb

      belongs_to :generation
      has_many :pokemon_pokedex_numbers, inverse_of: :pokemon
      has_many :pokedexes, through: :pokemon_pokedex_numbers
      has_many :version_groups, through: :pokedexes
      has_many :versions, through: :version_groups
    
      accepts_nested_attributes_for :pokemon_pokedex_numbers, reject_if: :all_blank, allow_destroy: true

Pokedex.rb

      has_many :pokedex_version_groups
      has_many :version_groups, through: :pokedex_version_groups
      has_many :pokemon_pokedex_numbers
      has_many :pokemon, through: :pokemon_pokedex_numbers

PokemonPokedexNumber

      belongs_to :pokemon
      belongs_to :pokedex

PokedexVersionGroup

      belongs_to :pokedex
      belongs_to :version_group

VersionGroup.rb

      belongs_to :generation
      has_one :pokedex_version_group
      has_one :pokedex, through: :pokedex_version_group
      has_many :versions

Version.rb

    belongs_to :version_group

Before using the has\_many through another has\_many through in Pokemon I had to loop through pokedexes-&gt;version\_groups-&gt;versions to print each version but I found with the current way the following controller:

      # GET /pokemon
      def index
        @pokemon = Pokemon.includes(:versions).all
      end

eager loads all the tables (6 queries) and I can just loop versions directly on the pokemon.

    pokemon.versions.each
## [9][new rails migration automatically has status of "up"](https://www.reddit.com/r/rails/comments/i94ahr/new_rails_migration_automatically_has_status_of_up/)
- url: https://www.reddit.com/r/rails/comments/i94ahr/new_rails_migration_automatically_has_status_of_up/
---
Not sure what's happening, but when I create a new rails migration, the status is automatically "up" but when I check the schema, it didn't make the changes.

Have the following versions:

    $ rails -v
    Rails 5.2.4.3
    $ ruby -v
    ruby 2.6.0p0 (2018-12-25 revision 66547) [x86_64-darwin18]

When checking the migration files, previous migrations started with v5.1:

\`class MyMigration &lt; ActiveRecord::Migration\[5.1\]\`

But my new migration has v5.2:

\`class MyOtherMigration &lt; ActiveRecord::Migration\[5.2\]\`

&amp;#x200B;

Edit:

After removing the migration and creating a new migration, I'm getting the following error whenever I try to check status of migrations:

    ERROR:  column "name" of relation "books" already exists (PG::DuplicateColumn)

When I check `schema_migrations`, the new migration timestamp is not listed

When I go into `$ rails c`, and check the Book object, the "name" field is there. My `schema.rb` file does not have the "name" field though. I'm guessing, my previous migration somehow migrated without updating the `schema.rb` file and after removing the migration without performing the "down" action on it made my state like this. I did try running the "down" command before removing the migration, but it didn't work.

&amp;#x200B;

Edit2:

I updated my migration file to both have an "up" and "down" that removes the new field:

    remove_column :books, :name, :string

After saving the file, it somehow migrating it right away. I can now perform things like "`$ rails db:migrate:status`" without getting an error. Checking "`$ rails c`", the new field is removed. I did not explicitly run an "up" or "down" migrate command after saving the migration file. It somehow ran it automatically..

&amp;#x200B;

Edit3: 

Not sure what was happening, but decided just to reset (Warning: this drops the data in db)

    $ rake db:reset
    $ rails db:migrate

Now I'm back to the state of "up" migrations happening right away and "down" migrations do not work. \*sigh\*
## [10][Question about text encryption in Rails](https://www.reddit.com/r/rails/comments/i8zp1u/question_about_text_encryption_in_rails/)
- url: https://www.reddit.com/r/rails/comments/i8zp1u/question_about_text_encryption_in_rails/
---
I want to encrypt the content of a column in a model in Rails. Here's the scenario:

* There are Users (devise) and Posts. Each User can have many Posts
* Each Post has a title and a content column 
* When someone logs in to the application, then he has access to all posts/content 
* The content is of course a column of type text.

I want to limit access to certain Post content, leaving the title intact for search purposes. I want the app to prompt the user for a password (when he enables somehow this feature) and If he enters the right one, show the text. 

One silly idea is to redirect the user to an Unlock this Post page where there could be a single text\_field form that posts to a controller action. The verification of the right passcode is done and If legit, he will be redirected to the show decrypted Post page. 

I don't want to use a gem for this (lockbox or attr\_encrypted). The text can be encrypted with a before\_save callback in the Post model. 

Have you ever ran into a similar issue and If yes, what did you do ?
## [11][SSO server in rails, is there any gems?](https://www.reddit.com/r/rails/comments/i8vq7h/sso_server_in_rails_is_there_any_gems/)
- url: https://www.reddit.com/r/rails/comments/i8vq7h/sso_server_in_rails_is_there_any_gems/
---
Actually, I'm looking for a way to make my app the SSO server and users will be able to login to other services I made with a single account.
## [12][Looking for Rails tutor/mentor](https://www.reddit.com/r/rails/comments/i8lzk7/looking_for_rails_tutormentor/)
- url: https://www.reddit.com/r/rails/comments/i8lzk7/looking_for_rails_tutormentor/
---
I'm currently building an app for a friend, which I'd like to eventually turn into a multi-tenant SaaS app. I'm self-taught through online tutorials, so I've been able to make decent progress. However, I'm starting to hit a wall with certain features.

I plan to tackle these features over the next few weeks: multi-tenancy, setting up robust 3rd party API integrations (Stripe, Mailgun), background jobs, and deployment beyond Heroku.

I'm looking for someone that is able to meet once a week via Zoom and do short PR reviews async during the week.

Happy to pay a reasonable weekly or monthly fee.
