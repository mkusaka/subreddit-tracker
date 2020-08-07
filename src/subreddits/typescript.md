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
## [2][Announcing TypeScript 4.0 RC](https://www.reddit.com/r/typescript/comments/i51era/announcing_typescript_40_rc/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-4-0-rc/
---

## [3][Typescript compiler errors with typescript are a nightmare and could be improved by showing the *canonical* type.](https://www.reddit.com/r/typescript/comments/i4t8mi/typescript_compiler_errors_with_typescript_are_a/)
- url: https://www.reddit.com/r/typescript/comments/i4t8mi/typescript_compiler_errors_with_typescript_are_a/
---
I'm dealing with this error right now:

    TS2345: Argument of type 'MemoExoticComponent&lt;ForwardRefExoticComponent&lt;IProps &amp; RefAttributes&lt;HTMLElement&gt;&gt;&gt;' is not assignable to parameter of type 'ReactRefComponent&lt;IProps &amp; IRefProps&gt;'.   Type 'ReactElement&lt;any, string | ((props: any) =&gt; ReactElement&lt;any, string | ... | (new (props: any) =&gt; Component&lt;any, any, any&gt;)&gt; | null) | (new (props: any) =&gt; Component&lt;any, any, any&gt;)&gt; | null' is not assignable to type 'Element'.     Type 'null' is not assignable to type 'Element'.

... this is a nightmare.

It forces me to jump between all the types.

Why not just show the above BUT show the canonical type as an 'interface' declaration showing what the type SHOULD be if I were to code it without using type or interface references/inheritance.

So it would just say something like:

    the canonical type is:
    
    interface Foo {
        readonly bar: string;
        ....
    }

This would make my life 100x easier!
## [4][What's the best way to initiate classes from an api call without lying to typescript?](https://www.reddit.com/r/typescript/comments/i5cevm/whats_the_best_way_to_initiate_classes_from_an/)
- url: https://www.reddit.com/r/typescript/comments/i5cevm/whats_the_best_way_to_initiate_classes_from_an/
---
This is just an example to illustrate my issue.

Say I have a class

    class Person {
      
      constuctor(public firstName: string, public lastName: string) {}
 
      getName() {
        return this.firstName + " " + this.lastName;
      }
    }

