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
## [2][A library to generate the maximally efficient series of unique keys for a given alphabet.](https://www.reddit.com/r/typescript/comments/hrm45x/a_library_to_generate_the_maximally_efficient/)
- url: https://github.com/gactjs/key
---

## [3][Why does my typescript file complain about not finding the declared variable and assign its number?](https://www.reddit.com/r/typescript/comments/hrkrr3/why_does_my_typescript_file_complain_about_not/)
- url: https://www.reddit.com/r/typescript/comments/hrkrr3/why_does_my_typescript_file_complain_about_not/
---
 I'm getting an error say that it cannot find age, when trying to assign to 18, 

Does someone know why?

    declare var age: number  
    
    age = 18 // error


When I try to build binder with webpack it say variable cannot be found.
## [4][Marking Optional Parameters in Callbacks As NonOptional?](https://www.reddit.com/r/typescript/comments/hrmj9s/marking_optional_parameters_in_callbacks_as/)
- url: https://www.reddit.com/r/typescript/comments/hrmj9s/marking_optional_parameters_in_callbacks_as/
---
From the official docs:

## Optional Parameters in Callbacks [\#](https://www.typescriptlang.org/docs/handbook/declaration-files/do-s-and-don-ts.html#optional-parameters-in-callbacks)

*Don’t* use optional parameters in callbacks unless you really mean it:

    /* WRONG */ 
    interface Fetcher {    
      getObject(done: (data: any, elapsedTime?: number) =&gt; void): void; 
    } 

This has a very specific meaning: the `done` callback might be invoked with 1 argument or might be invoked with 2 arguments. The author probably intended to say that the callback might not care about the elapsedTime parameter, but **there’s no need to make the parameter optional to accomplish this – it’s always legal to provide a callback that accepts fewer arguments.**

*Do* write callback parameters as non-optional:

    /* OK */ 
    interface Fetcher {    
      getObject(done: (data: any, elapsedTime: number) =&gt; void): void; 
    }

\----------

Unless callbacks are unusual, I seem to remember a lot of lint errors thrown by TSC along the lines of `Expected 2 arguments, but received only 1`. I have a hard time squaring this with the above, can anyone clarify this?
## [5][Cool helper type that will allow you to recursively replace all types in an object with another type.](https://www.reddit.com/r/typescript/comments/hr19a2/cool_helper_type_that_will_allow_you_to/)
- url: https://www.reddit.com/r/typescript/comments/hr19a2/cool_helper_type_that_will_allow_you_to/
---
Figured this out recently, when I had to transform large plain objects of different shapes, and replace some leaf value with a different one. `Replaced` type takes 4 generic arguments:

* `T` the object to be transformed.
* `TReplace` the type to be replaced with
* `TWith` the type to use instead of `TReplace`
* `TKeep` the type to not transform. Defaults to primitive values.  


Obviously to be used as the return type of some function that does the actual transformation.

    type Primitive = string | number | bigint | boolean | null | undefined;
    
    type Replaced&lt;T, TReplace, TWith, TKeep = Primitive&gt; = T extends TReplace | TKeep
        ? (T extends TReplace
            ? TWith | Exclude&lt;T, TReplace&gt;
            : T)
        : {
            [P in keyof T]: Replaced&lt;T[P], TReplace, TWith, TKeep&gt;
        }
    
    type Foo = symbol;
    type Bar = "3";
    
    type ToReplace = {
        x?: {
            y: Foo | number;
            z: string;
        }[]
    } | undefined;
    
    type WasReplaced = Replaced&lt;ToReplace, Foo, Bar&gt;;
    /* Output:
    {
        x?: {
            y: number | "3";
            z: string;
        }[];
    }
    */
## [6][Is reverting back to require due to no declaration file a valid reason?](https://www.reddit.com/r/typescript/comments/hr8v9r/is_reverting_back_to_require_due_to_no/)
- url: https://www.reddit.com/r/typescript/comments/hr8v9r/is_reverting_back_to_require_due_to_no/
---
    import audioConcat from 'audioconcat';
    
    // Could not find a declaration file for module 'audioconcat'. 
    // '/path/node_modules/audioconcat/index.js' implicitly has an 
    // 'any' type.
     
    // Try `npm install @types/audioconcat` if it exists or add a 
    // new declaration (.d.ts) file containing `declare module 'audioconcat';`ts(7016)

I need practice writing declaration files so I'll do it this time.

In general will you guys always do that or just go back to a `require` statement? I notice those don't seem to need declaration files, at least according to ESLint and TSC
## [7][Call a new function after assigning value through await asynchronous call](https://www.reddit.com/r/typescript/comments/hr8mfk/call_a_new_function_after_assigning_value_through/)
- url: https://www.reddit.com/r/typescript/comments/hr8mfk/call_a_new_function_after_assigning_value_through/
---
Hello everyone, I'm having this issue where I have an array of instructors with 3 elements (employee id, name, email). I need to try and have the names searchable in an input field, I thought it would be easiest to create a new variable with just an array of names.

