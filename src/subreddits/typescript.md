# typescript
## [1][Who's hiring Typescript developers - June](https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/)
- url: https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/
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
## [2][Positional Lists in TypeScript](https://www.reddit.com/r/typescript/comments/h7tms0/positional_lists_in_typescript/)
- url: https://medium.com/@wim.jongeneel1/positional-lists-in-typescript-e6fa939a8a43?source=friends_link&amp;sk=ade329e32c20447c8679769a7ad0e7d3
---

## [3][How do you manage complex state models? (NgRx etc)](https://www.reddit.com/r/typescript/comments/h7u1z0/how_do_you_manage_complex_state_models_ngrx_etc/)
- url: https://www.reddit.com/r/typescript/comments/h7u1z0/how_do_you_manage_complex_state_models_ngrx_etc/
---
When working with rather complex state models it becomes quite messy using the spread operator to partially update the state only. Also you lose with the spread operator the actual class instance.

For example my state model (just an example, not my real state)


```typescript
class Content {
    articles: Article[];
    //..
    getArticle(id: string) { return; }
}

class AppState {
    loggedIn: boolean;
    user: User;
    content: Content;
    uiState: UIState;
}
```

When doing action-&gt;state reduction you would have a reducer like this

```typescript
on(action, (state) =&gt; {
    return {...state, loggedIn: true, content: {...state.content, articles: articles};
}
```
where you basically lose class instances. That gets also more and more unsuitable the more my state model grows in complexity.
I wonder if there are libraries out there that help with this uses cases. Do you know any, or do you limit yourself to work with plain objects and a shallow store all the time? NgRx for examples write in their docs it's recommended to keep the store state as shallow as possible, and now I wonder if this is because of exactly this reason or if there is another reason involved here. Seems like an arbitrary limitation because they don't provide a library to easily patch frozen objects.

I ask because I created a solution for that which allows to patch frozen objects (by returning a new patched one which contains only updates where actual change happened), without using that dirty spread operator and wonder if there are other pitfalls.

My current code (which works for frozen states and returns a new one instead of modifying the original):

```typescript
on(action, (state) =&gt; {
    return applyPatch(state, state =&gt; {
        state.loggedIn = true;
        state.content.articles = articles;
    });
}
```
## [4][Arbitrary amount of generics in function call?](https://www.reddit.com/r/typescript/comments/h7ivji/arbitrary_amount_of_generics_in_function_call/)
- url: https://www.reddit.com/r/typescript/comments/h7ivji/arbitrary_amount_of_generics_in_function_call/
---
The mysql package allows you to make an SQL query with an arbitrary amount of statements. In my case, I've got a stored procedure with an OUT parameter:

    db.query('call getSalaries(1, 2, @average); select @average)'); // two statements

This would return an array where the first element is a Salary[] and the second element is an array of numbers (the average - though in this particular case, that array would contain a single number). In general, depending on the statements in the query and how many there are, I'm going to have an array of elements of different types, with corresponding length.

I would like to be able to call query in different ways. In general, if I call db.query&lt;A, B, C, ...etc&gt;(myQuery), then I know I'm going to get an array where the first element is A[], the second B[], the third C[], and so on. How would I go about modeling that?

I hope the question is clear. Thanks!
## [5][How to exclude folders when include follows a glob pattern?](https://www.reddit.com/r/typescript/comments/h7o0fg/how_to_exclude_folders_when_include_follows_a/)
- url: https://www.reddit.com/r/typescript/comments/h7o0fg/how_to_exclude_folders_when_include_follows_a/
---
I think the `**` means "include everything". Below I was hoping excluding `__tests__` would take precedence over this broad match in include but it did not. How would you modify tsconfig to include everything in `src` except for folders and files inside `__tests__` ?
## [6][What is a valid use case for Interfaces with indexers in typescript?](https://www.reddit.com/r/typescript/comments/h7k8y7/what_is_a_valid_use_case_for_interfaces_with/)
- url: https://www.reddit.com/r/typescript/comments/h7k8y7/what_is_a_valid_use_case_for_interfaces_with/
---
I could not quite come up with a valid use case to justify why typescript supports interface with indexers. plain javascript provides you with indexing objects then why one would need to define an indexer in an interface.?
## [7][how can I keep absolute path rule to transcompiled js file?](https://www.reddit.com/r/typescript/comments/h7megx/how_can_i_keep_absolute_path_rule_to/)
- url: https://www.reddit.com/r/typescript/comments/h7megx/how_can_i_keep_absolute_path_rule_to/
---
&amp;#x200B;

