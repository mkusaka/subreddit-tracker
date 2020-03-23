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
## [2][JavaScript file structure in Rails 6?](https://www.reddit.com/r/rails/comments/fna1n4/javascript_file_structure_in_rails_6/)
- url: https://www.reddit.com/r/rails/comments/fna1n4/javascript_file_structure_in_rails_6/
---
Hello,

I'm developing a Rails application and I'm at the point of development where I need to add some JS. I can't find much documentation or tutorials how to do that correctly. I need to have some page specific JavaScript and the best thing I found so far was:
- under the packs folder create a `page1.js` where I have my JavaScript for page1
- in the view of page1 (`page1.html.erb`) include that pack with a JavaScript pack tag 

Allthough it works, it doesn't seem like the best solution and also sometimes this `page1.js` appears to be running on other pages as well.

Can you guys help me with this? 
Thanks a lots
## [3][To add different texts in the same popup-view](https://www.reddit.com/r/rails/comments/fniyhd/to_add_different_texts_in_the_same_popupview/)
- url: https://www.reddit.com/r/rails/comments/fniyhd/to_add_different_texts_in_the_same_popupview/
---
I have a popup in views/application/generic\_popup.html.erb

    &lt;div&gt;Warning&lt;/div&gt;
    &lt;div class="popup_body"&gt;text to show&lt;/div&gt;
    

I want to open this popup in my page on `views/cars/show.html.erb`

    &lt;%= link_to generic_popup_path, data: { load_popup: true, remote: true } do %&gt;Submit&lt;% end %&gt;

and I want to show in the `&lt;div class="popup_body"&gt;`, in generic\_popup, another text, for example `&lt;div class="popup_body"&gt;you can't submit it&lt;/div&gt;`

&amp;#x200B;

I want to open the SAME popup also in another page on `views/drivers/show.html.erb`

and to show A DIFFERENT TEXT always replacing `&lt;div class="popup_body"&gt;`, in generic\_popup, for example `&lt;div class="popup_body"&gt;The driver is stupid&lt;/div&gt;`

**How to do? Is it possible?**
## [4][Interacting with sass-lang?](https://www.reddit.com/r/rails/comments/fn7vf9/interacting_with_sasslang/)
- url: https://www.reddit.com/r/rails/comments/fn7vf9/interacting_with_sasslang/
---
Is there a way to interact with scss file? like passing a variable or a value to scss to determine with file to import in that scss?
## [5][how to load all user email in background.](https://www.reddit.com/r/rails/comments/fnemmw/how_to_load_all_user_email_in_background/)
- url: https://www.reddit.com/r/rails/comments/fnemmw/how_to_load_all_user_email_in_background/
---
i need when create post immediate send mail not waiting when search . appear this solution load all email in background . i need how to load all user mails in background . i use `delayed job gem`
## [6][Rails migration runs without errors but doesnt create tsvector column, GIN index or TRIGGER](https://www.reddit.com/r/rails/comments/fmwhoe/rails_migration_runs_without_errors_but_doesnt/)
- url: https://www.reddit.com/r/rails/comments/fmwhoe/rails_migration_runs_without_errors_but_doesnt/
---
Hi,

I'm currently implementing search functionality in my rails app using pg\_search for PostgreSQL full text search. Unfortunately, I'm having a problem with getting the migration to add the tsvector column, GIN INDEX and the TRIGGER in the listings table. The migration runs successfully but doesn't create the things specified in the migration file.

