# typescript
## [1][Who's hiring Typescript developers - February](https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/)
- url: https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/
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
## [2][Typescript I love you but this documentation is ironic](https://www.reddit.com/r/typescript/comments/ewjhiy/typescript_i_love_you_but_this_documentation_is/)
- url: https://i.redd.it/0xaegxrcu1e41.jpg
---

## [3][Is there a way to filter objects in an array using string from an array?](https://www.reddit.com/r/typescript/comments/ewxjlh/is_there_a_way_to_filter_objects_in_an_array/)
- url: https://www.reddit.com/r/typescript/comments/ewxjlh/is_there_a_way_to_filter_objects_in_an_array/
---
I don’t init is that makes sense so I’m going to try and explain it further;

I have an array of lectures and an array of lecture IDs.

I want to filter a list of all lectures down to an array of only the ones that match the IDs of the ones in the array of IDs.
## [4][Is letting .js files generate side by side with .ts a bad practice?](https://www.reddit.com/r/typescript/comments/ewsi3a/is_letting_js_files_generate_side_by_side_with_ts/)
- url: https://www.reddit.com/r/typescript/comments/ewsi3a/is_letting_js_files_generate_side_by_side_with_ts/
---
It's the simplest pattern but also doubles the file volume per folder. What do you guys think about leaving projects that way vs specifying an output directory? What would guide your decision to move to the later or some other convention?
## [5][Need help with dynamic properties](https://www.reddit.com/r/typescript/comments/eworn8/need_help_with_dynamic_properties/)
- url: https://www.reddit.com/r/typescript/comments/eworn8/need_help_with_dynamic_properties/
---
Hi. I have a type A with several properties which can be of different keys: strings, numbers, types B, irrelevant. I want to copy over certain fields from object a1 being type A to object a2, that is also of type A, but I do not know in advance which fields i want to copy over. The fields that I want to copy over I get in an array of string literals which tell me which field i want to copy.   
I do not understand how to "type" this correctly..

Example on codesandbox: [https://codesandbox.io/s/typescript-playground-export-fjol9](https://codesandbox.io/s/typescript-playground-export-fjol9)  


https://preview.redd.it/sk2k3apwc4e41.png?width=738&amp;format=png&amp;auto=webp&amp;s=9dd30a17952037ccb437a9a0b4887fe080029d30
## [6][Adding type safety to a HTTP API client both compile-time and runtime](https://www.reddit.com/r/typescript/comments/ewbx66/adding_type_safety_to_a_http_api_client_both/)
- url: https://medium.com/smartly-io/oats-how-we-learned-to-stop-worrying-and-love-types-aa0041aaa9cc
---

## [7][TS1005 error on unary assignment, what is the error here?](https://www.reddit.com/r/typescript/comments/ew9p6b/ts1005_error_on_unary_assignment_what_is_the/)
- url: https://www.reddit.com/r/typescript/comments/ew9p6b/ts1005_error_on_unary_assignment_what_is_the/
---
    let port: number = 8080;
    
    // TS1005: ';' expected.
    process.env.PORT &amp;&amp; port = parseInt(process.env.PORT!);

Typescript couldn't detect that it only runs if POST is defined, so the ! was needed. But the error above still lints on the assignment ( = ) operator. What is the issue exactly and how should I fix this?
## [8][Access values in a string array class property through an import](https://www.reddit.com/r/typescript/comments/ewceyp/access_values_in_a_string_array_class_property/)
- url: https://www.reddit.com/r/typescript/comments/ewceyp/access_values_in_a_string_array_class_property/
---
    // I need to access the value in: new Nouns().getNouns()[0];
    export class Nouns {
        constructor(){
        }
    
       getNouns():string[] {
          return this.nouns;
        }
    
        public nouns:string[] = [
        "Armour",
        "Barrymore",
        "Cabot",
        "Catholicism",
        "Chihuahua",
        ........
        ]

Here is my attempt :   


    import { Nouns } from "./nouns";
    import { Adjectives } from "./adjectives";
    
     constructor() { 
        this.myName = this.generateName();
        console.log(`A new Turtle ${this.myName} was born !`); 
        console.log(Nouns);
      } // The result : A new Turtle undefined undefined was born !
    
    
      generateName(): string {
        let nameArray:string[] = new Nouns().getNouns();
        let adjectiveArray:string[] = new Adjectives().getAdjectives();
        let randomSeed:number = Math.random()*adjectiveArray.length;
        let randomSeed2:number = Math.random()*nameArray.length;
    
        return `${adjectiveArray[randomSeed]} ${nameArray[randomSeed2]}`;
      }
    
    

What am I doing wrong here ?
## [9][A fully typed CLI library with GUI support](https://www.reddit.com/r/typescript/comments/evswnd/a_fully_typed_cli_library_with_gui_support/)
- url: https://github.com/hediet/ts-cli/blob/master/cli/README.md
---

## [10][C# like Namespaces for folders?](https://www.reddit.com/r/typescript/comments/evzjww/c_like_namespaces_for_folders/)
- url: https://www.reddit.com/r/typescript/comments/evzjww/c_like_namespaces_for_folders/
---
Namespaces in C# are tied to folder structure, for example ``using Project.Game.Models`` can import all classes from models.

```
/game
-- /foo.tsx
-- /bar.tsx
-- /models
----- /tree.tsx
----- /car.tsx
/app.tsx
```

Is there a way to do this in TypeScript by exporting/importing/declaring namespaces rather than modeules?

App.tsx
```
import {Game} from './game'; &lt;-- This does not work?
import {Game} from './game/models/tree.tsx'; &lt;-- This works but you can't import multiple classes?

const App: React.FC = () =&gt; {
    return (
       &lt;Game.FooClass.FooComponent fun={true} /&gt;
       &lt;Game.Models.TreeClass.TreeComponent visible={true} /&gt;
    )
```

Tree.tsx
```
export namespace Game.Models {
    export class Tree { ... }
}
```
## [11][An Angular Based Web App with two forms maintains its state between steps and validates the user's input](https://www.reddit.com/r/typescript/comments/ew2mvj/an_angular_based_web_app_with_two_forms_maintains/)
- url: https://github.com/shpotainna/passengers
---

