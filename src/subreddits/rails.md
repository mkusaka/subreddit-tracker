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
## [2][Best new thing to learn after Rails?](https://www.reddit.com/r/rails/comments/fo05t5/best_new_thing_to_learn_after_rails/)
- url: https://www.reddit.com/r/rails/comments/fo05t5/best_new_thing_to_learn_after_rails/
---
I heard it might be JS. Did they mean vanilla JS? God that is ugly. I keep forgetting parentheses and ending the sentence with ;  I heard also buzzwords like MongoDB, Angular. I see myself as a Rails dev, but what should a take a look next? I would think some JS framework, and only using Rails as an API? 

&amp;#x200B;

What technology do you think is really useful to learn when working with Ruby on Rails?
## [3][Babel Es6 With Rails 6](https://www.reddit.com/r/rails/comments/fo4y2g/babel_es6_with_rails_6/)
- url: https://www.reddit.com/r/rails/comments/fo4y2g/babel_es6_with_rails_6/
---
Hi

I can't seem to get react written in ES6 style to compile with Rails 6 and Webpacker. Maybe I'm missing a config? Not sure how to configure babel to tell it to understand ES6.
## [4][How would one go about finding Jr Dev Position?](https://www.reddit.com/r/rails/comments/fnzvnd/how_would_one_go_about_finding_jr_dev_position/)
- url: https://www.reddit.com/r/rails/comments/fnzvnd/how_would_one_go_about_finding_jr_dev_position/
---
I (almost) got 2 personal projects, one on Rails 4 which I kind of abandoned, because it was useful at the start for me but not anymore. The current one I see potential in (Rails 6), a couple weeks in the future I see launching 1-st or 2nd iteration. This could be a good or bad thing - I know what I want to accomplish, and keep improving it, because the thing I don't like is publishing something I would critique for edge cases bugs or overall user experience. But enough excuses. I am getting into the business, so I would be fine with minimum salary/hourly in your country, working remote. Is it hard in this climate, and I am asking thinking before Corona outbreak. I got my Git reps for those projects on bitbucket, so I can arrange for potential employers to have a peak; I do not know what is the procedure. I am not big on CVs, just like the idea of showing the way I code and if there are any takers, start working remotely for minimal compensation. My written English, my 2nd language, is quite good, not sure about spoken, haven't had opportunities to practice it much.

&amp;#x200B;

Any general advice guys on how to get into business having some code to show off? With tests and all :)
## [5][best command to kill all running `rails server` in macos](https://www.reddit.com/r/rails/comments/fnyi8f/best_command_to_kill_all_running_rails_server_in/)
- url: https://www.reddit.com/r/rails/comments/fnyi8f/best_command_to_kill_all_running_rails_server_in/
---

```
kill -9 `cat tmp/pids/server.pid`
issues
- only works if you're cd'd to the application directory
- stops only that instance

killall ruby || killall rails
issues
- works on some types of applications only

killall puma
issues
- im not sure this has ever worked for me. the only reason i try this is because _ps aux_ shows several puma entries

```

if you guys have a good ~/bin/kill-all-rails snippet, please share.
## [6][How many comments in the latest 24 hours posted by the (current)User?](https://www.reddit.com/r/rails/comments/fnqgzm/how_many_comments_in_the_latest_24_hours_posted/)
- url: https://www.reddit.com/r/rails/comments/fnqgzm/how_many_comments_in_the_latest_24_hours_posted/
---
In the `models/user.rb` I have

      has_many :comments

and in `models/comment.rb`, obviously, I have

    #  id         :integer          not null, primary key
    #  user_id    :integer          not null
    #  text       :text             not null
    #  created_at :datetime         not null
    #  updated_at :datetime         not null

&amp;#x200B;

**How to check how many comments the User posted in the latest 24 horus?**

Shoud I add it in `models/user.rb`, right?

I was thinking to 

      def comments_latest_24
        Comments.where('created_at &gt;= ?', 24.hours.ago)
      end

and in view to use `&lt;%= @comments_latest_24.count %&gt;`

is it right?

**Is it counting only the comments posted by THE User? Or the comments posted by everyone?!**
## [7][undefined method authenticate_user! after Rails 6 upgrade](https://www.reddit.com/r/rails/comments/fnzwwp/undefined_method_authenticate_user_after_rails_6/)
- url: https://www.reddit.com/r/rails/comments/fnzwwp/undefined_method_authenticate_user_after_rails_6/
---
Hi I did a Rails 5.2 to Rails 6 upgrade but then the authenticate\_user! from devise doesn't seem to be recognized. Anyone encounter a similar problem?
## [8][optimizing DB query based on whether last_updated_at field at has changed in last x hours](https://www.reddit.com/r/rails/comments/fnqjdw/optimizing_db_query_based_on_whether_last_updated/)
- url: https://www.reddit.com/r/rails/comments/fnqjdw/optimizing_db_query_based_on_whether_last_updated/
---
I am relatively new to understanding indexing and not sure if this is an appropriate use case.  I have a model of about 1000 records, and it is associated with another model of about lets say 3000 records, and another with 10s of thousands of records.  I plan to create a scheduled job that runs every few hours or perhaps 24 hours based on whether or not certain fields have changed and if so, if they need to be updated/saved with changes grabbed from a site. 

