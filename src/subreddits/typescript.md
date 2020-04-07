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
## [2][Is it possible to type this in TypeScript?](https://www.reddit.com/r/typescript/comments/fwi9h8/is_it_possible_to_type_this_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/fwi9h8/is_it_possible_to_type_this_in_typescript/
---
Hey everyone, having troubles typing something in TypeScript, I understand the error more or less but I don't know how I could do it differently.

Playground link:

https://www.typescriptlang.org/play/?ssl=26&amp;ssc=38&amp;pln=1&amp;pc=1#code/JYWwDg9gTgLgBAbzgVwM4FMDKMCGN1wC+cAZlBCHAORTo4DGMVA3AFCsnIB2jwEXKDAGEKYADboAHtjzoAPABU4U-FwAmqOACV09aGrmoYUYFwDmAGjg4uATwB89gBQBKRKzhw9XI3ADaRrJWGDAy+AC6cAC8gli4+IrOCIQubB5e-L4h+FDRcE4A1ui2AFxwRbYQJHAKVgBuOGLI6GU2tm5R9u6eniFh6E5OgfgdXU5IAHRTw+hWfhXhZQ1NBCmp6Z60MMhQAjNsnoRpnt6+ZugwOXmFxWUVVTWj3T1bO3vx6PPF4QdE7JsXN6IOAhMrZdBQKznGBlaFXI6sQjsUw5EgMAgABQhqH4zxIwCgRgAcjgQC0QcZTGZfmIcMTSeSjCZzGwkaxTvAwNj+P08mh0CJwBJpB85FjCfxnOtWFyJVx+hNoU4qPjCTASWSqOtZTj5R8JiFlbT6ZqrFRMCBgDAABZa5hAA

Code: 

    import { useState } from 'react';

    function useComplexState&lt;T extends Record&lt;string, any&gt;&gt;() {
      const [state, setState] = useState&lt;T&gt;({});

      const setter = (key: keyof T, value: any) =&gt; {
        setState((state) =&gt; state != null ? ({ ...state, [key]: value }) : state);
        return state;
      };

      const getter = (key: keyof T) =&gt; {
        return state[key];
      }

      return { set: setter, get: getter };
    }

    interface Person {
      firstName: string;
      lastName: string;
    }

    const personState = useComplexState&lt;Person&gt;();

    personState.get('firstName');
    personState.set('lastName', 'Smith');

Basically, I want:
 
 * `T` to always be an object `{}` so no one should be able to do `useComplexState&lt;boolean&gt;`
 * Ideally, `T` shouldn't be required, so if I use `useComplexState()` I should just have a generic object Record&lt;string, any&gt; inside.
 * I'd like to have a default value of `{}` for the `state`.
## [3][Help with TestBed](https://www.reddit.com/r/typescript/comments/fwgu9e/help_with_testbed/)
- url: https://www.reddit.com/r/typescript/comments/fwgu9e/help_with_testbed/
---
Hi all.    

I'm trying to set up a test for my service using TestBed. That service has some dependencies of other services. I've created .stub files for those dependencies and exported them. Problem is when I try to use these stubs in the test it still seems like the test methods use the real service and not the test. Here's my setup (the fake names are just for this thread):    

    describe('ServiceToTest', () =&gt; {
        let serviceDependency1Stub: ServiceDependency1;
    
        beforeEach(() =&gt; {
            TestBed.configureTestingModule({
                imports: [],
                providers: [
                    ServiceToTest,
                    { provide: ServiceDependency1, useClass: ServiceDependency1Stub },
                    { provide: ServiceDependency2, useClass: ServiceDependency2Stub }
                ]
            });
    
            serviceDependency1Stub = TestBed.inject(ServiceDependency1);
        });
    
        it('description', fakeAsync(() =&gt; {
            // Error appears here when trying to use the BehaviorSubject on the mock:
            // Property 'behaviorSubjectTest' does not exist on type                 
           'ServiceDependency1'
            serviceDependency1Stub.behaviorSubjectTest.next(new Test());
            ... more logic
        }));
    });
    

I understand the error, that's because behaviorSubjectTest property is on the stub and not the real service. But why doesn't it use the stub service when I listed it as a useClass in the providers?

