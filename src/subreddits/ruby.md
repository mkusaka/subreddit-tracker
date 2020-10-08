# ruby
## [1][Polyphony has an io_uring backend!](https://www.reddit.com/r/ruby/comments/j79rmc/polyphony_has_an_io_uring_backend/)
- url: https://www.reddit.com/r/ruby/comments/j79rmc/polyphony_has_an_io_uring_backend/
---
I am pleased to announce the release of [Polyphony](https://github.com/digital-fabric/polyphony) version 0.46.0, which includes a full-blown io\_uring backend. The io\_uring backend uses a relatively recent new I/O API in Linux that is the future of I/O processing. The new backend provides more than double the performance of the libev backend, achieving **over 120K requests/second** using a basic "Hello world" HTTP server.

Those of you unfamiliar with io\_uring can head over to the [Lord of the io\_uring website](https://unixism.net/loti/index.html). The io\_uring backend is automatically used on Linux with kernel version 5.6 or higher.

For more information on the new io\_uring backend, including preliminary benchmark numbers and design details you can consult the [pull request](https://github.com/digital-fabric/polyphony/pull/44). The source code for the io\_uring backend in all its glory can be found[ here](https://github.com/digital-fabric/polyphony/blob/master/ext/polyphony/backend_io_uring.c).

For more information on Polyphony please consult the (slightly outdated) [Polyphony docs](https://digital-fabric.github.io/polyphony/).
## [2][Getting started with Tailwind CSS on Rails 6](https://www.reddit.com/r/ruby/comments/j78xd7/getting_started_with_tailwind_css_on_rails_6/)
- url: https://rubyyagi.com/tailwind-css-on-rails-6-intro/
---

## [3][What makes Ruby such a good language for creating DSLs ?](https://www.reddit.com/r/ruby/comments/j79fd2/what_makes_ruby_such_a_good_language_for_creating/)
- url: https://www.reddit.com/r/ruby/comments/j79fd2/what_makes_ruby_such_a_good_language_for_creating/
---
I have no formal education in computer science, I have bussiness background, but was always interested in programming since young age. Lately I am working through [https://teachyourselfcs.com/](https://teachyourselfcs.com/)

I am now working through classic SICP, apparently it later discusses creating language as another way of abstraction (complementing functional programming and OOP).

Now, I have an ecommerce company where I have created some tools in Python. Mainly doing some http requests, automating sending emails,  downloading xml files, processing them, filtering products from warehouses etc. So mostly command line tools. Probably it will be rewritten soon to a server program that continously works in the background.

I am looking for a language where I can experiment with ideas of: functional programming (doesn't have to be pure FP), OOP (I am very interested in fully grasping OOP, it so my worst area right now), metaprogramming.

So right now I am somewhat productive in Python and need to decide, do I fully commit to Python or pivot to some other language. I don't have that much time to waste but on the other side, I would rather learn something where I can have more freedom to explore and learn advanced subjects.

Here is my question: why is Ruby so known for creating DSL languages ? Why is metaprogramming such and important concept there and not in Python ? Python have decorators, metaclassess and AST module, how is Ruby in this regard ? Why is everyone in Python scared to touch those tools ? Are Python metaclasses not powerfull enough ?

I also hate the word 'pythonic' and community insitance on using uniform programming style for each problem.
## [4][How to generate fake text in multiple languages](https://www.reddit.com/r/ruby/comments/j79nvn/how_to_generate_fake_text_in_multiple_languages/)
- url: https://www.reddit.com/r/ruby/comments/j79nvn/how_to_generate_fake_text_in_multiple_languages/
---
I'm trying to generate random words in multiple languages. Something like Faker::Lorem.words

How would I do this in Ruby
## [5][μ-observers - Simple and powerful implementation of the observer pattern.](https://www.reddit.com/r/ruby/comments/j6qle5/μobservers_simple_and_powerful_implementation_of/)
- url: https://www.reddit.com/r/ruby/comments/j6qle5/μobservers_simple_and_powerful_implementation_of/
---
&amp;#x200B;

[https:\/\/github.com\/serradura\/u-observers](https://preview.redd.it/ankry7il8or51.png?width=616&amp;format=png&amp;auto=webp&amp;s=369eb78bb7cfec252615ff3d85b48995e7c4a6ab)

**Compatibility**

|Ruby|ActiveRecord/Rails|
|:-|:-|
|\&gt;= 2.2.0|\&gt;=3.2, &lt; 6.1|

*Note: The ActiveRecord isn't a dependency, but you could add a module to enable some static methods that were designed to be used with its* [*callbacks*](https://guides.rubyonrails.org/active_record_callbacks.html)*.*

**About**

This gem implements the observer pattern [\[1\]](https://en.wikipedia.org/wiki/Observer_pattern)[\[2\]](https://refactoring.guru/design-patterns/observer) (also known as publish/subscribe). It provides a simple mechanism for one object to inform a set of interested third-party objects when its state changes.

Ruby's standard library [has an abstraction](https://ruby-doc.org/stdlib-2.7.1/libdoc/observer/rdoc/Observable.html) that enables you to use this pattern. But its design can conflict with other mainstream libraries, like the [ActiveModel/ActiveRecord](https://api.rubyonrails.org/classes/ActiveModel/Dirty.html#method-i-changed), which also has the [changed](https://ruby-doc.org/stdlib-2.7.1/libdoc/observer/rdoc/Observable.html#method-i-changed) method. In this case, the behavior of the Stdlib will be compromised.

Because of this issue, I decided to create a gem that encapsulates the pattern without changing the object's implementation so much. The Micro::Observers includes just one instance method in the target class (its instance will be the observed subject).

**Usage examples**

*Attaching and notifying observers*

[https://github.com/serradura/u-observers#usage](https://github.com/serradura/u-observers#usage)

https://preview.redd.it/468dj9y59or51.png?width=1471&amp;format=png&amp;auto=webp&amp;s=652d3f79f6a37c28fbeb16e06d0d5b595ed43803

*ActiveRecord integration*

[https://github.com/serradura/u-observers#activerecord-and-activemodel-integrations](https://github.com/serradura/u-observers#activerecord-and-activemodel-integrations)

https://preview.redd.it/emcwu2kp9or51.png?width=1259&amp;format=png&amp;auto=webp&amp;s=3e109c21d9183d7188929194b130ac06f2d7b761

What do you think about it? Does it look good?
## [6][How's the Performance of Ruby 3.0.0-preview1?](https://www.reddit.com/r/ruby/comments/j6ssd0/hows_the_performance_of_ruby_300preview1/)
- url: https://www.fastruby.io/blog/rails/performance/ruby/hows-the-performance-of-ruby-3.0.0-preview1.html
---

## [7][Multitenancy with Postgres schemas: key concepts explained](https://www.reddit.com/r/ruby/comments/j6wxp3/multitenancy_with_postgres_schemas_key_concepts/)
- url: https://blog.arkency.com/multitenancy-with-postgres-schemas-key-concepts-explained/
---

## [8][Generating an Increasing unique number](https://www.reddit.com/r/ruby/comments/j6sbjt/generating_an_increasing_unique_number/)
- url: https://www.reddit.com/r/ruby/comments/j6sbjt/generating_an_increasing_unique_number/
---
 I want to create an unique value. preferably something that's user friends  i.e. a-z and maybe numbers. 

Approach 1. 

for generating numbers 

  
I did something like this. 

     (Time.now.to_f * 100000).to_i

And locally it does look like it does it.   


    2.4.7 :050 &gt; 10.times do 
    2.4.7 :051 &gt;     puts (Time.now.to_f * 100000).to_i
    2.4.7 :052?&gt;   end
    160207178626801
    160207178626804
    160207178626805
    160207178626806
    160207178626807
    160207178626808
    160207178626809
    160207178626810
    160207178626811
    160207178626812
     =&gt; 10

Approach 2.   


Generating a hex and then converting to base 26. However I'm not sure if this will be unique.   


Can someone help me out with this?
## [9][Array.select question](https://www.reddit.com/r/ruby/comments/j6vkqb/arrayselect_question/)
- url: https://www.reddit.com/r/ruby/comments/j6vkqb/arrayselect_question/
---
&amp;#x200B;

[Hey guys when I use this code it just returns the original array. When I use array.select! and mutate the original array it works but why? Shouldn't array.select return an array with all elements that return true without modifying the existing array? Thanks.](https://preview.redd.it/swvwmukumpr51.png?width=279&amp;format=png&amp;auto=webp&amp;s=a6efc4d8f30b961b46b5445971035b594c4005cd)
## [10][Does ruby apply options in the hashbang even if it's called explicitly?](https://www.reddit.com/r/ruby/comments/j6muvf/does_ruby_apply_options_in_the_hashbang_even_if/)
- url: https://www.reddit.com/r/ruby/comments/j6muvf/does_ruby_apply_options_in_the_hashbang_even_if/
---
I put `-wU` in the hashbang. Do I still need the `encoding` magic comment for case I decided to call `ruby file.rb` instead of just `./file.rb`? If I don't, is this consistent across all platforms?

What about when there's `env` in the hashbang?

Edit: okay, it *does* read the options, sorry for asking before trying. But still I can't test on other OSs than Linux, can please someone here try?

Edit 2: Looking through the source, there's a testcase for it:
* Link fixed on current commit: &lt;https://github.com/ruby/ruby/blob/62abdbadf2937372924ef68aadff5191fc0f0880/test/ruby/test_rubyoptions.rb#L421&gt;
* Link fixed on `master` (no guarantee it'll survive till the end) &lt;https://github.com/ruby/ruby/blob/master/test/ruby/test_rubyoptions.rb#L421&gt;
