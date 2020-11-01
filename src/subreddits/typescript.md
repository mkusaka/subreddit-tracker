# typescript
## [1][Who's hiring Typescript developers - November](https://www.reddit.com/r/typescript/comments/jlsywo/whos_hiring_typescript_developers_november/)
- url: https://www.reddit.com/r/typescript/comments/jlsywo/whos_hiring_typescript_developers_november/
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
## [2][Query regarding Excess Property Checks in TypeScript](https://www.reddit.com/r/typescript/comments/jlyvu3/query_regarding_excess_property_checks_in/)
- url: https://www.reddit.com/r/typescript/comments/jlyvu3/query_regarding_excess_property_checks_in/
---
Hi Everyone, 

I am unable to understand on why is `mySquare1` valid while `mySquare2` ends up in an error?

    interface SquareConfig {
      color: string;
      width: number;
    }

    function createSquare(config: SquareConfig): { color: string; area: number } {
      return { color: config.color || "red", area: config.width ? config.width*config.width : 20 };
    }

    let mySquare1 = { color:'red', width: 10, opacity: 15}
    let mySquare2 = createSquare(mySquare1);
    let mySquare = createSquare({ width: 100, color: 'white', opacity: 15 }); // Error
    // Argument of type '{ width: number; color: string; opacity: number; }' is not assignable to parameter of type 
    'SquareConfig'. Object literal may only specify known properties, and 'opacity' does not exist in type 
    'SquareConfig'.
## [3][Semicolons, yay or nay?](https://www.reddit.com/r/typescript/comments/jlsqxj/semicolons_yay_or_nay/)
- url: https://www.reddit.com/r/typescript/comments/jlsqxj/semicolons_yay_or_nay/
---
What is your preference for typescript?

