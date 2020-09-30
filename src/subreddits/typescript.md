# typescript
## [1][Who's hiring Typescript developers - September](https://www.reddit.com/r/typescript/comments/ik9rft/whos_hiring_typescript_developers_september/)
- url: https://www.reddit.com/r/typescript/comments/ik9rft/whos_hiring_typescript_developers_september/
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
## [2][I tried to write clean architecture based back-end application in TypeScript](https://www.reddit.com/r/typescript/comments/j27rqe/i_tried_to_write_clean_architecture_based_backend/)
- url: https://github.com/pvarentsov/typescript-clean-architecture
---

## [3][Excluding the keys of a base type](https://www.reddit.com/r/typescript/comments/j2bggl/excluding_the_keys_of_a_base_type/)
- url: https://www.reddit.com/r/typescript/comments/j2bggl/excluding_the_keys_of_a_base_type/
---
I have a base type (ex: Record&lt;string, any&gt;) that I want to ensure is true for all sub-type instances. For each sub-type, I have a set of specific keys which will be valid which return a specific set of values.

For example: type MyType extends Record&lt;string, any&gt; = { a: string, b: number }

I am looking for a way to get the type "a" | "b" from MyType. However, using "keyof MyType" == "string | number".

What is the right way to extract only ("a" | "b") into a type of its own, given MyType as the input?

