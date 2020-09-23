# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/iya619/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/iya619/personal_projects_show_off_your_own_project_andor/
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
## [3][flag_shih_tzu gem , return same value for any key](https://www.reddit.com/r/rails/comments/iy9u1k/flag_shih_tzu_gem_return_same_value_for_any_key/)
- url: https://www.reddit.com/r/rails/comments/iy9u1k/flag_shih_tzu_gem_return_same_value_for_any_key/
---
I'm using **flag\_shih\_tzu** gem and at my User model I include that lines  


    include FlagShihTzu
    
       has_flags  2 =&gt; :admin,
                  3 =&gt; :moderator

I have a column **flags** as well in User table  **t.integer "flags", default: 1, null: false**  


but when I set Flags value to 2 or 3 it doesn't matter It will always return me the same value admin? for example  


    user = User.first
    user.flags = 2
    user. save 
    
    User.first.admin? 
    true 
    
    user = User.second  
    user.flags = 3
    user.save
    
    User.second.admin?
    true
    User.second.moderator?
    false
## [4][Where to start](https://www.reddit.com/r/rails/comments/iy3dmv/where_to_start/)
- url: https://www.reddit.com/r/rails/comments/iy3dmv/where_to_start/
---
Hey everyone,

I have an interview for a Java/Ruby position next Monday. I am proficient in Java but I have no exposure to Ruby and I am looking for resources to get started. I read the first few chapters of 'Eloquent Ruby' this afternoon and it seems to share similarities with Python/Elixir - both of which I have experience with. 

Any advice would be appreciated :)

Thanks in advance!
## [5][Why is dependent: :destroy not working for me?](https://www.reddit.com/r/rails/comments/ixwox5/why_is_dependent_destroy_not_working_for_me/)
- url: https://www.reddit.com/r/rails/comments/ixwox5/why_is_dependent_destroy_not_working_for_me/
---
I have two models:

    class Detection &lt; ApplicationRecord
      has_many :links, foreign_key: 'start_id', dependent: :destroy
      has_many :links, foreign_key: 'end_id', dependent: :destroy
    end
    
    class Link &lt; ApplicationRecord
      belongs_to :start, class_name: 'Detection', foreign_key: :start_id
      belongs_to :end, class_name: 'Detection', foreign_key: :end_id
    end

When I delete a detection, I want the associated links to be destroyed. However, I'm getting a failed foreign-key constraint from pg when I try to delete a Detection, suggesting that my associations are not being described correctly here:

    PG::ForeignKeyViolation: ERROR: update or delete on table "detections" violates foreign key constraint "fk_rails_8633ff3bb0" on table "links" DETAIL: Key (id)=(1) is still referenced from table "links".


How do I tell Rails to delete any links where start\_id or end\_id match the deleted Detection's id?
## [6][Is there a better way to do what I just did??](https://www.reddit.com/r/rails/comments/ixw4kt/is_there_a_better_way_to_do_what_i_just_did/)
- url: https://www.reddit.com/r/rails/comments/ixw4kt/is_there_a_better_way_to_do_what_i_just_did/
---
Hi all, I’m a Rails beginner and just implemented a feature that has resulted in using multiple conditionals in my controllers and views, and am just wondering if there might be a more elegant way to do this?

I’m working on a practice app that lets you create online Invoices, and now finished a letting you create Estimates as well....which are simply Invoice records, but with a “estimate” Boolean attribute set to TRUE. 

So I have a “New Invoice” button that calls the ‘new’ action in the Invoices controller, AND I have a “New Estimate” button that also calls that same action, but passes in a “estimate: true” hash. 

Then, in the controller, I use an if/else conditional to check for that hash, and if it’s there, I create a new record with the ‘estimate’ Boolean set to true. Otherwise, I create a new record without setting the estimate boolean. 

Finally, for my Invoice views (show and  \_form, along with Active Mailer views) I also do a conditional check so the views properly display the words “Invoice” or “Estimate” depending on what the record’s estimate Boolean is, and also displays a different ‘Back’ or ‘Cancel’ button which brings you back to either the Invoices or Estimates index page (speaking of which, buttons to the Invoices page and Estimates page work the same way, with a conditional passed to the Index controller action, and that action using a conditional to serve up either invoices or estimates).

