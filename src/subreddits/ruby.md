# ruby
## [1][Started my first roguelike with Dragon Ruby](https://www.reddit.com/r/ruby/comments/fexv07/started_my_first_roguelike_with_dragon_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fexv07/started_my_first_roguelike_with_dragon_ruby/
---
I decided to try out Dragon Ruby while also learning to develop my first roguelike. I was able to finish generating a basic dungeon map using the assets from [https://0x72.itch.io/dungeontileset-ii](https://0x72.itch.io/dungeontileset-ii) over the last few days. So far it's been pretty fun and I've already learned a lot about the Ruby programming language and 2D game development in general.

https://i.redd.it/bo2cbv0h7al41.gif
## [2][Improving my arrays](https://www.reddit.com/r/ruby/comments/ff3rqu/improving_my_arrays/)
- url: https://www.reddit.com/r/ruby/comments/ff3rqu/improving_my_arrays/
---
For a friend's game in Ruby 1.8 I'm coding something and while visually I've got it under control, I feel like I'm not organizing my data properly:

The idea is that the player can unlock several page entries in a book, and then, within each entry, there are nine other pieces of information to unlock. The thing is, I don't think I'm thinking right when it comes to how should this information be unlocked.

For example, an array should be all false at first, or should it? I need to initialize the data, and right now, I'm doing it like this:

[https://i.imgur.com/nR3eziA.png](https://i.imgur.com/nR3eziA.png)

I hope someone can help me figure out a better way of arranging this example $unlockedEntries array, as to store the data. 

https://preview.redd.it/hpr80kex6cl41.png?width=739&amp;format=png&amp;auto=webp&amp;s=167710e831f6c14956e7110771a1b3b3e1fd59ed
## [3][Learning Ruby without webdev?](https://www.reddit.com/r/ruby/comments/fer4mz/learning_ruby_without_webdev/)
- url: https://www.reddit.com/r/ruby/comments/fer4mz/learning_ruby_without_webdev/
---
Ruby seems pretty huge for web development type stuff, but I've also heard of its scripting capabilities applicable to ops and security, which interest me as I've come off knowing some rudimentary programming knowledge in Java and Python, not yet having JS knowledge(HTML/CSS of course). 

Could someone really get a feel for the Ruby language, hence "learning" it to a degree, without stepping much into web development?
## [4][How to parse nested json in ruby?](https://www.reddit.com/r/ruby/comments/feg2a1/how_to_parse_nested_json_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/feg2a1/how_to_parse_nested_json_in_ruby/
---
Hey guys,

I'm trying to use an API that returns nested json. I'm using json.parse but it is splitting the json file into two large hashes, when I want to go down one or two more levels. Does anyone know how I could do this? Thanks.
## [5][barcharts gem - print barcharts in your terminal in ascii (e.g. *) or with unicode block elements (e.g. ▏▎▍▋▊▉)](https://www.reddit.com/r/ruby/comments/feazzy/barcharts_gem_print_barcharts_in_your_terminal_in/)
- url: https://github.com/feedreader/pluto/tree/master/barcharts
---

## [6][Eager loading Active Storage models to avoid N+1 issues](https://www.reddit.com/r/ruby/comments/fec6wr/eager_loading_active_storage_models_to_avoid_n1/)
- url: https://blog.saeloun.com/2020/03/06/eagerload-active-storage-models
---

## [7][How do I know if the statement is true or false in Ruby programming? is my understanding correct?](https://www.reddit.com/r/ruby/comments/fe7cne/how_do_i_know_if_the_statement_is_true_or_false/)
- url: https://i.redd.it/vf1m59ki0zk41.png
---

## [8][What's new in Kiba ETL v3 (visually explained)](https://www.reddit.com/r/ruby/comments/fdyca1/whats_new_in_kiba_etl_v3_visually_explained/)
- url: https://thibautbarrere.com/2020/03/05/new-in-kiba-etl-v3/
---

## [9][Why Pry is one of the most important tools a junior Rubyist can learn](https://www.reddit.com/r/ruby/comments/fdh5gh/why_pry_is_one_of_the_most_important_tools_a/)
- url: https://www.reddit.com/r/ruby/comments/fdh5gh/why_pry_is_one_of_the_most_important_tools_a/
---
As programmers we often have to mentally run code. To imagine how a program will behave given certain inputs. This is hard enough for experienced developers. But for juniors? It can seem impossible. In this article, Melissa Williams argues that pry is an invaluable tool for junior rubyists because it allows them to see exactly what is going on as their code is run. [https://www.honeybadger.io/blog/debugging-ruby-with-pry/](https://www.honeybadger.io/blog/debugging-ruby-with-pry/)
## [10][Mistakes I've made treating file paths as strings](https://www.reddit.com/r/ruby/comments/fdnxz8/mistakes_ive_made_treating_file_paths_as_strings/)
- url: https://www.reddit.com/r/ruby/comments/fdnxz8/mistakes_ive_made_treating_file_paths_as_strings/
---
It bugs me about how long I've manipulated paths as strings without having a failure, but in the last couple of months things have broken in Ruby and Node projects. So I wrote about [the issues I've faced and how we should approach path names to avoid them](https://philna.sh/blog/2020/03/04/mistakes-treating-paths-as-strings/) (TL;DR use the standard lib `Pathname` and `File.join` in Ruby).
