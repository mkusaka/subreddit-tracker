# ruby
## [1][Ruby on Rails authorization using CanCanCan](https://www.reddit.com/r/ruby/comments/gcoypt/ruby_on_rails_authorization_using_cancancan/)
- url: https://www.reddit.com/r/ruby/comments/gcoypt/ruby_on_rails_authorization_using_cancancan/
---
Hi ruby family,

As an initiative to give back to the community, I have started writing a series of blogs on ruby and ruby on rails. Planning to create more content in the future to help share the knowledge. I just published a post about Authorization on Ruby on Rails using CanCanCan. Do check it out and let me know your thoughts.

[https://addytalks.tech/2020/05/03/ruby-on-rails-authorization-with-cancancan/](https://addytalks.tech/2020/05/03/ruby-on-rails-authorization-with-cancancan/)
## [2][Coronavirusapi.com is seeking help from Ruby developers](https://www.reddit.com/r/ruby/comments/gcawyk/coronavirusapicom_is_seeking_help_from_ruby/)
- url: https://twitter.com/josephdelong/status/1256625135831977984?s=20
---

## [3][Rio - Ruby I/O Facilitator - Oldie But Goodie!](https://www.reddit.com/r/ruby/comments/gcfb5o/rio_ruby_io_facilitator_oldie_but_goodie/)
- url: http://rio4ruby.com/rio/
---

## [4][How to build first telegram bot with Ruby? 8 lines of code and 5 minutes is enough!](https://www.reddit.com/r/ruby/comments/gc3g2v/how_to_build_first_telegram_bot_with_ruby_8_lines/)
- url: https://youtu.be/WefRd8W5KuA
---

## [5][A Ruby IO object that can be composed of other IO objects](https://www.reddit.com/r/ruby/comments/gcam26/a_ruby_io_object_that_can_be_composed_of_other_io/)
- url: https://github.com/agis/multi_io
---

## [6][GitHub README vs GitBook for documenting a medium-complexity gem?](https://www.reddit.com/r/ruby/comments/gcf3zp/github_readme_vs_gitbook_for_documenting_a/)
- url: https://www.reddit.com/r/ruby/comments/gcf3zp/github_readme_vs_gitbook_for_documenting_a/
---
I have a gem, [tabulo](https://github.com/matt-harvey/tabulo) that is getting to the point where the README is pretty long, and contains more than just a bare introduction to the gem. It really acts as a user manual for the gem. I like to have a README that's pretty substantial, as I think this is much more accessible than just pointing to the [rubydocs](https://www.rubydoc.info/gems/tabulo) and leaving it at that. However, it bothers me that there is no sidebar navigation&amp;mdash;meaning readers have to keep jumping back to the top to see the contents. It also feels like the README is getting too long for just one page.

Lately I've been looking at [GitBook](https://gitbook.com) and wondering whether to migrate the content of the README there. It provides a very nicely formatted layout with sidebar navigation, as well as a prompt that lets users easily suggest edits to the documentation. See the GitBook for [FoalTS](https://https://foalts.gitbook.io/docs/), as one example.

My question is: Do the readability and navigability benefits of putting the user guide in GitBooks (or something similar) outweigh the "accessibility" benefits (for want of a better word) of having the user guide simply be the GitHub project README? Having everything in the README means that once users (or potential users) land on the GitHub repo, the "complete" guide to using the gem is "zero clicks" away; it can be read through continually simply by scrolling down. But it's not as navigable or as slick as GitBook. I'm trying to weigh these things up. If anyone has any experience, thoughts or feedback on this, that would be greatly appreciated.
## [7][Seeking advice on the Ruby language learning curves](https://www.reddit.com/r/ruby/comments/gc1pjc/seeking_advice_on_the_ruby_language_learning/)
- url: https://www.reddit.com/r/ruby/comments/gc1pjc/seeking_advice_on_the_ruby_language_learning/
---
For the people that already into the in-depth level of the Ruby language, what is your thought on the current standing of Ruby? and what was the point of realization that makes you feel like the needs for Ruby?

I have had experiences with the Ruby on Rails framework and only come to learn the Ruby language on-demand for debugging. I have been spoiled by the framework convention itself and left me with an uncertain/unfulfillment feeling. 

I'm all ear for any advice and starting-point recommendation that might not just simply on the topic of learning the coding language.
## [8][Trying out an Amazon Coding Problem on Run Length Encoding, doing it in Ruby and the one used is Java, not sure why solution isn't working](https://www.reddit.com/r/ruby/comments/gccw6y/trying_out_an_amazon_coding_problem_on_run_length/)
- url: https://www.reddit.com/r/ruby/comments/gccw6y/trying_out_an_amazon_coding_problem_on_run_length/
---
https://www.youtube.com/watch?time_continue=52&amp;v=mjZpZ_wcYFg&amp;feature=emb_logo

so basically I have this:

    def run_length_encoding(str)
      count = 1
      new_str = ''
      previous_char = str[0]

      str.chars[1..-1].each_with_index do |char, i |
        if char == previous_char
          count =+ 1
        else
          new_str += (count.to_s + previous_char)
          count = 1
          previous_char = char
        end
      end

      new_str
    end

the idea is that if you have say 'aabbbdddda' string input, you should get '2a3b4d1a', so order matters, thus we use an array as the person in the video does.  However I'm not getting results when I try it in repl.it. I try my method and get:

    run_length_encoding('aabbccccda')
    =&gt; "1a1b1c1d"

If anyone has ideas please feel free to let me know.  I'm doing it slightly different than the exact why the person in the video is - it seems they assign `0` to the `previous_char` and I'm not sure on that, I assign the first letter in the string - but I'm not trying to copy exactly anyways.
## [9][Can't use ruby installed gem](https://www.reddit.com/r/ruby/comments/gc5jmu/cant_use_ruby_installed_gem/)
- url: https://www.reddit.com/r/ruby/comments/gc5jmu/cant_use_ruby_installed_gem/
---
I am new to ruby and would like to use the ruby gem [Word-Search Puzzle Generator](https://www.rubydoc.info/gems/wordsearch-puzzle)

I am using Mac OS X 10.14.6.

This is what I did:

&amp;#x200B;

1. Updated brew using `brew update &amp;&amp; brew upgrade`
2. Ran `brew doctor` to ensure there were no problems (no problems were found).
3. Ran `brew install ruby`
4. Ran echo `'export PATH="/usr/local/opt/ruby/bin:$PATH"' &gt;&gt; ~/.bash_profile`
5. Ran `source ~/.bash_profile`
6. Ran `gem install wordsearch`. This returned the following error:  
`ERROR:  While executing gem ... (Errno::EACCES)`

`Permission denied @ rb_sysopen - /usr/local/lib/ruby/gems/2.7.0/gems/wordsearch-1.0.0/ext/wordsearch/accent.c`

7. So I tried `sudo gem install wordsearch`. This worked.

8. However, when I run w`ordsearch nitwit blubber oddment tweak` it says ***command not found.***

Any ideas? Thanks!
## [10][Guild renamed to Ractor](https://www.reddit.com/r/ruby/comments/gbhkjn/guild_renamed_to_ractor/)
- url: https://www.reddit.com/r/ruby/comments/gbhkjn/guild_renamed_to_ractor/
---
[https://github.com/ko1/ruby/blob/ractor/ractor.ja.md](https://github.com/ko1/ruby/blob/ractor/ractor.ja.md)
