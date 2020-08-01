# typescript
## [1][Who's hiring Typescript developers - August](https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/)
- url: https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/
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
## [2][a tsconfig guide that actually makes sense](https://www.reddit.com/r/typescript/comments/i1m8lg/a_tsconfig_guide_that_actually_makes_sense/)
- url: https://www.stackchief.com/blog/tsconfig%20%7C%20the%20missing%20docs
---

## [3][How do you generate, organize and maintain your seeds/fixtures data for integration or e2e testing](https://www.reddit.com/r/typescript/comments/i1nmsk/how_do_you_generate_organize_and_maintain_your/)
- url: https://www.reddit.com/r/typescript/comments/i1nmsk/how_do_you_generate_organize_and_maintain_your/
---
Hello all,

I'd like to write mainly e2e and integration tests but the main challenge to me is to put my system in a certain state.

I'm unable to find good resources about how to generate, organize and maintain fixtures / seeds data.

Everytime I need to start writing an e2e test this part is the most challenging.

I just want to know if there are better alternatives than manually writing your set of fixtures that you reuse and modify per test case.

Is reusing other parts of your own API to generate data in other tests a bad thing? Like imagine testing the case "As a user I can only see my own books", do you insert a user and books in the database directly before running the test or do you use your own API operations like "user register", "user add books" to populate the database?

Do you always start from an empty database or do you have a consistent set of data that you reset between tests? (The consistent set seems easier from a test preparation perspective but is also less explicit as you can't see the whole picture in the test case).
## [4][Problem exporting typings from a package: module is imported from 'package-name/build'](https://www.reddit.com/r/typescript/comments/i1ozwt/problem_exporting_typings_from_a_package_module/)
- url: https://www.reddit.com/r/typescript/comments/i1ozwt/problem_exporting_typings_from_a_package_module/
---
Hi, I'm having trouble configuring the package.json and tsconfig.json for a package with the following structure:

    package-name
        build               \
            lib             |
                index.js    |
            test            | tsc output
                specs.js    |
            index.js        /

        src                 \
            lib             |
                index.ts    |
            test            | source files
                specs.ts    |
            index.ts        /
        package.json
        tsconfig.json

When I'm trying to import something from such a package, the auto-import (I'm using Webstorm if that makes a difference)
is always something like `import { Something } from 'package-name/build';` while I would want it to be simply
`import { Something } from 'package-name';`.

Imports from `package-name/build` seems to work at runtime, but my IDE can not resolve it properly. If I manually fix
the import from `package-name/build` to `package-name`, my IDE can resolve the types and it also works fine at runtime.

I tried a lot of things with the `main` and `types` properties in the package.json, also with the `sourceRoot`,
`baseUrl` and `rootDir` properties in the tsconfig.json but to no avail, I never succeeded at making auto-imports be
imported from `package-name`.

