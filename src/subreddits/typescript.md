# typescript
## [1][Who's hiring Typescript developers - May](https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/)
- url: https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/
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
## [2][Phaser Integration with Angular and Electron](https://www.reddit.com/r/typescript/comments/gd3z69/phaser_integration_with_angular_and_electron/)
- url: https://www.reddit.com/r/typescript/comments/gd3z69/phaser_integration_with_angular_and_electron/
---
Hello, I'm new to this sub. I'm starting to do some work with Phaser, as I am pretty familiar with typescript, html, and css. I started a project and got some work done before figuring I should strip back the assets and some implementation to provide a clean little template. I have a repo for it and thought I might share:

&amp;#x200B;

[https://github.com/TBosak/game-template](https://github.com/TBosak/game-template)

&amp;#x200B;

Enjoy! Feel free to fork and improve my mess.
## [3][Calling ReasonML from TypeScript (or Flow), the easy way.](https://www.reddit.com/r/typescript/comments/gdaoab/calling_reasonml_from_typescript_or_flow_the_easy/)
- url: https://www.hackdoor.io/articles/myNBk3px/calling-reasonml-from-typescript-or-flow-the-easy-way
---

## [4][Retturn type of a function that creates nth-dimensional arrays](https://www.reddit.com/r/typescript/comments/gd9hem/retturn_type_of_a_function_that_creates/)
- url: https://www.reddit.com/r/typescript/comments/gd9hem/retturn_type_of_a_function_that_creates/
---
I'm writing a function that will create and return an Nth-dimensional array, where N is set by a parameter, and the array can be filled with an initial value. I'm struggling to think of the correct return type for this function and would love some help.

Something like is what I'm after:

    const createNthDimensionalArray = &lt;T&gt;(
        NDimensions: number,
        initialValues?: T
    ): T[][] =&gt; {
        // Array creation happens here
    }

But the return type of the above function (`T[][]`) will only work for 2-dimensional arrays. Ideally I need a type that can ensure the arrays are only filled with the same type as the `initialValues` parameter, while being able to support however many dimensions the user needs, if such a thing exists!

&amp;#x200B;

Currently I'm using `any[]` as the return type, but I'm sure there's a better way.
## [5][Learning typescript as a junior developer](https://www.reddit.com/r/typescript/comments/gd23ti/learning_typescript_as_a_junior_developer/)
- url: https://www.reddit.com/r/typescript/comments/gd23ti/learning_typescript_as_a_junior_developer/
---
**TL;DR: is it common for junior developers to learn Typescript at work / trough a course offered by a company?**  


I'm a self taught developer and currently still in the process of learning. Currently I'm working on some more advanced topics, including Typescript. I feel the industry is not always expecting a junior to write TS compared to more experienced developers.  


The current course I'm following is not very good in terms of learning TS. Therefore I'm considering following a course fully dedicated to Typescript.  


I'd like to hear your thoughts on this: do you recommend to take such a course or shouldn't it hurt to start working without being comfortable with it (and waiting for a company to introduce me to it with either a course or something else)?
## [6][Typescript and react conditional render](https://www.reddit.com/r/typescript/comments/gcqdmt/typescript_and_react_conditional_render/)
- url: https://www.reddit.com/r/typescript/comments/gcqdmt/typescript_and_react_conditional_render/
---
I am new to typescript and not an expert in FE development. I've encountered issue that seems pretty basic, but I failed to found any solution. Maybe I just don't know how to google it properly.

In react component I have a button, that is disabled on some condition, which triggers a component's function:

    import React, {Component} from 'react';
    
    type DraftCompany = {
        id: null
        name: string,
    };
    
    type Company = Omit&lt;DraftCompany, 'id'&gt; &amp; {
        id: number;
    };
    
    type Props = {
        company: Company | DraftCompany,
        onDeleteCompany: (companyId: number) =&gt; void,
    }
    
    class CompanyRow extends Component &lt;Props&gt; {
        handleDeleteCompany = () =&gt; {
            this.props.onDeleteCompany(this.props.company.id);
        };
    
        render = () =&gt; {
            return (
                &lt;div&gt;
                    &lt;div&gt;{this.props.company.name}&lt;/div&gt;
                    &lt;div&gt;
                        &lt;button disabled={this.props.company.id === null} onClick={this.handleDeleteCompany}/&gt;
                    &lt;/div&gt;
                &lt;/div&gt;
            )
        }
    }
    
    export default CompanyRow;

I am getting typescript error on calling `this.props.onDeleteCompany(this.props.company.id);` that says that there is a chance I will pass `null` as a parameter. I fully understand why typescript gives me this error, the question is: **what would be the best way to deal with this error?**

&amp;#x200B;

I have found multiple ways:

&amp;#x200B;

1 ) Add `if` guard:

    handleDeleteCompany = () =&gt; {
            if (this.props.company.id) {
                this.props.onDeleteCompany(this.props.company.id);
            }
        };

