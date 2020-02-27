# ruby
## [1][Introducing Hanami::API](https://www.reddit.com/r/ruby/comments/f9tmgv/introducing_hanamiapi/)
- url: https://www.reddit.com/r/ruby/comments/f9tmgv/introducing_hanamiapi/
---
It's a minimal, extremely fast, lightweight Ruby framework for HTTP APIs.

(The article contains benchmark results)

[http://hanamirb.org/blog/2020/02/26/introducing-hanami-api.html](http://hanamirb.org/blog/2020/02/26/introducing-hanami-api.html)
## [2][Freelancing in Rails](https://www.reddit.com/r/ruby/comments/fa7p6q/freelancing_in_rails/)
- url: https://www.reddit.com/r/ruby/comments/fa7p6q/freelancing_in_rails/
---
hey.. how are you all?? I hope you all doing fine.. I want advice from experience people.. Anyone here doing freelancing work in ruby on rails?? can you tell me the struggles of finding a rails job in freelancing?? if someone learning rails and want a career of rails in freelancing, what advice would you like to give??
## [3][Checking something in the db (SQLite3) + Sinatra](https://www.reddit.com/r/ruby/comments/fa29sf/checking_something_in_the_db_sqlite3_sinatra/)
- url: https://www.reddit.com/r/ruby/comments/fa29sf/checking_something_in_the_db_sqlite3_sinatra/
---
Greetings.

I have a route a like this :

    post '/login' do 
     # Doing something here 

And a SQLite3 database.

I run this query:

`db.execute("SELECT password FROM users where password=?", password)`

And it seems fine, but when I want to check if it's what I entered in the form or not, Sinatra gives me an error and says as the query is a `nilClass`, it can't do anything on the damn thing.

I just want to make sure the password is in the DB or not.

P.S : It's a learning project, so I don't want to use any ORM.
## [4][I've made a simple state machine in a class](https://www.reddit.com/r/ruby/comments/f9w7n5/ive_made_a_simple_state_machine_in_a_class/)
- url: https://www.reddit.com/r/ruby/comments/f9w7n5/ive_made_a_simple_state_machine_in_a_class/
---
&amp;#x200B;

https://i.redd.it/nnpm6kj1taj41.gif

here is the code in action

here is the code

    class Player
        attr_gtk
        attr_accessor :grid, :inputs, :state, :outputs
    
        def tick 
            defaults
            render
            process_inputs
            calc_extended_stats
            move_player
        end
    
        def defaults
            state.pos_x ||= state.wlk_spd0
            state.pos_y ||= 400
            state.hero_sheet ||= "kalvgv idle.png"
            state.anim_spd ||= 20
            state.pose ||= 1
            state.wlk_spd ||= 10
            state.body ||= 3
            state.mind ||= 4
            state.step_x ||= 29
            state.step_y ||= 13
            state.is_walking ||= false
            state.sides ||= 0
            state.verts ||= 0
            state.name ||= "Kalvgv"
        end
    
        def character ordinal_x, ordinal_y, x, y, file, sizx, sizy, w1, h1
            [
              {
                x: x,
                y: y,
                w: w1,
                h: h1,
                path: "sprites/#{file}",
                tile_x: ordinal_x * sizx,
                tile_y: ordinal_y * sizy,
                tile_w: sizx,
                tile_h: sizy
              },
              #[x, y, 64, 32, 'sprites/watermark.png'],
            ]
          end
    
          def calc_extended_stats
            state.health ||= 1
            state.health = (state.body * 1.5) + state.mind + 2.5
            state.damage = ((state.body + state.mind)/2).floor
            state.wlk_spd = 16 - (state.body + state.mind)
            state.step_x = 29 + state.body
            state.step_y = 13 + state.body
          end
    
          def render
            outputs.sprites &lt;&lt; character(state.tick_count.div(state.anim_spd).mod(7),state.pose,state.pos_x, state.pos_y,state.hero_sheet, 18, 24,36,48)
            outputs.sprites &lt;&lt; [1080,0,200,720,"sprites/character_sheet.png"]
            outputs.labels &lt;&lt; [1160,720-32,state.name,10,23,150,255]
            outputs.labels &lt;&lt; [1080+20,720-96,"Body: "+state.body.to_s,10,23,120,255]
            outputs.labels &lt;&lt; [1080+20,720-96-32,"Mind: "+state.mind.to_s,10,23,120,255]
            outputs.labels &lt;&lt; [1080+20,720-96-32-64-32,"HP: "+state.health.to_s,10,23,120,255]
            outputs.labels &lt;&lt; [1080+20,720-96-32-64-64,"SPD: "+state.anim_spd.to_s,10,23,120,255]
            outputs.labels &lt;&lt; [1080+20,720-96-32-64-96,"Step: "+((state.step_x + state.step_y)* 0.5).to_s,10,23,120,255]
            outputs.labels &lt;&lt; [1080+20,720-96-96-96-32,"Damage: "+state.damage.to_s,10,23,120,255]
          end
    
          def move_player
            if state.is_walking == true &amp;&amp; state.tick_count % (state.anim_spd) == 1
                state.pos_x += state.sides * state.step_x
                state.pos_y += state.verts * state.step_y
            end
          end
    
          def process_inputs
            if inputs.keyboard.key_up.up || inputs.keyboard.key_up.down || inputs.keyboard.key_up.right || inputs.keyboard.key_up.left
                state.hero_sheet = "kalvgv idle.png"
                state.anim_spd = 60 - (state.body + 10 + state.mind + 10) - (state.mind * 2)
                state.is_walking = false
            end
    
            if state.is_walking == false &amp;&amp; state.verts != 0
                state.sides = 0
                state.verts = 0
            end
    
            if inputs.keyboard.key_down.up || inputs.keyboard.key_down.down || inputs.keyboard.key_down.right || inputs.keyboard.key_down.left
                state.hero_sheet = "kalvgv_walk_sheet.png"
                state.anim_spd = state.wlk_spd
                state.is_walking = true
            end
            if inputs.keyboard.key_held.up
                #state.pos_y += state.step_y
                #state.pos_x += state.step_x
                state.sides = 1
                state.verts = 1
                puts state.pos_y
                puts state.pos_x
                state.pose = 3
            end
            if inputs.keyboard.key_held.down
                #state.pos_y -= state.step_y
                #state.pos_x -= state.step_x
                state.sides = -1
                state.verts = -1
                puts state.pos_y
                puts state.pos_x
                state.pose = 0
            end
            if inputs.keyboard.key_held.left
               #state.pos_x -= state.step_x
                #state.pos_y += state.step_y
                state.sides = -1
                state.verts = 1
                puts state.pos_x
                puts state.pos_y
                state.pose = 2
            end
            if inputs.keyboard.key_held.right 
                #state.pos_x += state.step_x
                #state.pos_y -= state.step_y
                state.sides = 1
                state.verts = -1
                puts state.pos_x
                puts state.pos_y
                state.pose = 1
            end
          end
        end

basically the character, Player is created, and the defaults are set, render is made, then we calculate the stats from the two investable stats, body and mind, then we move the player. I know there should be moving the player before the stats, but it's just a temporary system. eventually I'll add leveling for the player and that will trigger stat calculation
## [5][started to learn ruby and run into a bit of a problem](https://www.reddit.com/r/ruby/comments/f9qr4i/started_to_learn_ruby_and_run_into_a_bit_of_a/)
- url: https://www.reddit.com/r/ruby/comments/f9qr4i/started_to_learn_ruby_and_run_into_a_bit_of_a/
---
hey guys,

I'm a noob at ruby and just wondering why the following code doesn't work.

&amp;#x200B;

char\_name = "john"

char\_age = "40"

\#---------------------------Story Begins Here----------------------------------#

puts "this is a story"

puts ("all about "+char\_name+" who was "+char\_age+"")

puts (""+char\_name+" was "+char\_age+" and he liked apples")

puts ("because he was "+char\_age+" and his name was "+char\_name+" ")

puts "this is\\nmostly true"

phrase = "but who really knows?"

puts phrase

\#puts phrase.include? "?"

\#puts phrase \[15,20\]

\#-------------------------------New Chapter------------------------------------#

puts "hello #{char\_name}, are you here to buy apples?"

num\_apples = "#{1+3}"

num\_apples.gsub ("#{1+3}", "#{1+2}")

puts "yes i would like to buy #{num\_apples} apples please"

&amp;#x200B;

No error shows up but the last part with the variable num\_apples.gsub doesnt change the amount from 4 to 3. is there a different way i should be doing this or can this just not be used for numbers and math.

&amp;#x200B;

thank you for helping :)
## [6][Accessing data between several Rails apps](https://www.reddit.com/r/ruby/comments/f9qxen/accessing_data_between_several_rails_apps/)
- url: https://www.reddit.com/r/ruby/comments/f9qxen/accessing_data_between_several_rails_apps/
---
The company I work for has the domain split between a few monoliths, each is its own Rails app with its own daabase. All of the apps depend on one monolith that handles users, roles and other things, let's call it UsersApp.

