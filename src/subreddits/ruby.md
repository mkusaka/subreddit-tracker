# ruby
## [1][Best way to handle money?](https://www.reddit.com/r/ruby/comments/j0gsrt/best_way_to_handle_money/)
- url: https://www.reddit.com/r/ruby/comments/j0gsrt/best_way_to_handle_money/
---
I have a simple ruby script to generate payroll data. I am adding and multiplying but at the end of the day, I only need to support USD (for now). 

My database type is decimal(8,4) because I was told that’s accounting standards, so future proofing. 

However, what is the best way to handle addition and multiplication within my script for money? Is it BigDecimal? I don’t expect the numbers to be big (max $10,000) but doing it the “right way” is probably best. 

Thank you!
## [2][REPL playground for various Ruby Engines](https://www.reddit.com/r/ruby/comments/j0nsey/repl_playground_for_various_ruby_engines/)
- url: https://rubyapi.org/repl
---

## [3][Gem for updating dependencies of your project](https://www.reddit.com/r/ruby/comments/j07s0c/gem_for_updating_dependencies_of_your_project/)
- url: https://www.reddit.com/r/ruby/comments/j07s0c/gem_for_updating_dependencies_of_your_project/
---
Hi everybody,

Me and a friend made a simple tool to update dependencies of your ruby projects. Please check it out and let us know what you think. 

[https://github.com/spandx/dependency\_bumper](https://github.com/spandx/dependency_bumper)
## [4][Ruby 3.0.0 preview1 released](https://www.reddit.com/r/ruby/comments/iznkst/ruby_300_preview1_released/)
- url: https://www.ruby-lang.org/en/news/2020/09/25/ruby-3-0-0-preview1-released/
---

## [5][TIL that DateTime is considered deprecated](https://www.reddit.com/r/ruby/comments/izr2lm/til_that_datetime_is_considered_deprecated/)
- url: https://bugs.ruby-lang.org/issues/15712#note-4
---

## [6][Does anyone have a copy of the book `Building Git` that I could borrow? Broke college student.](https://www.reddit.com/r/ruby/comments/j0068g/does_anyone_have_a_copy_of_the_book_building_git/)
- url: https://www.reddit.com/r/ruby/comments/j0068g/does_anyone_have_a_copy_of_the_book_building_git/
---
I'm happy to pay for the book at a later date, once I graduate. I believe in supporting creators, however; I can't afford the book and really itching to work through it this weekend.  


It's a shot in the dark but if anyone is willing to help, I would be forever grateful.
## [7][Question about calling method / if statements](https://www.reddit.com/r/ruby/comments/izx3q5/question_about_calling_method_if_statements/)
- url: https://www.reddit.com/r/ruby/comments/izx3q5/question_about_calling_method_if_statements/
---
&amp;#x200B;

&amp;#x200B;

[Hey guys so in rspec I get an error on the first picture but not on the second picture. The if statement is checking for true or false then moving onto \\"check in successful\\" I don't understand why you don't have to call the .add\_occupant method again after the if statement to actually complete the task. That \\"if\\" line doesn't actually complete the method does it? If so, why? Thanks!](https://preview.redd.it/z8x96an6gep51.png?width=532&amp;format=png&amp;auto=webp&amp;s=6b1ced3e5b3b440e45f97f8e1f969f3ab8f2f3d9)

https://preview.redd.it/s2fzz2ywfep51.png?width=480&amp;format=png&amp;auto=webp&amp;s=1691d279438663d1e510947f64257a22acbcd18e
## [8][HereDocs tokens as method argument](https://www.reddit.com/r/ruby/comments/izk8o8/heredocs_tokens_as_method_argument/)
- url: https://medium.com/rubycademy/heredocs-tokens-as-method-argument-8fe3ed723d96
---

## [9][Wanted to share my method for managing custom fields in Ruby on Rails](https://www.reddit.com/r/ruby/comments/izo8s6/wanted_to_share_my_method_for_managing_custom/)
- url: https://syndicode.com/blog/manage-custom-fields-for-an-activerecord-object-in-rails/
---

## [10][There are so many ways to set/get an array index... confused!](https://www.reddit.com/r/ruby/comments/izwu5h/there_are_so_many_ways_to_setget_an_array_index/)
- url: https://www.reddit.com/r/ruby/comments/izwu5h/there_are_so_many_ways_to_setget_an_array_index/
---
Hi,

I get that Ruby is all dynamic and proud of it, but how the heck am I to know what is the preferable way to access/set arrays? the same goes for a lot of operations...

Here my array annotations:

&amp;#x200B;

    a = [1, ["a", "b"], 4] \\ arr[1][0] \\ arr.dig(3,0) \\ value_at, a[2,3] = .. \\ 
    
    a[2..3] \\ slice, a.[]=(0, "first") \\ a.[](2) \\ %w(a b c) \\ %W({a} b c) \\ .to_ary \\ .to_arr

edit: I forgot `[1,2][0] \\ &lt;&lt; \\ unshift(0) \\ push(1,2,3)`

It should be renamed to Ruby++ or even better Common Ruby :)

Thanks!

&amp;#x200B;

&amp;#x200B;

https://preview.redd.it/2tptv7d0lep51.jpg?width=318&amp;format=pjpg&amp;auto=webp&amp;s=847854de90f6c9111938aae375a2fe8965efa2f9
