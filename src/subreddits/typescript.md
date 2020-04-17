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
## [2][A friend sent me this comic. Had to add it to the internal TypeScript guide I wrote for the team](https://www.reddit.com/r/typescript/comments/g2ve22/a_friend_sent_me_this_comic_had_to_add_it_to_the/)
- url: https://i.imgur.com/YMo4iC0.jpg
---

## [3][Strategies for migrating to TypeScript](https://www.reddit.com/r/typescript/comments/g31q57/strategies_for_migrating_to_typescript/)
- url: https://2ality.com/2020/04/migrating-to-typescript.html
---

## [4][Best approach to share a module with types definition between personal projects](https://www.reddit.com/r/typescript/comments/g2yiua/best_approach_to_share_a_module_with_types/)
- url: https://www.reddit.com/r/typescript/comments/g2yiua/best_approach_to_share_a_module_with_types/
---
Hello to everyone :)  
I've started a personal project a few months ago that consisted in a simple Angular app and a few cloud functions on Firebase.

Time has passed and now I'm struggling to keep all my type definitions synced between projects. 

Is there any best practice to share type definitions between projects without publish any package on npm?
## [5][Is there a more idiomatic way to do this in TypeScript?](https://www.reddit.com/r/typescript/comments/g2uro9/is_there_a_more_idiomatic_way_to_do_this_in/)
- url: https://www.reddit.com/r/typescript/comments/g2uro9/is_there_a_more_idiomatic_way_to_do_this_in/
---
Hi y'all, I was recently reading the wikipedia page on [Tagged Unions](https://en.wikipedia.org/wiki/Tagged_union), and under the TS example it showed this block of code as an example:

    type Tree = {
        type:
        | ["node", {
            left: Tree,
            right: Tree
        }]
        | ["leaf", {
            value: string
        }]
    }

    function visit(t: Tree) {
        switch (t.type[0]) {
            case "leaf": {
                const s = t.type[1]
                console.log(s.value)
                break
            }
            case "node": {
                const s = t.type[1]
                visit(s.left)
                visit(s.right)
                break
            }
            default:
                assertUnreachable(t.type[0])
        }
    }
Is this really the idimotic to define/visit a tree or is there a better way? It looks especially clunky especially since TS has the ability to differentiate types and do exhaustivness checking.
## [6][Introducing List Comprehension for TypeScript](https://www.reddit.com/r/typescript/comments/g2mnqp/introducing_list_comprehension_for_typescript/)
- url: https://medium.com/@wim.jongeneel1/introducing-list-comprehension-for-typescript-6204d9b1003e?source=friends_link&amp;sk=09d35f8af4313c7dafde2c8131391a9d
---

## [7][Why is redux's combine reducer complaining about method overload when using typescript and react-redux-firebase?](https://www.reddit.com/r/typescript/comments/g2yr2q/why_is_reduxs_combine_reducer_complaining_about/)
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
## [8][Made this earlier today and thought it was a pretty cool pattern: Implicitly type the second item of a tuple (action payload) based off the first item of a tuple (action type)](https://www.reddit.com/r/typescript/comments/g2sgtj/made_this_earlier_today_and_thought_it_was_a/)
- url: https://gist.github.com/tvler/ccbb3136b0d8c6c48ca53d98382dcb05
---

## [9][How to override node module type declaration](https://www.reddit.com/r/typescript/comments/g2gr9k/how_to_override_node_module_type_declaration/)
- url: https://www.reddit.com/r/typescript/comments/g2gr9k/how_to_override_node_module_type_declaration/
---
Trying to override express.js Request.query type. Have this setup where I can add a type for new attribute, but cant override existing one.

    declare namespace Express { 
        export interface Request {         
            attr1: string, // this is working         
            query: any      // this is not working, query is still type of Query 
        } 
    }

ts.config

    "typeRoots" : [
         "./src/types",
         "node_modules/@types"        
    ]

Am I doing something wrong or is there a different approach?
## [10][A generic middleware pattern in Typescript](https://www.reddit.com/r/typescript/comments/g1wzl6/a_generic_middleware_pattern_in_typescript/)
- url: https://evertpot.com/generic-middleware/
---

## [11][GitHub Action: Typescript compile failsafe](https://www.reddit.com/r/typescript/comments/g26o6y/github_action_typescript_compile_failsafe/)
- url: https://www.reddit.com/r/typescript/comments/g26o6y/github_action_typescript_compile_failsafe/
---
I got annoyed that I kept forgetting to compile my TS before pushing. So I made a GitHub action to compile it for me and make another commit if I ever forget to compile it. I'm posting a link to the action here if you guys ever need that failsafe in the background for anything.

https://github.com/marketplace/actions/typescript-tsc-build-push
