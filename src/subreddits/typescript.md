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
## [2][tsdotnet/linq 1.0 release](https://www.reddit.com/r/typescript/comments/gkmceo/tsdotnetlinq_10_release/)
- url: https://www.reddit.com/r/typescript/comments/gkmceo/tsdotnetlinq_10_release/
---
Fan of LINQ in .NET? Here's the TypeScript / JavaScript variant. :)

Full API Docs: [https://tsdotnet.github.io/linq/](https://tsdotnet.github.io/linq/)

GitHub: [https://github.com/tsdotnet/linq](https://github.com/tsdotnet/linq)  (100% test coverage.)

NPM: [https://www.npmjs.com/package/@tsdotnet/linq](https://www.npmjs.com/package/@tsdotnet/linq)

    npm i @tsdotnet/linq

Check out other [`tsdotnet`](https://github.com/tsdotnet) packages.

&amp;#x200B;

[tsdotnet](https://preview.redd.it/45dm3qln91z41.png?width=200&amp;format=png&amp;auto=webp&amp;s=dc78de3b5905fa4a3adb15b957de6800c9f3705b)
## [3][The Great CoffeeScript to Typescript Migration of 2017](https://www.reddit.com/r/typescript/comments/gki9k1/the_great_coffeescript_to_typescript_migration_of/)
- url: https://dropbox.tech/frontend/the-great-coffeescript-to-typescript-migration-of-2017
---

## [4][How to start contributing to Typescript?](https://www.reddit.com/r/typescript/comments/gkhkyh/how_to_start_contributing_to_typescript/)
- url: https://www.reddit.com/r/typescript/comments/gkhkyh/how_to_start_contributing_to_typescript/
---
I really like typescript and what it tries to do and solve. I guess it's a fairly complex project that can level up my coding and engineering skills.

Right now I'm a mid-level typescript user, I use it in a front-end React project and a sometimes on Nestjs project. What I want to do is read "Effective Typescript" book and start answering typescript questions on SO and start looking at PRs and Issues on the Github repo. 

What I don't know is how much people outside Microsoft are involved in coding and how to get to know the codebase in like 3 month and what are the prerequisites? do we need compiler knowledge? (I know some, lexical analysis, syntax analysis and...)
## [5][I created a zero-dependncy promise-based job queue with concurrency limiting](https://www.reddit.com/r/typescript/comments/gksmub/i_created_a_zerodependncy_promisebased_job_queue/)
- url: https://github.com/Tanuel/async-queue
---

## [6][Visual Code Plugin suggestion to help TypeScript become static type like](https://www.reddit.com/r/typescript/comments/gkdobw/visual_code_plugin_suggestion_to_help_typescript/)
- url: https://www.reddit.com/r/typescript/comments/gkdobw/visual_code_plugin_suggestion_to_help_typescript/
---
EDIT: Thank you everyone &lt;3. Changing the config file really helped!

EDIT 2: Thank you everyone again! I am sorry for the noob question I can not tell you how much I learned today and how grateful I am for your patience. (1) It was pointed out various times TypeScript differs from JavaScript because of Static Typing(dug lol), but as another redditor pointed out it is within itself. We cannot handle outside resources coming in. (2) The analysis on my code and insight how I could do things better. From a java only background all this new stuff from Angular is coming at me at 100mph and I am just trying to digest it all as best as I can. Thank you, thank you, thank you!

TLDR: Any plugins to make TypeScript static typish like Java.

&amp;#x200B;

&amp;#x200B;

Are there any plugins to help TypeScript static type. I know that at compile time it won't matter but I was hoping during writing code it would highlight things red or say you are missing a return, conflicting types, or something like that. I come from a Java background and I personally like things being of static typed. I don't mind parsing, or the extra steps. I have gotten too many bugs where an Object is returned, but not of a Object I made. Or sometimes a string is returned instead of a Object, then I send that string down a path where an object is suppose to be (despite explicitly asking for an object, not a string) and then indexOf doesn't work and I just don't like that. I know python offers something like that,but browsing through the plugins I couldn't find anything. Any recommendations?
## [7][Complex Type Deductions (my brain hurts)](https://www.reddit.com/r/typescript/comments/gkgl1o/complex_type_deductions_my_brain_hurts/)
- url: https://www.reddit.com/r/typescript/comments/gkgl1o/complex_type_deductions_my_brain_hurts/
---
I'm running into a **very** specific problem with type recognition.  My application is very type safe in a series of situations.  We have all of our language strings defined in one giant object.  The general structure for this object is shown below.  Below that are a series of types and examples that explain how we index the object.  All of our components and pages extend the base classes Component and Page defined below the examples.  The base classes do a number of things but here I've kept them "simple" to avoid ballooning this post anymore than it already is.  The way this structure is defined, it makes the rule that **all modules must contain a pages object**.  And that works for our use-cases.  However, I just set it up so that every instance of PageProperties always has a title property in it.  However, when it gets to the statement at the very bottom, TSC can't seem to deduce anything about the type of Page.prototype.pageProperties, even though it is a fact that pageProperties will always have a title element.  I want this to be enforced by the type-system, that way when a developer adds a new page with properties, their build will fail unless they provide a title.  


Can anyone think of a way to type-enforce this?

    const applicationConstants = {
        modules: {
            moduleA: {
                pages: {
                    moduleAFirstPage: {
                        title: "Module A first page",
                        someAFirstProperty: "Foo"
                    }
                },
                components: {
                    moduleAFirstComponent: {
                        someAComponentProperty: "Foo2"
                    }
                }
            },
            moduleB: {
                pages: {
                    moduleBFirstPage: {
                        title: "Module B first page",
                        someBFirstProperty: "Bar"
                    },
                    moduleBSecondPage: {
                        title: "Module B second page",
                        someBSecondProperty: "Baz"
                    }
                },
                components: {
                    moduleBFirstComponent: {
                        someBComponentFirstProperty: "Bar2"
                    }
                }
            }
        }
    } as const;
    
    const MODULES = "modules" as const;
    const PAGES = "pages" as const;
    type ModuleKey = typeof MODULES;
    type PagesKey = typeof PAGES;
    type ApplicationConstants = typeof applicationConstants;
    type ModuleName = keyof ApplicationConstants[ModuleKey];
    type SectionName&lt;Module extends ModuleName&gt; = keyof ApplicationConstants[ModuleKey][Module];
    type ComponentName&lt;
        Module extends ModuleName,
        Section extends SectionName&lt;Module&gt;
    &gt; = keyof ApplicationConstants[ModuleKey][Module][Section];
    type PageName&lt;Module extends ModuleName&gt; = ComponentName&lt;Module, PagesKey&gt;;
    type ComponentProperties&lt;
        Module extends ModuleName,
        Section extends SectionName&lt;Module&gt;,
        ComponentId extends ComponentName&lt;Module, Section&gt;
    &gt; = ApplicationConstants[ModuleKey][Module][Section][ComponentId];
    type PageProperties&lt;
        Module extends ModuleName,
        PageId extends PageName&lt;Module&gt;
    &gt; = ComponentProperties&lt;Module, PagesKey, PageId&gt;;
    
    // @ts-expect-error
    const invalidModuleName: ModuleName = "moduleD";
    
    const validModuleName: ModuleName = "moduleB";
    
    type ModuleBName = typeof validModuleName;
    
    // @ts-expect-error
    const invalidPageName: PageName&lt;ModuleBName&gt; = "moduleAFirstPage";
    
    const validPageName: PageName&lt;ModuleBName&gt; = "moduleBSecondPage";
    
    type ModuleBSecondPage = PageProperties&lt;ModuleBName, typeof validPageName&gt;;
    
    // @ts-expect-error
    const invalidPageProperties: ModuleBSecondPage = applicationConstants.modules.moduleA.pages.moduleAFirstPage;
    
    const validPageProperties: ModuleBSecondPage = applicationConstants.modules.moduleB.pages.moduleBSecondPage;
    
    class Component&lt;
        Module extends ModuleName,
        Section extends SectionName&lt;Module&gt;,
        ComponentId extends ComponentName&lt;Module, Section&gt;
    &gt; {
        public readonly componentProperties: ComponentProperties&lt;Module, Section, ComponentId&gt;;
    
        public constructor(module: Module, section: Section, componentId: ComponentId) {
            this.componentProperties = applicationConstants[MODULES][module][section][componentId];
        }
    }
    
    class Page&lt;
        Module extends ModuleName,
        PageId extends PageName&lt;Module&gt;
    &gt; extends Component&lt;Module, PagesKey, PageId&gt; {
    
        public readonly pageProperties: PageProperties&lt;Module, PageId&gt;;
    
        public constructor(module: Module, pageId: PageId) {
            super(module, PAGES, pageId);
    
            this.pageProperties = this.componentProperties;
    
            // At this point, when every page is constructed
            // I would like to do something with the title property
            // And every page has a title property
            // But unfortunately this throws an error which amounts to the idea that
            // none of the properties of pageProperties can be deduced here
            console.log(this.pageProperties.title);
        }
    
    }
## [8][Making GraphQL Magic with Sqlmancer](https://www.reddit.com/r/typescript/comments/gk2jby/making_graphql_magic_with_sqlmancer/)
- url: https://www.reddit.com/r/typescript/comments/gk2jby/making_graphql_magic_with_sqlmancer/
---
The official beta version of Sqlmancer has been released! Sqlmancer is a Node.js library that empowers you to effortlessly and efficiently translate GraphQL queries into SQL statements. I just wrote up a brief article that shows how to get started and showcases some basic features.

[https://medium.com/@danielrearden/making-graphql-magic-with-sqlmancer-d69c8203c087?sk=e4a66985889eee34f0b67d39f3bd5741](https://medium.com/@danielrearden/making-graphql-magic-with-sqlmancer-d69c8203c087?sk=e4a66985889eee34f0b67d39f3bd5741)

I also put together a CodeSandbox container with the code from the article. You can see the library in action without forking or installing anything :)

The library's still in development, so any feedback is always welcome. If you're interested in contributing, let me know or just open an issue. And if you like the library, please star it on GitHub!
## [9][Importing typescript types in js files . How to organize the type files?](https://www.reddit.com/r/typescript/comments/gk9njk/importing_typescript_types_in_js_files_how_to/)
- url: https://www.reddit.com/r/typescript/comments/gk9njk/importing_typescript_types_in_js_files_how_to/
---
So I want to write a node js module in javascript that looks something like this :

    //myModule.js 
    const foo = function (aBigObject) {
    	[
    		function fn1(_) {/*lots of code*/},
    		function fn2(_) {/*lots of code*/},
    		/*a lot of other functions*/
    	].forEach(fn =&gt; fn(aBigObject))
    }

I created a file `myModule.d.ts` in the same folder that `myModule.js` is located , and declared the appropriates types in it .

Now when I import `foo` in some other project I have intellisense for its parameters .

My issues is that when I try to edit something in `myModule.js` I do not have the very much needed intellisense for `aBigObject`. So in `myModule.js` file I use a JSDoc comment and import the types from the file `myModule.d.ts`.

1)Why do I have to do that manually via a JSDoc comment and vscode does not understand automatically the type?

Another issue that I face is that there are internal types that are needed in `myModules.js` while I develop it .They are imported via JSDoc comments. They really make developing way much easier . Nobody will have a need to use these internal types except for me that I develop `myModule.js` .

2)The question is , where do I put the internal types and why? Till so far I have them in file `myModule.d.ts`. Would it be wise to create a separate file for them , e.g. `myModules.internals.d.ts`?

I am not into going full ts and then compiling it to js and d.ts (and to be honest I do not even know how to do it) . I am just not into compiling because it requires much more knowledge than what I already do.

Edit : Another thing that I have noticed is that maybe it is better avoiding naming the `d.ts` file with the same name as the js file because like this I need to write less types . 

For example I have to:

    export declare function foo({/*lot of types and properties*/}) : void;

to get intellisense when I import `foo` in another project . 

But if I want to have intellisense for the parameters in `myModule.js`  while I am developing then I have to define in the `d.ts` file again like this :

    type foo = ({/*lot of types and properties*/}) =&gt; void;

and import it via a JSDoc comment.

If I just write only the `type ...` into `d.ts` file and do not write `declare ...` and import the type via JSDoc comment into `myModule.js` and the `d.ts` file does not have the same name as `myModule` then I have both intellisense when I develop `myModule.js` and when I import it in another project ,without the need to define the same type two times .

3)Is there any bad practice with that?
## [10][[Podcast] A visit to Deno Land with Core Member Kit Kelly (JS Party #127)](https://www.reddit.com/r/typescript/comments/gkhsdf/podcast_a_visit_to_deno_land_with_core_member_kit/)
- url: https://changelog.com/jsparty/127
---

## [11][How to I type a function that it has to not contain return ?](https://www.reddit.com/r/typescript/comments/gk63ea/how_to_i_type_a_function_that_it_has_to_not/)
- url: https://www.reddit.com/r/typescript/comments/gk63ea/how_to_i_type_a_function_that_it_has_to_not/
---
I thought that this :

    const foo:() =&gt; void = function() {
    	return true;
    }
    foo();

would lint error,but it does not.Am I doing something wrong with the type?How to make it lint error ?
