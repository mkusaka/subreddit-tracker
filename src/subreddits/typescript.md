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
## [2][Moving from TypeScript to Rust / WebAssembly](https://www.reddit.com/r/typescript/comments/hnw5v6/moving_from_typescript_to_rust_webassembly/)
- url: https://nicolodavis.com/blog/typescript-to-rust/
---

## [3][Extracting and Hashing Lazy-Loaded CSS in Angular](https://www.reddit.com/r/typescript/comments/hnz6r0/extracting_and_hashing_lazyloaded_css_in_angular/)
- url: https://volosoft.com/blog/Extracting-and-Hashing-Lazy-Loaded-CSS-in-Angular
---

## [4][Is it possible to type a function that has been declared with the function keyword?](https://www.reddit.com/r/typescript/comments/hnodve/is_it_possible_to_type_a_function_that_has_been/)
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
## [5][How should I manage type definition for REST calls?](https://www.reddit.com/r/typescript/comments/hnbrrf/how_should_i_manage_type_definition_for_rest_calls/)
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
## [6][Why is this object possibly null?](https://www.reddit.com/r/typescript/comments/hnnavt/why_is_this_object_possibly_null/)
- url: https://www.reddit.com/r/typescript/comments/hnnavt/why_is_this_object_possibly_null/
---
I'm sure part of the answer is that there may be no match, hence blank array, which I guess counts as an empty object. Is that right?

If so, does the null type equaling an object bug in JS not apply to TS?

       public fileName: string;
    
       constructor(
          contentType: contentTypes
          , fileName: string
          , folderPath: string
       ) {
          this.contentType = this.getContentType();
          this.fileName = fileName;
          this.folderPath = folderPath;
    
          this.fullFilePath = this.setFullFolderPath();
       }
    
       // --------------- Internal methods
    
       private getContentType() {
          // grab the first number
          const firstNumberPattern: RegExp = /^\d+/;
          const firstNumber: number | null = 
        Number(this!.fileName!.match(firstNumberPattern)[0]);
        /*
          this: this
          Object is possibly 'null'.ts(2531)
        */

How should the error be cleared? Seems wierd that the error is on `this`. And `!` is not asserting not null, as expected. And adding a union of null onto `firstNumber` also didn't help.
## [7][Dynamic string types](https://www.reddit.com/r/typescript/comments/hnh2up/dynamic_string_types/)
- url: https://www.reddit.com/r/typescript/comments/hnh2up/dynamic_string_types/
---
I've got a type
    type bin = 0|1;
I'd like to make a type that would be
    type key = bin + '_' + bin;
This is to access properties of an object that has keys like 0_1
Is this possible
## [8][GUI live editor for complex Object](https://www.reddit.com/r/typescript/comments/hngywz/gui_live_editor_for_complex_object/)
- url: https://www.reddit.com/r/typescript/comments/hngywz/gui_live_editor_for_complex_object/
---
Hi everyone, I'm searching a GUI editor for complex objects. I tried dat.GUI but it's not working well with arrays and optional (undefined) properties. Anyone can suggest me an alternative or at least knows a package like this? I'm starting to create one on my own but it's a lot of work and I have other priorities. If you want to see what I need to edit checkout this: https://github.com/matteobruni/tsparticles/wiki/tsParticles-Options
## [9][How to get auto completion for file paths](https://www.reddit.com/r/typescript/comments/hnd28j/how_to_get_auto_completion_for_file_paths/)
- url: https://www.reddit.com/r/typescript/comments/hnd28j/how_to_get_auto_completion_for_file_paths/
---
When typing import or require statements my editor offers autocompletion for the file path
Is there a way to achieve this elsewhere for the fs module for example.
In case it's relevant I use emacs with tide
## [10][Assigning only the properties which is defined in interface](https://www.reddit.com/r/typescript/comments/hn0xu8/assigning_only_the_properties_which_is_defined_in/)
- url: https://www.reddit.com/r/typescript/comments/hn0xu8/assigning_only_the_properties_which_is_defined_in/
---
Hi,
Is it possible to add only the properties which is defined in the interface?

Eg: 


Interface Car {


model:string 


year:number


 }


// req={model:"gla",year:2020,brand:"Tesla"}

function abc(req){


let car:Car



}



I need to assign only the properties which are defined in the interface. The method I'm currently using is car.model =req.model. Now the object is really long so the code is getting bigger and bigger.
## [11][Priority Queue in TypeScript](https://www.reddit.com/r/typescript/comments/hmsedt/priority_queue_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/hmsedt/priority_queue_in_typescript/
---
 Queues and stacks are used in programming to process a collection of items in a certain order. Those data structures work either with a *first-in-first-out* or *last-in-first-out* system in which the order of processing is always directly coupled to the order of insertion. However, sometimes you need to have more control over the processing order then offered by the insertion order. 

[https://medium.com/@wim\_jongeneel/priority-queue-in-typescript-6ef23116901?source=friends\_link&amp;sk=a35b45b2a26f4cd5daafb84feafb6c88](https://medium.com/@wim_jongeneel/priority-queue-in-typescript-6ef23116901?source=friends_link&amp;sk=a35b45b2a26f4cd5daafb84feafb6c88)
