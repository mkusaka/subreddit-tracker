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
## [2][(Beta) I've written a guide to architecting Node TypeScript apps - still WIP but has a runnable companion repo and full documentation](https://www.reddit.com/r/typescript/comments/hed665/beta_ive_written_a_guide_to_architecting_node/)
- url: https://jbreckmckye.gitbook.io/node-ts-architecture/
---

## [3][What's coming in TypeScript 4?](https://www.reddit.com/r/typescript/comments/hdvk70/whats_coming_in_typescript_4/)
- url: https://httptoolkit.tech/blog/whats-coming-in-typescript-4
---

## [4][Help reating this type](https://www.reddit.com/r/typescript/comments/heehno/help_reating_this_type/)
- url: https://www.reddit.com/r/typescript/comments/heehno/help_reating_this_type/
---
     translationServiceStub?: Promise&lt;{
            [name: string]: Function;
        }&gt;;

I see this is an optional type. It is a returned promise. The promise contains an object. The object has a key name that is a string. the value of name is a function (in other words name is a method).

Is that accurate?

If so can someone tell me why it wasn't written more concisely like this:

     translationServiceStub?: Promise&lt;{
            name: Function;
        }&gt;;

Or even:

     translationServiceStub?: Promise&lt;{
            name: () =&gt; void // or whatever the expected return type is
        }&gt;;
## [5][Adding types for a third party imported package](https://www.reddit.com/r/typescript/comments/hea2qv/adding_types_for_a_third_party_imported_package/)
- url: https://www.reddit.com/r/typescript/comments/hea2qv/adding_types_for_a_third_party_imported_package/
---
Edit: Solved... kind of? I was using Checkbox like `&lt;Checkbox /&gt;` without any props, so I tried adding the 'id' prop and it worked, but then I tried removing the 'id' prop again and it still worked. I'm guessing the change caused something that was being cached to update, and whatever it was that was cached was causing the error?

I'm trying to use this package: [https://www.npmjs.com/package/react-simple-checkbox](https://www.npmjs.com/package/react-simple-checkbox)

The issue is it doesn't seem to have been written with TypeScript in mind. I'm brand new to TypeScript but I'm guessing this is a relatively common problem.

The error I got was:

&gt;`Could not find a declaration file for module 'react-simple-checkbox'. '/node_modules/react-simple-checkbox/dist/bundle.js' implicitly has an 'any' type. Try 'npm install @types/react-simple-checkbox' if it exists or add a new declaration (.d.ts) file containing 'declare module 'react-simple-checkbox';'ts(7016)`

I tried to npm install the types, but they don't exist, so the other option is to 'add a new declaration file'. I made a d.ts file, declaring prop types according to the npm page:

    declare module "react-simple-checkbox" {  
      export default function Checkbox(props: {  
        backAnimationDuration?: number;  
        borderThickness?: number;  
        checked?: boolean;  
        className?: string;  
        color?: string | object;  
        delay?: number;  
        id?: string;  
        onChange?: Function;  
        size?: number;  
        tickAnimationDuration?: number;  
        tickSize?: number;  
      }) : JSX.Element;
    }

Now the first error is gone, but I'm getting a new error...

&gt;`JSX element type 'Checkbox' does not have any construct or call signatures.ts(2604)`

I'm pretty lost on this one. Looking up the error I just get a bunch of results where people needed to install the correct `@types` file for whatever package they're using... but that doesn't help me since this package doesn't have a types file.

Could somebody point me in the right direction?
## [6][How to correctly use iterators in classes?](https://www.reddit.com/r/typescript/comments/he0o69/how_to_correctly_use_iterators_in_classes/)
- url: https://www.reddit.com/r/typescript/comments/he0o69/how_to_correctly_use_iterators_in_classes/
---
Hi.  I have a class called `List`: the Singly Linked List data structure.

I wanted to implement an iterator for the class, so that I could use something like a for…of loop.

**TLDR** How do I implement an iterator to a class which can be used by all the objects extentiated from the class?
## [7][PUZZLE: How to infer a type from function arguments](https://www.reddit.com/r/typescript/comments/he0z3h/puzzle_how_to_infer_a_type_from_function_arguments/)
- url: https://www.reddit.com/r/typescript/comments/he0z3h/puzzle_how_to_infer_a_type_from_function_arguments/
---
So I have a situation I can't seem to figure out whether it is even possible.

