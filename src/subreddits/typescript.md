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
## [2][Type-safe lenses](https://www.reddit.com/r/typescript/comments/foo055/typesafe_lenses/)
- url: https://github.com/hoppinger/ts-lenses
---

## [3][MySQL Parser in Typescript](https://www.reddit.com/r/typescript/comments/fojj5j/mysql_parser_in_typescript/)
- url: https://github.com/stevenmiller888/ts-mysql-parser
---

## [4][Redux Toolkit v1.3.0 final: New `createAsyncThunk` and `createEntityAdapter` APIs, Immer 6.0, smaller bundle sizes!](https://www.reddit.com/r/typescript/comments/fo5uvd/redux_toolkit_v130_final_new_createasyncthunk_and/)
- url: https://github.com/reduxjs/redux-toolkit/releases/tag/v1.3.0
---

## [5][Is typescript programmers should know about prototype in JavaScript in deep?](https://www.reddit.com/r/typescript/comments/fok6fv/is_typescript_programmers_should_know_about/)
- url: https://www.reddit.com/r/typescript/comments/fok6fv/is_typescript_programmers_should_know_about/
---
If yes why?
## [6][A rule of thumb for working with null values](https://www.reddit.com/r/typescript/comments/fo5sim/a_rule_of_thumb_for_working_with_null_values/)
- url: https://effectivetypescript.com/2020/03/24/null-values-to-perimeter/
---

## [7][How to type return type of a function based on given argument](https://www.reddit.com/r/typescript/comments/fo2rsk/how_to_type_return_type_of_a_function_based_on/)
- url: https://www.reddit.com/r/typescript/comments/fo2rsk/how_to_type_return_type_of_a_function_based_on/
---
Probably the title is wrong, i couldn't figure out how to explain it. Here is a simplified version of the problem.

    type Foo&lt;T&gt; = { getValue: () =&gt; T; };
    
    function createFoo&lt;T&gt;(value: T): Foo&lt;T&gt; { return { getValue: () =&gt; value }; }
    
    let items = { apple: createFoo(1), orange: createFoo('hello'), };
    
    function combine(items) { // }
    
    let combined = combine({ apple, orange });
    
    /* combined's type should be { apple: number, orange: string } */

Thank you :)
## [8][Ran into a TS issue. Maybe a TSC problem?](https://www.reddit.com/r/typescript/comments/fo7j2s/ran_into_a_ts_issue_maybe_a_tsc_problem/)
- url: https://www.reddit.com/r/typescript/comments/fo7j2s/ran_into_a_ts_issue_maybe_a_tsc_problem/
---
Here's the [code](https://pastebin.com/cQrWSrd3)

Compiler [error](https://pastebin.com/auYcY96J)

Edit: Switched to Markdown. I really need to see if there's a way to set that to default. Also, apologies for the syntax highlighting in Pastebin. Is there a better service to share TS code?
## [9][How to type object with variable key names?](https://www.reddit.com/r/typescript/comments/fnu9ha/how_to_type_object_with_variable_key_names/)
- url: https://www.reddit.com/r/typescript/comments/fnu9ha/how_to_type_object_with_variable_key_names/
---
For practice I'm working on an algorithm that assigns values as (object) key names. The keys are paired to an array.

This is in raw  JS. In TS I don't know how to handle objects with dynamic key names:

    { joe : [{...}, {...}, {...}], sally : [{...}, {...}]  } 

(key names will vary)

Anyone know how to properly type these?
## [10][typedi cheat sheet](https://www.reddit.com/r/typescript/comments/fo1equ/typedi_cheat_sheet/)
- url: https://www.reddit.com/r/typescript/comments/fo1equ/typedi_cheat_sheet/
---
Trying to grok `typedi` and came up with this. Is following right?

1. `@Inject()` injects `Container` into decorated class properties,  
*Hint: **InjectContainer**IntoProp*
2. `@Service()` injects `Container` into decorated classes via constructor params,  
*Hint: **InjectContainer**IntoClassViaConstructorParams*
3. `@Service{'&lt;name&gt;')` injects `Container` into decorated classes and provides that classes via `&lt;name&gt;` via `Container.get('&lt;name&gt;')`,  
*Hint: **InjectContainer**IntoClassViaConstructorParamsAnd**ProvideClass**ViaAlias*

1 is mandatory for classes without a constructor. 2 is mandatory if a constructor with params injects the dependency. 3 is the one which actually act as a *service*.

A more intuitive naming scheme would have been:  
1: `@InjectToProp()`   
2: `@InjectViaConstructor()`  
3: `@ProvideAs('&lt;name&gt;')`
## [11][What is the common convention for interface names?](https://www.reddit.com/r/typescript/comments/fng5ji/what_is_the_common_convention_for_interface_names/)
- url: https://www.reddit.com/r/typescript/comments/fng5ji/what_is_the_common_convention_for_interface_names/
---
Hello all,

I am learning Typescript and I was naming my interfaces happily until I arrived to Mongodb/Mongoose.  


For example, I had before an interface called User. However, when I create my model schema, in mongo without Typescript I would have called this User as well. Therefore, that would create a conflict.  


Is there a common convention about this?  


I see this guy in the link down below uses IUser but I have read also that Typescript community does not like using Iwhatever... and personally I do not like that either.

[https://brianflove.com/2016/10/04/typescript-declaring-mongoose-schema-model/](https://brianflove.com/2016/10/04/typescript-declaring-mongoose-schema-model/)  


Therefore, do you know what is the most common way to naming mongo stuff and interfaces?  


Thank you in advance.
