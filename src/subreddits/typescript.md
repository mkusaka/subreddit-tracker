# typescript
## [1][Who's hiring Typescript developers - September](https://www.reddit.com/r/typescript/comments/ik9rft/whos_hiring_typescript_developers_september/)
- url: https://www.reddit.com/r/typescript/comments/ik9rft/whos_hiring_typescript_developers_september/
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
## [2][Array lookup in TypeScript should return "T | undefined" and not "T"](https://www.reddit.com/r/typescript/comments/ivlxdk/array_lookup_in_typescript_should_return_t/)
- url: https://www.reddit.com/r/typescript/comments/ivlxdk/array_lookup_in_typescript_should_return_t/
---
Is there a workaround for this? I found an issue report with no progress:

https://github.com/microsoft/TypeScript/issues/11122

&gt; const xs = [1,2,3];
&gt; const x5 = xs[5]; // type is number, expected number | undefined
&gt;
&gt; The type system doesn't know of care about the length of xs, so I would expect any lookup in the array to return T | undefined.

This has bitten me a few times and seems like an obvious hole. The best I've got is to define my own "safeArrayGet(arr, index)" function.
## [3][code.store : A new TypeScript &amp; GraphQL backend as a service in beta.](https://www.reddit.com/r/typescript/comments/iv50px/codestore_a_new_typescript_graphql_backend_as_a/)
- url: https://www.reddit.com/r/typescript/comments/iv50px/codestore_a_new_typescript_graphql_backend_as_a/
---
Hi guys!

We've been working on this project for one year now. [code.store](https://code.store/) is a GraphQL + TypeScript back-end as a service with a focus on micro/macro-services re-use.  
· You can create scalable, serverless, back-end, and re-use them in different projects with a single command line.  
· We wanted to make it as simple as possible, so you can work on business logic, and we try to take care of everything else: database and code generation, automatic versioning, continuous integration, and deployments, backups and data migrations, and scaling.  
We see potential use cases:  
· Back-end for mobile applications, PWA, or any front-end apps.  
· Company-wide catalog of reusable, live microservices, that can be instantly deployed on new projects or applications.

It's our first public release, we would be happy with any feedback if you have time to play with it.

[https://code.store](https://code.store/)
## [4][Any chrome users? How do you deal with the irrelevant google results popping up on the url bar? This extension fixes that problem, but it doesn't work. Can you update it?](https://www.reddit.com/r/typescript/comments/ivq9ky/any_chrome_users_how_do_you_deal_with_the/)
- url: https://chrome.google.com/webstore/detail/dynamichistory/ehkdegpnplleadlmjoaidmjiabocgpok
---

## [5][Short question on p5 for TS](https://www.reddit.com/r/typescript/comments/ivgt1j/short_question_on_p5_for_ts/)
- url: https://www.reddit.com/r/typescript/comments/ivgt1j/short_question_on_p5_for_ts/
---
Hello, I am using [this template](https://github.com/Gaweph/p5-typescript-starter). 

I've changed sketch.ts and changed the module compiler option to ES6 from none as without that I get an error, although I am not sure if doing that was the correct way to rectify the error. What does the module option do?

And part 2: So I've written a short wrapper class around rect() to use as a base class for whatnot in the future.

When I try to import it into sketch.ts, it screws up the setup and draw functions, I get the warning they are not being used and the animation won't load. But when I import it into .d.ts and put it in the declare global it all works fine, why? 

Thank you in advance
## [6][TS and Functional Programming?](https://www.reddit.com/r/typescript/comments/iv9ogt/ts_and_functional_programming/)
- url: https://www.reddit.com/r/typescript/comments/iv9ogt/ts_and_functional_programming/
---
Is it possible to use functional programming concepts with typescript? Such as object composition and function composition? What I mean is to not use classes. I am trying to convince my team (React.js) to use typescript but they're reluctant because they believe it is not possible to use functional programming with TS.
Help me change their minds, or change mine haha
## [7][Dependency injection frameworks](https://www.reddit.com/r/typescript/comments/ivg3o8/dependency_injection_frameworks/)
- url: https://www.reddit.com/r/typescript/comments/ivg3o8/dependency_injection_frameworks/
---
Hi all! I'm currently working on a typescript project and I'm fairly new to the framework. I've been looking around for DI frameworks (inversify, tsyringe etc) but I don't know which one I should choose. I know inversify is by far the most popular but I find tsyringe more elegant. Coming from Java, I'd love to find a typescript framework similar to Spring. Any suggestions? Thank you!
## [8][npm ERR! Failed at the functions@ lint script. npm ERR! This is probably not a problem with npm. There is likely additional logging output above. · Issue #783 · firebase/firebase-functions · GitHub](https://www.reddit.com/r/typescript/comments/ivmocs/npm_err_failed_at_the_functions_lint_script_npm/)
- url: https://github.com/firebase/firebase-functions/issues/783
---

## [9][Extended Cheatsheet about TypeScript](https://www.reddit.com/r/typescript/comments/iupwfs/extended_cheatsheet_about_typescript/)
- url: https://github.com/f0lg0/ts-extended-cheatsheet
---

## [10][is it worth adding TypeScript for a small, solo project?](https://www.reddit.com/r/typescript/comments/iuqcd6/is_it_worth_adding_typescript_for_a_small_solo/)
- url: https://www.reddit.com/r/typescript/comments/iuqcd6/is_it_worth_adding_typescript_for_a_small_solo/
---
I'm a long time Java developer. I've done a fair amount of front-end work as well, mostly copy-past-modify from other existing pages/components, just enough to get stuff working without fully understanding all the crazy js, frameworks, etc.   

But now I'm trying to dive in and understand it all better. Maybe even for a job switch to a primary nodeJS position. I've been following various tutorials and courses like Pluralsight and newline. I'm trying to develop a project that processes tweets, parses them, writes them to a db like Mongo, and then has a GraphQL API to query the db. This is just for learning, I learn best by developing. The newline course looks really good so far, but they introduce TypeScript. I thought Typescript sounded good, coming from my strongly-typed background, but it seems to introduce more problems than it solves, so I'm wondering if I should just remove it and go back to plain js. I would still like to use some kind of objects or classes though. I've seen courses on js classes and modules, so maybe i need to go through more of them.  

[Here's an example, my current blocker, on a StackOverflow post](https://stackoverflow.com/questions/63929850/how-to-use-twits-typescript-definitions-callback-function-not-assignable). I had a basic twitter user timeline request working fine in plain js. But then I tried converting it to Typescript and got stumped.
## [11][Awkward Type Discrimination](https://www.reddit.com/r/typescript/comments/iuk91f/awkward_type_discrimination/)
- url: https://www.reddit.com/r/typescript/comments/iuk91f/awkward_type_discrimination/
---
Consider this:

    type T1 =
        | {s: string}
        | {n: number}
        | void
    ;
    
    type T2 = {x: string} &amp; T1
    
    function fn(arg: T2): number {
        if ('s' in arg) {
            return arg.s.length;
        } else if ('n' in arg) {
            return arg.n;
        } else {
            return arg.x.length;
        }
    }

I really wish I could use `arg.s !== undefined` instead of `'s' in arg`, but TS doesn't accept it for some reason (TS2339).

Anyway, let this be a lesson for you if you're trying to do something like this.

EDIT: I foolishly used `{}` in place of `void` above, this version behaves as expected, AND it excludes incorrect types. Special thanks to u/ministerkosh
