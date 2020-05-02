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
## [2][I reverse engineered a npm registry from scratch. It is written entirely in TypeScript and has a low amount of dependencies.](https://www.reddit.com/r/typescript/comments/gc3q61/i_reverse_engineered_a_npm_registry_from_scratch/)
- url: https://github.com/Tanuel/baggy
---

## [3][Using Typescript in a Javascript React app](https://www.reddit.com/r/typescript/comments/gc6h9o/using_typescript_in_a_javascript_react_app/)
- url: https://www.reddit.com/r/typescript/comments/gc6h9o/using_typescript_in_a_javascript_react_app/
---
I have a React app written entirely in Javascript. I want to start using Typescript but I don't want to go back and convert any of my old code to TS. The app is too large. I'd like to only use TS moving forward.

Are there any guides on how I can set this up? I use webpack.
## [4][Why am I getting this TS error?](https://www.reddit.com/r/typescript/comments/gbpqiy/why_am_i_getting_this_ts_error/)
- url: https://i.redd.it/un2hvbg4r7w41.png
---

## [5][Open Source Typescript Express projects](https://www.reddit.com/r/typescript/comments/gc4xe3/open_source_typescript_express_projects/)
- url: https://www.reddit.com/r/typescript/comments/gc4xe3/open_source_typescript_express_projects/
---
Hello, I've been learning Typescript recently and decided to try to implement what I've learning in a new project.

I would like to have a look at some major projects in Typescript Express so that I can catch some best practices, design patterns, etc.

I'd appreciate if anyone could share links to some open source projects/companies that used these technologies for their product.
## [6][The Jupiter (YC S19) Stack - From TypeScript to Kubernetes and back — how we write, build, and push code at Jupiter](https://www.reddit.com/r/typescript/comments/gbl21c/the_jupiter_yc_s19_stack_from_typescript_to/)
- url: https://starship.jupiter.co/jupiter-stack/
---

## [7][Dynamic import performance evaluation](https://www.reddit.com/r/typescript/comments/gbt5k3/dynamic_import_performance_evaluation/)
- url: https://www.reddit.com/r/typescript/comments/gbt5k3/dynamic_import_performance_evaluation/
---
Hello everyone.
In order to reduce cold start time of my serverless functions, I was thinking about using dynamic imports inside each function in order to avoid initialization of dependencies I may not use.
I know that once a module is normally imported, it is preserved in the global scope and I can use it without any overhead. 
Is the same for dynamic imports? 
I know that it could be a silly question but I feel like I'm missing a piece of the puzzle here.
## [8][Why are these functions not declared as private methods?](https://www.reddit.com/r/typescript/comments/gbify9/why_are_these_functions_not_declared_as_private/)
- url: https://www.reddit.com/r/typescript/comments/gbify9/why_are_these_functions_not_declared_as_private/
---
I'm looking at the source code for the Angular Material website ([https://github.com/angular/material.angular.io](https://github.com/angular/material.angular.io)).

In this example:

    import {Injectable} from '@angular/core';
    
    
    /**
     * Class for managing stylesheets. Stylesheets are loaded into named slots so that they can be
     * removed or changed later.
     */
    @Injectable()
    export class StyleManager {
      /**
       * Set the stylesheet with the specified key.
       */
      setStyle(key: string, href: string) {
        getLinkElementForKey(key).setAttribute('href', href);
      }
    
      /**
       * Remove the stylesheet with the specified key.
       */
      removeStyle(key: string) {
        const existingLinkElement = getExistingLinkElementByKey(key);
        if (existingLinkElement) {
          document.head.removeChild(existingLinkElement);
        }
      }
    }
    
    function getLinkElementForKey(key: string) {
      return getExistingLinkElementByKey(key) || createLinkElementWithKey(key);
    }
    
    function getExistingLinkElementByKey(key: string) {
      return document.head.querySelector(`link[rel="stylesheet"].${getClassNameForKey(key)}`);
    }
    
    function createLinkElementWithKey(key: string) {
      const linkEl = document.createElement('link');
      linkEl.setAttribute('rel', 'stylesheet');
      linkEl.classList.add(getClassNameForKey(key));
      document.head.appendChild(linkEl);
      return linkEl;
    }
    
    function getClassNameForKey(key: string) {
      return `style-manager-${key}`;
    }
    