I'm not sure at which level I yet want to to check, but I will want to check one of the above models to see whether `last_updated_at` has changed say with the last hour or few, or perhaps 24 hours - still figuring out the time frame. 

Is an index just not a use case for the above example(I hope it makes some sense)?  Also available to the models is logidze so we could check relative changes, if I understand correctly.
## [9][API design using fast jsonapi](https://www.reddit.com/r/rails/comments/fnlwqe/api_design_using_fast_jsonapi/)
- url: https://www.reddit.com/r/rails/comments/fnlwqe/api_design_using_fast_jsonapi/
---
Hey, I'm new to Rails here! And no, this isn't a query on the framework itself

I'm designing an API for a smart home app. This is a sample response 

    {
      "data": [
        {
          "id": "1",
          "type": "Gateway",
          "attributes": {
            "name": "Soho Home",
            "serial_number": "1234567890",
            "integrator_id": 1,
            "rooms": {
              "data": [
                {
                  "id": "1",
                  "type": "Room",
                  "attributes": {
                    "name": "Foyer",
                    "devices": {
                      "data": [
                        {
                          "id": "1",
                          "type": "device",
                          "attributes": {
                            "name": "Corner Lamp",
                            "device_type": "light",
                            "percentage": "1.0"
                          },
                          "relationships": {
                            "room": {
                              "data": {
                                "id": "1",
                                "type": "room"
                              }
                            }
                          }
                        },
                        {
                          "id": "2",
                          "type": "device",
                          "attributes": {
                            "name": "Curtain",
                            "device_type": "curtain",
                            "percentage": "1.0"
                          },
                          "relationships": {
                            "room": {
                              "data": {
                                "id": "1",
                                "type": "room"
                              }
                            }
                          }
                        }
                      ]
                    }
                  },
                  "relationships": {
                    "devices": {
                      "data": [
                        {
                          "id": "1",
                          "type": "device"
                        },
                        {
                          "id": "2",
                          "type": "device"
                        }
                      ]
                    },
                    "gateway": {
                      "data": {
                        "id": "1",
                        "type": "gateway"
                      }
                    }
                  }
                },
                {
                  "id": "2",
                  "type": "Room",
                  "attributes": {
                    "name": "Cabin",
                    "devices": {
                      "data": []
                    }
                  },
                  "relationships": {
                    "devices": {
                      "data": []
                    },
                    "gateway": {
                      "data": {
                        "id": "1",
                        "type": "gateway"
                      }
                    }
                  }
                }
              ]
            }
          },
          "relationships": {
            "rooms": {
              "data": [
                {
                  "id": "1",
                  "type": "room"
                },
                {
                  "id": "2",
                  "type": "room"
                }
              ]
            },
          }
        }
      ]
    }

Now, I'm using the fast jsonapi gem provided by Netflix, and my question is whether I load all that data down to three-four levels in one call or do I utilise the link option provided in the gem to make continous calls to get data at each level

&amp;#x200B;

For example, lets assume I have four homes in my account. Each home has a gateway. Each gateway has rooms. Each room has devices. And I toggle those devices.

The API will be consumed by an Android/iOS app. Does it make sense to load everything on startup in one call (all the gateways and their nested attributes) and operate the device, or utilise the JSON API spec and make a subsequent calls to get the rooms, then another call for the devices, and then operate them?

I've not written an API in a while and this is my first time writing one in Rails. I absolutely love Rails so far but this bit of decision making is beyond my current capacity and I would be happy if someone chimed in!
## [10][Connecting users - Friendship](https://www.reddit.com/r/rails/comments/fnp4lg/connecting_users_friendship/)
- url: https://www.reddit.com/r/rails/comments/fnp4lg/connecting_users_friendship/
---
Hey Guys,   


I'm currently working on a simple enough Social Media platform for a college project.  
Am looking for a good(recent) guide to adding friendships between users. I have followed a guide or two already but seem to be hitting a dead end every time.   


The guilds I've tried (with no luck):  
 [https://smartfunnycool.com/friendships-in-activerecord/](https://smartfunnycool.com/friendships-in-activerecord/)   
 [http://railscasts.com/episodes/163-self-referential-association?view=asciicast](http://railscasts.com/episodes/163-self-referential-association?view=asciicast) 

Can anyone advise on a good guide or tutorial?  


Cheers
## [11][How to change table in db item?](https://www.reddit.com/r/rails/comments/fnq3jr/how_to_change_table_in_db_item/)
- url: https://www.reddit.com/r/rails/comments/fnq3jr/how_to_change_table_in_db_item/
---
One of my tables looks like this...

      create_table "artists", force: :cascade do |t|
        t.string "name"
        t.integer "user_id", null: false
        t.datetime "created_at", precision: 6, null: false
        t.datetime "updated_at", precision: 6, null: false
        t.index ["user_id"], name: "index_artists_on_user_id"
      end

But I need to delete the null: false from the user\_id to look like this...

      create_table "artists", force: :cascade do |t|
        t.string "name"
        t.integer "user_id"
        t.datetime "created_at", null: false
        t.datetime "updated_at", null: false
        t.index ["user_id"], name: "index_artists_on_user_id"
      end

How do I go about that? 

Thanks!