So that’s how it works, and it DOES work, but I just wonder if there might be a better, cleaner way to do something like this? I haven’t figured one out despite some research, but Rails is filled with so many clever tools, I thought I’d ask here!

Thanks very much!
## [7][Becoming a regular Rails contributor?](https://www.reddit.com/r/rails/comments/ixn3qh/becoming_a_regular_rails_contributor/)
- url: https://www.reddit.com/r/rails/comments/ixn3qh/becoming_a_regular_rails_contributor/
---
Are there any regular Rails contributors here?

It's been an old goal of mine to become one. Now I"m not expecting to reach Tenderlove's level, but I want to get to the level where I have at least a vague idea how to solve regular bugs in ActiveRecord. That I know enough about the internals I can come up with solutions given enough time. When I now look at Rails bugs, especially ActiveRecord, I usually wouldn't know even what files are causing the problem - is it AssociationReflection? AssociationProxy maybe?. When I look at someone's bug fix it sorta makes sense,  but the interaction between all the different classes and modules is too much for my brain to handle.

I guess I need to practice more, but it seems like ordinary bugs are often times too hard for me to practice on. It seems like regular contributors have a tons of contextual knowledge on ActiveRecord that I'm missing.

P.S : I'm not sure why but I'm more interested in ActiveRecord than other parts of Rails, and it makes sense to me to focus on one part in any case.

Has anyone here made the transition to becoming a regular contributor and care to share how this came to be?
## [8][3rd party vendor SaaS is sending us us 'action' variable in POST](https://www.reddit.com/r/rails/comments/ixj48w/3rd_party_vendor_saas_is_sending_us_us_action/)
- url: https://www.reddit.com/r/rails/comments/ixj48w/3rd_party_vendor_saas_is_sending_us_us_action/
---
Hi,

Got a problem. 3rd party vendor SaaS is sending us 'action' variable in POST, which is getting overwritten by the Rails routing with the method name.

Do you have an idea for a workaround? 

The only idea we have now is to create a front app that will receive such payload and will change the action parameter to action2 (or something) and then it will send it to our rails app.
## [9][Decorating API response before](https://www.reddit.com/r/rails/comments/ixkeic/decorating_api_response_before/)
- url: https://www.reddit.com/r/rails/comments/ixkeic/decorating_api_response_before/
---
Hi everyone, so I'm back after [this](https://www.reddit.com/r/rails/comments/it350w/connecting_to_wordpress_db_from_rails/). I've essentially created an adapter to connect the crud operations from WP API so that it can be used in our app. We've stored some of the posts' data inside our DB and sync regularly for updated content.

Currently, the front end is fetching the posts via the WP API but now, we want to add some fields to the posts records using information we have in our DB.

The original plan would be to simply send the related information (with post ids) to the front end and cross-check them with the fetched post ids, then combine them. However, it doesn't feel really clean IMO.

My question is would it be better to fetch the posts on the back end, "decorate" the response with the new fields, then sending the decorated response to the front end? Seems like a roundabout way of doing things but is there a pros/cons to it? Or if anyone else have any suggestions would be much appreciated!
## [10][Simple Blog Section For Webapp](https://www.reddit.com/r/rails/comments/ixivht/simple_blog_section_for_webapp/)
- url: https://www.reddit.com/r/rails/comments/ixivht/simple_blog_section_for_webapp/
---
Hi there :-)

&amp;#x200B;

I built a webapp that recently went live, and I'd like to start creating blogposts for Content Marketing purposes. Does anyone in the community happen to have suggestions/advice regarding building a simple blog section for a functioning webapp?

&amp;#x200B;

Simple:

\- Users (Devise) are able to view and comment

\- Administrators can create posts

&amp;#x200B;

I got comfy-blog up and running, but the styling is a little too bare-boned for me. At this point, I'm thinking of custom coding it, SpinaCMS, ButterCMS, or maybe even just using Wordpress. 

&amp;#x200B;

