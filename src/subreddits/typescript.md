# typescript
## [1][Who's hiring Typescript developers - June](https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/)
- url: https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/
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
## [2][What are the actual reasons some devs don't like Typescript?](https://www.reddit.com/r/typescript/comments/gvrrsu/what_are_the_actual_reasons_some_devs_dont_like/)
- url: https://www.reddit.com/r/typescript/comments/gvrrsu/what_are_the_actual_reasons_some_devs_dont_like/
---
I usually hear "excess code writing/verbosity." Personally I think it covers for not wanting to learn another language. Or maybe the dev is in a field with fairly low complexity or routine work.

I just hooked up part of a CLI program across 5 different files with abstract base classes, interfaces enforcing consistency between a main controller class, and an array of instances of various types. Would have been impossible for it to work on the first go but TS tells you where there are issues every step of the way. I personally think one actually codes faster on the first pass too, when you consider the superior feedback loop.
## [3][Understanding "baseUrl" and "paths" in TypeScript with * glob](https://www.reddit.com/r/typescript/comments/gvq2mh/understanding_baseurl_and_paths_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/gvq2mh/understanding_baseurl_and_paths_in_typescript/
---
I have a monorepo created with yarn workspaces and the following folder structure:

```
├── foo                   (workspace @project/foo)
│   ├── package.json
│   └── src
│       └── index.ts
├── bar                   (workspace @project/bar)
│   ├── package.json
│   └── src
│       └── index.ts
├── package.json          (monorepo root)
└── tsconfig.json         (base tsconfig)
```

And the following settings in `tsconfig.json`:

```ts
{
  "compilerOptions": {
    "baseUrl": ".",
    "module": "commonjs",
    "paths": {
      "@project/foo/*": "./packages/foo/src/*",
      "@project/bar/*": "./packages/bar/src/*"
    },
    ...
  }
}
```

In the `@project/bar` workspace, I want to import modules from `@project/foo`:


```ts
import foo from "@project/foo";
```

But I'm getting the following error:

&gt; Cannot find module '@project/foo' or its corresponding type declarations.ts(2307)

If I remove the `*` symbols from both the keys and the values of the "paths" object, the code compiles. Why is that? How can I keep the `*` glob pattern and make non-relative imports to my local modules?
## [4][Define a type of Record where the key should be of type T](https://www.reddit.com/r/typescript/comments/gvlse5/define_a_type_of_record_where_the_key_should_be/)
- url: https://www.reddit.com/r/typescript/comments/gvlse5/define_a_type_of_record_where_the_key_should_be/
---
I have a `create` method inside a users-service like so:

```
async create(fields: Record&lt;string, any&gt;) {}
```

but the `fields` is always an object of keys where the keys are a partial of `User` keys, so how do I achieve this typing? Something like: `Record&lt;K extends keyof User, V&gt;` ?
## [5][Help! "Require is not defined."](https://www.reddit.com/r/typescript/comments/gvfxaa/help_require_is_not_defined/)
- url: https://www.reddit.com/r/typescript/comments/gvfxaa/help_require_is_not_defined/
---
C++ dev here making some javascript. Thought it would be better to switch to typescript. Stuff compiles to app.js just fine but at runtime I get "require is not defined." I can confirm that require is being created in the js file.

I've scoured the internet. This is apparently a common problem with tons of solutions, none of which work. What do I need to include to make "require" frigging work? Some people said it's because commonjs is included in the tsconfig. Okay cool. What do I do then? Tried changing that. Tried npm installing commonjs. I imagine something this simple should have a simple answer but apparently not...
## [6][Refactor to use RxJS operators - help needed](https://www.reddit.com/r/typescript/comments/gv7qnh/refactor_to_use_rxjs_operators_help_needed/)
- url: https://www.reddit.com/r/typescript/comments/gv7qnh/refactor_to_use_rxjs_operators_help_needed/
---
Hi.    

So I'm struggling with nested subscriptions. It is working, but I know it's an anti-pattern, and I want to refactor it to use RxJS operators. But I'm struggling a bit...    

Here's the code:  


    this.getArray()
        .pipe(takeUntil(this.destroyed$))
        .subscribe((result) =&gt; {
            result.map((value) =&gt; {
                if (value.Id !== null) {
                    this.someService
                        .getById(value.Id)
                        .subscribe(
                            (data) =&gt; (value.Name = data.Name)
                        );
                }
            });
            ... some more logic
        }); 

`getArray()` returns an `Observable&lt;SomeClass[]&gt;` and `getById()` returns an `Observable&lt;SomeClass&gt;`. I guess I want to use `mergeMap` here? So I need to use the `getById()` for each `value.Id` in the array returned from `getArray()`. I'm a bit lost so any help and feedback are appreciated.    

Thank you.
## [7][Styled components and typescript](https://www.reddit.com/r/typescript/comments/gv8jw3/styled_components_and_typescript/)
- url: https://www.reddit.com/r/typescript/comments/gv8jw3/styled_components_and_typescript/
---
Hi anyone knows how to make this code fully typed with Typescript ?

&amp;#x200B;

    interface Size {
      small: number;
      medium: number;
      large: number;
      xLarge: number;
    }
    const size: Size = {
      small: 500,
      medium: 960,
      large: 1140,
      xLarge: 1400,
    };
    
    export const below = Object.keys(size).reduce((acc, label) =&gt; {
      acc[label] = (...args) =&gt; css`
        @media (max-width: ${size[label] / 16}em) {
          ${css(...args)}
        }
      `;
      return acc;
    }, {});
## [8][Peertube v2.2.0 released!](https://www.reddit.com/r/typescript/comments/gv898u/peertube_v220_released/)
- url: https://github.com/Chocobozzz/PeerTube/releases/tag/v2.2.0
---

## [9][Is strict mode worth it?](https://www.reddit.com/r/typescript/comments/guu3a8/is_strict_mode_worth_it/)
- url: https://www.reddit.com/r/typescript/comments/guu3a8/is_strict_mode_worth_it/
---
I just got caught out so I've turned on strictNullChecks, but apart from that is strict mode worth it?
## [10][Can I improve my knowledge and more deeply understand OOP with TypeScript?](https://www.reddit.com/r/typescript/comments/gux7fw/can_i_improve_my_knowledge_and_more_deeply/)
- url: https://www.reddit.com/r/typescript/comments/gux7fw/can_i_improve_my_knowledge_and_more_deeply/
---
A couple years a go I switched careers and became a front end dev. I love it but I have a lot to learn obviously. One of the things I really want to get a better grasp of is OOP. I have a pretty solid JavaScript knowledge but I have read a lot that it isn't truly an OOP language. I would like to learn TS to start using with React and Redux anyway so I was wondering if TS would help me to learn OOP properly, and if so, can you suggest any resources.   


Thank you very much in advance.
## [11][What's the purpose of having a type variable extend any?](https://www.reddit.com/r/typescript/comments/gupooo/whats_the_purpose_of_having_a_type_variable/)
- url: https://www.reddit.com/r/typescript/comments/gupooo/whats_the_purpose_of_having_a_type_variable/
---
I've randomly run into this looking at random TS snippets online or library source codes. Here's an example of it inside [`ts-toolbelt`](https://github.com/pirix-gh/ts-toolbelt/blob/0d5f604/src/Any/Compute.ts#L13).

I thought the following two examples were equivalent; are they not?

    type WithAny&lt;T extends any&gt; = T;
    type WithoutAny&lt;T&gt; = T;
