# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/fx6je4/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/fx6je4/personal_projects_show_off_your_own_project_andor/
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
## [3][Rails 6 + docker-compose : simplest possible Hello World project™](https://www.reddit.com/r/rails/comments/g2c48p/rails_6_dockercompose_simplest_possible_hello/)
- url: https://www.reddit.com/r/rails/comments/g2c48p/rails_6_dockercompose_simplest_possible_hello/
---
Hello mates, I spend some time to find the simplest possible configuration to start a Rails project using Docker.

[https://github.com/bdavidxyz/simplest-rails-docker](https://github.com/bdavidxyz/simplest-rails-docker)

I picked some ideas from other projects from GitHub, trying to remove as much files and as much lines per file as possible.

With Ruby-on-Rails, you can't start a project without a properly initialized, configured, and started database. You also can't start the local server without a webpack dev server running (from Rails 6 I guess). So that's make 3 services running at least before even to print a "hello world" page.

This kind of project is extremely useful to test various tools and versions without polluting your local machine. I find it also more professional to use Docker instead of brittle, quickly-too-old documentation. So it could be a starting point for a more serious configuration.

What do you think ? What could be simplified ? Do you have more simple examples ? How do you test tools in isolation ? Thanks for your returns, and stay safe :)
## [4][From Agency Life to Software Development: Q&amp;A with Steve Polito](https://www.reddit.com/r/rails/comments/g25cia/from_agency_life_to_software_development_qa_with/)
- url: https://www.reddit.com/r/rails/comments/g25cia/from_agency_life_to_software_development_qa_with/
---
The Remote Ruby crew is joined by Steve Polito, a developer working at an agency, to have a Q&amp;A about transitioning into being a full-time software developer working on products. If you’re a junior developer or in a similar situation, [give this episode of Remote Ruby a listen!](https://remoteruby.transistor.fm/74)
## [5][GraphQL: making an order management system, trying to add many to many](https://www.reddit.com/r/rails/comments/g2cfrk/graphql_making_an_order_management_system_trying/)
- url: https://www.reddit.com/r/rails/comments/g2cfrk/graphql_making_an_order_management_system_trying/
---
Oh hi, it's me again! :)

Still trying to do a GraphQL api with rails, but I've hit a snag. I am doing an order management system, and it's not going to be worth much unless I can add products to an order. I have set them up with a Product model and an Order model with a has\_and\_belongs\_to\_many relationship. ie. Product has\_many orders and Order has\_many products. that part seems to be working fine, I made a join table and as far as I can see it works.

The problem arrises when I am using GraphQL. The documentation is hard to find once you move away from the most basic queries and mutations. alas, I am in the process for doing the CreateOrder mutation for my api, as I need to add products on creation, i.e whatever the user orders.

&amp;#x200B;

https://preview.redd.it/yyt09xqjo5t41.png?width=1240&amp;format=png&amp;auto=webp&amp;s=a12f18a5b1f13e8642952ba88aef2d7a04930caf

first I had the ProductType for an argument, but I realized that's a non starter and found out I need to create an input type for it, so I did:

&amp;#x200B;

https://preview.redd.it/l4shn61uo5t41.png?width=1148&amp;format=png&amp;auto=webp&amp;s=65a980bee0b3c47d68aa12e49a7fb103592c0c87

At this point this, I don't know if this is even the correct way of doing it, because I wasn't able to find a Ruby example online. Either way, GraphQL isn't complaining at least. So I try calling the mutation like so:

&amp;#x200B;

https://preview.redd.it/svajfnw6p5t41.png?width=438&amp;format=png&amp;auto=webp&amp;s=27b074ad895e58763f2c990ed948ca579dedb00b

and it is at this point that everything falls apart as rails gives me the error:

    Product(#375140) expected, got #&lt;Types::ProductInputType:0x00007fb1a791b8c8 u/context=#&lt;Query::Context ...&gt;, u/ruby_style_hash={:id=&gt;1}, u/arguments_by_keyword={:id=&gt;#&lt;Types::BaseArgument:0x00007fb1aa96ec00 u/name="id", u/type_expr=Integer, u/description=nil, u/null=false, u/default_value=:__no_default__, u/owner=Types::ProductInputType, u/as=nil, u/loads=nil, u/keyword=:id, u/prepare=nil, u/ast_node=nil, u/from_resolver=false, u/method_access=true, u/type=#&lt;GraphQL::Schema::NonNull u/of_type=GraphQL::Types::Int&gt;&gt;, :name=&gt;#&lt;Types::BaseArgument:0x00007fb1aa96d3c8 u/name="name", u/type_expr=String, u/description=nil, u/null=true, u/default_value=:__no_default__, u/owner=Types::ProductInputType, u/as=nil, u/loads=nil, u/keyword=:name, u/prepare=nil, u/ast_node=nil, u/from_resolver=false, u/method_access=true, u/type=GraphQL::Types::String&gt;, :organization_id=&gt;#&lt;Types::BaseArgument:0x00007fb1aa96c6f8 u/name="organizationId", u/type_expr=Integer, u/description=nil, u/null=true, u/default_value=:__no_default__, u/owner=Types::ProductInputType, u/as=nil, u/loads=nil, u/keyword=:organization_id, u/prepare=nil, u/ast_node=nil, u/from_resolver=false, u/method_access=true, u/type=GraphQL::Types::Int&gt;, :description=&gt;#&lt;Types::BaseArgument:0x00007fb1a6c734b8 u/name="description", u/type_expr=String, u/description=nil, u/null=true, u/default_value=:__no_default__, u/owner=Types::ProductInputType, u/as=nil, u/loads=nil, u/keyword=:description, u/prepare=nil, u/ast_node=nil, u/from_resolver=false, u/method_access=true, u/type=GraphQL::Types::String&gt;, :inventory=&gt;#&lt;Types::BaseArgument:0x00007fb1a6c72400 u/name="inventory", u/type_expr=Integer, u/description=nil, u/null=true, u/default_value=:__no_default__, u/owner=Types::ProductInputType, u/as=nil, u/loads=nil, u/keyword=:inventory, u/prepare=nil, u/ast_node=nil, u/from_resolver=false, u/method_access=true, u/type=GraphQL::Types::Int&gt;, :images=&gt;#&lt;Types::BaseArgument:0x00007fb1a6c712f8 u/name="images", u/type_expr=String, u/description=nil, u/null=true, u/default_value=:__no_default__, u/owner=Types::ProductInputType, u/as=nil, u/loads=nil, u/keyword=:images, u/prepare=nil, u/ast_node=nil, u/from_resolver=false, u/method_access=true, u/type=GraphQL::Types::String&gt;, :price=&gt;#&lt;Types::BaseArgument:0x00007fb1a6c70128 u/name="price", u/type_expr=Float, u/description=nil, u/null=true, u/default_value=:__no_default__, u/owner=Types::ProductInputType, u/as=nil, u/loads=nil, u/keyword=:price, u/prepare=nil, u/ast_node=nil, u/from_resolver=false, u/method_access=true, u/type=GraphQL::Types::Float&gt;}&gt; which is an instance of Types::ProductInputType(#372160)

&amp;#x200B;

it would be nice if someone in this awesome community could give me a pointer in the right direction, either by telling me what's wrong here, or by sharing a tutorial that covers something similar, because I haven't been able to find one on my own :/

&amp;#x200B;

Best Regards
## [6][How do you break out of a for loop that has nested? My attributes?](https://www.reddit.com/r/rails/comments/g2c38b/how_do_you_break_out_of_a_for_loop_that_has/)
- url: https://www.reddit.com/r/rails/comments/g2c38b/how_do_you_break_out_of_a_for_loop_that_has/
---
My code looks like this

&lt;% for @leads in @leads.nested_lead %&gt;
      &lt;%= @leads.contact %&gt;
&lt;% end %&gt;

---------- This is where my error comes in - - - - - - - -
When I do
&lt;%= @leads.count %&gt;

I'm getting an error for undefined method for .count
Or any over method I call that exists in the @leads table. 🤔
## [7][Event Members](https://www.reddit.com/r/rails/comments/g27vbr/event_members/)
- url: https://www.reddit.com/r/rails/comments/g27vbr/event_members/
---
Hey all,   


My Events page is up and running, I'm looking to add member to the event now.

I would like to have a group created when the event is created allocating the event creator as a group member, group admin(bool) and link the event to the group.

  
The event has many users and should be indicated via a event group.

Looking for the best/easiest way to implement this.  


All idea's/comments welcome.  


Thanks
## [8][Are there any *recent* Doorkeeper+Devise tutorials to make a Oauth2 server provider?](https://www.reddit.com/r/rails/comments/g23ewp/are_there_any_recent_doorkeeperdevise_tutorials/)
- url: https://www.reddit.com/r/rails/comments/g23ewp/are_there_any_recent_doorkeeperdevise_tutorials/
---
I've been struggling a lot to put together an Oauth2 provider. The tutorials I've found are either from 2015-2016 (with code that needs adjustments), or make some substantial assumptions (such as building on top of previous code from a previous guide).

I tried to follow these guides:

https://dev.mikamai.com/2015/02/11/oauth2-on-rails/  
https://dev.mikamai.com/2015/03/02/oauth2-on-rails-the-client-application/

But I get some errors about missing parameters. As I am totally new to this topic I've no idea what exactly I'm supposed to look for, or what to change.

My intention is to make an oauth2 provider for my Rails and Node applications.

Any tutorials you may know? Or perhaps somewhere I could investigate further?
## [9][Using Optimizer Hints in Rails](https://www.reddit.com/r/rails/comments/g1xxys/using_optimizer_hints_in_rails/)
- url: https://blog.saeloun.com/2020/04/14/using-optimizer-hints-in-rails-6
---

## [10][Dealing with uniqueness constraint updates using PostgresQL](https://www.reddit.com/r/rails/comments/g255cs/dealing_with_uniqueness_constraint_updates_using/)
- url: https://www.reddit.com/r/rails/comments/g255cs/dealing_with_uniqueness_constraint_updates_using/
---
Hey guys. I have a few , uuid, string-based columns that have uniqueness constraints. Sometimes the uuid needs to be removed on update. Since there are objects that do not yet have a uuid (i.e. an empty string). I get hit with a database error. Is there any way to keep the uniqueness uniqueness constraint with this in mind?
## [11][Technical interview](https://www.reddit.com/r/rails/comments/g1w9op/technical_interview/)
- url: https://www.reddit.com/r/rails/comments/g1w9op/technical_interview/
---
I m thinking to apply volunteer work but not sure if my technical skills in programming is great. What should I do? How to prepare for technical interview. I can’t solve the question but I can understand the solution after looking at it. Any suggestions in this regards will be great.
## [12][Deploy to Heroku: assets:precompile error on Heroku but can't reproduce locally](https://www.reddit.com/r/rails/comments/g1xtqy/deploy_to_heroku_assetsprecompile_error_on_heroku/)
- url: https://www.reddit.com/r/rails/comments/g1xtqy/deploy_to_heroku_assetsprecompile_error_on_heroku/
---
Hi,

I've just finished this feature but I am getting an assets:precompile error that I can't figure out.  On Heroku, the error is:

`----&gt; Preparing app for Rails asset pipeline`

`Running: rake assets:precompile`

`rake aborted!`

`\`NoMethodError: undefined method []' for nil:NilClass\`\``

`\`/tmp/build_26e19a2ff8e992cdff1a9b487057f4a3/config/initializers/shrine.rb:4:in &lt;main&gt;'\`\``

&amp;#x200B;

I tried to reproduce locally with RAILS\_ENV=production rake assets:precompile --trace.Maybe one of you can notice something:

`** Invoke assets:precompile (first_time)`

`** Invoke assets:environment (first_time)`

`** Execute assets:environment`

`** Invoke environment (first_time)`

`** Execute environment`

`** Invoke yarn:install (first_time)`

`** Invoke webpacker:yarn_install (first_time)`

`** Execute webpacker:yarn_install`

`yarn install v1.21.1`

`[1/4] Resolving packages...`

`success Already up-to-date.`

`Done in 0.97s.`

`** Execute yarn:install`

`yarn install v1.21.1`

`[1/4] Resolving packages...`

`[2/4] Fetching packages...`

`info fsevents@1.2.12: The platform "linux" is incompatible with this module.`

`info "fsevents@1.2.12" is an optional dependency and failed compatibility check. Excluding it from installation.`

`[3/4] Linking dependencies...`

`warning " &gt; webpack-dev-server@3.10.3" has unmet peer dependency "webpack@^4.0.0 || ^5.0.0".`

`warning "webpack-dev-server &gt; webpack-dev-middleware@3.7.2" has unmet peer dependency "webpack@^4.0.0".`

`[4/4] Building fresh packages...`

`Done in 252.40s.`

`** Execute assets:precompile`

`I, [2020-04-15T13:23:44.967735 #5912] INFO -- : Writing /vagrant/Rails/ink/public/assets/manifest-cadda289ef9c70eaa0879a36e6263cb33f7523a16b3ef862e0b8609cdc2bdab1.js`

`I, [2020-04-15T13:23:44.969095 #5912] INFO -- : Writing /vagrant/Rails/ink/public/assets/manifest-cadda289ef9c70eaa0879a36e6263cb33f7523a16b3ef862e0b8609cdc2bdab1.js.gz`

`I, [2020-04-15T13:23:44.969485 #5912] INFO -- : Writing /vagrant/Rails/ink/public/assets/application-d0ff5974b6aa52cf562bea5921840c032a860a91a3512f7fe8f768f6bbe005f6.css`

`I, [2020-04-15T13:23:44.969644 #5912] INFO -- : Writing /vagrant/Rails/ink/public/assets/application-d0ff5974b6aa52cf562bea5921840c032a860a91a3512f7fe8f768f6bbe005f6.css.gz`

`I, [2020-04-15T13:23:44.969801 #5912] INFO -- : Writing /vagrant/Rails/ink/public/assets/notes-d0ff5974b6aa52cf562bea5921840c032a860a91a3512f7fe8f768f6bbe005f6.css`

`I, [2020-04-15T13:23:44.969933 #5912] INFO -- : Writing /vagrant/Rails/ink/public/assets/notes-d0ff5974b6aa52cf562bea5921840c032a860a91a3512f7fe8f768f6bbe005f6.css.gz`

`I, [2020-04-15T13:23:44.970108 #5912] INFO -- : Writing /vagrant/Rails/ink/public/assets/scaffolds-d06093ee516fe24cb3fe95eb72cd21d1cddaec3bfa31d359826f6e6cafd43bd2.css`

`I, [2020-04-15T13:23:44.970236 #5912] INFO -- : Writing /vagrant/Rails/ink/public/assets/scaffolds-d06093ee516fe24cb3fe95eb72cd21d1cddaec3bfa31d359826f6e6cafd43bd2.css.gz`

`** Invoke webpacker:compile (first_time)`

`** Invoke webpacker:verify_install (first_time)`

`** Invoke webpacker:check_node (first_time)`

`** Execute webpacker:check_node`

`** Invoke webpacker:check_yarn (first_time)`

`** Execute webpacker:check_yarn`

`** Invoke webpacker:check_binstubs (first_time)`

`** Execute webpacker:check_binstubs`

`** Execute webpacker:verify_install`

`** Invoke environment`

`** Execute webpacker:compile`

`Compiling...`

`Compiled all packs in /vagrant/Rails/ink/public/packs`

`Hash: 26849029fbedbedced37`

`Version: webpack 4.42.1`

`Time: 24098ms`

`Built at: 2020-04-15 13:24:14`

`Asset Size Chunks Chunk Names`

`css/application-7891a208.css 826 KiB 0, 1 [emitted] [immutable] [big] application`

`css/application-7891a208.css.gz 106 KiB [emitted]`

`js/application-7da76fec4ff71205bb73.js 157 KiB 0, 1 [emitted] [immutable] application`

`js/application-7da76fec4ff71205bb73.js.LICENSE.txt 489 bytes [emitted]`

`js/application-7da76fec4ff71205bb73.js.gz 47.6 KiB [emitted]`

`js/application-7da76fec4ff71205bb73.js.map 657 KiB 0, 1 [emitted] [dev] application`

`js/application-7da76fec4ff71205bb73.js.map.gz 186 KiB [emitted]`

`js/notes-aa3d2eba82abbbda5606.js 89.4 KiB 1 [emitted] [immutable] notes`

`js/notes-aa3d2eba82abbbda5606.js.LICENSE.txt 489 bytes [emitted]`

`js/notes-aa3d2eba82abbbda5606.js.gz 31.1 KiB [emitted]`

`js/notes-aa3d2eba82abbbda5606.js.map 457 KiB 1 [emitted] [dev] notes`

`js/notes-aa3d2eba82abbbda5606.js.map.gz 139 KiB [emitted]`

`manifest.json 791 bytes [emitted]`

`manifest.json.gz 220 bytes [emitted]`

`Entrypoint application [big] = css/application-7891a208.css js/application-7da76fec4ff71205bb73.js js/application-7da76fec4ff71205bb73.js.map`

`Entrypoint notes = js/notes-aa3d2eba82abbbda5606.js js/notes-aa3d2eba82abbbda5606.js.map`

`[1] (webpack)/buildin/module.js 552 bytes {0} {1} [built]`

`[2] ./app/javascript/packs/notes.js 2.53 KiB {0} {1} [built]`

`[3] ./app/javascript/packs/application.js 850 bytes {0} [built]`

`[7] ./app/javascript/channels/index.js 205 bytes {0} [built]`

`[8] ./app/javascript/channels sync _channel\.js$ 160 bytes {0} [built]`

`[9] ./app/javascript/stylesheets/application.scss 39 bytes {0} [built]`

`+ 5 hidden modules`

&amp;#x200B;

`WARNING in asset size limit: The following asset(s) exceed the recommended size limit (244 KiB).`

`This can impact web performance.`

`Assets:`

`css/application-7891a208.css (826 KiB)`

&amp;#x200B;

`WARNING in entrypoint size limit: The following entrypoint(s) combined asset size exceeds the recommended limit (244 KiB). This can impact web performance.`

`Entrypoints:`

`application (983 KiB)`

`css/application-7891a208.css`

`js/application-7da76fec4ff71205bb73.js`

`WARNING in webpack performance recommendations:`

`You can limit the size of your bundles by using import() or require.ensure to lazy load some parts of your application.`

`For more info visit` [`https://webpack.js.org/guides/code-splitting/`](https://webpack.js.org/guides/code-splitting/)

`Child mini-css-extract-plugin node_modules/css-loader/dist/cjs.js??ref--7-1!node_modules/postcss-loader/src/index.js??ref--7-2!node_modules/sass-loader/dist/cjs.js??ref--7-3!app/javascript/stylesheets/application.scss:`

`Entrypoint mini-css-extract-plugin = *`

`[0] ./node_modules/css-loader/dist/cjs.js??ref--7-1!./node_modules/postcss-loader/src??ref--7-2!./node_modules/sass-loader/dist/cjs.js??ref--7-3!./app/javascript/stylesheets/application.scss 1.45 MiB {0} [built]`

`+ 1 hidden module`

Thanks in advance!