The migration file:

    class AddTsVectorColumns &lt; ActiveRecord::Migration[6.0]
      def change
        def up
          add_column :listings, :tsv, :tsvector
          add_index :listings, :tsv, using: 'gin'
    
          execute &lt;&lt;-SQL
            CREATE OR REPLACE FUNCTION update_listings_tsv() RETURNS trigger AS $$
            BEGIN
              new.tsv := (
                SELECT
                  setweight(to_tsvector(l.item_name), 'A') ||
                  setweight(to_tsvector(coalesce((string_agg(brands.name, ' ')), '')), 'A') ||
                  setweight(to_tsvector(coalesce((string_agg(colours.name, ' ')), '')), 'B') ||
                  setweight(to_tsvector(l.description), 'B') ||
                  setweight(to_tsvector(categories.name), 'B') ||
                  setweight(to_tsvector(sub_categories.name), 'B') ||
                  setweight(to_tsvector(sizes.name), 'B') ||
                  setweight(to_tsvector(users.username), 'C')
    
    
                FROM listings l
                JOIN users ON users.id = l.user_id
                JOIN categories ON categories.id = l.category_id
                JOIN sub_categories ON sub_categories.id = l.sub_category_id
                JOIN sizes ON sizes.id = l.size_id
                JOIN conditions ON conditions.id = l.condition_id
    
                -- Associative Tables
                JOIN brands_listings ON brands_listings.listing_id = brands_listings.brand_id
                JOIN brands ON brands.id = brands_listings.brand_id
                JOIN colours_listings ON colours_listings.listing_id = colours_listings.colour_id
                JOIN colours ON colours.id = colours_listings.colour_id
                WHERE l.id = new.id
                GROUP BY l.id, users.id, categories.id, sub_categories.id, sizes.id, conditions.id
              );
              RETURN new;
            END;
            $$ LANGUAGE plpgsql;
    
            CREATE TRIGGER tsvectorupdate BEFORE INSERT OR UPDATE
            ON listings FOR EACH ROW EXECUTE PROCEDURE update_listings_tsv();
          SQL
        end
    
        def down
          execute &lt;&lt;-SQL
            DROP TRIGGER tsvectorupdate
            ON listings
          SQL
    
          remove_index :listings, :tsv
          remove_column :listings, :tsv
        end
      end
    end

