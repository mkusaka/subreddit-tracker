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
## [2][Creating an open source team management app. Server and client are both typed and use Prisma, Nexus, Next and Apollo. Would love your feedbacks, and contributions!](https://www.reddit.com/r/typescript/comments/goau56/creating_an_open_source_team_management_app/)
- url: https://github.com/troup-io
---

## [3][Generic not working for types imported via jsdoc comment , in js files , but works via typescript types defined in jsdoc comment . How to fix that ?](https://www.reddit.com/r/typescript/comments/goj3yj/generic_not_working_for_types_imported_via_jsdoc/)
- url: https://i.redd.it/r9j7js42gb051.png
---

## [4][Is a function a valid void type?](https://www.reddit.com/r/typescript/comments/go4ahi/is_a_function_a_valid_void_type/)
- url: https://www.reddit.com/r/typescript/comments/go4ahi/is_a_function_a_valid_void_type/
---
While I was refactoring some code, I changed some currying functions to simple functions. Obviously I started with the interfaces and I realized this strange thing:

    interface FuncInterface {
      fi: (id: number) =&gt; void;
    }
    const func = (id: number) =&gt; () =&gt; {
      console.log(id);
    };
    const implementation: FuncInterface = {
      fi: func,
    };

TS doesn't complain although `func` is returning a function.

Why is this happening?
## [5][Confused by convention of using the word "Intrinsic"](https://www.reddit.com/r/typescript/comments/gohl2r/confused_by_convention_of_using_the_word_intrinsic/)
- url: https://www.reddit.com/r/typescript/comments/gohl2r/confused_by_convention_of_using_the_word_intrinsic/
---
**Problem:**

I understand  what intrinsic means generally, but I'm  wondering if there is some concrete meaning it has in the context of Typescript.  


**Background:**

