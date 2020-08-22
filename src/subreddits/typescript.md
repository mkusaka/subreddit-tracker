# typescript
## [1][Who's hiring Typescript developers - August](https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/)
- url: https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/
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
## [2][How do I define `types` in `declaration.d.ts` when I am referring to 2 paths (`next-mdx-remote/render-to-string` &amp; `next-mdx-remote/hydrate`) of 1 module ,i.e, `next-mdx-remote`?](https://www.reddit.com/r/typescript/comments/ieelgi/how_do_i_define_types_in_declarationdts_when_i_am/)
- url: https://www.reddit.com/r/typescript/comments/ieelgi/how_do_i_define_types_in_declarationdts_when_i_am/
---
I have the following imports:

```ts
import renderToString from 'next-mdx-remote/render-to-string'
import hydrate from 'next-mdx-remote/hydrate'
```

And there are no typings for `next-mdx-remote` so I put the following in `declaration.d.ts`:

```ts
declare module 'next-mdx-remote'
```

But it still errors out.

I have to do the following to get rid of the errors:

```ts
declare module 'next-mdx-remote/render-to-string'
declare module 'next-mdx-remote/hydrate'
```

Is this correct?
## [3][I hate mocking Typescript classes with Jest](https://www.reddit.com/r/typescript/comments/idvumf/i_hate_mocking_typescript_classes_with_jest/)
- url: https://www.reddit.com/r/typescript/comments/idvumf/i_hate_mocking_typescript_classes_with_jest/
---
I find mocking typescript classes with Jest, so painful, does anyone else find this?

I'm trying to mock a single class, and it just turns into so much code.

Are there are good resources / libraries for mocking classes with typescript?
## [4][How to create a project with express+react+typescript?](https://www.reddit.com/r/typescript/comments/ied04j/how_to_create_a_project_with/)
- url: https://www.reddit.com/r/typescript/comments/ied04j/how_to_create_a_project_with/
---
Hi  
   How to create a project with express+react+typescript? Any official way?  
thanks  
Peter
## [5][Announcing TypeScript 4.0](https://www.reddit.com/r/typescript/comments/ideto6/announcing_typescript_40/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-4-0/
---

## [6][How do you guys handle complex input arguments in tests?](https://www.reddit.com/r/typescript/comments/ie17q2/how_do_you_guys_handle_complex_input_arguments_in/)
- url: https://www.reddit.com/r/typescript/comments/ie17q2/how_do_you_guys_handle_complex_input_arguments_in/
---
I’ve been running typescript for a while now and I like the experience about 99% of the time. The last thing is handling complex input arguments in tests. Mocking classes are easy (with ts-mockito), but types are a different story.

Say that I have a helper function that gets called from a dependency, in my case Apollo server. This receives an argument of type A. A includes a lot of mandatory fields, some of which are other complex types. Since A isn’t a class I can’t just instantiate a new one, which (as I understand it) forces me to type it out as an object of type A.

Is there a better way or some shortcut?
## [7][Getting Started with TypeScript : Set up and Intro](https://www.reddit.com/r/typescript/comments/idupum/getting_started_with_typescript_set_up_and_intro/)
- url: https://blogs.rajankalwar.com.np/getting-started-with-typescript-set-up-and-intro-cke2ca7db00f7b8s1doa6817v
---

## [8][Command line tool to organize your Typescript imports](https://www.reddit.com/r/typescript/comments/idqc9r/command_line_tool_to_organize_your_typescript/)
- url: https://www.npmjs.com/package/import-conductor
---

## [9][Deriving sub interface using utility type Omit and Pick . Some issues with refactoring and linting .](https://www.reddit.com/r/typescript/comments/idu0k5/deriving_sub_interface_using_utility_type_omit/)
- url: https://www.reddit.com/r/typescript/comments/idu0k5/deriving_sub_interface_using_utility_type_omit/
---
Here is one example :

    interface mega {
    	a : string,
    	bb : number,
    	c : boolean
    }
    
    type mini1 = Pick&lt;mega,"a"|"bb"/*does give intellisense*/&gt;;
    type mini2 = Omit&lt;mega,"a"/*does not give intellisense*/,"d"/*does not even exist and does not lint error*/&gt;;

1)You get intellisense for the properties for Pick but not for Omit . Why ? 

2)You refactor a to aa and it does not work in both Pick and Omit . Why ?

3)Omit does not lint error if you omit properties that do not exist . Why ?
## [10][Learning typescript](https://www.reddit.com/r/typescript/comments/id8ju3/learning_typescript/)
- url: https://www.reddit.com/r/typescript/comments/id8ju3/learning_typescript/
---
Hi,

I want to learn typescript, not for a server-side or a client-side application but for writing standalone applications. I am not very familiar with javascript, I mainly write in Python. What topics should I learn before typescript? From my understanding, I need at least to learn the basics of node.js and javascript. Is there a good tutorial or course that covers all these topics? 

Thanks!
## [11][Does interface segregation principle has anything to do with preventing collisions of interfaces that are merged?](https://www.reddit.com/r/typescript/comments/ideebf/does_interface_segregation_principle_has_anything/)
- url: https://www.reddit.com/r/typescript/comments/ideebf/does_interface_segregation_principle_has_anything/
---
The question in the title .

Sometimes I find my self having to create bigger interfaces by composing the smaller ones . The issue here is that the smaller ones get edited once in a while and nothing guarantees that they do not collide when they are merged .

As far as I understand ISP has nothing to do with preventing collisions .

Then how does someone prevents the collisions ? Unit testing ? 

I did [a post](https://www.reddit.com/r/typescript/comments/hjif2w/checking_for_interface_collisions/) some time ago , similar to this one and got no real help .
