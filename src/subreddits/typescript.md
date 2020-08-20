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
## [2][Learning typescript](https://www.reddit.com/r/typescript/comments/id8ju3/learning_typescript/)
- url: https://www.reddit.com/r/typescript/comments/id8ju3/learning_typescript/
---
Hi,

I want to learn typescript, not for a server-side or a client-side application but for writing standalone applications. I am not very familiar with javascript, I mainly write in Python. What topics should I learn before typescript? From my understanding, I need at least to learn the basics of node.js and javascript. Is there a good tutorial or course that covers all these topics? 

Thanks!
## [3][How to correct relative paths to other file types in same directory when compiling TS?](https://www.reddit.com/r/typescript/comments/icz5ig/how_to_correct_relative_paths_to_other_file_types/)
- url: https://www.reddit.com/r/typescript/comments/icz5ig/how_to_correct_relative_paths_to_other_file_types/
---
So let's say we have ts and css file in the "src" directory. In ts file we have import './style.css'. After compiling TS to "lib" directory, that path is no longer valid for compiled file, because css file is not compiled. How would you make it change to '../src/style.css'?
## [4][For a function that gets as an argument a single object with optional properties that have initial values , how to type the argument ?](https://www.reddit.com/r/typescript/comments/idacac/for_a_function_that_gets_as_an_argument_a_single/)
- url: https://www.reddit.com/r/typescript/comments/idacac/for_a_function_that_gets_as_an_argument_a_single/
---
For example :

    function foo({a = 1 , b} = {}) {
        console.log(a,b);
    }
    foo({});//1 undefined
    foo();//1 undefined
    foo({a : 2});//2 undefined
    foo({b : 2});//1 2
    foo({a : 3 ,b : 2});//3 2

How do I type the argument of that function in typescript ?
## [5][TypeScript dialect with Extension Methods?](https://www.reddit.com/r/typescript/comments/id67j7/typescript_dialect_with_extension_methods/)
- url: https://www.reddit.com/r/typescript/comments/id67j7/typescript_dialect_with_extension_methods/
---
The philosophy of TypeScript is to not do anything that's not JavaScript Standard, so the official compiler probably never support it.

TypeScript became quite a powerful language and pretty much the only thing I miss is the Extension Methods.

I wonder if there are any modification of TypeScript compiler or dialect that would do that?

P.S.

Just realised that I'm missing one more thing - proper function overloading.
## [6][Specify generic type must have at least one property](https://www.reddit.com/r/typescript/comments/icr0cr/specify_generic_type_must_have_at_least_one/)
- url: https://www.reddit.com/r/typescript/comments/icr0cr/specify_generic_type_must_have_at_least_one/
---
I have a type like:

export type Response&lt;T&gt; = { valid: boolean } &amp; T;I am trying to say that the resulting inferred type must at least have \`valid: boolean\` as a property. I believe this is wrong though because when I am trying to infer promise based Response types, then I am getting an error that my return type does not include \`then\`, \`catch\`, \`finally\` (like a promise).

I want typescript to know that valid will always be there, but whatever the type the dev passes in as T, then my response will also have all of those properties.

is there a better way to do what I am trying to do? Thank you
## [7][Union and Intersection Types are confusing because they describe the result, not the operation it self](https://www.reddit.com/r/typescript/comments/iculfz/union_and_intersection_types_are_confusing/)
- url: https://www.reddit.com/r/typescript/comments/iculfz/union_and_intersection_types_are_confusing/
---
I think it's counter-intuitive to use Set Theory or Venn Diagrams to understand/explain Union and Intersection operation in TS, because Union and Intersection in TS actually describe the result of the operation, not the operation itself.

* **Union** `|` **in TS === XOR**  
Union Operator `|`  in TS is actually doing *XOR / Either Or* **on types**, but the *result* is equivalent to doing *Set Intersection* **on type properties**
* **Intersection** `&amp;` **in TS === Combining**  
Intersection Operator `&amp;` in TS actually doing *Combining* **on types**, but the *result* is equivalent to doing *Set Union* **on type properties**

### Example:
Let `item` a value of arbitrary type.\
Let `Type_A = {foo: number, bar: int}`, `Type_B = {bar: int, baz: string}`

* **the Intersection Type `Type_A &amp; Type_B`:** \
Defining `item` of `Type_A &amp; Type_B` means `item` is the **combination of** `Type_A` and `Type_B`, which is actually `{foo: number, bar: int, baz: string}`. Then it is safe to allow `item.foo`, `item.bar`, and `item.baz`.\
Hence, we have
`Type_A &amp; Type_B = {foo: number, bar: int, baz: string}`

* **the Union Type `Type_A | Type_B`:**\
Defining `item` of `Type_A | Type_B` means `item` is **either** `Type_A` **or** `Type_B`. Then it is unsafe to allow `item.foo`, because that's an error if the actual `item` is of `Type_B`. Similarly, accessing `item.baz` is also unsafe. \
Hence, we have
`Type_A | Type_B = {bar: int}`

I have also put this answer in SO: [https://stackoverflow.com/a/63485446/11554998](https://stackoverflow.com/a/63485446/11554998)

But I'm not sure if it really makes sense to other people. Feel free to share your thoughts. I hope I understand the problem correctly and my post can be helpful. :D
## [8][Which ORM should I use?](https://www.reddit.com/r/typescript/comments/icbey2/which_orm_should_i_use/)
- url: https://www.reddit.com/r/typescript/comments/icbey2/which_orm_should_i_use/
---
Which ORM should I use to handle a SQL database in a TypeScript project?
## [9][Dynamic type systems are not inherently more open](https://www.reddit.com/r/typescript/comments/icd0sn/dynamic_type_systems_are_not_inherently_more_open/)
- url: https://lexi-lambda.github.io/blog/2020/01/19/no-dynamic-type-systems-are-not-inherently-more-open/
---

## [10][Pascal interpreter written 100% in Typescript](https://www.reddit.com/r/typescript/comments/icoobk/pascal_interpreter_written_100_in_typescript/)
- url: https://github.com/komninoschat/psi
---

## [11][How to enforce nested objects have matching keys?](https://www.reddit.com/r/typescript/comments/iccp8r/how_to_enforce_nested_objects_have_matching_keys/)
- url: https://www.reddit.com/r/typescript/comments/iccp8r/how_to_enforce_nested_objects_have_matching_keys/
---
I am trying to create a localization object for my React Native app. This is the structure:

    const strings = {
      ...
    
      en: {
        string_1: "foo",
        string_2: "bar",
        ...
      },
    
      fr: { ... },
    
      ...
    }

I want to enforce through TypeScript that each locale matches the other in terms of keys. For example, if `"en"` has key `"string_1"`, then every other locale should also have key `"string_1"`. Or if `"fr"` has key `"string_3"`, then every other locale should also have key `"string_3"`.

My reasoning is that I don't want to add support for a language that does not have all the keys used in the app, so it would be nice to know if/what is missing.

Is this possible to do? And is it possible to do for an arbitrary amount of languages and keys without having to manually create/update an interface?
