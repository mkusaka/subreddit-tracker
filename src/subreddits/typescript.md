# typescript
## [1][Who's hiring Typescript developers - March](https://www.reddit.com/r/typescript/comments/fbll8c/whos_hiring_typescript_developers_march/)
- url: https://www.reddit.com/r/typescript/comments/fbll8c/whos_hiring_typescript_developers_march/
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
## [2][Enums, string/number types of interfaces/classes cannot be used as index signatures](https://www.reddit.com/r/typescript/comments/fp5wfk/enums_stringnumber_types_of_interfacesclasses/)
- url: https://www.reddit.com/r/typescript/comments/fp5wfk/enums_stringnumber_types_of_interfacesclasses/
---
I just came across this post [https://github.com/microsoft/TypeScript/issues/37448](https://github.com/microsoft/TypeScript/issues/37448)...

Does anyone have the problem here? Are there any strategies you'd found to overcome it?
## [3][Converting Node.js (benchmarks game) programs to TypeScript](https://www.reddit.com/r/typescript/comments/foxvol/converting_nodejs_benchmarks_game_programs_to/)
- url: https://www.reddit.com/r/typescript/comments/foxvol/converting_nodejs_benchmarks_game_programs_to/
---
I've been trying to convert Node.js (benchmarks game) programs to TypeScript — using *trial and error* and *dumb luck*.

[That has worked](https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/node-typescript.html) for:

* binary-trees
* mandelbrot
* n-body
* pidigits
* regex-redux

But the converted `fannkuch-redux` program was considerably slower, and I failed to convert these programs:

* spectral-norm
* reverse-complement
* k-nucleotide
* fasta

Please help!

Please convert those 5 Node.js programs and [contribute working TypeScript equivalents to the benchmarks game](https://salsa.debian.org/benchmarksgame-team/benchmarksgame/blob/master/CONTRIBUTING.md#contribute-source-code-for-measurement).
## [4][Please suggest advanced generics course](https://www.reddit.com/r/typescript/comments/fp862c/please_suggest_advanced_generics_course/)
- url: https://www.reddit.com/r/typescript/comments/fp862c/please_suggest_advanced_generics_course/
---
I have a Pluralsight subscription but it doesn't have anything deep on topic. Personally I prefer video courses, but any suggestion is appreciated. Thanks!
## [5][Possible to type based on entered params?](https://www.reddit.com/r/typescript/comments/foz02q/possible_to_type_based_on_entered_params/)
- url: https://www.reddit.com/r/typescript/comments/foz02q/possible_to_type_based_on_entered_params/
---
I'm curious if this is possible in TypeScript. I'm creating a `urlHelper` function where a user supplies an entity, method and path. I'm trying to see if I can limit the values in `path` based on what the user put in for `entity`/`method` and what is available in a const. Here's what I've got so far:


    const endpoints = {
      company: {
        get: {
          address: '/company/{id}/address',
          owner: '/company/{id}/owner'
        },
        post: {
          details: '/company/{id}'
        }
      },
      user: {
        get: {
          profile: '/me/profile'
        }
      }
    } as const
    
    type Entity = keyof typeof endpoints // "company" | "user"
    type Method = 'get' | 'post'
    type Path = ????
    
    const urlHelper = (entity: Entity, method: Method, path: Path) =&gt; {
      // Do some stuff
    }


What I'm looking for is if a user puts "company" for `entity`, then "get" for `method`, `path` can only be "address" or "owner". Is this doable?
## [6][Type-safe lenses](https://www.reddit.com/r/typescript/comments/foo055/typesafe_lenses/)
- url: https://github.com/hoppinger/ts-lenses
---

## [7][Best option to make a Libary?](https://www.reddit.com/r/typescript/comments/forltt/best_option_to_make_a_libary/)
- url: https://www.reddit.com/r/typescript/comments/forltt/best_option_to_make_a_libary/
---
Hi everyone!

I'm currently working a project which requires a self made library to make calls to a soap server. This library I need to make needs to be Typescript and needs to be imported using gitlab (no need for npm compatibilities). 

The point is that this library will wrap every header and soap method, so I can just create a connector and let the library/wrapper works.

What is the best option to make a library like that and make it exportable with it's definition types? It is very important that the definitions are available in the project that imports this library. I've checked some guides, but all of them are npm-oriented or jus too simple for my case.

I hope I'm asking in the right place,

thanks!
## [8][Enforce Immutability With Typescript To Boost Refactor Process](https://www.reddit.com/r/typescript/comments/foqn34/enforce_immutability_with_typescript_to_boost/)
- url: https://medium.com//enforce-immutability-with-typescript-to-boost-refactor-process-70055dac5d52?source=friends_link&amp;sk=a750d3062f3fbba9f2d3c156f5050800
---

## [9][MySQL Parser in Typescript](https://www.reddit.com/r/typescript/comments/fojj5j/mysql_parser_in_typescript/)
- url: https://github.com/stevenmiller888/ts-mysql-parser
---

## [10][What's wrong with this type predicate?](https://www.reddit.com/r/typescript/comments/fore56/whats_wrong_with_this_type_predicate/)
- url: https://www.reddit.com/r/typescript/comments/fore56/whats_wrong_with_this_type_predicate/
---
I'm trying to use a [type predicate](https://www.typescriptlang.org/docs/handbook/advanced-types.html#using-type-predicates) to type something, and for some reason, it doesn't appear to be working.

Here's some example code:

    const getIsTypeA = (thingIAmChecking, value): thingIAmChecking is SomeGenericType&lt;TypeA&gt; =&gt; 
      value === 'some other value'

    const myFn = (thingIAmChecking: SomeGenericType&lt;TypeA | TypeB&gt;) =&gt; {
      const value = getValue();
      if (getIsTypeA(thingIAmChecking, value)) {
        thingIAmChecking.onlyTypeAHasThis // ERROR
      }
    }

Is there something I am fundamentally misunderstanding about type predicates? I thought if I return a boolean value, typescript will understand that inside the block of the `if`, the type is `SomeGenericType&lt;TypeA&gt;`.

I can post a more specific example if needed, and thanks in advance for any help with understanding what's going on here!
## [11][`as` vs. `new class` in object literals?](https://www.reddit.com/r/typescript/comments/foua9e/as_vs_new_class_in_object_literals/)
- url: https://www.reddit.com/r/typescript/comments/foua9e/as_vs_new_class_in_object_literals/
---
```tsx
// const initialState = {
//   counter: 0 as number
// }


const initialState = new class {
  counter: number = 0
}

type State = typeof initialState;
```

[View Poll](https://www.reddit.com/poll/foua9e)
