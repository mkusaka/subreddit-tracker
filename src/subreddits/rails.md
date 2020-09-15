# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/ipfxw8/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/ipfxw8/personal_projects_show_off_your_own_project_andor/
---
In this thread you can showcase your personal pet project to other redditors.

Need help with a specific problem or just wanna have some extra eyeballs on your code? Ask away!

A suggested format to get you started:

1. **Name of your project**
2. **A short description**
3. **Application stack**
4. **Link to Live app**
5. **Link to GitHub**
6. **You experience level**
7. **Other information or areas that you would like advice on**

 

^(Many thanks to Kritnc for getting the ball rolling.)
## [2][How to write this piece of code.](https://www.reddit.com/r/rails/comments/it7j0j/how_to_write_this_piece_of_code/)
- url: https://www.reddit.com/r/rails/comments/it7j0j/how_to_write_this_piece_of_code/
---
I have a requirement where I need to scope based on a few conditions. like so

    posts = case when type == 'deleted'
    posts.where(deleted: true)
    when type == 'block'
    posts.where(blocked: true)
    else 
    # return empty
    end

And then I'll loop through this like 

`posts.find_each(batch_size: 10) ...`

Now, how I do I return something to signify empty AR object list so that the find\_each fails and it doesn't get inside the loop

I was thinking I'd use  `nil`  and then user `posts&amp;.find_each`   


Is there a better way to do this?
## [3][Connecting to WordPress DB from rails](https://www.reddit.com/r/rails/comments/it350w/connecting_to_wordpress_db_from_rails/)
- url: https://www.reddit.com/r/rails/comments/it350w/connecting_to_wordpress_db_from_rails/
---
Hey all,

Our company has a DB or blog articles up in WordPress and a separate website made with RoR. Now, we need to be able to fetch the lists of articles but through the Rails back end and display it on the website.

Has anyone done this type of work before or have any solutions to this? Preferably one that is stable and up to date. Thanks!
## [4][need help with an array of objects](https://www.reddit.com/r/rails/comments/isvu94/need_help_with_an_array_of_objects/)
- url: https://www.reddit.com/r/rails/comments/isvu94/need_help_with_an_array_of_objects/
---
Hello all -

I have an array of objects (users). each object has an id, username, email, etc...
I'm trying to loop over the array and print ONLY the username (in rails console)

I've done quite a bit of googling, and can't get this to work - it always seems to print out the ENTIRE user object when I loop over the array.

what I want to do is something like:

user.each do |x|
puts x.username
end

I've tried x['username'], x[username], x.username, x[:username]. kind of at a loss here

Thanks in advance
## [5][your suggestion for a good rails API documentation](https://www.reddit.com/r/rails/comments/islkz9/your_suggestion_for_a_good_rails_api_documentation/)
- url: https://www.reddit.com/r/rails/comments/islkz9/your_suggestion_for_a_good_rails_api_documentation/
---
Hello, i started to learn both ruby and rails in the past two weeks, the problem is i'm really struggling while reading the rails API documentation, in their main site [https://api.rubyonrails.org/](https://api.rubyonrails.org/) i can't find a way to get all the inherited/included methods for a class, i was looking for the `render` method in `ActionController::API` and didn't find it so i went to `ActionController::Metal` and `ActionController::Base` and still didn't find it i really had to do a google search to find it, i'm just looking for a documentation to show me all the available methods you can use inside a class just same as the MDN documentation for JavaScript.

