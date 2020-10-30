# typescript
## [1][Who's hiring Typescript developers - October](https://www.reddit.com/r/typescript/comments/j2xtzq/whos_hiring_typescript_developers_october/)
- url: https://www.reddit.com/r/typescript/comments/j2xtzq/whos_hiring_typescript_developers_october/
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
## [2][What's the point of tuples? Why would you ever use a tuple when an object works just fine?](https://www.reddit.com/r/typescript/comments/jkpz5l/whats_the_point_of_tuples_why_would_you_ever_use/)
- url: https://www.reddit.com/r/typescript/comments/jkpz5l/whats_the_point_of_tuples_why_would_you_ever_use/
---
I just don't get the point of tuples. If anything, they just make it more confusing for your coworkers because they may try to add different data types to the array, and they'll be confused as hell when an error is thrown.
## [3][Solid TypeScript + ExpressJS Tutorial](https://www.reddit.com/r/typescript/comments/jkjze0/solid_typescript_expressjs_tutorial/)
- url: https://www.reddit.com/r/typescript/comments/jkjze0/solid_typescript_expressjs_tutorial/
---
Hey all, 

Getting into Typescript for a potential new job and was wondering if anyone has any solid tutorial recommendations? I've had a quick play around with TypeScript and read to docs. Looking for something that includes 

* Typescript 
* ExpressJs
* Postgres (or any DB really) 

Bonus:

* TypeORM
* Docker
## [4][Starting a new library, what are the most modern settings i can use for node/browser?](https://www.reddit.com/r/typescript/comments/jkom4a/starting_a_new_library_what_are_the_most_modern/)
- url: https://www.reddit.com/r/typescript/comments/jkom4a/starting_a_new_library_what_are_the_most_modern/
---
... I want to start a new library.  Say I want to build to run in Node 15, and the most modern of browsers.  What are the best settings to use in my tsconfig?  I want zero backwards compatibility, and I want the output from TSC to run on both.  What can I get away with right now?
## [5][Programmatically create declaration file](https://www.reddit.com/r/typescript/comments/jklm3e/programmatically_create_declaration_file/)
- url: https://www.reddit.com/r/typescript/comments/jklm3e/programmatically_create_declaration_file/
---
I know you read the title, and thought "sure, just use the `tsc file.js --declaration --emitDeclarationOnly --outFile file.d.ts` command". But thats not what I'm looking for.
I'm looking to parse the output of an API that describes exactly what fields a particular record could have, and use that to create a declaration file of interfaces and stuff. I can make elemental interfaces by hand, and have my code use those interfaces to interpret the API's response on the field types, and output something Typescript can recognize.

I want to do something like below. Is there ANYTHING that can do something like this?

```js
  const fields = [ //Example of my API's explanation of field types.
  {
    fieldType: 'string',
    fieldName: 'name',
  },{
    fieldType: 'number',
    fieldName: 'age',
  },{
    fieldType: 'multiselect',
    fieldName: 'gender',
    values: ['Male, Female']
  } 
]

// the code I would write with this imaginary library...
import {typeMaker} from 'type-maker' // doesn't exist

recordInterface = new typeMaker.Interface();
fields.forEach(f =&gt; {
  if (['string', 'number', 'date'].includes(f.fieldType)) { //like if its a built in type
    recordInterface.addProperty({ name: f.fieldName, type: f.fieldType })
  } else if (f.fieldType === 'multiselect') { //custom logic for custom type
    recordInterface.addProperty({ name: f.fieldName, type: f.values })
  }
}

//then just write it to afile.
recordInterface.writeToDeclarationFile('./types.d.ts')

// The contents of the types.d.ts file would look like this
interface Record {
  name: string;
  age: number;
  gender: 'Male' | 'Female';
}
```
## [6][Type definition for a function that returns Promise&lt;T&gt; but only when it's given a promise?](https://www.reddit.com/r/typescript/comments/jkcvld/type_definition_for_a_function_that_returns/)
- url: https://www.reddit.com/r/typescript/comments/jkcvld/type_definition_for_a_function_that_returns/
---
I'm writing a timing library so I can add a tracer to detect latency in our app.