[View Poll](https://www.reddit.com/poll/jlsqxj)
## [4][Strict version of built-in types (which use `any`)](https://www.reddit.com/r/typescript/comments/jlhzuv/strict_version_of_builtin_types_which_use_any/)
- url: https://www.reddit.com/r/typescript/comments/jlhzuv/strict_version_of_builtin_types_which_use_any/
---
Hello,

in Typescript, some function calls like

    const user = JSON.parse(something}
    
    console.log(user.metadata.name)

Or 

    fetch('/')
      .then(res =&gt; res.json())
      .then(user =&gt; { console.log(user.metadata.name) })

might result in runtime TypeError as they use \`any\` type for native JS and DOM APIs (instead of \`unknown\` as it didn't exist back then I suppose) .

  
(Links to definitions for the examples above: [https://github.com/microsoft/TypeScript/blob/71a2c59c557d871a1d38d10df83cfc36dc10d887/lib/lib.es5.d.ts#L1059](https://github.com/microsoft/TypeScript/blob/71a2c59c557d871a1d38d10df83cfc36dc10d887/lib/lib.es5.d.ts#L1059) and [https://github.com/microsoft/TypeScript/blob/71a2c59c557d871a1d38d10df83cfc36dc10d887/lib/lib.dom.d.ts#L2557](https://github.com/microsoft/TypeScript/blob/71a2c59c557d871a1d38d10df83cfc36dc10d887/lib/lib.dom.d.ts#L2557))

&amp;#x200B;

Is there a community version of these native types which would use \`unknown\` instead of  \`any\` for these kinds of functions?  


Or is there any other way to tackle this problem without requiring me to remember all the cases where typescript has got \`any\` as the built-in type?
## [5][Good practices to follow for a Typescript project?](https://www.reddit.com/r/typescript/comments/jlhuc4/good_practices_to_follow_for_a_typescript_project/)
- url: https://www.reddit.com/r/typescript/comments/jlhuc4/good_practices_to_follow_for_a_typescript_project/
---
Hi I'm going to start working on my first big project with Typescript using React Native. I'm still planning the whole system architecture but I need some help with how I should structure the code and files, document it, etc. Are there any resources that you would recommend for me?
## [6][Pathing packages in a TS monorepo?](https://www.reddit.com/r/typescript/comments/jl9ud4/pathing_packages_in_a_ts_monorepo/)
- url: https://www.reddit.com/r/typescript/comments/jl9ud4/pathing_packages_in_a_ts_monorepo/
---
Anyone have experience with pathing packages in monorepos?  I've got a monorepo with several packages, including a component library and a dashboard. I'm trying to import the component library in my dashboard package as import { Button } from '@company/ui' but it doesnt give me type definitions on my Button component unless I import from @company/ui/lib

I of course can always just import from /lib, but I feel like I'm doing something wrong here.
## [7][Justified use of an exclamation mark](https://www.reddit.com/r/typescript/comments/jl87jc/justified_use_of_an_exclamation_mark/)
- url: https://www.reddit.com/r/typescript/comments/jl87jc/justified_use_of_an_exclamation_mark/
---
Hey! I just found a justified use case for an exclamation mark:

    function parse(tokens: Token[]) {
        while (tokens.length) {
            const token = tokens.shift()!;
            ...
        }
    }

Since we check `tokens.length` as part of the while loop, we can be sure `tokens.shift()` returns a token.

So far, I haven't found many other cases of exclamation mark that I could justify. Have you found any that make sense?
## [8][Share business logic between projects](https://www.reddit.com/r/typescript/comments/jkx8xr/share_business_logic_between_projects/)
- url: https://www.reddit.com/r/typescript/comments/jkx8xr/share_business_logic_between_projects/
---
Hello everyone, 

this is a problem I'm trying to solve since months and I've still not found any satisfactory solution.

&amp;#x200B;

My company has been using Firebase Cloud functions for most of its activities, but we switched part of our infrastructure in a Google App Engine Container. Similar technologies, different benefits, make sense.

&amp;#x200B;

The problem I'm facing is that I don't really know how to share the common logic between projects. Here's some info.  


\- I'm using TS in both my projects  
\- I'd like to keep these projects private  
\- The two projects just differ in the controller part, nothing else  
\- I've setup a light CI with one of the two projects. Once I push a commit on master branch, Google Cloud Build download the repo, build it and deploy it in the App Engine container.   


How should I organize my projects? I'm not willing to stay one week more with so much duplicated business logic, but I'm not able to see a solution to my problem YET.  


Thank you
## [9][Property '' is missing in type '{}' but required in type ''.](https://www.reddit.com/r/typescript/comments/jkxrjq/property_is_missing_in_type_but_required_in_type/)
- url: https://www.reddit.com/r/typescript/comments/jkxrjq/property_is_missing_in_type_but_required_in_type/
---
On the bold code I'm getting this error after setting the first property as mandatory.

Any idea how to fix this?

Property 'categId' is missing in type '{}' but required in type 'Produit'. 

export class Produit {  
  categId: string;  
  id?: string;  
  nom?: string;  
  image?: string;  
  prix?: string;  
  description?: string;  
 constructor(**args: Produit = {}**) {  
 *this*.categId = args.categId;  
 *this*.id = args.id;  
 *this*.nom = args.nom;  
 *this*.image = args.image;  
 *this*.prix = args.prix;  
 *this*.description = args.description;  
  }  
}
## [10][What's the point of tuples? Why would you ever use a tuple when an object works just fine?](https://www.reddit.com/r/typescript/comments/jkpz5l/whats_the_point_of_tuples_why_would_you_ever_use/)
- url: https://www.reddit.com/r/typescript/comments/jkpz5l/whats_the_point_of_tuples_why_would_you_ever_use/
---
I just don't get the point of tuples. If anything, they just make it more confusing for your coworkers because they may try to add different data types to the array, and they'll be confused as hell when an error is thrown.
## [11][Starting a new library, what are the most modern settings i can use for node/browser?](https://www.reddit.com/r/typescript/comments/jkom4a/starting_a_new_library_what_are_the_most_modern/)
- url: https://www.reddit.com/r/typescript/comments/jkom4a/starting_a_new_library_what_are_the_most_modern/
---
... I want to start a new library.  Say I want to build to run in Node 15, and the most modern of browsers.  What are the best settings to use in my tsconfig?  I want zero backwards compatibility, and I want the output from TSC to run on both.  What can I get away with right now?
