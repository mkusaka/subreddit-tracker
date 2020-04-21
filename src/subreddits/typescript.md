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
## [2][How should I write an interface for a curried function?](https://www.reddit.com/r/typescript/comments/g5c8pe/how_should_i_write_an_interface_for_a_curried/)
- url: https://www.reddit.com/r/typescript/comments/g5c8pe/how_should_i_write_an_interface_for_a_curried/
---
I am just started learning typescript and I want to know if this is the correct way to write the interface for a curried function

    interface Add {
      (a: number): (b: number) =&gt; number;
    }
    const add: Add = (a) =&gt; (b) =&gt; a + b
    console.log(add(1)(2))

Or do I have to explicitly add a return type to the function even though I already added it to the interface

    interface Add {
      (a: number): (b: number) =&gt; number;
    }
    const add: Add = (a) =&gt; (b): number =&gt; a + b
    console.log(add(1)(2))
    

My linter is warning me about not having the return type on the function. Should I add the return type twice (that feels like double work)? Also should I be using interfaces or types, or something else entirely that I don't know about?

This is my linter message btw *Missing return type on function.eslint*[*@typescript-eslint/explicit-function-return-type*](https://github.com/typescript-eslint/typescript-eslint/blob/v2.28.0/packages/eslint-plugin/docs/rules/explicit-function-return-type.md)

Thanks y'all
## [3][generate TypeScript functions from SQL queries?](https://www.reddit.com/r/typescript/comments/g57odt/generate_typescript_functions_from_sql_queries/)
- url: https://www.reddit.com/r/typescript/comments/g57odt/generate_typescript_functions_from_sql_queries/
---
In the Go library [sqlc](https://github.com/kyleconroy/sqlc), you write SQL queries, annotate it in certain way, and the library generates typesafe functions based on the queries and the annotations. I was wondering if anyone knows of a TypeScript library that does exactly this
## [4][Myzod v1.0.0-alphar release - Schema Validation and Type Inference](https://www.reddit.com/r/typescript/comments/g5eybi/myzod_v100alphar_release_schema_validation_and/)
- url: https://www.reddit.com/r/typescript/comments/g5eybi/myzod_v100alphar_release_schema_validation_and/
---
[Myzod](https://www.npmjs.com/package/myzod) is a runtime validation library who's goal is to use only typescript concepts to build type infer-able schemas.  The purpose of [myzod](https://www.npmjs.com/package/myzod) is to no longer have to match declared typescript types to the result of separately maintained validation logic, and in so doing minimise discrepancies between runtime  and the compile time types. Myzod is also inspired by [@hapi/joi](https://www.npmjs.com/package/@hapi/joi) and offers a similar validation api.  


At this time I am about ready to release version 1.0.0 but I am hoping to get more eyes on it before I do. Any feedback, issues, feature requests, or PRs would be extremely valuable.

Thanks in advance.
## [5][Deonify: For NPM module authors that would like to support Deno but do not want to write and maintain a port.](https://www.reddit.com/r/typescript/comments/g4op75/deonify_for_npm_module_authors_that_would_like_to/)
- url: https://github.com/garronej/denoify
---

## [6][Type Tetris - a typescript clone of Tetris, remix it!](https://www.reddit.com/r/typescript/comments/g50g2e/type_tetris_a_typescript_clone_of_tetris_remix_it/)
- url: https://glitch.com/~type-tetris
---

## [7][I made a template to help you make new Google Apps Script projects using TypeScript and Webpack](https://www.reddit.com/r/typescript/comments/g4t9ba/i_made_a_template_to_help_you_make_new_google/)
- url: https://github.com/iansan5653/gas-ts-template
---

## [8][JSCasts ep17 - build an entire startup in Node.js part2 - hooks and authentication](https://www.reddit.com/r/typescript/comments/g4z79l/jscasts_ep17_build_an_entire_startup_in_nodejs/)
- url: https://youtu.be/wOSPDtvJcyQ
---

## [9][Why does this generic constructor not seem to get type checked at all? Can it be?](https://www.reddit.com/r/typescript/comments/g4wyt7/why_does_this_generic_constructor_not_seem_to_get/)
- url: https://www.reddit.com/r/typescript/comments/g4wyt7/why_does_this_generic_constructor_not_seem_to_get/
---
Hi I am trying to use a generic  factory like class to create 1 type many times with the same passed in  arguments. My issue is that the passed in arguments don't seem to be  type checked against the generic type's constructor that is being  created. I tried to boil it down to as simple as possible example:

&amp;#x200B;

    abstract class SharedBaseClass {
        constructor(public name: string) {}
        abstract sayHi(): void;
    }
    interface IConstructor&lt;T&gt;
    {
        new (...args: any[]): T;
    }
    class Creator&lt;InputType extends SharedBaseClass&gt; extends SharedBaseClass
    {
        inputs: Array&lt;InputType&gt;;
    
        constructor(private inputNames: string[],
                    inputConstructor: IConstructor&lt;InputType&gt;,
                    ...params: ConstructorParameters&lt;IConstructor&lt;InputType&gt;&gt;) {
            super("Creator");
            this.inputs = new Array&lt;InputType&gt;(this.inputNames.length);
            this.inputNames.forEach((name: string, idx: number) =&gt; {
                this.inputs[idx] = new inputConstructor(name, ...params);
            });
        }
        runAll() {
            this.inputs.forEach((input: InputType, idx: number) =&gt; {
                input.sayHi();
            });
        }
        sayHi() {
            console.log("Hi from Creator");
        }
    }
    
    class Input1 extends SharedBaseClass {
        constructor(name: string, public age: number, public otherArg: number) {
            super(name);
        }
        sayHi() {
            console.log("Hi from ", this.name, " ", this.age, " ", this.otherArg);
        }
    }

Allowed invocations:

&amp;#x200B;

     let c = new Creator&lt;Input1&gt;(["a", "b", "c"], Input1); // no args
     let c = new Creator&lt;Input1&gt;(["a", "b", "c"], Input1, 11); // too few args
     let c = new Creator&lt;Input1&gt;(["a", "b", "c"], Input1, 11, 14); // correct # args
     let c = new Creator&lt;Input1&gt;(["a", "b", "c"], Input1, 11, 14, 15); // too many args 
    
     c.runAll();

 

Is there anyway I can have the generic type's constructor arguments actually validated correctly here?

Thanks!
## [10][How to keep types intact when returning Promise from try-catch?](https://www.reddit.com/r/typescript/comments/g4vn2d/how_to_keep_types_intact_when_returning_promise/)
- url: https://www.reddit.com/r/typescript/comments/g4vn2d/how_to_keep_types_intact_when_returning_promise/
---
I have a bunch of functions that I want to batch together with Promise.all, but I want each individual function to manage its own error/loading states.

How would you write the following code in such a way that "x" is of the type "SomeType"? Right now, since the catch block doesn't return anything, the compiler is unhappy. I would like to avoid having to give "x" the type "SomeType | null" if possible.  

&amp;#x200B;

`getSomePromise(): Promise&lt;SomeType&gt; {`  
 `let promise = new Promise();`  
 `this.setState({ isLoading: true, isError: false });`  
 `try {`  
 `promise = someHTTPcallFunction();`  
 `} catch (err) {`  
 `this.setState({ isError: true });`  
 `}`  
 `return promise;`  
  `}`

`const x = await getSomePromise();`

`const [xData, otherPromiseDataHere] = await Promise.all([x, otherPromisesHere])`
## [11][Hydro-SDK - Author native Flutter experiences in Typescript and deliver updates directly to users over the air and out of band](https://www.reddit.com/r/typescript/comments/g4vcel/hydrosdk_author_native_flutter_experiences_in/)
- url: https://github.com/chgibb/hydro-sdk
---

