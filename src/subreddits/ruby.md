# ruby
## [1][Rails 6.1 deprecates rails db:structure:dump and rails db:structure:load | The Official BigBinary Blog | BigBinary](https://www.reddit.com/r/ruby/comments/ixjs28/rails_61_deprecates_rails_dbstructuredump_and/)
- url: https://blog.bigbinary.com/2020/09/22/rails-6-1-deprecates-rails-db-structure-dump.html
---

## [2][Ruby Courses - Advanced Ruby: Behind the Magic](https://www.reddit.com/r/ruby/comments/ix35gf/ruby_courses_advanced_ruby_behind_the_magic/)
- url: https://www.reddit.com/r/ruby/comments/ix35gf/ruby_courses_advanced_ruby_behind_the_magic/
---
It seems like biggest ruby resources are screencasts, blogs and books. And yes, they offer tons of value. 

But I’m wondering if you folks know of effective courses to learn advanced ruby/rails concepts?

Also, has anyone tried this one “Advanced Ruby: Behind the Magic”

https://courses.gorails.com/advanced-ruby-for-rails-devs

(To me, 300 bucks seems to be on the higher end for 30 screencasts..)
## [3][Any downsides storing database host/username/password in ruby ENV?](https://www.reddit.com/r/ruby/comments/ixazdi/any_downsides_storing_database/)
- url: https://www.reddit.com/r/ruby/comments/ixazdi/any_downsides_storing_database/
---
I have this code

YAML.load(File.open(env\_file))\[ENV\["RACK\_ENV"\]\].each { |key, value| ENV\[key.to\_s\] = value } if File.exists?(env\_file)

This file loads all the secret variables, client API secrets, db secrets and everything into ENV.

&amp;#x200B;

Then, different gems just call ENV\['aa\_secret\_key'\] to get the secret and boot up the app

secret\_env\_file.yml is not uploaded on source control. So, everything stays safe.

&amp;#x200B;

Are there any downsides of using ENV to store all those information while the app boots up? I think ENV\['anything\_password'\] is still accessible later on as long as the app is still up and running. Like for sinatra/rails etc.
## [4][Any good Ruby library to load test a Rails application?](https://www.reddit.com/r/ruby/comments/ix6ltq/any_good_ruby_library_to_load_test_a_rails/)
- url: https://www.reddit.com/r/ruby/comments/ix6ltq/any_good_ruby_library_to_load_test_a_rails/
---
Hello fellow ruby developers,

I'm currently maintaining [Gatling load test](https://gatling.io/) scenario for a quite decent Rails app (\~250k rpm).

It's working well but it is very hard to maintain and implement new features on the scenario because everything's written in Scala which my team and I are not quite familiar with.

The scenario we are writing are RPM-based meaning that we are trying our best to get as close as possible from production in terms of rpm for most endpoints.

Does anyone know if there is a Ruby implementation somewhere that we could try?

I heard that Ruby wasn't made for this kind of task and languages/frameworks such as Gatling, NodeJs, ... and they were way better at it. Could someone explain me why?

Cheers from France!
## [5][Issues with pattern matching on JSON / array of hashes](https://www.reddit.com/r/ruby/comments/ix7a1h/issues_with_pattern_matching_on_json_array_of/)
- url: https://www.reddit.com/r/ruby/comments/ix7a1h/issues_with_pattern_matching_on_json_array_of/
---
Is there a way to get pattern matching to work on an array of hashes? My use case is a JSON document ([here](https://www.dph.illinois.gov/sitefiles/COVIDHospitalRegions.json?nocache=1)) that has a few arrays of hashes. I don't know if this is a potential use for pattern matching, but I would *like* to use it to match against particular items, acting sort of like a `where` clause:

    require 'net/http'
    require 'json'

    HOSPITAL_DATA_URI = URI("https://www.dph.illinois.gov/sitefiles/COVIDHospitalRegions.json?nocache=1").freeze

    hospital_data = JSON.parse(Net::HTTP.get(HOSPITAL_DATA_URI), symbolize_names: true)
    region_values = hospital_data[:regionValues]

    case region_values
      in [{id: 10, **unimportant_attributes} =&gt; region_10]
        puts 'We found region 10!'
    else
      puts 'aw shucks'
    end

    binding.irb

Does anyone know if this is possible?

EDIT - wrote a StackOverflow question for it as well: https://stackoverflow.com/questions/63998870/how-can-i-pattern-match-on-an-array-of-hashes-in-ruby-2-7
## [6][What's a snippet of code that makes you happiest that you learnt Ruby?](https://www.reddit.com/r/ruby/comments/iwy3c4/whats_a_snippet_of_code_that_makes_you_happiest/)
- url: https://www.reddit.com/r/ruby/comments/iwy3c4/whats_a_snippet_of_code_that_makes_you_happiest/
---
For me currently it's:

    puts`dd`

8 character identity algorithm (outputs any input without any alteration). :) Keen to see other examples.
## [7][Secure way of doing OAuth for SPA &amp; Native Apps](https://www.reddit.com/r/ruby/comments/iwztni/secure_way_of_doing_oauth_for_spa_native_apps/)
- url: https://blog.joshsoftware.com/2020/05/18/secure-way-of-doing-oauth-for-spa-native-apps/
---

## [8][Compile Ruby to C](https://www.reddit.com/r/ruby/comments/iwr3pn/compile_ruby_to_c/)
- url: https://github.com/agrafix/rubyspeed
---

## [9][An SMS Reminder Service With Ruby &amp; A Raspberry Pi](https://www.reddit.com/r/ruby/comments/iwoh4f/an_sms_reminder_service_with_ruby_a_raspberry_pi/)
- url: https://emmanuelhayford.com/building-an-sms-reminder-with-ruby-raspberry-pi/
---

## [10][An RSpec custom matcher to test application code that logs information into log files](https://www.reddit.com/r/ruby/comments/iwl1fq/an_rspec_custom_matcher_to_test_application_code/)
- url: https://github.com/juanmanuelramallo/rspec-log_matcher
---

