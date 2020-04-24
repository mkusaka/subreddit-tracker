# typescript
## [1][Who's hiring Typescript developers - April](https://www.reddit.com/r/typescript/comments/fsojx3/whos_hiring_typescript_developers_april/)
- url: https://www.reddit.com/r/typescript/comments/fsojx3/whos_hiring_typescript_developers_april/
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
## [2][I got sick of constantly configuring packages, especially in mono repos, so I created a tool that takes away a lot of the configuration. ts-engine provides build, lint, test and typechecking support via single dependency 😁](https://www.reddit.com/r/typescript/comments/g6uphc/i_got_sick_of_constantly_configuring_packages/)
- url: https://ts-engine.dev
---

## [3][Best RESTful framework with good TypeScript support?](https://www.reddit.com/r/typescript/comments/g74vn8/best_restful_framework_with_good_typescript/)
- url: https://www.reddit.com/r/typescript/comments/g74vn8/best_restful_framework_with_good_typescript/
---
I'm used to working with Python backends; what I'm looking for is something like [FastAPI](https://github.com/tiangolo/fastapi) which uses type annotations to validate requests/response structures among other things. It's very handy. I've done a lot of looking but can't find a TypeScript-first API framework with similar features. Anyone have any recommendations?
## [4][Generate TS Types from a Postgres Database](https://www.reddit.com/r/typescript/comments/g6n301/generate_ts_types_from_a_postgres_database/)
- url: https://github.com/kristiandupont/kanel
---

## [5][Currency internationalization | eo-locale](https://www.reddit.com/r/typescript/comments/g74ok8/currency_internationalization_eolocale/)
- url: https://eo-locale.netlify.app/blog/money/
---

## [6][[🤘] Rockplate - Templating language with built in Linter, official VSCode IntelliSense extension, 5kb min.gzipped - written in TypeScript. Feedbacks are very welcome, thank you](https://www.reddit.com/r/typescript/comments/g6po38/rockplate_templating_language_with_built_in/)
- url: https://rockplate.github.io
---

## [7][Codedoc: Easily create beautiful and modern docs/wiki for your software projects](https://www.reddit.com/r/typescript/comments/g6nplp/codedoc_easily_create_beautiful_and_modern/)
- url: https://codedoc.cc
---

## [8][I've created small lib allowing you to create type-safe filters by assigning them to draft object (npm filterer)](https://www.reddit.com/r/typescript/comments/g659qu/ive_created_small_lib_allowing_you_to_create/)
- url: https://v.redd.it/4dynosj9leu41
---

## [9][Made a Animal Crossing Scraper with TypeScript!](https://www.reddit.com/r/typescript/comments/g6rrc4/made_a_animal_crossing_scraper_with_typescript/)
- url: https://github.com/matthewwwillard/ac_scrapper
---

## [10][Overriding class attribute using reflect metadata](https://www.reddit.com/r/typescript/comments/g6rll0/overriding_class_attribute_using_reflect_metadata/)
- url: https://www.reddit.com/r/typescript/comments/g6rll0/overriding_class_attribute_using_reflect_metadata/
---
Hi everyone,

I need your help with this overriding reusable controller's swagger description using decorators.

[https://stackoverflow.com/questions/61394012/override-swagger-attribute-in-a-reusable-controller](https://stackoverflow.com/questions/61394012/override-swagger-attribute-in-a-reusable-controller)
## [11][Trying to extend Array's interface for modules](https://www.reddit.com/r/typescript/comments/g6l2r3/trying_to_extend_arrays_interface_for_modules/)
- url: https://www.reddit.com/r/typescript/comments/g6l2r3/trying_to_extend_arrays_interface_for_modules/
---
Greetings,

i'm trying to extend Array's interface to have a more familiar and safe interface.

I have this module in "Array.ts"

    export{}
    
    declare global 
        {
        interface Array&lt;T&gt;
            {
            resize(size: number, fill?: T): Array&lt;T&gt;;
            at(index: number): T
            size(): number;
            empty(): boolean;
            }
        }
    
    if (!Array.prototype.resize) 
        {
        Array.prototype.resize = function&lt;T&gt;(size: number, fill?: T): T[] 
            {
            if(fill === undefined || size &lt;= this.size())
                { this.length = size; }
            else 
                {
                for(let i = this.size(); i &lt; size; i++)
                    { this[i] = fill; }
                }
            return this;
            }
        }
    
    if(!Array.prototype.at)
        { 
        Array.prototype.at = function&lt;T&gt;(index: number): T
            {
            if(index &lt; 0 || index &gt;= this.size()) { throw "Index out of array bounds exception."; }
            return this[index];
            }
        }
    
    if (!Array.prototype.size) 
        { Array.prototype.size = function(): number { return this.length; } }
    
    if (!Array.prototype.empty)
        { Array.prototype.empty = function(): boolean { return this.size() === 0; } }

And i'm importing it as follows:

    import {} from "./Util/Array.js";

Previously i was using it in a single file and it worked as expected (there was just the resize function).

Then i replaced all my instances of .length with .size() and added the import in all the modified files. The IDE doesn't report any error. However when i try running the project i get errors.

*  On Firefox: `TypeError: this.arr.size is not a function`
* On Edge: `Object doesn't support property or method 'size'`

Is this related to modules being cached rather than loaded twice? Even if that was the case shouldn't the prototype of array have gained the methods on the first time the module was imported?

Any idea on how to solve that?
