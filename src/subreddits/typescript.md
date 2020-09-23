# typescript
## [1][Who's hiring Typescript developers - September](https://www.reddit.com/r/typescript/comments/ik9rft/whos_hiring_typescript_developers_september/)
- url: https://www.reddit.com/r/typescript/comments/ik9rft/whos_hiring_typescript_developers_september/
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
## [2][Why does setting 'as const' allow a string to match a string union type?](https://www.reddit.com/r/typescript/comments/iy8b7t/why_does_setting_as_const_allow_a_string_to_match/)
- url: https://www.reddit.com/r/typescript/comments/iy8b7t/why_does_setting_as_const_allow_a_string_to_match/
---
I had to use this inside a fetch request, I consider it hacky but it works:

        mode: 'cors' as const, // allows string to match a string union type

I learned this here but never found an explanation for it. why does asserting the string as a const variable suddenly let it match a string union type (which I saw in the error popup)?
## [3][Most Popular Programming Languages on GitHub - 2011/2019](https://www.reddit.com/r/typescript/comments/ixqlpu/most_popular_programming_languages_on_github/)
- url: https://youtu.be/eCUy0F-oVXA
---

## [4][Generics and Interfaces Confusion](https://www.reddit.com/r/typescript/comments/ixwhns/generics_and_interfaces_confusion/)
- url: https://www.reddit.com/r/typescript/comments/ixwhns/generics_and_interfaces_confusion/
---
Hi, I'm new to TypeScript. I have to create multiple tests and I want to ensure the setup is strongly typed.I was reading through the Generics documentation and I can understand the simple cases but I get lost with the advanced examples.

I tried doing something similar to what's described [here](https://www.typescriptlang.org/docs/handbook/generics.html) in the last example.

There are two test function signatures that varies depending on the tested value and the condition value:

    interface ITestFuncWithCond&lt;ValueType, ConditionType&gt; {
       (value: ValueType, condition: ConditionType): boolean;
    };
    
    interface ITestFuncNoCond&lt;ValueType&gt; {
       (value: ValueType): boolean;
    };

The idea is to do something like this with the tests (note the types):

    // Example A
    // No condition param
    const testA: Tester = createTest(StringTestIsEmpty);
    // Value param type must be string.
    testA.test("myString");
    
    // Example B
    // condition param is number
    const testB: Tester = createTest(StringLengthEquals, 2);
    // Value param type must be string.
    testB.test("myString");
    
    // Example C
    // condition param is string
    const testC: Tester = createTest(StringEquals, "myString");
    // Value param type must be string.
    testC.test("myString");
    
    // Example D
    // condition param is number
    const testD: Tester = createTest(NumberEquals, 10);
    // Value param type must be number.
    testD.test(5);

And this is the pseudo code that I have for this, but as you can probably see, I need some guiding words to untangle this.

    // Pseudo Internal setup
    
    interface ITestFuncWithCond&lt;ValueType, ConditionType&gt; {
       (value: ValueType, condition: ConditionType): boolean;
    };
    
    interface ITestFuncNoCond&lt;ValueType&gt; {
       (value: ValueType): boolean;
    };
    
    class Tester {
       name: TestName;
       // ConditionType???
       condition?: ConditionType;
       // ValueType??? ConditionType ???
       testImpl: ITestFuncWithCond&lt;ValueType, ConditionType&gt; | ITestFuncNoCond&lt;ValueType&gt;
    
       // Test constructor with different signatures.
       constructor(testName: TestName);
       constructor(testName: TestName, condition: ConditionType);
       constructor(testName: TestName, condition?: ConditionType) {
          this.name = testName; this.condition = condition;
          // Here, the testName is used to fetch the actual test function (testImpl)
          testImpl = getTestImpl(testName)
       
          // After the testImpl has been fetched, check if condition type "C" is as defined in "testImpl".
          if (condition) {
             if ((testImpl condition parameter type) === ConditionType) {
                // OK
             } else {
                // NOT OK
             }
          }
       }
    
       // ValueType should be string for test A, B and C, and should be number for test D.
       // "this.condition" should be undefined in test A.
       // "this.condition" should be "number for test B and D, and "string" for testC.
       test(value: ValueType): boolean {
          if (this.condition) {
             return testImpl(value, this.condition);
          } else {
             return testImpl(value);
          }
       }
    }
## [5][Convert class to interface during runtime](https://www.reddit.com/r/typescript/comments/ixngm0/convert_class_to_interface_during_runtime/)
- url: https://www.reddit.com/r/typescript/comments/ixngm0/convert_class_to_interface_during_runtime/
---
Hello everyone,
I'm using react/redux with typescript (functional) and integrating a 3rd party api client. The models defined in the 3rd party are class definitions, however the redux store complaints when trying to add class objects to it. Is there a way to create an interface that uses the field definitions of a class or turn the class definition into a native object during runtime? I want to avoid having to redefine the models as interfaces in order to prevent possible schema desync between my and the 3rd party library.

