# javascript
## [1][WTF Wednesday (January 01, 2020)](https://www.reddit.com/r/javascript/comments/eijusf/wtf_wednesday_january_01_2020/)
- url: https://www.reddit.com/r/javascript/comments/eijusf/wtf_wednesday_january_01_2020/
---
Post a link to a GitHub repo that you would like to have reviewed, and brace yourself for the comments!
Whether you're a junior wanting your code sharpened or a senior interested in giving some feedback and have some time to spare, 
this is the place.

[Named after this comic](https://davidwalsh.name/demo/code-review.png)
## [2][[AskJS] Share with me your uplifting stories of supporting an AngularJS project in the year 2020](https://www.reddit.com/r/javascript/comments/ejlpbf/askjs_share_with_me_your_uplifting_stories_of/)
- url: https://www.reddit.com/r/javascript/comments/ejlpbf/askjs_share_with_me_your_uplifting_stories_of/
---
Because I’m about to throw myself out a window.
## [3][Vanilla JavaScript and HTML - No frameworks. No libraries. No problem.](https://www.reddit.com/r/javascript/comments/ejouma/vanilla_javascript_and_html_no_frameworks_no/)
- url: https://johnpapa.net/render-html-2/
---

## [4][[AskJS] tagtical - Tagged templates to enhance your text processing experience](https://www.reddit.com/r/javascript/comments/ejocox/askjs_tagtical_tagged_templates_to_enhance_your/)
- url: https://www.reddit.com/r/javascript/comments/ejocox/askjs_tagtical_tagged_templates_to_enhance_your/
---
Hi

I'm working on a [side project](https://github.com/customcommander/tagtical) and I'm looking for early feedback:

* Does it look useful?
* Any suggestions?
* Why would you not use it?
* etc.

Anything that can help me improve and grow ;-) Thanks!

### Use case: pluralisation

Depending on the value of `n` you need to produce a different text, e.g. "There is 1 fox", "There are 10 foxes", etc.

A typical implementation would be:

    `There ${n === 1 ? 'is' : 'are'} ${n} ${n === 1 ? 'fox' : 'foxes'}`
    //=&gt; e.g. "There is 1 fox"
    //=&gt; e.g. "There are 10 foxes"

Simple enough yet already difficult for the brain to parse. The code sort of hides the intent, which is a shame IMHO.

How about this?

    pluralize`There is/are ${n} fox(es)`;
    //=&gt; e.g. "There is 1 fox"
    //=&gt; e.g. "There are 10 foxes"

### Use case: default value

Similar concept; you can embed default values to use when their interpolated siblings are empty. The default values are removed from the string when they're not needed:

    defaults`Hi ${username}/guest, you have ${num}/no new emails`;
    //=&gt; e.g. "Hi John, you have 10 new emails"
    //=&gt; e.g. "Hi John, you have no new emails"
    //=&gt; e.g. "Hi guest, you have no new emails"
## [5][[AskJS] Is the callback pattern dead?](https://www.reddit.com/r/javascript/comments/ejmo38/askjs_is_the_callback_pattern_dead/)
- url: https://www.reddit.com/r/javascript/comments/ejmo38/askjs_is_the_callback_pattern_dead/
---
I heard a lead dev say the callback pattern is dead.  Is that too much of a blanket statement.  
In React adding a callback to a certain setState call can be incredibly useful.

Is there a better non-callback way to do this?
## [6][How My Painstakingly Made Open Source YouTube-Alternative Became Top Trending On Github](https://www.reddit.com/r/javascript/comments/ejbl62/how_my_painstakingly_made_open_source/)
- url: https://medium.com/@NodeTube_org/how-my-painstakingly-made-open-source-youtube-alternative-became-top-trending-on-github-3ef9c6fb1370
---

## [7][How to build a blog with Gatsby.js in a few easy steps](https://www.reddit.com/r/javascript/comments/ej8mjk/how_to_build_a_blog_with_gatsbyjs_in_a_few_easy/)
- url: https://nosleepjavascript.com/gatsby-blog/
---

## [8][JSED - JavaScript Encoded Data. A solution to the local AJAX CORS error](https://www.reddit.com/r/javascript/comments/ejo4to/jsed_javascript_encoded_data_a_solution_to_the/)
- url: https://leeviolander.github.io/JavaScriptEncodedData/
---

## [9][[AskJS] How do I send out hundreds if not thousands of AJAX requests client-side without my app failing?](https://www.reddit.com/r/javascript/comments/ejqw45/askjs_how_do_i_send_out_hundreds_if_not_thousands/)
- url: https://www.reddit.com/r/javascript/comments/ejqw45/askjs_how_do_i_send_out_hundreds_if_not_thousands/
---
Basically I have a huge array of usernames, and for every element in that array  I send a GET request to an API (redditapi in this case) which in turn recursively calls GET requests until there is no more data to GET (usually this is only about 5-7 recursive requests per user at most). However, about half of those requests fail with an error of "err insufficient resources" which I deem is due to the fact I am absolutely blasting out requests. I have tried using Promise.all to no avail, I have tried adding setTimeouts to delay requests being sent out and all this does is wait until each request is done before it sends the next one. I'd like to do something like send out 5-10 requests, wait for them to succeed or fail and then send the next batch of 10 users. Any help is appreciated.
## [10][Chrome Extension That Automatically Skips YouTube Ads](https://www.reddit.com/r/javascript/comments/eiz4sz/chrome_extension_that_automatically_skips_youtube/)
- url: https://github.com/penge/skip-ad
---

## [11][Elementum.js: The simplest framework to work with vanilla WebComponents](https://www.reddit.com/r/javascript/comments/ej5yok/elementumjs_the_simplest_framework_to_work_with/)
- url: https://lucasmenendez.github.io/elementum.js/
---

