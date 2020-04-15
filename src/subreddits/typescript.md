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
## [2][Intro to Game development with Typescript and PixieJS](https://www.reddit.com/r/typescript/comments/g1j93n/intro_to_game_development_with_typescript_and/)
- url: https://nosleepjavascript.com/intro-to-gamedev/
---

## [3][Opinions on strongly typed Reflect.getMetadata or better practices](https://www.reddit.com/r/typescript/comments/g1rlcf/opinions_on_strongly_typed_reflectgetmetadata_or/)
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
## [4][People who import their (global) types, where do you put them?](https://www.reddit.com/r/typescript/comments/g165mn/people_who_import_their_global_types_where_do_you/)
- url: https://www.reddit.com/r/typescript/comments/g165mn/people_who_import_their_global_types_where_do_you/
---
Having not felt like I've settled on an approach I like yet, I'm curious how other people organize their projects in regard to their type definitions.

Edit: by 'global' I meant more like 'common' or 'shared' or whatever. Bad Ben, Bad!
## [5][Conditional types change behaviour if under type alias](https://www.reddit.com/r/typescript/comments/g14n1c/conditional_types_change_behaviour_if_under_type/)
- url: https://www.reddit.com/r/typescript/comments/g14n1c/conditional_types_change_behaviour_if_under_type/
---
I recently wanted to make sure that an array of strings contained all the strings of a string union.

I found a simple way, but when attempting to generalize the pattern, I came across some behaviour I cannot explain. Perhaps someone has an idea? (disclaimer: I may be missing something obvious)

&amp;#x200B;

    type UnionEquality&lt;A, B&gt; = A extends B ? (B extends A ? true : never) : never;
    type Foo = UnionEquality&lt;"foo" | "bar", "foo"&gt;; // true
    type Foo2 = "foo" extends "foo" | "bar" ? ("foo" | "bar" extends "foo" ? true : never) : never; // never