or if you could help me with the correct way to read the official documentation i would be really thankful
## [6][Error with blogger app](https://www.reddit.com/r/rails/comments/ism550/error_with_blogger_app/)
- url: https://www.reddit.com/r/rails/comments/ism550/error_with_blogger_app/
---
Hello, I am completely new to Rails and I am currently following a tutorial to create a blogger app ([http://tutorials.jumpstartlab.com/projects/blogger.html](http://tutorials.jumpstartlab.com/projects/blogger.html)) but I have run into a problem even though I believe that I did everything exactly as I was supposed to.

&amp;#x200B;

[The error](https://preview.redd.it/u4m1qf3ck4n51.png?width=1091&amp;format=png&amp;auto=webp&amp;s=70a41e1c025ddfe21ba4112a0aa8adc1d4f06127)

&amp;#x200B;

[My code](https://preview.redd.it/oghgpxapk4n51.png?width=1846&amp;format=png&amp;auto=webp&amp;s=c09dc49b9c4904c4cc60c1848d967cce44d9c0a0)

Could anyone help identifying where the problem is?
## [7][i made a ruby gem API wrapper for Google Pagespeed Insights](https://www.reddit.com/r/rails/comments/is782v/i_made_a_ruby_gem_api_wrapper_for_google/)
- url: https://www.reddit.com/r/rails/comments/is782v/i_made_a_ruby_gem_api_wrapper_for_google/
---
Hi, i posted the gem here because maybe could be useful for somebody someday. If you want, you can star and contribute to the project on github ;), Thanks you.

Repo: [https://github.com/kevyder/ruby\_pagespeed\_insights](https://github.com/kevyder/ruby_pagespeed_insights)

Rubygems: [https://rubygems.org/gems/ruby\_pagespeed\_insights](https://rubygems.org/gems/ruby_pagespeed_insights)
## [8][Video Tutorial: How to Host Two Different Rails Apps on One EC2 Instance](https://www.reddit.com/r/rails/comments/irz7l8/video_tutorial_how_to_host_two_different_rails/)
- url: https://www.reddit.com/r/rails/comments/irz7l8/video_tutorial_how_to_host_two_different_rails/
---
I just posted my latest AWS + Rails tutorial: 2 rails apps on the same server.  This was a particularly difficult topic for me when I started learning how to host apps on EC2.  

My hope is that I can help others figure this out much more quickly and with significantly less effort.  The configuration and techniques used in this tutorial tie several previous tutorials together.  I fully understand this could be accomplished with Docker, but I am attempting to lower the barrier to entry for beginners.  Using Docker felt like unnecessary complexity at this point, though I may cover the topic in the future.

If you are interested in learning, check it out and let me know what you think

[https://youtu.be/3U-uJzNEm1c](https://youtu.be/3U-uJzNEm1c)
## [9][Ruby on Rails Livestreams: Sundae Club](https://www.reddit.com/r/rails/comments/is1su4/ruby_on_rails_livestreams_sundae_club/)
- url: https://www.reddit.com/r/rails/comments/is1su4/ruby_on_rails_livestreams_sundae_club/
---
Hi everyone,

I posted some Ruby on Rails livestreaming on Reddit at the end of last year and got some positive feedback (thanks to everyone who got in touch). I'm now working a 9-5 and can't do ad-hoc streams, so I thought I'd do something a little more regular on Sunday mornings. I've been coaching at a some coding events and workshops throughout quarantine, so thought the streams may be a resource for people that are learning, but are looking for some more general questions to be answered, or to see how someone else approaches a task in Rails, rather than specific tutorials.

The channel is called Sundae Club. As well as each Sunday, I’m hoping to do some additional livestreams for the UI design stuff (Sketch/Figma) in any spare time I can find.

Here’s this week’s stream:

https://youtu.be/vQjw5uUAx5k

We covered adding our initial models and associations, we also added Tailwind as a CSS library.

You can find the first stream from last week [here](https://youtu.be/2_pUPKCVX28).
## [10][nokogiri installation error](https://www.reddit.com/r/rails/comments/is1w0a/nokogiri_installation_error/)
- url: https://www.reddit.com/r/rails/comments/is1w0a/nokogiri_installation_error/
---
Hello.

I am using ubuntu and when I use bundle install I get an error with nokogiri installation. How can I fix it?

gem\_make.out :

    current directory: /home/domke89/.rvm/gems/ruby-2.7.0/gems/nokogiri-1.10.10/ext/nokogiri
    /usr/share/rvm/rubies/ruby-2.7.0/bin/ruby -I /usr/share/rvm/rubies/ruby-2.7.0/lib/ruby/2.7.0 -r ./siteconf20200913-67510-1ka5qgs.rb extconf.rb --use-system-libraries
    checking if the C compiler accepts ... yes
    Building nokogiri using system libraries.
    checking for xmlParseDoc() in libxml/parser.h... yes
    checking for xsltParseStylesheetDoc() in libxslt/xslt.h... yes
    checking for exsltFuncRegister() in libexslt/exslt.h... yes
    checking for xmlHasFeature()... yes
    checking for xmlFirstElementChild()... yes
    checking for xmlRelaxNGSetParserStructuredErrors()... yes
    checking for xmlRelaxNGSetParserStructuredErrors()... yes
    checking for xmlRelaxNGSetValidStructuredErrors()... yes
    checking for xmlSchemaSetValidStructuredErrors()... yes
    checking for xmlSchemaSetParserStructuredErrors()... yes
    creating Makefile
    
    current directory: /home/domke89/.rvm/gems/ruby-2.7.0/gems/nokogiri-1.10.10/ext/nokogiri
    make "DESTDIR=" clean
    
    current directory: /home/domke89/.rvm/gems/ruby-2.7.0/gems/nokogiri-1.10.10/ext/nokogiri
    make "DESTDIR="
    compiling html_document.c
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./html_document.h:4,
                     from html_document.c:1:
    html_document.c: In function ‘new’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    html_document.c:16:3: note: in expansion of macro ‘rb_scan_args’
       16 |   rb_scan_args(argc, argv, "0*", &amp;rest);
          |   ^~~~~~~~~~~~
    html_document.c: At top level:
    cc1: warning: unrecognized command line option ‘-Wno-self-assign’
    cc1: warning: unrecognized command line option ‘-Wno-parentheses-equality’
    cc1: warning: unrecognized command line option ‘-Wno-constant-logical-operand’
    compiling html_element_description.c
    compiling html_entity_lookup.c
    compiling html_sax_parser_context.c
    compiling html_sax_push_parser.c
    compiling nokogiri.c
    compiling xml_attr.c
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_attr.h:4,
                     from xml_attr.c:1:
    xml_attr.c: In function ‘new’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    xml_attr.c:61:3: note: in expansion of macro ‘rb_scan_args’
       61 |   rb_scan_args(argc, argv, "2*", &amp;document, &amp;name, &amp;rest);
          |   ^~~~~~~~~~~~
    xml_attr.c: At top level:
    cc1: warning: unrecognized command line option ‘-Wno-self-assign’
    cc1: warning: unrecognized command line option ‘-Wno-parentheses-equality’
    cc1: warning: unrecognized command line option ‘-Wno-constant-logical-operand’
    compiling xml_attribute_decl.c
    compiling xml_cdata.c
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_cdata.h:4,
                     from xml_cdata.c:1:
    xml_cdata.c: In function ‘new’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    xml_cdata.c:23:3: note: in expansion of macro ‘rb_scan_args’
       23 |   rb_scan_args(argc, argv, "2*", &amp;doc, &amp;content, &amp;rest);
          |   ^~~~~~~~~~~~
    xml_cdata.c: At top level:
    cc1: warning: unrecognized command line option ‘-Wno-self-assign’
    cc1: warning: unrecognized command line option ‘-Wno-parentheses-equality’
    cc1: warning: unrecognized command line option ‘-Wno-constant-logical-operand’
    compiling xml_comment.c
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_comment.h:4,
                     from xml_comment.c:1:
    xml_comment.c: In function ‘new’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    xml_comment.c:21:3: note: in expansion of macro ‘rb_scan_args’
       21 |   rb_scan_args(argc, argv, "2*", &amp;document, &amp;content, &amp;rest);
          |   ^~~~~~~~~~~~
    xml_comment.c: At top level:
    cc1: warning: unrecognized command line option ‘-Wno-self-assign’
    cc1: warning: unrecognized command line option ‘-Wno-parentheses-equality’
    cc1: warning: unrecognized command line option ‘-Wno-constant-logical-operand’
    compiling xml_document.c
    xml_document.c: In function ‘dealloc’:
    xml_document.c:49:25: warning: passing argument 2 of ‘rb_st_foreach’ from incompatible pointer type [-Wincompatible-pointer-types]
       49 |   st_foreach(node_hash, dealloc_node_i, (st_data_t)doc);
          |                         ^~~~~~~~~~~~~~
          |                         |
          |                         int (*)(xmlNode *, xmlNode *, xmlDoc *) {aka int (*)(struct _xmlNode *, struct _xmlNode *, struct _xmlDoc *)}
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/intern.h:39,
                     from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2148,
                     from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_document.h:4,
                     from xml_document.c:1:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/st.h:141:31: note: expected ‘int (*)(st_data_t,  st_data_t,  st_data_t)’ {aka ‘int (*)(long unsigned int,  long unsigned int,  long unsigned int)’} but argument is of type ‘int (*)(xmlNode *, xmlNode *, xmlDoc *)’ {aka ‘int (*)(struct _xmlNode *, struct _xmlNode *, struct _xmlDoc *)’}
      141 | int rb_st_foreach(st_table *, st_foreach_callback_func *, st_data_t);
          |                               ^~~~~~~~~~~~~~~~~~~~~~~~~~
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_document.h:4,
                     from xml_document.c:1:
    xml_document.c: In function ‘duplicate_document’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    xml_document.c:330:6: note: in expansion of macro ‘rb_scan_args’
      330 |   if(rb_scan_args(argc, argv, "01", &amp;level) == 0)
          |      ^~~~~~~~~~~~
    xml_document.c: In function ‘new’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    xml_document.c:357:3: note: in expansion of macro ‘rb_scan_args’
      357 |   rb_scan_args(argc, argv, "0*", &amp;rest);
          |   ^~~~~~~~~~~~
    xml_document.c: In function ‘create_entity’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    xml_document.c:436:3: note: in expansion of macro ‘rb_scan_args’
      436 |   rb_scan_args(argc, argv, "14", &amp;name, &amp;type, &amp;external_id, &amp;system_id,
          |   ^~~~~~~~~~~~
    xml_document.c: In function ‘canonicalize’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    xml_document.c:511:3: note: in expansion of macro ‘rb_scan_args’
      511 |   rb_scan_args(argc, argv, "03", &amp;mode, &amp;incl_ns, &amp;with_comments);
          |   ^~~~~~~~~~~~
    xml_document.c: At top level:
    cc1: warning: unrecognized command line option ‘-Wno-self-assign’
    cc1: warning: unrecognized command line option ‘-Wno-parentheses-equality’
    cc1: warning: unrecognized command line option ‘-Wno-constant-logical-operand’
    compiling xml_document_fragment.c
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_document_fragment.h:4,
                     from xml_document_fragment.c:1:
    xml_document_fragment.c: In function ‘new’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    xml_document_fragment.c:17:3: note: in expansion of macro ‘rb_scan_args’
       17 |   rb_scan_args(argc, argv, "1*", &amp;document, &amp;rest);
          |   ^~~~~~~~~~~~
    xml_document_fragment.c: At top level:
    cc1: warning: unrecognized command line option ‘-Wno-self-assign’
    cc1: warning: unrecognized command line option ‘-Wno-parentheses-equality’
    cc1: warning: unrecognized command line option ‘-Wno-constant-logical-operand’
    compiling xml_dtd.c
    compiling xml_element_content.c
    compiling xml_element_decl.c
    compiling xml_encoding_handler.c
    compiling xml_entity_decl.c
    compiling xml_entity_reference.c
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_entity_reference.h:4,
                     from xml_entity_reference.c:1:
    xml_entity_reference.c: In function ‘new’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    xml_entity_reference.c:18:3: note: in expansion of macro ‘rb_scan_args’
       18 |   rb_scan_args(argc, argv, "2*", &amp;document, &amp;name, &amp;rest);
          |   ^~~~~~~~~~~~
    xml_entity_reference.c: At top level:
    cc1: warning: unrecognized command line option ‘-Wno-self-assign’
    cc1: warning: unrecognized command line option ‘-Wno-parentheses-equality’
    cc1: warning: unrecognized command line option ‘-Wno-constant-logical-operand’
    compiling xml_io.c
    xml_io.c: In function ‘io_read_callback’:
    xml_io.c:20:22: warning: passing argument 1 of ‘rb_rescue’ from incompatible pointer type [-Wincompatible-pointer-types]
       20 |   string = rb_rescue(read_check, (VALUE)args, read_failed, 0);
          |                      ^~~~~~~~~~
          |                      |
          |                      VALUE (*)(VALUE *) {aka long unsigned int (*)(long unsigned int *)}
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_io.h:4,
                     from xml_io.c:1:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:1987:17: note: expected ‘VALUE (*)(VALUE)’ {aka ‘long unsigned int (*)(long unsigned int)’} but argument is of type ‘VALUE (*)(VALUE *)’ {aka ‘long unsigned int (*)(long unsigned int *)’}
     1987 | VALUE rb_rescue(VALUE(*)(VALUE),VALUE,VALUE(*)(VALUE,VALUE),VALUE);
          |                 ^~~~~~~~~~~~~~~
    xml_io.c:20:47: warning: passing argument 3 of ‘rb_rescue’ from incompatible pointer type [-Wincompatible-pointer-types]
       20 |   string = rb_rescue(read_check, (VALUE)args, read_failed, 0);
          |                                               ^~~~~~~~~~~
          |                                               |
          |                                               VALUE (*)(void) {aka long unsigned int (*)(void)}
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_io.h:4,
                     from xml_io.c:1:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:1987:39: note: expected ‘VALUE (*)(VALUE,  VALUE)’ {aka ‘long unsigned int (*)(long unsigned int,  long unsigned int)’} but argument is of type ‘VALUE (*)(void)’ {aka ‘long unsigned int (*)(void)’}
     1987 | VALUE rb_rescue(VALUE(*)(VALUE),VALUE,VALUE(*)(VALUE,VALUE),VALUE);
          |                                       ^~~~~~~~~~~~~~~~~~~~~
    xml_io.c: In function ‘io_write_callback’:
    xml_io.c:47:20: warning: passing argument 1 of ‘rb_rescue’ from incompatible pointer type [-Wincompatible-pointer-types]
       47 |   size = rb_rescue(write_check, (VALUE)args, write_failed, 0);
          |                    ^~~~~~~~~~~
          |                    |
          |                    VALUE (*)(VALUE *) {aka long unsigned int (*)(long unsigned int *)}
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_io.h:4,
                     from xml_io.c:1:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:1987:17: note: expected ‘VALUE (*)(VALUE)’ {aka ‘long unsigned int (*)(long unsigned int)’} but argument is of type ‘VALUE (*)(VALUE *)’ {aka ‘long unsigned int (*)(long unsigned int *)’}
     1987 | VALUE rb_rescue(VALUE(*)(VALUE),VALUE,VALUE(*)(VALUE,VALUE),VALUE);
          |                 ^~~~~~~~~~~~~~~
    xml_io.c:47:46: warning: passing argument 3 of ‘rb_rescue’ from incompatible pointer type [-Wincompatible-pointer-types]
       47 |   size = rb_rescue(write_check, (VALUE)args, write_failed, 0);
          |                                              ^~~~~~~~~~~~
          |                                              |
          |                                              VALUE (*)(void) {aka long unsigned int (*)(void)}
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_io.h:4,
                     from xml_io.c:1:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:1987:39: note: expected ‘VALUE (*)(VALUE,  VALUE)’ {aka ‘long unsigned int (*)(long unsigned int,  long unsigned int)’} but argument is of type ‘VALUE (*)(void)’ {aka ‘long unsigned int (*)(void)’}
     1987 | VALUE rb_rescue(VALUE(*)(VALUE),VALUE,VALUE(*)(VALUE,VALUE),VALUE);
          |                                       ^~~~~~~~~~~~~~~~~~~~~
    xml_io.c: At top level:
    cc1: warning: unrecognized command line option ‘-Wno-self-assign’
    cc1: warning: unrecognized command line option ‘-Wno-parentheses-equality’
    cc1: warning: unrecognized command line option ‘-Wno-constant-logical-operand’
    compiling xml_libxml2_hacks.c
    compiling xml_namespace.c
    compiling xml_node.c
    xml_node.c: In function ‘reparent_node_with’:
    xml_node.c:304:29: warning: passing argument 1 of ‘xmlFree’ discards ‘const’ qualifier from pointer target type [-Wdiscarded-qualifiers]
      304 |       xmlFree(reparentee-&gt;ns-&gt;prefix);
          |               ~~~~~~~~~~~~~~^~~~~~~~
    xml_node.c:304:29: note: expected ‘void *’ but argument is of type ‘const xmlChar *’ {aka ‘const unsigned char *’}
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_node.h:4,
                     from xml_node.c:1:
    xml_node.c: In function ‘duplicate_node’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    xml_node.c:550:12: note: in expansion of macro ‘rb_scan_args’
      550 |   n_args = rb_scan_args(argc, argv, "02", &amp;r_level, &amp;r_new_parent_doc);
          |            ^~~~~~~~~~~~
    xml_node.c: In function ‘new’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    xml_node.c:1393:3: note: in expansion of macro ‘rb_scan_args’
     1393 |   rb_scan_args(argc, argv, "2*", &amp;name, &amp;document, &amp;rest);
          |   ^~~~~~~~~~~~
    xml_node.c: At top level:
    cc1: warning: unrecognized command line option ‘-Wno-self-assign’
    cc1: warning: unrecognized command line option ‘-Wno-parentheses-equality’
    cc1: warning: unrecognized command line option ‘-Wno-constant-logical-operand’
    compiling xml_node_set.c
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_node_set.h:4,
                     from xml_node_set.c:1:
    xml_node_set.c: In function ‘slice’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    xml_node_set.c:294:5: note: in expansion of macro ‘rb_scan_args’
      294 |     rb_scan_args(argc, argv, "11", NULL, NULL);
          |     ^~~~~~~~~~~~
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_node_set.h:4,
                     from xml_node_set.c:1:
    xml_node_set.c: In function ‘init_xml_node_set’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2799:117: warning: passing argument 3 of ‘rb_define_method0’ from incompatible pointer type [-Wincompatible-pointer-types]
     2799 | #define rb_define_method(klass, mid, func, arity) rb_define_method_choose_prototypem3((arity),(func))((klass),(mid),(func),(arity));
          |                                                                                                                     ^~~~~~
          |                                                                                                                     |
          |                                                                                                                     VALUE (*)(VALUE,  VALUE) {aka long unsigned int (*)(long unsigned int,  long unsigned int)}
    xml_node_set.c:479:3: note: in expansion of macro ‘rb_define_method’
      479 |   rb_define_method(klass, "to_a", to_array, 0);
          |   ^~~~~~~~~~~~~~~~
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2148,
                     from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_node_set.h:4,
                     from xml_node_set.c:1:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2775:27: note: expected ‘VALUE (*)(VALUE)’ {aka ‘long unsigned int (*)(long unsigned int)’} but argument is of type ‘VALUE (*)(VALUE,  VALUE)’ {aka ‘long unsigned int (*)(long unsigned int,  long unsigned int)’}
     2775 | RB_METHOD_DEFINITION_DECL(rb_define_method, (2,3), (VALUE klass, const char *name), (klass, name))
          |                           ^~~~~~~~~~~~~~~~
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/intern.h:1042:82: note: in definition of macro ‘RB_METHOD_DEFINITION_DECL_C’
     1042 |     __attribute__((__unused__,__weakref__(#def),__nonnull__ nonnull))static void defname(RB_UNWRAP_MACRO decl,VALUE(*func)funcargs,int arity);
          |                                                                                  ^~~~~~~
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/intern.h:1074:1: note: in expansion of macro ‘RB_METHOD_DEFINITION_DECL_1’
     1074 | RB_METHOD_DEFINITION_DECL_1(def,nonnull,def##0 ,0 ,decl,vars,(VALUE)) \
          | ^~~~~~~~~~~~~~~~~~~~~~~~~~~
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2775:1: note: in expansion of macro ‘RB_METHOD_DEFINITION_DECL’
     2775 | RB_METHOD_DEFINITION_DECL(rb_define_method, (2,3), (VALUE klass, const char *name), (klass, name))
          | ^~~~~~~~~~~~~~~~~~~~~~~~~
    xml_node_set.c: At top level:
    cc1: warning: unrecognized command line option ‘-Wno-self-assign’
    cc1: warning: unrecognized command line option ‘-Wno-parentheses-equality’
    cc1: warning: unrecognized command line option ‘-Wno-constant-logical-operand’
    compiling xml_processing_instruction.c
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_processing_instruction.h:4,
                     from xml_processing_instruction.c:1:
    xml_processing_instruction.c: In function ‘new’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    xml_processing_instruction.c:20:3: note: in expansion of macro ‘rb_scan_args’
       20 |   rb_scan_args(argc, argv, "3*", &amp;document, &amp;name, &amp;content, &amp;rest);
          |   ^~~~~~~~~~~~
    xml_processing_instruction.c: At top level:
    cc1: warning: unrecognized command line option ‘-Wno-self-assign’
    cc1: warning: unrecognized command line option ‘-Wno-parentheses-equality’
    cc1: warning: unrecognized command line option ‘-Wno-constant-logical-operand’
    compiling xml_reader.c
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_reader.h:4,
                     from xml_reader.c:1:
    xml_reader.c: In function ‘from_memory’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    xml_reader.c:533:3: note: in expansion of macro ‘rb_scan_args’
      533 |   rb_scan_args(argc, argv, "13", &amp;rb_buffer, &amp;rb_url, &amp;encoding, &amp;rb_options);
          |   ^~~~~~~~~~~~
    xml_reader.c: In function ‘from_io’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    xml_reader.c:577:3: note: in expansion of macro ‘rb_scan_args’
      577 |   rb_scan_args(argc, argv, "13", &amp;rb_io, &amp;rb_url, &amp;encoding, &amp;rb_options);
          |   ^~~~~~~~~~~~
    xml_reader.c: At top level:
    cc1: warning: unrecognized command line option ‘-Wno-self-assign’
    cc1: warning: unrecognized command line option ‘-Wno-parentheses-equality’
    cc1: warning: unrecognized command line option ‘-Wno-constant-logical-operand’
    compiling xml_relax_ng.c
    compiling xml_sax_parser.c
    compiling xml_sax_parser_context.c
    compiling xml_sax_push_parser.c
    compiling xml_schema.c
    compiling xml_syntax_error.c
    compiling xml_text.c
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_text.h:4,
                     from xml_text.c:1:
    xml_text.c: In function ‘new’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    xml_text.c:18:3: note: in expansion of macro ‘rb_scan_args’
       18 |   rb_scan_args(argc, argv, "2*", &amp;string, &amp;document, &amp;rest);
          |   ^~~~~~~~~~~~
    xml_text.c: At top level:
    cc1: warning: unrecognized command line option ‘-Wno-self-assign’
    cc1: warning: unrecognized command line option ‘-Wno-parentheses-equality’
    cc1: warning: unrecognized command line option ‘-Wno-constant-logical-operand’
    compiling xml_xpath_context.c
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xml_xpath_context.h:4,
                     from xml_xpath_context.c:1:
    xml_xpath_context.c: In function ‘evaluate’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    xml_xpath_context.c:200:6: note: in expansion of macro ‘rb_scan_args’
      200 |   if(rb_scan_args(argc, argv, "11", &amp;search_path, &amp;xpath_handler) == 1)
          |      ^~~~~~~~~~~~
    xml_xpath_context.c: At top level:
    cc1: warning: unrecognized command line option ‘-Wno-self-assign’
    cc1: warning: unrecognized command line option ‘-Wno-parentheses-equality’
    cc1: warning: unrecognized command line option ‘-Wno-constant-logical-operand’
    compiling xslt_stylesheet.c
    In file included from /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby.h:33,
                     from ./nokogiri.h:33,
                     from ./xslt_stylesheet.h:4,
                     from xslt_stylesheet.c:1:
    xslt_stylesheet.c: In function ‘transform’:
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2508:48: warning: cast discards ‘const’ qualifier from pointer target type [-Wcast-qual]
     2508 |        (rb_scan_args_verify(fmt, varc), vars), (char *)fmt, varc)
          |                                                ^
    /usr/share/rvm/rubies/ruby-2.7.0/include/ruby-2.7.0/ruby/ruby.h:2340:9: note: in expansion of macro ‘rb_scan_args0’
     2340 |         rb_scan_args0(argc,argvp,fmt,\
          |         ^~~~~~~~~~~~~
    xslt_stylesheet.c:141:5: note: in expansion of macro ‘rb_scan_args’
      141 |     rb_scan_args(argc, argv, "11", &amp;xmldoc, &amp;paramobj);
          |     ^~~~~~~~~~~~
    At top level:
    xslt_stylesheet.c:112:13: warning: ‘swallow_superfluous_xml_errors’ defined but not used [-Wunused-function]
      112 | static void swallow_superfluous_xml_errors(void * userdata, xmlErrorPtr error, ...)
          |             ^~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    cc1: warning: unrecognized command line option ‘-Wno-self-assign’
    cc1: warning: unrecognized command line option ‘-Wno-parentheses-equality’
    cc1: warning: unrecognized command line option ‘-Wno-constant-logical-operand’
    linking shared-object nokogiri/nokogiri.so
    Cleaning files only used during build.
    rm -rf /home/domke89/.rvm/gems/ruby-2.7.0/gems/nokogiri-1.10.10/ports/archives
    
    current directory: /home/domke89/.rvm/gems/ruby-2.7.0/gems/nokogiri-1.10.10/ext/nokogiri
    make "DESTDIR=" install
    make: /usr/bin/mkdir: Command not found
    make: *** [Makefile:202: .sitearchdir.-.nokogiri.time] Error 127
    
    make install failed, exit code 2
## [11][Digital signage using rails framework?](https://www.reddit.com/r/rails/comments/is1bjq/digital_signage_using_rails_framework/)
- url: https://www.reddit.com/r/rails/comments/is1bjq/digital_signage_using_rails_framework/
---
So, i was looking at this: [https://github.com/wassgha/digital-signage](https://github.com/wassgha/digital-signage)

Does anyone know if there's a rails based digital signage thing?

I think it should be possible to create this in Rails, would it be hard?