I assign the value to my instructors with this call

    this.allInstructors = await this.instructorsEndpoint().toPromise();

I thought if I added .then(otherFunction()) after toPromise() I could have that function run after the promise is filled but this turned out not to be the case. (this is what I mean)

    this.allInstructors = await this.instructorsEndpoint().toPromise().then(this.transferInstructorData);

When I try to compile I get this error

     Type 'void | InstructorData[]' is not assignable to type 'InstructorData[]'.
    [cad-classes-ui-7cd547bf4-8w8s5 ui]   Type 'void' is not assignable to type 'InstructorData[]'.
    [cad-classes-ui-7cd547bf4-8w8s5 ui]

Any ideas for what to do?
## [8][Omit an extending interface in React types](https://www.reddit.com/r/typescript/comments/hr87p0/omit_an_extending_interface_in_react_types/)
- url: https://www.reddit.com/r/typescript/comments/hr87p0/omit_an_extending_interface_in_react_types/
---
Hey. I just know the basic of the basics in Typescrip and I have a doubt.

So, I using attributes from React but I'm not getting how to exclude some attributes from an element plus a whole extending interface.

`import React, {`  
 `ButtonHTMLAttributes,`  
 `DOMAttributes`  
`} from "react";`

`type BtnAttrs = Omit&lt;ButtonHTMLAttributes&lt;HTMLButtonElement&gt;, "attr1" | "attr2&gt;;`

But ButtonHTMLAttributes is extending HTMLAttributes and this one DOMAttributes. How may I omit DOMAttributes as well?
## [9][How can I type the ArrayConstructor to get rid of Unsafe spread of an any value in an array?](https://www.reddit.com/r/typescript/comments/hr29y9/how_can_i_type_the_arrayconstructor_to_get_rid_of/)
- url: https://www.reddit.com/r/typescript/comments/hr29y9/how_can_i_type_the_arrayconstructor_to_get_rid_of/
---
```ts
    export function createMultiple&lt;T&gt;(size: number, factory: () =&gt; T): T[] {
      return [...Array(size)].map(() =&gt; factory());
    } // red squiggle under ...Array(size): Unsafe spread of an any value in an array

    // Arrange
    class Test {}
    function factory(): Test {
      return new Test();
    }
    // Act
    const zombieArray = createMultiple(5, factory);

    // Assert
    assert(zombieArray[0] instanceof Test); // OK, it passes
    assertEquals(zombieArray, [{}, {}, {}, {}, {}]); // OK, it passes
```
https://www.typescriptlang.org/play/?ssl=1&amp;ssc=1&amp;pln=15&amp;pc=73#code/AQ4UwDwBwewJwC7AGYFcB2BjBBLG7hM4wBDBMAWVQBtcpqwAeAFQD4AKAZxwC8wAuYOlQBbAEZg4AGhQls8AJ6D2ASmABeVsGYrBzANoBdYAG8AsAChQoYglRwC+gHQuAgnDgkFXXmBWGnERIodlUNLWQ5BEVVFQBuS2sAX2AAelTgYgATYE4AR1QcAHMihmAMLMlgFyd3T29uPl1gAFV0ThJkMFyoYhIcmGRgEgIRhWAANxJqVG6cUdGPL0tE0HTgOpGisFWQTGoSTk5tME4kEyTdlAxsPAJI+ThvZuZT86uQW3sCdDAAdxOZ1UCSsoEuoJA61c2CumHwZ2APBg4hwYE243UhD65CotBw9DA7AArDIHtEnvEVhC0hlXEdJAgrutDpwGewkSi0UsFPoAAzGeZnEaYMCDQEIeI04AAeQA0jIcEgoCzTkyMirEABRArTTjs5FiVHomT6C4yM2mJLmq2WwyS9ZyhVKlWcIA

(the error is from `typescript-eslint`, so it is not highlighted)
## [10][I made a Snapchat clone in the browser!](https://www.reddit.com/r/typescript/comments/hqfzmn/i_made_a_snapchat_clone_in_the_browser/)
- url: https://v.redd.it/gnbkqf8duma51
---

## [11][How to resolve New Relic for Node.js halted startup due to an error: Error: Failed to connect to collector?](https://www.reddit.com/r/typescript/comments/hr4f5g/how_to_resolve_new_relic_for_nodejs_halted/)
- url: https://www.reddit.com/r/typescript/comments/hr4f5g/how_to_resolve_new_relic_for_nodejs_halted/
---
I've cloned a fully functioning repo. The project is in Typescript and I'm adding some unit tests using mocha. The project uses NewRelic and when I run the project I get the error: New Relic for Node.js halted startup due to an error: Error: Failed to connect to collector. 

There are more details posted in the following StackOverflow post. Any advice on this will be greatly appreciated!

[https://stackoverflow.com/questions/62899405/how-to-resolve-new-relic-for-node-js-halted-startup-due-to-an-error-error-fail](https://stackoverflow.com/questions/62899405/how-to-resolve-new-relic-for-node-js-halted-startup-due-to-an-error-error-fail)
