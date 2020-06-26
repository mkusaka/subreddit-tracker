# ruby
## [1][Anonymous Struct literal `${a:1, b:2}` by ko1](https://www.reddit.com/r/ruby/comments/hg2bm1/anonymous_struct_literal_a1_b2_by_ko1/)
- url: https://github.com/ruby/ruby/pull/3259
---

## [2][Ruby can't access file it creates. ls -l can't view into that folder unless I pwd into it?](https://www.reddit.com/r/ruby/comments/hg0iu5/ruby_cant_access_file_it_creates_ls_l_cant_view/)
- url: https://www.reddit.com/r/ruby/comments/hg0iu5/ruby_cant_access_file_it_creates_ls_l_cant_view/
---
I've got a web application I am troubleshooting that' I'm not sure why it's failing. below is a segment of the code where it fails trying to open the file with:

&gt; No such file or directory @ rb_sysopen - /data/domains/default/files/users/205/20200625_DiF4HR63uq9X9JEL783g5v/asdf

      params[:message][:attachments].to_a.each do |attachment|
            @tempattachment = current_user.attachments.find(attachment)
            tempattachment_path = "/data/domains/default/files/users/#{current_user.id}/#{@tempattachment.updated_at.to_formatted_s(:number)[0..7]}_#{@tempattachment.magicid}/"
            tempattachment_path_before = File.expand_path(tempattachment_path + @tempattachment.filename.tr(' \'()@#+=^%$&amp;;,!{}[]~`','_'))
            tempopenedfile = File.open(tempattachment_path_before)

however the file (asdf, no extension) is there, was created by the _sfta user account running the ruby web app and the permissions show it should be able to access it. As root I also can't directly access that path as a whole but if I CD into the parent then ls it works (even as root):

    []# ls -la /data/domains/default/files/users/205/20200625_DiF4HR63uq9X9JEL783g5v/asdf
    ls: cannot access /data/domains/default/files/users/205/20200625_DiF4HR63uq9X9JEL783g5v/asdf: No such file or directory
    []# ls -la /data/domains/default/files/users/205/20200625_DiF4HR63uq9X9JEL783g5v
    ls: cannot access /data/domains/default/files/users/205/20200625_DiF4HR63uq9X9JEL783g5v: No such file or directory
    []# ls -la /data/domains/default/files/users/205/20200625_DiF4HR63uq9X9JEL783g5v/
    ls: cannot access /data/domains/default/files/users/205/20200625_DiF4HR63uq9X9JEL783g5v/: No such file or directory
    []# ls -la /data/domains/default/files/users/205
    total 4
    drwxr-xr-x.   3 _sfta _sfta   44 Jun 25 22:06 .
    drwxrwx---. 137 _sfta _sfta 4096 Jun 25 21:57 ..
    drwxrwx---.   2 _sfta _sfta   17 Jun 25 22:06 20200626_DiF4HR63uq9X9JEL783g5v
    []# ls -la /data/domains/default/files/users/205/20200626_DiF4HR63uq9X9JEL783g5v/
    total 4
    drwxrwx---. 2 _sfta _sfta  17 Jun 25 22:06 .
    drwxr-xr-x. 3 _sfta _sfta  44 Jun 25 22:06 ..
    -rw-rw----. 1 _sfta _sfta 833 Jun 25 22:06 asdf
    []#

sudo su _sfta then running:

    bash-4.2$ ls -la /data/domains/default/files/users/205/20200625_DiF4HR63uq9X9JEL783g5v
    total 4
    drwxrwx---. 2 _sfta _sfta  17 Jun 25 22:06 .
    drwxr-xr-x. 3 _sfta _sfta  44 Jun 25 22:06 ..
    -rw-rw----. 1 _sfta _sfta 833 Jun 25 22:06 asdf

What am I missing with this File.Open that it can't access the file?
## [3][Rubyist teaches Elasticsearch in a fun way](https://www.reddit.com/r/ruby/comments/hfpr6p/rubyist_teaches_elasticsearch_in_a_fun_way/)
- url: https://realptsdengineer.com/learn-elasticsearch-fun-way/
---