It works, but I don't like the idea of adding such guards into every function, if someone removes `disabled` logic, I want to receive console error telling me about it immediately, not to have it be silently swallowed. In my project I have a lot of such code that relies on render, I doubt it is a best practice to add such checks everywhere. Correct me pls if I am wrong.

&amp;#x200B;

2) Apply `as` operator  to field :

     handleDeleteCompany = () =&gt; {
            this.props.onDeleteCompany(this.props.company.id as number);
        };

It works, but looks kinda hacky.

&amp;#x200B;

3) Apply `!` operator to field:

     handleDeleteCompany = () =&gt; {
            this.props.onDeleteCompany(this.props.company.id!);
        };

It works, but looks kinda hacky too.

&amp;#x200B;

4) Apply `as` operator to whole object and pass it to function:

     &lt;button disabled={this.props.company.id === null} 
             onClick={() =&gt; this.handleDeleteCompany(this.props.company as Company)}/&gt;

and

     handleDeleteCompany = (company: Company) =&gt; {
            this.props.onDeleteCompany(company.id as number);
        };

It works, but it looks like I am unnecessary passing the value I could have grabbed in function itself from props. I am not sure it is best practice to do such things.

&amp;#x200B;

I am sure there should be some pure typescript solution like defining Props type as a union or using conditional types with some combination of `any` and `never`. But I haven't figured it out.

As I said, I am not a FE expert so there is a high chance I am wrong about what is considered best practice. Help me please.

&amp;#x200B;

