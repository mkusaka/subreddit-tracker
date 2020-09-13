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
## [2][From Rust to TypeScript](https://www.reddit.com/r/typescript/comments/irhluf/from_rust_to_typescript/)
- url: https://valand.dev/blog/post/from-rust-to-typescript
---

## [3][Pipe and compose](https://www.reddit.com/r/typescript/comments/irxcjg/pipe_and_compose/)
- url: https://www.reddit.com/r/typescript/comments/irxcjg/pipe_and_compose/
---
How would you type a compose function , pipe function in Typescript  ?

&amp;#x200B;

`const compose = (...fns) =&gt; (x) =&gt; fns.reduceRight((acc, fn) =&gt; fn(acc), x);`

`const pipe = (...fns) =&gt; (x) =&gt; fns.reduce((acc, fn) =&gt; fn(acc), x);`
## [4][Typescript type definition for role property with Identity Server 4 token auth](https://www.reddit.com/r/typescript/comments/irwt8d/typescript_type_definition_for_role_property_with/)
- url: https://www.reddit.com/r/typescript/comments/irwt8d/typescript_type_definition_for_role_property_with/
---
I'm wondering if anyone's encountered this before....but using the token auth implementation of Identity Server 4, the role property can either be a string or an array of strings, coming back, depending on whether or not a user has more than one role:

https://github.com/IdentityServer/IdentityServer4/issues/2468 

I'm wondering if anyone's encountered this and come up with a good workaround, rather than my dirty `Array.isArray()` check on the role property?
## [5][Good resources for learning more deep metaprogramming of typescript?](https://www.reddit.com/r/typescript/comments/irvlhf/good_resources_for_learning_more_deep/)
- url: https://www.reddit.com/r/typescript/comments/irvlhf/good_resources_for_learning_more_deep/
---
I want to improve my typescript skills by learning more complex, advanced techniques with typescript to develop better APIs etc for parts of the app. Any good resources such as sites or blogs?

Edit: I'll read the Typescript handbook. Surprisingly haven't read the whole thing all this time...
## [6][Fully tested apps?](https://www.reddit.com/r/typescript/comments/iriilx/fully_tested_apps/)
- url: https://www.reddit.com/r/typescript/comments/iriilx/fully_tested_apps/
---
I am looking for a React Typescript project with a full suite of unit, integration and e2e testing. I need to get better at unit, integration and e2e testing and how to write succinct functions and components that are easy to test.
## [7][Do I need babel when using typescript?](https://www.reddit.com/r/typescript/comments/irkgdu/do_i_need_babel_when_using_typescript/)
- url: https://www.reddit.com/r/typescript/comments/irkgdu/do_i_need_babel_when_using_typescript/
---
Hello guys, I am pretty new to this ecosystem and I am getting really overwhelmed with all the informations.I am setting up environment for my little project and I was wondering if I even need babel when I am using typescript? If I understand correctly typescript transpiles my modern code into es3 if I set "target" to ES3 in tsconfig, so why would I use babel? Am I missing something?

Edit: my project is using react if that changes anything
## [8][I made a ts template repo on github that has CI/CD actions configured, husky lint-staged, test running, eslint, prettier, and nodemon for people's usage in hacktoberfest.](https://www.reddit.com/r/typescript/comments/ir66mv/i_made_a_ts_template_repo_on_github_that_has_cicd/)
- url: https://github.com/jakehamtexas/ts-lib-starter
---

## [9][Disable any `any`](https://www.reddit.com/r/typescript/comments/irin9e/disable_any_any/)
- url: https://www.reddit.com/r/typescript/comments/irin9e/disable_any_any/
---
Hi there! Is there a way to forcefully replace all `any` with `unknown`, e. g. also for imported libraries? I would like to experiment around with a bit more strict typing in TypeScript, but unfortunately so far I only found `noImplicitAny` and `noExplicitAny` (from eslint).
## [10][What are some of the best engineered simple CRUD application using Typescript, Redux, Axios and implementing localization?](https://www.reddit.com/r/typescript/comments/irc69u/what_are_some_of_the_best_engineered_simple_crud/)
- url: https://www.reddit.com/r/typescript/comments/irc69u/what_are_some_of_the_best_engineered_simple_crud/
---
Trying to find a simple example, and improve upon it. It's the best way to learn. Finding the application that's the most well-made and looking at the best practices.
## [11][Could I get some advice about an error I'm getting?](https://www.reddit.com/r/typescript/comments/irahqn/could_i_get_some_advice_about_an_error_im_getting/)
- url: https://www.reddit.com/r/typescript/comments/irahqn/could_i_get_some_advice_about_an_error_im_getting/
---
I'm writing a small number sequences game for my son to play - I've written a library in typescript with a buddy of mine - all it does is generate number sequences;

[https://github.com/mikeyhogarth/number-sequences/tree/v1.2.1](https://github.com/mikeyhogarth/number-sequences/tree/v1.2.1)

I then started creating a UI to use that library;

[https://github.com/mikeyhogarth/number-sequences-game](https://github.com/mikeyhogarth/number-sequences-game) (it's deployed if anyone wants to take a look, but don't expect much at the moment!).

When I first imported the \`number-sequences\` library, the compiler was throwing an error;

    Could not find a declaration file for module 'number-sequences'. '/Users/mikey/Development/number-sequences-game/node_modules/number-sequences/dist/index.js' implicitly has an 'any' type.
      Try `npm install @types/number-sequences` if it exists or add a new declaration (.d.ts) file containing `declare module 'number-sequences';`ts(7016)

I got around this by adding the following line to \`react-app.d.ts\`

**declare module "number-sequences";**

But that feels like something I shouldn't have had to do - or is it? Is there something I could have done in the original library to prevent having to do this?
