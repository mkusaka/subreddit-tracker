# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gvu083/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gvu083/personal_projects_show_off_your_own_project_andor/
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
## [2][Slides What Comes After MVC - Peter Harkins](https://www.reddit.com/r/rails/comments/gvs0dh/slides_what_comes_after_mvc_peter_harkins/)
- url: https://www.reddit.com/r/rails/comments/gvs0dh/slides_what_comes_after_mvc_peter_harkins/
---
Peter talked so fast. I can't keep up, let alone digest the ideas presented.

Does anyone have the slides and notes to [https://push.cx/2015/railsconf](https://push.cx/2015/railsconf) that they can share?

Thanks.
## [3][Services and tools we use for development](https://www.reddit.com/r/rails/comments/gvpvlp/services_and_tools_we_use_for_development/)
- url: https://www.reddit.com/r/rails/comments/gvpvlp/services_and_tools_we_use_for_development/
---
Tools and services that help to automatize development flow. Easy to install with increasing team effectiveness.

[https://jtway.co/services-and-tools-we-use-for-development-2749af5ad08a](https://jtway.co/services-and-tools-we-use-for-development-2749af5ad08a)
## [4][Question about how a Rails backend works...](https://www.reddit.com/r/rails/comments/gvm3r6/question_about_how_a_rails_backend_works/)
- url: https://www.reddit.com/r/rails/comments/gvm3r6/question_about_how_a_rails_backend_works/
---
This is a general programming question and may not be specific to Rails. However, it's an issue I've run into building an app with a Rails backend.

Say I make an app and store user information in the Rails server-- such as a login and pw-- and push it all to github... if somebody clones that repo, runs it and tries to login with a login/pw that I saved will they be able to do it?
## [5][Dynamically added inputs are not sent as parameters...](https://www.reddit.com/r/rails/comments/gvnje5/dynamically_added_inputs_are_not_sent_as/)
- url: https://www.reddit.com/r/rails/comments/gvnje5/dynamically_added_inputs_are_not_sent_as/
---
[The same thing that happened to this guy is happening to me](https://stackoverflow.com/questions/52200128/rails-dynamic-add-nested-form-error-not-show-in-parameters), in my case I have two inputs `&lt;%= number_field_tag 'conciliacion_diferencias[]', nil, class: 'form-control', placeholder: 'ID de la venta' %&gt;` and I added a third one with Javascript, but only the first two are being sent to the controller...

This is the HTML I see when I use inspect in the parent `div` of the 3 inputs... I don't see any difference...

```
&lt;div id="difs-inputs" class="col-sm-6"&gt;
&lt;input type="number" name="conciliacion_diferencias[]" id="conciliacion_diferencias_" class="form-control" placeholder="ID de la venta"&gt;
&lt;br&gt;
&lt;input type="number" name="conciliacion_diferencias[]" id="conciliacion_diferencias_" class="form-control" placeholder="ID de la venta"&gt;
&lt;br&gt;
&lt;input id="conciliacion_diferencias_" type="number" name="conciliacion_diferencias[]" placeholder="ID de la venta" class="form-control"&gt;&lt;/div&gt;
```

I added the third input with the following JS function... 

```
                function addFields() {
                    var dif_container = document.getElementById("difs-inputs");
                    var tx_container = document.getElementById("txs-inputs");

                    var dif_br = document.createElement("br");
                    var tx_br = document.createElement("br");

                        var dif_input = document.createElement("input");
                        dif_input.id = "conciliacion_diferencias_";
                        dif_input.type = "number";

                        dif_input.name = "conciliacion_diferencias[]";
                        dif_input.placeholder = "ID de la venta"
                        dif_input.className = "form-control"
                        dif_container.appendChild(dif_br);
                        dif_container.appendChild(dif_input);
                }
```

But when I put "1", "2" and "3" in each one of the 3 inputs (if it helps you to understand, [I made a screenshot](https://i.imgur.com/D28KDZO.jpg)) and submit I only see the first two values/parameters sent, in other words I see this in the params...

```
... "conciliacion_diferencias"=&gt;["1", "2"], ...
```

I don't get what's happening... Why are the dynamically created inputs not shown/sent in the parameters...
## [6][Can I ignore some Activerecords syntax in our app ?](https://www.reddit.com/r/rails/comments/gvpxar/can_i_ignore_some_activerecords_syntax_in_our_app/)
- url: https://www.reddit.com/r/rails/comments/gvpxar/can_i_ignore_some_activerecords_syntax_in_our_app/
---
Is is possible to ignore some syntaxs when we compile it run in local or prod for our app?
I want to ignore  these: params[:model].permit! , Model.where( "id IN #{ids}.join(',')" )  literal string in WHERE  got brakeman warning SQL injection, and permit!
Is it possible to ignore this and make warning before run server ? 
Sorry I am typing on phone :)
## [7][Building A Drag &amp; Drop Block Builder](https://www.reddit.com/r/rails/comments/gvpt0e/building_a_drag_drop_block_builder/)
- url: https://www.reddit.com/r/rails/comments/gvpt0e/building_a_drag_drop_block_builder/
---
I want to build a drag and drop newsletter builder, something similar like in the image below. Can you give me any tips or suggestions?

https://preview.redd.it/kc3l2ua6kn251.png?width=1914&amp;format=png&amp;auto=webp&amp;s=498698bf60b2777f8f5c52fd8e499b144b7db854
## [8][Question about a potential rails project](https://www.reddit.com/r/rails/comments/gvk154/question_about_a_potential_rails_project/)
- url: https://www.reddit.com/r/rails/comments/gvk154/question_about_a_potential_rails_project/
---
Want to a project where I build a website like thumbtack. If anyone doesn’t know what thumbtack is its basically a website that matches customers with local professionals for a wide range of services that someone may need; an example could be locksmith. How hard would it be to create something like this ? Also how would I go about getting the information from all local businesses near someone’s location for their particular task ? Would I use an api and then store that Information in a database ? If so what api could I use ? Any advise for this project would be much appreciated. Thank you for your time and help  !
## [9][Resources for Fancy Links? Specifically trying to populate user/email field if redirected to login after click](https://www.reddit.com/r/rails/comments/gv3ydx/resources_for_fancy_links_specifically_trying_to/)
- url: https://www.reddit.com/r/rails/comments/gv3ydx/resources_for_fancy_links_specifically_trying_to/
---
The business that I’m designing for has a lot of stubborn, crotchety customers, and in this new infrastructure, I’m trying to save the employees as much time as possible by empowering customers (who don’t know how to search their email, or google a tracking number, or make it two clicks past a login page).

What’s the best way for including params in a link that then update the value of an input?

I feel like I could just sit down and build a param trigger to render JS and put logic in a js.erb file, but something in my gut is saying that tailoring to unique params situations is usually a job for a pack full of jquery
## [10][Do you prefer to use Auth0? 😊](https://www.reddit.com/r/rails/comments/gvj86a/do_you_prefer_to_use_auth0/)
- url: https://www.reddit.com/r/rails/comments/gvj86a/do_you_prefer_to_use_auth0/
---
This is how to implement to get access token with client credentials in Ruby.

[https://noknow.info/it/ruby/auth0/how\_to\_implement\_to\_get\_access\_token\_with\_client\_credentials](https://noknow.info/it/ruby/auth0/how_to_implement_to_get_access_token_with_client_credentials)
## [11][Make user_id in Invoice table the same as id in User table](https://www.reddit.com/r/rails/comments/gv89lw/make_user_id_in_invoice_table_the_same_as_id_in/)
- url: https://www.reddit.com/r/rails/comments/gv89lw/make_user_id_in_invoice_table_the_same_as_id_in/
---
I'd like to make the. id of each user created match the user\_id in my invoices table. They are associated by the same customer\_name, which the :customer\_name is in both tables. Currently, when I create and invoice, user\_id: nil. I'm not sure if I have to make an action in my controller, write an association in my routes, or define a method in my invoice model. Any help or things I can try is greatly appreciated.