## [4][HEY’s Gemfile](https://www.reddit.com/r/ruby/comments/hfhrcb/heys_gemfile/)
- url: https://gist.github.com/dhh/782fb925b57450da28c1e15656779556
---

## [5][Ruby / gRPC Skeleton. I tried to make an ideal jump off point for Ruby micro-services with docker-compose, github-actions, hanami-model, and postgres. I'd love your feedback for improvements!](https://www.reddit.com/r/ruby/comments/hfpyfg/ruby_grpc_skeleton_i_tried_to_make_an_ideal_jump/)
- url: https://github.com/dudo/ruby_grpc_skeleton
---

## [6][til-rb, a gem to help you write and manage a TIL repo](https://www.reddit.com/r/ruby/comments/hflzrs/tilrb_a_gem_to_help_you_write_and_manage_a_til/)
- url: https://www.reddit.com/r/ruby/comments/hflzrs/tilrb_a_gem_to_help_you_write_and_manage_a_til/
---
Hey there 👋

&amp;#x200B;

I recently published an early version of [a small gem](https://github.com/pjambet/til-rb/) I've been working on to manage a "TIL repo", like [this one](https://github.com/jbranchaud/til) that made the rounds recently. It essentially creates a new commit for you, using the GH API, with a new file containing the content of the new TIL, and maintains a table of content in the [README.md](https://README.md) file.

&amp;#x200B;

Building it was really interesting and [I wrote a post about it](https://blog.pjam.me/posts/til-cli/), thought it would be worth sharing here!

&amp;#x200B;

I've been using it for a week or so, and while I'm sure it has a lot of bugs (almost total lack of validation on user input for instance), it has worked reasonably well so far.

&amp;#x200B;

Let me know if you have questions or feedback.
## [7][Exception tracking: Airbrake vs Honeybadger vs Sentry vs AppSignal](https://www.reddit.com/r/ruby/comments/hfk47b/exception_tracking_airbrake_vs_honeybadger_vs/)
- url: https://www.reddit.com/r/ruby/comments/hfk47b/exception_tracking_airbrake_vs_honeybadger_vs/
---
Which one do you use / have recently decided for? And why?

What do you hate about the current ones that you (have to) use?

Did I miss an important one for the ruby ecosystem?
## [8][Moving from EventMachine to Async](https://www.reddit.com/r/ruby/comments/hfgsdf/moving_from_eventmachine_to_async/)
- url: https://blog.joshsoftware.com/2020/06/19/moving-from-eventmachine-to-async/
---

## [9][Lingohub - removing unused translation keys](https://www.reddit.com/r/ruby/comments/hfk697/lingohub_removing_unused_translation_keys/)
- url: https://www.reddit.com/r/ruby/comments/hfk697/lingohub_removing_unused_translation_keys/
---
Is there any way to remove unused keys on Lingohub?

I'm just deleting a bunch of unused keys in the translation yml file. Is there any good and safe way to do this?

edit:

Answer from Lingohub team:

&gt;Hi xyz,LingoHub will automatically deactivate keys when they are removed from a file. So for example you had a file with the following keys:A B CAnd you then upload the same file with the keysA CThe B key would be deactivated and not shown in the editor or exported anymore.Please note, this behavior is changeable through project settings. So make sure you check them before.Hope that helps.Best,

This is a partial answer that I wanted. Can I somehow reverse these keys if the mistake is created?or create more environments?

stackoverflow q: [https://stackoverflow.com/questions/62556336/lingohub-removing-unused-translation-keys](https://stackoverflow.com/questions/62556336/lingohub-removing-unused-translation-keys)

  
Lingohub service for translates: [https://lingohub.com/](https://lingohub.com/)
## [10][Introducing UndercoverCI - a GitHub code review app that finds untested methods you've changed](https://www.reddit.com/r/ruby/comments/hf8070/introducing_undercoverci_a_github_code_review_app/)
- url: https://medium.com/@mrgrodo/introducing-undercoverci-98c6c50793f2
---