The output of the migration:

    D, [2020-03-22T13:23:39.965602 #70710] DEBUG -- :   primary::SchemaMigration Create (0.4ms)  INSERT INTO "schema_migrations" ("version") VALUES ($1) RETURNING "version"  [["version", "20200313140854"]]
    D, [2020-03-22T13:23:39.968646 #70710] DEBUG -- :    (0.9ms)  COMMIT
    I, [2020-03-22T13:23:39.969992 #70710]  INFO -- : Migrating to AddTsVectorColumns (20200321060306)
    == 20200321060306 AddTsVectorColumns: migrating ===============================
    == 20200321060306 AddTsVectorColumns: migrated (0.0000s) ======================
    
    D, [2020-03-22T13:23:39.973220 #70710] DEBUG -- :    (0.4ms)  BEGIN
    D, [2020-03-22T13:23:39.975042 #70710] DEBUG -- :   primary::SchemaMigration Create (0.4ms)  INSERT INTO "schema_migrations" ("version") VALUES ($1) RETURNING "version"  [["version", "20200321060306"]]
    D, [2020-03-22T13:23:39.977113 #70710] DEBUG -- :    (0.4ms)  COMMIT
    D, [2020-03-22T13:23:39.986874 #70710] DEBUG -- :   ActiveRecord::InternalMetadata Load (0.8ms)  SELECT "ar_internal_metadata".* FROM "ar_internal_metadata" WHERE "ar_internal_metadata"."key" = $1 LIMIT $2  [["key", "environment"], ["LIMIT", 1]]
    D, [2020-03-22T13:23:39.993805 #70710] DEBUG -- :    (0.2ms)  BEGIN
    D, [2020-03-22T13:23:39.999419 #70710] DEBUG -- :   ActiveRecord::InternalMetadata Create (4.2ms)  INSERT INTO "ar_internal_metadata" ("key", "value", "created_at", "updated_at") VALUES ($1, $2, $3, $4) RETURNING "key"  [["key", "environment"], ["value", "development"], ["created_at", "2020-03-22 03:23:39.992922"], ["updated_at", "2020-03-22 03:23:39.992922"]]
    D, [2020-03-22T13:23:40.002363 #70710] DEBUG -- :    (0.7ms)  COMMIT
    D, [2020-03-22T13:23:40.003788 #70710] DEBUG -- :    (0.4ms)  SELECT pg_advisory_unlock(6123583507257778380)
    D, [2020-03-22T13:23:40.258046 #70710] DEBUG -- :    (0.6ms)  SELECT "schema_migrations"."version" FROM "schema_migrations" ORDER BY "schema_migrations"."version" ASC

Error when trying to access the tsv column (which should have been created):

    I, [2020-03-22T10:49:43.273416 #64203]  INFO -- : Started GET "/api/v1/public/search?q=Adidas" for ::1 at 2020-03-22 10:49:43 +1000
    D, [2020-03-22T10:49:43.367465 #64203] DEBUG -- :    (6.1ms)  SELECT "schema_migrations"."version" FROM "schema_migrations" ORDER BY "schema_migrations"."version" ASC
    I, [2020-03-22T10:49:43.383171 #64203]  INFO -- : Processing by Api::V1::Public::ListingsController#search as */*
    I, [2020-03-22T10:49:43.383234 #64203]  INFO -- :   Parameters: {"q"=&gt;"Adidas"}
    D, [2020-03-22T10:49:43.638039 #64203] DEBUG -- :   Listing Load (12.4ms)  SELECT "listings".* FROM "listings" INNER JOIN (SELECT "listings"."id" AS pg_search_id, (ts_rank(("listings"."tsv"), (to_tsquery('simple', ''' ' || 'Adidas' || ' ''' || ':*')), 0)) AS rank FROM "listings" WHERE (("listings"."tsv") @@ (to_tsquery('simple', ''' ' || 'Adidas' || ' ''' || ':*')))) AS pg_search_8a836f245cd6a84ba9cbd1 ON "listings"."id" = pg_search_8a836f245cd6a84ba9cbd1.pg_search_id ORDER BY pg_search_8a836f245cd6a84ba9cbd1.rank DESC, "listings"."id" ASC LIMIT $1  [["LIMIT", 100]]
    D, [2020-03-22T10:49:43.638832 #64203] DEBUG -- :   ↳ app/controllers/application_controller.rb:47:in `public_render_listings_helper'
    I, [2020-03-22T10:49:43.639121 #64203]  INFO -- : Completed 500 Internal Server Error in 256ms (ActiveRecord: 22.8ms | Allocations: 45123)
    
    
    F, [2020-03-22T10:49:43.640013 #64203] FATAL -- :
    ActiveRecord::StatementInvalid (PG::UndefinedColumn: ERROR:  column listings.tsv does not exist
    LINE 1: ...SELECT "listings"."id" AS pg_search_id, (ts_rank(("listings"...
                                                                 ^
    ):

Extract from the sql schema (with config.active\_record.schema\_format = :sql) -- -- PostgreSQL database dump --

    SET statement_timeout = 0;
    SET lock_timeout = 0;
    SET client_encoding = 'UTF8';
    SET standard_conforming_strings = on;
    SET check_function_bodies = false;
    SET client_min_messages = warning;
    SET row_security = off;
    
    --
    -- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: -
    --
    
    CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;
    
    
    --
    -- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: -
    --
    
    COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';
    
    
    SET search_path = public, pg_catalog;
    
    SET default_tablespace = '';
    
    SET default_with_oids = false;
    
    --
    -- Name: listings; Type: TABLE; Schema: public; Owner: -
    --
    
    CREATE TABLE listings (
        id bigint NOT NULL,
        item_name character varying,
        description character varying,
        price integer,
        shipping_cost integer,
        created_at timestamp without time zone NOT NULL,
        updated_at timestamp without time zone NOT NULL,
        category_id bigint,
        sub_category_id bigint,
        size_id bigint,
        user_id bigint,
        full_item_name text,
        cached_votes_total integer DEFAULT 0,
        cached_votes_score integer DEFAULT 0,
        cached_votes_up integer DEFAULT 0,
        cached_votes_down integer DEFAULT 0,
        cached_weighted_score integer DEFAULT 0,
        cached_weighted_total integer DEFAULT 0,
        cached_weighted_average double precision DEFAULT 0.0,
        flaremeter integer DEFAULT 0,
        condition_id bigint,
        cover_photo_data text
    );
    
    
    --
    -- Name: listings_id_seq; Type: SEQUENCE; Schema: public; Owner: -
    --
    
    CREATE SEQUENCE listings_id_seq
        START WITH 1
        INCREMENT BY 1
        NO MINVALUE
        NO MAXVALUE
        CACHE 1;
    
    
    --
    -- Name: listings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
    --
    
    ALTER SEQUENCE listings_id_seq OWNED BY listings.id;
    
    
    --
    -- PostgreSQL database dump complete
    --
    
    SET search_path TO "$user", public;
    
    INSERT INTO "schema_migrations" (version) VALUES
    ('20200223103428'),
    ('20200313111818'),
    ('20200313140854'),
    ('20200321060306');
## [7][VS Code Extention help](https://www.reddit.com/r/rails/comments/fmwbw4/vs_code_extention_help/)
- url: https://www.reddit.com/r/rails/comments/fmwbw4/vs_code_extention_help/
---
What is the extention used by pressing Ctrl + Shift to auto complete CSS classes in erb files.  


ex - 

 `.nav navbar`     \-&gt;  Ctrl + Shift -&gt; `&lt;div class="nav navbar"&gt;&lt;/div&gt;`
## [8][Sessions + Devise - [How to force sign out a user from all its session?]](https://www.reddit.com/r/rails/comments/fmx7ve/sessions_devise_how_to_force_sign_out_a_user_from/)
- url: https://www.reddit.com/r/rails/comments/fmx7ve/sessions_devise_how_to_force_sign_out_a_user_from/
---
I want to clear all sessions of a user on some attribute changes(email, password, etc.)

I use the ActiveRecord sessions table along with devise to handle user sessions. Since a user can have multiple sessions and each session is a unique entry in the table it's becoming difficult for me to figure out the most optimum approach to clear all the sessions for the user.

&amp;#x200B;

Note - the table is quite big so iterating over the sessions and doing marshal dump in a batch won't be helpful.
## [9][whats the point of integration test?](https://www.reddit.com/r/rails/comments/fmrycp/whats_the_point_of_integration_test/)
- url: https://www.reddit.com/r/rails/comments/fmrycp/whats_the_point_of_integration_test/
---
If you are already doing unit testing and system testing, whats the value of doing integration testing? How can multiple components together fail an integration test when they each pass their own unit test (and the system test passes too)?
## [10][Are Companies still hiring?](https://www.reddit.com/r/rails/comments/fm6nwe/are_companies_still_hiring/)
- url: https://www.reddit.com/r/rails/comments/fm6nwe/are_companies_still_hiring/
---
I am a part time Rails developer searching for a full time role. It looks like I am facing some tough trends with companies everywhere closing and doing layoffs. 
I am seeing Jr developers getting laid off everywhere. Do I stand a chance in the job search right now with less than a year of professional experience.
## [11][Wanting to help other RoR developers](https://www.reddit.com/r/rails/comments/fm6vfn/wanting_to_help_other_ror_developers/)
- url: https://www.reddit.com/r/rails/comments/fm6vfn/wanting_to_help_other_ror_developers/
---
I really enjoy helping other programmers so i'd like to offer my time to anyone is stuck on something in Rails. I've been developing in RoR for 6 years and own a small custom software development company working with some pretty large projects.

I mostly enjoy working with with the backend focusing on code design and testing in rspec, but if I can help you in something else, I'm happy to give it a shot. Even if you're brand new and it's just a discussion to ask some questions and maybe help connect the dots for you.

Some requirements:

* I'd like to video conference the session (you sharing your screen) so we can record it and pop it up on YouTube for others. I'd like to keep it less than 2 hours.
* You would need a mic/camera and good internet connection
* I think it would be best suited those who are a beginner/intermediate in the framework. Preferably I'd like us to focus on an actual problem in the session rather than it turning into a straight video tutorial.
* It would be helpful to get an idea on what the challenge you're facing, just so I know I can actually help you and not waste our time

If you're keen just leave a comment or send me a PM with your challenge and I'll be in touch. If it's not within my skillset maybe someone else might offer some time!

&amp;#x200B;

EDIT:  I'm not looking to charge anything, I'm just trying to stay busy, take a break from personal projects, and avoid video games.
