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
## [2][W3C Compatible Keycodes for Typescript](https://www.reddit.com/r/typescript/comments/iexbyq/w3c_compatible_keycodes_for_typescript/)
- url: https://github.com/ashubham/w3c-keys
---

## [3][How to best document your TS React project?](https://www.reddit.com/r/typescript/comments/ieqouv/how_to_best_document_your_ts_react_project/)
- url: https://www.reddit.com/r/typescript/comments/ieqouv/how_to_best_document_your_ts_react_project/
---
One thing that doesn't seem to be consistent in my code writing is the way I document my code.

Of course, I know about TSDoc (one extreme) and I know of the principle that good code should be self-explanatory (other extreme).

I'd love to know how you document your TS (React) code and if you know of any good open-source code that inspires your coding style.
## [4][Why does TypeScript still think this number could be undefined?](https://www.reddit.com/r/typescript/comments/if2zc6/why_does_typescript_still_think_this_number_could/)
- url: https://www.reddit.com/r/typescript/comments/if2zc6/why_does_typescript_still_think_this_number_could/
---
I've just spent 20+ minutes playing around with a function like this and I'm lost. Is TypeScript's type inference not able to figure this out, or am I missing something obvious? 

(I'm running pretty low on sleep so it's very possibly the latter)

```typescript
function myFunc(a?: number, b?: number): number {
  if (!a &amp;&amp; !b) throw new Error('No value provided');

  const num = a ?? b;

  return num; // Type 'number | undefined' is not assignable to type 'number'
}
```
## [5][I'm thinking of learning Typescript as my first experience with JS, so I've got a few questions.](https://www.reddit.com/r/typescript/comments/iewlun/im_thinking_of_learning_typescript_as_my_first/)
- url: https://www.reddit.com/r/typescript/comments/iewlun/im_thinking_of_learning_typescript_as_my_first/
---
Do Typescript compilers translate whatever you write into JS that can run anywhere else, or is it a different process?

How different would my experience be if I try to learn JS first before TS? 

I've got a good amount of experience (albeit amateur college-project stuff) with Python and I did Java in high school, how much of that helps in learning TS?
## [6][How to separate internal and external types from TypeScript's compiler?](https://www.reddit.com/r/typescript/comments/iet52c/how_to_separate_internal_and_external_types_from/)
- url: https://www.reddit.com/r/typescript/comments/iet52c/how_to_separate_internal_and_external_types_from/
---
As the title states, how do I go about separating out internal and external types?

For example, let's say I'm writing a simple React component library. This library should export my single component and the prop types it takes. Nothing more.

Now this component is rather big, so it's built up of multiple smaller components. This is no problem, and my bundler will roll them all into a single `index.js` (well, a `index.js`, `index.module.js`, `index.umd.js`, etc.). Problem is, the types for the helper components are still built and exported. 

I'll get a `.d.ts` file defining the prop types of `helperComponent.tsx` when it doesn't even exist anymore as it's been entirely in-lined in `index.js`. And as it's just an internal component its types don't matter one bit to the end user anyways. It's written in TypeScript simply because it increases my development speed, but I don't want to be punished for using TS. Its types have no use being in the bundle. The thing they define no longer exists.

How would I go about ensuring these utterly useless `.d.ts` files aren't being added to my build outputs? Some `.ts`/`.tsx` files are entirely internal and I do not need or want definitions being built for them. Do I just need to manually prune them every time?
## [7][[Robot] Generate Google Slides from Wikipedia content](https://www.reddit.com/r/typescript/comments/ievf4r/robot_generate_google_slides_from_wikipedia/)
- url: https://github.com/vilmacio22/gslides-maker
---

## [8][How do I define `types` in `declaration.d.ts` when I am referring to 2 paths (`next-mdx-remote/render-to-string` &amp; `next-mdx-remote/hydrate`) of 1 module ,i.e, `next-mdx-remote`?](https://www.reddit.com/r/typescript/comments/ieelgi/how_do_i_define_types_in_declarationdts_when_i_am/)
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
## [9][I hate mocking Typescript classes with Jest](https://www.reddit.com/r/typescript/comments/idvumf/i_hate_mocking_typescript_classes_with_jest/)
- url: https://www.reddit.com/r/typescript/comments/idvumf/i_hate_mocking_typescript_classes_with_jest/
---
I find mocking typescript classes with Jest, so painful, does anyone else find this?

I'm trying to mock a single class, and it just turns into so much code.

Are there are good resources / libraries for mocking classes with typescript?
## [10][Rosebox: Be part of the journey](https://www.reddit.com/r/typescript/comments/iejjn0/rosebox_be_part_of_the_journey/)
- url: https://www.reddit.com/r/typescript/comments/iejjn0/rosebox_be_part_of_the_journey/
---
Hey everyone 👋 ! Hope your weekends are going well

I want to announce that Rosebox now has a roadmap project and milestones to reflect the status of releases. You can also submit a request to add type-support for a particular property by creating an issue with the label \`PropType\`. You can even suggest an improvement or report a bug :)

Roadmap: [https://github.com/hugonteifeh/rosebox/projects/12](https://github.com/hugonteifeh/rosebox/projects/12)

Milestones: [https://github.com/hugonteifeh/rosebox/milestones](https://github.com/hugonteifeh/rosebox/milestones)
## [11][How to create a project with express+react+typescript?](https://www.reddit.com/r/typescript/comments/ied04j/how_to_create_a_project_with/)
- url: https://www.reddit.com/r/typescript/comments/ied04j/how_to_create_a_project_with/
---
Hi  
   How to create a project with express+react+typescript? Any official way?  
thanks  
Peter