React has [JSX.IntrinsicElements](https://www.typescriptlang.org/docs/handbook/jsx.html#intrinsic-elements), Which is described as:  


*something intrinsic to the environment (browser)*  
e.g.  native elements (e.g. strings like "div", "span") as opposed to components  


But TypeScript when I  look through the [typescript source](https://github.com/microsoft/TypeScript/blob/master/tests/baselines/reference/intrinsics.js), I can't tell if Intrinsic simply denotes:  
\- inheritance (e.g. "base type"),   or   
\- environment feature (e.g. "div")

  
This is further muddied, as I wonder "Intrinsic" is alluding to something  like an [intrinsic function](https://en.wikipedia.org/wiki/Intrinsic_function)

*a function whose implementation is handled specially by the* [*compiler*](https://en.wikipedia.org/wiki/Compiler)*.*  


Can anyone put me on the right track?
## [6][Best resources for learning typescript](https://www.reddit.com/r/typescript/comments/gnvyd3/best_resources_for_learning_typescript/)
- url: https://www.reddit.com/r/typescript/comments/gnvyd3/best_resources_for_learning_typescript/
---
As a longtime JS dev who now accepts TS as my lord and saviour - what are the best resources for 'deeply' learning TS fundamentals?

Are there any equivalent resources like 'Eloquent Javascript', 'You don't know JS yet' etc. for TS?

Video series/course suggestions also welcome.

Thanks!
## [7][Failing function type check?](https://www.reddit.com/r/typescript/comments/go4kj4/failing_function_type_check/)
- url: https://www.reddit.com/r/typescript/comments/go4kj4/failing_function_type_check/
---
Why the first interface is not checking the optional argument and the second one does?

    interface FuncInterface1 {
      fi1(id?: number): void;
    }
    interface FuncInterface2 {
      fi2: (id?: number) =&gt; void;
    }
    const func = (id: number) =&gt; {
      console.log(id);
    };
    const imp1: FuncInterface1 = {
      fi1: func,     // OK
    };
    const imp2: FuncInterface2 = {
      fi2: func,     // typescript argument error
    };
## [8][Make typescript understand that the esmodule I import has its declaration files somewhere else .](https://www.reddit.com/r/typescript/comments/go5iy7/make_typescript_understand_that_the_esmodule_i/)
- url: https://www.reddit.com/r/typescript/comments/go5iy7/make_typescript_understand_that_the_esmodule_i/
---
For example lets say I have the following project :

    ./myEsModule.js
    ./myFolder
    ./myFolder/types.d.ts

How do I make typescript understand that my esmodule has its declaration files somewhere else , without moving the declaration files and changing their name .
## [9][Can't seem to index a mapped type.](https://www.reddit.com/r/typescript/comments/gnxcq4/cant_seem_to_index_a_mapped_type/)
- url: https://www.reddit.com/r/typescript/comments/gnxcq4/cant_seem_to_index_a_mapped_type/
---
I'm trying to learn advanced types, specifically, mapped types and I'm a little confused with their behavior.

Right now my code looks like this:

    type StatNameNoHp = 'atk' | 'def' | 'spa' | 'spd' | 'spe';
    type StatName = 'hp' | StatNameNoHp;
    type StatsTable&lt;T = number&gt; = {
      [stat in StatName]: T
    };
    const x: StatsTable = { hp: 0, atk: 0, def: 0, spa: 0, spd: 0, spe: 0 };
    for (const stat in x) {
      console.log(x[stat]); //error
    }
    

I get an error saying:

    Element implicitly has an 'any' type because expression of type 'string' can't be used to index type 'StatsTable&lt;number&gt;'.
      No index signature with a parameter of type 'string' was found on type 'StatsTable&lt;number&gt;'

I'm not too sure what's happening here. Referencing x\['hp'\] seems to work perfectly fine. Any help would be appreciated.
## [10][Typing some complex lodash stuff](https://www.reddit.com/r/typescript/comments/gnsseg/typing_some_complex_lodash_stuff/)
- url: https://www.reddit.com/r/typescript/comments/gnsseg/typing_some_complex_lodash_stuff/
---
I was trying to type return value of \_.[defaultsDeep](https://lodash.com/docs/4.17.15#defaultsDeep)

Heres the return type for 2 arguments

`type Primitive = string | number | boolean | null`  
`type DeepDefault&lt;T, K&gt; = T extends Primitive ? T : Omit&lt;T, keyof K&gt; &amp; {`  
  `[Key in keyof K]: Key extends keyof T ? DeepDefault&lt;T[Key], K[Key]&gt; : K[Key]`  
`} ;`  
Heres the link to [playground](https://www.typescriptlang.org/play/?ssl=5&amp;ssc=4&amp;pln=1&amp;pc=1#code/C4TwDgpgBACgTgSwLYOAgbtAvFAzsRAOwHMoAfKQgVyQCMI5ypaB7FgGwgENCnr32AKEGhIUACIQIYSQDMuVdsAA8AFQA0UANIA+KDlVQIAD2ARCAE1yxEKNJigB+KIYBcUAPJ21mgNYQQFlltPQAyKABvQSgoAG0tAKgEXn9A4K0AXXcEkCNTcysoVKCXJwkpGQh5RRVVeICMzS16kAy9bJaMwQBfKABuYVFoFloAKwBGfUjoqC53cfUZ2nd8ImJFmIh3KJiY2QQ4fHcCKggZ7pnZbZm9g6OoAHJEYgALYAfznsHwYbGAJimOygAGN5hsoBYVgRkusZltprs8BBgSxLO55OxcGcYhc9tdEftDsB3A8AO5wVHED44r4iH5QACSHjAYC4U0k0jkCiUyhGE00fL+OkEKMI+CgLBZc0ZzNZgJm0oWS3cACJVjCVeDQVAlTFIVA1dCSJq4fjdoT7icIOCYljRfqMVjzuCrgjzXdiY9nm9qVALhcgA)   
The problem is typing the general case of N parameters   
I know i can just add overloads for up to \~10 parameters for example like so  
`DeepDefaultFor3&lt;T,K,J&gt; = DeepDefault&lt;T, DeepDefault&lt;K,J&gt;&gt;`  
but that doesn't seem right  
Does any one have ideas about hot one would type  \_.[defaultsDeep](https://lodash.com/docs/4.17.15#defaultsDeep) of N parameters  when N is unknown?
## [11][Using JSDoc tags to test functions [Prototype]](https://www.reddit.com/r/typescript/comments/gn6sit/using_jsdoc_tags_to_test_functions_prototype/)
- url: https://i.redd.it/8gqa8evpfvz41.gif
---

