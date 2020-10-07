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
## [2][How to use jQuery Plugins like owl-carousel in typeScript?](https://www.reddit.com/r/typescript/comments/j6npge/how_to_use_jquery_plugins_like_owlcarousel_in/)
- url: https://www.reddit.com/r/typescript/comments/j6npge/how_to_use_jquery_plugins_like_owlcarousel_in/
---
Hi!  
I am new to typeScript. I can't figure out how to install the owl-carousel in my typeScript project.

Any help regarding this will be appreciated! Thanks in advance.
## [3][Building a Game With TypeScript. New episode! Colors and Layers](https://www.reddit.com/r/typescript/comments/j6858v/building_a_game_with_typescript_new_episode/)
- url: https://medium.com/@gregsolo/building-a-game-with-typescript-colors-and-layers-337b0e4d71f?source=friends_link&amp;sk=45ea8fe1af4d87bbd4b1b12b43c4557b
---

## [4][Is "keyof typeof" the correct way to assert an enum's keys?](https://www.reddit.com/r/typescript/comments/j6e1my/is_keyof_typeof_the_correct_way_to_assert_an/)
- url: https://www.reddit.com/r/typescript/comments/j6e1my/is_keyof_typeof_the_correct_way_to_assert_an/
---
I was trying to use an enum's keys to make a selection on it:

    enum myEnum {
      a = "a value"
      b = "b value"
    }
    
    const selection = myEnum[someVar as keyof typeof myEnum];

Is this the correct usage? I tried to only use `keyof`, but a whole bunch of generic object keys came back in a lint error. 

The above passes lint but feels wrong to me, I'm not sure exactly what the keys in `typeof myEnum` are but I doubt it is just `a` and `b`.
## [5][Why is my higher-order factory function not throwing a type error, and how can I make it work?](https://www.reddit.com/r/typescript/comments/j64mkf/why_is_my_higherorder_factory_function_not/)
- url: https://www.reddit.com/r/typescript/comments/j64mkf/why_is_my_higherorder_factory_function_not/
---
`heroFactory([10, 10])` works, but I would also like to return many heroes into an array.

`const heros = createMultipleInstantiatedClasses(3, heroFactory);` should not typecheck as it should accept only a function with an argument `factory: (position: Vector2) =&gt; T` but it does, and I get a load of `undefined` positions as a result. What is wrong with my type signature?

If I change to `const heros = createMultipleInstantiatedClasses(3, heroFactory([10, 10]));` then I get ``&lt;Hero&gt;` is not a function`` error.

Can anyone help? Thanks.

https://www.typescriptlang.org/play?#code/KYDwDg9gTgLgBDAnmYcBqwDGNoCY4C8cA2gHYCuAtgEbBQA0cFNdAugNwBQnmANgIYBnQXAASdCHADenOHEiCAljEURSXOZjWCYUctmgAKBctWkAXOiw4ouAJTTZchAAtFggHQmVawvIhKPupOAL6cYZygkLBwAGbkpNhmcJhQwPwwwACy5LwqYLzAAJKkOvykKhnAACYAwgLCwIIAPAAqAHyGTsy0UADysfVCgk2WPXT0TrH8BlCIlsYBpmqWGLP2hO1wrZx2lq3ErI5yaTDkUKRwAIJQUPyIHrFQEJSGUnCFpADmMC5jVL0BkNGiIQoxprNEHYuBEotB4PFEkE4C4JAAxGY2RCLQJmVbWPB7MQSY5wU7nS6kYAAd2Jzxxy1I0PC3C0pXgqOefk5EAxkMMxAAjAAGRgi1jQuAAeilcD6AGkeNoORIRERUulMjk8ooCsV2eVKpk6g0RoJDABmRg8vlY5lK0oQQoeXgQL6GHmSmV0yTvbx4khwEVi4VwI4hQAQBA7BE7gC63R7VV7ZcQfdJ-LiVnAEtVgLFFFTqnAwRHxFy-UsgpYc3mCzVi-RSySK5mLNnSLn84Xi2Go0A

```typescript
export type Vector2 = [number, number];

class Hero {
  position;
  constructor(position: Vector2) {
    this.position = position;
  }
}

export function createMultipleInstantiatedClasses&lt;T&gt;(
  numberOfClasses: number,
  factory: (position: Vector2) =&gt; T
): T[] {
  return Array.from({ length: numberOfClasses }, factory);
}

export function heroFactory(position: Vector2): Hero {
  return new Hero(position);
}

const hero = heroFactory([10, 10]); // OK
const heros = createMultipleInstantiatedClasses(3, heroFactory);

console.log(hero); // Hero { position: [ 10, 10 ] } 
console.log(heros); // [ Hero { position: undefined }, Hero { position: undefined }, Hero { position: undefined } ] 
```
## [6][Express REST APIs in TypeScript?](https://www.reddit.com/r/typescript/comments/j5vdz8/express_rest_apis_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/j5vdz8/express_rest_apis_in_typescript/
---

Is there a cleaner way to do this like C# classes and methods?

```csharp

