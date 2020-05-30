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
## [2][TS Blog Post: Changes to How We Manage DefinitelyTyped](https://www.reddit.com/r/typescript/comments/gsx450/ts_blog_post_changes_to_how_we_manage/)
- url: https://devblogs.microsoft.com/typescript/changes-to-how-we-manage-definitelytyped/
---

## [3][React Web &amp; Electron shared code mono repository](https://www.reddit.com/r/typescript/comments/gtas6x/react_web_electron_shared_code_mono_repository/)
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
## [4][How can I implement a search function for a drop-down menu?](https://www.reddit.com/r/typescript/comments/gte3ox/how_can_i_implement_a_search_function_for_a/)
- url: https://www.reddit.com/r/typescript/comments/gte3ox/how_can_i_implement_a_search_function_for_a/
---
A combo-box perhaps? I'm trying to create a search function that can search for a code within a drop-down menu that has loads of them.
## [5][Does tsconfig need ending / for directories?](https://www.reddit.com/r/typescript/comments/gt50hv/does_tsconfig_need_ending_for_directories/)
- url: https://www.reddit.com/r/typescript/comments/gt50hv/does_tsconfig_need_ending_for_directories/
---
In the documentation examples I did not see a trailing forward slash on the exclude properties.

Include did have it, but it was using both a `**` and `*` matcher. So because it was looking for a pattern, I understand the need.