Playground:
https://www.typescriptlang.org/play?#code/C4TwDgpgBAksEFsDOUC8UBKEDGB7ATgCYA8Sw+AlgHYDmANFAIZUgB8A3FAPRdQCyEZsCjBcUMPlwA3CoWiMoAI0ZJooSFDL4ArtmDb8EBlVzCkjEFABEzS1uo0oFFLgDWFqwChq8fADNGbGg+EDhEFAgAD3gqQhQw5CgAb25eASERMQlpWXkqKCiwABsKbAphErIoXD8oKUYSwihXCBAUZibgAAsICnxNSDK-UqhDfXx89QgkTyg5pgAuTXIHOln5xSWqbQRFCHxPAF9PTymoAGlWpABBPW0GtGbWmv5Q+GROHjR0e1ooAB8oNtdvtTuBoJc2gARaZ9CBNdA2KwA6yKKyfXhyJBwwgnPzaKh6Ci4fI0CDABLEAAqBWiEFiKEhNzuDVYAApUlAAMJdZhkkRdZxQEpUNRiJkw7GGTpiQKCiBSaBYnFKCC8mQEAB06zm5XCSxCCSQDB1TkISypAEpPAa3uEANpUgC6yVNYwM+T1yHtsid7COJzwVCqXqQtqNjySpsYSyshEYwEYVjW8yUSwArAHPEGqvUitoINdHmSKe82aGGEjLRjlpRaJqoAB1AiuBzanPCPMFgBCxfJCXL72NqKs1c5wL2+Abzfwrfr2ZJuYaBa5fdLiEH4Ur2FHNdsUAAhEemxlROJJIRdPIROCCvhJP0eoYoH4CExPVQ87JrDuniBNUAA
## [4][What am I missing here?](https://www.reddit.com/r/typescript/comments/j27k9e/what_am_i_missing_here/)
- url: https://www.reddit.com/r/typescript/comments/j27k9e/what_am_i_missing_here/
---
[Playground](https://www.typescriptlang.org/play?#code/C4TwDgpgBAKgjFAvFA2gcgM7AE4EsB2A5mgDRQDeAtgQPwBcU+ArpQEYTZmUCGAHvYxbtsAXwC6AbgBQoSLABMSVGm7Zs3EKVgxJUmeGgwYSKVDNQAPrDinzVmPKnSpEXmAD22YFABmTfADGwLju+FAAFtwYUAA8MAB8UAAU3Azc+CBkANYMWRAg7j6wAJRpULjR5CgA+lnlYXkFRToMMDVZYiJS5LZm2BDATNgN9VDcUABkE2MoHVAAhIjI-gAmED4EECvSXVIEwBw+3AHQAFLuBDEAysDcwQEwZABK-sGUEAkUvWMBJ2DASWADBe+DeHzIADcGDc7rgHqUoKx3O4ADYQdI7PQBUJYKBPeAMc6XLB4IhkeCJZA9cw-P4A4CQ4pfGk0-qDYbJWQQQpQCFIJZQTA4AjEJlTb4s8xJeaRDCAlBwMRkNDUfBoJkWKzABViAB0qtiyAhurRRGA4TFEwlkuSMqi8sVyp4vHVli1Ov1fCg8SNJogZot0hpXV22PwuPx8kJF3wMX8WXw7gA7vgUEqFJTmeZjnTASQIUzqTaw7iMExWFylNrFUGbWyhmEkgBBNQaXUVFvqEBJAuW62S40QCEcbtMJCJfFtMsVgwoAAMYj1OYg-yS065ZCYxWKtbMIaxOO8k6URbMQtJxGB8BI3xUrc0V-kN5EQA)

I want to create a simple type validator for some data I get coming in.

This is a sub-example but it already doesn't work. What am I doing wrong?
## [5][I created a FULLY TYPED Twitter API Client for Node.js](https://www.reddit.com/r/typescript/comments/j1bidy/i_created_a_fully_typed_twitter_api_client_for/)
- url: https://www.reddit.com/r/typescript/comments/j1bidy/i_created_a_fully_typed_twitter_api_client_for/
---
I created a FULLY TYPED Twitter API Client for Node.js 🔧  
This client is taking full advantage of modern TypeScript capabilites!

✅ **Fully typed!** Both for query parameters and responses.  
✅ Includes 90% of the **official Twitter API** endpoints.  
✅ **Promise-based!** No ugly callbacks.  
✅ Inbuilt in-memory **cache** for rate-limit friendly usage.

Check it out 👇

[https://github.com/Silind/twitter-api-client](https://github.com/Silind/twitter-api-client)
## [6][Is it standard to use the not null operator everywhere Env variables are used?](https://www.reddit.com/r/typescript/comments/j1iuc4/is_it_standard_to_use_the_not_null_operator/)
- url: https://www.reddit.com/r/typescript/comments/j1iuc4/is_it_standard_to_use_the_not_null_operator/
---
It's by far my most common use of !. In strict mode all Env variables have an implicit type of `string | undefined`.

Just curious if you guys also use ! with them. Or if you do something else such as asserting `process.env` as `any`
## [7][Trying to wrap my head around type inference failing](https://www.reddit.com/r/typescript/comments/j1bcj4/trying_to_wrap_my_head_around_type_inference/)
- url: https://www.reddit.com/r/typescript/comments/j1bcj4/trying_to_wrap_my_head_around_type_inference/
---
Hi, I've run into Typescript failing to infer something and I'm trying to understand why. This simple example works:

```
type Group&lt;A, B&gt; = {
    createA: () =&gt; A,
    aToB: (a: A) =&gt; B,
};

function doSomething&lt;A, B&gt;(group: Group&lt;A, B&gt;) {
    const a = group.createA();
    const b = group.aToB(a);

    return b;
}

const example = doSomething({
    createA: () =&gt; 123,
    aToB: (value) =&gt; value.toString(),
});
```

So `example` is inferred as `Group&lt;number, string&gt;`, aToB's value parameter is number, `aToB` returns a string. All makes sense.

But if I add another function to `Group` that takes `B` as a parameter:

```
type Group&lt;A, B&gt; = {
    createA: () =&gt; A,
    aToB: (a: A) =&gt; B,
    logB: (b: B) =&gt; void
};

function doSomething&lt;A, B&gt;(group: Group&lt;A, B&gt;) {
    const a = group.createA();
    const b = group.aToB(a);

    group.logB(b);

    return b;
}

const example = doSomething({
    createA: () =&gt; 123,
    aToB: (value) =&gt; value.toString(),
    logB: (value) =&gt; console.log(value)
});
```

`A` remains a number, but `B` becomes `unknown`.

I'm trying to figure out why this is. I assume now that `B` is used as a parameter, that stops Typescript from being able to infer `B` as being the result of `createA`, and since there's no type on `logB: (value)`, it gives up and calls it unknown.

I can get around this by `logB: (value: string) =&gt; console.log(value)`, but is there a way to tell Typescript to always infer `B` from the return type of `createA`, even if it's used as a parameter elsewhere?
## [8][Looking for collaborators for open source project.](https://www.reddit.com/r/typescript/comments/j0vzdn/looking_for_collaborators_for_open_source_project/)
- url: https://www.reddit.com/r/typescript/comments/j0vzdn/looking_for_collaborators_for_open_source_project/
---
Hi, good evening!

My name is William and I am a senior engineer based in the UK.

I am looking for a couple of junior developers who would like to work with me on an open source project that aims to help photographers and other media artist with a workflow manager.

If anyone is looking for a project to work on to learn more about Vue and TypeScript, I would be very happy to mentor them.
​
Cheers!

---
EDIT

Thanks to everyone who expressed interested in Photion!

This is the main repo we should be working on:
https://github.com/photion/web-admin

For anyone interested into joining, here is a slack invite:
https://join.slack.com/t/photion/shared_invite/zt-hnqy02xy-wB0vJqZv_lhrB~7qWv6~VA
## [9][Typesafe, future proof way to get from "unknown" to "SomeInterfaceType" without "any"?](https://www.reddit.com/r/typescript/comments/j17c1h/typesafe_future_proof_way_to_get_from_unknown_to/)
- url: https://www.reddit.com/r/typescript/comments/j17c1h/typesafe_future_proof_way_to_get_from_unknown_to/
---
Is there a way to build functions with the signature of either `isSomeInterface(x:unknown):x is SomeInterface` or `assertSomeInterface(x:unknown):asserts x is SomeInterface` in a way that doesn't use `any`, and will break if I add a property to `SomeInterface` without changing the type guard?

I want a series of refinements, by the end of which the compiler knows that `x is SomeInterface` and it can therefore let me know if a change happens to `SomeInterface` and that's no-longer knowable.

Thoughts?
## [10][Type for an instance of a class specifically?](https://www.reddit.com/r/typescript/comments/j1405d/type_for_an_instance_of_a_class_specifically/)
- url: https://www.reddit.com/r/typescript/comments/j1405d/type_for_an_instance_of_a_class_specifically/
---
In this example, is it possible to have \`printName\` only accept an instance of a Cat? Currently it also accepts objects matching the interface defined by Cat, which may lead to bugs down the road

    class Cat {
        constructor(public name: string) {
            if (name.length &lt;= 1) {
                throw new Error('Name is too short!');
            }
        }
    }
    
    function printName(cat: Cat) {
        console.log(cat.name);
    }
    
    // printName(new Cat('')); // Good, would like to enforce this usage
    printName({name: ''}); // No warning from TypeScript :(

This workaround enforces that usage

    function printName(cat: Cat) {
        if (cat instanceof Cat) {
          console.log(cat.name);
        }
        throw new Error('Invalid cat');
    }

However, can this be checked at compile-time through a type instead?
## [11][a7ul/esbuild-node-tsc Build your Typescript Node.js projects using blazing fast esbuild.](https://www.reddit.com/r/typescript/comments/j0myuc/a7ulesbuildnodetsc_build_your_typescript_nodejs/)
- url: https://github.com/a7ul/esbuild-node-tsc
---

