# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][3rd party vendor SaaS is sending us us 'action' variable in POST](https://www.reddit.com/r/rails/comments/ixj48w/3rd_party_vendor_saas_is_sending_us_us_action/)
- url: https://www.reddit.com/r/rails/comments/ixj48w/3rd_party_vendor_saas_is_sending_us_us_action/
---
Hi,

Got a problem. 3rd party vendor SaaS is sending us 'action' variable in POST, which is getting overwritten by the Rails routing with the method name.

Do you have an idea for a workaround? 

The only idea we have now is to create a front app that will receive such payload and will change the action parameter to action2 (or something) and then it will send it to our rails app.
## [3][Becoming a regular Rails contributor?](https://www.reddit.com/r/rails/comments/ixn3qh/becoming_a_regular_rails_contributor/)
- url: https://www.reddit.com/r/rails/comments/ixn3qh/becoming_a_regular_rails_contributor/
---
Are there any regular Rails contributors here?

It's been an old goal of mine to become one. Now I"m not expecting to reach Tenderlove's level, but I want to get to the level where I have at least a vague idea how to solve regular bugs in ActiveRecord. That I know enough about the internals I can come up with solutions given enough time. When I now look at Rails bugs, especially ActiveRecord, I usually wouldn't know even what files are causing the problem - is it AssociationReflection? AssociationProxy maybe?. When I look at someone's bug fix it sorta makes sense,  but the interaction between all the different classes and modules is too much for my brain to handle.

I guess I need to practice more, but it seems like ordinary bugs are often times too hard for me to practice on. It seems like regular contributors have a tons of contextual knowledge on ActiveRecord that I'm missing.

P.S : I'm not sure why but I'm more interested in ActiveRecord than other parts of Rails, and it makes sense to me to focus on one part in any case.

Has anyone here made the transition to becoming a regular contributor and care to share how this came to be?
## [4][Decorating API response before](https://www.reddit.com/r/rails/comments/ixkeic/decorating_api_response_before/)
- url: https://www.reddit.com/r/rails/comments/ixkeic/decorating_api_response_before/
---
Hi everyone, so I'm back after [this](https://www.reddit.com/r/rails/comments/it350w/connecting_to_wordpress_db_from_rails/). I've essentially created an adapter to connect the crud operations from WP API so that it can be used in our app. We've stored some of the posts' data inside our DB and sync regularly for updated content.

Currently, the front end is fetching the posts via the WP API but now, we want to add some fields to the posts records using information we have in our DB.

The original plan would be to simply send the related information (with post ids) to the front end and cross-check them with the fetched post ids, then combine them. However, it doesn't feel really clean IMO.

My question is would it be better to fetch the posts on the back end, "decorate" the response with the new fields, then sending the decorated response to the front end? Seems like a roundabout way of doing things but is there a pros/cons to it? Or if anyone else have any suggestions would be much appreciated!
## [5][Simple Blog Section For Webapp](https://www.reddit.com/r/rails/comments/ixivht/simple_blog_section_for_webapp/)
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
## [6][Adobe-Approved Trustlist Document Signing with Rails (SAD story)](https://www.reddit.com/r/rails/comments/ix58uo/adobeapproved_trustlist_document_signing_with/)
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
## [7][I just refused a Django/React contract because I love Rails too much. Was it a good idea ?](https://www.reddit.com/r/rails/comments/ixm5u0/i_just_refused_a_djangoreact_contract_because_i/)
- url: https://www.reddit.com/r/rails/comments/ixm5u0/i_just_refused_a_djangoreact_contract_because_i/
---

## [8][Recommended machines for pro Rails development](https://www.reddit.com/r/rails/comments/ix1dro/recommended_machines_for_pro_rails_development/)
- url: https://www.reddit.com/r/rails/comments/ix1dro/recommended_machines_for_pro_rails_development/
---
Hey guys!  I’m in the market for a new development machine.

I’ve been learning on a four year old Acer 11 chromebook that I threw the Gallium distribution of Linux on once I started learning to code.  I got it for around $200 and besides the Linux glitches and the fact the 11 inch tn screen about near makes my eyes bleed, it’s been a really solid development machine, but it’s starting to show its age.  

I would replace it with another Linux run chromebook in a heartbeat, but the guy who heads Gallium’s development says that chromebooks and Linux have diverged in irreconcilable ways, so I’m trying to work out my options. 

Spec wise, the four gigs of ram and old Celeron have been doable so far, but I imagine once I get into larger projects I’ll need more.  I want to be as budget conscious as I can, while still getting a machine that will meet a good share of my future needs.

So I have two main questions for you guys:

\-What specs are actually needed for pro Rails development?

\-What machines have you had a good experience with, regardless of price?
## [9][Setting up a Rails API with React,](https://www.reddit.com/r/rails/comments/iwhjm9/setting_up_a_rails_api_with_react/)
- url: https://www.reddit.com/r/rails/comments/iwhjm9/setting_up_a_rails_api_with_react/
---
I'm fairly new to the Rails/JS environment. I've been struggling for a while to get a workflow going that feels comfortable to use.   
The main goal is to get the development environment dockerized, so as to not pollute the local system, while still preserving nice features, such as debuggers, live reload, etc.

Most guides on the internet either use two separate projects, which would be easy to set up, but annoying to work with, as.  
Making an --api rails project prohibits the use of the builtin webpacker functionality.  
Making a normal project, but then using webpack means that you end up with tons of unused templates, that are still required by rails, as well as resources having the non-rest endpoints.  


Any ideas/advice are welcome..
## [10][Anniversary (checking only day and month)](https://www.reddit.com/r/rails/comments/iwed56/anniversary_checking_only_day_and_month/)
- url: https://www.reddit.com/r/rails/comments/iwed56/anniversary_checking_only_day_and_month/
---
Using a scope like this

        scope :anniversary, lambda { 
          where('created_at &gt; ?', 1.year.ago)
        }

I can check only if **an event was created 1 year ago**. But now I want to change it.

I want to check if it was created also 2, 3, 4, etc. years ago. **How to compare only the month and the day?**

I want to have events created the same day of the month from all the years.
## [11][File types as part of ActiveStorage uploads](https://www.reddit.com/r/rails/comments/ivyqcw/file_types_as_part_of_activestorage_uploads/)
- url: https://www.reddit.com/r/rails/comments/ivyqcw/file_types_as_part_of_activestorage_uploads/
---
Hi /r/rails

I am working on an internal Rails app around adding products to our warehouse inventory to replace a paper based filing system.

When an employee adds a product, they can (if its available) add the OEM's product data sheet to their upload.

There are also other PDF documents that become part of the product's folder including:

* Procurement (who would add receipts for buys from the suppliers/manufacturers)
* Occupational safety team (They have a PDF form they fill out OK'ing warehousing the item and if it has to be stored at different heights or separated from other items)
* Receiving (they scan items when received and would upload an exported PDF from their receiving system)
* Loss prevention team (some items are kept caged and they identify these)
* Legal (A few items are export controlled and legal provides an opinion on if they need to be identified as such)

I'd like to be able to identify these files as part of display of files and workflow through the app

What is the easiest way to manage this in rails?

My first thought was to have my product model have

`has_many :files`

Then create a file model and table that has an enum for all the types of files ("MSDS", "Receipt", "Receiving", "Safety", "Security", "Legal") and then this file model to have:

    has_many_attached :msds
    has_many_attached :receipt
    has_many_attached :receiving_reports
    has_many_attached :safety
    has_many_attached :security
    has_many_attached :legal

Is this the best way to go about this? Am I over or under thinking this?
