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
## [2][Common JSON patterns in Haskell, Rust and TypeScript](https://www.reddit.com/r/typescript/comments/fvwm3o/common_json_patterns_in_haskell_rust_and/)
- url: https://codetalk.io/posts/2020-04-05-common-json-patterns-in-haskell-rust-and-javascript.html
---

## [3][RFC - TypeScript Mixin Pattern](https://www.reddit.com/r/typescript/comments/fvyklu/rfc_typescript_mixin_pattern/)
- url: https://www.bryntum.com/blog/the-mixin-pattern-in-typescript-all-you-need-to-know-part-2/
---

## [4][what is the diference between ambient declaration, ambient module and ambient module declaration?](https://www.reddit.com/r/typescript/comments/fvxbu6/what_is_the_diference_between_ambient_declaration/)
- url: https://www.reddit.com/r/typescript/comments/fvxbu6/what_is_the_diference_between_ambient_declaration/
---
I am a little confused by these definitions, I saw the official typescript documentation and I don't understand the most adequate definition of `ambient` and how it affects when the word module (`ambient module` and `ambient module declaration`) is used together.
## [5][typing for dynamic objects ?](https://www.reddit.com/r/typescript/comments/fvt2w8/typing_for_dynamic_objects/)
- url: https://www.reddit.com/r/typescript/comments/fvt2w8/typing_for_dynamic_objects/
---
in typescript if you don't define a type for an object, you get intellisense out of the box. 

```
const container = {
  row: { flexDirection: "row" },
  ...
}
```

how ever if you provide typing for the inner objects you lose typing on the outer one 

```
const container: Record&lt;string, ViewStyle&gt; = {
  row: { flexDirection: "row" },
  ...
}
```

is there a way to get both such that when u are using the `container` in the snippet above you get `row` as intellisense
## [6][Some questions regarding defining types with an interface .](https://www.reddit.com/r/typescript/comments/fvi102/some_questions_regarding_defining_types_with_an/)
- url: https://www.reddit.com/r/typescript/comments/fvi102/some_questions_regarding_defining_types_with_an/
---
I have this interface :

    interface myInterface {
    	state?: { [x: string]: any },
    	[x: string]: Function
    }
    
    let myObj : myInterface = {
        state : {
            query : "hello",
            color : "#666666",
        },
        method() {
            console.log(this.query);
        }
    }

1. How do I define that the `Function` is supposed to be non arrow function ?
2. The `[x:string]: Function` will shadow the `state?:{[x:string]:any}` , that means that vs code will except every property to have a type `Function` and will show an error when I will define a `state` property that is of type different from `Function`. How do I prevent the shadowing?
3. Is there any way to make vs code understand that `state` property has type `{[x:string]: infer the type}` . I am asking that because in my example it is clear that `this.query` is supposed to have `string` type but vs code shows `any` .
## [7][Should a union extend one of its subtypes?](https://www.reddit.com/r/typescript/comments/fvj90u/should_a_union_extend_one_of_its_subtypes/)
- url: https://www.reddit.com/r/typescript/comments/fvj90u/should_a_union_extend_one_of_its_subtypes/
---
Hi.

I've been trying to write tests that assert that the type my generics evaluate to are assignable to the type I expect it to be.


I've run into a really weird problem. This is my assertion type (in reality it goes both ways but lets keep it simple here):  

```
type Assert&lt;T, k&gt; = T extends K ? true : never
```

This is how I test that the generic of the type some variable should be a string type:

```
const dummyVariableName: Assert&lt;MyGeneric&lt;typeof someVariable&gt;, string&gt; = true;
dummyVariableName; // No Unused variables check
```

If it compiles that means it works.

However I am getting an issue with Union types.

```
// Expected is true, as I expect since I can assign string to string | undefined
type Expected = Assert&lt;string, string | undefined&gt; 

// WhatTheHell is also true, which I really didn't expect.
type WhatTheHell = Assert&lt;string | undefined, string&gt;;
```

My reasoning is that this shouldn't work because string | undefined cant be assigned to string:

```
const optionalValue: string | undefined = 'some string' as any;
const value: string = optionalValue;
```
Produces:

```
Type 'string | undefined' is not assignable to type 'string'.
  Type 'undefined' is not assignable to type 'string'.
```

Is this a tsconfig issue? This is really weird type behaviour, and I cannot test whether my Generic type evaluates to a string or to string | undefined.

Any insight would be much appreciated.
## [8][Question: Is there a generic name for this type function: (state, fn) =&gt; fn(state)](https://www.reddit.com/r/typescript/comments/fuxfwc/question_is_there_a_generic_name_for_this_type/)
- url: https://www.reddit.com/r/typescript/comments/fuxfwc/question_is_there_a_generic_name_for_this_type/
---
example;

    type Item {
        title: string,
        price: number
    }
    
    type State { 
        items: Item[]
    } 
    
    let state: State = { 
        items: []
    }
    
    function foo&lt;T, U&gt;(state: T, fn: (v: T) =&gt; U): U {}
    
    let totalPrice = foo(state, ({ items }) =&gt; items.reduce((acc, next) =&gt; acc + next.price, 0));
    
    let mostExpensive = foo(state, ({ items }) =&gt; Math.max(...items.map(item =&gt; item.price)));
    
    let firstItem = foo(state, ({ items }) =&gt; items[0]);

How do you guys call such function like \`foo\`.
## [9][Idea: exporting "function bundles"](https://www.reddit.com/r/typescript/comments/fumqcz/idea_exporting_function_bundles/)
- url: https://www.reddit.com/r/typescript/comments/fumqcz/idea_exporting_function_bundles/
---
It would be cool if there was a CLI of some sort that accepted a file and converted each of its exports into a bundled, portable module.

I was wishing something like this existed recently when I was migrating a single API endpoint (written in TypeScript) from an AWS instance (I'll call this Repo #1) to Google Cloud Functions (Repo #2). Even though all the code was written and working in Repo #1, it was a huge chore to move it to Repo #2 because the endpoint code used a few classes and utility functions that were scattered throughout my codebase. All said and done, I ended up copying and pasting a dozen files (mostly containing simple util functions) over to Repo #2 and having to do a lot of manual rewriting of \`import\` statements to account for the changes in project structure. It was very annoying and dumb.

So I was wishing I could have just exported the endpoint from Repo #1 to a bundled, single-file module that I could drop into Repo #2. Ideally it would do tree shaking so it only imports the code from Repo #1 that it actually needs to run. Everything was implemented in strict TypeScript so it was all types and there were no weird global variables or what have you.

Does anyone have a sense of how difficult this would be to do? I don't have a good grasp of the inner workings of the TS engine or parser.
## [10][Variant: More advanced discriminated unions in TypeScript](https://www.reddit.com/r/typescript/comments/fu9wkk/variant_more_advanced_discriminated_unions_in/)
- url: https://github.com/paarthenon/variant
---

## [11][Humpf: Damped Spring motion as a function of time](https://www.reddit.com/r/typescript/comments/fu74v8/humpf_damped_spring_motion_as_a_function_of_time/)
- url: https://github.com/etienne-dldc/humpf
---