[tsconfig.json](https://preview.redd.it/d1om8x8mvh451.png?width=782&amp;format=png&amp;auto=webp&amp;s=7102291aa13a282ecf40a1a0947c471def50bb49)

I've set the paths and baseUrl!

&amp;#x200B;

&amp;#x200B;

[import crawler](https://preview.redd.it/pv0kv4drvh451.png?width=728&amp;format=png&amp;auto=webp&amp;s=5a5ec00ca6006e44d33dedca08a2171028f22415)

it works well. but problem is when I run the code..!

&amp;#x200B;

&amp;#x200B;

[compiled one](https://preview.redd.it/y15psx5uvh451.png?width=882&amp;format=png&amp;auto=webp&amp;s=1df94a8c01447172032a1ac37acd8ad7377b73c5)

it compiled well as well.

&amp;#x200B;

[error](https://preview.redd.it/h3goftwwvh451.png?width=1168&amp;format=png&amp;auto=webp&amp;s=86e09e854c055ed9196a15f1f3901e73b6ca15d7)

this is what I got.

&amp;#x200B;

I think the absolute rule didn't apply when I run the code. 

&amp;#x200B;

any helps..?
## [8][Node.js, Dependency Injection, Layered Architecture, and TDD: A Practical Example Part 1](https://www.reddit.com/r/typescript/comments/h13hr7/nodejs_dependency_injection_layered_architecture/)
- url: https://carlosgonzalez.dev/posts/node-js-di-layered-architecture-and-tdd-a-practical-example-part-1/
---

## [9][How to warn on number to boolean coercion?](https://www.reddit.com/r/typescript/comments/h7aw3i/how_to_warn_on_number_to_boolean_coercion/)
- url: https://www.reddit.com/r/typescript/comments/h7aw3i/how_to_warn_on_number_to_boolean_coercion/
---
Hi, I was wondering if it's possible to get Typescript to warn whenever a possible number is coerced into a boolean. For example:

```
const x: number | null = 5;

if(x) {} // warn

const y = !x //warn
```

The reason for this is that I would like to avoid numbers potentially becoming falsey due to 0.
## [10][How to create a Conditional Type/Typeguard for this?](https://www.reddit.com/r/typescript/comments/h7dob0/how_to_create_a_conditional_typetypeguard_for_this/)
- url: https://www.reddit.com/r/typescript/comments/h7dob0/how_to_create_a_conditional_typetypeguard_for_this/
---
So I've got a data structure (`SingleValue`) that I want to be able to handle in 2 different ways. The first is when only one or no instances of it is needed, in which case the variable holding the value should be of type `SingleValue | null`. The other case is when many are needed, in which case the variable holding the value should be of type `SingleValue[]`.

How can I set up a single variable and a typeguard function for a type that is conditionally one of those two types based on a generic type argument.

This is what I have at the moment but it's not quite right:

```
type SingleValue = { /* ... */ };  
type Value&lt;Multiple extends boolean&gt; = Multiple extends true  
  ? SingleValue[]  
  : SingleValue | null;  

function isMultipleValue&lt;Multiple extends boolean&gt;(  
 value: Value&lt;Multiple&gt;  
): value is Value&lt;true&gt; {  
 return Array.isArray(value);  
}
```

The error I get is:

```
A type predicate's type must be assignable to its parameter's type.
  Type 'readonly SingleValue[]' is not assignable to type 'Value&lt;Multiple&gt;'.
```

(this error happens on the 3rd to last line)

Note: I cannot switch to just using an array for the one/none case due to how other parts of my project are set up and the implementation of 3rd party libraries. For the same reason `Multiple` but be a generic type argument.
## [11][Duplication of code for ORMs](https://www.reddit.com/r/typescript/comments/h15lwb/duplication_of_code_for_orms/)
- url: https://www.reddit.com/r/typescript/comments/h15lwb/duplication_of_code_for_orms/
---
First of all, I am a newcomer to Typescript, and my previous projects involved C# and WPF.

On the C# side of things, I have always used Entity Framework as ORM, which was a mighty powerful ORM. It handled code-first as well as db-first approaches quite easily.

As far as I can tell, there is nothing in Typescript that handles db-first approach. But the bigger problem for me is duplication of code when using ORMs. Entity Framework could easily infer the database schema from the classes only. In Typescript, if I want to define a schema, I will have to do it twice, once as a plain old TS class, and secondly for the ORM. An example of such a code is [here](https://sequelize.org/master/manual/typescript.html).

I have found typegoose to do something akin to Entity Framework when it comes to ORM. TS-Mongoose looked promising until I ran into the issue [here](https://www.reddit.com/r/typescript/comments/h0e3ff/help_with_tsmongoose/).

Has anyone found a less tedious solution for Postgresql too?

&amp;#x200B;

EDIT: Found it. It's TypeORM.
