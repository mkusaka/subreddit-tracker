# javascript
## [1][WTF Wednesday (July 08, 2020)](https://www.reddit.com/r/javascript/comments/hnjqff/wtf_wednesday_july_08_2020/)
- url: https://www.reddit.com/r/javascript/comments/hnjqff/wtf_wednesday_july_08_2020/
---
Post a link to a GitHub repo that you would like to have reviewed, and brace yourself for the comments!
Whether you're a junior wanting your code sharpened or a senior interested in giving some feedback and have some time to spare, 
this is the place.

[Named after this comic](https://davidwalsh.name/demo/code-review.png)
## [2][Showoff Saturday (July 11, 2020)](https://www.reddit.com/r/javascript/comments/hp8pa9/showoff_saturday_july_11_2020/)
- url: https://www.reddit.com/r/javascript/comments/hp8pa9/showoff_saturday_july_11_2020/
---
Did you find or create something cool this week in javascript? 

Show us here!
## [3][SbAdmin2 build with LitElement, what do you think?](https://www.reddit.com/r/javascript/comments/hpr9tf/sbadmin2_build_with_litelement_what_do_you_think/)
- url: https://slekrem.github.io/sb-admin-2-lit-version/
---

## [4][[AskJS] Which framework do you prefer from scraping data from website? (building a chrome extension)](https://www.reddit.com/r/javascript/comments/hplh9l/askjs_which_framework_do_you_prefer_from_scraping/)
- url: https://www.reddit.com/r/javascript/comments/hplh9l/askjs_which_framework_do_you_prefer_from_scraping/
---
I mostly develop in python. Recently built data scraping tools for a few websites to extract and recalculated users data in a more useful way. I used selenium due to ease of use and ability to use DOM to access the data.

Now I want to rebuild that python data scraper as a chrome extension. Obviously in javascript. Between security issues and javascript libraries I need to choose an architecture.

any tips/suggestions on javascript packages to work with?

Hoping to fast track tool selection before digging too deep into my spare time.  


\[edit: fixed grammar\]  

## [5][Is monorepo for you?](https://www.reddit.com/r/javascript/comments/hpr830/is_monorepo_for_you/)
- url: https://medium.com/@him_jar/is-monorepo-for-you-2020-78cc1717a4f1
---

## [6][The traversal order of object properties in ES6](https://www.reddit.com/r/javascript/comments/hp9nby/the_traversal_order_of_object_properties_in_es6/)
- url: https://2ality.com/2015/10/property-traversal-order-es6.html
---

## [7][An open-source web component for voice-enabled chatbots. Simply provide the URL to an FAQ as an attribute. Heuristic algorithms, tensorflow.js, etc. everything is fully coded in JavaScript.](https://www.reddit.com/r/javascript/comments/hphnha/an_opensource_web_component_for_voiceenabled/)
- url: https://github.com/uihilab/instant-expert
---

## [8][eSlayers part 4 - More user info](https://www.reddit.com/r/javascript/comments/hpo65c/eslayers_part_4_more_user_info/)
- url: https://dev.to/rembrandtreyes/eslayers-part-4-more-user-info-23oe
---

## [9][BloggerJS - script to modify the format of URLs in a Blogger / Blogspot blog, visually creating better navigation. The script "cleans" the URL, removing the date or date from it, as the case may be, as well as the making it a more convenient URL to share](https://www.reddit.com/r/javascript/comments/hpu1nk/bloggerjs_script_to_modify_the_format_of_urls_in/)
- url: https://github.com/jokenox/bloggerjs/
---

## [10][ToastMaker - Open Source JavaScript Library for Showing Toast Notifications on Web Page](https://www.reddit.com/r/javascript/comments/hp85j0/toastmaker_open_source_javascript_library_for/)
- url: https://toastmaker.wirehall.com/
---

## [11][[AskJS] Trick for destructuring re-assignment without parenthesis](https://www.reddit.com/r/javascript/comments/hp91ta/askjs_trick_for_destructuring_reassignment/)
- url: https://www.reddit.com/r/javascript/comments/hp91ta/askjs_trick_for_destructuring_reassignment/
---
For context of what I'm talking about, see either [here](https://stackoverflow.com/questions/27386234/object-destructuring-without-var) or [here on stackoverflow](https://stackoverflow.com/questions/48714689/javascript-re-assign-let-variable-with-destructuring/48714713) (short) or the [notes here on MDN](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment#Assignment_without_declaration) (detailed).

---
*Edit to summarize for the lazy ones: you want to do

     // beginning of function:    
     let { latitude, longitude } = startingCoordinates()
     // ...
     // other parts of function
     // ...
     { latitude, longitude } = updatedCoordinates()
but this is a syntax error on the second assignment; instead you have to do

     ;({ latitude, longitude } = updatedCoordinates())

---

I hate this requirement of parenthesis around an assignment, for me it seems to communicate things that are not true ("this is an expression, we are going to use the return value"). Also it doesn't allow for a semicolon-free coding style (which may be a good thing for some people, but I don't like it), since otherwise the parenthesis might be interpreted as trying to call the previous line as a function. Also it's cumbersome to wrap assignments.

So I've came up with the following trick for reassignment instead. You can simply write

     let {} = { latitude, longitude } = updatedCoordinates()

This works, needs no parenthesis, needs no semicolons, and doesn't pollute the namespace with any more variables. And while it still doesn't communicate the correct thing clearly ("a destructuring reassignment is happening here"), at least it doesn't seem to communicate anything else either (or worst case it communicates "what the heck is this").

That's it, just wanted to let y'all know about this, maybe someone else finds this useful too. And, of course, if someone has an even better solution, I'm all ears.

----

Offtopic: I don't feel like the [AskJS] tag rings very true here as there's no explicit question in my post, but the guide says it's also for "debating best practices", so I guess this post should be ok.
## [12][A New RegExp Engine in SpiderMonkey (Firefox 78+)](https://www.reddit.com/r/javascript/comments/hpbe5i/a_new_regexp_engine_in_spidermonkey_firefox_78/)
- url: https://hacks.mozilla.org/2020/06/a-new-regexp-engine-in-spidermonkey/
---

