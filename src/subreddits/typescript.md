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
## [2][TypeScript Jesus Anders Hejlsberg has been working on some fun new features: Template string types and mapped type `as` clauses](https://www.reddit.com/r/typescript/comments/ikc3h4/typescript_jesus_anders_hejlsberg_has_been/)
- url: https://github.com/microsoft/TypeScript/pull/40336
---

## [3][CyberCode Online: A mmorpg webgame that looks like (disguised as) VS Code](https://www.reddit.com/r/typescript/comments/ik0fxh/cybercode_online_a_mmorpg_webgame_that_looks_like/)
- url: https://i.redd.it/3znwfrhlzck51.png
---

## [4][Auto-generate enum from object keys?](https://www.reddit.com/r/typescript/comments/ikgvq1/autogenerate_enum_from_object_keys/)
- url: https://www.reddit.com/r/typescript/comments/ikgvq1/autogenerate_enum_from_object_keys/
---
I'm using framer motion and the typical little string typos are reminding me that maybe the time spent defining string enums is worth it.

Is there a way to use something like Object.keys to automatically generate a string enum with an object's keys? That would speed and clean things up quite a bit.
## [5][How to model a self referencing tuple array](https://www.reddit.com/r/typescript/comments/ik6jd6/how_to_model_a_self_referencing_tuple_array/)
- url: https://www.reddit.com/r/typescript/comments/ik6jd6/how_to_model_a_self_referencing_tuple_array/
---
I have a data structure that looks like this:

    [
        ["value1", ["value1"]],
        ["value2", ["value1", "value2", "value3"]],
        ["value3", ["value1", "value2", "value3"]],
        ["value4", ["value1", "value2", "value3", "value4"]],
    ]

On the surface it could be described as [string, string[]][], but there is a relationship between those values, the array in the second element of the tuple ([string, string[]]) must only include string values that match the first element of the tuple within the whole structure. 

For example,

This would be valid because "value1" and "value2" are defined as the first elements and the second element only references strings in the first element.

    [
        ["value1", ["value1"]],
        ["value2", ["value1", "value2"]]
    ]

However this would *not* be valid, because of the reference to "value3" which does not occur in any of the first elements:

    [
        ["value1", ["value1"]],
        ["value2", ["value1", "value2", "value3"]]
    ]

Thought I'd see if any of you gurus had an idea on how you'd model that in typescript or if the application will just need to enforce it at runtime. 

P.S. I know there are different ways to go about structuring this data, though in this case the assumption has to be that we cannot change from that structure.
## [6][What are the most crazy anti-TypeScript arguments you’ve heard?](https://www.reddit.com/r/typescript/comments/ik5w7l/what_are_the_most_crazy_antitypescript_arguments/)
- url: https://www.reddit.com/r/typescript/comments/ik5w7l/what_are_the_most_crazy_antitypescript_arguments/
---
“The syntax is ugly”  
“It only pretends to check types”  
“Whats the point if you just use “any””  
“I don’t like Angular for Angular specific things, therefore TypeScript is bad”  
“It‘s too much work to setup to be worth it”  
“Just don’t be a bad programmer and you won’t need TypeScript”  
  
I’ll admit there are some good reasons not to use it in extreme situations, but some people’s reluctance to it is completely absurd.
## [7][Literature](https://www.reddit.com/r/typescript/comments/ik4235/literature/)
- url: https://www.reddit.com/r/typescript/comments/ik4235/literature/
---
Hello guys,  
I signed up for a typescript beginners course before lockdown, unfortunately it hasn't happened yet and will be moved to late October. I'd like to get a little knowledge out front. Any books and/or youtubers you can recommend, learing typescript from scratch?

THANKS! :)
## [8][How do I create a data access layer correctly?](https://www.reddit.com/r/typescript/comments/ijtgcv/how_do_i_create_a_data_access_layer_correctly/)
- url: https://www.reddit.com/r/typescript/comments/ijtgcv/how_do_i_create_a_data_access_layer_correctly/
---
I use TypeORM to work with the database, but I think it would be better to write my own abstraction (correct me if I'm wrong). In DAL, do I need to describe competitive methods for getting data, such as getUser (id), or develop generic methods for any data models, such as findOne&lt;User&gt;(args)?  

Can you recommend resources on this topic?
## [9][Extending HTMLCollectionOf](https://www.reddit.com/r/typescript/comments/ijzofs/extending_htmlcollectionof/)
- url: https://www.reddit.com/r/typescript/comments/ijzofs/extending_htmlcollectionof/
---
So I want to add a method to the `HTMLCollectionOf` class (it's called `toArray`, in case you're curious).

I have augmented the scope as such (everthing is fine so far);

    declare interface HTMLCollectionOf&lt;T extends Element&gt; {
        toArray(): Array&lt;HTMLElement&gt;
        // NOTE: we can discuss the type checking / filtering of Elements vs HTMLElements later :)
    }

Now I want to add an implementation of the method to the class... usually something like `class.methodName = function() {}` or `class.prototype.methodName = function() {}` would work, depending on requirements... but it doesn't for `HTMLCollectionOf`.

&gt;'HTMLCollectionOf' only refers to a type, but is being used as a value here.

That error sorta makes sense but i'm not sure what I should be referencing to add my method.

Thanks in advance as always guys.
## [10][Does using await mean all the catches will just go to the try -catch block?](https://www.reddit.com/r/typescript/comments/ijl6il/does_using_await_mean_all_the_catches_will_just/)
- url: https://www.reddit.com/r/typescript/comments/ijl6il/does_using_await_mean_all_the_catches_will_just/
---
I’m using like 3 awaits. I use them because I need the data of the 3 awaits before doing anything else. I put them all in a try catch block. If any of them fail it’ll go to the catch block right? How do I know which one failed?
## [11][fp-ts equivalent of Scala Option.foreach](https://www.reddit.com/r/typescript/comments/ijlruv/fpts_equivalent_of_scala_optionforeach/)
- url: https://www.reddit.com/r/typescript/comments/ijlruv/fpts_equivalent_of_scala_optionforeach/
---
So Scala has been my big introduction to functional programming and I love it. Now in my TS projects I'm embracing the fp-ts library, which clearly adheres to a different style of functional structures than Scala does.

One big thing I've been looking for is foreach. In Scala if you want to execute code against the contents of an Option (if there's a value present), foreach is how you do it. I don't see any similarly named function in fp-ts tho. I've been going over the docs but I'm a bit unclear on what to use.

Guidance would be appreciated. Thanks.
