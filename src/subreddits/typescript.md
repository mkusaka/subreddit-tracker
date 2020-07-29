# typescript
## [1][Who's hiring Typescript developers - July](https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/)
- url: https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/
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
## [2][Add property to a type](https://www.reddit.com/r/typescript/comments/hzvwp6/add_property_to_a_type/)
- url: https://www.reddit.com/r/typescript/comments/hzvwp6/add_property_to_a_type/
---
Hi,
I am starting out with typescript and have been stuck on this problem for some time.

I have an object type whose keys are in an enum. So I made it the following way.

```
type MyType = {
    [x in myenum]: string;
}
```

This works. However I need to add one more property to it, say `id:number` I have tried using the extends method as well as directly adding it. But it did not work.

Any suggestions to make it work?
## [3][Is it possible to default export a type , and default export an interface without name ?](https://www.reddit.com/r/typescript/comments/hzzxti/is_it_possible_to_default_export_a_type_and/)
- url: https://www.reddit.com/r/typescript/comments/hzzxti/is_it_possible_to_default_export_a_type_and/
---
Questions in the title .
## [4][Would you guys advise a junior who's put a lot of time into Typescript to hold out for a TS job or pick up whatever comes along?](https://www.reddit.com/r/typescript/comments/hzn9h5/would_you_guys_advise_a_junior_whos_put_a_lot_of/)
- url: https://www.reddit.com/r/typescript/comments/hzn9h5/would_you_guys_advise_a_junior_whos_put_a_lot_of/
---
My situation is I've been studying programming intensely for 3 years now. JS was my first language, then I did a Wordpress plugin in PHP, and now I'm working with Typescript and React. I also know Express and Postgres.

I'm not afraid of a job that is focused on say Rails or Laravel development, I can mostly read Python (dabbled a bit in it early). So I'm confident that i could jump into any Ruby or Python job and be quite capable for an entry level developer. I just would prefer to develop my muscles using TS utility types, is/as, generics that take function arguments etc.

I'm moving to Atlanta which is an easier job market at the bottom that the top 3 where getting any offer at all can be a prize for a junior.

What do you guys think, take anything that comes along or try to get a TS job? I'm confident in my skils and think I could surely get a React or Angular job, I'm less certain about the quantity of jobs that use TS in Atlanta or the wisdom of wanting to work in that specific tech stack.

PS: To clarify, since TS can be on both ends of the stack, I'd be happy if I got to use TS on either end. No complaints about TS/React with a Rails API for example.
## [5][Stuck on "Module not found: LOL.png" for past 2 days, tried everything on stackoverflow](https://www.reddit.com/r/typescript/comments/hzlzqf/stuck_on_module_not_found_lolpng_for_past_2_days/)
- url: https://www.reddit.com/r/typescript/comments/hzlzqf/stuck_on_module_not_found_lolpng_for_past_2_days/
---
**Big shoutout to this community, y'all are there to share the love and frustration of being a developer.** Been stuck the past 2 days with this.

**ERROR**
```
./src/index.tsx
Failed to compile
Module not found: Can't resolve 'images/LOL.png'
```
----

Tried all the stuff on Stackoverflow like:

1. Using ``index.d.ts, or global.d.ts, or react-app-env.d.ts`` trick with ``/// &lt;reference path='./index.d.ts'/&gt;``
```
declare module '*.png';

// or
declare module "*.png" {
  const content: string;
  export default content;
}
```

2. Tried ``webpack file-loader``

3. Delete ``node_modules`` and try again

4. Changing up the import like
```
import * as LOL from 'images/LOL.png'
import * as LOL from '../images/LOL.png'

import LOL from 'images/LOL.png'
import LOL from '../images/LOL.png'
```

5. Adding ``"esModuleInterop": true,`` to ``tsconfig.json``

------

Still getting the same error. I'm sure this is just some easy thing I overloooked.

------

# EDIT: SOLVED thx to sub