I am trying to infer the type of a value in an object of type `T extends {}`, given the object `T` and `fieldName` with type `keyof T`. I want to use this value as a parameter to a function declared and returned in this scope. See below:

```
export function makeFormFieldProps&lt;T extends {}&gt;(
  formData: T,
  fieldName: keyof T
) {
  return {
    value: formData[fieldName],
    onChange: (val: &lt;type of formData[fieldName]&gt;) =&gt; { ... }
  };
}
```
Thanks!
## [8][[Crosspost] I End Up Writing JavaScript Code Anytime I Try To Write TypeScript Code. Is TypeScript Necessary?](https://www.reddit.com/r/typescript/comments/hdyvq9/crosspost_i_end_up_writing_javascript_code/)
- url: https://www.reddit.com/r/typescript/comments/hdyvq9/crosspost_i_end_up_writing_javascript_code/
---
I have 6 years of Vanilla JavaScript experience. I learned TypeScript a few months back and have basic to intermediate knowledge.

Since Deno was released with native TypeScript support I decided to force myself to always write TypeScript no matter the environment (Deno, Node.js, Front-end) when working on projects.

The reason for this is that it seems the JS ecosystem is heading towards more and more TypeScript adoption since most projects are being rewritten from Vanilla JavaScript to TypeScript these days.

I realized that the code I write always ends up looking like vanilla JavaScript anyway with just a few data types added here and there.

I find it cumbersome to always create data types like **interfaces** all the time so I just use a regular **object literal** most of the time. Also the code becomes bloated with too many unnecessary data types so I tend not to use types for basic variables.

Setting up a proper linter like ESLint helps me find bugs most of the time when writing just vanilla JavaScript,

My question is why is there so much hype about TypeScript claiming it helps prevent bugs and it's ideal for large JavaScript projects?
## [9][Implicit super calls](https://www.reddit.com/r/typescript/comments/hdxmj8/implicit_super_calls/)
- url: https://www.reddit.com/r/typescript/comments/hdxmj8/implicit_super_calls/
---
    class Parent {
      public a: string;
      
      constructor(setA: string) {
        this.a = setA;
      }
    }
    
    class Child extends Parent {}
    
    
    const childInstance = new Child("value passed in");
    
    console.log(childInstance.a); // prints

Any risks to relying on implicit super calls?

At first I thought argument ambiguity. But they are positional so it shouldn't matter.
## [10][[Newbie] Are getters syntatic sugar?](https://www.reddit.com/r/typescript/comments/hdmea1/newbie_are_getters_syntatic_sugar/)
- url: https://www.reddit.com/r/typescript/comments/hdmea1/newbie_are_getters_syntatic_sugar/
---
Instead of 

&amp;#x200B;

    constructor(public x: string){
    
    }
    get x () {
       
      return x;
    }

Why not just do 

&amp;#x200B;

    constructor(public x: string){
    
    }
    
    console.log(class.x)
    
    or even
    
    y = "y";
    constructor(public x: string){
    
    }
    console.log(class.y);
## [11][Can I annotate all TypeORM queries with a filter condition?](https://www.reddit.com/r/typescript/comments/hdph92/can_i_annotate_all_typeorm_queries_with_a_filter/)
- url: https://www.reddit.com/r/typescript/comments/hdph92/can_i_annotate_all_typeorm_queries_with_a_filter/
---
I'm trying to create a multi-tenant architecture. For this, I want to create a new EntityManager  
or Repository that automatically annotates all queries with an additional filter for the tenant\_id.

A query like:

    createQueryBuilder() 
        .select("user") 
        .from(User, "user") 
        .where("user.id = :id", { id: "[...]" })

Should be automatically transformed by the Repository or Manager to a query like:

    createQueryBuilder() 
        .select("user") .from(User, "user") 
        .where("user.id = :id", { id: "[...]" }) 
        .andWhere("user.tenantId = :tenantId", { tenantId: "[...]" });

Is this possible with TypeORM?

(The benefit with this approach is that it is more fault-tolerant since it would be impossible to forget filtering for the right tenant)
