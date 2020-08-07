# ruby
## [1][Open Source Status Update - July 2020 (dry-validation, dry-rails and hanami 2.0)](https://www.reddit.com/r/ruby/comments/i5bbj2/open_source_status_update_july_2020_dryvalidation/)
- url: https://solnic.codes/2020/08/06/open-source-status-update-july-2020/
---

## [2][Quieres aprender Ruby &amp; Ruby on Rails?](https://www.reddit.com/r/ruby/comments/i4ye6k/quieres_aprender_ruby_ruby_on_rails/)
- url: https://www.reddit.com/r/ruby/comments/i4ye6k/quieres_aprender_ruby_ruby_on_rails/
---
Querida comunidad queria comentarles que hemos creado un server en Discord el cual tiene como propósito enseñar sobre programación.

Creo que la mejor forma de aprender es enseñar, por supuesto no somos formadores, quizá sabemos un poco de tecnología pero exactamente ese es el propósito de este server "Aprender" y en específico del ecosistema "Ruby y Ruby On Rails".

En el canal de voz \[Grupo de lectura\] tendremos un proyecto que trataremos de implementar por 1 hora diaria de lunes a viernes (8.00pm Colombia), donde crearemos un ERP, queremos que aprendas temas como planeación, metodologías ágiles, buenas prácticas, test, desarrollo, deploy, documentation, trabajo en remoto, cómo trabajar en equipo \[roles, responsabilidades, tiempos de entrega, etc\], la idea principal es aprender haciendo/enseñando, así que cualquier tipo de aporte que nos lleve a lograr ese objetivo es bienvenido. 

