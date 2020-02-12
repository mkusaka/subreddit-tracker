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
## [2][Typeclasses in Typescript with fp-ts](https://www.reddit.com/r/typescript/comments/f2kokc/typeclasses_in_typescript_with_fpts/)
- url: https://paulgray.net/typeclasses-in-typescript/
---

## [3][PeerTube v2.1 released! (PeerTube is an Open Source &amp; Decentralized YouTube Alternative)](https://www.reddit.com/r/typescript/comments/f2a539/peertube_v21_released_peertube_is_an_open_source/)
- url: https://github.com/Chocobozzz/PeerTube/releases/tag/v2.1.0
---

## [4][Immutable objects in TypeScript](https://www.reddit.com/r/typescript/comments/f2nr76/immutable_objects_in_typescript/)
- url: https://nehalist.io/immutable-objects-in-typescript/
---

## [5][Issue understanding Generics in Typescript](https://www.reddit.com/r/typescript/comments/f2muhz/issue_understanding_generics_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/f2muhz/issue_understanding_generics_in_typescript/
---
I'm trying to understand generics in typescript using the example here, but am having issues:  [https://pastebin.com/Pekq7RC6](https://pastebin.com/Pekq7RC6)   


The generic gets created like this:  const miniMaxSum = &lt;arrType&gt;( ) but when I try to assign it to my variables and parameters inside the function, i get TS errors.

&amp;#x200B;
```
let maxSum : arrType = 0; // Error -&gt; '0' is assignable to the constraint of type 'arrType', but 'arrType' could be instantiated with a different subtype of constraint 'number'.
let minSum : arrType = 0;
```

```
maxSum = maxSum + item; // Type 'number' is not assignable to type 'arrType'. 'number' is assignable to the constraint of type 'arrType', but 'arrType' could be instantiated with a different subtype of constraint 'number'.
```
There is also a similar error why trying to add integers to the maxSum and minSum variables. I cant seem to understand why its happening,
## [6][What's the most advanced thing you know within TypeScript?](https://www.reddit.com/r/typescript/comments/f25ux0/whats_the_most_advanced_thing_you_know_within/)
- url: https://www.reddit.com/r/typescript/comments/f25ux0/whats_the_most_advanced_thing_you_know_within/
---
Hey,

I'm planning to give a presentation to provide insights about TypeScript at work soon and I was wondering what's the most advanced thing you can do with TypeScript?

So this is about TS itself and not an advanced project/app using it. A search on Google or the official docs gave me some examples but I was wondering if someone knows some handy things :) Or maybe some solution for a problem you would have needed a bigger workaround without TypeScript, etc.
## [7][Infer type from argument](https://www.reddit.com/r/typescript/comments/f2dcxh/infer_type_from_argument/)
- url: https://www.reddit.com/r/typescript/comments/f2dcxh/infer_type_from_argument/
---
 Hi,

