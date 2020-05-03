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
## [2][Typescript and react conditional render](https://www.reddit.com/r/typescript/comments/gcqdmt/typescript_and_react_conditional_render/)
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
## [3][I reverse engineered a npm registry from scratch. It is written entirely in TypeScript and has a low amount of dependencies.](https://www.reddit.com/r/typescript/comments/gc3q61/i_reverse_engineered_a_npm_registry_from_scratch/)
- url: https://github.com/Tanuel/baggy
---

## [4][Error in Typed Method Decorator](https://www.reddit.com/r/typescript/comments/gccnh3/error_in_typed_method_decorator/)
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
## [5][Emited code is ignoring nullable return value](https://www.reddit.com/r/typescript/comments/gcats8/emited_code_is_ignoring_nullable_return_value/)
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
## [6][Script unsigned error](https://www.reddit.com/r/typescript/comments/gchpqz/script_unsigned_error/)
- url: https://www.reddit.com/r/typescript/comments/gchpqz/script_unsigned_error/
---
I recently reinstalled Windows. Before I did, my TS files compiled fine from the command line with tsc.

Now, when I try to compile a TS file with tsc, I get an error that the script can't be run because it's unsigned (tsc.ps1 cannot be loaded).

Any idea why the before and after behavior is different?
## [7][Using Typescript in a Javascript React app](https://www.reddit.com/r/typescript/comments/gc6h9o/using_typescript_in_a_javascript_react_app/)
- url: https://www.reddit.com/r/typescript/comments/gc6h9o/using_typescript_in_a_javascript_react_app/
---
I have a React app written entirely in Javascript. I want to start using Typescript but I don't want to go back and convert any of my old code to TS. The app is too large. I'd like to only use TS moving forward.

Are there any guides on how I can set this up? I use webpack.
## [8][Open Source Typescript Express projects](https://www.reddit.com/r/typescript/comments/gc4xe3/open_source_typescript_express_projects/)
- url: https://www.reddit.com/r/typescript/comments/gc4xe3/open_source_typescript_express_projects/
---
Hello, I've been learning Typescript recently and decided to try to implement what I've learning in a new project.

I would like to have a look at some major projects in Typescript Express so that I can catch some best practices, design patterns, etc.

I'd appreciate if anyone could share links to some open source projects/companies that used these technologies for their product.
## [9][Why am I getting this TS error?](https://www.reddit.com/r/typescript/comments/gbpqiy/why_am_i_getting_this_ts_error/)
- url: https://i.redd.it/un2hvbg4r7w41.png
---

## [10][The Jupiter (YC S19) Stack - From TypeScript to Kubernetes and back — how we write, build, and push code at Jupiter](https://www.reddit.com/r/typescript/comments/gbl21c/the_jupiter_yc_s19_stack_from_typescript_to/)
- url: https://starship.jupiter.co/jupiter-stack/
---

## [11][Dynamic import performance evaluation](https://www.reddit.com/r/typescript/comments/gbt5k3/dynamic_import_performance_evaluation/)
- url: https://www.reddit.com/r/typescript/comments/gbt5k3/dynamic_import_performance_evaluation/
---
Hello everyone.
In order to reduce cold start time of my serverless functions, I was thinking about using dynamic imports inside each function in order to avoid initialization of dependencies I may not use.
I know that once a module is normally imported, it is preserved in the global scope and I can use it without any overhead. 
Is the same for dynamic imports? 
I know that it could be a silly question but I feel like I'm missing a piece of the puzzle here.