Intentaremos seguir un libro del profesor Javier Vazquez ([https://rclibros.es/producto/ruby-on-rails-aprende-a-crear-aplicaciones-web-desde-cero/](https://rclibros.es/producto/ruby-on-rails-aprende-a-crear-aplicaciones-web-desde-cero/)) que aborda el desarrollo de un software de este tipo y el cual encontrarás en el canal \[#libro\], es solo una guia como para estar en la misma página y que todos hablemos el mismo idioma, no es obligación que lo tengas para que puedas compartir en esta comunidad pero a mi te percepción será de mucha ayuda.

Desde ya te invitamos a que nos compartas que piensas de esta metodología y si es posible nos recomiendes material para ahondar mejor esa perspectiva de aprender haciendo/enseñando, siéntete libre de hacerlo en el canal \[#general\]!

&amp;#x200B;

Puedes unirte o invitar a tus amigos desde el siguiente link [https://discord.gg/P4wZGsQ](https://discord.gg/P4wZGsQ)
## [3][Required Module Data Instantiation](https://www.reddit.com/r/ruby/comments/i5bl7n/required_module_data_instantiation/)
- url: https://www.reddit.com/r/ruby/comments/i5bl7n/required_module_data_instantiation/
---
Hey everyone,  


I have been working on a gem that will help ease the creation of RPG-like games using Ruby. I was looking for some feedback on how I should be initializing required data for the modules to work appropriately.  


The repo is here if you are interested in perusing the code: [https://github.com/Rhoxio/adjective-rpg-engine](https://github.com/Rhoxio/adjective-rpg-engine) (Ignore the Readme - it's out of date right now.) The rdoc is here and is pretty much up to date: [https://rhoxio.github.io/adjective-rpg-engine/Adjective.html](https://rhoxio.github.io/adjective-rpg-engine/Adjective.html)  


An applicable case can be found here in a surrogate testing class: [https://github.com/Rhoxio/adjective-rpg-engine/blob/master/spec/surrogate\_classes/surrogate\_actor.rb](https://github.com/Rhoxio/adjective-rpg-engine/blob/master/spec/surrogate_classes/surrogate_actor.rb)  


Which uses a module like `Adjective::Temporable` found here: [https://github.com/Rhoxio/adjective-rpg-engine/blob/master/lib/adjective/concerns/temporable.rb](https://github.com/Rhoxio/adjective-rpg-engine/blob/master/lib/adjective/concerns/temporable.rb)  


As of right now, I have methods in the modules (`initialize_modulename(opts)`) that when called, sets instance variables and attr\_reader/accessor as I have seen fit. I assume they will change reader/writer/accessors from the default if they see the need arise - but I need to test it effectively and very sparingly use `attr_accessor` as it is due to the nature of the code and need the provided default setter/getters for the given attributes anyway.  


After doing some research on the topic, I have found patterns that use `define_method` and similar tricks, but when it comes down to it the default functionality *can* be overridden by simply defining a method of the same name. I am assuming (with great, and probably misplaced hope) that if they use the library, they will use the provided method names and and specifically included block and proc functionality if they wish to have it do something "different" than the default functionality provides.   


Am I missing something? It "feels" like having them call `initialize_modulename(opts)` in `initialize` from the parent class is an "extra step" or "poor implementation". I have looked into patterns, but many of them seem to ask that I marry together specific classes by name or role, and in this situation, seems to be completely out of context for the problem I am trying to solve. I still need them to be able to pass data down the chain easily and cleanly, essentially, with as little development overhead as possible.   


Thanks in advance! Also, any feedback about the code is appreciated. I've been coding in Ruby for about 5 years now, and I wanted the challenge of writing something more meta than my usual web dev stuff.
## [4][Lisamander and Robb clean up a Ruby on Rails codebase with Sorbet - LawIsCode on Twitch](https://www.reddit.com/r/ruby/comments/i4wk1g/lisamander_and_robb_clean_up_a_ruby_on_rails/)
- url: https://www.twitch.tv/videos/701837904
---

## [5][The Definitive Lemur Talk List](https://www.reddit.com/r/ruby/comments/i51r46/the_definitive_lemur_talk_list/)
- url: https://twitter.com/keystonelemur/status/1291488545450926086
---

## [6][What are some cool things you can do with Ruby parser.](https://www.reddit.com/r/ruby/comments/i4u9by/what_are_some_cool_things_you_can_do_with_ruby/)
- url: https://www.reddit.com/r/ruby/comments/i4u9by/what_are_some_cool_things_you_can_do_with_ruby/
---
I just learnt about what what parsers in ruby are and how they can rewrite source code. 
For someone who is trying to learn it's API's and AST do you can any cool suggestions that can be done with it.
Like I was thinking if trimming whitespaces from the end of the sentence could be one of them. 
Or enforcing only 1 space between = and the next character.
Any cool places ruby parsers are being used?
## [7][Creating and distributing a GUI app](https://www.reddit.com/r/ruby/comments/i4s771/creating_and_distributing_a_gui_app/)
- url: https://www.reddit.com/r/ruby/comments/i4s771/creating_and_distributing_a_gui_app/
---
Hello there, I have been searching the whole web for an answer but I was not able to. I am using Fedora as my operating system. I want to create a GUI app and distribute it as an executable file. Not as a gem or .rb files. I want to create cross platform executables. It would be perfect if you directed me to right places. I am looking for a way to develop GUI and distribute it.

&amp;#x200B;

\-Thanks
## [8][Record a Call in Ruby with Vonage Voice API WebSockets](https://www.reddit.com/r/ruby/comments/i4nt0z/record_a_call_in_ruby_with_vonage_voice_api/)
- url: https://www.nexmo.com/blog/2020/08/04/record-a-call-in-ruby-with-vonage-voice-api-websockets
---

## [9][Could you please help me with this JRuby question?](https://www.reddit.com/r/ruby/comments/i4qj12/could_you_please_help_me_with_this_jruby_question/)
- url: https://stackoverflow.com/q/63251312
---

## [10][Trouble With Deploying Rails App with Elastic Beanstalk](https://www.reddit.com/r/ruby/comments/i4nq6a/trouble_with_deploying_rails_app_with_elastic/)
- url: https://www.reddit.com/r/ruby/comments/i4nq6a/trouble_with_deploying_rails_app_with_elastic/
---
I've been trying to figure out how to deploy a simple rails app to elastic beanstalk using the cli on my macOS, but every time I get to `eb create` I always get the error at the bottom:


    2020/08/06 07:22:56.626563 [INFO] Executing instruction: StageApplication
    2020/08/06 07:22:56.626674 [INFO] extracting /opt/elasticbeanstalk/deployment/app_source_bundle to /var/app/staging/
    2020/08/06 07:22:56.626695 [INFO] Running command /bin/sh -c /usr/bin/unzip -q -o /opt/elasticbeanstalk/deployment/app_source_bundle -d /var/app/staging/
    2020/08/06 07:22:56.638657 [INFO] finished extracting /opt/elasticbeanstalk/deployment/app_source_bundle to /var/app/staging/ successfully
    2020/08/06 07:22:56.640331 [INFO] Executing instruction: RunAppDeployPreBuildHooks
    2020/08/06 07:22:56.640351 [INFO] The dir .platform/hooks/prebuild/ does not exist in the application. Skipping this step...
    2020/08/06 07:22:56.640356 [INFO] Executing instruction: stage ruby application
    2020/08/06 07:22:56.640360 [INFO] stage ruby application ....
    2020/08/06 07:22:56.640386 [INFO] Running command /bin/sh -c bundle config set --local deployment true
    2020/08/06 07:22:56.669550 [ERROR] An error occurred during execution of command [app-deploy] - [stage ruby application]. Stop running the command. Error: install dependencies in Gemfile failed with error Command /bin/sh -c bundle config set --local deployment true failed with error exit status 1. Stderr:rbenv: version `ruby-2.7.0' is not installed (set by /var/app/staging/.ruby-version)


After that I run `eb terminate` so that I don't incur any costs. But I don't understand the solution. Previously, the error said that 'ruby-2.6.3' is not installed' and I was unable to update to the latest version using rbenv so I removed it with homebrew and installed Ruby 2.7.0 using rvm, then tried to remove rails with `gem uninstall rails`. 

Then I tried installing rails again using `\curl -sSL https://get.rvm.io | bash -s stable --rails` which I got from the [RVM install guide](https://rvm.io/rvm/install), but then the above error happened. 

I imagine the answer would be to somehow update Ruby using rbenv? It was pretty difficult for me to figure that out, I believe since there is a version on Mac that is difficult to upgrade since the os uses it for legacy apps? Does anybody have a resource to help with that? I'm sure there's lots of flaws in my understanding of all this so any direction would be appreciated.
