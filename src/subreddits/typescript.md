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
## [2][A generic middleware pattern in Typescript](https://www.reddit.com/r/typescript/comments/g1wzl6/a_generic_middleware_pattern_in_typescript/)
- url: https://evertpot.com/generic-middleware/
---

## [3][GitHub Action: Typescript compile failsafe](https://www.reddit.com/r/typescript/comments/g26o6y/github_action_typescript_compile_failsafe/)
- url: https://www.reddit.com/r/typescript/comments/g26o6y/github_action_typescript_compile_failsafe/
---
I got annoyed that I kept forgetting to compile my TS before pushing. So I made a GitHub action to compile it for me and make another commit if I ever forget to compile it. I'm posting a link to the action here if you guys ever need that failsafe in the background for anything.

https://github.com/marketplace/actions/typescript-tsc-build-push
## [4][Is there a way to convert dynamic array to static type assertion?](https://www.reddit.com/r/typescript/comments/g1txjl/is_there_a_way_to_convert_dynamic_array_to_static/)
- url: https://www.reddit.com/r/typescript/comments/g1txjl/is_there_a_way_to_convert_dynamic_array_to_static/
---
is there a way to dynamically generate a const assertion in typescript ?

I know you can create a `readonlyArray` assertion based on a staticly typed array, eg.

     const staticTyped = [1, 2, 3] as const // readonly [1, 2, 3]
     type TypeStatic = typeof staticTyped[number] // 1 | 2 | 3
    

Now that works as expected and elegantly. However, I need to create a large readonly array that's generated dynamically.

    const generateNumRange = (from: number, to: number): number[] =&gt; {
      const arr = []
      for (let i = from; i &lt;= to; i++) {
         arr.push(i)
      }
      return arr
    }
    
    const dynamicTyped = [...generateNumRange(0, 255)] as const // readonly number[]
    type TypeDynamic = typeof dynamicTyped[number] // number (oof)

Now I do understand where the issue lies, however I can't think of a way, that would allow me to return the `dyanmicTyped` as `[0, 1, 2, ..., 255]` rather than `number[]`.

Is there any way to do this / perhaps even more elegant solution to the problem altogether?
## [5][Intro to Game development with Typescript and PixieJS](https://www.reddit.com/r/typescript/comments/g1j93n/intro_to_game_development_with_typescript_and/)
- url: https://nosleepjavascript.com/intro-to-gamedev/
---

## [6][Hook into TSC to perform a check at compile time?](https://www.reddit.com/r/typescript/comments/g1z4bj/hook_into_tsc_to_perform_a_check_at_compile_time/)
- url: https://www.reddit.com/r/typescript/comments/g1z4bj/hook_into_tsc_to_perform_a_check_at_compile_time/
---
I have some strings that form a query in my project, and I can check those strings for validity at run-time when they're being used.

It would be nice to catch any errors in the query strings at compile time. Is there a way I can hook into the compiler so that something I write is run when I execute `tsc`? If I can emit errors along with line information that would be helpful.

If anyone has any resources on creating something like this that would be awesome.
## [7][Due to Covid, my company have asked me to learn Angular/Typescript. Any thoughts on where to start would be much appreciated.](https://www.reddit.com/r/typescript/comments/g20u5m/due_to_covid_my_company_have_asked_me_to_learn/)
- url: https://www.reddit.com/r/typescript/comments/g20u5m/due_to_covid_my_company_have_asked_me_to_learn/
---

## [8][Opinions on strongly typed Reflect.getMetadata or better practices](https://www.reddit.com/r/typescript/comments/g1rlcf/opinions_on_strongly_typed_reflectgetmetadata_or/)
- url: https://www.reddit.com/r/typescript/comments/g1rlcf/opinions_on_strongly_typed_reflectgetmetadata_or/
---
Hi everyone, I'm playing a bit with `reflect-metatada` but something that's annoying me a lot is all the `any's` that are in the type definitions. For example:

    const value = Reflect.getMetadata('key', Object)

How do we know `key` is a valid metadata key? And `value` would be any, what if it could be also `undefined`?

I'm not saying that TypeScript should be automatically able to infer which `keys` and `values` are valid metadata in `Object` (even tho that'd be great) but at least, I think we should be able to do something like:

    interface UserMetadata {
      name: string;
      age: number;
    }
    
    const name = Reflect.getMetadata&lt;UserMetadata&gt;('name', Object);  // name is recognized as a valid key and is a string
    const age = Reflect.getMetadata&lt;UserMetadata&gt;('age', Object);  // age is recognized as a valid key an is a number
    const firstName = Reflect.getMetadata&lt;UserMetadata&gt;('firstName', Object);  // It complains because firstName is not a a valid key

Or am I missing something and this is possible? So far, I tried to improve the default typings for `getMetadata` but even tho I got the `key` validation part right, it's still returning `any`.

    function getMetadata&lt;T, K extends keyof T = keyof T&gt;(metadataKey: K, target: Object): T[K];

What do you think?
## [9][I created Cantara - A CLI tool to create (Serverless) Fullstack React apps with TypeScript](https://www.reddit.com/r/typescript/comments/g1p4m6/i_created_cantara_a_cli_tool_to_create_serverless/)
- url: https://dev.to/scriptify/cantara-a-cli-tool-to-create-serverless-fullstack-react-apps-in-minutes-23g
---

## [10][People who import their (global) types, where do you put them?](https://www.reddit.com/r/typescript/comments/g165mn/people_who_import_their_global_types_where_do_you/)
- url: https://www.reddit.com/r/typescript/comments/g165mn/people_who_import_their_global_types_where_do_you/
---
Having not felt like I've settled on an approach I like yet, I'm curious how other people organize their projects in regard to their type definitions.

Edit: by 'global' I meant more like 'common' or 'shared' or whatever. Bad Ben, Bad!
## [11][Conditional types change behaviour if under type alias](https://www.reddit.com/r/typescript/comments/g14n1c/conditional_types_change_behaviour_if_under_type/)
- url: https://www.reddit.com/r/typescript/comments/g14n1c/conditional_types_change_behaviour_if_under_type/
---
I recently wanted to make sure that an array of strings contained all the strings of a string union.

I found a simple way, but when attempting to generalize the pattern, I came across some behaviour I cannot explain. Perhaps someone has an idea? (disclaimer: I may be missing something obvious)

&amp;#x200B;

    type UnionEquality&lt;A, B&gt; = A extends B ? (B extends A ? true : never) : never;
    type Foo = UnionEquality&lt;"foo" | "bar", "foo"&gt;; // true
    type Foo2 = "foo" extends "foo" | "bar" ? ("foo" | "bar" extends "foo" ? true : never) : never; // never

I'd expect \`Foo\` and \`Foo2\` to evaluate to the same type (`never`), since one is just an expansion of the other.
