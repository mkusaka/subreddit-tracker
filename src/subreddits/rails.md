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
## [3][Is It Possible to Learn Rails Development by Volunteering?](https://www.reddit.com/r/rails/comments/fx080d/is_it_possible_to_learn_rails_development_by/)
- url: https://www.reddit.com/r/rails/comments/fx080d/is_it_possible_to_learn_rails_development_by/
---
I'm 52 years old former attorney. I suffer from a bi-polar condition that has rendered me disabled for 19 years. I'm still licensed but because of my condition I don't want to expose myself to the stress of practicing law. I've always enjoyed computers and in my college days I supported myself working as a  computer support technician. In fact, after I obtained my BS, I worked as a phone tech support for Packard Bell (arguably the worst computer manufacturer ever) until the company closed operations in California after the Northridge Earthquake of 1994.

Afterwards, I worked for an insurance sales company where I sold insurance by day and wrote the company's software by night. I used MS Access Basic. For a small company with fewer than 10 computers that were networked via a peer to peer network. It worked surprisingly well. I then went to law school and practiced for about 5 years before my illness disabled me.

As a therapy, I've studied linux and ROR.  I'd like to gain proficiency in Rails development to further my well being. My question to the group is it possible to do this by volunteering? I currently do not have the skill set nor experience to do this commercially. Also, by working for pay, I could jeopardize my disability status. If I lose my disability status, I could lose my medical coverage.
## [4][good resources for learning testing in Rails](https://www.reddit.com/r/rails/comments/fwkcyx/good_resources_for_learning_testing_in_rails/)
- url: https://www.reddit.com/r/rails/comments/fwkcyx/good_resources_for_learning_testing_in_rails/
---
I've posted about them before but was curious and went ahead in the curriculum, but as a part of their free extensive Rails course, they have a large section (14.5 hrs) of testing at AppAcademy Open

https://open.appacademy.io/learn/full-stack-online/rails/rails-testing--intro

Here is a look at most of it:

https://imgur.com/a/BTlm7v8

Just another resource for those out there who may feel they are fuzzy and this might help fill some gaps, or be the main learning path.
## [5][understanding has_many x through: y](https://www.reddit.com/r/rails/comments/fwr58t/understanding_has_many_x_through_y/)
- url: https://www.reddit.com/r/rails/comments/fwr58t/understanding_has_many_x_through_y/
---
I was watching this lesson and around the 3 min mark is the relevant material:

https://open.appacademy.io/learn/full-stack-online/sql/more-associations--has_many-through-----

one has to have an account(free), but here are images of the relevant code:

https://imgur.com/a/vWsjOF8

The instructor was emphasizing that the `through` ActiveRecord method in `through: :dogs` is a key with a value paired to the `:dogs`, and that `:dogs` here is a **method** - I guess in this case referring to the `House#dogs` method and NOT referring to the `Dog` class itself.  Same with the `source: :toys`, that `:toys` is also a method and NOT the `Toy` model, so I guess in that case it would be the `Dog#toys` method that the `:toys` in `source: :toys` is referring to.  Is the above understanding correct?
## [6][Yet another active form](https://www.reddit.com/r/rails/comments/fwk4ip/yet_another_active_form/)
- url: https://www.reddit.com/r/rails/comments/fwk4ip/yet_another_active_form/
---
Hey! I just want to share with you a gem we've been working on recently and it's about form objects. 

Me and my coworkers built an abstraction to handle these form objects in one of our client's projects, and provided it was so helpful we decided to extract it to a gem.

I would really appreciate every early feedback I can get, we just published v0.1.0. 

