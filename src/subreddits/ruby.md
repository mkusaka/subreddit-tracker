# ruby
## [1][μ-observers - Simple and powerful implementation of the observer pattern.](https://www.reddit.com/r/ruby/comments/j6qle5/μobservers_simple_and_powerful_implementation_of/)
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
## [2][dogweather/non_empty_array](https://www.reddit.com/r/ruby/comments/j6jbzv/dogweathernon_empty_array/)
- url: https://github.com/dogweather/non_empty_array
---

## [3][Does ruby apply options in the hashbang even if it's called explicitly?](https://www.reddit.com/r/ruby/comments/j6muvf/does_ruby_apply_options_in_the_hashbang_even_if/)
- url: https://www.reddit.com/r/ruby/comments/j6muvf/does_ruby_apply_options_in_the_hashbang_even_if/
---
I put `-wU` in the hashbang. Do I still need the `encoding` magic comment for case I decided to call `ruby file.rb` instead of just `./file.rb`? If I don't, is this consistent across all platforms?

What about when there's `env` in the hashbang?

Edit: okay, it *does* read the options, sorry for asking before trying. But still I can't test on other OSs than Linux, can please someone here try?
## [4][Help with Inject](https://www.reddit.com/r/ruby/comments/j6e5y5/help_with_inject/)
- url: https://www.reddit.com/r/ruby/comments/j6e5y5/help_with_inject/
---
&amp;#x200B;

[In this problem we take in an unknown quantity of numbers and find the common divisors between them all. The factors method finds all of the factors for each number. Once we do nums.map we get a 2D array, each array contains all of the factors of the original number. I don't understand why we can then just .inject that 2D array to get all of the common elements of the arrays? Thanks](https://preview.redd.it/4uy1vw86mjr51.png?width=655&amp;format=png&amp;auto=webp&amp;s=e45e022b73aed51f20f806da7374d66d5ddb9e5d)
## [5][Rails 6 and Stimulus.js - a quick and painless launch](https://www.reddit.com/r/ruby/comments/j62skn/rails_6_and_stimulusjs_a_quick_and_painless_launch/)
- url: https://longliveruby.com/articles/rails-6-stimulus-js
---

## [6][Is there a gem for XML to JSON convertion?](https://www.reddit.com/r/ruby/comments/j6ht5w/is_there_a_gem_for_xml_to_json_convertion/)
- url: https://www.reddit.com/r/ruby/comments/j6ht5w/is_there_a_gem_for_xml_to_json_convertion/
---
Hi there 
I have a request to process a XML file 
I got pretty surprised why an api return XML file instead of JSON 
So i am wondering if would be more productive have this file First Converted to a JSON before start my data manipulation code . 

Any thougths ? Is there any node package or Ruby gem for that?
Best Regards
## [7][Return to a pervious part of the code](https://www.reddit.com/r/ruby/comments/j6b428/return_to_a_pervious_part_of_the_code/)
- url: https://www.reddit.com/r/ruby/comments/j6b428/return_to_a_pervious_part_of_the_code/
---
Hello, Everyone!

I am working on creating a password manager out of boredom. What I have now is just a bare metal way of controlling users. I'm using pretty simple control flow and a method to achieve this. It works well, but one part of it I can't figure out.

\---Snip---

puts "Choose profile to begin."puts ""puts "BlakeDianeOther"def user()u = gets.chompenduser = user().downcasewhile user doif user == "blake"puts "Enter Password"breakelsif user == "diane"puts "Enter Password:"breakelsif user == "other"puts "Enter username:"breakelseputs "Please choose your profile..."breakendendif user == "blake"branpass = gets.chompelsif user == "diane"despass = gets.chompelsif user == "other"othuser = gets.chompelsereturnendbranpasscheck = "Password000"despasscheck = "Passwor001"othusercheck = "Johnny".downcasewhile user doif branpass == branpasscheck and user == "blake"puts "Unlocking: Blake"breakelsif despass == despasscheck and user == "diane"puts "Unlocking: Diane"breakelsif othuser == othusercheck and user == "other"puts "Redirecting to Password set..."breakelseputs "Password invalid..."return user()endend

\---Snip---

This basically lists the users and allows you to choose the user that you want. One you choose the user, it prompts for a password.

My issue is that when I get to the point of validating the predetermined password, if it is correct, it works, but it if is incorrect, I can't figure out how to get it to loop back to inputting the password until it is correct. Would define the password validation as a method and calling it when it is incorrect work?

Any help is appreciated!

&amp;#x200B;

EDIT: I apologize for the formatting of the code. 
## [8][Why You Should Migrate your Heroku Postgres Database to AWS RDS](https://www.reddit.com/r/ruby/comments/j61rm6/why_you_should_migrate_your_heroku_postgres/)
- url: https://pawelurbanek.com/heroku-postgres-aws-rds
---

## [9][GVL/GIL free concurrency?](https://www.reddit.com/r/ruby/comments/j67ayj/gvlgil_free_concurrency/)
- url: https://www.reddit.com/r/ruby/comments/j67ayj/gvlgil_free_concurrency/
---
I am looking for something like `Thread`s concurrency, but without GVL, so far I tried a couple of gems like Async and Polyphony, but those involve fibers, I am looking for something about true parallel execution.
## [10][Trouble with \W regex matcher in Ruby](https://www.reddit.com/r/ruby/comments/j65t4d/trouble_with_w_regex_matcher_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/j65t4d/trouble_with_w_regex_matcher_in_ruby/
---
I need to remove all non-word via regex and this is how I do it.   


 In English, it works fine. 

    2.4.7 :010 &gt; 'hello  $%'.gsub(/\W/, '')
     =&gt; "hello" 

But when I try with Japanese,   


    2.4.7 :009 &gt; 'こんにちは  $%'.gsub(/\W/, '')
     =&gt; ""

the word here is "hello" in Japanese but regex thinks it's a not a word and removes it   


Can someone help me why this is happening and how to fix this? I was under the impression that regex would handle different languages.   


PS. I tried the same in Python 2 and well and the behavior was just like here.
