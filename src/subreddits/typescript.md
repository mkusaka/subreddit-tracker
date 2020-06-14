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
## [2][Rosebox 0.2.3](https://www.reddit.com/r/typescript/comments/h8pev6/rosebox_023/)
- url: https://www.reddit.com/r/typescript/comments/h8pev6/rosebox_023/
---
[Rosebox](https://www.rosebox.dev/) is an active project with the ultimate goal to build an independent and complete JS/TS styling framework with an emphasis on writing declarative, reliable, and conflict-free style-code.
## [3][Why does this work?](https://www.reddit.com/r/typescript/comments/h89nb2/why_does_this_work/)
- url: https://www.reddit.com/r/typescript/comments/h89nb2/why_does_this_work/
---
I'm having trouble understanding this odd behavior...

```
interface Foo {
    n? : number;
    s: string;
}

const foo = { n : 1 } as Foo;  // OK

```

If  `s` here isn't optional, then why does this work? Isn't this an error? Shouldn't we be assigning `s` to something?
## [4][How can I get ts specific errors , for all my files , and log them in the terminal without building ?](https://www.reddit.com/r/typescript/comments/h8dfyq/how_can_i_get_ts_specific_errors_for_all_my_files/)
- url: https://www.reddit.com/r/typescript/comments/h8dfyq/how_can_i_get_ts_specific_errors_for_all_my_files/
---
I am taking about the errors which are provided when you tsc . Nut I want to get them without building .
## [5][webpack and conflicting types with identical versions in imported library.](https://www.reddit.com/r/typescript/comments/h8dqld/webpack_and_conflicting_types_with_identical/)
- url: https://www.reddit.com/r/typescript/comments/h8dqld/webpack_and_conflicting_types_with_identical/
---
I'm really stuck on this issue... 

If I have a react components library that I want to isolate I can't just setup a devDependency on @types/react because when I import it  I get a conflict saying that types were found in two locations.  

This is super silly because they're the same types and same version.

I *could* setup peerDependencies, I think, but then I can't have unit tests in the main code or use that as a module by itself.

It ends up causing a LOT of hassle and now I'm basically dead in the water until I completely refactor my  components/modules into separate npm projects.

Is there a workaround here?
## [6][Prevent tsc from renaming variables that are not exposed .](https://www.reddit.com/r/typescript/comments/h89fc9/prevent_tsc_from_renaming_variables_that_are_not/)
- url: https://www.reddit.com/r/typescript/comments/h89fc9/prevent_tsc_from_renaming_variables_that_are_not/
---
I have the following two functions : `getAllTextNodesDeep` , `recurseDomTreeNodes` , in my own personal browser-specific typescript library .

In my library there is an entry point file that exports both of them .

Each of the functions has its own file that default exports it .

The first function uses the second function.

When I `tsc` my library the compiler changes the internal variable names . So instead of `recurseDomTreeNodes` there is `recurseDomTreeNodes_1` all over the place inside the `getAllTextNodesDeep` function definition .

That cause no issues at all except in the following case :

I want to add some functions in the window object of the puppeteer `page` instance . When I do that , and then invoke `getAllTextNodesDeep` , the function looks for `recurseDomTreeNodes_1` and not `recurseDomTreeNodes` , which makes my function throw error.

Do you know if there is any way to make tsc not change internal names?

Or am I writing bad practice code and I should go for a dependency injection or something else ?
## [7][Positional Lists in TypeScript](https://www.reddit.com/r/typescript/comments/h7tms0/positional_lists_in_typescript/)
- url: https://medium.com/@wim.jongeneel1/positional-lists-in-typescript-e6fa939a8a43?source=friends_link&amp;sk=ade329e32c20447c8679769a7ad0e7d3
---

## [8][How do you manage complex state models? (NgRx etc)](https://www.reddit.com/r/typescript/comments/h7u1z0/how_do_you_manage_complex_state_models_ngrx_etc/)
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
## [9][Arbitrary amount of generics in function call?](https://www.reddit.com/r/typescript/comments/h7ivji/arbitrary_amount_of_generics_in_function_call/)
- url: https://www.reddit.com/r/typescript/comments/h7ivji/arbitrary_amount_of_generics_in_function_call/
---
The mysql package allows you to make an SQL query with an arbitrary amount of statements. In my case, I've got a stored procedure with an OUT parameter:

    db.query('call getSalaries(1, 2, @average); select @average)'); // two statements

This would return an array where the first element is a Salary[] and the second element is an array of numbers (the average - though in this particular case, that array would contain a single number). In general, depending on the statements in the query and how many there are, I'm going to have an array of elements of different types, with corresponding length.

I would like to be able to call query in different ways. In general, if I call db.query&lt;A, B, C, ...etc&gt;(myQuery), then I know I'm going to get an array where the first element is A[], the second B[], the third C[], and so on. How would I go about modeling that?

I hope the question is clear. Thanks!
## [10][How to exclude folders when include follows a glob pattern?](https://www.reddit.com/r/typescript/comments/h7o0fg/how_to_exclude_folders_when_include_follows_a/)
- url: https://www.reddit.com/r/typescript/comments/h7o0fg/how_to_exclude_folders_when_include_follows_a/
---
I think the `**` means "include everything". Below I was hoping excluding `__tests__` would take precedence over this broad match in include but it did not. How would you modify tsconfig to include everything in `src` except for folders and files inside `__tests__` ?
## [11][What is a valid use case for Interfaces with indexers in typescript?](https://www.reddit.com/r/typescript/comments/h7k8y7/what_is_a_valid_use_case_for_interfaces_with/)
- url: https://www.reddit.com/r/typescript/comments/h7k8y7/what_is_a_valid_use_case_for_interfaces_with/
---
I could not quite come up with a valid use case to justify why typescript supports interface with indexers. plain javascript provides you with indexing objects then why one would need to define an indexer in an interface.?
