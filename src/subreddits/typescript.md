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
## [2][Linkdash - Generate a handy dashboard of links in seconds!](https://www.reddit.com/r/typescript/comments/fxpx9q/linkdash_generate_a_handy_dashboard_of_links_in/)
- url: https://i.redd.it/3mgoes21prr41.gif
---

## [3][TypeScript Tutorial For Beginners](https://www.reddit.com/r/typescript/comments/fxo95b/typescript_tutorial_for_beginners/)
- url: https://afteracademy.com/blog/typescript-tutorial-for-beginners
---

## [4][Instantly find code using intuitive search and code rank — live demo (metacode.app)](https://www.reddit.com/r/typescript/comments/fxfeih/instantly_find_code_using_intuitive_search_and/)
- url: https://v.redd.it/ythfwk0dxnr41
---

## [5][I'm forced to discriminate a union type to then perform the same operation (Visitor Pattern for ASTs)](https://www.reddit.com/r/typescript/comments/fxolgd/im_forced_to_discriminate_a_union_type_to_then/)
- url: https://stackoverflow.com/questions/61111469/forced-to-discriminate-a-union-type-to-then-perform-the-same-operation-visitor?stw=2
---

## [6][Foal TS - April release (version 1.7) - File uploads to local or Cloud storage (AWS S3), OpenAPI spec generation, safe config access](https://www.reddit.com/r/typescript/comments/fx40ui/foal_ts_april_release_version_17_file_uploads_to/)
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
## [7][Do you guys verify your axios responses?](https://www.reddit.com/r/typescript/comments/fx75qu/do_you_guys_verify_your_axios_responses/)
- url: https://www.reddit.com/r/typescript/comments/fx75qu/do_you_guys_verify_your_axios_responses/
---
My App recieves a lot of data through axios, and currently I do not type check any of it. So basically I'm casting completely unknown values into typescript. Which is not ideal.

I wonder, does anyone check their axios responses somehow? Runtime typechecking doesn't really seem to be a thing for typescript.

I've read about user defined type guards, but I couldn't get them to do what I want them to do, which is something like this

    interface WhatIWant {
        ...
    }
    
    Axios.get(...).then(res =&gt; {
        if (verifyType(res.data, WhatIWant))
            setState(res.data)
        else throw new Error("Unexpected Type")
    }

Going full verbose shouldn't be the moral of the story and doesn't use my typescript types.

    let data;
    if (typeof res.data.foo === string) data.foo = res.data.foo;
    if (typeof res.data.bar === number) data.bar = res.data.bar;

I've heard of libraries like yup. But again, that doesn't solve anything if I'm not using my typescript types/interfaces.

Edit: Thanks for the answers! Great suggestions in this thread, I'll check them out when I have time :-)
## [8][Type Checking String Literal Access](https://www.reddit.com/r/typescript/comments/fxcjo4/type_checking_string_literal_access/)
- url: https://www.reddit.com/r/typescript/comments/fxcjo4/type_checking_string_literal_access/
---
I'm getting data from this API that has some irritatingly shaped data.  

```
interface Res{
    Title: string;
    'Publish Date': string;
    'Citation Count': number; 
}
```

I get type safety when accessing `obj.Title` but not when accessing `obj['Publish Date']`

I would prefer not to do manipulations of the API response if I don't have to. What are your suggestions?
## [9][Dealing with types at scale](https://www.reddit.com/r/typescript/comments/fx5nrx/dealing_with_types_at_scale/)
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
## [10][Type error while binding a value](https://www.reddit.com/r/typescript/comments/fx6bip/type_error_while_binding_a_value/)
- url: https://www.reddit.com/r/typescript/comments/fx6bip/type_error_while_binding_a_value/
---
Begginer question here..

const item = ({ index, value }) =&gt; \`${index}:${value}\`; 

Error in VS: Binding element 'index' implicitly has an 'any' type.ts(7031)

where do I specificy the type for the above line?
## [11][Can't solve this TypeScript error! :(](https://www.reddit.com/r/typescript/comments/fx2lha/cant_solve_this_typescript_error/)
- url: https://www.reddit.com/r/typescript/comments/fx2lha/cant_solve_this_typescript_error/
---
Was trying to work with a React app, but this Typescript error happens. Since I'm a noob at TS, not sure what's happening. Any clue, pros? :( 

https://preview.redd.it/vly86gw8zjr41.png?width=1053&amp;format=png&amp;auto=webp&amp;s=768d6e7d0d78be2634dd91747a1918354b9043be
