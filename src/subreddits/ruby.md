# ruby
## [1][GitHub Actions](https://www.reddit.com/r/ruby/comments/io6g0t/github_actions/)
- url: https://www.driftingruby.com/episodes/github-actions?utm_medium=social&amp;utm_campaign=weekly_episode&amp;utm_source=reddit
---

## [2][retest: a gem to help you refactor your ruby projects faster](https://www.reddit.com/r/ruby/comments/io0uae/retest_a_gem_to_help_you_refactor_your_ruby/)
- url: https://github.com/AlexB52/retest
---

## [3][Survey finds only 3% of Ruby on Rails developers use Windows](https://www.reddit.com/r/ruby/comments/inpgox/survey_finds_only_3_of_ruby_on_rails_developers/)
- url: https://developers.slashdot.org/story/20/09/06/0028214/survey-finds-only-3-of-ruby-on-rails-developers-use-windows
---

## [4][Advice on editing a YAML file with Ruby](https://www.reddit.com/r/ruby/comments/invxlx/advice_on_editing_a_yaml_file_with_ruby/)
- url: https://www.reddit.com/r/ruby/comments/invxlx/advice_on_editing_a_yaml_file_with_ruby/
---
Hello,

I'm currently having some difficulty working with yaml files.

What I'd like to be able to do is load one YAML file, and then load another YAML file which contains some changes or additions, update the original yaml hash based on those changes and then write the file back to yaml. Something like this: https://pastebin.com/waCwzcWL

The data hash appears to be updated correctly. However, when the hash is written to the file as YAML, the file experiences some unexpected changes. There are four top level keys in the original file. Any child keys of the FIRST key in the file where their value is a Hash or an Array have &amp;1, &amp;2, &amp;3 (etc., incrementing for each key in order) appended to the end of the line, and the key is also duplicated into the other top level keys, with *1, *2, *3 etc., appended to the end of the key's name. However the value is not duplicated.

Does anyone have any experience with this sort of thing and can offer some advice on what is happening here?

Many thanks in advance.
## [5][HashWithIndifferentAccess in Rails](https://www.reddit.com/r/ruby/comments/io1trg/hashwithindifferentaccess_in_rails/)
- url: https://www.sandipmane.dev/hash-with-indifferent-access-in-rails
---

## [6][Command-line Arguments in Ruby: Part I](https://www.reddit.com/r/ruby/comments/inmgds/commandline_arguments_in_ruby_part_i/)
- url: https://medium.com/rubycademy/command-line-arguments-in-ruby-part-i-8a89479cb70f
---

## [7][Integrate Doorkeeper gem – API only ruby on rails course (chapter 10)](https://www.reddit.com/r/ruby/comments/inne5q/integrate_doorkeeper_gem_api_only_ruby_on_rails/)
- url: https://duetcode.io/rails-api-only-course/integrate-doorkeeper-gem
---

## [8][converting a hash into date object, help!!](https://www.reddit.com/r/ruby/comments/inpti8/converting_a_hash_into_date_object_help/)
- url: https://www.reddit.com/r/ruby/comments/inpti8/converting_a_hash_into_date_object_help/
---
I'm new to programming and I can't quite understand how to convert a hash string into date object.

I created a hash out of file that contains date as a string in different formats (like *September 1988*, the other line would be  *July 11th 1960*, and sometimes *year* only)

    require 'date'
    
    def create_book_hash(book_array)
      {
        link: book_array[0],
        title: book_array[1],
        author: book_array[2],
        pages: book_array[3].to_i,
        date: book_array[4],
        rating: book_array[5].to_f,
        genre: book_array[6]
      }
    end
    
    def books_sorted_by_date (books_array)
        books_array.sort_by { |key| Date.strptime(key[:date], '%Y, %m') }
      end
    
    book_file= File.read("books.txt")
        .split("\n")
        .map { |line| line.split("|")}
        .map { |book_array| create_book_hash(book_array)}
    puts books_sorted_by_date(book_file)

I'm trying to sort books by date, so it would be in ascending order by year and since I have different string types,  i put a hash key as the first argument in `strptime` to access all the values in **:date .** And that gives me  `\`strptime': invalid date (Date::Error).` I don't understand why and what can I do to convert these strings into date objects?
## [9][Building a Discord bot in Ruby and vim - live stream starting soon](https://www.reddit.com/r/ruby/comments/imwwso/building_a_discord_bot_in_ruby_and_vim_live/)
- url: https://www.twitch.tv/drzel1
---

## [10][Using git with ruby -- gitti gem - (yet) another (lite) git command line wrapper / library](https://www.reddit.com/r/ruby/comments/imx3i8/using_git_with_ruby_gitti_gem_yet_another_lite/)
- url: https://github.com/rubycoco/gitti/tree/master/gitti
---