The last 4 functions getLinkElementForKey, getExistingLinkElementByKey, createLinkElementWithKey and getClassNameForKey are defined outside of the class StyleManager. They aren't used anywhere else besides this class.

What is the reason for doing this over defining them as private method inside the class?

I can't test it out because the project doesn't start on Windows as it calls a bash script in yarn start.
## [9][Purify 0.15 released! - A Functional programming library for TypeScript](https://www.reddit.com/r/typescript/comments/gb0qhw/purify_015_released_a_functional_programming/)
- url: https://www.reddit.com/r/typescript/comments/gb0qhw/purify_015_released_a_functional_programming/
---
Changelog: [https://gigobyte.github.io/purify/changelog/0.15](https://gigobyte.github.io/purify/changelog/0.15)

**If you are not familiar with functional programming practices please take a look at the documentation, otherwise read below.**

I've finalized the API for async error handling, I feel like this might be best DX for this kind of thing (working with Either and Maybe inside Promises). I'd very much appreciate any feedback on how easy it is to use on the back-end, especially in complicated async flows.

I've already compared purify to fp-ts a [couple](https://www.reddit.com/r/functionalprogramming/comments/ebg4pc/purify_014_released_a_functional_programming/fb5uv16/) of [times](https://www.reddit.com/r/typescript/comments/8y20no/pure_010_released_a_functional_programming/e29e4yl/).
## [10][Uncovered lines in test when using `foo || []` - guidance from experienced devs needed](https://www.reddit.com/r/typescript/comments/gbd0fm/uncovered_lines_in_test_when_using_foo_guidance/)
- url: https://www.reddit.com/r/typescript/comments/gbd0fm/uncovered_lines_in_test_when_using_foo_guidance/
---
Dear Friends of TypeScript,

Even though I have quite some experience with other languages I'm fairly new to TS and don't get all the parts just yet. I just wrote a function that checks if the value of an entry in a map is an array with more than one element and would like to get some feedback, especially since when testing, one particular line of code is marked as 'uncovered'. My question is if there's any other way to write that function to make sure that there are no uncovered lines?

Here is the function:

    /**
     *
     * @param wordMap - map with sorted word as index and array of words as value
     */
    function filterAnagrams(wordMap: Map&lt;string, string[]&gt;): string[] {
      let anagrams: string[] = [];
      for (let key of wordMap.keys()) {
        let words: string[] = wordMap.get(key) || [];
        if (words.length &gt; 1) {
          anagrams.push(...words);
        }
      }
      return anagrams;
    }

Here is the test:

    describe("filterAnagrams", function () {
      it("should return the values from the map where elements &gt; 1", function () {
        let tempMap = new Map();
        tempMap.set("foo", ["foo1", "foo2"]);
        tempMap.set("bar", ["bar1"]);
        let filteredValues: string[] = filterAnagrams(tempMap);
        expect(filteredValues).toEqual(["foo1", "foo2"]);
      });
    
      it("should return empty array since map is empty", function () {
        let tempMap = new Map();
        let filteredValues: string[] = filterAnagrams(tempMap);
        expect(filteredValues).toEqual([]);
      });
    });

The empty array at the end of line 3 in the \`filterAnagrams\` function  is marked as uncovered by jest:

    let words: string[] = wordMap.get(key) || [];

If I remove the || \[\] part I receive the error

    Type 'string[] | undefined' is not assignable to type 'string[]'.
    Type 'undefined' is not assignable to type 'string[]'.ts(2322)
## [11][Using fetch with Typescript and the Todoist API](https://www.reddit.com/r/typescript/comments/gb66fp/using_fetch_with_typescript_and_the_todoist_api/)
- url: https://medium.com//using-fetch-with-typescript-and-the-todoist-api-5203c5177ed5?source=friends_link&amp;sk=a16444467bf3dcfca20b102972fe8b43
---