You can try it yourself by pulling [this repo](https://gitlab.com/gpascal/export-types-example) (which uses yarn workspaces to link the packages locally):

    git clone https://gitlab.com/gpascal/export-types-example
    cd export-types-example
    yarn install

Then open the project in your IDE, go to `packages/package-import/src/lib/index.ts`, delete the import statement and try to auto-import `TestType` or `testFunction`.

Any help would be greatly appreciated.

---------------------

For reference, the tsconfig.json and package.json look like this:

tsconfig.json:

    {
        "compileOnSave": true,
        "compilerOptions": {
            "target": "es2017",
            "module": "commonjs",
            "outDir": "build",
            "noImplicitAny": false,
            "moduleResolution": "node",
            "inlineSources": true,
            "inlineSourceMap": true
        },
        "include": ["src"]
    }

package.json:

    {
        "name": "package-name",
        "version": "1.0.0",
        "main": "build/index.js",
        "types": "src/index.ts",
        "devDependencies": {
            "@types/node": "^14.0.27",
            "typescript": "^3.9.7"
        }
    }


Edit: added link to the test repo
## [5][What's the difference?](https://www.reddit.com/r/typescript/comments/i1djdm/whats_the_difference/)
- url: https://www.reddit.com/r/typescript/comments/i1djdm/whats_the_difference/
---
Hey guys, javascript dev here. I've been looking through the Typescript documentation trying to understand why is everybody in love with this programming language.

Apart from the types in javascript and the abstract classes, what are the differences between JS and TS? Apart from the previously said stuff, why is it better to use TS rather than JS?

&amp;#x200B;

Thank you
## [6][How do I reference an object property by string?](https://www.reddit.com/r/typescript/comments/i17xcj/how_do_i_reference_an_object_property_by_string/)
- url: https://www.reddit.com/r/typescript/comments/i17xcj/how_do_i_reference_an_object_property_by_string/
---
I've been stuck on this for some time trying to simplify my code.

I have an object for building a visualisation, it looks something like


```
const visualisation = {   
   ...   
  themes: {   
   hyper: {  
    ...  
}    
}  
```

The idea is that the theme is changeable, I have a dropdown where I have various values that match to the name of the themes.

When I reference my object though this works `visualiserThemes.themes['hyper']` and returns the correct object
However when I get a value as a string I cannot reference the object property anymore.
so this doesn't work
```const value = styleSelect.options[styleSelect.selectedIndex].value
```  
```
visualiserThemes.themes[value]
```  
where `value` is a string

Setting a custom const as a string manually though does work, like this

```
const value2 = 'hyper';
visualiserThemes.themes[value2]
```

The errors I'm getting 

```
Element implicitly has an 'any' type because expression of type 'string' can't be used to index type...  
No index signature with a parameter of type 'string' was found on type...
```

I really don't understand what I'm doing wrong, in both examples the value and value2 constants are strings but the former just errors out.
Sorry for the formatting, reddit has horrible code formatting.
## [7][How to resolve imports in Typescript Mocha Testing?](https://www.reddit.com/r/typescript/comments/i17kgp/how_to_resolve_imports_in_typescript_mocha_testing/)
- url: https://www.reddit.com/r/typescript/comments/i17kgp/how_to_resolve_imports_in_typescript_mocha_testing/
---
I'm having an issue importing various classes for mocha testing, for my Typescript project. I have included the following stackoverflow post for reference:

[https://stackoverflow.com/questions/63192719/how-to-resolve-class-is-not-a-constructor-error-in-typescript-mocha-testing](https://stackoverflow.com/questions/63192719/how-to-resolve-class-is-not-a-constructor-error-in-typescript-mocha-testing)

npm start works just fine but npm test is throwing an error where \[Class\] is not a constructor for most classes in my project.
## [8][Bueno - Modern, composable validation schemas tailored for TypeScript.](https://www.reddit.com/r/typescript/comments/i0oqw9/bueno_modern_composable_validation_schemas/)
- url: https://github.com/philipnilsson/bueno#quickstart
---

## [9][Building Angular Projects in Visual Studio!?](https://www.reddit.com/r/typescript/comments/i0ptcs/building_angular_projects_in_visual_studio/)
- url: https://www.reddit.com/r/typescript/comments/i0ptcs/building_angular_projects_in_visual_studio/
---
Hi! Im Gabby a Program Manager at Microsoft and I'm here to present my teams latest work:

The Angular Language Service. Now available in Visual Studio!

Read more here: [https://www.nowayshecodes.com/blog-tech/angular-language-service-for-visual-studio](https://www.nowayshecodes.com/blog-tech/angular-language-service-for-visual-studio) !!!

Reply if you have any questions :)

&amp;#x200B;
## [10][Dependency Injection: inversifyJS vs tsyringe](https://www.reddit.com/r/typescript/comments/i0jdmn/dependency_injection_inversifyjs_vs_tsyringe/)
- url: https://www.reddit.com/r/typescript/comments/i0jdmn/dependency_injection_inversifyjs_vs_tsyringe/
---
 I am new to DI.

I could not find a good comparison of the two frameworks. On the first look tsyringe looks more elegant to me. It seems to auto register classes. In inversifyJS you seem to have to bind every class to the container first before it couls be injected? On the other hanf inversifyJS is much more popular (but is also older). Why is this?

Any PROs &amp; CONs?

Thank you!
## [11][Running ts-node vs compiled JS files](https://www.reddit.com/r/typescript/comments/i07561/running_tsnode_vs_compiled_js_files/)
- url: https://www.reddit.com/r/typescript/comments/i07561/running_tsnode_vs_compiled_js_files/
---
New TS user here and while setting up a dev workflow, I am confused between the two for local development:  

1. Use `ts-node` to watch for changes and run `*.ts` files without transpiling to JS.
2. Transpile to JS first and then run the `*.js` files.

Option 1 looks more natural to me but I wouldn't know for sure if my transpiled JS files work as expected. Option 2 would make sure my transpiled JS files work as expected but that would require transpiling on EVERY file change. This question was asked [last year](https://www.reddit.com/r/typescript/comments/9o1zzp/tsnode_vs_compiling_for_development/) as well but I couldn't find a satisfactory answer.

Thanks!
