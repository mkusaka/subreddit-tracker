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
## [2][How to create a fully tree shakable icon library in Angular](https://www.reddit.com/r/typescript/comments/fnibx7/how_to_create_a_fully_tree_shakable_icon_library/)
- url: https://medium.com/angular-in-depth/how-to-create-a-fully-tree-shakable-icon-library-in-angular-c5488cf9cd76
---

## [3][What is the common convention for interface names?](https://www.reddit.com/r/typescript/comments/fng5ji/what_is_the_common_convention_for_interface_names/)
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
## [4][Need help with spread operator and Date in typescript](https://www.reddit.com/r/typescript/comments/fnhe74/need_help_with_spread_operator_and_date_in/)
- url: https://www.reddit.com/r/typescript/comments/fnhe74/need_help_with_spread_operator_and_date_in/
---
Consider the below code:
```
export function(date :string) : Date{  
 const dateNumbers : number[] = date.split(/,|-| |:/).map(numText =&gt; Number(numText))  
 console.log(...dateNumbers);  
 return new Date(...dateNumbers)  
}
```
I cannot use spread operator for the new Date in typescript, but I actually can use it in javascript, the error is Expected 0-7 arguments, but got 0 or more.ts(2556)

Please help me with that, I could not figure it out.
## [5][why am able to push this "any" into a typed array?](https://www.reddit.com/r/typescript/comments/fn2s73/why_am_able_to_push_this_any_into_a_typed_array/)
- url: https://www.reddit.com/r/typescript/comments/fn2s73/why_am_able_to_push_this_any_into_a_typed_array/
---
 Good Morning I'm beginning to learn nestJS for a backend and so far liking it... 

I'm trying to wire up a front end React as i go along. this code snippet has no back end, it's just an example.

i'm puzzled by this typeScript issue.

The front end makes a request to the backend and what comes back is unknown, i know that it's a TaskI, because that's what the backend is responding with, but i wanted to emulate the case where the front end typescript can't know that and has to "prove" it out. so i set the data to type unknown... 

I was expecting a TS complaint, but here it allows me to push an any into my TaskI array, do you know why this is?

I shouldn't be able to push an any into a typed array.

More generally, not having worked with type unknown before, this is the concept of using type unknown, to prove to the compiler that a type is indeed a type?

Thanks for your time and consideration.

 [https://codesandbox.io/s/practical-shape-p5rko](https://codesandbox.io/s/practical-shape-p5rko)
## [6][(Open Source) Corona Statistics Dashboard in Angular 9 (PWA) + NodeJS Scrapper](https://www.reddit.com/r/typescript/comments/fn63q5/open_source_corona_statistics_dashboard_in/)
- url: https://www.reddit.com/r/typescript/comments/fn63q5/open_source_corona_statistics_dashboard_in/
---
(Open Source) Corona Statistics Dashboard in Angular 9 (PWA) + NodeJS Scrapper

[https://coronastatistics.live/](https://coronastatistics.live/)

Product Hunt: [https://www.producthunt.com/posts/corona-virus-statistics-and-tracker](https://www.producthunt.com/posts/corona-virus-statistics-and-tracker)

Github: [https://github.com/OssamaRafique/Corona-Statistics-And-Tracker-Dashboard-Angular-9](https://github.com/OssamaRafique/Corona-Statistics-And-Tracker-Dashboard-Angular-9)

Buy me a Coffee [https://ko-fi.com/ossamarafique](https://ko-fi.com/ossamarafique)

Please give me a star and upvote if you like it.
## [7][Confused by generic related compilation error.](https://www.reddit.com/r/typescript/comments/fmgw5y/confused_by_generic_related_compilation_error/)
- url: https://www.reddit.com/r/typescript/comments/fmgw5y/confused_by_generic_related_compilation_error/
---
Basically the code that I have that is leading to the compilation error is similar to this:

```
class Box&lt;T&gt; {

}

class Gen&lt;T extends Box&lt;number&gt;&gt; {
    private action(array: Array&lt;T&gt;): Array&lt;T&gt; {
        return new Array()
    }

    private useAction() {
        let x:Box&lt;number&gt;[] = new Array&lt;Box&lt;number&gt;&gt;();
        // This leads to the compilation error
        this.action(x)
    } 
}
```

I would expect this to work.

I defined `action` to take an array of any type, hence the `array:Array&lt;T&gt;`

In `useAction` I am calling the `action` method.

I am calling it with an array of `Box&lt;number&gt;`

Since I defined it to take any type, I expected passing a variable of type `Box&lt;number&gt;` to `action` should then work. 

But it fails with a compilation error:

&gt; Argument of type 'Box&lt;number&gt;[]' is not assignable to parameter of
&gt; type 'T[]'.   Type 'Box&lt;number&gt;' is not assignable to type 'T'.
&gt;     'Box&lt;number&gt;' is assignable to the constraint of type 'T', but 'T' could be instantiated with a different subtype of constraint
&gt; 'Box&lt;number&gt;'.

I can't decipher the error message. Why is it failing and how do I fix it?

I asked the question on StackOverflow, if you know how to crack this riddle, please feel free to answer the question over there

[https://stackoverflow.com/questions/60789275/how-to-fix-compilation-error-when-trying-to-call-a-generic-typescript-method](https://stackoverflow.com/questions/60789275/how-to-fix-compilation-error-when-trying-to-call-a-generic-typescript-method)
## [8][Proper TypeGraphQL architecture](https://www.reddit.com/r/typescript/comments/fmcymr/proper_typegraphql_architecture/)
- url: https://www.reddit.com/r/typescript/comments/fmcymr/proper_typegraphql_architecture/
---
Trying to figure out the reference setup with TypeGraphQL and my feeling is...

- `@ObjectType` + the decorated class reflect the TS Type + SDL Type
- `@Resolver` + the decorated class reflect the resolver
- `@Service` + the decorated class which is DIed with `typedi` should reflect the data source

Where I am unsure is the last one, `@Service`: 

1. Does it makes sense?
2. Do I put actual db queries (in my case mongo native driver) there or do I make another extra class?
3. Why do I have to decorate also the resolver with `@Service`?
4. How would 3 differ re ActiveRecord and DataMapper pattern?
## [9][Introducing Gretchen: Making Fetch Happen in TypeScript](https://www.reddit.com/r/typescript/comments/fm038x/introducing_gretchen_making_fetch_happen_in/)
- url: https://www.truework.com/blog/engineering/2020-03-04-introducing-gretchen/
---

## [10][PostSass - A Sass x PostCSS cli tool, written in TypeScript (X-Post r/webdev)](https://www.reddit.com/r/typescript/comments/fmdebi/postsass_a_sass_x_postcss_cli_tool_written_in/)
- url: https://github.com/Tanuel/postsass
---

## [11][Seafox: self-hosted javascript parser written in Typescript](https://www.reddit.com/r/typescript/comments/fm02aq/seafox_selfhosted_javascript_parser_written_in/)
- url: https://github.com/KFlash/seafox
---

