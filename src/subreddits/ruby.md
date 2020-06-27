# ruby
## [1][Fun Facts about Ruby #6](https://www.reddit.com/r/ruby/comments/hgokpp/fun_facts_about_ruby_6/)
- url: https://i.redd.it/r9rmstmyde751.png
---

## [2][How to return the output of a If/Else in ruby](https://www.reddit.com/r/ruby/comments/hgqxln/how_to_return_the_output_of_a_ifelse_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/hgqxln/how_to_return_the_output_of_a_ifelse_in_ruby/
---
I have a method that returns one this if it is true and another thing if it is false.

    def has_perm?(item)
      return 
          if(true)
            item.this?
          else 
            item.that?
          end
    end

I wanted to write like this but it doesn't work. What is some clean way I can write code for this? The below looks fine but seems unnecessary assignment

    def has_perm?(item)
      result =  
          if(true)
            item.this?
          else 
            item.that?
          end
      result
    end
## [3][If you're using GraphQL with ruby these generators will save you many keystrokes](https://www.reddit.com/r/ruby/comments/hgb82l/if_youre_using_graphql_with_ruby_these_generators/)
- url: https://www.reddit.com/r/ruby/comments/hgb82l/if_youre_using_graphql_with_ruby_these_generators/
---
I released this library of graphql rails generators a few months ago after doing an ungodly amount of typing on a data-heavy application backed by a graphql API. The generators look at your Rails models and generate graphql types, mutations and input types from your database scheme (currently only activerecord supported, but other databases would be trivial to add). 

Hope people find some use from this and saves you a few thousand keystrokes. Cheers.

[https://github.com/ajsharp/graphql-rails-generators](https://github.com/ajsharp/graphql-rails-generators)
## [4][Add NEWS entries about JIT optimizations · ruby/ruby@200c5f4 · GitHub](https://www.reddit.com/r/ruby/comments/hg8787/add_news_entries_about_jit_optimizations/)
- url: https://github.com/ruby/ruby/commit/200c5f4075cb1d179c2eba5b30b5b0a500870f67
---

## [5][Acts As Tracked: Selectively track changes made on your ActiveRecord Models](https://www.reddit.com/r/ruby/comments/hg8evn/acts_as_tracked_selectively_track_changes_made_on/)
- url: https://www.reddit.com/r/ruby/comments/hg8evn/acts_as_tracked_selectively_track_changes_made_on/
---
Hi, i've made a gem to selectively track changes on AR models, where audited gem would be an overkill. ActsAsTracked can be plugged into ActiveRecord model, and then used whenever you need a history of changes and actors made on the record.

You can find docs in the [repository](https://github.com/ramblingcode/acts-as-tracked)

Blog post [here](https://www.ramblingcode.dev/posts/track_changes_on_your_activerecord_models/)

Example usage in this [project](https://github.com/ramblingcode/rails6-acts-as-tracked-usage)

Hope you find it useful.
## [6][Anonymous Struct literal `${a:1, b:2}` by ko1](https://www.reddit.com/r/ruby/comments/hg2bm1/anonymous_struct_literal_a1_b2_by_ko1/)
- url: https://github.com/ruby/ruby/pull/3259
---

## [7][Soo, I have this error on my kali nethunter its some ruby error how do I fix it any help?](https://www.reddit.com/r/ruby/comments/hgj59z/soo_i_have_this_error_on_my_kali_nethunter_its/)
- url: https://i.redd.it/4l259vl5gc751.png
---

## [8][Homework trouble](https://www.reddit.com/r/ruby/comments/hg9mns/homework_trouble/)
- url: https://www.reddit.com/r/ruby/comments/hg9mns/homework_trouble/
---
I've got some homework from my mentor who's helping me get back into Ruby after 3 years of not doing it.
He gave me this homework:

Fruit = Struct.new(:name, :color)

fruits = [
  Fruit.new("apple", "green"),
  Fruit.new("apricot", "orange")
  Fruit.new("banana", "yellow"),
  Fruit.new("grapes", "green"),
  Fruit.new("orange", "orange"),
  Fruit.new("lemon", "yellow"),
  Fruit.new("lime", "green"),
  Fruit.new("kiwi", "green")
]
```

How do I....

1. produce a list of all the green fruit
2. produce a list of all the fruits whose name starts with "a"
3. produce a list of all the fruits that aren't yellow
4. break the list down into a hash of smaller lists that are grouped by color?


And I am so lost already.
Can anyone help me with some hints?
## [9][Ruby can't access file it creates. ls -l can't view into that folder unless I pwd into it?](https://www.reddit.com/r/ruby/comments/hg0iu5/ruby_cant_access_file_it_creates_ls_l_cant_view/)
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
## [10][Rubyist teaches Elasticsearch in a fun way](https://www.reddit.com/r/ruby/comments/hfpr6p/rubyist_teaches_elasticsearch_in_a_fun_way/)
- url: https://realptsdengineer.com/learn-elasticsearch-fun-way/
---

