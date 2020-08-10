# typescript
## [1][Who's hiring Typescript developers - August](https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/)
- url: https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/
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
## [2][TypeScript/Express/React/Webpack starter project?](https://www.reddit.com/r/typescript/comments/i74q3u/typescriptexpressreactwebpack_starter_project/)
- url: https://www.reddit.com/r/typescript/comments/i74q3u/typescriptexpressreactwebpack_starter_project/
---
I'm trying to make a TypeScript/Express/React/Webpack project and there's a few bits that I can't get plumbed right. Specifically, I want to have a "good" dev experience without a ton of recompiling. 

As such, I want to run the project when developing using both:

1. Nodemon+TS-Node
2. Webpack Hot Module Replacement

I can't seem to find a combination of configurations that will allow me to get both of those working at the same time. Does anyone have any examples of this setup working well?

Thanks!
## [3][Game of Life in TypeScript - text/video tutorial](https://www.reddit.com/r/typescript/comments/i6mtcf/game_of_life_in_typescript_textvideo_tutorial/)
- url: https://medium.com/hypersphere-codes/conways-game-of-life-in-typescript-a955aec3bd49
---

## [4][MikroOrm v4 /w GraphQL example project](https://www.reddit.com/r/typescript/comments/i6kwnr/mikroorm_v4_w_graphql_example_project/)
- url: https://github.com/driescroons/mikro-orm-graphql-example
---

## [5][Typescript without a framework for frontend development](https://www.reddit.com/r/typescript/comments/i6m5iv/typescript_without_a_framework_for_frontend/)
- url: https://www.reddit.com/r/typescript/comments/i6m5iv/typescript_without_a_framework_for_frontend/
---
I want to use typescript in the front end without the use of a framework. Basically, just compiling it down to es3 javascript. In the tsconfig.json file, I've sent module to none but that's not exactly what I want. 

I do not want the import and export statements to appear in the javascript files after compiling and module none does that; however, in development, I want to be able to use import and export for better IntelliSense, using types, interfaces, and etc. 

How do you accomplish this?
## [6][When using @babel/preset-typescript, how do I prevent compilation unless type checks pass?](https://www.reddit.com/r/typescript/comments/i6no4g/when_using_babelpresettypescript_how_do_i_prevent/)
- url: https://www.reddit.com/r/typescript/comments/i6no4g/when_using_babelpresettypescript_how_do_i_prevent/
---
I've recently switched my Webpack config to using `@babel/preset-typescript` instead of `ts-loader` for easier integration with other tools (like React hot reloading). I know that Babel just strips the types and does not do any type checking by itself.

However, I do not want it to compile anything until my types are correct, the main reasons being:

- Saving a file with wrong types often indicates actual erors in the logic of the app, and hot-loading such a broken file can break my components, so I need to manually refresh the page.
- Sometimes IDEs just bug out and stop showing errors where they should be. I want to at least see them in the terminal, preferably the same one that runs my dev server.
- Production build should only be successful if all types are correct.

`ts-loader` had the correct behavior out of the box. How do I set up a project with Webpack and `@babel/preset-typescript` to behave in the same way?

Thanks!
## [7][defining a method in a different file than the class](https://www.reddit.com/r/typescript/comments/i6v0w4/defining_a_method_in_a_different_file_than_the/)
- url: https://www.reddit.com/r/typescript/comments/i6v0w4/defining_a_method_in_a_different_file_than_the/
---
i.e. keeping the class aware of the method, and the method aware that `this` refers to that class

is it possible without resorting to ugly hacks and/or completely breaking IDE intellisense?

method1.ts

    export default function method1 () : void
    {
        // somehow make this function aware that it will be a method of SomeClass
    
        this.propStr += 'o'
        this.method2()
    }

method2.ts

    export default function method2 () : void
    {
        // somehow make this function aware that it will be a method of SomeClass
    
        this.propNum += 1
    }

SomeClass.ts

    import method1 from './method1'
    import method2 from './method2'
    
    export default class SomeClass
    {
        private propStr:string = 'foo'
        private propNum:number = 123
    
        public method1 = method1.bind(this)
        public method2 = method2.bind(this)
    }
