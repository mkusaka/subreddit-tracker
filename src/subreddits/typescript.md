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
## [2][Announcing TypeScript 3.9](https://www.reddit.com/r/typescript/comments/gilb3e/announcing_typescript_39/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-3-9/
---

## [3][Can someone explain the need for an awaited type?](https://www.reddit.com/r/typescript/comments/gistgu/can_someone_explain_the_need_for_an_awaited_type/)
- url: https://www.reddit.com/r/typescript/comments/gistgu/can_someone_explain_the_need_for_an_awaited_type/
---
Title. It seems like the current type definition of the promise works as expected for all of my applications. What problem does `awaited T` solve?

I've read through the issue, and it still isn't making very much sense. Await seems to correctly resolve these types to what is expected, at least in my experience.

Any help would be appreciated, thank you!
## [4][Anyway to have a any type but still force the compiler to give the 'could be undefined' warning?](https://www.reddit.com/r/typescript/comments/giyaao/anyway_to_have_a_any_type_but_still_force_the/)
- url: https://www.reddit.com/r/typescript/comments/giyaao/anyway_to_have_a_any_type_but_still_force_the/
---

## [5][I am working on a "node typescript architecture" project which suggests a structure for Node TS apps. It's WIP but any feedback is appreciated.](https://www.reddit.com/r/typescript/comments/gixr3e/i_am_working_on_a_node_typescript_architecture/)
- url: https://www.reddit.com/r/typescript/comments/gixr3e/i_am_working_on_a_node_typescript_architecture/
---
This is still in quite an early stage, and the docs are very incomplete, but I've been working on a suggested Node Typescript Architecture example over on Github. It is based on what I've picked up over the last two or three years of writing Node apps.

The idea is, as concisely as I can make it: use IOC in your domain code, and functional programming to replace a DI container. Write a set of adapters that partially apply you domain functions at runtime to provide 'context'. These adapters can wrap the function to implement things like transactions, resource spinup / teardowns etc.

