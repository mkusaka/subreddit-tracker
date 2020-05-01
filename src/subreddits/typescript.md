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
## [2][Purify 0.15 released! - A Functional programming library for TypeScript](https://www.reddit.com/r/typescript/comments/gb0qhw/purify_015_released_a_functional_programming/)
- url: https://www.reddit.com/r/typescript/comments/gb0qhw/purify_015_released_a_functional_programming/
---
Changelog: [https://gigobyte.github.io/purify/changelog/0.15](https://gigobyte.github.io/purify/changelog/0.15)

**If you are not familiar with functional programming practices please take a look at the documentation, otherwise read below.**

I've finalized the API for async error handling, I feel like this might be best DX for this kind of thing (working with Either and Maybe inside Promises). I'd very much appreciate any feedback on how easy it is to use on the back-end, especially in complicated async flows.

I've already compared purify to fp-ts a [couple](https://www.reddit.com/r/functionalprogramming/comments/ebg4pc/purify_014_released_a_functional_programming/fb5uv16/) of [times](https://www.reddit.com/r/typescript/comments/8y20no/pure_010_released_a_functional_programming/e29e4yl/).
## [3][Uncovered lines in test when using `foo || []` - guidance from experienced devs needed](https://www.reddit.com/r/typescript/comments/gbd0fm/uncovered_lines_in_test_when_using_foo_guidance/)
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
## [4][Using fetch with Typescript and the Todoist API](https://www.reddit.com/r/typescript/comments/gb66fp/using_fetch_with_typescript_and_the_todoist_api/)
- url: https://medium.com//using-fetch-with-typescript-and-the-todoist-api-5203c5177ed5?source=friends_link&amp;sk=a16444467bf3dcfca20b102972fe8b43
---

## [5][Typescript generics 'lens' to constrain a type to selected fields?](https://www.reddit.com/r/typescript/comments/gb81fz/typescript_generics_lens_to_constrain_a_type_to/)
- url: https://www.reddit.com/r/typescript/comments/gb81fz/typescript_generics_lens_to_constrain_a_type_to/
---
... I don't think this is possible in typescript.  Maybe I'm wrong.

I have an API I'm trying to ship that I need a 'lens' to listen to a subscribe map and only update when specific fields have changed.

I want it to return a synthetic type contrained on the input.

the input API is simple. 

I can use fields: ReadonlyArray&lt;keyof Foo&gt;

... the problem I'm having is how to constrain the output to a map which only contains those input values.  

So if the map has keys name, address, city, state, zip, and the user specifies just city, state, and zip, I want a map with only city, state and zip.

I want name and address to be removed so that the caller doesn't attempt to read from name and address which could be stale.
## [6][Easiest way to return an interface/map with a collection of functions ?](https://www.reddit.com/r/typescript/comments/gaywkz/easiest_way_to_return_an_interfacemap_with_a/)
- url: https://www.reddit.com/r/typescript/comments/gaywkz/easiest_way_to_return_an_interfacemap_with_a/
---
I'm working on something in react and can't use classes as 'this' keeps being thrown away.  I want to avoid constantly having to avoid foo.bind(myInstance) too.

So instead I'm returning an interface with callbacks in it ... they're just functions.

so I'll have:

    interface MyInterface {
        myCallback1: () =&gt; void;
        myCallback2: () =&gt; void;

        ... etc
  }


and the way I'm returning this now is:

    function myCallback1() {
        // ... impl
    } 

    function myCallback2() {
        // ... impl
    } 

    return {
        myCallback1, myCallback2, etc...
    } 

... but this yields a ton of boilerplate.  Is there a cleaner way to return / organize everything better?  I was thinking namespaces but you can't return one as an object.  Maybe a static class? Not sure you can return one though... will have to play with that.

... EDIT.  Returning a static class seems to work but curious if that's the best strategy.

... EDIT2. It actually won't work because this is discarded and there's the need to call other functions. Without 'this' they're all in the same namespace.

.... EDIT3. It works if you give it a name like Foo then return Foo and only access functions by Foo and not 'this'
## [7][class-transformer as RxJS operator](https://www.reddit.com/r/typescript/comments/gaxnp9/classtransformer_as_rxjs_operator/)
- url: https://www.reddit.com/r/typescript/comments/gaxnp9/classtransformer_as_rxjs_operator/
---
I want to wrap `plainToClass` from `class-transformer` into an RxJS operator function so that I can use it in `pipe` to map a plain object to a class.

The original signature of `plainToClass` is:

    function plainToClass&lt;T, V&gt;(cls: ClassType&lt;T&gt;, plain: V[], options?: ClassTransformOptions): T[];
    function plainToClass&lt;T, V&gt;(cls: ClassType&lt;T&gt;, plain: V, options?: ClassTransformOptions): T;

My function would be:

    const toClass = cls =&gt; map(plain =&gt; plainToClass(cls, plain));

Simply, `plainToClass` takes a constructor and a literal object and uses that object to create an instance of your class. My function just wraps this in an RxJS `map` operator.

How should I write the type annotations for this, so that when `v` is a single object, it's return type will be a single instance, and when `v` is an array, the return type will be an array of instances?

E.g.:

    const toClass = &lt;T, V&gt;(cls: new() =&gt; T) =&gt; map&lt;V, T[] when V is Array else T&gt;((plain: V) =&gt; plainToClass(cls, plain));

&amp;#x200B;
## [8][Announcing TypeScript 3.9 RC](https://www.reddit.com/r/typescript/comments/gaecct/announcing_typescript_39_rc/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-3-9-rc/
---

## [9][A collection of challenging TypeScript exercises](https://www.reddit.com/r/typescript/comments/gabb46/a_collection_of_challenging_typescript_exercises/)
- url: https://github.com/mdevils/typescript-exercises
---

## [10][Restore mock of imported class function in ts-jest while testing](https://www.reddit.com/r/typescript/comments/gav2w0/restore_mock_of_imported_class_function_in_tsjest/)
- url: https://www.reddit.com/r/typescript/comments/gav2w0/restore_mock_of_imported_class_function_in_tsjest/
---
**Issue**

I am mocking a function (only one) which is imported from another class which is being called inside the main class. I am able to mock the function and return the values which I provide. But I am not able to restore the mocked function back to normal for the subsequent tests.

Any help would be appreciated!

Framework Using: jest + ts-jest

**Code**

\~main.ts

\`\`\`

import {SubClass} from './subclass.ts' 

export class MainClass { 

let sub: SubClass = new SubClass() 

public async Car(){ 

let start = await sub.key(); 

return start } }

\`\`\`

\~sub.ts

\`\`\`

export class SubClass{ 

public async key(){ return "you can start the car" } }

\`\`\`

\`\`\`  
import {SubClass} from './subclass.ts' 

import {MainClass} from './mainclass.ts' 

import {mocked} from 'ts-jest/utils'  

jest.mock(./subclass.ts) 

let main = new MainClass() 

let sub = new SubClass() 

let mockedFn = mocked(sub,true)  

mockedFn.key = jest.fn().mockImplementation(() =&gt; console.log('mocking succesfull'))  

afterEach(()=&gt;{ mockedFn.key.mockRestore() // tried mockClear(), mockReset() })

  it('test for main func mocked result',async ()=&gt;{ 

const result = await main.car()  

expect(result).toEqual("mocking succesfull") } 

it('test for main func result',async ()=&gt;{ 

const result = await main.car()  

expect(result).toEqual("you can start the car") // getting result as undefined 

}

\`\`\`
## [11][How can I have intellisense for a dictionary object but also type definition.](https://www.reddit.com/r/typescript/comments/gakqj5/how_can_i_have_intellisense_for_a_dictionary/)
- url: https://www.reddit.com/r/typescript/comments/gakqj5/how_can_i_have_intellisense_for_a_dictionary/
---
    /**
     * @type {{[x:string] : (type1) =&gt; type2}}
    */
    const obj;

This does not give me intellisense for the properties of the `obj`.
