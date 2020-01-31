# typescript
## [1][Who's hiring Typescript developers - January](https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/)
- url: https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/
---
The monthly thread for people to post openings at their companies.

* Please state the job location and include the keywords REMOTE, INTERNS and/or VISA when the corresponding sort of candidate is welcome. When remote work is not an option, include ONSITE.

* Please only post if you personally are part of the hiring company—no recruiting firms or job boards **Please report recruiters or job boards**. 

* Only one post per company. 

* If it isn't a household name, explain what your company does. Sell it.

* Please add the company email that applications should be sent to, or the companies application web form/job posting (needless to say this should be on the company website, not a third party site).


Commenters: please don't reply to job posts to complain about something. It's off topic here.

Readers: please only email if you are personally interested in the job. 

[Previous Hiring Threads](https://www.reddit.com/r/typescript/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AMonthly%2BHiring%2BThread)
## [2][Typescript I love you but this documentation is ironic](https://www.reddit.com/r/typescript/comments/ewjhiy/typescript_i_love_you_but_this_documentation_is/)
- url: https://i.redd.it/0xaegxrcu1e41.jpg
---

## [3][Adding type safety to a HTTP API client both compile-time and runtime](https://www.reddit.com/r/typescript/comments/ewbx66/adding_type_safety_to_a_http_api_client_both/)
- url: https://medium.com/smartly-io/oats-how-we-learned-to-stop-worrying-and-love-types-aa0041aaa9cc
---

## [4][TS1005 error on unary assignment, what is the error here?](https://www.reddit.com/r/typescript/comments/ew9p6b/ts1005_error_on_unary_assignment_what_is_the/)
- url: https://www.reddit.com/r/typescript/comments/ew9p6b/ts1005_error_on_unary_assignment_what_is_the/
---
    let port: number = 8080;
    
    // TS1005: ';' expected.
    process.env.PORT &amp;&amp; port = parseInt(process.env.PORT!);

Typescript couldn't detect that it only runs if POST is defined, so the ! was needed. But the error above still lints on the assignment ( = ) operator. What is the issue exactly and how should I fix this?
## [5][Access values in a string array class property through an import](https://www.reddit.com/r/typescript/comments/ewceyp/access_values_in_a_string_array_class_property/)
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
## [6][A fully typed CLI library with GUI support](https://www.reddit.com/r/typescript/comments/evswnd/a_fully_typed_cli_library_with_gui_support/)
- url: https://github.com/hediet/ts-cli/blob/master/cli/README.md
---

## [7][C# like Namespaces for folders?](https://www.reddit.com/r/typescript/comments/evzjww/c_like_namespaces_for_folders/)
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
## [8][An Angular Based Web App with two forms maintains its state between steps and validates the user's input](https://www.reddit.com/r/typescript/comments/ew2mvj/an_angular_based_web_app_with_two_forms_maintains/)
- url: https://github.com/shpotainna/passengers
---

## [9][How to check an Array of objects contains matching values with enum](https://www.reddit.com/r/typescript/comments/evrcei/how_to_check_an_array_of_objects_contains/)
- url: https://www.reddit.com/r/typescript/comments/evrcei/how_to_check_an_array_of_objects_contains/
---
Hello, I'm trying to check the API response in my angular service(array of objects) against Enum, so If the API returns items matching to the values in the enum, I want to filter matching ones.

here's my API call:

     return this.angularService.getRoles().pipe(map(resp =&gt; {
          let x = Object.keys(myEnum); // makes enum array
          console.log('resp',resp.filter(test =&gt; test['name'] === x )); // matching items are not filtered
          return (some values);
        }))

and my enum:

    export enum myEnum {
        User = 'User',
        Admin = 'Admin',
        Reporter = 'Reporter'
      }
## [10][How to globally load types from a package that typically isn't globally-loaded?](https://www.reddit.com/r/typescript/comments/evpy5n/how_to_globally_load_types_from_a_package_that/)
- url: https://www.reddit.com/r/typescript/comments/evpy5n/how_to_globally_load_types_from_a_package_that/
---
I prefer to use `import` only to include actual functionality in a file, and to use my `tsconfig` to include all necessary types.

Here's an example reducer within my app:

    import { createReducer } from '@ngrx/store';
    
    const initialState: State = ['a', 'b', 'c', 'd'];
    
    const reducer = createReducer(
      initialState,
    );
    
    export function appReducer(state: State, action: Action): State {
      return reducer(state, action);
    }
    
I've declared `State` in `custom_types/typings.d.ts`, and have included `custom_types` within `tsconfig.compileOptions.typeRoots`. This works fine.

`Action` is undefined. I could include `Action` in the `import` statement, but since it's only used for typing and not for functionality, I would prefer to include it through my `tsconfig`.

I've tried to globally include `@ngrx` typings in several ways:

- Adding `@ngrx/store` to `typeRoots`
- Adding `@ngrx/store/index.d.ts` to `typeRoots`
- Adding `import '@ngrx/store'` to `custom_types/typings.d.ts`

But the compiler continues to say that `Action` is undefined.

How can I globally include all `@ngrx` typings in my project?
## [11][Using redux-saga with Typescipt](https://www.reddit.com/r/typescript/comments/evg6rd/using_reduxsaga_with_typescipt/)
- url: https://www.reddit.com/r/typescript/comments/evg6rd/using_reduxsaga_with_typescipt/
---
I've been using `redux-thunk` for quite a while on a project of mine, but I've recently outgrown it and have turned to other libraries to handle async actions (mostly network requests). Of the options out there, I've found `redux-saga` and `redux-observable` to be the most promising options. After some research, I found that `redux-saga` seemed the best for my use-case. 

I then installed saga and the accompanying types, but I quickly found that saga's integration with TypeScript was less than ideal. For example, using a selector returns an `any` type as opposed to the proper type which corresponds to the output of the selector.

I did some follow-up research on this topic, and from what I can see, saga doesn't support this kind of type safety -- something that `redux-observables` does support. Is there some kind of work-around or alternate `@types` library which allows saga to support real type safety with TypeScript? 

I was excited to use saga, but if proper type checks do not properly work, I will be forced to use `redux-observable`.
