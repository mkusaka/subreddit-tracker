# typescript
## [1][Who's hiring Typescript developers - February](https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/)
- url: https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/
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
## [2][Forget NodeJS! Build native TypeScript applications with Deno 🦖](https://www.reddit.com/r/typescript/comments/f9fayv/forget_nodejs_build_native_typescript/)
- url: https://deepu.tech/deno-runtime-for-typescript/
---

## [3][Modern Typescript with Examples Cheat Sheet](https://www.reddit.com/r/typescript/comments/f9es9d/modern_typescript_with_examples_cheat_sheet/)
- url: https://www.reddit.com/r/typescript/comments/f9es9d/modern_typescript_with_examples_cheat_sheet/
---
I was looking for something like this but could find nothing that fitted. The goal is to be a reference when looking things up with great examples that only demonstrate the syntax and problem/solution, not a load of other stuff which tends to be the case in blog examples.

Once I am happy with it I will create an attractive formatted PDF to print and have that as a link too.

I have taken examples from mainly the official docs and various blog posts by some great authors including quite a bit of Marius Schulz.

The other goal is to try and include only what most people would need, and omit any old ways of doing things, so it is quite opinionated like that. 

In my personal TS journey this is all the stuff I have needed and wanted a reference to come back to, especially with easy to understand examples. Obvious things are left out:

[https://github.com/David-Else/modern-typescript-with-examples-cheat-sheet](https://github.com/David-Else/modern-typescript-with-examples-cheat-sheet)

Feedback and pull requests very welcome, especially factual corrections!
## [4][tsParticles - Resurrecting particles.js](https://www.reddit.com/r/typescript/comments/f9p0gf/tsparticles_resurrecting_particlesjs/)
- url: https://www.reddit.com/r/typescript/comments/f9p0gf/tsparticles_resurrecting_particlesjs/
---
tsparticles - Resurrecting particles.js

Hi all,

I'm resurrecting the old particles.js, updated and converted to typescript. I've implemented some new features and fixes requested by users on the old project.

[https://www.npmjs.com/package/tsparticles](https://www.npmjs.com/package/tsparticles)
## [5][Types for PlantUML Parser: In this post I demonstrate how to parse PlantUML with TypeScript in a few simple steps.](https://www.reddit.com/r/typescript/comments/f9c4hy/types_for_plantuml_parser_in_this_post_i/)
- url: https://duckpond.ch/plantuml-parser/javascript/2020/02/25/types-for-plantuml-parser.html
---

## [6][Is it not possible to only accept instances of a class that have a particular static method *on the class itself*?](https://www.reddit.com/r/typescript/comments/f9i892/is_it_not_possible_to_only_accept_instances_of_a/)
- url: https://www.reddit.com/r/typescript/comments/f9i892/is_it_not_possible_to_only_accept_instances_of_a/
---
I've tried about 100 different things, and basically want I want to achieve is simply


    interface Serializable {
      serialize(): this;
      static deserialize(input: string): this;
    }

Obviously there's no way to define static methods on an interface, but is there no other way to achieve this if I control how the parameters of a function are typed?
## [7][Serverside typescript bundling with webpack](https://www.reddit.com/r/typescript/comments/f9gkb4/serverside_typescript_bundling_with_webpack/)
- url: https://www.reddit.com/r/typescript/comments/f9gkb4/serverside_typescript_bundling_with_webpack/
---
As the title says, I’m trying to bundle a typescript app with webpack using ts-loader. I have several other typescripts module as dependencies which are declared in tsconfig.json paths section. Tsc works great. When using webpack and tsconfig-paths-webpack-plugin, those modules are treated as external and are not bundled (with the message no static exports). If I’m importing them using a relative path, then their transpiled code is added into the bundle. Is there a way to add their code to the bundle without changing every import in the project to a relative path?
## [8][CSS not loaded](https://www.reddit.com/r/typescript/comments/f9gakw/css_not_loaded/)
- url: https://www.reddit.com/r/typescript/comments/f9gakw/css_not_loaded/
---
Hi I tried to deploy to netlify a React TS app. The problem is that no of my css go built, I was using Styled-components together with typescript, is thier anything I should do before i build my project? Thank you all!
## [9][Test if a property exists on a disjoint type?](https://www.reddit.com/r/typescript/comments/f9fdrk/test_if_a_property_exists_on_a_disjoint_type/)
- url: https://www.reddit.com/r/typescript/comments/f9fdrk/test_if_a_property_exists_on_a_disjoint_type/
---
If I have a Foo with a 'name' but Bar doesn't have a name and I have:

  type MyType = Foo | Bar;

.. it SEEMS like I should be able to do:

  if (myType.name) {
  ... 
  }

But typescript says that this is an error as 'name' doesn't exist on Bar.

What's the best way to do this?  I don't want to cast it to &lt;any&gt; because I want to know that this type exists in Foo ... but I'm not actually sure it IS a foo yet (thus the test for .name)

Thoughts?
## [10][When creating a debounce decorator function, how do I inform TypeScript there will be args?](https://www.reddit.com/r/typescript/comments/f9asjm/when_creating_a_debounce_decorator_function_how/)
- url: https://www.reddit.com/r/typescript/comments/f9asjm/when_creating_a_debounce_decorator_function_how/
---
I was trying to convert a previous debounce decorator function I'd made in JS to TS, and am running into an issue where the TS interpreted is expecting 0 args, but I need to provide one. Here's the code:

    function debounce(fn: { (message?: any, ...optionalParams: any[]): void; (message?: any, ...optionalParams: any[]): void; apply?: any; }, minTime: number){
        let isCooldown = false;
        return function(){
            if (isCooldown){ return; }
            fn.apply(this, arguments);
            isCooldown = true;
            setTimeout(() =&gt; isCooldown = false, minTime)
        }
    }

    let log = debounce(console.log, 1000);

    log('Hi'); // runs immediately
    log('Ignore me!');

    setTimeout(() =&gt; log('Ignore me as well!'), 100); // ignored ( only 100 ms passed )
    setTimeout(() =&gt; log('Don\'t ignore me!!'), 1100); // Runs because enough time has passed
    setTimeout(() =&gt; log('Still ignored!'), 1500); //  ignored (less than 1000 ms from the last run)
    setTimeout(() =&gt; log('Good now you get it!!'), 2200); // Runs because enough time has passed

Everywhere I use the log function, tslint tells me expected 0 args but got 1. Any advice is greatly appreciated, still new to TypeScript
## [11][Search your codebase in natural language (Metacode, a vscode extension)](https://www.reddit.com/r/typescript/comments/f8vc6y/search_your_codebase_in_natural_language_metacode/)
- url: https://i.redd.it/hbys8ycwwwi41.gif
---

