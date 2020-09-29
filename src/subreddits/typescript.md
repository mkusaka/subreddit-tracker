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
## [2][I created a FULLY TYPED Twitter API Client for Node.js](https://www.reddit.com/r/typescript/comments/j1bidy/i_created_a_fully_typed_twitter_api_client_for/)
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
## [3][Is it standard to use the not null operator everywhere Env variables are used?](https://www.reddit.com/r/typescript/comments/j1iuc4/is_it_standard_to_use_the_not_null_operator/)
- url: https://www.reddit.com/r/typescript/comments/j1iuc4/is_it_standard_to_use_the_not_null_operator/
---
It's by far my most common use of !. In strict mode all Env variables have an implicit type of `string | undefined`.

Just curious if you guys also use ! with them. Or if you do something else such as asserting `process.env` as `any`
## [4][Trying to wrap my head around type inference failing](https://www.reddit.com/r/typescript/comments/j1bcj4/trying_to_wrap_my_head_around_type_inference/)
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
## [5][Looking for collaborators for open source project.](https://www.reddit.com/r/typescript/comments/j0vzdn/looking_for_collaborators_for_open_source_project/)
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
## [6][Typesafe, future proof way to get from "unknown" to "SomeInterfaceType" without "any"?](https://www.reddit.com/r/typescript/comments/j17c1h/typesafe_future_proof_way_to_get_from_unknown_to/)
- url: https://www.reddit.com/r/typescript/comments/j17c1h/typesafe_future_proof_way_to_get_from_unknown_to/
---
Is there a way to build functions with the signature of either `isSomeInterface(x:unknown):x is SomeInterface` or `assertSomeInterface(x:unknown):asserts x is SomeInterface` in a way that doesn't use `any`, and will break if I add a property to `SomeInterface` without changing the type guard?

I want a series of refinements, by the end of which the compiler knows that `x is SomeInterface` and it can therefore let me know if a change happens to `SomeInterface` and that's no-longer knowable.

Thoughts?
## [7][Type for an instance of a class specifically?](https://www.reddit.com/r/typescript/comments/j1405d/type_for_an_instance_of_a_class_specifically/)
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
## [8][a7ul/esbuild-node-tsc Build your Typescript Node.js projects using blazing fast esbuild.](https://www.reddit.com/r/typescript/comments/j0myuc/a7ulesbuildnodetsc_build_your_typescript_nodejs/)
- url: https://github.com/a7ul/esbuild-node-tsc
---

## [9][How to extend this inferred type?](https://www.reddit.com/r/typescript/comments/j0v4et/how_to_extend_this_inferred_type/)
- url: https://www.reddit.com/r/typescript/comments/j0v4et/how_to_extend_this_inferred_type/
---
Inside the if statement towards the bottom, I notice i am getting a lint on a type that seems to be inferred:

        const requestOptions = {
          method,
          mode: 'cors',
          headers: {
            authorization: `Bearer ${this.decryptedAccessToken}`,
            'content-type': options.contentType || 'application/json',
            ...options.specialHeaders,
          },
        };
    
        if (method === 'post') {
          requestOptions.body = body;
        }
    
    /*
        Property 'body' does not exist on type '{ method: "post" | "get"; mode: string; headers: any; }'.ts(2339)
    */

I tried various syntaxes for annotating and asserting an extended type that includes the `body` key but came up short. Can someone tell me how to properly type this to allow the new key `body`?
## [10][Optional property input into an optional parameter](https://www.reddit.com/r/typescript/comments/j0z2c7/optional_property_input_into_an_optional_parameter/)
- url: https://www.reddit.com/r/typescript/comments/j0z2c7/optional_property_input_into_an_optional_parameter/
---
The method buildEndpointUrl has an optional 2nd argument for queryParams. 

My intent below is to say "The options object sometimes isn't passed in. And then, sometimes it doesn't have a key for queryParams" 

I feel like only the first sentence is being implemented on line 3. But I can not put a ? on `.queryParams` there. 

So is a property always assumed optional if the parameter is defined as optional?

      public async get(urlPath: string, options?: GetRequestOptions): Promise&lt;object | void&gt; {
        const fullEndpointUrl = this.buildEndpointUrl(urlPath, options?.queryParams);
    
    private buildEndpointUrl
      &lt;QueryParamsShape extends { [idx: string]: string | string[] }&gt;(
        urlPath: string, queryParams?: QueryParamsShape
## [11][memoized-node-fetch: A wrapper around node-fetch (or any other fetch-like function) that returns a single promise until it resolves.](https://www.reddit.com/r/typescript/comments/j0ovcx/memoizednodefetch_a_wrapper_around_nodefetch_or/)
- url: https://github.com/chrispanag/memoized-node-fetch
---

