# ruby
## [1][Platformer game written in Ruby - Super Bombinhas](https://www.reddit.com/r/ruby/comments/fxqvdf/platformer_game_written_in_ruby_super_bombinhas/)
- url: https://www.reddit.com/r/ruby/comments/fxqvdf/platformer_game_written_in_ruby_super_bombinhas/
---
Hi!

I'd like to share a demo of my platformer game, it is [open source](https://github.com/victords/super-bombinhas) and completely written in Ruby, with the help of the [Gosu](https://www.libgosu.org/) and [MiniGL](https://github.com/victords/minigl) gems. Also, the Windows executable was created with the [Ocra](https://github.com/larsch/ocra/) gem.

Windows: [https://drive.google.com/open?id=1KWRfnx6eGbj37OK7dd6Pk8uVzRtIWWTx](https://drive.google.com/open?id=1KWRfnx6eGbj37OK7dd6Pk8uVzRtIWWTx)

Linux (Debian-based): [https://drive.google.com/open?id=1tjh\_4akR5b1QpAMqtG1zJHJ1YhW5i2T8](https://drive.google.com/open?id=1tjh_4akR5b1QpAMqtG1zJHJ1YhW5i2T8)

I'd appreciate any feedback. Thanks!
## [2][Minimalist tool to automate the writing of unit tests.](https://www.reddit.com/r/ruby/comments/fxevwc/minimalist_tool_to_automate_the_writing_of_unit/)
- url: https://github.com/fixrb/brutal
---

## [3][How would i read a certain number of lines from a file?](https://www.reddit.com/r/ruby/comments/fxncta/how_would_i_read_a_certain_number_of_lines_from_a/)
- url: https://www.reddit.com/r/ruby/comments/fxncta/how_would_i_read_a_certain_number_of_lines_from_a/
---
Hey guys,

So I'm doing some ruby work for uni and one of the tasks is to read a file and output a certain amount of lines in that file. It'll hopefully be explained better in the code below (the # is the instructions pretty much) :

    def write_data_to_file(a_file)
       a_file.puts('5') # 5 is the number of times it has to loop
       a_file.puts('Fred')
       a_file.puts('Sam')
       a_file.puts('Jill')
       a_file.puts('Jenny')
       a_file.puts('Zorro')
    end
    
    # I need to edit the following
    # so that it uses a loop which repeats
    # acccording to the number of lines in the File
    # which is given in the first line of the File
    
    def read_data_from_file(a_file)
      count = a_file.gets.to_i
      puts count.to_s
      puts a_file.gets
      puts a_file.gets
      puts a_file.gets
      puts a_file.gets
      puts a_file.gets
    end
    
    # I need to shorten the following code in main() as well but I'm pretty sure that it's already the best it can be
    
    def main
      a_file = File.new("mydata.txt", "w") # open for writing
      write_data_to_file(a_file)
      a_file.close
      
      a_file = File.new("mydata.txt", "r") # open for reading
      read_data_from_file(a_file)
      a_file.close
    end
    
    main()

I have no idea how to go about this. i was thinking changing the code in read\_data from\_file(a\_file) to somethiing like:

    def read_data_from_file(a_file)
      File.foreach('mydata.txt').with_index do |line|
         puts "#{line}"
      end
    end

but after that i am extremely clueless.

Any help is appreciated,

thank you :)

&amp;#x200B;

EDIT:

after using my brain for 10 minutes i realised it wasnt that hard and all i needed to do was create a simple loop based on the integer printed at the start (5) and use a count operation. 

i used :

    def read_data_from_file()
      a_file = File.new("mydata.txt", "r") # open for reading
      count = a_file.gets.to_i
      i=0
      while (i&lt;count)
        puts a_file.gets
        i=i+1
      end
      a_file.close
    end

in case anyone wants to know :)
## [4][has_one + belongs_to not working](https://www.reddit.com/r/ruby/comments/fxm6eo/has_one_belongs_to_not_working/)
- url: https://www.reddit.com/r/ruby/comments/fxm6eo/has_one_belongs_to_not_working/
---
I’m getting confused on the correct structure for my migration tables and the models.

I have three models: `Session`, `Paste` and `Invitee`.

A `Paste` is the main model for the user on the platform.

The user can also invite people via e-mail to give view access, which is called an `Invitee`.

A `Session` is created for the owner of the paste when the paste is created via `@paste.session.create({ :sid =&gt; … })` and also for the `Invitee` via: `@invitee.session.create({ :sid =&gt; … })`. So a `Session` can be used via other models such as a `Paste` or an `Invitee`.

This is what I think the relationship should be:

```
class Paste &lt; ActiveRecord::Base
  has_one :session
  has_many :invitees
end

class Invitee &lt; ActiveRecord::Base
  has_one :session
end

class Session &lt; ActiveRecord::Base
  belongs_to :paste
  belongs_to :invitee, :through =&gt; :paste
end
```

I’ve looked into polymorphic relationships, not sure if this is recommended for this use-case or not.

Unfortunately, when doing `@paste.session.create()` it says:

```
NoMethodError Exception: undefined method `create' for nil:NilClass
```

These are my tables:

Pastes
```
class CreatePastes &lt; ActiveRecord::Migration[6.0]
  def change
    create_table :pastes do |t|
      t.timestamps
    end
  end
end
```

Invitees
```
class CreateInvitees &lt; ActiveRecord::Migration[6.0]
  def change
    create_table :invitees do |t|
      t.string :email
      t.references :paste, foreign_key: true
      t.timestamps
    end
  end
end
```

Sessions
```
class CreateSessions &lt; ActiveRecord::Migration[6.0]
  def change
    create_table :sessions do |t|
      t.string :sid
      t.references :paste, foreign_key: true
      t.timestamps
    end
  end
end
```

This is the result of my `db/schema.rb`:
```
create_table "invitees", force: :cascade do |t|
  t.string "email"
  t.integer "paste_id"
  t.datetime "created_at", precision: 6, null: false
  t.datetime "updated_at", precision: 6, null: false
  t.index ["paste_id"], name: "index_invitees_on_paste_id"
end

create_table "pastes", force: :cascade do |t|
  t.datetime "created_at", precision: 6, null: false
  t.datetime "updated_at", precision: 6, null: false
end

create_table "sessions", force: :cascade do |t|
  t.string "sid"
  t.integer "paste_id"
  t.datetime "created_at", precision: 6, null: false
  t.datetime "updated_at", precision: 6, null: false
  t.index ["paste_id"], name: "index_sessions_on_paste_id"
end

add_foreign_key "invitees", "pastes"
add_foreign_key "sessions", "pastes"
```
## [5][Heredocs and how to use them in Ruby and Rails](https://www.reddit.com/r/ruby/comments/fxb8ks/heredocs_and_how_to_use_them_in_ruby_and_rails/)
- url: https://blog.saeloun.com/2020/04/08/heredoc-in-ruby-and-rails
---

## [6][Split method in Ruby](https://www.reddit.com/r/ruby/comments/fxd17z/split_method_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fxd17z/split_method_in_ruby/
---
I have an array: array = \["John Smith", "Bill Taylor", Troy Tate"\]

I am trying to use the split method to print either the first name or last name only from each name in the array. I am using the following, but the array.split(" ") is obviously only splitting the strings into two at the whitespace.

How can I use the '.string' method to only print either the first or last name. Or first character for that matter?

    array = ["John Smith", "Bill Taylor", Troy Tate"]
    array.each do |array|   
        puts array.split(" ") 
    end
## [7][Step by step guide on how to upgrade a Rails application](https://www.reddit.com/r/ruby/comments/fx8mg6/step_by_step_guide_on_how_to_upgrade_a_rails/)
- url: https://www.fastruby.io/blog/rails/upgrade/our-rails-upgrade-process.html
---

## [8][The Citadel Architecture at AppSignal](https://www.reddit.com/r/ruby/comments/fx8uh0/the_citadel_architecture_at_appsignal/)
- url: https://blog.appsignal.com/2020/04/08/the-citadel-architecture-at-appsignal.html
---

## [9][Which js script library should I use](https://www.reddit.com/r/ruby/comments/fxf8in/which_js_script_library_should_i_use/)
- url: https://www.reddit.com/r/ruby/comments/fxf8in/which_js_script_library_should_i_use/
---
We are building an interactive website as an assignment and we’re using ruby and Sinatra as a building environment.
Our instructor told us that we can use bootstrap library for css but not for javascript as it will cause us trouble.
What open source libraries are most compatible for this project or should I create the elements from scratch using internal js (that would be very complex)
## [10][How to build a conference line with Twilio and Ruby](https://www.reddit.com/r/ruby/comments/fwzum1/how_to_build_a_conference_line_with_twilio_and/)
- url: https://www.reddit.com/r/ruby/comments/fwzum1/how_to_build_a_conference_line_with_twilio_and/
---
It's a weird world we live in right now, so much remote work, so many people working from home. There's been an increase in the number of video calls, but what if you just want a voice conference call without choosing an app? You [can set up a conference line easily with Twilio, and this is how to use Ruby and Rails to do so](https://www.twilio.com/blog/build-conference-line-twilio-ruby).
