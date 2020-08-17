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
## [2][Dictionary object does not give undefined as possible type (even with strict null checks).](https://www.reddit.com/r/typescript/comments/ibbl0b/dictionary_object_does_not_give_undefined_as/)
- url: https://www.reddit.com/r/typescript/comments/ibbl0b/dictionary_object_does_not_give_undefined_as/
---
[Playground link](https://www.typescriptlang.org/play?#code/MYewdgzgLgBAJgS2FB4CGAnAnjAXDAbwG0APXaDBMAcwF08YwBXAWwCMBTDAXxgF5C3ANwAoEaEiwS-eEhTpsAOhJCA9KoAWIAG5cYOvdIA0MBFAgwIWgO4WoWAA4cGzdnrRg4jELFecMMAA+MEyeHABmVBxwQA) . Is this the expected behavior ?
## [3][TypeScript-first middleware for AWS Lambda functions](https://www.reddit.com/r/typescript/comments/ib1ju9/typescriptfirst_middleware_for_aws_lambda/)
- url: https://dbartholomae.github.io/lambda-middleware/
---

## [4][TypeScript Exercises: Now interactive and solutions are included!](https://www.reddit.com/r/typescript/comments/ib2g23/typescript_exercises_now_interactive_and/)
- url: https://www.reddit.com/r/typescript/comments/ib2g23/typescript_exercises_now_interactive_and/
---
I've re-created the TypeScript exercises and now they are interactive and fully accessible through a web-browser. Check out and have fun: [https://typescript-exercises.github.io/](https://typescript-exercises.github.io/)
## [5][Building a game with TypeScript. Drawing Grid 1/5](https://www.reddit.com/r/typescript/comments/iaxv3w/building_a_game_with_typescript_drawing_grid_15/)
- url: https://medium.com/@gregsolo/building-a-game-with-typescript-drawing-grid-1-5-aaf68797a0bb?sk=9663603b5086a6bbd30147fd26a4f457
---

## [6][Pwa](https://www.reddit.com/r/typescript/comments/ibahbh/pwa/)
- url: https://www.reddit.com/r/typescript/comments/ibahbh/pwa/
---
Hi is it possible to create Progressive web app (react - typescript) that you could block on computer, I just want it to work on mobile devices.
## [7][How do I publish a library with multiple index.d.ts?](https://www.reddit.com/r/typescript/comments/iarrr4/how_do_i_publish_a_library_with_multiple_indexdts/)
- url: https://www.reddit.com/r/typescript/comments/iarrr4/how_do_i_publish_a_library_with_multiple_indexdts/
---
So I want to publish a components library where the consumer can use components from the main index or a sub-folder where legacy components are stored.

`import Button from "my-library"`

`import Button from "my-library/legacy"`

In tsconfig.json, I have it generate the types:

        "declaration": true /* Generates corresponding '.d.ts' file. */,
        "declarationMap": true /* Generates a sourcemap for each corresponding '.d.ts' file. */,
        "declarationDir": "typing",

So it generates a typings folder with 2 index.d.ts files, one in root and one in \`legacy\` folder.  


My question is I could only define one entry point in the \`types\` field in package.json. How do I include the other typing file in the \`legacy\` subfolder? Is this where I use a triple slash directive?   


Thanks!
## [8][Need help getting a libraries typescript working](https://www.reddit.com/r/typescript/comments/ib19o6/need_help_getting_a_libraries_typescript_working/)
- url: https://www.reddit.com/r/typescript/comments/ib19o6/need_help_getting_a_libraries_typescript_working/
---
Trying to use a library called forerunnerDB. Pretty neat library, but its typescript declarations are lacking a bit. While I've gotten some familiarity with typescript itself, I'm having trouble getting their declaration working.  


[package.json](https://github.com/webtorrent/webtorrent/blob/master/package.json), it's missing the types declaration in there which is needed because they put their type declaration file in another folder.

[Core.js](https://github.com/Irrelon/ForerunnerDB/blob/master/js/lib/Core.js), the start point of the library. This is what is called when you do new forerunnerdb  


[types.d.ts](https://github.com/Irrelon/ForerunnerDB/blob/master/typescript/types.d.ts), their types file.  


Now I'm pretty confused as to how to get the types.d.ts declaration to relate Core in the types declaration to the core in Core.js. I'm further confused since the types in the file provided here are classes rather than interfaces.  


So, would anyone mind giving me some insights and hints as to what I can do to at least get this file to work as it's supposed to and maybe some hints as to how I can help them get it put together?
## [9][How to use ReturnType on a 2 level deep returned object?](https://www.reddit.com/r/typescript/comments/iasxn7/how_to_use_returntype_on_a_2_level_deep_returned/)
- url: https://www.reddit.com/r/typescript/comments/iasxn7/how_to_use_returntype_on_a_2_level_deep_returned/
---
The `pg-promise` library is kind of odd in that you first invoke it with an options object, and then invoke it again with a config object. The first and second invokation contain different properties  (and there are also static methods on the uninvoked import)

This is what I have now:

      private pgPromiseOptions: ReturnType&lt;typeof PGPromise&gt;;
    
      constructor(pgPromiseConfigured: any) {
        this.pgPromiseOptions = PGPromise({
          capSQL: true,
        });
    
        this.pgPromiseConfigured = pgPromiseConfigured;
      }

All of the following type annotations have failed. What is the correct way to do it?

    private pgPromiseConfigured: ReturnType&lt;typeof pgPromiseOptions&gt;;
    constructor(private pgPromiseConfigured: ReturnType&lt;typeof pgPromiseOptions&gt;
    
    private pgPromiseConfigured: ReturnType&lt;typeof PGPromise()&gt;;
    constructor(private pgPromiseConfigured: ReturnType&lt;typeof PGPromise()&gt;
## [10][How I rewrite these Go functions with TypeScript?](https://www.reddit.com/r/typescript/comments/iat063/how_i_rewrite_these_go_functions_with_typescript/)
- url: https://www.reddit.com/r/typescript/comments/iat063/how_i_rewrite_these_go_functions_with_typescript/
---
I'm trying to do this [tutorial](https://flaviocopes.com/go-git-contributions/) but I want to write with TypeScript! I like Go, however I'm studying more about Node. However, I'm having difficult with Date operations, could anyone help me?

&amp;#x200B;

    // getBeginningOfDay given a time.Time calculates the start time of that day 
    
    func getBeginningOfDay(t time.Time) time.Time { 
    
    	year, month, day := [t.Date](https://t.Date)() 
    
    	startOfDay := [time.Date](https://time.Date)(year, month, day, 0, 0, 0, 0, t.Location()) 
    
    	return startOfDay 
    
    } 
    
    // countDaysSinceDate counts how many days passed since the passed date 
    func countDaysSinceDate(date time.Time) int { 
    	days := 0 
    	now := getBeginningOfDay(time.Now()) 
    
    	for date.Before(now) { 
    		date = date.Add(time.Hour * 24) 
    		days++ 
    
    		if days &gt; daysInLastSixMonths { 
    			return outOfRange 
    		} 
    	} 
    	return days 
    }
    
    // printMonths prints the month names in the first line, determining when the month 
    // changed between switching weeks 
    func printMonths() { 
    	week := getBeginningOfDay(time.Now()).Add(-(daysInLastSixMonths * time.Hour * 24)) 
    	month := week.Month() fmt.Printf(" ") 
    	for { 
    		if week.Month() != month { 
    			fmt.Printf("%s ", week.Month().String()[:3]) 
    			month = week.Month() 
    		} else { 
    			fmt.Printf(" ") 
    		} 
    		week = week.Add(7 * time.Hour * 24) 
    		if week.After(time.Now()) { 
    			break 
    		} 
    	} 
    	fmt.Printf("\n") 
    }

&amp;#x200B;
## [11][How do you write shared libraries?](https://www.reddit.com/r/typescript/comments/iapnkw/how_do_you_write_shared_libraries/)
- url: https://www.reddit.com/r/typescript/comments/iapnkw/how_do_you_write_shared_libraries/
---
When I write a library for use in a browser I use `"lib": ["DOM"]` in my tsconfig, and for node I install `@types/node`, but what's the best practice for libraries that run on both?

I need to use types that exist in both of these, but not in the non-DOM lib files.
Is there a project that allows me to get the intersection of these types or do I just have to compile twice to check I'm not using a platform-specific API?
