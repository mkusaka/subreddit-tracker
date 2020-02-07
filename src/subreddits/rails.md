# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Ruby on Rails advanced courses?](https://www.reddit.com/r/rails/comments/f00t14/ruby_on_rails_advanced_courses/)
- url: https://www.reddit.com/r/rails/comments/f00t14/ruby_on_rails_advanced_courses/
---
Hello, I consider I have experience in Rails, but I want to go to the next level, for this I want to know if there's an advanced course on Rails which has some or most of the following topics:

* Examples of good use of rails conventions (service objects, query object, null objects, creating concerns to DRY your controllers/models, etc.).
* Make performant apps (performant queries, performant ruby code, fix n+1 queries, etc.).
* Examples of good use of design patterns (observable, chain of responsibility, factory, strategy pattern, etc.).
* Serializing data (blueprinter, fast json api, etc.)
* Make performant tests (build or build_stubbed instead of create, parallel tests, etc.)

Or even building your own gems, metaprogramming, etc. (I know this is ruby, but to be an advanced Rails developer you have to be an advanced Ruby developer too IMO)
## [3][There are two Ruby version manager RVM vs RBENV. what's your choice? and why?](https://www.reddit.com/r/rails/comments/f009mb/there_are_two_ruby_version_manager_rvm_vs_rbenv/)
- url: https://www.reddit.com/r/rails/comments/f009mb/there_are_two_ruby_version_manager_rvm_vs_rbenv/
---
I was at work, I installed ruby and rails using RBENV. 

My senior colleague made me install RVM and uninstalled RBENV. Insists that RVM handles better somehow. 

What's your thought on this? 

What is the industrial standard?
## [4][Yet another Arel Cheatsheet, but on Steroids](https://www.reddit.com/r/rails/comments/ezuuup/yet_another_arel_cheatsheet_but_on_steroids/)
- url: https://www.reddit.com/r/rails/comments/ezuuup/yet_another_arel_cheatsheet_but_on_steroids/
---
Hi guys,

I've created a cheatsheet of Arel that collects my personal experience with it. I've decided to share it since Arel is pretty much undocumented, but I think it's a really powerful tool.

[https://gist.github.com/ProGM/c6df08da14708dcc28b5ca325df37ceb](https://gist.github.com/ProGM/c6df08da14708dcc28b5ca325df37ceb)
## [5][Multiple database connections?](https://www.reddit.com/r/rails/comments/f02fxm/multiple_database_connections/)
- url: https://www.reddit.com/r/rails/comments/f02fxm/multiple_database_connections/
---
Hi. I have an ongoing project which asks for many DB connections (9 total) in:

-SQL Server
-MySQL
-Oracle SQL
-PostgreSQL

I never did something like this before, so i might ask for some advices or ideas.

I'm using Rails 6.

In resume, it's a data warehouse. I must delevelop an ETL module which process all the data from these 9 databases and puts into only one DB (which is PostgreSQL). I'd like to read your experiences, advices or something like this.
## [6][How to transition data in a JSONB column into separate columns for millions of rows in Rails/Postgres?](https://www.reddit.com/r/rails/comments/ezy875/how_to_transition_data_in_a_jsonb_column_into/)
- url: https://www.reddit.com/r/rails/comments/ezy875/how_to_transition_data_in_a_jsonb_column_into/
---
I've got a JSONB column in Postgres and I need to split out the JSONB data into individual columns. I've got all the new columns set up, but now I've got about 3 million records whose JSONB data needs to get split out into individual columns.

What's the most performant way to basically run an UPDATE on every single row?

Right now I'm doing this inside a Rails rake task:

```ruby
Track.all.find_each do |track|
  track.update_attributes(
    key: track.audio_features['key'],
    mode: track.audio_features['mode'],
    tempo: track.audio_features['tempo'],
    energy: track.audio_features['energy'],
    valence: track.audio_features['valence'],
    liveness: track.audio_features['liveness'],
    loudness: track.audio_features['loudness'],
    speechiness: track.audio_features['speechiness'],
    acousticness: track.audio_features['acousticness'],
    danceability: track.audio_features['danceability'],
    time_signature: track.audio_features['time_signature'],
    instrumentalness: track.audio_features['instrumentalness'],
  )
end
```

That...will take a while for 3 million rows. Is there a faster way to do that?
## [7][How to attach raw SQL to an existing Rails ActiveRecord chain?](https://www.reddit.com/r/rails/comments/ezwltk/how_to_attach_raw_sql_to_an_existing_rails/)
- url: https://www.reddit.com/r/rails/comments/ezwltk/how_to_attach_raw_sql_to_an_existing_rails/
---
I have a rule builder that ultimately builds up ActiveRecord queries by chaining multiple `where` calls, like so:

```rails
Track.where("tracks.popularity &lt; ?", 1).where("(audio_features -&gt;&gt; 'valence')::numeric between ? and ?", 2, 5)
```

Then, if someone wants to sort the results randomly, it would append `order("random()")`.

However, given the table size, `random()` is extremely inefficient for ordering, so I need to use Postgres TABLESAMPLE-ing. 

In a raw SQL query, that looks like this:

```sql
SELECT * FROM "tracks" TABLESAMPLE SYSTEM(0.1) LIMIT 250;
```

Is there some way to add that `TABLESAMPLE SYSTEM(0.1)` to the existing chain of ActiveRecord calls? Putting it inside a `where()` or `order()` doesn't work since it's not a WHERE or ORDER BY function.
## [8][Best way to bulk create objects?](https://www.reddit.com/r/rails/comments/ezw57a/best_way_to_bulk_create_objects/)
- url: https://www.reddit.com/r/rails/comments/ezw57a/best_way_to_bulk_create_objects/
---
I have an array of objects I want to send to my rails backend through the form of a param: `dish_section`:

    [{dish_id: 1, section_id: 1}, {dish_id: 2, section_id: 1}]

I've looked into a gem called [bulk\_insert](https://github.com/jamis/bulk_insert) but for the life of me I can't translate my para an array without getting some type of ActionControllerParameter error. I wanted to check here to see what's usually the best pattern to handle bulk inserts coming from a JSON API. I'm currently on Rails 5 for this project so `insert_all` unfortunately an option isn't an option.
## [9][Does Rails generate a JSON or HTML in tmp folder?](https://www.reddit.com/r/rails/comments/eztvim/does_rails_generate_a_json_or_html_in_tmp_folder/)
- url: https://www.reddit.com/r/rails/comments/eztvim/does_rails_generate_a_json_or_html_in_tmp_folder/
---
I'm curious as to whether Rails actually generates a Html or json temporarily when you call the "render" method in Controller
## [10][What are the implications of creating methods in model rather than in controller](https://www.reddit.com/r/rails/comments/eztu5y/what_are_the_implications_of_creating_methods_in/)
- url: https://www.reddit.com/r/rails/comments/eztu5y/what_are_the_implications_of_creating_methods_in/
---
Hi, I am following omniauth + devise integration.  
[https://github.com/heartcombo/devise/wiki/OmniAuth:-Overview](https://github.com/heartcombo/devise/wiki/OmniAuth:-Overview)  
In the above url the guide recommends to create method in User model whereas I can do User.where and User.create in controller  
I am trying to understand in what scenarios one has to create method in model and controller  
Thanks in Advance
## [11][Do you add IDs to association tables?](https://www.reddit.com/r/rails/comments/ezo6rc/do_you_add_ids_to_association_tables/)
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