public class RestAPI {
   [HTTPGet]
   [Authorize]
   public static string GetFoo() {
   	return new string("Welcome to Foo!")
   }
}
```

Express.js routes look like Python req res and its fugly.
## [7][Change types using TS Compiler API](https://www.reddit.com/r/typescript/comments/j5uodl/change_types_using_ts_compiler_api/)
- url: https://www.reddit.com/r/typescript/comments/j5uodl/change_types_using_ts_compiler_api/
---
Hi everyone,

I am using the TS Compiler in a project and I need to modify the inference of TS a bit. For example, if I have:

    let a = 2
    a = 'name'

the type nodes of both **a** identifiers are **numbers**. Of course this is expected in TS.

However, is there any way for the compiler to actually say that **a** is **number | string** ? I saw *typescript.d.ts* and it is possible to do some hacks (or maybe do my own separate code analysis), but if there is any TS library or codefix that does that (and other types of inference on top of the inference done by TS) I think it would be better.
## [8][Creating an express router in TypeScript](https://www.reddit.com/r/typescript/comments/j5kcvc/creating_an_express_router_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/j5kcvc/creating_an_express_router_in_typescript/
---
Trying to build a clean application for a personal project using SOLID design principles.

I was wondering how do you guys build a router in TypeScript? Do you use classes to initialise a router class?
## [9][How to get notified about finished recursive http calls with typescript?](https://www.reddit.com/r/typescript/comments/j5ri9g/how_to_get_notified_about_finished_recursive_http/)
- url: https://www.reddit.com/r/typescript/comments/j5ri9g/how_to_get_notified_about_finished_recursive_http/
---
I am trying to recursively fetch Objects from a Server. An Object can have multiple objects as children and it is not known how many children there are and how deep the objects are nested.

I have a function that fetches the data from the sever and creates typescript objects with it and adds the children to the according parent. My problem is, that I'm subscribing to the result in the function, so I never know when I am finished with fetching the objects from the server. Where do I have to subscribe to my result? And how can I still create my Typescript objects from the data when I don't subscribe in my fetch function?

This is what I currently have:

    fetchDataFromServer(token: HttpHeaders, url: string, parent: OwnObject) {
    
        let current: OwnObject;
        return this.http    //subscribe should be done somwhere else but how can I get notified about the subscribe only if the recursion is completely finished
          .get&lt;any&gt;(url, { headers: token }).subscribe(result =&gt; {
            if (Array.isArray(result)) {
              for (const resultChild of result) {
                if (!this.isOfTypeToShow(resultChild)) continue;
                else {
                  current = new OwnObject(resultChild.name, resultChild.type, resultChild.href, resultChild.id, []);
                  parent.childObject.push(current);
                  if (!resultChild.type.endsWith('type that has no children')) { //this type cannot have children so no further search is necessary
                    const urlOfChildren = this.createURLStringOfChildren(resultChild.relations, current);
                    if (urlOfChildren !== '') {
                      return this.fetchDataFromServer(token, urlOfChildren, current);
                    }
                  }
                }
              }
            } else {
              if (this.isOfTypeToShow(result)) {
                current = new OwnObject(result.name, result.type, result.href, result.id, []); 
                parent !== null ? parent.childObject.push(current) : this.rootObject.next(current);
                if (!result.type.endsWith('type that has no children')) {
                  const urlOfChildren = this.createURLStringOfChildren(result.relations, current);
                  if (urlOfChildren !== '') {
                    return this.fetchDataFromServer(token, urlOfChildren, current);
                  }
                }
              }
            }
          }));
      }

Sorry for the bad  formatting and maybe noobish question, I am using angular btw and would be very happy if you could help me a little bit

&amp;#x200B;

&amp;#x200B;

The method createURLStringOfChildren just creates a url string according to the children of the current element. With this string I fetch the children from the server.

The method isOfTypeToShow just returns true if the type of the element is interesting in the context (some types can be ignored)

&amp;#x200B;

I have also tried to use expand according to this example ([https://stackoverflow.com/questions/40529232/angular-2-http-observables-and-recursive-requests](https://stackoverflow.com/questions/40529232/angular-2-http-observables-and-recursive-requests)) example but it did not really help me with my problem. I have only limited knowledge about request calls in Typescript/Javascript.

So basically I just need to know how I can do my operations in the function without subscribing to it and how I can get notified when the recursive http-call is completely finished.

The function gets called from another component in the first place.
## [10][What problem would there be with running an Express API server using ts-node](https://www.reddit.com/r/typescript/comments/j5le0l/what_problem_would_there_be_with_running_an/)
- url: https://www.reddit.com/r/typescript/comments/j5le0l/what_problem_would_there_be_with_running_an/
---
I use ts-node-dev and the server seems to run fine locally. Is there a reason not to just skip compilation and run the server from source code during production?
## [11][Trying to understand this error message](https://www.reddit.com/r/typescript/comments/j5k6lo/trying_to_understand_this_error_message/)
- url: https://www.reddit.com/r/typescript/comments/j5k6lo/trying_to_understand_this_error_message/
---
I'm trying to assign a function to the onMouseDown property on a [styled component](https://github.com/styled-components/styled-components) but am getting an error message which I'm having an issue understanding the syntax of: `Type '(e: MouseEvent) =&gt; void' is not assignable to type '(event: MouseEvent&lt;HTMLDivElement, MouseEvent&gt;) =&gt; void'`
Are `&lt;HTMLDivElement, MouseEvent&gt;` two generics in this case?
