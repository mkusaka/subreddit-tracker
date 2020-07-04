# ruby
## [1][Ruby 3.0 JIT on Rails (slides from Ruby commiter k0kubun)](https://www.reddit.com/r/ruby/comments/hl1ydr/ruby_30_jit_on_rails_slides_from_ruby_commiter/)
- url: https://speakerdeck.com/k0kubun/ruby-3-dot-0-jit-on-rails
---

## [2][Best technologies for embedding dashboards/reports on a site run by Ruby on Rails.](https://www.reddit.com/r/ruby/comments/hl0vtz/best_technologies_for_embedding_dashboardsreports/)
- url: https://www.reddit.com/r/ruby/comments/hl0vtz/best_technologies_for_embedding_dashboardsreports/
---
We need to embed a fairly complex dashboard on a portal that will be seen by a user after they've logged in.

 Initially we were going to do this using a solution like Tableau or Power BI embedded, but due to costs involved, we feel it's more sustainable if we build our own solution using an existing dashboarding library instead.

I figure someone here has to have done something like this using Ruby on Rails. I'm looking for suggestions for what the best libraries might be - possibly something like Plotly?
## [3][Is module_function really the same as extend self ?](https://www.reddit.com/r/ruby/comments/hl1j9j/is_module_function_really_the_same_as_extend_self/)
- url: https://medium.com/rubycademy/is-module-function-really-the-same-as-extend-self-ac1e96a1cda0
---

## [4][Take the 2020 Ruby on Rails Developer Community Survey](https://www.reddit.com/r/ruby/comments/hkuw5o/take_the_2020_ruby_on_rails_developer_community/)
- url: /r/rails/comments/grlcry/open_2020_ruby_on_rails_developer_community_survey/
---

## [5][Polyphony - Fine-Grained Concurrency for Ruby](https://www.reddit.com/r/ruby/comments/hklp43/polyphony_finegrained_concurrency_for_ruby/)
- url: https://github.com/digital-fabric/polyphony
---

## [6][Rewrite Kernel#tap with Ruby by k0kubun](https://www.reddit.com/r/ruby/comments/hker0s/rewrite_kerneltap_with_ruby_by_k0kubun/)
- url: https://github.com/ruby/ruby/pull/3281
---

## [7][Fun Facts about Ruby #8](https://www.reddit.com/r/ruby/comments/hkf8zk/fun_facts_about_ruby_8/)
- url: https://twitter.com/RubyCademy/status/1278962830637662208
---

## [8][I thought I understood "self", but I don't](https://www.reddit.com/r/ruby/comments/hk6upv/i_thought_i_understood_self_but_i_dont/)
- url: https://www.reddit.com/r/ruby/comments/hk6upv/i_thought_i_understood_self_but_i_dont/
---
https://github.com/Ubuntu19019/learnruby/blob/master/Understanding%20self

So I'm doing a battleship project. It has an rspec file and I'm doing what it tells me to do. The attack method (which I believe is an instance method)  is tripping me up. 

I thought self is used for an instance of the class Board. new_board = Board.new(n).  new_board would be an instance. There's no function to return @grid. But attr_reader for @size. So I could do self.size.(some other instance method) right? Since I can't do that with @grid I thought I would have to do 
&gt;if @grid[position] == :S

That doesn't work. My theory is that since there's no method for returning @grid I can't call another method on it. If there was a method I could call either self.grid[position] == :S or @grid[position] == :S

If I had a method for returning @grid is it better to directly call @grid or do self.grid(is this an instance)?

I also need clarification on what self refers to. 

&gt;if self[position] == :S

Is self referring to an instance of the class? If that's the case what exactly does the initialize method do? Does initialize build the grid inside of the Board class, then self[postion] would be like new_board[position] since new_board is an instance of the Board class?

I'm really confused. After I digest this information I want to know more about class methods vs instance methods and why I would choose to make one over the other.
## [9][How often do you write scrapping scripts?](https://www.reddit.com/r/ruby/comments/hk8glk/how_often_do_you_write_scrapping_scripts/)
- url: https://www.reddit.com/r/ruby/comments/hk8glk/how_often_do_you_write_scrapping_scripts/
---
I was doing some website and the client needed me to copy some data from a blog. So naturally I said, this would be a good time to write a webscrapper. It took me a while to get the hang of Nokogiri but I did. Which made me wonder, how often do you write scrapping scripts in your experience?
## [10][Can you Scrap images with Nokogiri? Or do I have to go the Selenium route?](https://www.reddit.com/r/ruby/comments/hk8r1b/can_you_scrap_images_with_nokogiri_or_do_i_have/)
- url: https://www.reddit.com/r/ruby/comments/hk8r1b/can_you_scrap_images_with_nokogiri_or_do_i_have/
---
As the title says, I'm looking to get some images from a blog. Any help would be appreciated
