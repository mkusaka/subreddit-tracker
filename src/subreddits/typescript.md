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
## [2][Foal TS - April release (version 1.7) - File uploads to local or Cloud storage (AWS S3), OpenAPI spec generation, safe config access](https://www.reddit.com/r/typescript/comments/fx40ui/foal_ts_april_release_version_17_file_uploads_to/)
- url: https://www.reddit.com/r/typescript/comments/fx40ui/foal_ts_april_release_version_17_file_uploads_to/
---
Foal TS framework version 1.7 is officially released!

[Foal TS - April release \(version 1.7\)](https://preview.redd.it/bpntqk11lkr41.png?width=1020&amp;format=png&amp;auto=webp&amp;s=7de8eaa02ac49ea52d7368e494930e1c9118afaf)

**File uploads to local and Cloud storage**, better error description, safer config access, logging improved, you can discover the features of this version here: [https://github.com/FoalTS/foal/releases/tag/v1.7.0](https://github.com/FoalTS/foal/releases/tag/v1.7.0)

The **big feature** of this release is the new `@ValidateMultipartFormDataBody` hook which allows you to **validate and parse multipart/form-data requests** (used for file uploads).

Let’s say we have a profile page from which the user can upload their profile picture as well as their pseudonym. The image should not exceed a certain size and the pseudonym should be 30-characters long.

We want to upload the files in streaming (for performance reason) to our local server file system when coding the application.

Then, when deploying the application to production, we want to use AWS and store our images in an S3 bucket.

And finally, we would like to generate the OpenAPI specification and use Swagger UI to test our application.

=&gt; This new version allows you to achieve this with just a few lines of code and configuration.

More documentation here: [https://foalts.gitbook.io/docs/topic-guides/file-system/upload-and-download-files](https://foalts.gitbook.io/docs/topic-guides/file-system/upload-and-download-files)

[Page at http:\/\/localhost:3001\/swagger](https://preview.redd.it/tw35xed3lkr41.png?width=1446&amp;format=png&amp;auto=webp&amp;s=b9084c176684da1e120daa1ebc07e98ef4d1accf)

[api.controller.ts](https://preview.redd.it/dzn0pmd3lkr41.png?width=588&amp;format=png&amp;auto=webp&amp;s=e246d7c948b0dd098d16d28c804afdcba7953a82)

[openapi.controller.ts](https://preview.redd.it/y7eo8fd3lkr41.png?width=430&amp;format=png&amp;auto=webp&amp;s=d27d45618d0fdb18c8a9bc112c7967f0c47446c0)

[app.controller.ts \(the root controller of the app\)](https://preview.redd.it/xbp0ajd3lkr41.png?width=482&amp;format=png&amp;auto=webp&amp;s=2e82cf41d6da5235de2bb87c82e3da9f8b3cf70a)

[the configuration file config\/development.yml](https://preview.redd.it/x2ah4hd3lkr41.png?width=217&amp;format=png&amp;auto=webp&amp;s=1dbaed210f1d93fd3a909c5609ab82b67fecacaa)

[the configuration file config\/production.yml](https://preview.redd.it/4d5bzld3lkr41.png?width=203&amp;format=png&amp;auto=webp&amp;s=6d2804d009e8ebf59d7056877050023e9a3934b3)

Foal, in a few words, it's a Node.js framework:

&amp;#x200B;

* written in TypeScript
* provided with batteries included (Auth, OpenAPI, GraphQL, ORM, CLI, scripts, file storage)
* and with a simple and intuitive architecture (no magic, no over-engineering).  
 

And the must: it has more than 11,000 lines of documentation.

[https://foalts.org](https://foalts.org/)

[https://github.com/FoalTS/foal](https://github.com/FoalTS/foal)

[https://foalts.gitbook.io/docs/](https://foalts.gitbook.io/docs/)

\#TypeScript #JavaScript #NodeJS #FoalTS #AWS #S3
## [3][Type error while binding a value](https://www.reddit.com/r/typescript/comments/fx6bip/type_error_while_binding_a_value/)
- url: https://www.reddit.com/r/typescript/comments/fx6bip/type_error_while_binding_a_value/
---
Begginer question here..

const item = ({ index, value }) =&gt; \`${index}:${value}\`; 

Error in VS: Binding element 'index' implicitly has an 'any' type.ts(7031)

where do I specificy the type for the above line?
## [4][Dealing with types at scale](https://www.reddit.com/r/typescript/comments/fx5nrx/dealing_with_types_at_scale/)
- url: https://www.reddit.com/r/typescript/comments/fx5nrx/dealing_with_types_at_scale/
---
During development, we tend to consume types, i.e:

```ts
interface Foo {
  x: string
  y: number
  z: boolean
  //... more properties
}
```

Then in our functions, we just "re-use" that type `Foo`

```ts
type Fn = (f: Foo) =&gt; ReturnType
```

But what happens when `Foo` is an object containing a bunch of properties where let's say a function `SimpleFn` does not use the whole set of properties

```ts
type SimpleFn = (f: Pick&lt;Foo, 'x' | 'y'&gt;) =&gt; ReturnType
```

later we have a 

```ts
type SimpleFn2 = (g: Pick&lt;Foo, 'y' | 'z'&gt;) =&gt; ReturnType 
```

What option do you prefer?

a)
```ts
type SimpleFn = (f: Foo) =&gt; ReturnType
type SimpleFn2 = (g: Foo) =&gt; ReturnType
```

b)
```
type SimpleFn = (f: Pick&lt;Foo, 'x' | 'y'&gt;) =&gt; ReturnType
type SimpleFn2 = (g: Pick&lt;Foo, 'y' | 'z'&gt;) =&gt; ReturnType 
```

c) Try to have a very well defined domain where I can group all my types in a file and then reuse them
## [5][Announcing Chronograph 1.0.0 - An open-source reactive computational engine, implemented in TypeScript](https://www.reddit.com/r/typescript/comments/fwp8ga/announcing_chronograph_100_an_opensource_reactive/)
- url: https://www.bryntum.com/blog/announcing-chronograph-1-0-0/
---

## [6][Can't solve this TypeScript error! :(](https://www.reddit.com/r/typescript/comments/fx2lha/cant_solve_this_typescript_error/)
- url: https://www.reddit.com/r/typescript/comments/fx2lha/cant_solve_this_typescript_error/
---
Was trying to work with a React app, but this Typescript error happens. Since I'm a noob at TS, not sure what's happening. Any clue, pros? :( 

https://preview.redd.it/vly86gw8zjr41.png?width=1053&amp;format=png&amp;auto=webp&amp;s=768d6e7d0d78be2634dd91747a1918354b9043be
## [7][Why do these typings stop working when they are used inside an indexed type?](https://www.reddit.com/r/typescript/comments/fwv4ky/why_do_these_typings_stop_working_when_they_are/)
- url: https://www.reddit.com/r/typescript/comments/fwv4ky/why_do_these_typings_stop_working_when_they_are/
---

    type Thing = { name: string; };
    type Things = Array&lt;Thing&gt;;
    type ThingOrThings = Thing | Things;
    
    var thing: ThingOrThings = { name: "sam" };
    thing.name = "pat"
    // ts correctly infers that thing is a Thing that has .name
    var things: ThingOrThings = [{ name: "sam" },{ name: "ash" }];
    things[0].name = "pat"
    things.map(x =&gt; x)
    // ts correctly infers that things is a Things, an array I can index or .map()
    
    interface ThingMap { [key: string]: ThingOrThings; }
    var thingmap: ThingMap = { thing: { name: "sam" }, things: [{ name: "sam" },{ name: "ash" }] };
    thingmap.thing.name = "pat"; // ERROR: Property 'name' does not exist on type 'Things'. (2339)
    // How is thingmap.thing here functionally different from thing above? they have the same contents and the same type
    thingmap.things[0].name = "pat"; // ERROR: Property '0' does not exist on type 'ThingOrThings'. (7053)
    // How is thingmap.things here functionally different from things above? they have the same contents and the same type
    
Most of the question is in the comments above. When I use ThingOrThings directly as the type of an object, Typescript is quite capable of figuring out whether the object is a Thing or a Things, then allowing or disallowing me to manipulate it as such. However, when I put that exact same type into an indexed interface, and then put the exact same contents into two objects inside that interface, TS suddenly can't figure out how to handle the objects.

[Typescript Playground link](https://www.typescriptlang.org/play/?ssl=1&amp;ssc=1&amp;pln=19&amp;pc=1#code/C4TwDgpgBAKgFgSwHYHMoF4oG8pIIYC2EAXFAM7ABOyKA3FAL60BQoksiqZGUAgpZTwgAPPBoA+Fm2hjUAeUqyU3TEqgAfDjTItmANzyUowTilJKFSldlyESUAERlCDxlNMA6fER4OweYAdmAHpg424AYwB7AQgI4AAbEChkADMISm4TAONTFO48LVRcnLg8bi87fUNc7XNTS1NrAG0cb3snF0YAGja7UgdyuFcGAF13bWaABlHKn0w-AKCTbQ8CPDAACgAPDHEobYBKELDgSJjKOMTktIyssuBarnyoQqtu16RXgSEoAEkoBE8F9kAATCC7GJQNYbTbHZjIYAZVJ4CIyUwAWQ2NmaAGsICBSBRqKhRvUaI1tPQGNUjCtUOswOTUFiwDwcPSzDZ2gNnAQRh9OWRSK1bEReV0GL0xR0hiNRm5WKZGR5OXNoAt-IEWKEoAAJKIAdxenJVnKgcAy0FSAFckPEEFF8AkklBQQhUulLkhHqlKFECE80HgAEZRPQQAD8uQJFrwEZj5DsgKdSJ9BSQoMTfOg0iVNDNTWms3avi1Dh1YQNxoQ9wLG1VTQtVqgtvtwEdztd7s9Vp9rf9gaFrzDEejJljZQTE6TPmiPog6c+WZnOeM4AgzCAA
)
## [8][Building Vue Enterprise Application: Part 3. The Store](https://www.reddit.com/r/typescript/comments/fwkytl/building_vue_enterprise_application_part_3_the/)
- url: https://medium.com//building-vue-enterprise-application-part-3-the-store-dbda0e4bb117?source=friends_link&amp;sk=e394ce16a71eef6187eb578e3d0783fe
---

## [9][Is it possible to type this in TypeScript?](https://www.reddit.com/r/typescript/comments/fwi9h8/is_it_possible_to_type_this_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/fwi9h8/is_it_possible_to_type_this_in_typescript/
---
Hey everyone, having troubles typing something in TypeScript, I understand the error more or less but I don't know how I could do it differently.

Playground link:

https://www.typescriptlang.org/play/?ssl=26&amp;ssc=38&amp;pln=1&amp;pc=1#code/JYWwDg9gTgLgBAbzgVwM4FMDKMCGN1wC+cAZlBCHAORTo4DGMVA3AFCsnIB2jwEXKDAGEKYADboAHtjzoAPABU4U-FwAmqOACV09aGrmoYUYFwDmAGjg4uATwB89gBQBKRKzhw9XI3ADaRrJWGDAy+AC6cAC8gli4+IrOCIQubB5e-L4h+FDRcE4A1ui2AFxwRbYQJHAKVgBuOGLI6GU2tm5R9u6eniFh6E5OgfgdXU5IAHRTw+hWfhXhZQ1NBCmp6Z60MMhQAjNsnoRpnt6+ZugwOXmFxWUVVTWj3T1bO3vx6PPF4QdE7JsXN6IOAhMrZdBQKznGBlaFXI6sQjsUw5EgMAgABQhqH4zxIwCgRgAcjgQC0QcZTGZfmIcMTSeSjCZzGwkaxTvAwNj+P08mh0CJwBJpB85FjCfxnOtWFyJVx+hNoU4qPjCTASWSqOtZTj5R8JiFlbT6ZqrFRMCBgDAABZa5hAA

Code: 

    import { useState } from 'react';

    function useComplexState&lt;T extends Record&lt;string, any&gt;&gt;() {
      const [state, setState] = useState&lt;T&gt;({});

      const setter = (key: keyof T, value: any) =&gt; {
        setState((state) =&gt; state != null ? ({ ...state, [key]: value }) : state);
        return state;
      };

      const getter = (key: keyof T) =&gt; {
        return state[key];
      }

      return { set: setter, get: getter };
    }

    interface Person {
      firstName: string;
      lastName: string;
    }

    const personState = useComplexState&lt;Person&gt;();

    personState.get('firstName');
    personState.set('lastName', 'Smith');

Basically, I want:
 
 * `T` to always be an object `{}` so no one should be able to do `useComplexState&lt;boolean&gt;`
 * Ideally, `T` shouldn't be required, so if I use `useComplexState()` I should just have a generic object Record&lt;string, any&gt; inside.
 * I'd like to have a default value of `{}` for the `state`.
## [10][utils-decorators library](https://www.reddit.com/r/typescript/comments/fwq433/utilsdecorators_library/)
- url: https://www.reddit.com/r/typescript/comments/fwq433/utilsdecorators_library/
---
Hi,

I just published a new version of my decorators library [utils-decorators](https://github.com/vlio20/utils-decorators).

It already contained decorators such as debounce, memozie, refreshable and many more. Now I am introducing some new and useful decorators which can be handy both for node and web applications.

The added decorators are:

`@timeout -` Retries execution of the decorated method. The method will be invoked extra x + 1 (where x is the retries values) in the worst case (when all invocation failed).

`@multiDispatch -` Will throw an error (`TimeoutError`) if the decorated method returned promise won't be resolved within the provided timeout (timeout is in milliseconds

`@retry -` Will invoke the decorated method the amount of provided times in parallel. This decorator is useful when you want to increase the chance of successful invocation of the decorated method.

`@delegate -` For a given input, if within the time period of the resolving of the promise of the first invocation the decorated method was invoked multiple times (with the same input) the response would be the promise that was generated by the first invocation.  
This way this decorator reduces the amount of calls to the implementation of the decorated method, for example accessing the database.  


I would be happy to hear your feedback and suggest improvements to the library. The library has 100% of unit tests code coverage but most importantly it has a 100 score in mutation tests (using [stryker](https://stryker-mutator.io/)).
## [11][Hey, I made a super small super type safe library to sort any type of array](https://www.reddit.com/r/typescript/comments/fwfwli/hey_i_made_a_super_small_super_type_safe_library/)
- url: https://www.reddit.com/r/typescript/comments/fwfwli/hey_i_made_a_super_small_super_type_safe_library/
---
check it out  
[sort-es](https://sort-es.netlify.com/) : Blazing fast, tree-shakeable, type-safe, modern utility library to sort any type of array