## [8][Help: Redux reducer object pattern / better options?](https://www.reddit.com/r/typescript/comments/i6d7na/help_redux_reducer_object_pattern_better_options/)
- url: https://www.reddit.com/r/typescript/comments/i6d7na/help_redux_reducer_object_pattern_better_options/
---
I've been playing around with using an object, hashed by action types, that maps to a state reducer. I've done this with JS in the past and the TypeScript(ing) is a bit complicated for my liking.

Here's my sample code - 

    const INCREMENT = 'increment'
    type IncrementAction = { type: typeof INCREMENT; amount: number }
    export const increment = (amount: number): IncrementAction =&gt; ({ type: INCREMENT, amount })
    
    const DECREMENT = 'decrement'
    type DecrementAction = { type: typeof DECREMENT; amount: number }
    export const decrement = (amount: number): DecrementAction =&gt; ({ type: DECREMENT, amount })
    
    type StoreState = { readonly amount: number }
    type StoreActions = IncrementAction | DecrementAction
    
    type StoreReducer&lt;S, A extends StoreActions&gt; = {
      [key in A['type']]?: (state: S, action: A extends StoreActions ? (A['type'] extends key ? A : never) : never) =&gt; S
    }
    
    const reducers: StoreReducer&lt;StoreState, StoreActions&gt; = {
      [INCREMENT]: (state, action) =&gt; ({ amount: state.amount + action.amount }),
      [DECREMENT]: (state, action) =&gt; ({ amount: state.amount - action.amount }),
    }
    
    const initialState: StoreState = { amount: 0 }
    
    export default (state: StoreState = initialState, action: StoreActions) =&gt; {
      const reducer = reducers[action.type]
      if (reducer) return reducer(state, action as never) // sad panda
      return state
    }

So this will "work" - within each reducer function, \`(state, action)\` both parameters get inferred properly based on the key - which is really cool!  

My problems with it are 

* The StoreReducer type makes my brain bleed.  This seems overly complicated and I swear I'm going to look at that later and think "wtf is this thing?"
* in the actual reducer function, typescript seems to get "confused" and the type of the action parameter from the returned function always comes back as \`never\`, so I need to cast the action to never, which.... seems a bit wrong, and is making me a sad panda.

 I'm wondering if anyone else has done this in the past and what your approach was.
## [9][Retrieving the type of the keys for an object with dynamic keys](https://www.reddit.com/r/typescript/comments/i6dx8u/retrieving_the_type_of_the_keys_for_an_object/)
- url: https://www.reddit.com/r/typescript/comments/i6dx8u/retrieving_the_type_of_the_keys_for_an_object/
---
Is it possible to retrieve the type set for the keys in a dynamic object like so:   


    type Treatments = {
      [treatment: string]: 'on' | 'off' | string;
    };
    
    const addTreatment = (prev: Treatments, treatment: ***Treatments Key Type***): Treatments =&gt; {
        return {
          ...prev,
          [treatment]: value,
        }
     };
## [10][Is there any such think as type assertion via is , but via a more than one parameter asserting function ?](https://www.reddit.com/r/typescript/comments/i68ngl/is_there_any_such_think_as_type_assertion_via_is/)
- url: https://www.reddit.com/r/typescript/comments/i68ngl/is_there_any_such_think_as_type_assertion_via_is/
---
Or at least do they plan to add that feature in the future in typescript ?

Here is an example of what I am talking about :

```
type shallowSerializableReference = unknown[] | {[x:string] : unknown};
function shouldDiff(oldVal: unknown, newVal: unknown): oldVal is shallowSerializableReference , newVal is shallowSerializableReference {
	return (
		(isOfInternalClass(oldVal, "Object") &amp;&amp; isOfInternalClass(newVal, "Object")) ||
		(isOfInternalClass(oldVal, "Array") &amp;&amp; isOfInternalClass(newVal, "Array"))
	);
}
```
## [11][Array type of unknown length with first element type1 and the rest of the elements type2 .](https://www.reddit.com/r/typescript/comments/i5wgu2/array_type_of_unknown_length_with_first_element/)
- url: https://www.reddit.com/r/typescript/comments/i5wgu2/array_type_of_unknown_length_with_first_element/
---
How to define that type ?