I created fresh react app with typescript and then modified ``react-app-env.d.ts`` with ``typeRoots: [ 'src/react-app-env.d.ts' ]``.

```
declare module '*.png';
```
## [6][Why dependency injection(DI)?](https://www.reddit.com/r/typescript/comments/hzi88y/why_dependency_injectiondi/)
- url: https://www.reddit.com/r/typescript/comments/hzi88y/why_dependency_injectiondi/
---
An example using typeorm and typedi;

    // resolver.ts
    import { Resolver } from "type-graphql";
    import { Inject } from "typedi";
    import { User } from "./user.entity";
    import { UserService } from "./user.service";
    
    @Resolver(User)
    export class UserResolver {
      @Inject("UserService")
      public readonly userService: UserService;
    
      // ...
    }
    =================
    // service.ts
    import { Service } from "typedi";
    import { InjectRepository } from "typeorm-typedi-extensions";
    import { Repository } from "typeorm";
    import { User } from "./user.entity";
    
    @Service("UserService")
    export class UserService {
      @InjectRepository(User)
      protected readonly userRepo: Repository&lt;User&gt;;
    
      // ...
    }

And

    // resolver.ts
    import { Resolver } from "type-graphql";
    import { User } from "./user.entity";
    import { UserService } from "./user.service";
    
    @Resolver(User)
    export class UserResolver {
      public readonly userService = new UserService();
    
      // ...
    }
    =================
    // service.ts
    import { getRepository } from "typeorm";
    import { User } from "./user.entity";
    
    export class UserService {
      protected readonly userRepo = getRepository(User);
    
      // ...
    }

Both works but I genuinely want to know what advantage typedi provides here or is it just a fancy way of writing it
## [7][Why do you need Unit Testing in React Native application | Video Tutorial](https://www.reddit.com/r/typescript/comments/hzfqdn/why_do_you_need_unit_testing_in_react_native/)
- url: https://youtu.be/JGseou6bwYs
---

## [8][Can I defined how two types union?](https://www.reddit.com/r/typescript/comments/hzoybx/can_i_defined_how_two_types_union/)
- url: https://www.reddit.com/r/typescript/comments/hzoybx/can_i_defined_how_two_types_union/
---
Say I have this type:

    interface MyType&lt;T, U&gt; {
      foo&lt;V&gt;(cb: (param: T) =&gt; V): MyType&lt;V, U&gt;
      bar&lt;W&gt;(cb: (param: U) =&gt; W): MyType&lt;W, U&gt;
    }

Is it possible for me to tell TypeScript that `MyType&lt;T, never&gt; | MyType&lt;never, U&gt;` is equivalent to `MyType&lt;T, U&gt;`?

## Why I want this:

If I have some value that implicitly gets typed as `MyType&lt;T, never&gt; | MyType&lt;never, U&gt;`, I can't use the value as expected.

Example:

    declare function myFunction1(): MyType&lt;number, never&gt; | MyType&lt;never, string&gt;;
    
    myFunction1()
      .foo((myNumber) =&gt; { /* ... */ })
      .bar((myString) =&gt; { /* ... */});

This results in: `This expression is not callable. Each member of the union type '(&lt;V&gt;(cb: (param: number) =&gt; V) =&gt; MyType&lt;V, never&gt;) | (&lt;V&gt;(cb: (param: never) =&gt; V) =&gt; MyType&lt;V, string&gt;)' has signatures, but none of those signatures are compatible with each other.ts(2349)`

But I want it to work like this:

    declare function myFunction2(): MyType&lt;number, string&gt;;
    
    myFunction2()
      .foo((myNumber) =&gt; { /* ... */ })
      .bar((myString) =&gt; { /* ... */});