Here is a [playground](https://www.typescriptlang.org/play?#code/JYWwDg9gTgLgBAJQKYEMDGMA0cDeBhCcCAOyWJgF84AzKQuAcilQwYG4BYAKG5gE8wSOABEoKajALgUxPnAC8ubnBVxgAEwBccYgFcANvuWriKEEm0BnGFGDEA5pm4VOPLv0FwpYGXMUB5EGAYAB5RcUlCH1lsBg0GAD44ADIlLlU1LR1dEAAjJChXF25eASEABTowSwU0jLQo321vXzgAHxExCRaY4xUSYSR9JBgkHr5tAAoG6VkASSy9PIKASgUkgDcIDScuChKuNH0USxrxhAgAdzgkAA9R4nUzqJIyeBDKiGqknD64AAsZOphoNhqNxrVJmt5D8-hkYP9gJYAHRgKoogZDEZjRqySYIpGo9HImbRPjIjQrVwZYpuDLMR4FSHQ2HpDKqZgwXRQYhwSZw9kqELqYAbBICwUZYWihI4AkotFfFGk3zI0zmCghAD0IrFEslQt14rZBtNcBCuV0MBgJDgIssKFyw3U8jliIVxJVsgp6gU8kUekMVBIeH0wDQAGtXfLkYDHiCseDcXwKFrjWbBdqjfqpTqZfqVn99nsDndILA7UhqCgDPBzlc2EA)
## [7][[QUESTION] typeorm mysql one-to-one relation](https://www.reddit.com/r/typescript/comments/gcwzis/question_typeorm_mysql_onetoone_relation/)
- url: https://www.reddit.com/r/typescript/comments/gcwzis/question_typeorm_mysql_onetoone_relation/
---
Membership Entity:

    @Entity() 
    export default class Membership extends BaseEntity {
         @PrimaryColumn("uuid")
         id: string;
    
         @OneToOne(type =&gt; Member, member =&gt; member.membership)
         member: Member 
    } 

Member Entity:

    @Entity() 
    export default class Membership extends BaseEntity {
         @PrimaryColumn("uuid")
         id: string;
    
         @OneToOne(type =&gt; Membership, membership =&gt; membership.member)
         @JoinColumn()
         membership: Membership 
    } 

This setup works fine when I create a member (A) and hook it up to a membership (M). However, I get a `Duplicate Entry Error`  
 when I try to add another member (B) and hook it up to the same membership (M).

Can anyone please help me out. I'm pretty new to typeorm and mysql and would love to have your opinion on the issue.

Thanks!
## [8][I reverse engineered a npm registry from scratch. It is written entirely in TypeScript and has a low amount of dependencies.](https://www.reddit.com/r/typescript/comments/gc3q61/i_reverse_engineered_a_npm_registry_from_scratch/)
- url: https://github.com/Tanuel/baggy
---

## [9][Error in Typed Method Decorator](https://www.reddit.com/r/typescript/comments/gccnh3/error_in_typed_method_decorator/)
- url: https://www.reddit.com/r/typescript/comments/gccnh3/error_in_typed_method_decorator/
---
In my use case of method decorators I wanted them to work like a abstract method constraint.

So i want a method decorator to "force" a function signature but allow fewer arguments. Now to do that, simply adding a function as a type in the method decorator doesn't work. Adding the possible function types also doesn't work. So if anyone has a correction, feel free to comment it.

This is the code I came up with:  


    class Test {
      // works
      @Decorator('abc')
      test0(): string {
        return '';
      }
    
      // fails
      @Decorator('abc')
      test1({abc}: {abc: string}): string {
        return '';
      }
    
      // fails
      @Decorator('abc')
      test2({abc}: {abc: string}, context: any): string {
        return '';
      }
    }
    
    interface Signature&lt;T&gt; extends Function {
      (args: T, context: any): any;
      (args: T): any;
      (): any;
    }
    
    function Decorator&lt;T&gt;(...args: Array&lt;keyof T&gt;) {
      return function (target: any, key: string, descriptor: TypedPropertyDescriptor&lt;Signature&lt;T&gt;&gt;) {
        // do something
      }
    }

I also created a REPL [here](https://repl.it/repls/BossyConsiderateDistributionsoftware).
## [10][Emited code is ignoring nullable return value](https://www.reddit.com/r/typescript/comments/gcats8/emited_code_is_ignoring_nullable_return_value/)
- url: https://www.reddit.com/r/typescript/comments/gcats8/emited_code_is_ignoring_nullable_return_value/
---
I have a method declared with this signature on my TypeScript file:

https://preview.redd.it/3qmjo8gk3ew41.png?width=728&amp;format=png&amp;auto=webp&amp;s=174f3d04f08004df3be6f9dc44ca9990bc8bee4d

And when I compile this to JS+D.TS, it suddenly lose nullability in all types:

&amp;#x200B;

https://preview.redd.it/njmzewnv3ew41.png?width=770&amp;format=png&amp;auto=webp&amp;s=cd4fa11a8c584e1196766f5bec47ac81b305c24a

And TypeScript thinks I'm returning a non-nullable T (String in this case):

&amp;#x200B;

https://preview.redd.it/vx8uwmlf5ew41.png?width=162&amp;format=png&amp;auto=webp&amp;s=d3df21438004a7fbbac59ac0b35a4734ec682017

The d.ts file is generated correctly, but if hover over method's name, it also shows incorrect typing:

&amp;#x200B;

https://preview.redd.it/gybtbmh44ew41.png?width=739&amp;format=png&amp;auto=webp&amp;s=74e2bb369bc9e640152814a48fcb6a372fbb5b89

On the original TS file, it shows the correct typings:

&amp;#x200B;

https://preview.redd.it/m7fsrspb4ew41.png?width=907&amp;format=png&amp;auto=webp&amp;s=c624d027adc1813362623c4f177a3f2f3561598f

I don't know exactly what's wrong with this. I've searched across Google and haven't found anything. I've also tried with full nullable (T | null | undefined) on the return type but still doesn't work.

Note: Obviously, I can't mark neither "this" nor return type as optional ("?").
## [11][Using Typescript in a Javascript React app](https://www.reddit.com/r/typescript/comments/gc6h9o/using_typescript_in_a_javascript_react_app/)
- url: https://www.reddit.com/r/typescript/comments/gc6h9o/using_typescript_in_a_javascript_react_app/
---
I have a React app written entirely in Javascript. I want to start using Typescript but I don't want to go back and convert any of my old code to TS. The app is too large. I'd like to only use TS moving forward.

Are there any guides on how I can set this up? I use webpack.