Now, say that I make an API call to get a Person array, so I might write it like this

    function fetchPeople(url: string) {
      return Axios
        .get&lt;Person[]&gt;(url)
        .then(res =&gt; res.data.map(p =&gt; new Person(p.firstName, p.lastName));
    }

Now, this does work, but it involves lying to Typescript, because of course Typescript would also now accept 

    function fetchNames(url: string) {
      return Axios
        .get&lt;Person[]&gt;(url)
        .then(res =&gt; res.data.map(p =&gt; p.getName());
    }

and then that would give a runtime error.

Another possibility is to type out an interface `IPerson` with properties `firstName: string` and `lastName: string`. The problem I have with this is that in my actual use case this involves hundreds of lines of code to create a load of interfaces that will be used exactly once each. So that seems like a waste of time just for a bit of extra type checking.

The third possibility is to just pass `any` as the type to `Axios.get`. This is even more honest than the first option, but it also does away with my typechecking. 

Does anyone have a sensible solution?
## [5][How to properly organize modules and types?](https://www.reddit.com/r/typescript/comments/i55esz/how_to_properly_organize_modules_and_types/)
- url: https://www.reddit.com/r/typescript/comments/i55esz/how_to_properly_organize_modules_and_types/
---
Hello people!So Im a junior dev and I have a bunch of different Apps at work, but since  they all have the same "business logic", I had a package in the past for that. But now, since I'm adding new platforms, I decided to create a new one using tree-shaking, but I wanted to use the correct pattern/architecture for that. This new package is mainly for React App (they are all frontend)

This package should contain a lot of stuff: Services, Models, Typings, etc (that's why I'm going to use Tree-Shaking), but I don't what's the best way to organize the code and the types. Because theses modules inside the package might have some dependencies between each other, but most of those are types only (which would be removed when compiling to plain js), so for example, let's say I have these 2 interfaces at 2 different files.

&gt;// A/index.ts - interface A at A/index.ts  
&gt;  
&gt;interface A {    key: string; }  
&gt;  
&gt;// B/index.ts - interface B at B/index.ts  
&gt;  
&gt;interface B {    foo(a : A): void;}

the file 'B' only needs the types from file 'A', not actual code and I &gt;&gt;think&lt;&lt; that is something to worry about.So with the goal of having a Tree-Shakable/Type friendly Util's Package for my web app, I came up with the following solution:

Let's say I want to create a User model, I would then have:

// user/getters.ts -&gt; all the implementations of the model

// user/schema.d.ts -&gt;  all the type external type definitions

// user/index.ts -&gt; put everything together

I would then have some sort of index.d.ts at the root of my application that would import all the .d.ts files, making all those types global, thus removing the need to import types across files.

What you guys think? Is this good pattern? Is it bad? Is is horrible? I would gladly take on suggestions :D, or links to reference material, I don't know.

PS: This is the first time I'm asking a code-related question in a long while, so I'm sorry if it's a little confusing \\o
## [6][Using the new variadic tuple types to convert callback functions to promises](https://www.reddit.com/r/typescript/comments/i58b83/using_the_new_variadic_tuple_types_to_convert/)
- url: https://www.reddit.com/r/typescript/comments/i58b83/using_the_new_variadic_tuple_types_to_convert/
---
I've been messing around with the new variadic tuples in typescript 4.0 and was able to write a higher order function that converts the popular (err, result) "callback" format fn's into a type-safe promise.

Thought this was a really cool use-case of what's now possible, so wanted to share:

[Typescript Playground Link](https://www.typescriptlang.org/play?ts=4.0.0-beta#code/FAFwngDgpgBAwgQwDZIEYIMYGsCiAnPGAXhnzwHtCAfGAZxDwEsA7AcxhuYFcUBuUSLEQp02ADwAlAHzEYACigEAXPGRpMuAgBoYeKLR4gVEgJTEZCZmGADoMAGJdmGEI3LMA6oxAALYevEAFRgoAA8QKGYAE1oYSzAAbQBdHWlZOQA6LIQ8VloVBKyMwJ1VEQ1JKSSzIgsrfmAMd3oYCAoAW0ZaRgAzMH9RLFkxAEkQ8MiYuKtk1Kk5HqcMFUdnV3cvXwGKkbmamUzs3PyYEZMVAAUOrqhK8xgAb2AYXSgQLjxmGGYoAHcYK7kTq0KByOR6WjkJAANygOj0ACsoC59o9ni8YItnIcMjk8joFMoygFNHh4fpDMZUU8MRi9O9PiECDAAPyvJEuQl4MxKCFQ2HgilIEAmfi0gC+JnRkv44pswCxLjcXx8UBQ5E2fjUgzk6OYCHaUBU9CYbC06IQrCN3y47VQinNLwwqBUXJUZEoHBtKHJBmFxoYLFYqOh5EYUWAZhpMGdcm4PpgAAMRgBydowAAkD31hvFOktsCzBfFidFwDlwCQbxgqvVAEFaIDgbASG0gV1ev1tRo5LWkBrvFrytgy8A5AhaGBnPJqeiqyAYKw9G8g7IEL8EN4a2r+w2mzc5AAiPvkQ86ADMAEYyy9kIoQHIl1AV2wy5K5GWgA)

Also release it as an NPM package if anyone wants to help improve it (or fix my mistakes)!

[https://github.com/ChuckJonas/ts-promisify-callback](https://github.com/ChuckJonas/ts-promisify-callback)
## [7][Optional but not undefined?](https://www.reddit.com/r/typescript/comments/i54lu8/optional_but_not_undefined/)
- url: https://www.reddit.com/r/typescript/comments/i54lu8/optional_but_not_undefined/
---
Is it possible to create a type where a property is optional but cannot be set to undefined? Or does TypeScript simply not allow this as any omitted property's value is undefined?

Edit: This is the best I could come up with but it doesn't work so I'm concluding that it's not possible.

```  
type Foo = { foo?: string }  
type Bar&lt;T&gt; = { [P in keyof T]: T[P] extends undefined ? never : T[P] }  
type Baz = Bar&lt;Foo&gt;  

const a: Baz = {};  
const b: Baz = { foo: "abc" };  
const c: Baz = { foo: undefined }; // trying to make this not allowed  
```
## [8][How is the job market right now?](https://www.reddit.com/r/typescript/comments/i564s9/how_is_the_job_market_right_now/)
- url: https://www.reddit.com/r/typescript/comments/i564s9/how_is_the_job_market_right_now/
---
Would love to know if the pandemic had much of a negative impact, and also how much working remotely is opening up hiring abroad/generating a more global competition.
## [9][Typescript keeps being unhappy with optional parameter](https://www.reddit.com/r/typescript/comments/i4zp5v/typescript_keeps_being_unhappy_with_optional/)
- url: https://www.reddit.com/r/typescript/comments/i4zp5v/typescript_keeps_being_unhappy_with_optional/
---
I've been working on a domino game in React/Typescript that is based on Mexican Train.  I have a class called `BoneArray` that is an abstract class for any array of dominos (bones as they're called in the game).  This is a class that any array of dominos (the player hand, the bone pile, etc), might inherit from.

I'm having some trouble with a generic method for reduce.  The `bones` property is protected, because I want to limit what users of this class can do with it, but I'm also trying to write wrappers around some of the JavaScript array methods so they can remain available.

The trouble I'm having with reduce is with the optional `initialValue` parameter.

Here's the code:

```
public reduce&lt;T&gt;(
    callbackfn: (
        accum: T, 
        current: Domino, 
        index?: number, 
        array?: Domino[]) =&gt; T, initialValue: T): T {
            return this.bones.reduce(callbackfn, initialValue);
}
```

It works as is.  For a Hand object that inherits from BoneArray `this.score = this.reduce&lt;number&gt;((accum, domino) =&gt; accum + domino.score, 0);` works fine, I get the appropriate scores for the dominoes in my hand.  But I can't figure out how to make `initialValue` an optional parameter.  I've tried `initialValue?: T` and I've tried `initialValue: T | undefined` and I've tried both together, `initialValue?: T | undefined` and none of them work.  I keep getting an error saying there's no matching overload for initialValue to be undefined.  I *can* make the value optional with `initialValue?: T | any` but I feel like `any` is a gross bandaid fix, I have a bit of an aversion to `any`.

How do I make `initialValue` an optional parameter, as it is in the regular JavaScript `Array.prototype.reduce`?
## [10][Watching and compiling packages in a workspace monorepo with typescript?](https://www.reddit.com/r/typescript/comments/i4xe5e/watching_and_compiling_packages_in_a_workspace/)
- url: https://www.reddit.com/r/typescript/comments/i4xe5e/watching_and_compiling_packages_in_a_workspace/
---
So I'm using yarn workspaces to split a project up into "micro packages" for better code separation. This is a backend node project. I've got your standard setup of a `packages/` folder which also includes a `packages/app` which is the entry point to run

    packages/
        app/
        module-1/
        module-2/

Running `yarn run build` correctly builds each module so I can import it, and running `yarn run dev` runs `ts-node-dev src/index.ts` in the app package, all that is gravy

The part that I'm stumbling on is during development; when using `yarn run dev`, changes to the `module-1` package don't trigger a restart. If I want to work on module-1, I have to use `yarn workspace @group/module-1 run build:watch` alongside the dev script

Is there a way to set it up so that running my dev script can start a watch on all the separate packages?
## [11][Decouple TypeScript code with Event Driven architectures. I made a fun nerdy way to do this using IOC and DI. Let me know what you guys think. Is this useful? Star and share if you like!](https://www.reddit.com/r/typescript/comments/i4tkwm/decouple_typescript_code_with_event_driven/)
- url: https://github.com/bobbylite/telephone-ts
---

