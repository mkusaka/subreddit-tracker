# ruby
## [1][Strict interfaces in Ruby](https://www.reddit.com/r/ruby/comments/jackyk/strict_interfaces_in_ruby/)
- url: https://github.com/featurist/interfacable
---

## [2][Question about Ruby class instance variable](https://www.reddit.com/r/ruby/comments/ja6ds8/question_about_ruby_class_instance_variable/)
- url: https://www.reddit.com/r/ruby/comments/ja6ds8/question_about_ruby_class_instance_variable/
---
&amp;#x200B;

[Hey guys so the instructions said \\"Start by creating an instance variable, self.root.node that stores the knight's initial position in an instance of your class. I don't understand this. Wouldn't self.root.node be a class METHOD??? wouldn't it be an instance variable if it was like @ root = PolyTreeNode.new\(start\_pos\). See the part I highlighted for what I'm talking about. Also, I initialized my @ root = PolyTreeNode.new\(start\_pos\) and didn't put it in this build\_move\_tree method. Not sure if that matters but if it does I would love to know why. Thanks.](https://preview.redd.it/4olkecczess51.png?width=746&amp;format=png&amp;auto=webp&amp;s=0b936032a6afde377baac6ee0154b46e9301933c)
## [3][Software transactional memory proposal for Threads and Ractors](https://www.reddit.com/r/ruby/comments/j9y0dj/software_transactional_memory_proposal_for/)
- url: https://www.reddit.com/r/ruby/comments/j9y0dj/software_transactional_memory_proposal_for/
---
[https://bugs.ruby-lang.org/issues/17261](https://bugs.ruby-lang.org/issues/17261)
## [4][The Ruby world beyond Ruby on Rails - web frameworks](https://www.reddit.com/r/ruby/comments/j9ojp0/the_ruby_world_beyond_ruby_on_rails_web_frameworks/)
- url: https://longliveruby.com/articles/rails-alternatives-ruby-web-frameworks
---

## [5][Action Mailer queueing, but not sending](https://www.reddit.com/r/ruby/comments/j9uhjk/action_mailer_queueing_but_not_sending/)
- url: https://www.reddit.com/r/ruby/comments/j9uhjk/action_mailer_queueing_but_not_sending/
---
Hello! I have no idea what I'm doing.

A ruby project was handed over to me but I have 0 experience on it and the last guy left no documentation. I've been able to work my way around some things, but there is currently an issue with the mailer. It won't send anything, but it leaves no error and it even logs that it's been enqueued.

https://preview.redd.it/atcfdq9zzos51.png?width=1278&amp;format=png&amp;auto=webp&amp;s=25ea148a0a58a725aa0da7442daf073202944ed6

For what I've researched, the problem is that Sidekiq is not active (and it isn't), but even when I run Sidekiq manually, it'll still never send anything. 

This is the ucrrent mailer configuration. The gmail and it's password are both correct. The system once used to work because the gmail account has literal thousands of sent emails, but it suddenly stopped working a few weeks ago.

https://preview.redd.it/mnz4yloh0ps51.png?width=556&amp;format=png&amp;auto=webp&amp;s=bf8276d53b21f90a70b97415bb1fa3388385a07a

The server is running on AWS, and again, I have no clue of what i'm doing or what's going on, so any help is going to be really appreciated, and a lot of patience too!
## [6][Ruby + Hacktoberfest Meetup! Tonight! JWTs](https://www.reddit.com/r/ruby/comments/j9q4hx/ruby_hacktoberfest_meetup_tonight_jwts/)
- url: https://www.reddit.com/r/ruby/comments/j9q4hx/ruby_hacktoberfest_meetup_tonight_jwts/
---
🎟[https://live.remo.co/e/israelrb-hacktoberfest/register](https://live.remo.co/e/israelrb-hacktoberfest/register)

Join us for the 3rd israel.rb virtual meetup!

This time it'll be a Hacktoberfest edition with special guest Dan Moore from FusionAuth discussing JSON Web Tokens (JWTs) and what Rails developers need to know about them.

We'll also have lightning talks from the community! If you're interested in giving a 5-10 minute lightning talk on an open source theme, let us know!  


https://preview.redd.it/avcky8gksns51.jpg?width=820&amp;format=pjpg&amp;auto=webp&amp;s=d21cf36839efd639be7c3b527175b6aba0c245f6

⏰ 19:00 Israel / 17:00 London / 12:00 NYC / 09:00 SF
## [7][Passing arrays into object intialize and unit testing?](https://www.reddit.com/r/ruby/comments/j9ycvr/passing_arrays_into_object_intialize_and_unit/)
- url: https://www.reddit.com/r/ruby/comments/j9ycvr/passing_arrays_into_object_intialize_and_unit/
---
I have been able to unit test my car class but I am having a lot of trouble with cardealer:

    class Car
      def initialize(name, rank, number)
        @name = name
        @rank = rank
        @number = number
      end
      attr_reader :name
      attr_reader :rank
      attr_reader :number
    end

&amp;#x200B;

    require_relative "car"
    
    class CarDealer
      def initialize(cars)
        @cars[] = cars
      end
    
      attr_reader :cars
    end
    

&amp;#x200B;

    require "test/unit"
    require_relative "car"
    require_relative "cardealer"
    
    class CarDealerTest &lt; Test::Unit::TestCase
      def setup
        @cars = []
        IO.foreach("cars.txt") do |line|
          data = line.split
          @cars.push Car.new(data[0], data[1].to_i, data[2].to_i)
        end
    
        @car_dealer = CarDealer.new(@cars)
      end
    
      def test_car_dealer_cars
        assert_equal(@cars, @car_dealer.cars)
      end
    
      def test_car_dealer_car_1_name
        assert_equal(@cars[0].name, @car_dealer.cars[0].name)
      end
    
    end
    

&amp;#x200B;

When I try to even just print the array objects the output is nil.

I changed this to just printing when the following which was there instead was giving the errors:

    ##teamcity[testFailed name = 'test_car_dealer_car_1_name' message = 'NoMethodError: undefined method `|[|]=|' for nil:NilClass' details = '/home/user/cloud/code/RubymineProjects/cardealer.rb:5:in `initialize|'|n    /home/user/cloud/code/RubymineProjects/cardealertest.rb:13:in `new|'|n    /home/user/cloud/code/RubymineProjects/cardealertest.rb:13:in `setup|'|n    (eval):12:in `run|'|n    /home/user/.local/share/JetBrains/Toolbox/apps/RubyMine/ch-0/202.7660.39/plugins/ruby/rb/testing/patch/testunit/test/unit/ui/teamcity/testrunner.rb:93:in `start_mediator|'|n    /home/user/.local/share/JetBrains/Toolbox/apps/RubyMine/ch-0/202.7660.39/plugins/ruby/rb/testing/patch/testunit/test/unit/ui/teamcity/testrunner.rb:81:in `start|'' error = 'true' nodeId = 'CarDealerTest.test_car_dealer_car_1_name' timestamp = '2020-10-12T20:57:12.195+0100']

&amp;#x200B;

      def test_car_dealer_cars
        assert_equal(@cars, @car_dealer.cars)
      end
    
      def test_car_dealer_car_0_name
        assert_equal(@cars[0].name, @car_dealer.cars[0].name)
      end

cars.text file:

    Mondeo 8 8
    Fiesta 6 7

Please tell me what I am doing wrong.
## [8][Ruby 2.7.2 has revised Rdoc for Array](https://www.reddit.com/r/ruby/comments/j9791u/ruby_272_has_revised_rdoc_for_array/)
- url: https://www.reddit.com/r/ruby/comments/j9791u/ruby_272_has_revised_rdoc_for_array/
---
I've  spent a good part of this year working on revisions to the RDoc for  Array.  The revisions include all methods documentation in addition to  much of the introductory text (class documentation).  Lots more example  code.

Check it out:

\- 2.7.2: [https://ruby-doc.org/core-2.7.2/Array.html](https://ruby-doc.org/core-2.7.2/Array.html)

\- 2.7.1: [https://ruby-doc.org/core-2.7.1/Array.html](https://ruby-doc.org/core-2.7.1/Array.html)
## [9][Ruby: The not so good parts](https://www.reddit.com/r/ruby/comments/j9c8z7/ruby_the_not_so_good_parts/)
- url: https://www.chrismytton.com/ruby-the-not-so-good-parts/
---

## [10][Integrate Bootstrap 4 on Rails 6 Application using Webpack and Yarn || No gem needed](https://www.reddit.com/r/ruby/comments/j9mwmm/integrate_bootstrap_4_on_rails_6_application/)
- url: https://www.youtube.com/watch?v=Zhw_GSoWQrc&amp;t=19s
---