The project is here: [https://github.com/jbreckmckye/node-typescript-architecture](https://github.com/jbreckmckye/node-typescript-architecture)

There is example code which implements a simple app using Postgres (with transactions) and RabbitMQ events, and Express.js bindings. It is runnable if you have Docker installed (at least Docker for Mac, I have no idea about Docker for windows).

There is also the beginning of a gitbook (but it's still very incomplete).

Any feedback would be invaluable.
## [6][Is using an assertion function to validate a number fits a range valid use? Is there a library of assertion functions?](https://www.reddit.com/r/typescript/comments/gisvg1/is_using_an_assertion_function_to_validate_a/)
- url: https://www.reddit.com/r/typescript/comments/gisvg1/is_using_an_assertion_function_to_validate_a/
---
... so, I'm really new to TypeScript, but I was thinking, there's not really a way to validate a numerical range in typescript directly (without using something like a potentially massive union) ..  but is there a reason why you shouldn't use assertion function to validate that a number fits a specific range?  My current use case is just to validate that a number is zero or positive.  

Secondly, does anyone out here know of any useful libraries of assertion functions?  This seems like a feature that could be quite useful to use to build up quite a suite of things with.
## [7][ncdc - a typescript consumer driven contract tool](https://www.reddit.com/r/typescript/comments/giiukv/ncdc_a_typescript_consumer_driven_contract_tool/)
- url: https://www.reddit.com/r/typescript/comments/giiukv/ncdc_a_typescript_consumer_driven_contract_tool/
---
I think I've made some good progress on my first npm package, [ncdc](https://github.com/tamj0rd2/ncdc). ncdc (node cdc) is a tool that takes a consumer contract (written in yaml) and tests the specified endpoints against a producer. You can also run a single command to mock those endpoints.

The point of using this tool is to ensure that each endpoint you use in development is called in the same way and responds with the same data that you expect from your live environment. There's not much point in testing against mock APIs that don't strongly resemble your live APIs.

It easily integrates with typescript to prevent you from having to write JSON schemas for your already defined Typescript interfaces, types and enums. With this, you're guaranteed that your code, mock API endpoints and real endpoints are all consistent with each other.

I made this after struggling for a while with another CDC tool we were using at work. It wasn't great at handling optional fields, so making a tool that can validate using JSON schemas without us having to basically duplicate our contracts seemed like a great solution.

https://github.com/tamj0rd2/ncdc

https://www.npmjs.com/package/ncdc
## [8][Noob question but how would you type a code like this?](https://www.reddit.com/r/typescript/comments/gipmo7/noob_question_but_how_would_you_type_a_code_like/)
- url: https://www.reddit.com/r/typescript/comments/gipmo7/noob_question_but_how_would_you_type_a_code_like/
---
I'm using Inertia.js which is why I even needed this snippet. How can I type the hook so it works just like `useState`?

```
import { useEffect, useRef, useState } from "react";

export default function useStateIfMounted(initialState: any) {
    const isMounted = useRef(true);

    useEffect(() =&gt; {
        return () =&gt; {
            isMounted.current = false;
        };
    }, []);

    const [state, setState] = useState(initialState);

    function setStateIfMounted(newState: any) {
        if (isMounted.current) {
            setState(newState);
        }
    }

    return [state, setStateIfMounted];
}
```
## [9][Question about conditional types.](https://www.reddit.com/r/typescript/comments/gino4p/question_about_conditional_types/)
- url: https://www.reddit.com/r/typescript/comments/gino4p/question_about_conditional_types/
---
I'm trying to write a function for automatic handling of HTTP responses of certain response codes. Though, the context doesn't matter. Essentially, I have a type with some kind of discriminator, named `code`. If code is equal to a given value, call one function, otherwise call another function. The hard part is getting the called functions to have the correct types. Explaining it is pretty hard, but I think the code will do the question more justice. 

I've written what I'd expect to work, but I can't get it to compile. Interestingly, Typescript infers the types correctly for the parameters to the passed-in functions. Enough talk. The code: 

    const ifOne = &lt;R extends {code: 1 | 2 | 3 }&gt;(
        param: R, 
        oneHandler: (result: R extends {code : 1} ? R : never) =&gt; void,
        otherHandler: (result: R extends {code : 1} ? never : R) =&gt; void,
    ) =&gt; {
        if (param.code === 1) {
            oneHandler(param);
        } else {
            otherHandler(param);
        }
    }
    
    type Type = {
        code: 1;
        value: "Hello";
    } | {
        code: 2;
        value: "Red";
    } | {
        code: 3;
        value: "Dit";
    }
    
    const param: Type = {
        code: 2,
        value: "Red"
    };
    
    ifOne&lt;Type&gt;(param, 
        one =&gt; {
    
        },
        other=&gt; {
            // I'd expect this function to be called. 
        }
    );

So If I write this in vscode, the types of `one` and `other` are correctly inferred. It knows `one` is of type `{code: 1, value: "Hello"}`, and `other` is of  type `{code: 2, value: "Red"} | {code: 3, value: "Dit"}`. But, the compilation error is in the function `ifOne`. It refuses to call `oneHandler` and `twoHandler` with the parameter `param`. With the error: 

" Argument of type 'R' is not assignable to parameter of type 'R extends { code: 1; } ? R : never'.
  Type '{ code: 1 | 2 | 3; }' is not assignable to type 'R extends { code: 1; } ? R : never'"

Perhaps this is not possible, in which case I can find another way to do what I need, but I'm curious if I'm missing something obvious here, or if this is something Typescript doesn't support. 

Thanks
## [10][Noob question about type declaration files: a package I'm writing a file for default exports an object literal, how do I declare that in the .d.ts file?](https://www.reddit.com/r/typescript/comments/gihrne/noob_question_about_type_declaration_files_a/)
- url: https://www.reddit.com/r/typescript/comments/gihrne/noob_question_about_type_declaration_files_a/
---
Declaring it as a class or function works but adds extra props and methods to the imported object (such as `arguments`, `prototype` etc). I'd like it where the imported object is used like so:

    import foo from 'fooLibrary';

    foo.validMethod();

So it has 'static' methods but is not written out like an actual class. I can't export an object literal in the .d.ts file. Exporting it as an interface doesn't work as the IDE complains that `foo` is being used as a value when it should be used as a type. Thanks in advance to any respondents.

EDIT: I *think* I've worked it out. Just need to use and approach `declare module 'fooLibrary'` like it is an object literal from the start. So `declare module 'fooLibrary' { export function validMethod(): string; }` would produce the desired result above.
## [11][Unionize and Objectify: a trick to bring conditional types to objects](https://www.reddit.com/r/typescript/comments/gif63z/unionize_and_objectify_a_trick_to_bring/)
- url: https://effectivetypescript.com/2020/05/12/unionize-objectify/
---