Cheers!
## [11][Adobe-Approved Trustlist Document Signing with Rails (SAD story)](https://www.reddit.com/r/rails/comments/ix58uo/adobeapproved_trustlist_document_signing_with/)
- url: https://www.reddit.com/r/rails/comments/ix58uo/adobeapproved_trustlist_document_signing_with/
---
&amp;#x200B;

We're creating a primitive version of DocuSign within our app for document processing in the COVID age.

Apparently, PDF's have both an ink signature and a digital signature that is like a checksum for the PDF, so the user can be sure it hasn't been changed since the signature was applied.

I found this quick and dirty document signing Ruby script: [https://gist.github.com/matiaskorhonen/223bd527279cf49bed1e](https://gist.github.com/matiaskorhonen/223bd527279cf49bed1e) , which appears to work as intended with a self-signed cert, but Acrobat doesn't trust the self-signed certs.

In order to get Adobe to recognize the digital signature, it has to come from one of the certificate providers on this list: [https://helpx.adobe.com/acrobat/kb/approved-trust-list1.html](https://helpx.adobe.com/acrobat/kb/approved-trust-list1.html)

I have spent over a month trying to get a certificate. As far as I can tell, almost all of them try to sell you a USB stick which contains the private key and is a required hardware module for compliancy with the US Esign Act. I don't understand: how am I supposed to use a physical stick with our AWS-hosted website? Who would want the certificate on a USB stick?

One provider told me that I could host the certificate on AWS CloudHSM, which seems like it would be the best place to keep it, as there are both Ruby and Linux client libraries for interacting with it. However, this costs $1.80/hour, aka $12,000/year, so even this is an imperfect solution.

The prices I've been quoted have been horrendous:

$5000/year for 50,000 signatures per year (This price is 50% off!)

Most of them also don't offer money-back guarantees, which I really want in order to ensure that the Adobe recognizes the signature created by the crude Origami method linked above.

After several weeks, providing a utility bill in my name, my notarized passport, and a bank statement one AATL provider, along with a phonebook verification to my company's founder, and asking finance to give them some more company bills, I finally got a certificate for $100/year for 3 years with refund guarantee. HOWEVER, despite them SWEARING to me that they would allow me to host the certificate on AWS Cloud HSM, they said "oh haha!!! We don't do that for document signing certificates!!! hahaha!!!! Why would you think that!!!" So I am begging them to make an exception for me (upgrade to ENTERPRISE!) but I'm not optimistic.

So what am I supposed to do?

\#1 Is there any way to use these stupid USB sticks with a Rails app hosted on AWS?

\#2 Do I really need the digital signatures? Why don't I just leave the PDF's without a digital signature and just an ink signature? I highly doubt that every PDF I ever open on my computer has a digital signature.

\#3 Can I just keep our own server-side checksums of the PDF's in case there is ever a dispute?

\#4 Can I try using one of the foreign providers to save some money? Or are these certificates scoped by country?

\#5 Do any of you have any other recommendations on how to solve this?
## [12][Recommended machines for pro Rails development](https://www.reddit.com/r/rails/comments/ix1dro/recommended_machines_for_pro_rails_development/)
- url: https://www.reddit.com/r/rails/comments/ix1dro/recommended_machines_for_pro_rails_development/
---
Hey guys!  I’m in the market for a new development machine.

I’ve been learning on a four year old Acer 11 chromebook that I threw the Gallium distribution of Linux on once I started learning to code.  I got it for around $200 and besides the Linux glitches and the fact the 11 inch tn screen about near makes my eyes bleed, it’s been a really solid development machine, but it’s starting to show its age.  

I would replace it with another Linux run chromebook in a heartbeat, but the guy who heads Gallium’s development says that chromebooks and Linux have diverged in irreconcilable ways, so I’m trying to work out my options. 

Spec wise, the four gigs of ram and old Celeron have been doable so far, but I imagine once I get into larger projects I’ll need more.  I want to be as budget conscious as I can, while still getting a machine that will meet a good share of my future needs.

So I have two main questions for you guys:

\-What specs are actually needed for pro Rails development?

\-What machines have you had a good experience with, regardless of price?
