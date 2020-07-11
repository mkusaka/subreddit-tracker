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
## [2][[ Help ] Create new method (wrapping)](https://www.reddit.com/r/typescript/comments/hp5gao/help_create_new_method_wrapping/)
- url: https://www.reddit.com/r/typescript/comments/hp5gao/help_create_new_method_wrapping/
---
Hi you! As the title says I need to create a new method and I know there is a prototype way but isn't encouraged.
I was told about wrapping and then I looked into the documentation and Stack Overflow but I couldn't find any clear explanation/answer.

I want achieve something like this:
numberInstance.myMethod(anyNumber)

Thanks for helping me !
## [3]["--isolatedModules flag" in typescript](https://www.reddit.com/r/typescript/comments/hp440a/isolatedmodules_flag_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/hp440a/isolatedmodules_flag_in_typescript/
---
&gt; All files must be modules when the '--isolatedModules' flag is provided

I've googled this error but I honestly don't think I'm violating anything. 

    import React from 'react';
    import Hero from '../components/Hero';
    
    const About: React.FC = () =&gt; (
      &lt;Hero /&gt;
    );
    
    export default About;
    
The error occurs at "(1,1)"..so  the first character. 

I have an almost exactly similar file that doesn't produce the error. It is also super strange because when I was working on this file yesterday and saved it in this state, it was rendering fine. But right now as I run the app again..I get that error. Does anyone know what it could be?

One answer from S/O: This error happens when there is no import or export statement in a file (these make a file a module). But I do have an export statement..
## [4][How to share common typings in a project? [Lerna / Workspaces]](https://www.reddit.com/r/typescript/comments/hor78i/how_to_share_common_typings_in_a_project_lerna/)
- url: https://www.reddit.com/r/typescript/comments/hor78i/how_to_share_common_typings_in_a_project_lerna/
---
I did bit of a searching and couldn't find a real example anywhere.  

I'm about to create a project (both backend and frontend) that might share common typings (especially models etc). I haven't set up the lerna yet (the only share-able components are these typings at the moment). What's the best way to organize this? Can someone shed some light on this?  

Ideally, I guess, this common types would export two modules (if I understand the `module`s correctly) that matches the package names of the frontend and backend packages.

It'd be great if I get some ELI5 on `module` vs `namespace`, and how to organize typings that spans across multiple files.

Thanks!  