I would like to know if in general, one should include an ending forward slash when specifying folders, and the property is not requesting a matcher. I did add it on `outDir` to be safe; some programmers can be pedantic about style though
## [6][[QUESTION] Non-null assertion](https://www.reddit.com/r/typescript/comments/gst9x0/question_nonnull_assertion/)
- url: https://www.reddit.com/r/typescript/comments/gst9x0/question_nonnull_assertion/
---
[Playground Link](https://www.typescriptlang.org/play/?ssl=7&amp;ssc=1&amp;pln=9&amp;pc=1#)

Is there any way to avoid the non-null assertion on line 52?
I want something similar to type predicate. When `isEmpty()` returns false I want typescript to know `this.head`. and `this.tail` are not null.

Also, in the `get()` method, I know `cur` is not null because I did an index check first. Can I somehow tell Typescript that `cur` will never be null as well?
## [7][Node library distribution best practices](https://www.reddit.com/r/typescript/comments/gszn32/node_library_distribution_best_practices/)
- url: https://www.reddit.com/r/typescript/comments/gszn32/node_library_distribution_best_practices/
---
Hi all, I recently released a small [library](https://github.com/mjhart/help-header) targeting node apps (express in particular). I am building a commonjs module because I still see a lot of node apps using commonjs, so I _think_ that would be most useful for my consumers.

I've done a fair amount of node application development, but this is my first time releasing a library to be incorporated into multiple other apps. I am really unsure of the tradeoffs among the various distribution options available. Any pointers or feedback would be greatly appreciated!
## [8][Architecture , compile and bundle steps on creating a lodash like library in ts .](https://www.reddit.com/r/typescript/comments/gsvio9/architecture_compile_and_bundle_steps_on_creating/)
- url: https://www.reddit.com/r/typescript/comments/gsvio9/architecture_compile_and_bundle_steps_on_creating/
---
So I want to create a library . My architecture looks like this :

    project-folder
    ├─node_modules
    ├─src
    │ └─modules
    │   ├─private
    │   │ ├─privateFn1.ts
    │   │ ├─privateFn2.ts
    │   │ ├─ ...
    │   │ └─privateFnN.ts
    │   ├─fn1.ts
    │   ├─fn2.ts
    │   ├─ ...
    │   └─fnN.ts
    ├─package.json
    └─tsconfig.json

Each `fn*.ts` exports a single function that has name `fn*` (i.e. the same name as the file name).

Each `privateFn*.ts` exports a single reference that is meant to be used by `fn*.ts` .

How would you create :

1. `./dist/cjs/index.js` which is the main entry of the `package.json` , i.e. commonjs format for exports .
2. `./dist/esm/index.js` which is the browser entry of the `package.json` , i.e. esmodules format for exports .
3. `./dist/types/index.d.ts` which is the typings entry of the `package.json` .

which :

* exports each one of those functions exported by `fn*.ts` .
* imports nothing (i.e. the file is bundled and has no depedency) .
* and default exports an object which contains as properties each one of those functions exported by `fn*.ts` .

I am asking for a more clean and fast way  from what I have done so far :

Here is my `tsconfig.json` :

```
{
	"compilerOptions": {
		"module": "ESNext",
		"target": "ESNext",
		"declaration": true,
		"rootDir": "./src",
		"outDir": "./built"
	},
	"include": ["./src"]
}
```
Here is what my npm scripts do :

1. I remove `./dist` `./src/index.ts` `./built` if they exist from any older build step .
2. I have created a node js script which I execute and it creates the file `./src/index.ts` which imports and exports all the functions and default exports an object which has as properties all those functions .
3. I use tsc to compile all ts files deep in `./src` with output to `./built`.
4. I use rollup cli to bundle `./built/index.js` to `./dist/cjs/index.js`
5. I use rollup cli to bundle `./built/index.js` to `./dist/mjs/index.js`
6. I use rollup with `rollup-plugin-dts` with a config file to bundle `./built/index.d.ts` to `./dist/types/index.d.ts`
7. I remove `./built` `./src/index.ts` .

I have created a final npm script called `build` which executed all the other scripts.So I do `npm run build`.

Edit : By the way is there an official bundler for ts created by Microsoft ? [Regarding bundling d.ts there is a 5 years feature request that is still open](https://github.com/microsoft/TypeScript/issues/4433) . What a disappointment .
## [9][Resources for Generics and Decorators](https://www.reddit.com/r/typescript/comments/gsz2hn/resources_for_generics_and_decorators/)
- url: https://www.reddit.com/r/typescript/comments/gsz2hn/resources_for_generics_and_decorators/
---
Hey guys,

Noob here,sorry if this is a generic question but can anyone point me to any reosurces which dumbs down generics and decorators,and why it would useful. Somr real life use cases,would be great too 

Thanks in advance
## [10][Hey, I need feedback on my first TypeScript - Node app.](https://www.reddit.com/r/typescript/comments/gsmupu/hey_i_need_feedback_on_my_first_typescript_node/)
- url: https://www.reddit.com/r/typescript/comments/gsmupu/hey_i_need_feedback_on_my_first_typescript_node/
---
GitWiz is a portal to search for public repos from multiple version control platforms.

https://i.redd.it/2lngawns4n151.gif

Here's the link to the Heroku deployment - [https://gitwiz.herokuapp.com](https://gitwiz.herokuapp.com)

and the link to the GitHub repo -  [https://github.com/KSSBro/gitwiz](https://github.com/KSSBro/gitwiz) 

I'm currently working on adding more platforms.
## [11][Looking for a TypeScript interface generator that will generate keys in consistent order, and re-use existing interfaces regardless of parent key name (perhaps via "shape" checksums?)](https://www.reddit.com/r/typescript/comments/gspj8d/looking_for_a_typescript_interface_generator_that/)
- url: https://www.reddit.com/r/typescript/comments/gspj8d/looking_for_a_typescript_interface_generator_that/
---
* I have millions of JSON+XML files collected over many years, mostly in 3rd party formats that I don't control (web scraping of JSON data).
* But there's only about  5 different "high-level-formats" involved.  But within each of the 5 formats I need to expect that the exact layout of the JSON/XML structures will have all sorts of unexpected slight variations over time.

I'm looking for an NPM package or similar that:

* From sample objects (parsed from JSON/XML)... generates TypeScript interfaces - there's plenty of packages that do this... but from what I've seen in them, I'm going to end up with 100s/1000s of redundant interfaces simply due to key order and differing parent key names.  To explain further...
* If data samples have the keys in different orders
 *  e.g. `{b:2,a:1}` and another similar file like... `{a:1,b:2}` ... they should still be treated as exactly the same interface.  This means that the keys will need to be sorted alphabetically for consistency.  
* Recursion to be handled with separate interfaces. i.e. Each interface is just one flat levels of properties (which may refer to other generated interfaces).
* In terms of what the interfaces themselves are named... The parent key of an object shouldn't be important, I want objects of the same shape to use the same generated interface, regardless of what the parent keys' names were from the samples.
 *  i.e. In the interface generation libs I've seen so far, the names of the interfaces are taken from the parent key, i.e. `{userinfo: {name:'John, age:30}}` would generate an interface named `Userinfo`.  And another sample object like `{person: {name:'John, age:30}}` would generate another interface named `Person` ... even though the contents of the interfaces are exactly the same.  I only want to generate one interface per "shape".  I'm guessing that this means that the interface names will need to be checksums of their shape.  Not sure if there is another more simple solution than that, aside from having some kind of lookup-existing-interfaces-by-checksum database or something.

Any pointers to some libraries that might be handy here?

I'm also not sure if I should be looking at "JSONschema" generation first, and then generating the interfaces from that?  Although I'm not really doing "validation"... I'm processing existing data with old immutable data formats, and just trying to make it easier to figure out how much the data shapes actually vary within each "high-level-format", and then make my code that processes it all type safe so I can use ctrl-space intellisense/autocomplete to write the code easily.

Maybe there's some even better way of dealing with this than TypeScript interfaces?  Keen for any tips related to processing large amounts of similar-but-slightly-different JSON/XML formats like this. Without needing to write tons of `if` / `switch` statements.
