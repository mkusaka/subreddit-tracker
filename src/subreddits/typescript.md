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
## [2][Extended Cheatsheet about TypeScript](https://www.reddit.com/r/typescript/comments/iupwfs/extended_cheatsheet_about_typescript/)
- url: https://github.com/f0lg0/ts-extended-cheatsheet
---

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
## [4][is it worth adding TypeScript for a small, solo project?](https://www.reddit.com/r/typescript/comments/iuqcd6/is_it_worth_adding_typescript_for_a_small_solo/)
- url: https://www.reddit.com/r/typescript/comments/iuqcd6/is_it_worth_adding_typescript_for_a_small_solo/
---
I'm a long time Java developer. I've done a fair amount of front-end work as well, mostly copy-past-modify from other existing pages/components, just enough to get stuff working without fully understanding all the crazy js, frameworks, etc.   

But now I'm trying to dive in and understand it all better. Maybe even for a job switch to a primary nodeJS position. I've been following various tutorials and courses like Pluralsight and newline. I'm trying to develop a project that processes tweets, parses them, writes them to a db like Mongo, and then has a GraphQL API to query the db. This is just for learning, I learn best by developing. The newline course looks really good so far, but they introduce TypeScript. I thought Typescript sounded good, coming from my strongly-typed background, but it seems to introduce more problems than it solves, so I'm wondering if I should just remove it and go back to plain js. I would still like to use some kind of objects or classes though. I've seen courses on js classes and modules, so maybe i need to go through more of them.  

[Here's an example, my current blocker, on a StackOverflow post](https://stackoverflow.com/questions/63929850/how-to-use-twits-typescript-definitions-callback-function-not-assignable). I had a basic twitter user timeline request working fine in plain js. But then I tried converting it to Typescript and got stumped.
## [5][Awkward Type Discrimination](https://www.reddit.com/r/typescript/comments/iuk91f/awkward_type_discrimination/)
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
## [6][How are .ts files run directly?](https://www.reddit.com/r/typescript/comments/iuj8fp/how_are_ts_files_run_directly/)
- url: https://www.reddit.com/r/typescript/comments/iuj8fp/how_are_ts_files_run_directly/
---
I'm using ts-node-dev to run an express.js server and just curious how these types of libraries are able to run .ts files without compilation. Do they have their own interpreter that emulates the JS interpreter, but with full Typescript support?
## [7][Tricking Typescript into Spellchecking](https://www.reddit.com/r/typescript/comments/itzquc/tricking_typescript_into_spellchecking/)
- url: https://github.com/kkuchta/TSpell
---

## [8][Need tests review](https://www.reddit.com/r/typescript/comments/iuorx0/need_tests_review/)
- url: https://www.reddit.com/r/typescript/comments/iuorx0/need_tests_review/
---
Hi! I'm new to testing. I use jest. Currently I'm making a full stack app which will be a very simple CRM. I'm working on a backend now, using express and typeorm. I made 2 custom repositories and I made tests for them. Could you please check out and let me know if I'm doing this in proper way ?

Repo: https://github.com/nemmtor/sale-helper/tree/typeorm/server

Repo tests are inside src/repositories
## [9][Dependency Injection in TypeScript from scratch](https://www.reddit.com/r/typescript/comments/iucio0/dependency_injection_in_typescript_from_scratch/)
- url: https://medium.com/@Sentinelone_tech/dependency-injection-in-typescript-from-scratch-d1a4422043a0
---

## [10][Naming convention for plain objects vs classed objects?](https://www.reddit.com/r/typescript/comments/iu8pqf/naming_convention_for_plain_objects_vs_classed/)
- url: https://www.reddit.com/r/typescript/comments/iu8pqf/naming_convention_for_plain_objects_vs_classed/
---
Let's say that I have an interface representing the attributes of a Point {x: number; y: number;} and I also have a Point class which has these attributes (implements the interface) but also has methods such as shift, rotate, etc.  

Is there any standardized convention on how to name the attributes/properties interface vs. the class itself?  I use this pattern a lot and I've been inconsistent with it.  PointProps for the properties and Point for the class?  Point for the properties and PointEditor for the class?
## [11][I've been using Typescript (migrated after 15 years of Java) for about 1.5 years. Here's what I've learned and how my programming style has changed.](https://www.reddit.com/r/typescript/comments/itmcvp/ive_been_using_typescript_migrated_after_15_years/)
- url: https://www.reddit.com/r/typescript/comments/itmcvp/ive_been_using_typescript_migrated_after_15_years/
---
- I've found I really like inner functions.  I'll put functions within functions to group similar code.  In Java this wasn't possible (at least 8 ... haven't checked lately).  It's better to keep things close together especially if the function isn't reusable.

- I'm calling all interfaces as IFoo rather than just Foo. Interfaces are a bit different than in Java and are far more important.

- Almost all my code is migrating to namespaces and I'm migrating almost entirely from static functions and static classes.  Webpack can't tree shake them and I think this is just cleaner either way.

- I'm moving to a lot more functions, maps, streams, etc.  

- I use undefined a lot more than null.  

- Types and generics work pretty well but every once in a while I get stuck and they can kind of get confusing with edge cases.

Curious how it's changed your programming style.