I'd expect \`Foo\` and \`Foo2\` to evaluate to the same type (`never`), since one is just an expansion of the other.
## [6][Open-sourced a complex desktop application written entirely in Typescript: Server, cli, native-like frontend](https://www.reddit.com/r/typescript/comments/g0ri8q/opensourced_a_complex_desktop_application_written/)
- url: https://www.reddit.com/r/typescript/comments/g0ri8q/opensourced_a_complex_desktop_application_written/
---
[https://github.com/deepkit/deepkit](https://github.com/deepkit/deepkit)

Product page: [https://deepkit.ai/](https://deepkit.ai/)

It's a machine learning platform for experiment tracking, execution, and debugging. The GUI is a real-time interface with native-like desktop interface which means all data is streamed using reactive RPC APIs via websockets. There's no REST or polling involved.

Additional new TS open-source byproducts that were created specifically for this product and might be of interest to you:

* [Marshal.ts](https://github.com/marcj/marshal.ts) \- The fastest universal Typescript data serializer and validator incl. Mongo-ORM abstraction
* [Glut.ts](https://github.com/marcj/glut.ts) \- A reactive real-time client-server framework with distributed entity/file abstraction, distributed data exchange, and automatic entity syncing, based on RxJs and websockets
* [angular-desktop-ui](https://github.com/marcj/angular-desktop-ui) \- A GUI library for native-like GUI widget based on Angular 9+
* [npm-local-development](https://github.com/marcj/npm-local-development) A \`npm link\` alternative that actually works with complex setups.
* [typedoc-plugin-lerna-packages](https://github.com/marcj/typedoc-plugin-lerna-packages) \- TS auto doc generator for lerna packages

Some notes:

* Package management with Lerna
* oclif for cli (with custom build), Angular 9+ for GUI, Electron 8+, custom client-server framework using Glut.ts, Mongo DB
* Development time roughly 1.5 year (80h/week, total EUR \~450k), a single developer (for the main product + those OSS libs)
* multi component build based on classic Makefile and webpack

I post it here to give people the opportunity to look how such a complex product is organised, that it is possible to write something like that in TS as a solo developer, and to ask questions about the development process and the like.
## [7][Type-hinted object initializer in parent constructor?](https://www.reddit.com/r/typescript/comments/g13okd/typehinted_object_initializer_in_parent/)
- url: https://www.reddit.com/r/typescript/comments/g13okd/typehinted_object_initializer_in_parent/
---
Is it possible to have a type-hinted object initializer inside a constructor on a parent class? Ex:
```
class Model {
    modelField: string;
    constructor(initializer: SomeSpecialType){
        Object.assign(this, initializer);
    }
}

class User extends Model{
    email: string;
}

const user = new User({
    email: '...' // only get type hints for the email property
})
```

I know it is possible using a static method:

```
class Model{
    someField: string;
    static new&lt;T extends Model&gt;(this: new() =&gt; T, data: Omit&lt;T, keyof Model&gt;) {
        return Object.assign(new this, data);
    }
}

const user = User.new({
    email: 'some email'
})
```

I tried to use the same parameters inside the constructor but TS is telling me that the `this` keyword cannot be used in a `constructor` not generics can be used there.
## [8][Utility function to fix type widening?](https://www.reddit.com/r/typescript/comments/g10np6/utility_function_to_fix_type_widening/)
- url: https://www.reddit.com/r/typescript/comments/g10np6/utility_function_to_fix_type_widening/
---
I (or rather, a library I'm using) have a lot of interfaces for various options where a lot is optional.

```
interface Options {
  a?: number;
  b?: number;
  c?: number;
}
```

I create objects of the options I'm going to use, but also need to refer to these myself. I want the type checking when making the object, but after that I don't want things I've not specified to be available, or things I _have_ specified to be possibly undefined, like here:

```
const o: Options = { a: 1 } // Good: Object is type checked
o.a.toFixed(0)              // Bad: `a` is possibly undefined
```

Could just skip the type, but then there's no type checking of the options:

```
const o = { a: 1, d: 2 } // Bad: `d` isn't an option
o.a.toFixed(0)           // Good: `a` is defined
```

So, I created a helper function:

```
const createOptions = &lt;O extends Options&gt;(options: O): O =&gt; options
const o = createOptions({ a: 1 }) // Good: Object is type checked
o.a.toFixed(0)                    // Good: `a` is defined
```

This works, but having to create a function like this for every option type is getting annoying and messy. Is it possible to create a single generic helper function for this?

---

My first newb attempt was the following, but here Typescript requires me to supply _2_ types, rather than just the _1_ (`Options`) that should be necessary.

```
const create = &lt;T, U extends T&gt;(obj: U): U =&gt; obj
const o = create&lt;Options&gt;({ a: 1 }) // Bad: Typescript wants me to specify U
```

How can I write this identity function so I only need to specify `T`, and have Typescript infer `U` itself from the object I pass in?


[Playground Link](https://www.typescriptlang.org/play/#code/JYOwLgpgTgZghgYwgAgPIAczAPYgM7IDeAsAFDLJwD8AXMiAK4C2ARtANxkUu33NtRO5ZAl6NW0MgF8yZAPRzkAITgATOgAVsePMBYAbAJ7IGIVRBigIqsglx4wybAAY6GLPeQBeIpToBGZBlSFwA6OFCwbAAxYAAPawAKZwBKWVIFZTU6ADlsZDBDdBQEAAsIBABrUABzW3tHbECfQj9kQOCm8MiY+KTU9MyAcWxsVQByYDxSuiUGRxAIawL8hCgIOEgnRacYAvK8FBhsKGQIADdoY0Li+vxHNY3Idxx8b2QAHlQzuMgzAhe9gAfIlsJhXng3Ck3N4gU5wfY7g4nAAmd6PTYQQH4RKtOABIJpEIo7pRWIJVTJInyRQqdTIABKEAAjgxgOsCEwUFFkHhighgDBjGibhA8EiHutMe8PgAVAA0yAAqj8-qoCLKQdgWAArOhK6HK2FOXUSpwAZnRUsgXwR+BBeIJUhSQmw5tJvQpVLIQA)
## [9][migrating code from javascript](https://www.reddit.com/r/typescript/comments/g0zyqv/migrating_code_from_javascript/)
- url: https://www.reddit.com/r/typescript/comments/g0zyqv/migrating_code_from_javascript/
---
does all my logic, classes, async/await code stay and work the same? Im just concerned that some of the stuff I wrote in javascript may not work in ts.

thanks!!
## [10][Are Tests Necessary in TypeScript?](https://www.reddit.com/r/typescript/comments/g0qpxe/are_tests_necessary_in_typescript/)
- url: https://www.executeprogram.com/blog/are-tests-necessary-in-typescript
---

## [11][Typescript type cast](https://www.reddit.com/r/typescript/comments/g0ldya/typescript_type_cast/)
- url: https://www.reddit.com/r/typescript/comments/g0ldya/typescript_type_cast/
---
I could do 
```
const test = &lt;ILalamove&gt;partners[claimedPartner];
const riderInfos = await test.GetDriver(claimedJob.id,
                                        claimedJob.riderId,
                                        job.isTest,
                                    );
```
But not
```
const riderInfos = await &lt;ILalamove&gt;partners[claimedPartner].GetDriver(claimedJob.id,
                                        claimedJob.riderId,
                                        job.isTest,
                                    );
```
nor
```
const riderInfos = await partners[claimedPartner] as &lt;ILalamove&gt; 
                  .GetDriver(claimedJob.id,claimedJob.riderId,job.isTest,);
```
Do we have better way to do this casting?

Thanks
