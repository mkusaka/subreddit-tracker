# ruby
## [1][Ruby on Rails in a week](https://www.reddit.com/r/ruby/comments/jkpvwf/ruby_on_rails_in_a_week/)
- url: https://www.simplethread.com/ruby-on-rails-in-a-week/
---

## [2][Future of Ruby - The Journey to One Million Websocket Connections](https://www.reddit.com/r/ruby/comments/jki55j/future_of_ruby_the_journey_to_one_million/)
- url: https://www.youtube.com/watch?v=Dtn9Uudw4Mo
---

## [3][Packwerk at Shopify](https://www.reddit.com/r/ruby/comments/jkwp16/packwerk_at_shopify/)
- url: https://youtube.com/watch?v=olEA157z7kU
---

## [4][how is this converted to a hash in a function call?](https://www.reddit.com/r/ruby/comments/jktr6s/how_is_this_converted_to_a_hash_in_a_function_call/)
- url: https://www.reddit.com/r/ruby/comments/jktr6s/how_is_this_converted_to_a_hash_in_a_function_call/
---
Suppose I have a method like this 

    def fun(a, options = {})
      puts [a, options].inspect
    end
    fun(1, a: 3, b: 4) =&gt;  [1, {:a=&gt;3, :b=&gt;4}]

I don't understand how this gets converted to a hash. Also, is it possible to have this in the middle somewhere? like this  


    def fun(a, options = {}, c)
      puts [a, options].inspect
    end
    fun(1, a: 3, b: 4, 10)

without have to wrap \`a: 3, b: 4\` into  {}?
## [5][Why I love ruby, a complete RPC system using redis in 57 lines of code :)](https://www.reddit.com/r/ruby/comments/jki6b7/why_i_love_ruby_a_complete_rpc_system_using_redis/)
- url: https://www.reddit.com/r/ruby/comments/jki6b7/why_i_love_ruby_a_complete_rpc_system_using_redis/
---
So this RPC system has some advantages...
1. I programmed it as a toy (which was fun for me)
2. It's simple, if there is a bug in 57 lines of code, it's easy to find and fix
3. It uses redis, which I like, is simple, and fast
4. It handles exceptions which is probably necessary if anyone ever wanted to actually use it
5. there are no timeouts for long running code, cause sometimes timeouts are frustrating
6. there are timeouts for return values because those should be fast anyway
7. usage is simple so you don't have to understand what's going on to use it
8. if it's not already obvious rrecv and rsend invocations can occur on different processes (or with a little more configuration on different machines)  (all classes used as arguments should be available to all processes since we're using Marshaling)

```
rrecv(:name_of_thing_that_can_be_called) do |args|
#do whatever processing you want
#whatever is returned by this block is returned to the caller
end
```
```
result = rsend(:name_of_thing_that_can_be_called, &lt;whatever args you want&gt;, &lt;can be many&gt;)
```
every time you define a new callable thing with rrecv, it will spin up a new thread to respond to those requests

future work, add the ability to control how many workers rrecv creates, and possibly allow other types of workers other than just threads.



```
require 'securerandom'
require 'redis'

class Message
  attr_reader :payload, :guid

  def initialize(*args)
    @guid = SecureRandom.uuid
    @payload = args
  end
end

class RPCExceptionContainer
  attr_reader :saved_exception

  def initialize(exception)
    @saved_exception = exception
  end
end

def rsend(key, *args)
  @r ||= Redis.new
  Message.new(*args).tap do |message|
    @r.rpush(key, Marshal.dump(message))
    Marshal.load(@r.blpop(message.guid)[1]).tap do |return_value|
      if return_value.class == RPCExceptionContainer
        raise return_value.saved_exception
      else
        return return_value
      end
    end
  end
end

def rrecv(key)
  @keys ||= Hash.new
  if(!@keys.has_key?(key))
    r = Redis.new
    Thread.new do
      loop do
        Marshal.load(r.blpop(key)[1]).tap do |message|
          begin
            return_value = yield(message.payload)
            r.rpush(message.guid, Marshal.dump(return_value))
            r.expire(message.guid, 30)
          rescue
            r.rpush(message.guid, Marshal.dump(RPCExceptionContainer.new($!)))
            r.expire(message.guid, 30)
          end
        end
      end
    end
    @keys[key] = true
  else
    raise "You Called rrecv twice on the same key, this is almost certainly not what you intended"
  end
end

rrecv(:hello) do |args|
  puts "Hello #{args.reduce(:+)}"
  raise "what what"
  "success"
end
x = rsend(:hello, "hi there", "nobody") rescue "bobby farquard"
puts x

rrecv(:there) do |args|
  puts "There #{args.reduce(:+)}"
  "success"
end
y = rsend(:there, "something else")
puts y


x = rsend(:hello, "hi there", "nobody") rescue "bobby farquard again"
puts x
```
## [6][[outdated?] Build Awesome Command-Line Applications in Ruby 2](https://www.reddit.com/r/ruby/comments/jkhsa5/outdated_build_awesome_commandline_applications/)
- url: https://www.reddit.com/r/ruby/comments/jkhsa5/outdated_build_awesome_commandline_applications/
---
I am interested in using Ruby. My typical interaction with my OS is through the CLI. I'm not particularly interested in using shell scripts, thus Ruby would be a suitable replacement. I have the book as mentioned in the topic, which calls for Ruby 2.0 or 2.1. An Amazon review stated the code and gems in this book is outdated. Might there be a better resource regarding Ruby use as a shell replacement? Or will studying this code from 2013 really do me in a bad way until I'm more knowledgeable. Thank you
## [7][Need help on Chess Game](https://www.reddit.com/r/ruby/comments/jkmvoz/need_help_on_chess_game/)
- url: https://www.reddit.com/r/ruby/comments/jkmvoz/need_help_on_chess_game/
---
Hey guys, I'm making a chess game and I'm having a lot of trouble. We are on the inheritance section of the course so each piece is a sub class of "Piece".

In the directions, it says:  Your Piece will need to (1) track its position and (2) hold a reference to the Board. Don't allow a piece to move into a square already occupied by the same color piece, or to move a sliding piece past a piece that blocks it.

My TA said that I am supposed to put board in the initialize method for piece, and pass the board when creating each piece

I created all of the pieces in a board method where I iterated through each of the board spots in the first two and last two rows. However, I cant pass a fully formed board until the last piece has been created! So how do I do this?

Also, the board will change as pieces change spots, how do I update each piece's board attribute to be accurate???

Thank you!
## [8][Right-hand assignment has been converted into a one-line pattern matching](https://www.reddit.com/r/ruby/comments/jk8icl/righthand_assignment_has_been_converted_into_a/)
- url: https://github.com/ruby/ruby/pull/3703
---

## [9][Backing Rails enums with Postgres ENUM type](https://www.reddit.com/r/ruby/comments/jk77cb/backing_rails_enums_with_postgres_enum_type/)
- url: https://www.reddit.com/r/ruby/comments/jk77cb/backing_rails_enums_with_postgres_enum_type/
---
[https://diffickens.com/rails/2020/10/02/back-your-enums-with-enums.html](https://diffickens.com/rails/2020/10/02/back-your-enums-with-enums.html)

I haven't blogged since ages, but there was a discussion regarding how to store Rails enums, so I decided to give writing another try. :-)
## [10][Designed a simple rails app that communicates with Google Drive API.](https://www.reddit.com/r/ruby/comments/jkavpd/designed_a_simple_rails_app_that_communicates/)
- url: https://github.com/AbhishekSinhaCoder/Ruby-on-Rails-Challenge
---