Edit: Big thanks to u/intrepidsovereign for creating a [boilerplate](https://github.com/RyanChristian4427/example-monorepo-ts)
## [5][Typescript generating a type error when calling a function (in a variable) from an event handler](https://www.reddit.com/r/typescript/comments/hos6d2/typescript_generating_a_type_error_when_calling_a/)
- url: https://www.reddit.com/r/typescript/comments/hos6d2/typescript_generating_a_type_error_when_calling_a/
---
I have the following class component:  


    class MyComponent extends Component&lt;MyProps&gt; {
    ...
       public componentDidMount() {
       ...
          window.addEventListener('resize', this.resizeHandler)
       }
       
       private sendResizeInfo = _.debounce((info) =&gt; {...});
    
       private resizeHandler() {
          this.sendResizeInfo('extra_info')
       }
    ...
    }

I get the following error:

`Uncaught [TypeError: sendResizeInfo is not a function]`

when the resizeHandler is called. Does anyone know why?
## [6][Help with falsey value in union type?](https://www.reddit.com/r/typescript/comments/hoq4dp/help_with_falsey_value_in_union_type/)
- url: https://www.reddit.com/r/typescript/comments/hoq4dp/help_with_falsey_value_in_union_type/
---
My error:

    if (obj.details.onDay === 0) {
     // ^^ throws TS error ^^
    }
    
    (property) onDay: 1 | 2 | 3 | 4 | 5 | 6 | "Day" | "Weekday" | "Weekend day"
    This condition will always return 'false' since the types '1 | 2 | 3 | 4 | 5 | 6 | "Day" | "Weekday" | "Weekend day"' and '0' have no overlap.ts(2367)

Types and structures:

    monthly: {
        // ...
    
        details?: {
          // ...
    
          onDay: MonthlyDayOfWeekInterval;
        }
      }
    
    export type MonthlyDayOfWeekInterval = DayOfWeekInterval | 'Day' | 'Weekday' | 'Weekend day';
    export type DayOfWeekInterval = 0 | 1 | 2 | 3 | 4 | 5 | 6;

I am new to union types, my only assumption is because it's a falsey 0?
## [7][Help with using Parameters&lt;T&gt; utility type while calling methods via string method names in generic function?](https://www.reddit.com/r/typescript/comments/hos02f/help_with_using_parameterst_utility_type_while/)
- url: https://www.reddit.com/r/typescript/comments/hos02f/help_with_using_parameterst_utility_type_while/
---
Hi everyone,

I'm running into an issue while using the `Parameters&lt;T&gt;` utility type while trying to write a function that can call dynamically call methods via their method names in a type-safe way.

This is what I have so far:

```
class Foo {
  greet(name: string) {
    console.log(`Hello, ${name}!`);
  }

  greetNumber(num: number) {
    console.log(`Hello, ${num} v2!`);
  }
}

function callFooDynamic&lt;K extends "greet" | "greetNumber"&gt;(
  foo: Foo,
  target: K,
  ...args: Parameters&lt;Foo[K]&gt;
) {
  foo[target](...args);
}

const foo = new Foo();
callFooDynamic(foo, "greet", "world");
```

However, I get this error when compiling:

```
error TS2556: Expected 1 arguments, but got 0 or more.

20   foo[target](...args);
                 ~~~~~~~

  src/lib/dynamic.ts:4:9
    4   greet(name: string) {
              ~~~~~~~~~~~~
    An argument for 'name' was not provided.

```

I would've expected that the `args` array would have been of the correct type, not sure why TypeScript isn't inferring the type correctly? Thanks in advance for your help!
## [8][Q: Why aren’t TS private methods, actually private?](https://www.reddit.com/r/typescript/comments/hoctyy/q_why_arent_ts_private_methods_actually_private/)
- url: https://www.reddit.com/r/typescript/comments/hoctyy/q_why_arent_ts_private_methods_actually_private/
---
It’s compile time only. If you use compiled code in vanilla JavaScript, you can access the private instance properties like any other property.

It could compile classes down to old function style classes, and put all the private members inside the function, and then not attach them to the instance.

Is it because typescript would then have to also “adjust” all this.&lt;private-thing&gt; references? That seems... arbitrary as a blocker.

What other problems would their be?
## [9][Anyone willing to mentor?](https://www.reddit.com/r/typescript/comments/ho63g6/anyone_willing_to_mentor/)
- url: https://www.reddit.com/r/typescript/comments/ho63g6/anyone_willing_to_mentor/
---
Not exactly mentor, but I guess more of a guide. I am fairly capable of learning by myself but there are times where I don’t understand a concept and it would be easier to have someone walk me thru and explain to me. I’m not looking for someone for me to ask questions 24/7 but rather someone who is willing to answer the questions I have from time to time. The TS discord is great but I often get ignored when I ask a question
## [10][Optional parameters, specifying undefined vs null](https://www.reddit.com/r/typescript/comments/ho736x/optional_parameters_specifying_undefined_vs_null/)
- url: https://www.reddit.com/r/typescript/comments/ho736x/optional_parameters_specifying_undefined_vs_null/
---
       protected combineAndSave(
          audiosAndPauseFiles: Array&lt;string&gt;
          , savePath: string
          , filePrefix?: string
          , fileName?: string
       ): void {

I noticed when calling this, I had to specify `undefined` on the 3rd arg (4th will be implicitly undefined if not specified). Is that the best way to do it? In general I think `null` is more explicit but maybe that does apply for this case.

Also if anyone knows a way to more tightly constrain this method so that one of the last two params must be specifed, I appreciate any further advice on that front.
## [11][How to make compiled JS enforce required properties?](https://www.reddit.com/r/typescript/comments/ho69ir/how_to_make_compiled_js_enforce_required/)
- url: https://www.reddit.com/r/typescript/comments/ho69ir/how_to_make_compiled_js_enforce_required/
---
I am trying to make an API and TS looks like a good way to do it as it's sort of self-documenting.  

However, when I try to test whether the compiled JS throws an error when the user does not include a required param/property, it does not and just gives it a value of undefined.

1. Is there a way to force the compiled JS to throw errors if the user is using the API incorrectly?
2. Is there a better 'standard' way to enforce correct usage of the API by the user rather than throwing errors?

&amp;#x200B;

Example code:

    // index.ts
    class Example {
      req1: string;
      req2: string;
    
      constructor(req1: string, req2: string) {
        this.req1 = req1;
        this.req2 = req2;
      }
    }
    
    module.exports = Example;
    
    /** ts file is compiled into index.js */
    
    // JS User that consumes index.js
    const Example = require('index');
    
    const ex1 = new Example("Hello"); // Expect API to throw error here as req2 was not passed
    console.log(ex1); // Actual: Example { req1: "Hello", req2: undefined }
    
    const ex2 = new Example("Hello", 2); // Expect API to throw error here as req2 is an incorrect type
    console.log(ex2); // Actual: Example { req1: "Hello", req2: 2 }

&amp;#x200B;

The only way I've found to enforce error throwing is doing type checks, etc. and throwing my own errors in the constructor, but then that's like regular JS, isn't it?
