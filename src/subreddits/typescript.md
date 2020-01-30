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
## [2][A fully typed CLI library with GUI support](https://www.reddit.com/r/typescript/comments/evswnd/a_fully_typed_cli_library_with_gui_support/)
- url: https://github.com/hediet/ts-cli/blob/master/cli/README.md
---

## [3][C# like Namespaces for folders?](https://www.reddit.com/r/typescript/comments/evzjww/c_like_namespaces_for_folders/)
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
## [4][An Angular Based Web App with two forms maintains its state between steps and validates the user's input](https://www.reddit.com/r/typescript/comments/ew2mvj/an_angular_based_web_app_with_two_forms_maintains/)
- url: https://github.com/shpotainna/passengers
---

## [5][How to check an Array of objects contains matching values with enum](https://www.reddit.com/r/typescript/comments/evrcei/how_to_check_an_array_of_objects_contains/)
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
## [6][How to globally load types from a package that typically isn't globally-loaded?](https://www.reddit.com/r/typescript/comments/evpy5n/how_to_globally_load_types_from_a_package_that/)
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
## [7][Using redux-saga with Typescipt](https://www.reddit.com/r/typescript/comments/evg6rd/using_reduxsaga_with_typescipt/)
- url: https://www.reddit.com/r/typescript/comments/evg6rd/using_reduxsaga_with_typescipt/
---
I've been using `redux-thunk` for quite a while on a project of mine, but I've recently outgrown it and have turned to other libraries to handle async actions (mostly network requests). Of the options out there, I've found `redux-saga` and `redux-observable` to be the most promising options. After some research, I found that `redux-saga` seemed the best for my use-case. 

I then installed saga and the accompanying types, but I quickly found that saga's integration with TypeScript was less than ideal. For example, using a selector returns an `any` type as opposed to the proper type which corresponds to the output of the selector.

I did some follow-up research on this topic, and from what I can see, saga doesn't support this kind of type safety -- something that `redux-observables` does support. Is there some kind of work-around or alternate `@types` library which allows saga to support real type safety with TypeScript? 

I was excited to use saga, but if proper type checks do not properly work, I will be forced to use `redux-observable`.
## [8][Searching for JavaScript/TypeScript senior dev who can share some Best Practices](https://www.reddit.com/r/typescript/comments/evbddt/searching_for_javascripttypescript_senior_dev_who/)
- url: https://www.reddit.com/r/typescript/comments/evbddt/searching_for_javascripttypescript_senior_dev_who/
---
I am searching for JavaScript/TypeScript senior dev who can share some Best Practices and answer some of my questions. I can surely pay you for it :)

Topics:

1. Runtime type checking, best practices for JSOn parsing
2. Shared code for server/client apps (interfaces, enums, locales)

Medium (or similar) articles are too basic...
## [9][Get available properties of interface without initializing](https://www.reddit.com/r/typescript/comments/ev806n/get_available_properties_of_interface_without/)
- url: https://stackoverflow.com/questions/59953416/get-available-properties-of-interface-without-initializing
---

## [10][TypeScript advice](https://www.reddit.com/r/typescript/comments/ev6sd5/typescript_advice/)
- url: https://medium.com/@Armandotrue/typescript-advice-4a46b6ae5a7
---

## [11][Understanding any and unknown in TypeScript. Difference between never and void](https://www.reddit.com/r/typescript/comments/euoyos/understanding_any_and_unknown_in_typescript/)
- url: https://wanago.io/2020/01/27/understanding-any-and-unknown-in-typescript-difference-between-never-and-void/
---

