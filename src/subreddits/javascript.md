# javascript
## [1][Showoff Saturday (January 25, 2020)](https://www.reddit.com/r/javascript/comments/etpuy6/showoff_saturday_january_25_2020/)
- url: https://www.reddit.com/r/javascript/comments/etpuy6/showoff_saturday_january_25_2020/
---
Did you find or create something cool this week in javascript? 

Show us here!
## [2][JavaScript libraries are almost never updated once installed](https://www.reddit.com/r/javascript/comments/eus6a0/javascript_libraries_are_almost_never_updated/)
- url: https://blog.cloudflare.com/javascript-libraries-are-almost-never-updated/
---

## [3][CoreJS (used by Babel, Angular) author posted a comment on their repo 16 days ago saying "after some days I'll be in prison", then stops committing to the repo 13 days ago - claims financial problems](https://www.reddit.com/r/javascript/comments/eul7lg/corejs_used_by_babel_angular_author_posted_a/)
- url: https://github.com/zloirock/core-js/issues/747#issuecomment-573318269
---

## [4][EdgeDB 1.0 Alpha 2 with a new JS driver](https://www.reddit.com/r/javascript/comments/euwwt2/edgedb_10_alpha_2_with_a_new_js_driver/)
- url: https://edgedb.com/blog/edgedb-1-0-alpha-2/
---

## [5][Understanding Recursion &amp; Memoization via JavaScript](https://www.reddit.com/r/javascript/comments/eusuhu/understanding_recursion_memoization_via_javascript/)
- url: https://alligator.io/js/understanding-recursion/
---

## [6][Advanced Node.Js: A Hands on Guide to Event Loop, Child Process and Worker Threads in Node.Js](https://www.reddit.com/r/javascript/comments/euigv5/advanced_nodejs_a_hands_on_guide_to_event_loop/)
- url: https://blog.soshace.com/advanced-node-js-a-hands-on-guide-to-event-loop-child-process-and-worker-threads-in-node-js/
---

## [7][[AskJS] best practice for resolving Promises externally](https://www.reddit.com/r/javascript/comments/ev1p51/askjs_best_practice_for_resolving_promises/)
- url: https://www.reddit.com/r/javascript/comments/ev1p51/askjs_best_practice_for_resolving_promises/
---
We've run a few times into problems where we need to resolve/reject a Promise externally.  A common case where this happens is waiting for an external event, say a process to respond to a request of which there might be a few outstanding. I wanted to share a pattern we've been using and see if there are more elegant ways of dealing with this

    let callbacks = {};
    let reqId = 0;
    
    function makeRequest(reqArgs) {
        return new Promise((resolve, reject) =&gt; {
            let myReqId = reqId++;
            callbacks[myReqId] = {resolve, reject, myReqId};
            sendRequest(reqArgs);
        })
    }
    
    // this function is called some time later to process a response
    function processResponse(resp) {
        const reqId = resp.reqId;
        const cb = callbacks[reqId];
        delete callbacks[reqId];  // resolve/reject only once
        if(resp.success) cb.resolve();
        else             cb.reject(resp);
    }
    
    function sendRequest(reqArgs) {
        /// make the request
    }
    
    /// usage 
    makeRequest(someArgs).then(...).catch(...)
## [8][Using JavaScript to Rapidly Process Large JSON and CSV Files on the Command-Line](https://www.reddit.com/r/javascript/comments/eutujz/using_javascript_to_rapidly_process_large_json/)
- url: https://medium.com/@philipp.y.wille/processing-large-data-files-with-pixie-329b92d29a0b
---

## [9][JavaScript obfuscator](https://www.reddit.com/r/javascript/comments/euuqzy/javascript_obfuscator/)
- url: https://obfuscator.io/
---

## [10][Simple Online JavaScript Beautifier Tool](https://www.reddit.com/r/javascript/comments/ev1a5x/simple_online_javascript_beautifier_tool/)
- url: https://jsbeautifier.online/
---

## [11][Google Remarketing: Learn To Retarget Your Customers[FREE EDUONIX COURSE]](https://www.reddit.com/r/javascript/comments/ev3gwp/google_remarketing_learn_to_retarget_your/)
- url: https://www.freetechcourse.xyz/2020/01/google-remarketing-learn-to-retarget.html
---