I love TypeScript, but let's say you're writing plain JS, without any of that pseudo TypeScript with the comments (params i think they're called?)

  
(I'm using VS Code by the way)

let's say i have a simple function ...

`let returnTypeOfArgument = input =&gt; {`  
 `return typeof input;`  
`};`  
`console.log(returnTypeOfArgument("abcdefghij"));`

Is there a way that we can infer type from argument that the function is being called with?

So, here i immediately call the function with a string, can the function body infer input to be type string because that's what it's being called with outside the body, and therefore i can get some basic type checking, without any use of TypeScript features?

Thank you
## [8][Why does my type guard not work?](https://www.reddit.com/r/typescript/comments/f2bk53/why_does_my_type_guard_not_work/)
- url: https://www.reddit.com/r/typescript/comments/f2bk53/why_does_my_type_guard_not_work/
---
EDIT: ignore, stupid error, solved now!

Hi, I am getting the error:
```
Argument of type '{} | "string" | "number" | "bigint" | "boolean" | "symbol" | "undefined" | "object" | "function"' is not assignable to parameter of type 'string'.
  Type '{}' is not assignable to type 'string'.(2345)
```

but TS should know it is a string, I am checking just before with `typeof`.

I want the callback function to see if the value of the key `fileContents` is an object, then it assumes it is JSON and uses `JSON.stringify`, or if it is a string to just write a string.

https://www.typescriptlang.org/play/?ssl=1&amp;ssc=1&amp;pln=36&amp;pc=1#code/LAKAxg9gdgzgLgAjgTwA4FMDKYBOBLVOAFXQFtUAbAQznQC4EBvUBBAMzwvQDkrT6E8fFADmAbhbtO6AMLRaUODAaMAvggA+guMPGhVAbQC6CALwIDk1sxCs7Urr34MA5By4BGFwBord97Ly6IrKTH72ACboMLgEcHjQrugAHnyU6D7hCKp+qr62dowOPHwCbtIATD7FcorBSkmp5FwIVdmSRhIgoKAA9ABU-Sz9ZqZj4xOTU9Mzs3Pzs8MIAOr4tEgQCBF4MADWSwuHR8cnU8O9oFQwyFBg7ACut-HQCADua+hEEAAiO7sAFAEnAIhHhRN4tjQqAxQaIAJQMAAKOAgpB26AAPAA3CB4CIAPjC3RAOWJKAw2HwhBIzRo6AAdGwIDgAKJUMAAC3+tFp61MhJsrDwbAQ3LIlDpjOktQUSlG5hcEAARgArdBgOAuOFE+zvPC0L6-Pb-LI8iW0KWOUr5eysABSmAA8tx6bCRMLkGLeQyAjL6jAIVB7hQKBCKnC-HCuqx1OgKDB0AhhaKzdQLb6giF5Ug0OgICKXG6tTq7HqDT8-l7zT7pMCIanJRm6iEo5JSapW8SgA

Is there a better way to type this, or a way to solve the error? If I use `as string` it compiles and runs.
## [9][Visual Debugging in VS Code](https://www.reddit.com/r/typescript/comments/f1shn6/visual_debugging_in_vs_code/)
- url: https://i.redd.it/pdi1ucn2f4g41.gif
---

## [10][Is there any ORM that has dynamically typed query results?](https://www.reddit.com/r/typescript/comments/f1zky5/is_there_any_orm_that_has_dynamically_typed_query/)
- url: https://www.reddit.com/r/typescript/comments/f1zky5/is_there_any_orm_that_has_dynamically_typed_query/
---
Is there any ORM (or query builder) that provides dynamic typing for the result of a query? For example, if I call `findOne` in TypeORM like this:

```
const user = await connection.getRepository(User).findOne({ select: ['firstName'] })
```

it would be nice if the `user` variable was typed to include only the `firstName` field and not all fields associated with the model. Ditto for eagerly loaded associations. However, that doesn't appear to be the case. Are the any libraries that do in fact do that?
## [11][utility type that equals a specific key within an interface given the type which that key is mapped to](https://www.reddit.com/r/typescript/comments/f23wx8/utility_type_that_equals_a_specific_key_within_an/)
- url: https://www.reddit.com/r/typescript/comments/f23wx8/utility_type_that_equals_a_specific_key_within_an/
---
I'm trying to find or write a utility type which would have the behavior of the hypothetical `ReverseLookup` type in this example:

```
interface User {
  name: string
}

interface Article {
  title: string
}

interface Shapes {
  author: User,
  editor: User,
  article: Article,
}

// these would compile
const a: ReverseLookup&lt;Shapes, User&gt; = 'author'; // ✓
const b: ReverseLookup&lt;Shapes, User&gt; = 'editor'; // ✓

// these would cause an error
const c: ReverseLookup&lt;Shapes, User&gt; = 'article'; // ✗ ts error
const d: ReverseLookup&lt;Shapes, Article&gt; = 'editor'; // ✗ ts error
```

Where `ReverseLookup` equals a specific key within an interface given the type which that key is mapped to. So given `Shapes` and `User`, `ReverseLookup` is any key within `Shapes` which is mapped to type `User` .
