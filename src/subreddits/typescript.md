# typescript
## [1][Who's hiring Typescript developers - May](https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/)
- url: https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/
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
## [2][Deno 1.0 Released](https://www.reddit.com/r/typescript/comments/gja7lo/deno_10_released/)
- url: https://deno.land/v1
---

## [3][Microsoft announces TypeScript 3.9](https://www.reddit.com/r/typescript/comments/gj89g2/microsoft_announces_typescript_39/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-3-9/
---

## [4][tsdotnet/linq - v1.0.0-preview](https://www.reddit.com/r/typescript/comments/gjhumz/tsdotnetlinq_v100preview/)
- url: https://www.reddit.com/r/typescript/comments/gjhumz/tsdotnetlinq_v100preview/
---
[www.npmjs.com/package/@tsdotnet/linq](https://www.npmjs.com/package/@tsdotnet/linq)

[github.com/tsdotnet/linq](https://github.com/tsdotnet/linq)

[tsdotnet.github.io/linq](https://tsdotnet.github.io/linq/)

No readme or examples yet but JS docs are complete.

Check out the core `Linq&lt;T&gt;` class for creating the chain here:[https://github.com/tsdotnet/linq/blob/master/src/linq.ts](https://github.com/tsdotnet/linq/blob/master/src/linq.ts)

Feedback wanted, encouraged, and welcomed.
## [5][Question: How to structure an NPM package with general utilities](https://www.reddit.com/r/typescript/comments/gjmf3y/question_how_to_structure_an_npm_package_with/)
- url: https://stackoverflow.com/q/61798645/39321
---

## [6][Understanding Typescript](https://www.reddit.com/r/typescript/comments/gjmbph/understanding_typescript/)
- url: https://www.youtube.com/watch?v=gs0lpdB9okM
---

## [7][Some articles mention how TypeScript is slower than JavaScript. Why does it matter when it's compiled anyway?](https://www.reddit.com/r/typescript/comments/gjhgku/some_articles_mention_how_typescript_is_slower/)
- url: https://www.reddit.com/r/typescript/comments/gjhgku/some_articles_mention_how_typescript_is_slower/
---
&amp;#x200B;

So I've seen in some places that mention JavaScript as being faster. But when TypeScript will just become JavaScript in production, does this matter at all, and wont they just be the same speed in the end.

&amp;#x200B;

[src](https://jaxenter.com/energy-efficient-programming-languages-137264.html)
## [8][exporting typescript types in a package](https://www.reddit.com/r/typescript/comments/gjlumo/exporting_typescript_types_in_a_package/)
- url: https://www.reddit.com/r/typescript/comments/gjlumo/exporting_typescript_types_in_a_package/
---
***Outline***

I'm writing a package (let's call it myPackage) for shared react components and their corresponding typescript types. I'm using npm symlink to link myPackage to a typescript/react app I'm working on.

the structure of myPackage is as follows:

    multilist
    --index.tsx
    --types.d.ts
    index.tx 

this is the content of the multilist/types.d.ts file

    type ILookup = {
      readonly id: number;
      readonly text: string;
    };
    
    export type MultilistProps = {
      readonly isEditing: boolean;
      readonly id: string;
      readonly labelText: string;
      readonly value: ILookup | undefined;
      readonly responseOptions: ILookup[];
      readonly singleSelect: boolean;
      readonly invalid: boolean;
      readonly onChangeHandler: (value: any) =&gt; void;
    };

this is the content of the index.ts file

    import { Multilist } from './multilist';
    import { MultilistProps as _MultilistProps } from './multilist/types';
    export { Multilist };
    export type MultilistProps = _MultilistProps;

***Problem***

When I import from myPackage, the type MultilistPropsdoesn't retain it's type structure. A mouseover in VSCode looks like: type MulitlistProps: Any.

If I declare and export the MultilistProps type directly in myPackage/index.ts file; the problem is resolved.

If anybody has any idea's I would be very grateful!  


edit: formatting
## [9][I have made a test type to check if there are any collisions in keys of obj1 &amp; obj2 . How to get a more helpful error message?](https://www.reddit.com/r/typescript/comments/gjlkse/i_have_made_a_test_type_to_check_if_there_are_any/)
- url: https://www.reddit.com/r/typescript/comments/gjlkse/i_have_made_a_test_type_to_check_if_there_are_any/
---
First of all is there such a thing as test types in TS or are they just a hack?

With the help of these two posts\[[1](https://www.reddit.com/r/typescript/comments/gj7pah/i_have_troubles_understanding_isnever/)\]\[[2](https://www.reddit.com/r/typescript/comments/gab4ic/how_to_make_ts_or_eslint_or_vscode_warn_me_about/)\] and after getting a little bit more comfortable with typescript I managed to do this :

    type obj1 = {
    	a : string,
    	b : number,
    }
    
    type obj2 = {
    	a : number,
    	c : string
    }
    
    type MustBeNever&lt;T extends never&gt; = T;
    
    type TestNoCollisionsInMixin = MustBeNever&lt;keyof obj1 &amp; keyof obj2&gt;;

to check whether the mixin of two objects (merging one object into another) has any key collision.

You can see that the error message I get is not really helpful . Is there any way to make the error message or the code more helpful ?
## [10][Some useful classes/modules I've retooled for modern times.](https://www.reddit.com/r/typescript/comments/gj7phf/some_useful_classesmodules_ive_retooled_for/)
- url: https://www.reddit.com/r/typescript/comments/gj7phf/some_useful_classesmodules_ive_retooled_for/
---
[https://github.com/tsdotnet](https://github.com/tsdotnet)

Originally called the "TypeScript .NET Library", this has gone through a lot of iterations and now has arrived as individual modules that have minimal dependencies and work well with tree-shaking.

I'm proud to release this now as the foundational classes are rock solid and now have full API docs.

The end goal is to release a new version of LINQ for TypeScript which I've been curating for years now.

# LINQ and Iterables

[https://github.com/tsdotnet/linq/tree/dev/src](https://github.com/tsdotnet/linq/tree/dev/src)

Currently still in development, I realized that extensions weren't gonna be a thing in TS or JS for some time or never.  So switching to a `pipe(filter, filter, filter)` style pattern would not only be more flexible, but would ensure consumers only import what they need.

About 95% of the filters (LINQ methods) are done, and with the intention that they can stand alone or be used with other filters.  Currently working on an API to handle better type inferrence and simplify building a query.

# Queue&lt;T&gt;

[https://github.com/tsdotnet/queue](https://github.com/tsdotnet/queue)

A solid example of a class missing from JS that needs to be simple and fast.

JavaScript arrays are terrible at behaving like queues as `.shift()` is slow.  This queue beats the performance of both arrays and other linked lists.

# LinkedNodeList&lt;T&gt;

[https://github.com/tsdotnet/linked-node-list](https://github.com/tsdotnet/linked-node-list)

Simple, flexible, and reasonably light weight.  Nodes are defined by the consumer and can contain any properties outside of `previous` and `next`.

I'm surprised how often I look to use a linked list to do things that JavaScript arrays are terrible at.

# LinkedList&lt;T&gt;

[https://github.com/tsdotnet/linked-list](https://github.com/tsdotnet/linked-list)

Robust with familiar API.  Uses `LinkedNodeList&lt;T&gt;` under the hood.

TypeScript will likely prevent any slip-ups and inappropriate modifications of the link chain of a `LinkedNodeList&lt;T&gt;.`  `LinkedList&lt;T&gt;` adds more features similar to .NET that allow for working directly with nodes without breaking the chain.

...

Enjoy, hopefully someone finds these fully typed and documented classes useful.  More to come.
## [11][I have troubles understanding IsNever.](https://www.reddit.com/r/typescript/comments/gj7pah/i_have_troubles_understanding_isnever/)
- url: https://www.reddit.com/r/typescript/comments/gj7pah/i_have_troubles_understanding_isnever/
---
Why does this work :

    type IsNever&lt;A&gt; = (A extends never ? never : false) extends never ? true : false;

while this :

    type IsNever&lt;A&gt; = A extends never ? true : false

does not.

Edit : [Here](https://www.reddit.com/r/typescript/comments/gj7pah/i_have_troubles_understanding_isnever/fqleyy5?utm_source=share&amp;utm_medium=web2x) is how I understand the issue is solved .
