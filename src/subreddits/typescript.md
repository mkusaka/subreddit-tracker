# typescript
## [1][Who's hiring Typescript developers - March](https://www.reddit.com/r/typescript/comments/fbll8c/whos_hiring_typescript_developers_march/)
- url: https://www.reddit.com/r/typescript/comments/fbll8c/whos_hiring_typescript_developers_march/
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
## [2][I’m going to give a free and remote talk about problem solving with TypeScript compiler!](https://www.reddit.com/r/typescript/comments/frlquz/im_going_to_give_a_free_and_remote_talk_about/)
- url: https://www.meetup.com/at-wix/events/269726440/
---

## [3][tsPEG - PEG Parser generator for TypeScript - v1.3.0 just released!](https://www.reddit.com/r/typescript/comments/frq0cq/tspeg_peg_parser_generator_for_typescript_v130/)
- url: https://www.reddit.com/r/typescript/comments/frq0cq/tspeg_peg_parser_generator_for_typescript_v130/
---
The latest version of tsPEG has just been released! Check it out at [github.com/EoinDavey/tsPEG](https://github.com/EoinDavey/tsPEG), or see the list of new features [here](https://vey.ie/2020/03/30/tsPEG-v1-3-0.html).
## [4][Exporting Packages](https://www.reddit.com/r/typescript/comments/frlugf/exporting_packages/)
- url: https://www.reddit.com/r/typescript/comments/frlugf/exporting_packages/
---
Hello! I am exporting a package I have created with common types and interfaces between my backend and frontend.

What I am trying to achieve is the ability to import from the package like so:

`import { something } from @quwackers/common/utils/stuff`

Currently I am building to *dist/* however this means I can not access the above, i.e.

`dist/`  
`├── types/`  
`├── common/`  
`└── utils/`

I determined I could build simply to the root and import as indicated above, but I want to know if there is a '*cleaner'* or *'better'* generally accepted way to export my package.
## [5][When compiling, can you specify that you want comments removed from the output JavaScript files but left in for the generated declaration files!](https://www.reddit.com/r/typescript/comments/frj1hn/when_compiling_can_you_specify_that_you_want/)
- url: https://www.reddit.com/r/typescript/comments/frj1hn/when_compiling_can_you_specify_that_you_want/
---

## [6][Why does this question not have a decent answer?](https://www.reddit.com/r/typescript/comments/frgb87/why_does_this_question_not_have_a_decent_answer/)
- url: https://www.reddit.com/r/typescript/comments/frgb87/why_does_this_question_not_have_a_decent_answer/
---
https://stackoverflow.com/questions/46696266/where-can-i-find-documentation-for-typescripts-built-in-types-and-standard-libr

I don't understand how one can overlook the need for standard library reference documentation.
## [7][Generic Evaluation Display](https://www.reddit.com/r/typescript/comments/fr85n8/generic_evaluation_display/)
- url: https://www.reddit.com/r/typescript/comments/fr85n8/generic_evaluation_display/
---
Hi. This might seem a little esoteric but this is my situation.

I have been having fun writing somewhat advanced Generic types.

They work and the types do what I want them to do.

My issue is how typescript displays them. For example when I force an error message I get something like this:  


```
The expected type comes from property 'd' which is declared here on type
'InferObjectShape&lt;{ a: StringType; b: BooleanType; }&gt; &amp; InferObjectShape&lt;{ c: StringType; d: StringType; }&gt;'  
```  


However I know that `'InferObjectShape&lt;{ a: StringType; b: BooleanType; }&gt;` gets evaluated to `{a: string; b: string}`.

Is there anyway to tell typescript to go ahead and evaluate the generics fully?

It would be more useful if the message was:
```
The expected type comes from property 'd' which is declared here on type
'{ a: string; b: boolean; } &amp; { c: string; d: string }'  
```
  
or better yet (but not longer related to generics per se:
```
The expected type comes from property 'd' which is declared here on type
'{ a: string; b: boolean; c: string; d: string }'  
```

It would be much more useful for the user to see the evaluate object rather than Generics that don't mean anything to them. Is there any option for this?

Thanks.
## [8][Question about a Plugin System](https://www.reddit.com/r/typescript/comments/fr5w6w/question_about_a_plugin_system/)
- url: https://www.reddit.com/r/typescript/comments/fr5w6w/question_about_a_plugin_system/
---
Hello everyone,

I was thinking about a runtime plugin system for my library ([https://github.com/matteobruni/tsparticles/](https://github.com/matteobruni/tsparticles/)).

Using the library as sample some features can be external plugins like PolygonMask (that also adds a dependency to pathseg).

My idea is about adding the plugin script in the HTML and it will be added to the main container without any additional line of code.

Somebody know how to do it? No external libraries, I want to keep the project dependency-free (except pathseg, chrome forced me to add it) for now.
## [9][React+Redux+Redux-Thunk action creator right types ???](https://www.reddit.com/r/typescript/comments/frb7yc/reactreduxreduxthunk_action_creator_right_types/)
- url: https://www.reddit.com/r/typescript/comments/frb7yc/reactreduxreduxthunk_action_creator_right_types/
---
Hello community, I would like to ask for a help or advise with `react`, `typescript` application that uses `Redux Thunk` as a middleware.

My application uses [ducks file structure for Redux](https://medium.com/@scbarrus/the-ducks-file-structure-for-redux-d63c41b7035c).

I'm trying to _dispatch_ an `action` which is `Redux-Thunk action creator` which also dispatching an `Redux-Thunk action creator` (Sorry for my English I will try to explain it with code).

### Lets say I have an action creator like that:

```ts
    /** 
     * Not going to work unless I will explicitly change return type to `any` or do `as any` when dispatching this action from another action creator.
     */ 
export function setData(
  value: boolean
): ThunkAction&lt;void, AppState, null, SetDataAction&gt; {
  return (dispatch: Dispatch, getState: () =&gt; AppState): void =&gt; {
      dispatch({
        type: SHOW_ALL_DATA,
        payload: false,
      });
    }
  };
}
```

#### Than later, somewhere in my app I have another **_action_** that `dispatch` the action above, like that:

```ts
export function setSomeExtraData(
  value: boolean
): ThunkAction&lt;void, AppState, null, SetExtraDataAction&gt; {
  return (dispatch: Dispatch, getState: () =&gt; AppState): void =&gt; {
    dispatch(setData(true)); // not working
    // dispatch((setData as any)(true));working
    }
  };
}
```

#### Eror mesage looks like that:

```
No overload matches this call.
  Overload 1 of 3, '(action: RSAAAction&lt;any, unknown, unknown&gt;): Promise&lt;RSAAResultAction&lt;unknown, unknown&gt;&gt;', gave the following error.
    Argument of type 'ThunkAction&lt;Promise&lt;void&gt;, null, CombinedState&lt;{ auth: CombinedState&lt;{ username: string; authToken: string; signInTimestamp: number; }&gt;; currentUser: CombinedState&lt;{ user: User | null; isLoading: boolean; loaded: boolean; }&gt;; ... 12 more ...; inflightRequests: CombinedState&lt;...&gt;; }&gt;, LogsActionTypes&gt; | undefined' is not assignable to parameter of type 'RSAAAction&lt;any, unknown, unknown&gt;'.
      Type 'undefined' is not assignable to type 'RSAAAction&lt;any, unknown, unknown&gt;'.
      ...
      ...
```

So what we got. Ususaly when I'm not dispatching an `thunk-action creator` I don't even need to specify return type like I did here: 

`ThunkAction&lt;void, AppState, null, SetDataAction&gt;`

typescript more smarter than me to pick the right types, but with my case, I'm looking for a right types so Typescript wouldn't complain.

As a reference I saw this [StackOverflow post](https://stackoverflow.com/questions/43013204/how-to-dispatch-an-action-or-a-thunkaction-in-typescript-with-redux-thunk), but none of the solution is working for me. Please let me know if right type is exist for my case or `as any` is only one option here (except refactoring the code to avoid dispatching `redux-thunk action creator`)

Thank you!
## [10][Is Typescript + React production ready ?](https://www.reddit.com/r/typescript/comments/frajar/is_typescript_react_production_ready/)
- url: https://www.reddit.com/r/typescript/comments/frajar/is_typescript_react_production_ready/
---
i am a backend developer learning Typescript and Angular. Is the Typescript + react production ready and has anyone used it in production ? I am asking this because most of the tutorial i see for react is using js.
## [11][Announcing TypeScript 3.9 Beta](https://www.reddit.com/r/typescript/comments/fqnxb2/announcing_typescript_39_beta/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-3-9-beta/
---