I'd like to have it return a function such that if it's given () =&gt; T 

then it will return T 

... but if it's given Promise&lt;T&gt; it will return Promise&lt;T&gt;

so that the caller than would have to await the result from the tracer.

In the past I've done just traceSync and traceAsync but it seem silly to have two methods.

Internally I'm just testing if the value from the delegate is a Promise and then I then() it... 

I can't figure out how to define the types!  Any advice here?

The naive impl returns T | Promise&lt;T&gt; which isn't what I want...
## [7][Tracking down specific TypeScript source code files/function that is causing high CPU usage for TypeScript IDE/editor language services](https://www.reddit.com/r/typescript/comments/jk7avy/tracking_down_specific_typescript_source_code/)
- url: https://www.reddit.com/r/typescript/comments/jk7avy/tracking_down_specific_typescript_source_code/
---
I've got a TypeScript/Node project where I'm getting very high CPU usage while the project is open in vscode or phpstorm... quite often even when I'm not editing any files and the project is just open in a window in the background.  

My fans have been whirring up and down heaps over the last couple of months.  They're cranking right now while I've been typing this up, as they are most of the time lately.

It's just the one project with the problem, my other typescript project don't seem to have it - so it's not due to low end hardware (I've got a i7 7700 CPU, 32GB RAM and all files are on SSDs).

I have a few clues on which files it might be, but there's at least 10 of them, and they all contain a lot of code.  Some of it is big code-generated discriminated unions, which I think is a likely cause... but not sure which one(s).  I'm been programming in this style for longer than the CPU issue has been happening, but maybe I'm going too far with it lately given that typescript uses duck typing.

Are there any methods of finding out which source code files / functions might be causing the high CPU usage to the typescript service running in either vscode/phpstorm?  Or just by running `tsc` or some other tool outside of any editor?

I'm not sure exactly what it might involve... I'm hoping for maybe some kind of benchmarking/reporting thing that could point to specific `.ts` files, and ideally lines of code or type/function names that might be hogging the typescript service.

Last resort is to just start removing code by trial and error, but I have a feeling that approach might take me weeks, so I'm hopeful that maybe there's some better way.
## [8][VS Code extension for organizing classes/interfaces](https://www.reddit.com/r/typescript/comments/jk9vnq/vs_code_extension_for_organizing_classesinterfaces/)
- url: https://www.reddit.com/r/typescript/comments/jk9vnq/vs_code_extension_for_organizing_classesinterfaces/
---
I use linters to help expose the ordering of private/public instances, endpoints, etc. But I still find myself spending more time than I need just getting classes organized and structured well. Are there any extensions you have used that help auto-format in terms of organization?

Thanks!
## [9][I'm having trouble understanding how to polyfill nodejs dependencies with webpack 5. Are there any resources other than the docs that can help me understand this whole 'alias' thing?](https://www.reddit.com/r/typescript/comments/jjr8l0/im_having_trouble_understanding_how_to_polyfill/)
- url: https://www.reddit.com/r/typescript/comments/jjr8l0/im_having_trouble_understanding_how_to_polyfill/
---

## [10][What is late binding?](https://www.reddit.com/r/typescript/comments/jjm0kj/what_is_late_binding/)
- url: https://www.reddit.com/r/typescript/comments/jjm0kj/what_is_late_binding/
---
I only know TS and no other language. Does it make sense to talk about late binding in TS and if yes what is late binding?
## [11][const _ = this;](https://www.reddit.com/r/typescript/comments/jj0zlr/const_this/)
- url: https://www.reddit.com/r/typescript/comments/jj0zlr/const_this/
---
.Net background and I need to know why this code exists in some libraries. What good does it do?

const \_ = this;  
\_.localVariable = &lt;some arbitrary value&gt;;

Would it not be the same as just using this.localVariable?
