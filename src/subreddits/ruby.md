# ruby
## [1][I created a gem for creating auto-similar fractals](https://www.reddit.com/r/ruby/comments/g0grqs/i_created_a_gem_for_creating_autosimilar_fractals/)
- url: https://github.com/StefanoMartin/FractalsRB
---

## [2][Improve code readability with closures in Ruby](https://www.reddit.com/r/ruby/comments/g0flu7/improve_code_readability_with_closures_in_ruby/)
- url: https://railsguides.net/improve-code-readability-with-closures-in-ruby/
---

## [3][DudePolicy - gem for policy objects in Ruby on Rails application from perspective of current_user/current_account](https://www.reddit.com/r/ruby/comments/g0g3rj/dudepolicy_gem_for_policy_objects_in_ruby_on/)
- url: https://github.com/equivalent/dude_policy
---

## [4][Is Ruby worth Learning in 2020?](https://www.reddit.com/r/ruby/comments/g0efr7/is_ruby_worth_learning_in_2020/)
- url: https://www.reddit.com/r/ruby/comments/g0efr7/is_ruby_worth_learning_in_2020/
---
I'm a young programmer and have learnt c++ in school, so I'm familiar with OOP. I wanted to try out a new language(as I never found c++ interesting) and so I selected Ruby because of its simplicity. I watched few tutorials on youtube a tried to implement those on my own and I have to say that it's way easier than c++ and gives the same output as c++ with a lot less of code!!! So I was really impressed and thought to learn it in depth so I went to [Udemy.com](https://Udemy.com) and purchased a course of 48 hours which had a really good user review. But now some of my friends are suggesting me to learn python and said that learning ruby will be a waste of time n 2020 and instead of it I should learn Python. As being a student from India(from a middle-class family) I'm really scared whether my money has got wasted. Please share your opinion on what should I do as Udemy has a feature of refund within 30 days. Your help will be really appreciated!
## [5][Grape: Include is still better for code splitting](https://www.reddit.com/r/ruby/comments/g0e5lm/grape_include_is_still_better_for_code_splitting/)
- url: https://www.reddit.com/r/ruby/comments/g0e5lm/grape_include_is_still_better_for_code_splitting/
---
Grape offers a [mount method](https://github.com/ruby-grape/grape#modules) for combining parts (endpoints) of API.

    class AddressAPI &lt; Grape::API
      get '/addresses' do
        # logic here
      end
    
      params do
        # params declaration
      end
      post '/address' do
        # logic here
      end
    
      params do
        # params declaration
      end
      patch '/address' do
        # logic here
      end
    end
    
    class API &lt; Grape::API
      version 'v1', using: :path
    
      mount AddressAPI
    end

Naturally, we would like to avoid one long file (class), so some developers could use this way for code splitting.

However, there is another one way which every Ruby developer knows.

    module AddressAPI
      extend ActiveSupport::Concern
    
      included do
        get '/addresses' do
          # logic here
        end
    
        params do
          # params declaration
        end
        post '/address' do
          # logic here
        end
    
        params do
          # params declaration
        end
        patch '/address' do
          # logic here
        end
      end
    end
    
    class API &lt; Grape::API
      version 'v1', using: :path
    
      include AddressAPI
    end

**Note:** I could implement it without ActiveSupport, but Grape depends on it, so I prefer to benefit from dependencies I have in the stack.

A project I work on uses the last approach, I don't know the reason, most likely, the mount wasn't supported back then. So, I was curios which approach is better in terms of memory.

I added params to the post and path endpoints:

    requires :address, type: Hash do
      requires :street, type: String
      requires :postal_code, type: Integer
      optional :city, type: String
      optional :tags, type: Array[String]
    end

then installed [memory\_profiler](https://rubygems.org/gems/memory_profiler) to Ruby 2.7.1 and measured the API declaration.

**Mount**:

    Total allocated: 549351 bytes (4881 objects)
    Total retained:  146918 bytes (1020 objects)

**Include:**

    Total allocated: 419311 bytes (3938 objects)
    Total retained:  79702 bytes (583 objects)

The difference is essential, the mount needs extract *126.99 Kb* of the memory, *65.6 Kb* of *126.99 Kb* wasn't removed by garbage collector, very likely, it will stay.

A reason for this difference is [inheritance](https://github.com/ruby-grape/grape/blob/002280415a46b1cabea565533141298781a48553/lib/grape/dsl/routing.rb#L95), `AddressAPI` inherits settings of `API`. The inheritance logic is [quite complex](https://github.com/ruby-grape/grape/blob/9feefdff949078a266784e0d2d6f74213f7cf729/lib/grape/util/inheritable_setting.rb#L50-L61), it even involves copies in time. Settings contain lots of things:

* declared params
* validations
* helpers
* callbacks (`before`, `after` etc)
* rescue handlers

Ok, let's make it even more fun and extract the post endpoint to a separate class and mount into `AddressAPI`:

    class PostAddressAPI &lt; Grape::API
      params do
        requires :address, type: Hash do
          requires :street, type: String
          requires :postal_code, type: Integer
          optional :city, type: String
          optional :tags, type: Array[String]
        end
      end
      post '/address' do
        # nothing here
      end
    end

**Outcome:**

    Total allocated: 656751 bytes (5654 objects)
    Total retained:  203502 bytes (1398 objects)

To be fair, I checked the same approach with the include:

    Total allocated: 420367 bytes (3942 objects)
    Total retained:  80758 bytes (587 objects)

The difference is *230.84 Kb*! If you use the mount for code splitting, probably, you need to consider the include.
## [6][Feature Flags: The stupid simple way to de-stress production releases](https://www.reddit.com/r/ruby/comments/g0ivt2/feature_flags_the_stupid_simple_way_to_destress/)
- url: https://boringrails.com/articles/feature-flags-simplest-thing-that-could-work/
---

## [7][Help Needed: Where to Put Ruby Docs?](https://www.reddit.com/r/ruby/comments/g01xyg/help_needed_where_to_put_ruby_docs/)
- url: https://www.reddit.com/r/ruby/comments/g01xyg/help_needed_where_to_put_ruby_docs/
---
Where, other than in the Ruby project itself, can enhanced Ruby documentation be shared to a broad readership?

(By enhanced, I mean substantially beyond what's in the Ruby doc itself.)

TL;DR

Quite a while back I had the notion to found a GitHub project where I'd write about Ruby.  My goals included:

\- Telling examples for everything.

\- Hard-nosed precision (e.g., Hash-convertible objects).

\- Completeness (e.g., all exceptions raised).

Shortly, though, I realized that some of this would represent enhancements to Ruby's own documentation.  So, I asked myself, shouldn't it actually be put into the Ruby documentation itself, instead of in a separate (and far less prominent) place?  The Ruby doc lives in comments in the Ruby project's code, so that's where the enhancements would go.

Last year I experimented by basically rewriting the documentation for Ruby's ENV object.  Happy to say that all that got merged (in small batches) and then released with version 2.7.  Here are the before and after versions:

\* \[ENV version 2.6.6\]([https://ruby-doc.org/core-2.6.6/ENV.html](https://ruby-doc.org/core-2.6.6/ENV.html))

\* \[ENV version 2.7.0\]([https://ruby-doc.org/core-2.7.0/ENV.html](https://ruby-doc.org/core-2.7.0/ENV.html))

All well and good.

This year I've been working on classes Hash and Array.

I have not, though, been able to get anything (again, in small batches) reviewed and merged.  I don't complain about that:  Everyone's busy, just like us.

The problem is, I have a huge backlog of documentation now:

\* All 75 methods in \[Hash\]([https://github.com/BurdetteLamar/AboutRuby/blob/master/core/Hash/api/markdown.md#hash-api](https://github.com/BurdetteLamar/AboutRuby/blob/master/core/Hash/api/markdown.md#hash-api)).

\* Two-thirds of the 95 methods in \[Array\]([https://github.com/BurdetteLamar/AboutRuby/blob/master/core/Array/api/markdown.md#array-api](https://github.com/BurdetteLamar/AboutRuby/blob/master/core/Array/api/markdown.md#array-api)).

I've stashed all that in my own project, About Ruby.  But I don't think that very many in the huge Ruby world will find it there.

So my question is this:  Is there someplace else I can put this where people are more likely to find it?
## [8][Kind - Basic type system (in runtime) for Ruby](https://www.reddit.com/r/ruby/comments/g064hg/kind_basic_type_system_in_runtime_for_ruby/)
- url: https://www.reddit.com/r/ruby/comments/g064hg/kind_basic_type_system_in_runtime_for_ruby/
---
[https://github.com/serradura/kind](https://github.com/serradura/kind)  


I would like to share the Kind gem with you. I made it because I would like an abstraction to check the data of my method arguments. e.g:

    def sum(a, b)
      Kind.of.Numeric(a) + Kind.of.Numeric(b)
    end
    
    sum(1, 1)   # 2
    sum('1', 1) # Kind::Error ("\"1\" expected to be a kind of Numeric")

Until now, the project covers [23 kinds](https://github.com/serradura/kind#built-in-type-checkers) (the most relevant from Ruby's core).

It allows the creation of custom types. e.g:

    module Account
      class User
        Kind::Types.add(self)
      end
    end
    
    Kind.of.Account::User(Account::User.new)  # #&lt;Account::User:0x0000...&gt;

And one last thing, the gem exposes a constant to represent a value equivalent to `nil` ([Kind::Undefined](https://github.com/serradura/kind#kindundefined)). And to handle it and nil values, the gem has the [Kind::Maybe](https://github.com/serradura/kind#kindmaybe) to help the creation of computations which will avoid null pointer exceptions. e.g:

    result1 =
      Kind::Maybe[5]
        .then { |value| value * 5 }
        .then { |value| value + 17 }
        .value_or(0)
    
    puts result # 42
    
    #---
    
    result2 =
      Kind::Maybe[5]
        .then { nil }
        .then { |value| value * 5 }
        .value_or(0)
    
    puts result # 0

So, if you be interested. Please check out the project and send me feedback about it. Thanks!  


PS: I'm a big fan of \`dry-types\` and I mention this great lib in the [project's README](https://github.com/serradura/kind#kind-) (before the table of contents).
## [9][Seeking Rails Developers to help with an open source project.](https://www.reddit.com/r/ruby/comments/g07jdt/seeking_rails_developers_to_help_with_an_open/)
- url: https://www.reddit.com/r/ruby/comments/g07jdt/seeking_rails_developers_to_help_with_an_open/
---
Hey there!

**About me:** I am a technical wizard with nearly 20 years development experience, everything from Perl to Java to Rails. I love Rails because it's so easy to prototype something completely from zero.

I have a project that I released as open source, and I'd be interested in gathering collaborators. It helps people who can't attend anime tradeshows (due to COVID-19) make a virtual store. We don't charge people nor vendors to use it - it's currently just a "Buy Now" button on Paypal.

I'd love support implementing other features from a good Rails developer - plus it's a great way to show public contributions on your Github profile (if you care about that sort of thing... some recruiters check your profile).

Project url: [https://github.com/ryankopf/cononline](https://github.com/ryankopf/cononline)

I've got a small budget of $500 for the right dev(s) to help a bit.

Anyone interested?

Ryan
## [10][Help with recursion (backtracking) in Ruby](https://www.reddit.com/r/ruby/comments/g017te/help_with_recursion_backtracking_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/g017te/help_with_recursion_backtracking_in_ruby/
---
&gt;!I'm trying to make a sudoku solver in Ruby, and it doesn't go back to discard a wrong solution and keep trying new ones.!&lt;

&gt;!https://gist.github.com/bulgakke/fdb25526f2e514fe3f33ba1279e32355 - This is the same principle stripped of all the sudoku stuff; the goal is to form an array in which every number must be less than the previous one, recursively.!&lt;

&gt;!What is wrong with it?!&lt;

&gt;!https://gist.github.com/bulgakke/c521b1c92e2144ebfcf24ecd2352c6f8 - here is the actual sudoku solver code piece compared to one I found online. Where is the fundamental difference here (the two non-influential differences I spotted are: the second one works with string instead of two-dimensional arrays and checks the validity after trying the number, not before)?!&lt;

UPD: Solved. Thanks to u/fedekun, I changed \`@zero\_index\` to be \`zero\_index\`

Wasn't fully aware what are instance variables; apparently, there is only one instance (duh) of it throughout all the calls, now I see

&amp;#x200B;

UPD2: Done!  [https://github.com/bulgakke/sudoku\_solvers](https://github.com/bulgakke/sudoku_solvers) 