Edit: as pointed out by others in the comments, I actually want to cast to a native object not to an interface
## [6][Extending class methods to wrap all class instance methods](https://www.reddit.com/r/typescript/comments/ixrblq/extending_class_methods_to_wrap_all_class/)
- url: https://www.reddit.com/r/typescript/comments/ixrblq/extending_class_methods_to_wrap_all_class/
---
I am working with a fetching class that has methods to fetch different types of data. EX fetchDataA() fetchDataB(). I would like to extend this class basically so that I can provide an implicit wrapper to each instance method so I can specify request retrying, only fetching a single result, etc.   


How can I extend all the instance methods of a class without wrapping each individually?
## [7][A SQL database implemented purely in TypeScript type annotations.](https://www.reddit.com/r/typescript/comments/iww4hs/a_sql_database_implemented_purely_in_typescript/)
- url: https://github.com/codemix/ts-sql
---

## [8][When are we likely to see updated Clipboard access from typescript?](https://www.reddit.com/r/typescript/comments/ix861c/when_are_we_likely_to_see_updated_clipboard/)
- url: https://www.reddit.com/r/typescript/comments/ix861c/when_are_we_likely_to_see_updated_clipboard/
---
Currently the Clipboard interface in lib.d.ts only has readText and writeText.  When are we likely to see read and write, given they are pretty widely supported now?

Trying to decide if I should try to hack around this or wait.

Thanks
## [9][Debug Visualizer A VS Code extension for visualizing data structures while debugging. Like the VS Code's watch view, but with rich visualizations of the watched value.](https://www.reddit.com/r/typescript/comments/iwpzbm/debug_visualizer_a_vs_code_extension_for/)
- url: https://marketplace.visualstudio.com/items?itemName=hediet.debug-visualizer&amp;1
---

## [10][Jest Testing: How to import react components when the whole react app is in a typescript namespace and using C# ScriptBundle](https://www.reddit.com/r/typescript/comments/ix2rmw/jest_testing_how_to_import_react_components_when/)
- url: https://www.reddit.com/r/typescript/comments/ix2rmw/jest_testing_how_to_import_react_components_when/
---
 I'm having my whole ts react project in a namespace and now I want to cover this project with tests but I can't figure out how to import props so, I can access desired components that I want to test or import 'render from testing-library/react' in my test. For Example, my component looks like this:  
**MyComponent.tsx**

    module MainModule {
    
        interface Props extends ReactTabComponentProps {
        }
    
        interface State extends ReactComponentState {
        }
    
        export class MyComponent extends ReactTabComponent&lt;Props, State&gt; {
            constructor(props: Props) {
                super(props);
    
                this.state = {
                };
            }
    
            setTitle = () =&gt; {
                this.props.tab.setTitle(
                    &lt;span&gt;
                        &lt;Icon name="calculator" fixedWidth /&gt;
                        MainComponent
                    &lt;/span&gt;
                );
            }
    
            render() {
                return (
                    
                );
            }
        }
    }


My test file looks like this:  
**MyComponent.test.tsx**

    module MainModule {
    
     import { MainComponent } from '../Main/components/MainComponent';
     import { render } from '@testing-library/react';
     
        describe("mainComponentTest", () =&gt; {
            it("should renders correctly", async () =&gt; { 
               
            });
    
    
        });
    }

 With this, I get TS errors, Error: `(TS)Import declarations in a namespace cannot reference a module.`   
 And if I remove module from my test or write imports outside of module then I get the error:   
 

`(TS)Cannot compile modules using option 'outFile' unless the '--module'flag is 'amd' or 'system'.`  
Any help would be highly appreciated!
## [11][Safe Lookup Table in typescript](https://www.reddit.com/r/typescript/comments/iwsfwl/safe_lookup_table_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/iwsfwl/safe_lookup_table_in_typescript/
---
Is it possible to create a lookup table in typescript with functions that accept different type arguments?
I'm have a function that returns  a discriminated union with a type property that determines the type of data it holds, I would like to then use the type property as a key to an object with functions that accept types that the discriminated union was formed from
For example

    type a = {
        type:"obj1";
        foo:string;
    }
    type b={
        type:"obj2";
        bar:number;
    }
    type union= a or b;
    const lookupTable={
        obj1:({foo}:a)=&gt;foo.toLowerCase();
        obj2:({bar}:b)=&gt;(bar**2).toString(); // all the functions return the same type
    }
    //That all works, my issue comes when attempting to do something like
    const unionObj:union = {type:"obj1", foo:"str"}; // my code returns this from a function
    const result = lookupTable[unionObj.type](unionObj);
    // complicated type errors cause it can't tell that unionObj has been narrowed down to the right type

Can anyone please help with this
