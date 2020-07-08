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
## [2][How should I manage type definition for REST calls?](https://www.reddit.com/r/typescript/comments/hnbrrf/how_should_i_manage_type_definition_for_rest_calls/)
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
## [3][Dynamic string types](https://www.reddit.com/r/typescript/comments/hnh2up/dynamic_string_types/)
- url: https://www.reddit.com/r/typescript/comments/hnh2up/dynamic_string_types/
---
I've got a type
    type bin = 0|1;
I'd like to make a type that would be
    type key = bin + '_' + bin;
This is to access properties of an object that has keys like 0_1
Is this possible
## [4][GUI live editor for complex Object](https://www.reddit.com/r/typescript/comments/hngywz/gui_live_editor_for_complex_object/)
- url: https://www.reddit.com/r/typescript/comments/hngywz/gui_live_editor_for_complex_object/
---
Hi everyone, I'm searching a GUI editor for complex objects. I tried dat.GUI but it's not working well with arrays and optional (undefined) properties. Anyone can suggest me an alternative or at least knows a package like this? I'm starting to create one on my own but it's a lot of work and I have other priorities. If you want to see what I need to edit checkout this: https://github.com/matteobruni/tsparticles/wiki/tsParticles-Options
## [5][How to get auto completion for file paths](https://www.reddit.com/r/typescript/comments/hnd28j/how_to_get_auto_completion_for_file_paths/)
- url: https://www.reddit.com/r/typescript/comments/hnd28j/how_to_get_auto_completion_for_file_paths/
---
When typing import or require statements my editor offers autocompletion for the file path
Is there a way to achieve this elsewhere for the fs module for example.
In case it's relevant I use emacs with tide
## [6][Assigning only the properties which is defined in interface](https://www.reddit.com/r/typescript/comments/hn0xu8/assigning_only_the_properties_which_is_defined_in/)
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
## [7][Priority Queue in TypeScript](https://www.reddit.com/r/typescript/comments/hmsedt/priority_queue_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/hmsedt/priority_queue_in_typescript/
---
 Queues and stacks are used in programming to process a collection of items in a certain order. Those data structures work either with a *first-in-first-out* or *last-in-first-out* system in which the order of processing is always directly coupled to the order of insertion. However, sometimes you need to have more control over the processing order then offered by the insertion order. 

[https://medium.com/@wim\_jongeneel/priority-queue-in-typescript-6ef23116901?source=friends\_link&amp;sk=a35b45b2a26f4cd5daafb84feafb6c88](https://medium.com/@wim_jongeneel/priority-queue-in-typescript-6ef23116901?source=friends_link&amp;sk=a35b45b2a26f4cd5daafb84feafb6c88)
## [8][Newbie question for dev coming from static world.](https://www.reddit.com/r/typescript/comments/hmvdux/newbie_question_for_dev_coming_from_static_world/)
- url: https://www.reddit.com/r/typescript/comments/hmvdux/newbie_question_for_dev_coming_from_static_world/
---
I am new to typescript so I apologize if this sounds dumb. 

I can not figure out how to get types to work with any modules I have not written. Here an example. 

I am trying to use the http library to set up a simple server with one endpoint. I have tsc setup and @types/node installed but my code still isn’t providing me types when using vscode. Is there something I have to do to get the types to show up in intellisense? When I use app.on("request", (req, res) =&gt; {}) my IDE doesn't know that the listener is a RequestListener type. Should it auto complete to that or do I have to manually search for every type in the @types/package?

I went to the @types/nodes page on GitHub and their is no documentation their. Is that normal for types? Is their a rule of thumb of where I can find what type different objects and functions are in any given package?
## [9][Return an object with a generic string key](https://www.reddit.com/r/typescript/comments/hmw8ee/return_an_object_with_a_generic_string_key/)
- url: https://www.reddit.com/r/typescript/comments/hmw8ee/return_an_object_with_a_generic_string_key/
---
I have a generic function that will return data in an object like:
```ts
{
   data: T[],
   isLoading: boolean,
   error: Error | null,
}
```
This correctly gives me a return type of `T[]` for data.

I want to be able to pass a key name as a `string` so that the key does not always have to be data.

When I change my function to take in a `dataKey: string` and try to return the same object like: 
```ts
{
   [dataKey]: data,
   error,
   isLoading,
}
```

I get typescripts implied return type of 
```ts
{
    [x: string]: boolean | T[] | Error | null;
    error: Error | null;
    isLoading: boolean;
}
```

It seems like when I use the dynamic [dataKey], then my `data` type takes all the return object's possible types. What am I doing wrong here? Thank you
## [10][Best way to setup reloading ?](https://www.reddit.com/r/typescript/comments/hmpxnn/best_way_to_setup_reloading/)
- url: https://www.reddit.com/r/typescript/comments/hmpxnn/best_way_to_setup_reloading/
---
What is your preferred choice:

1. tsc + node
2. nodemon + ts-node
3. ts-node-dev
4. ?
## [11][How can i achieve object copy ??](https://www.reddit.com/r/typescript/comments/hmp6zr/how_can_i_achieve_object_copy/)
- url: https://www.reddit.com/r/typescript/comments/hmp6zr/how_can_i_achieve_object_copy/
---
const target = { a: 1, b: { c: 2} };

const source = { a: 4, b: { d: 3} };

const res = Object.assign(target, source);

console.log(res)

//Output :   { a: 4, b: { d: 3 } };

//Expected : { a: 4, b: { c: 2, d: 3 } };

Any help ???
