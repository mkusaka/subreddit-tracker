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
## [2][Implicit inferencing of generic type in class through constructor property does not work](https://www.reddit.com/r/typescript/comments/g3lftr/implicit_inferencing_of_generic_type_in_class/)
- url: https://www.reddit.com/r/typescript/comments/g3lftr/implicit_inferencing_of_generic_type_in_class/
---
Here is my [Code on TypeScript playground](https://www.typescriptlang.org/v2/en/play?#code/KYDwDg9gTgLgBASwHY2FAZgQwMbDgFVBgAUoIw0YFgBnOAbwCg4W4ByG4GKpAcxoB0IAJ4AvfMIpsAXHCQBXALYAjNAG5mrBDQCSKNFlwB5JLOUQIAG2CYkGgL6NGoSLDjYISGvDBkKsYT0AE1B8CABlGCh5bBh5KGAgo2UAKwBZTDBZejgAbWQQkEQkOABrYGEIdAIiUnJKahoAXVkFFTQ4ezgAXgZNFg4uHn4hMQkpWQAGABp+xF19DBxgE1kAJkZHZ3BoeGxLTBo6AHFgJDQEbD1UJdwAHlPzqEu6-ypaAD4+1ndPb2jYtAABS+BAAN0wqDgvnqAWCoQiURicQSyXSmWyeQKoGKZQqVTgjwu2FeDVoLTkSlUUE6AEoGFsfmB5MpLJc4LwuKTYI0gfSmD8fgB6IX0LraOAwCx44BgSUACwlNCQCDAFHgNHluwEcAASjZLHAAO7QSxBX4hY0QeRmuAJFElTDQvwdCCpYCxARzH72+IlMVwQ6Es7E7nvGgaH6OLYigCi4DZ2AQqEswjg8hoyF4NRAJBdPNocHQ0GDTxe+fDVqgpRojA8XngLkTyZ6cmARtLxOuBmWd0IubDjQ+IIrgSQhTCkQBKMSaIyYFpGnr3jgTYQydedF6a6TMAEnLzsPDfLUcBF8qDMEkeH7h7ejScIp0ijAzZTaeQ6AufE7zxJFcaBUyHkXh5QVPA3RSD0fDIMEEBCc1kHA34G2nEsgggQskAgeATWrOs-ngBAXzfVtzg7Ik-27W5gBHI8xwnRFp3iWdUnnRcCIbRASMuDc-C3bjX14vcD0HWgTzPIULzoK8KAYQZuCzQQRHEa8ZHTJBSmwo0kGmeZqMMFZTA0rSIB0+wgA) for interactive debugging

Here a Text version of my code:

    export interface TextProperties {
        'settings.xyzType': number;
        isInterfaceOn: boolean;
    }
    
    export const propertyIndexToStructuredObjMap: { [index in keyof TextProperties]: number } = {
        'settings.xyzType': 0,
        isInterfaceOn: 2
    }
    
    export class GenericInterface&lt;GenericProperties&gt; {
        constructor(private propertyIndexToStructureObjMap: { [index in keyof GenericProperties]: number }) {}
    
        public getProperties() {
            //{} is too keep this snippet short. Real world code would return a proper object.
            return {} as GenericProperties;
        }
    }
    
    //Explicitely using TextProperties for GenericProperties works
    const explicit = new GenericInterface&lt;TextProperties&gt;(propertyIndexToStructuredObjMap);
    const expliitProps = explicit.getProperties(); //has type TextProperties
    
    //Implicitely infering GenericProperties through the object provided in the constructor does not work
    const implicit = new GenericInterface(propertyIndexToStructuredObjMap);
    const implicitProps = implicit.getProperties(); //has type {'settings.xyzType': unknown, isInterfaceOn: unknown}

As described in the comments, submitting the type for the generic in the class explicitly makes it work like a charm and the getProperties function returns the correct type.

Implicit inferencing does not work completely. It returns an Object with the correct keys, but the types of it's values are all unknown.

Is there a way to make the implicit inferencing work in my case, so I don't have to explicitly state the type all the time?
## [3][Fancy Emitter - A new take on JavaScript's EventEmitter class. Strongly Typed and makes use of the newest features offered in ES6+](https://www.reddit.com/r/typescript/comments/g3nhha/fancy_emitter_a_new_take_on_javascripts/)
- url: https://github.com/mothepro/fancy-emitter
---

## [4][A friend sent me this comic. Had to add it to the internal TypeScript guide I wrote for the team](https://www.reddit.com/r/typescript/comments/g2ve22/a_friend_sent_me_this_comic_had_to_add_it_to_the/)
- url: https://i.imgur.com/YMo4iC0.jpg
---

## [5][Gact Store White Paper](https://www.reddit.com/r/typescript/comments/g3m9ts/gact_store_white_paper/)
- url: https://github.com/gactjs/store/blob/master/docs/white-paper.md
---

## [6][Gact Store React Bindings](https://www.reddit.com/r/typescript/comments/g3mbxj/gact_store_react_bindings/)
- url: https://github.com/gactjs/react-store
---

## [7][Gact Store](https://www.reddit.com/r/typescript/comments/g3mbkf/gact_store/)
- url: https://github.com/gactjs/store
---

## [8][ts lerna boilerplate](https://www.reddit.com/r/typescript/comments/g399q5/ts_lerna_boilerplate/)
- url: https://www.reddit.com/r/typescript/comments/g399q5/ts_lerna_boilerplate/
---
Hi,

I've created typescript (project reference, path mapping, ts-node-dev) + lerna + yarn workspace + jest (ts-jest) + eslint (\`overrides\` for supporting difference rules against js and ts) + etc boilerplate with explanation.

[https://github.com/jjangga0214/ts-yarn-lerna-boilerplate](https://github.com/jjangga0214/ts-yarn-lerna-boilerplate)

Hope it can be helpful to someone.
## [9][Why is redux's combine reducer complaining about method overload when using typescript and react-redux-firebase?](https://www.reddit.com/r/typescript/comments/g2yr2q/why_is_reduxs_combine_reducer_complaining_about/)
- url: https://www.reddit.com/r/typescript/comments/g2yr2q/why_is_reduxs_combine_reducer_complaining_about/
---
**The problem**

I have a basic project with just a minimal boilerplate, because I wanted to try out `react-redux-firebase`. So I set everything up according to the typescript example from here: https://github.com/prescottprue/react-redux-firebase/tree/master/examples/complete/typescript

But this part:

    const rootReducer = combineReducers&lt;RootState&gt;({
      firebase: firebaseReducer,
      firestore: firestoreReducer
    });

Is complaining:

    &gt; firestore: Reducer&lt;FirestoreReducer.Reducer, any&gt; No overload matches
    &gt; this call. Overload 1 of 3,

I made a minimal codesandbox demo to reproduce the problem:

https://codesandbox.io/s/great-perlman-p6g7p?file=/src/store/index.ts

You can see the error on line `54` altough the app is running just fine. On my local client it won't render, I just see the error in the browser (overload...). 

What am I doing wrong here? It has to be something about typescript, I'm a real beginner when it comes to typescript.
## [10][Strategies for migrating to TypeScript](https://www.reddit.com/r/typescript/comments/g31q57/strategies_for_migrating_to_typescript/)
- url: https://2ality.com/2020/04/migrating-to-typescript.html
---

## [11][Best approach to share a module with types definition between personal projects](https://www.reddit.com/r/typescript/comments/g2yiua/best_approach_to_share_a_module_with_types/)
- url: https://www.reddit.com/r/typescript/comments/g2yiua/best_approach_to_share_a_module_with_types/
---
Hello to everyone :)  
I've started a personal project a few months ago that consisted in a simple Angular app and a few cloud functions on Firebase.

Time has passed and now I'm struggling to keep all my type definitions synced between projects. 

Is there any best practice to share type definitions between projects without publish any package on npm?