[Playground Link](https://www.staging-typescript.org/play?noLib=true&amp;target=7&amp;jsx=0&amp;ts=4.0.0-beta#code/JYOwLgpgTgZghgYwgAgLIE8Aq6AOEA8mANMgKoB8yA3gFDLIwD2j+AauQBQIBGAXMhxxwocALb9MASmQBeSq0n8M2PGxIUA3HWTdh+AOqce-QcLH9S0ucn2K0WXAX3ryWgL40aAEwgIANsIoMACuIAhgwIwgyKLoAGKh4ZEgAIwcdsqO+CDBotzQJCAQAG7QlAA+9ioERaVQJADOYFCgAOaunrEJYRFRaZLaAHRMjBwcsQByuflQVpS09PQA9ABUyIMbyCtL2m4D9IO6UGOxAMrNbXPU2strG4NbO-R7Wt6+AVBBib3RXd-JACZ0koHKocnkCsgmi0QO1Xn8eoD0kMRid0FMIbNZPMbshVutNttdvt1kc0ecYa0rgtFvj7o9iVogA)
## [9][A Mental Model to think in Typescript](https://www.reddit.com/r/typescript/comments/hz93ap/a_mental_model_to_think_in_typescript/)
- url: https://leandrotk.github.io/tk/2020/07/a-mental-model-to-think-in-typescript/
---

## [10][Help with axios method type safety inside function](https://www.reddit.com/r/typescript/comments/hzcyxu/help_with_axios_method_type_safety_inside_function/)
- url: https://www.reddit.com/r/typescript/comments/hzcyxu/help_with_axios_method_type_safety_inside_function/
---
Hello guys, I am trying to type safe a react hook that uses axios for making requests, I am not being able to use axios\[method\] inside the function it gives me the error, I don't get it. It works if I only put the get in the enum though...

https://preview.redd.it/y9ieobjgwld51.png?width=739&amp;format=png&amp;auto=webp&amp;s=bcc38555dc3e8d671c85dd8abb873dcbcfaaa8e1

error is: This expression is not callable.Each member of the union type '(&lt;T = any, R = AxiosResponse&lt;T&gt;&gt;(url: string, data?: any, config?: AxiosRequestConfig | undefined) =&gt; Promise&lt;R&gt;) | (&lt;T = any, R = AxiosResponse&lt;T&gt;&gt;(url: string, config?: AxiosRequestConfig | undefined) =&gt; Promise&lt;...&gt;)' has signatures, but none of those signatures are compatible with each other.

&amp;#x200B;

EDIT: fixed by changing to:

`const response = await axios(url, { method, data: body })`

if anyone knows a workaround to keep it with the other syntax I would appreciate.
## [11][How would you compose a curried function?](https://www.reddit.com/r/typescript/comments/hzfp3d/how_would_you_compose_a_curried_function/)
- url: https://www.reddit.com/r/typescript/comments/hzfp3d/how_would_you_compose_a_curried_function/
---
Hello,

So I'm trying to get to grips with functional programming in Typescript and I figured I'd implement the needed functionality myself to understand best. The goal is to have point free code style and implement the framework for that as I go.

Currently I'm a bit stuck at the following problem:

    const fnAnd = (arg1: any) =&gt; (arg2: any) =&gt; Boolean(arg1) &amp;&amp; Boolean(arg2);
    const fnKey = &lt;T extends object&gt;(key: keyof T) =&gt; (val: T) =&gt; val?.[key];
    const fnSame = &lt;T&gt;(arg1: T) =&gt; (arg2: T) =&gt; arg1 === arg2;
    
    const isZero = fnSame(0); // (arg2: number) =&gt; boolean
    
    interface Vector {
      x: number;
      y: number;
    }
    
    const vecX = fnKey&lt;Vector&gt;('x'); // (val: Vector) =&gt; number
    const vecY = fnKey&lt;Vector&gt;('y'); // (val: Vector) =&gt; number
    
    // how to do this bit point-free?
    const isZeroVector = // (vec: Vector) =&gt; boolean
        (vec: Vector) =&gt;
            fnAnd
                (fnCompose(isZero, vecX)(vec))
                (fnCompose(isZero, vecY)(vec));

What is the fp solution for composing with curried multi-arity functions like fnAnd here?