What am I doing wrong? When I see what's available on serviceDependency1Stub I see all the methods and properties on the real service and not the stub.
## [4][Hey, I made a super small super type safe library to sort any type of array](https://www.reddit.com/r/typescript/comments/fwfwli/hey_i_made_a_super_small_super_type_safe_library/)
- url: https://www.reddit.com/r/typescript/comments/fwfwli/hey_i_made_a_super_small_super_type_safe_library/
---
check it out  
[sort-es](https://sort-es.netlify.com/) : Blazing fast, tree-shakeable, type-safe, modern utility library to sort any type of array
## [5][what is the difference between `Ts` errors vs `Eslint` errors when using `@typescript-eslint/parser` ?](https://www.reddit.com/r/typescript/comments/fwfe0x/what_is_the_difference_between_ts_errors_vs/)
- url: https://www.reddit.com/r/typescript/comments/fwfe0x/what_is_the_difference_between_ts_errors_vs/
---
the reason I'm asking this is, I wanna check for both type errors as well as eslint errors before commit my code, does running eslint covers all cases or should I run `tsc` or `typescript` too ?
## [6][Return type of promise + Try/Catch](https://www.reddit.com/r/typescript/comments/fwf9im/return_type_of_promise_trycatch/)
- url: https://www.reddit.com/r/typescript/comments/fwf9im/return_type_of_promise_trycatch/
---
Hello all,

I am getting the typical error of "A function whose declared type is neither void nor any must return a value" in the following code:  

```
const getStuff= async (  
  { setFieldValue, setFieldError, values: { stuff } }
): Promise&lt;AxiosResponse&lt;Stuff&gt;&gt; =&gt; {  
 try {  
 const res = await queryStuff(stuff);  
 if (res.status === 200) {  
 if (res.data.length &gt; 0) {  
        setFieldValue('stuff', res.data);  
      } else {  
        setFieldError('stuff', 'blabla error message');  
      }  
    }  
  } catch (err) {  
    setFieldError('stuff', 'blablabla error message');  
 if (err.response) {  
 // return error  
    } else if (err.request) {  
 // server did not respond  
    } else {  
 // something happened setting up the request  
    }  
  }  
};
```

QueryStuff calls axios and because getStuff is async then I must return a promise. The problem is that axios will throw an error upon an error so I am not sure how to refactor this. setFieldValue/setFieldError are Formik functions that return void.

How would you refactor the code to avoid this?

Thank you in advance.
## [7][Type checking a JSON file?](https://www.reddit.com/r/typescript/comments/fwgob0/type_checking_a_json_file/)
- url: https://www.reddit.com/r/typescript/comments/fwgob0/type_checking_a_json_file/
---
    export var tileList: hexes.TileList = {
      "3": {color: "yellow", stop: {value: 10}},
      "4": {color: "yellow", stop: {value: 10}},
      ...

My editor will happily do type checking inside this object within a .ts file.

However, if I take object and drop it in a JSON file, I lose the ability to type check it. Not just in realtime in my editor, but at all.

My workaround is to have a header file that contains `import blah blah; var tileList: hexes.TileList = `, then a script that concatenates that with the json file and pipes the result to tsc, but that seems so hacky-y.

Is there a better way, for one-time checking or for realtime checking in any editor or IDE?
## [8][Typescript - VSCode, object generator based on interface.](https://www.reddit.com/r/typescript/comments/fw85yx/typescript_vscode_object_generator_based_on/)
- url: https://www.reddit.com/r/typescript/comments/fw85yx/typescript_vscode_object_generator_based_on/
---
Hi folks, 

Using Typescript in VSCode, is there any extension to create an object using all properties of an interface. Example -

`interface IFields {`

`name: string;`

`age: number;`

`address: IAddress;`

`}`

&amp;#x200B;

`interface IAddress {`

`city: string;`

`zip: number;`

`}`

&amp;#x200B;

And then I can get something similar to -

const obj = {

  name: '',

  age: 0,

  address: {

city: '',

zip: 0

  }

};

&amp;#x200B;

Thanks!
## [9][How to structure a type where a particular field is occasionally not set?](https://www.reddit.com/r/typescript/comments/fw7e3n/how_to_structure_a_type_where_a_particular_field/)
- url: https://www.reddit.com/r/typescript/comments/fw7e3n/how_to_structure_a_type_where_a_particular_field/
---
I'm new to Typescript and this is tripping me up a bit. 

I've got a type for an object that is written to a database. Before it's written to the database, it does not have an ID, but anytime after that it does have an ID.

For simplicity say my type is something like:

    interface DatabaseObject {
        id: string,
        name: string
    }

When my Typescript backend is initially creating that object, I don't have an id string to set, which makes the type invalid. However, if I make `id` optional in the type, then I have to constantly check if `id` is set in the other 99% of my code, so I'd prefer to keep id non-optional.   


What's the best way to handle this?
## [10][Common JSON patterns in Haskell, Rust and TypeScript](https://www.reddit.com/r/typescript/comments/fvwm3o/common_json_patterns_in_haskell_rust_and/)
- url: https://codetalk.io/posts/2020-04-05-common-json-patterns-in-haskell-rust-and-javascript.html
---

## [11][CSS in TS?](https://www.reddit.com/r/typescript/comments/fw4ydy/css_in_ts/)
- url: https://www.reddit.com/r/typescript/comments/fw4ydy/css_in_ts/
---
What are you using for CSS in TypeScript? Only good library that I've found for type-safe CSS is Typestyle, so curious about what all of you are using. Also, is there a way to have type-safe CSS in other CSS in JS libraries like Emotion or Styled Components?
