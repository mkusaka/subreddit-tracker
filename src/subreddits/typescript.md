# typescript
## [1][Who's hiring Typescript developers - April](https://www.reddit.com/r/typescript/comments/fsojx3/whos_hiring_typescript_developers_april/)
- url: https://www.reddit.com/r/typescript/comments/fsojx3/whos_hiring_typescript_developers_april/
---
The monthly thread for people to post openings at their companies.

* Please state the job location and include the keywords REMOTE, INTERNS and/or VISA when the corresponding sort of candidate is welcome. When remote work is not an option, include ONSITE.

* Please only post if you personally are part of the hiring company—no recruiting firms or job boards **Please report recruiters or job boards**. 

* Only one post per company. 

* If it isn't a household name, explain what your company does. Sell it.

* Please add the company email that applications should be sent to, or the companies application web form/job posting (needless to say this should be on the company website, not a third party site).


Commenters: please don't reply to job posts to complain about something. It's off topic here.

Readers: please only email if you are personally interested in the job. 

Posting top level comments that aren't job postings, [that's a paddlin](https://i.imgur.com/FxMKfnY.jpg)

[Previous Hiring Threads](https://www.reddit.com/r/typescript/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AMonthly%2BHiring%2BThread)
## [2][What's the point of this isArray function?](https://www.reddit.com/r/typescript/comments/ga4a70/whats_the_point_of_this_isarray_function/)
- url: https://www.reddit.com/r/typescript/comments/ga4a70/whats_the_point_of_this_isarray_function/
---
So I'm in a project where one of the developers added this function:

  
`export const isArray = &lt;A&gt;(a: A): boolean =&gt; Array.isArray(a);`  


What's the point of it?  
Why?
## [3][Thoughts on event emitter library design](https://www.reddit.com/r/typescript/comments/ga1svd/thoughts_on_event_emitter_library_design/)
- url: https://www.reddit.com/r/typescript/comments/ga1svd/thoughts_on_event_emitter_library_design/
---
I'm working on a [library](https://github.com/badgateway/ketting), and I will have to implement a feature to subscribe to state changes.

The obvious go-to for this is Node's EventEmitter, but I'm curious if this is still what people will want in a newly developed api.

One idea was to use [for await (...of..)](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/for-await...of) as a way to process events. This has some nice benefits for trickling down exceptions. I don't really know what to do with an exception in an EventEmitter callback.

I guess I'm looking for some opinions. How would you ideally work with a library that exposes events? I need to support multiple subscribers and would like to do something that's close to the 2020 frontend meta-game, but does not require a lock-in into a larger ecosystem/framework.
## [4][Defining a function that has as a parameter a single object that has initial values.](https://www.reddit.com/r/typescript/comments/ga97kq/defining_a_function_that_has_as_a_parameter_a/)
- url: https://www.reddit.com/r/typescript/comments/ga97kq/defining_a_function_that_has_as_a_parameter_a/
---
Lets suppose I have this function :

    function foo({ a = 1, b }) {
    	/*some code*/
    }

I want property `a` to be optional.When no value is provided for `a` then `a` will be equal to `1`, regardless of whether a value was provided to `b`.Correct me if I am wrong but as far as I understand,for that to happen I have to manually code it inside the function body.

Can I at least write something in typescript (or JSDoc comment) so my IDE(vscode) understands what I want to do?More specifically I want when I type `foo` and hover over it,my IDE will show me that `a` has default value `1` and is supposed to be a number if I give a different value.

Is it possible to add also some comments about the functionality of the parameter `a` that will also be visible when I hover over the function (and not just when I hover over the parameter)?

This is what I have done so far :

    /**
     * @param {object} obj
     * @param {number} [obj.a=1]
     * @param {number} obj.b
    */
    function foo({ a = 1, b }) {
    	/*some code*/
    }

In vscode it does not show me initial value. For the comment about each parameter I have to hover over each individual parameter.
## [5][When should you write a type and when should you let it be inferred?](https://www.reddit.com/r/typescript/comments/g9olpy/when_should_you_write_a_type_and_when_should_you/)
- url: https://effectivetypescript.com/2020/04/28/avoid-inferable/
---

## [6][If you had to teach 30+ JS devs TypeScript, what resources would you give them?](https://www.reddit.com/r/typescript/comments/g9j5rb/if_you_had_to_teach_30_js_devs_typescript_what/)
- url: https://www.reddit.com/r/typescript/comments/g9j5rb/if_you_had_to_teach_30_js_devs_typescript_what/
---
Hello!

My company is currently in the architectural stage of building the new version of our web app. For that, our stack is now:

- Next.js
- preact
- Apollo
- mobx

As a way for us to ease our development process and get some assurances about our data structures, we're moving to TypeScript as well. We ran a poll recently, and about half have our JS devs used TypeScript before, and would like to use it again. The other half have heard of it and want to learn. 

For those 15+ JS devs that haven't used TypeScript before, we need some resources for teaching them. For the most part, they have at least 2 years of experience with JS each, and potentially more development experience. 

What resources would you all recommend to get JS devs up to speed as quickly as possible? Right now we front-end masters, which has a decent TypeScript course, and I've also suggested the TS Handbook on the docs site.

Thanks!
## [7][Tribute to Kurt Godel](https://www.reddit.com/r/typescript/comments/g9pxbg/tribute_to_kurt_godel/)
- url: https://medium.com/@damodharanjay/logic-verification-and-god-el-3849d1724275
---

## [8][Using async/await to avoid stack overflow error.](https://www.reddit.com/r/typescript/comments/g9kml7/using_asyncawait_to_avoid_stack_overflow_error/)
- url: https://gist.github.com/Jason5Lee/938b451b62e47f52806f5d24b9820644
---

## [9][Needed some reference project](https://www.reddit.com/r/typescript/comments/g9nvpy/needed_some_reference_project/)
- url: https://www.reddit.com/r/typescript/comments/g9nvpy/needed_some_reference_project/
---
Hi everyone, 

We are writing a library in typescript. So can you guys give us reference to some popular typescript projects to look into?
## [10][Object["unsafe Key"] type?](https://www.reddit.com/r/typescript/comments/g9pqhj/objectunsafe_key_type/)
- url: https://www.reddit.com/r/typescript/comments/g9pqhj/objectunsafe_key_type/
---
    class Example {
        unsafe(data: Array&lt;string&gt;): void {
            //
        }
    }
    new Example().unsafe(JSON.parse('')["whoGivesAF"])

According to my linter and current version of tsc; this is grand. Fine. Got me thinking though...

&amp;#x200B;

Was there a documented discussion from the lang devs about why TS would allow such dirty, typeless assertion as this to be valid? I would (naïvely) presume that it was to allow us to take data from AJAX calls and not have to check types for every property (definitely a potential for runtime exceptions)?

&amp;#x200B;

Would love to get a wee bit more of an insight as to where, specifically, the borders are, on this one.
## [11][Exploring TypeScript from a Business and Software Development Perspective](https://www.reddit.com/r/typescript/comments/g9lg1m/exploring_typescript_from_a_business_and_software/)
- url: https://www.monterail.com/blog/typescript-business-development?utm_medium=social&amp;utm_source=reddit&amp;utm_campaign=typescript
---

