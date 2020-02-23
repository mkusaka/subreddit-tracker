# ruby
## [1][Fusrodah - a Roda + Sequel + Webpack/TailwindCSS boilerplate](https://www.reddit.com/r/ruby/comments/f85el7/fusrodah_a_roda_sequel_webpacktailwindcss/)
- url: https://www.reddit.com/r/ruby/comments/f85el7/fusrodah_a_roda_sequel_webpacktailwindcss/
---
I've been enjoying Roda as my sideproject/fun stack but constantly having to rebuild my stack every time I start a side project with some basic stuff like Webpack and whatnot so I made one for cloning. Inspired by [roda-sequel-stack](https://github.com/jeremyevans/roda-sequel-stack) but catered to my needs.

It combines both [jeremyevans/roda-sequel-stack](https://github.com/jeremyevans/roda-sequel-stack) and [taniarascia/webpack-boilerplate](https://github.com/taniarascia/webpack-boilerplate)

So here you go, for anyone who want to start a new project with a fully working webpack stack.

[choonggg/fusrodah](https://github.com/choonggg/fusrodah)

Forgive the name, it was in my head every time I think of Roda
## [2][Pending cops in Rubocop to make upgrade easier](https://www.reddit.com/r/ruby/comments/f89eny/pending_cops_in_rubocop_to_make_upgrade_easier/)
- url: https://prathamesh.tech/2020/02/23/pending-cops-to-make-updating-rubocop-easier/
---

## [3][Static vs Dynamic typing in the context of game development (why I use Ruby).](https://www.reddit.com/r/ruby/comments/f7udb6/static_vs_dynamic_typing_in_the_context_of_game/)
- url: https://www.twitch.tv/videos/555300001?filter=highlights&amp;sort=time
---

## [4][iniparser v1.0 gem - read / parse INI configuration, settings and data files into a hash (incl. named subsections)](https://www.reddit.com/r/ruby/comments/f86smc/iniparser_v10_gem_read_parse_ini_configuration/)
- url: https://github.com/datatxt/iniparser
---

## [5][The redo Keyword in Ruby](https://www.reddit.com/r/ruby/comments/f7pi1f/the_redo_keyword_in_ruby/)
- url: https://medium.com/rubycademy/the-redo-keyword-in-ruby-3f150d69e3c2
---

## [6][Aaron Patterson's thoughts on the pipeline emoji](https://www.reddit.com/r/ruby/comments/f7d117/aaron_pattersons_thoughts_on_the_pipeline_emoji/)
- url: https://github.com/ruby/ruby/pull/2273/
---

## [7][Namespacing Rails application?](https://www.reddit.com/r/ruby/comments/f7idfp/namespacing_rails_application/)
- url: https://www.reddit.com/r/ruby/comments/f7idfp/namespacing_rails_application/
---
I maintain a "majestic" monolith eCommerce platform  with 3 very distinct parts -- api, admin, store. And I'd love to namespace this application, so this distinction could be move obvious and I can easily use mutant testing.  


Has anyone went through such a hassle with their Rails apps? Any blog posts/ experience you can share?  


Namespacing seems obvious for controller code, but not that 'obvious' for code that is being shared between all these "distinct parts". Should I namespace models as well?   


Quite a lot of questions once I start thinking about it and very little information on such a 'transition' in the internet.
## [8][Best Developer Tools Trick](https://www.reddit.com/r/ruby/comments/f7czxv/best_developer_tools_trick/)
- url: https://blog.driftingruby.com/best-developer-tool-trick/
---

## [9][Need help with some refactoring in a cfndsl CloudFormation template](https://www.reddit.com/r/ruby/comments/f7kn5v/need_help_with_some_refactoring_in_a_cfndsl/)
- url: https://www.reddit.com/r/ruby/comments/f7kn5v/need_help_with_some_refactoring_in_a_cfndsl/
---
In [cfndsl](https://github.com/cfndsl/cfndsl), you can create an AWS CloudFormation template like this:

```
CloudFormation do
  EC2_Instance(:Example) do
    ImageId 'ami-12345678'
    Type 't1.micro'
  end
end
```

However real-world templates are usually much larger, but that's not the point.

Let's say that I have dozens of CloudFormation templates that share the same `EC2_Instance`, each with a possibly different `ImageId` and `Type`.

How can I extract this into a method (e.g. `ec2_instance`) so that I can refactor the `CloudFormation` block into this:

```
def ec2_instance(ami, type)
   # ???
end

CloudFormation do
  ec2_instance('ami-12345678', 't1.micro')
end

CloudFormation do
  ec2_instance('ami-90111213', 't2.small')
end

```

Thanks in advance.
## [10][Ruby Quiz - Code Challenge #18 - Up-to-Date? Version Check All Your Libraries](https://www.reddit.com/r/ruby/comments/f7dct8/ruby_quiz_code_challenge_18_uptodate_version/)
- url: https://github.com/planetruby/quiz/tree/master/018
---

