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
## [2][Help: Redux reducer object pattern / better options?](https://www.reddit.com/r/typescript/comments/i6d7na/help_redux_reducer_object_pattern_better_options/)
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
## [3][Retrieving the type of the keys for an object with dynamic keys](https://www.reddit.com/r/typescript/comments/i6dx8u/retrieving_the_type_of_the_keys_for_an_object/)
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
## [4][Is there any such think as type assertion via is , but via a more than one parameter asserting function ?](https://www.reddit.com/r/typescript/comments/i68ngl/is_there_any_such_think_as_type_assertion_via_is/)
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
## [5][Array type of unknown length with first element type1 and the rest of the elements type2 .](https://www.reddit.com/r/typescript/comments/i5wgu2/array_type_of_unknown_length_with_first_element/)
- url: https://www.reddit.com/r/typescript/comments/i5wgu2/array_type_of_unknown_length_with_first_element/
---
How to define that type ?
## [6][Other languages with mapped/conditional types?](https://www.reddit.com/r/typescript/comments/i5vmni/other_languages_with_mappedconditional_types/)
- url: https://www.reddit.com/r/typescript/comments/i5vmni/other_languages_with_mappedconditional_types/
---
I’ve been using TS for a while now after only writing JS in the past. I’m currently working polymorphic data clients that have conditionally typed response types that would not be possible without conditional/mapped types.

This made me wonder, are there any other languages that use (structural) conditonal/mapped types? And if not, does anyone know why?

(I’ve glanced at C#, Swift, and Rust, and searched the web so perhaps I’ve glossed over some things)

Edit: I’ve added a more detailed example in the comments to explain what I’m after. Perhaps polymorphic is the wrong term here.
## [7][Did you use typescript first and then JavaScript?](https://www.reddit.com/r/typescript/comments/i5pp1l/did_you_use_typescript_first_and_then_javascript/)
- url: https://www.reddit.com/r/typescript/comments/i5pp1l/did_you_use_typescript_first_and_then_javascript/
---
Curiously I had to do this for a university homework in Angular Ionic, without any JS experience it was quite difficult at the time, but since we were learning Java I liked the "types" aspect, now I'm learning more JS when I have time and I want to dedicate more time to TS again later.
## [8][Type parameter within a function call?](https://www.reddit.com/r/typescript/comments/i5vvwq/type_parameter_within_a_function_call/)
- url: https://www.reddit.com/r/typescript/comments/i5vvwq/type_parameter_within_a_function_call/
---
    import { useForm } from "react-hook-form";
    
    interface IFormInput {
      firstName: String;
      gender: GenderEnum;
    }
    
    export default function App() {
      const { register, handleSubmit } = useForm&lt;IFormInput&gt;();

I'm uncomfortable with this syntax. From the object destructure I can tell `useForm` returns an object that has keys for `register` and `handleSubmit`. So it is unclear what `IFormInput`  is doing here, if anything.

Any clarification is appreciated, thanks.

PS: Also, I vaguely recall reading that standard data typings (other than maybe `Function`)  should not be capitalized. Is the capitalized `String` above probably a minor mistake? [https://react-hook-form.com/get-started#Registerfields](https://react-hook-form.com/get-started#Registerfields)
## [9][Is it possible to access static methods on generic types?](https://www.reddit.com/r/typescript/comments/i5ytvz/is_it_possible_to_access_static_methods_on/)
- url: https://www.reddit.com/r/typescript/comments/i5ytvz/is_it_possible_to_access_static_methods_on/
---
I need to know if something like this is possible:

`class Database {`  
 `static newInstance(): Database {`  
 `return new Database();`  
`}`  
`}`  
`function create&lt;T extends Database&gt;(): T {`  
 `return T.newInstance();`  
`}`
## [10][typecasting clutter](https://www.reddit.com/r/typescript/comments/i5x5s7/typecasting_clutter/)
- url: https://www.reddit.com/r/typescript/comments/i5x5s7/typecasting_clutter/
---
    foo = &lt;string&gt;foo
    bar = bar as string

compiles to:

    foo = foo;
    bar = bar;

isn't the compiler supposed to discard these? am I missing something?
## [11][Managing Types across multiple front-ends and services](https://www.reddit.com/r/typescript/comments/i5rutg/managing_types_across_multiple_frontends_and/)
- url: https://www.reddit.com/r/typescript/comments/i5rutg/managing_types_across_multiple_frontends_and/
---
Hey all,

I’m curious if anyone has any good suggestions on ways to manage types across multiple frontends and services that have overlapping uses of types or use similar types. I’m trying to slowly move our org to typescript and would like a good way to manage types that won’t get super unruly.