Now let's say you have BillingApp, and it needs to pull the data from UsersApp. For example forms in BullingApp need to have the user's data to show it in the form drop downs etc., what ways do you know of handling this issue and what are the tradeoffs? In our case the relationship is read only (BillingApp always reads and never writes), but I'd like to get a more general discussion going.

Just to help get a discussion going here are the ways I know of:

1. Using ActiveResource in BillingApp to access everything, or similar technique to just pull everything through an API that UsersApp exposes.
2. Pulling and syncing the data from UsersApp into BillingApp (basically duplicating the database tables in BillingApp). Let's say the data isn't time sensitive, the sync can happen once a day so "real timeness" isn't an issue here.
3. Packing UsersApp as a gem, and having BillingApp access the models/data directly.
## [7][.destroy Getting Killed](https://www.reddit.com/r/ruby/comments/f9ll24/destroy_getting_killed/)
- url: https://www.reddit.com/r/ruby/comments/f9ll24/destroy_getting_killed/
---
I'm running into an issue where I have a model with many children (has\_many) relationships. 

When I call .destroy on this model, the process grows so large that it's getting killed.

I've tried overloading .destroy on that model and some of the associated model to use a `.each(&amp;:destroy)` where possible but that did not solve the issue.

Any suggestions on destroying a model where the associated models have 100k-1M rows?
## [8][Ruby one of the highest-paid programming languages globally in 2020](https://www.reddit.com/r/ruby/comments/f9aumo/ruby_one_of_the_highestpaid_programming_languages/)
- url: https://learnworthy.net/highest-paid-programming-languages-in-2020/
---

