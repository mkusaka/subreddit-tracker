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
## [2][Q: Why aren’t TS private methods, actually private?](https://www.reddit.com/r/typescript/comments/hoctyy/q_why_arent_ts_private_methods_actually_private/)
- url: https://www.reddit.com/r/typescript/comments/hoctyy/q_why_arent_ts_private_methods_actually_private/
---
It’s compile time only. If you use compiled code in vanilla JavaScript, you can access the private instance properties like any other property.

It could compile classes down to old function style classes, and put all the private members inside the function, and then not attach them to the instance.

Is it because typescript would then have to also “adjust” all this.&lt;private-thing&gt; references? That seems... arbitrary as a blocker.

What other problems would their be?
## [3][Anyone willing to mentor?](https://www.reddit.com/r/typescript/comments/ho63g6/anyone_willing_to_mentor/)
- url: https://www.reddit.com/r/typescript/comments/ho63g6/anyone_willing_to_mentor/
---
Not exactly mentor, but I guess more of a guide. I am fairly capable of learning by myself but there are times where I don’t understand a concept and it would be easier to have someone walk me thru and explain to me. I’m not looking for someone for me to ask questions 24/7 but rather someone who is willing to answer the questions I have from time to time. The TS discord is great but I often get ignored when I ask a question
## [4][Optional parameters, specifying undefined vs null](https://www.reddit.com/r/typescript/comments/ho736x/optional_parameters_specifying_undefined_vs_null/)
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
## [5][How to make compiled JS enforce required properties?](https://www.reddit.com/r/typescript/comments/ho69ir/how_to_make_compiled_js_enforce_required/)
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
## [6][Are Promise and Array type constructors?](https://www.reddit.com/r/typescript/comments/hobm5f/are_promise_and_array_type_constructors/)
- url: https://www.reddit.com/r/typescript/comments/hobm5f/are_promise_and_array_type_constructors/
---
I know the utility type constructors are, but not these two.

If so, what other notable ones are there?
## [7][Is ReturnType a good utility when you're not sure about the return type?](https://www.reddit.com/r/typescript/comments/hobs51/is_returntype_a_good_utility_when_youre_not_sure/)
- url: https://www.reddit.com/r/typescript/comments/hobs51/is_returntype_a_good_utility_when_youre_not_sure/
---
I've used it a couple times when the docs and google searches just couldn't identify the obscure type being returned from some method (I'm looking at you, Google Cloud).

       tempFolder: ReturnType&lt;typeof tmp.dirSync&gt;

I feel like it's a good fallback, a "You tell me what it is, Typescript". On the other hand it's less insightful from a quick glance compared to an explicit type declaration.

Do you guys use this one often? Why or why not? I'm curious what others consider the boundary to be between acceptable use and overuse.
## [8][Moving from TypeScript to Rust / WebAssembly](https://www.reddit.com/r/typescript/comments/hnw5v6/moving_from_typescript_to_rust_webassembly/)
- url: https://nicolodavis.com/blog/typescript-to-rust/
---

## [9][Extracting and Hashing Lazy-Loaded CSS in Angular](https://www.reddit.com/r/typescript/comments/hnz6r0/extracting_and_hashing_lazyloaded_css_in_angular/)
- url: https://volosoft.com/blog/Extracting-and-Hashing-Lazy-Loaded-CSS-in-Angular
---

## [10][Is it possible to type a function that has been declared with the function keyword?](https://www.reddit.com/r/typescript/comments/hnodve/is_it_possible_to_type_a_function_that_has_been/)
- url: https://www.reddit.com/r/typescript/comments/hnodve/is_it_possible_to_type_a_function_that_has_been/
---
The TypeScript documentation includes lots of examples of typing arrow functions, but is it possible to use type when using the `function` keyword?

```js
type T = (n: string) =&gt; string

const arrowGreet: T = (name) =&gt; `Hello ${name}`

// How can I type fnGreet as T?
function fnGreet(name) {
    return `Hello ${name}`
}
```
## [11][How should I manage type definition for REST calls?](https://www.reddit.com/r/typescript/comments/hnbrrf/how_should_i_manage_type_definition_for_rest_calls/)
- url: https://www.reddit.com/r/typescript/comments/hnbrrf/how_should_i_manage_type_definition_for_rest_calls/
---
For example, let's say my app allow users to do CRUD operation on country resources.
My GET response would be an array of Countries, typed like this:
        
    interface Country {
      states: State[];
      name: string;
      id: number;
      abbreviation: string;
    }
    
    interface State {
      name: string;
      id: number;
      abbreviation: string;
      cities: City[];
    }
    
    interface City {
      name: string;
      id: number;
      abbreviation: string;
    }


Let's assume that creating a Country involves the user sending States and Cities in the same payload.

Since we're creating resources, they wouldn't have an id assigned to them yet, it doesn't make sense for the POST payload to contain ids. To adequately type my POST payload, I'd have to duplicate Country, State, City, and remove Id or make it an optional field.

    interface PostCountry {
      states: State[];
      name: string;
      abbreviation: string;
    }
    
    interface PostState {
      name: string;
      abbreviation: string;
      cities: City[];
    }
    
    interface PostCity {
      name: string;
      abbreviation: string;
    }


I'm new to Typescript and is currently converting an app to Typescript. How do you guys manage something like this? Is it inevitable to have multiple type definitions for each REST operation?
