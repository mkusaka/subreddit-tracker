# ruby
## [1][started to learn ruby and run into a bit of a problem](https://www.reddit.com/r/ruby/comments/f9qr4i/started_to_learn_ruby_and_run_into_a_bit_of_a/)
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
## [2][Accessing data between several Rails apps](https://www.reddit.com/r/ruby/comments/f9qxen/accessing_data_between_several_rails_apps/)
- url: https://www.reddit.com/r/ruby/comments/f9qxen/accessing_data_between_several_rails_apps/
---
The company I work for has the domain split between a few monoliths, each is its own Rails app with its own daabase. All of the apps depend on one monolith that handles users, roles and other things, let's call it UsersApp.

Now let's say you have BillingApp, and it needs to pull the data from UsersApp. For example forms in BullingApp need to have the user's data to show it in the form drop downs etc., what ways do you know of handling this issue and what are the tradeoffs? In our case the relationship is read only (BillingApp always reads and never writes), but I'd like to get a more general discussion going.

Just to help get a discussion going here are the ways I know of:

1. Using ActiveResource in BillingApp to access everything, or similar technique to just pull everything through an API that UsersApp exposes.
2. Pulling and syncing the data from UsersApp into BillingApp (basically duplicating the database tables in BillingApp). Let's say the data isn't time sensitive, the sync can happen once a day so "real timeness" isn't an issue here.
3. Packing UsersApp as a gem, and having BillingApp access the models/data directly.
## [3][Ruby one of the highest-paid programming languages globally in 2020](https://www.reddit.com/r/ruby/comments/f9aumo/ruby_one_of_the_highestpaid_programming_languages/)
- url: https://learnworthy.net/highest-paid-programming-languages-in-2020/
---

## [4][.destroy Getting Killed](https://www.reddit.com/r/ruby/comments/f9ll24/destroy_getting_killed/)
- url: https://www.reddit.com/r/ruby/comments/f9ll24/destroy_getting_killed/
---
I'm running into an issue where I have a model with many children (has\_many) relationships. 

When I call .destroy on this model, the process grows so large that it's getting killed.

I've tried overloading .destroy on that model and some of the associated model to use a `.each(&amp;:destroy)` where possible but that did not solve the issue.

Any suggestions on destroying a model where the associated models have 100k-1M rows?
## [5][Ruby Conferences &amp; Camps in 2020 - What's Upcoming? Anything Missing? Updates Welcome](https://www.reddit.com/r/ruby/comments/f9fohh/ruby_conferences_camps_in_2020_whats_upcoming/)
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
## [6][Logging Accessible via Active Record](https://www.reddit.com/r/ruby/comments/f9lizu/logging_accessible_via_active_record/)
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
## [7][Rails 6 fixes a bug where after_commit callbacks are called on failed update in a transaction block | The Official BigBinary Blog | BigBinary](https://www.reddit.com/r/ruby/comments/f997tk/rails_6_fixes_a_bug_where_after_commit_callbacks/)
- url: https://blog.bigbinary.com/2020/02/25/rails-6-fixes-a-bug-where-after_commit-callbacks-are-called-on-failed-update-in-a-transaction-block.html
---

## [8][Deep dive into rackup](https://www.reddit.com/r/ruby/comments/f978se/deep_dive_into_rackup/)
- url: https://deepdive.sh/deep-dive-into-rackup
---

## [9][Ruby on Rails Mistakes could kill Your Production Servers](https://www.reddit.com/r/ruby/comments/f96x3n/ruby_on_rails_mistakes_could_kill_your_production/)
- url: https://pawelurbanek.com/rails-mistakes-downtime
---

## [10][Hate for Medium](https://www.reddit.com/r/ruby/comments/f927ed/hate_for_medium/)
- url: https://www.reddit.com/r/ruby/comments/f927ed/hate_for_medium/
---
Several times, I've considered biting the bullet, putting my hand in my pocket, ponying up the $5, or whatever it is, for a Medium subscription. Some of the content is pretty good. I've enjoyed content by /u/mehdifarsi, /u/eric_programmer, /u/vfreefly and others in the past.

But, today, the top article in my *Medium Daily Digest* points me to \***medium link alert**\* - [Tracking User Streaks in Ruby](https://levelup.gitconnected.com/tracking-user-streaks-in-ruby-a49e90ce46a1). If you don't want to follow that link, fair enough, but in it, the author presents a 14 line poorly formatted method containing errors (both syntactic and semantic) that purports to count user activity streaks.

Now, I don't want to hate on the author - if this post was on a personal blog or a free aggregation site, I'd be happy to comment with some positive constructive criticism - but for a service that pesters me regularly for $5 per month to deliver this kind of "Ideas and perspectives you won’t find anywhere else" from "the world’s most insightful writers, thinkers, and storytellers" and call it "the smartest takes on topics that matter" is an outright scam!

If this outfit is taking $5/month off what must be an awful lot of people, surely they can afford to pay some subject matter experts to do the most cursory of proof reading? Can't they?

Anyway, rant over. Happy Birthday Ruby 🎂 🎉