[https://github.com/rootstrap/yaaf](https://github.com/rootstrap/yaaf)
## [7][How to implement an unknown 'end_time' as default in datetime column?](https://www.reddit.com/r/rails/comments/fwmwkb/how_to_implement_an_unknown_end_time_as_default/)
- url: https://www.reddit.com/r/rails/comments/fwmwkb/how_to_implement_an_unknown_end_time_as_default/
---
I'm creating a table that uses \`t.datetime: start\_time, default: Time.now\` and I need to set \`end\_time\` as  a default of 'open ended' or nil but would rather not put \`nil\` into my db fields. Is there a way to set an unknown value in datetime as default?
## [8][Deep namespaces practices](https://www.reddit.com/r/rails/comments/fwlzrm/deep_namespaces_practices/)
- url: https://www.reddit.com/r/rails/comments/fwlzrm/deep_namespaces_practices/
---
I'm working on rails api-only app. I've decided to split app into two major modules Auth and Social, so eg. comments_controller is in Api::V1::Social:: CommentsController, Api namespace look unnecessary, but in future is possible to add web interface, so I think it's ok to made namespace now. To keep my errors standardized I've decided to make new module ErrorResponses which have several error classes, I've decided to create errors modules independent for Auth module and Social module, to make easier to maintain by different people or split app into microservices easier, so I end with something like that:
render Api::V1:: Social:: ErrorResponses::WrongProfile
I thinkwd about that a lot, and in my opinion it's ok, but I have some doubts about such deep nested namespaces, could you tell me if my solution is elegant?
## [9][A hands-on tutorial to debugging your code with pry-byebug](https://www.reddit.com/r/rails/comments/fwgqcp/a_handson_tutorial_to_debugging_your_code_with/)
- url: https://www.reddit.com/r/rails/comments/fwgqcp/a_handson_tutorial_to_debugging_your_code_with/
---
Howdy folks!

Like most new developers, I started as a puts developer. Then, I discovered pry-byebug and debugging things got a lot easier.

So for those of you who can’t wrap their head around pry-byebug debugger, I wrote a hands-on tutorial that’ll walk you through its most basic commands.

You’ll learn how to: ⏸ Pause code execution 🆕 Get current values ⏭ Move on to the next line and wait ⏏️ Run code execution until the end of the process

https://remimercier.com/pry-byebug-tutorial/

Hope it'll help!
## [10][Can a some one please help me out before I go insane](https://www.reddit.com/r/rails/comments/fwjxrq/can_a_some_one_please_help_me_out_before_i_go/)
- url: https://www.reddit.com/r/rails/comments/fwjxrq/can_a_some_one_please_help_me_out_before_i_go/
---
I am new to rails and I am trying to make a tinder clone using this tutorial:

[https://www.youtube.com/watch?v=P5gAaZq-sPs](https://www.youtube.com/watch?v=P5gAaZq-sPs)   (22min)

I have set up slider  but the buttons just don't do anything this this is the code for the buttons:

&lt;button class="controls" id="decline"&gt;No&lt;/button&gt;

&lt;button class="controls" id="approve"&gt;Yes&lt;/button&gt;

The button seem not to be linked with the actual slider, what am I doing wrong?

&amp;#x200B;

This is the js:

&amp;#x200B;

$(function(){

var $activeSlide = $('#slides .slide:first-child');

&amp;#x200B;

// show first slide

$activeSlide.addClass("showing");

&amp;#x200B;

$("#decline").on("click", function(){

goToSlide('decline');

&amp;#x200B;

});

&amp;#x200B;

$("#approve").on("click", function(){

goToSlide('approve');

console.log(action);

&amp;#x200B;

});

&amp;#x200B;

function goToSlide(action){

$activeSlide.removeClass("showing");

&amp;#x200B;

// send data to controller

if(action == "approve"){

&amp;#x200B;

} else {

&amp;#x200B;

$activeSlide.addClass("showing");

&amp;#x200B;

//slides\[currentSlide\].className = 'slide';

//currentSlide = (n+slides.length)%slides.length;

//slides\[currentSlide\].className = 'slide showing';

}

});
## [11][Rails 4 callback sequence](https://www.reddit.com/r/rails/comments/fwhyyb/rails_4_callback_sequence/)
- url: https://www.reddit.com/r/rails/comments/fwhyyb/rails_4_callback_sequence/
---
Hi, I hope you all are healthy.

I need to know how to enforce sequence in callbacks.

e.g -&gt; `after_create :fun_a, :fun_b, :fun_c`  
In above code, suppose I want AR object to call these callbacks in order: `fun_a -&gt; fun_b -&gt; fun_c`  
Is there anyway to achieve this?  


Thanks :)
## [12][Getting started with GraphQL: Trying to do an update mutation](https://www.reddit.com/r/rails/comments/fw9271/getting_started_with_graphql_trying_to_do_an/)
- url: https://www.reddit.com/r/rails/comments/fw9271/getting_started_with_graphql_trying_to_do_an/
---
EDIT: Solution for how I fixed this is at the bottom!

&amp;#x200B;

Hi there.

So over the past few weeks I've been messing around with Ruby, Rails and lately also GraphQL. At the moment I have a demo project which is an api server with users and posts in a one to many relationships, and so far I've managed to do a lot of different queries with them, and also mutations to create new ones of either type.

&amp;#x200B;

The problem now is that I am trying to make an update post mutation with the following code:

&amp;#x200B;

https://preview.redd.it/6m3icfx83ar41.png?width=1092&amp;format=png&amp;auto=webp&amp;s=5ea26d914acccbb9b9b836666042135175cb3da8

so I can make a query that looks like this:

&amp;#x200B;

https://preview.redd.it/zqn47ujb3ar41.png?width=1714&amp;format=png&amp;auto=webp&amp;s=50ac65745efc8623d134a7d2ae0ea5790d7cf66f

problem I am having though is that it wants me to provide the user id of the post to update, even though I am not interested in updating that field.

&amp;#x200B;

I guess my question is, how do I need to structure the mutation so that I can update any of the fields independently of one another, instead of having to provide everything?

I know there is probably something really simple that I am missing, I have found the documentation either really hard to find, seemingly outdated, or extremely basic.

&amp;#x200B;

best regards

&amp;#x200B;

EDIT: You are all an amazing community, with your help the problem was solved like so:

&amp;#x200B;

https://preview.redd.it/n95dei3t9dr41.png?width=1062&amp;format=png&amp;auto=webp&amp;s=62f3a5456eb250cf82f65b964d9f6651d4e467ad