## [9][Logging Accessible via Active Record](https://www.reddit.com/r/ruby/comments/f9lizu/logging_accessible_via_active_record/)
- url: https://www.reddit.com/r/ruby/comments/f9lizu/logging_accessible_via_active_record/
---
Hey all! 

I'm looking to have logs written in a way that reads are accessible via Active Record but I'm not sure what the options are for this.

For example:  
`User.find(123).logs.first.lines.where(type: 'Info').where('message like ?', 'My Error Message')`

In short, I'd like to write logs to a standard SQL database, without the overhead of writing logs to a SQL database. (the Lines model includes +1M records)

I'm assuming this is possible but my Googling is coming up short. 

Can any recommend a technology, method of doing this or direction that I should be looking?

Thanks in advance!
## [10][Ruby Conferences &amp; Camps in 2020 - What's Upcoming? Anything Missing? Updates Welcome](https://www.reddit.com/r/ruby/comments/f9fohh/ruby_conferences_camps_in_2020_whats_upcoming/)
- url: https://www.reddit.com/r/ruby/comments/f9fohh/ruby_conferences_camps_in_2020_whats_upcoming/
---
Hello,

   To celebrate ruby's birthday I've updated the [Ruby Conferences &amp; Camps in 2020 - What's
Upcoming?](http://planetruby.github.io/calendar/2020) page. Try the [rubyconf command line tool](https://github.com/textkit/whatson) (packaged up in the whatson gem) e.g.

    $ rubyconf

printing as of Feb/25th:

    Upcoming Ruby Conferences:
    
    in 24d   Wrocław &lt;3 Ruby (wroclove.rb), Fri-Sun Mar/20-22 (3d) @ Wrocław, Poland
    in 31d   Ruby Retreat New Zealand (NZ), Fri-Mon Mar/27-30 (4d) @ Mt Cheeseman (near Christchurch), New Zealand
    in 37d   RubyDay Italy, Thu Apr/2 (1d) @ Verona, Veneto, Italy
    in 38d   Ruby by the Bay (Ruby for Good, West Coast Edition), Fri-Mon Apr/3-6 (4d) @ Marin Headlands (near San Francisco), California, United States
    in 39d   Ruby Wine 2.0, Sat Apr/4 (1d) @ Chisinau, Moldova
    in 44d   RubyKaigi, Thu-Sat Apr/9-11 (3d) @ Nagano, Japan
    in 53d   RubyConf Belarus (BY), Sat Apr/18 (1d) @ Minsk, Belarus
    in 60d   RubyConf India, Sat+Sun Apr/25+26 (2d) @ Goa, India
    in 70d   RailsConf (United States), Tue-Thu May/5-7 (3d) @ Portland, Oregon, United States
    in 80d   Balkan Ruby, Fri+Sat May/15+16 (2d) @ Sofia, Bulgaria
    in 102d  Ruby Unconf Hamburg, Sat+Sun Jun/6+7 (2d) @ Hamburg, Germany
    in 102d  Saint P Rubyconf, Sat+Sun Jun/6+7 (2d) @ Saint Petersburg, Russia
    in 129d  Brighton RubyConf, Fri Jul/3 (1d) @ Brighton, Sussex, England, United Kingdom
    in 142d  RubyNess, Thu+Fri Jul/16+17 (2d) @ Inverness, Scotland, United Kingdom
    in 149d  RubyConf Kenya, Thu-Sat Jul/23-25 (3d) @ Nairobi, Kenya
    in 178d  European Ruby Konference (EuRuKo), Fri+Sat Aug/21+22 (2d) @ Helsinki, Finnland
    in 189d  Rails Camp West, Tue-Fri Sep/1-4 (4d) @ Diablo Lake, Washington, United States


Anything missing? Updates welcome, see [`data/conferences2020.yml`](https://github.com/planetruby/calendar/blob/master/_data/conferences2020.yml) file
in the `planetruby/calendar` repo.

What's your favorite ruby conference or camp? Let us know. Cheers. Prost.
