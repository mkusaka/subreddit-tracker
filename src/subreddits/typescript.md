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
## [2][[NOOB QUESTION] How is typescript better if it is just compiled to javascript?](https://www.reddit.com/r/typescript/comments/gttwke/noob_question_how_is_typescript_better_if_it_is/)
- url: https://www.reddit.com/r/typescript/comments/gttwke/noob_question_how_is_typescript_better_if_it_is/
---
Since it is compiled to JS, the performance cannot be better than a natively written JS. 

Besides static type checking, what benefits do TS offer? I don't think it offers a performance boost. I'm confused if I should port all my JS code to TS.
## [3][Not clear on generic syntax](https://www.reddit.com/r/typescript/comments/gttpuz/not_clear_on_generic_syntax/)
- url: https://www.reddit.com/r/typescript/comments/gttpuz/not_clear_on_generic_syntax/
---
From docs:

    function identity&lt;T&gt;(arg: T): T {
      return arg;
    }

My issue is with the initial `&lt;T&gt;` just before the parameter definitions. That one appears to assert that the function `identity` is of type `T`. But the later 2 usages then suggest that the function returns itself, which I doubt is very useful above.

If the above is not a type assertion, how do you distinguish a type assertion ("This is asserted to be a string") from a "generic assertion"?

And what is even the purpose of putting it there? I guess just to say "This is entering the parameter definition and function body as a generic type"?

&amp;#x200B;

Here is some sample code I am working on:

    class SetupWizard {
      public printInstructions(option: optionClass) {
         console.log(optionClass.instructions)
      }
    }

Should the function header instead change to this:

    public printInstructions&lt;optionClass&gt;(option: optionClass)

All examples on the TS site use single character generic types. If you guys have an opinion on sticking to that as a convention, I am also curious why.
## [4][Generic error: Property 'print' does not exist on type 'C'](https://www.reddit.com/r/typescript/comments/gtv3f3/generic_error_property_print_does_not_exist_on/)
- url: https://www.reddit.com/r/typescript/comments/gtv3f3/generic_error_property_print_does_not_exist_on/
---
Not sure but the two errors below might be related. How do you guys clear "property does not exist" errors on generic types?

More generally, let's say there were 5+ possible instance types for the parameter `childInstance` below. would you type the property as a union, or keep it as a generic as shown? I am curious when you guys would move away from unions here; I'm guessing it has to do with the unknowns regarding possible new types that may be passed in, in the future.

    class Child1 { public print() { console.log("from Child1 instance"); } }
    class Child2 { public print() { console.log("from Child2 instance"); } }
    
    class Parent { 
        public print&lt;C&gt;(childInstance: C) { 
            childInstance.print(); // Property 'print' does not exist on type 'C'.(2339)
        } 
    }
    
    const child1 = new Child1();
    const child2 = new Child2();
    const parent = new Parent(); // Cannot redeclare block-scoped variable 'parent'.(2451)
    // input.tsx(18, 16): 'parent' was also declared here.
    
    parent.print(child1);
    parent.print(child2);
## [5][does typescript compile unused codes?](https://www.reddit.com/r/typescript/comments/gtwp3j/does_typescript_compile_unused_codes/)
- url: https://www.reddit.com/r/typescript/comments/gtwp3j/does_typescript_compile_unused_codes/
---
i am starting to new project with nrwl nx 
there are shared codes between frontend and backend and mobile apps.

i could not find a document about this subject

is there anyone point me the right documentation
## [6][Generic constraints](https://www.reddit.com/r/typescript/comments/gtlbz9/generic_constraints/)
- url: https://www.reddit.com/r/typescript/comments/gtlbz9/generic_constraints/
---
Coming from .Net, much of my Typescript code tends to look like C# so I apologise in advance.

Im confused by generics at the best of times but in C# this is quite easy.

I have the following class:

&amp;#x200B;

https://preview.redd.it/2qsovt47ry151.png?width=449&amp;format=png&amp;auto=webp&amp;s=0bc871f63da34ddac2d4932dc6ad9f0c147a8b02

In the constructor, I shouldn't (I think) need to pass itemType.  I should just be able to new T() inside the foreach but I cant see how to declare T as newable.  In C# I would just use

where T: SomeInterface, new()

Can anyone else?

&amp;#x200B;

Thanks!
## [7][[AskTS] how to render .tsx to html without react dependency?](https://www.reddit.com/r/typescript/comments/gtodnf/askts_how_to_render_tsx_to_html_without_react/)
- url: https://www.reddit.com/r/typescript/comments/gtodnf/askts_how_to_render_tsx_to_html_without_react/
---
I have some components in .tsx files that I want to render to HTML strings. Right now, there’s a dependency on react and react-dom/server. It’s working, but it feels like there should be an easier way that doesn’t require the full React library as well as another supplemental library. 

Any ideas or suggestions for other packages are welcome!
## [8][React Web &amp; Electron shared code mono repository](https://www.reddit.com/r/typescript/comments/gtas6x/react_web_electron_shared_code_mono_repository/)
- url: https://www.reddit.com/r/typescript/comments/gtas6x/react_web_electron_shared_code_mono_repository/
---
Hi,

I'm trying to build an application which will be available in the Browser &amp; native via Electron. Since both applications share components &amp; redux functionality I would like to have a shared source folder. I'm usually a C++ &amp; Golang developer, but I already have 2 medium projects worth of experience with create-react-app, electron-forge &amp; jsx. Since this project will probably be maintained for some years (hopefully) I thought Typescript would be the better option in this case. Since I already think in types when coding the transition from plain jsx shouldn't be to difficult.

I started the prototype with create-react-app + typescript &amp; electron-react-boilerplate, but soon found limitations with this setup. create-react-app doesn't allow source files outside of it's root, so the wanted folder structure shown below would be impossible. Also the although the electron boilerplate seems to be configured pretty feature rich, it's hard for me to tell exactly what configurations are really necessary and which "nice to have". 

**So basically my Question is:**

- Is typescript the right choice?
- What folder structure do you recommend?
- Should I use boilerplates or go from scratch?
- What is easier to keep up to date (boilerplate vs. scratch)

### Goals Summary

- typescript
- mono repo
- shared code
- react web
- react electron


#### Project Structure

```sh
.
├── electron
│   ├── package.json
│   └── src
├── shared
│   ├── package.json (if needed)
│   └── src
└── web
    ├── package.json
    └── src
```

### Packages

- react
- electron
- redux
- sass
- testing
- types
- and more
## [9][TS Blog Post: Changes to How We Manage DefinitelyTyped](https://www.reddit.com/r/typescript/comments/gsx450/ts_blog_post_changes_to_how_we_manage/)
- url: https://devblogs.microsoft.com/typescript/changes-to-how-we-manage-definitelytyped/
---

## [10][How can I implement a search function for a drop-down menu?](https://www.reddit.com/r/typescript/comments/gte3ox/how_can_i_implement_a_search_function_for_a/)
- url: https://www.reddit.com/r/typescript/comments/gte3ox/how_can_i_implement_a_search_function_for_a/
---
A combo-box perhaps? I'm trying to create a search function that can search for a code within a drop-down menu that has loads of them.
## [11][Trigger button click from another component using service](https://www.reddit.com/r/typescript/comments/gtjc8f/trigger_button_click_from_another_component_using/)
- url: https://www.reddit.com/r/typescript/comments/gtjc8f/trigger_button_click_from_another_component_using/
---
How to append div that trigger button click from another component using service? Any tips how to do this?
