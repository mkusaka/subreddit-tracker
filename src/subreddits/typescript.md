# typescript
## [1][Who's hiring Typescript developers - July](https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/)
- url: https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/
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
## [2][epubjs library](https://www.reddit.com/r/typescript/comments/hjrtwl/epubjs_library/)
- url: https://www.reddit.com/r/typescript/comments/hjrtwl/epubjs_library/
---
I am new with typescript and I had before .net / C# experience, nowadays I am dealing with typescript epubjs library. I developed epub-reader component and it is working as it should work, but there are cases with some epub files, where first page of book is always aligned on left side, when normally is on center. Those alignments are setup on popup window where you can choose "spread to one page" OR "spread to two pages"

&amp;#x200B;

dialog example : 

    &lt;mat-radio-group [(ngModel)]="epubTextOptions.selectedTwoPages" class="inline-button"&gt; &lt;mat-radio-button [value]="false" [checked]="!epubTextOptions.selectedTwoPages" class="inline-button"&gt;{{'reader.textOptions.dialog.onePage' | translate}}&lt;/mat-radio-button&gt; &lt;mat-radio-button [value]="true" [checked]="epubTextOptions.selectedTwoPages" class="inline-button"&gt;{{'reader.textOptions.dialog.twoPages' | translate}}&lt;/mat-radio-button&gt; &lt;/mat-radio-group&gt;

 typescript example:  

    export class MyClass implements OnInit {   book: Book;   rendition: Rendition; private myMethod() { this.rendition = this.book.renderTo(this.ePub.nativeElement, {       width: '100%',       height: '100%' }); this.rendition.spread(this.epubTextOptions.selectedTwoPages ? 'always' : 'none'); } }

 html component : 

    &lt;div id="epubContent" class="d-flex align-items-center" [style.height]="contentHeight + 'px'"&gt; &lt;app-button *ngIf="displayPrevPageButton() &amp;&amp; !isMobileDevice()" (btnClick)="prevPage()"               btnClass="mat-icon-button epub-action"&gt;&lt;span class="far fa-chevron-left"&gt;&lt;/span&gt;&lt;/app-button&gt; &lt;div #ePub class="epub-container"&gt;&lt;/div&gt; &lt;app-button *ngIf="displayNextPageButton() &amp;&amp; !isMobileDevice()" (btnClick)="nextPage()"               btnClass="mat-icon-button epub-action"&gt;&lt;span class="far fa-chevron-right"&gt;&lt;/span&gt;&lt;/app-button&gt; &lt;/div&gt;

 Everything is implemented from epubjs library examples and adapted to reader, but as I said with most of epub spread() functionality is working, but there are epub files, where first page align to left side when "one page" is selected instead of center align. Maybe some how I can check if file have no more than one page and if yes, then force to align to center?
## [3][Type 'string' is not assignable to type '"X" | "Y" | "Z"'](https://www.reddit.com/r/typescript/comments/hjvw3n/type_string_is_not_assignable_to_type_x_y_z/)
- url: https://www.reddit.com/r/typescript/comments/hjvw3n/type_string_is_not_assignable_to_type_x_y_z/
---
Recently I had to use the `ReturnType` generic type constructor to build a type for me, that was the return shape of an obsurely documented method in Google Text To Speech's API:

    protected async fetchAndWriteAudio(options: {
          input : { text : string }
          , voice : { languageCode : string, ssmlGender: voiceGender }
          , audioConfig : { 
             audioEncoding : "AUDIO_ENCODING_UNSPECIFIED" | "LINEAR16" | "MP3" | "OGG_OPUS" }
       }, fileNameAndPath: string)
          : Promise&lt;ReturnType&lt;typeof TextToSpeechClient.prototype.synthesizeSpeech&gt;&gt; {

But now the title error shows on the error-commented lines below. I thought a specific string value, depicted as a quad union from the `ReturnType` constructor, should match one of those specific strings. Instead, none match. Why is that and what is the proper fix?

       public async makeSentenceAudio(subFolderPath: string): Promise&lt;void&gt; {
          // setup the foreign and english text string
          const foreignSentenceText = this.textTranslate(
             this.sentence.englishVersion
             , translationDirection.toForeign);
          const englishText = this.sentence.englishVersion;
    
          
          /**** Generate Audios ****/
          // define audio options (including text)
          const sharedOptions = {
             input: { text : "" }
             , voice: { languageCode : "", ssmlGender : voiceGender.male }
             , audioConfig : {audioEncoding : "OGG_OPUS"}
          }
    
          const foreignAudioOptions = {
             ...sharedOptions
             , input: {text: ""}
          }
    
          const englishAudioOptions = {
             ...sharedOptions
             , input: {text: ""}
          }
    
          this.fetchAndWriteAudio(foreignAudioOptions, "placeholder");
          this.fetchAndWriteAudio(englishAudioOptions, "placeholder");

.

    // the error message. There was also one on englishAudioOptions
    
    const foreignAudioOptions: {
        input: {
            text: string;
        };
        voice: {
            languageCode: string;
            ssmlGender: voiceGender;
        };
        audioConfig: {
            audioEncoding: string;
        };
    }
    Argument of type '{ input: { text: string; }; voice: { languageCode: string; ssmlGender: voiceGender; }; audioConfig: { audioEncoding: string; }; }' is not assignable to parameter of type '{ input: { text: string; }; voice: { languageCode: string; ssmlGender: voiceGender; }; audioConfig: { audioEncoding: "OGG_OPUS" | "AUDIO_ENCODING_UNSPECIFIED" | "LINEAR16" | "MP3"; }; }'.
    
      The types of 'audioConfig.audioEncoding' are incompatible between these types.
        Type 'string' is not assignable to type '"OGG_OPUS" | "AUDIO_ENCODING_UNSPECIFIED" | "LINEAR16" | "MP3"'. ts(2345)
## [4][TIL What TypeScript Is](https://www.reddit.com/r/typescript/comments/hj8azp/til_what_typescript_is/)
- url: https://www.reddit.com/r/typescript/comments/hj8azp/til_what_typescript_is/
---
So last night, I finally watched a short intro on what exactly typescript is, and... Wow.

I wish I had known about this years ago; as someone who loves languages like C++, and never really was a fan of JS, TS makes the web world so much more inviting.

Anyway, happy to be here, sad it took me so long to learn about it.
## [5][Checking for interface collisions .](https://www.reddit.com/r/typescript/comments/hjif2w/checking_for_interface_collisions/)
- url: https://www.reddit.com/r/typescript/comments/hjif2w/checking_for_interface_collisions/
---
Imagine having multiple different interfaces that all together are mixed into a big interface .

Before some time I [asked](https://www.reddit.com/r/typescript/comments/gab4ic/how_to_make_ts_or_eslint_or_vscode_warn_me_about/) how can I check for collisions when I merge two interfaces and I got an answer that worked . Here is a more simple version of it :

    type isNever&lt;T extends never&gt; = T;
    //I just look if I get a linting error in the following type 
    //If yes then there is a collision
    type mustBeNever = isNever&lt;keyof interface1 &amp; keyof interface2&gt;;

Now I have to merge more than two interfaces . With that way , for n interfaces I need to do `n!/(2!(n-2)!) = n(n-1)/2` manual checks (if you are interested in how that number is calculated take a look [here](https://en.wikipedia.org/wiki/Combination)) , something that is not practical .

Is there any solution via using a type function like I have already done that is not unpractical or should I create a unit test ?
## [6][How to ensure the right function is called with the right arguments?](https://www.reddit.com/r/typescript/comments/hjlrr1/how_to_ensure_the_right_function_is_called_with/)
- url: https://www.reddit.com/r/typescript/comments/hjlrr1/how_to_ensure_the_right_function_is_called_with/
---
Suppose I've got a login and a logout function:

    login(creds: { user: string, pass: string }): string;
    logout(token: string): void;

And a handler decides which function to call depending on a variable:

    let f = command === 'login' ? login : logout;
    let x = command === 'login' ? creds : token;
    return f(x);

I know that if f is logout, then x is going to be a string. Likewise if f is login, then x is going to be a credentials object. Can you help to model that?

Thanks!
## [7][skipLibCheck but for .ts files?](https://www.reddit.com/r/typescript/comments/hjgd30/skiplibcheck_but_for_ts_files/)
- url: https://www.reddit.com/r/typescript/comments/hjgd30/skiplibcheck_but_for_ts_files/
---
I am trying to use the [tfjs](https://github.com/tensorflow/tfjs) library for machine learning in a typescript project with `strictNullChecks`. Unfortunately, tfjs doesn't use strictNullChecks internally. As a result, TypeScript fails to compile my project, because there are .ts files in tfjs where null or undefined is passed to functions without being explicitly allowed.

I know that `skipLibCheck` option exists for these cases, but it seems like it only affects `.d.ts` files. In the case of tfjs, there are actual `.ts` files (e.g. `node_modules/@tensorflow/tfjs-core/src/environment.ts`) which cause the failure. This seems strange, because while these files exist, they are not used in my project (only .d.ts definitions and pre-built .js files are required).

It also seems like adding "node_modules/@tensorflow" to tsconfig's "exclude" does not  fix the problem.

What can I do to make my project build without issues?

Thanks!
## [8][How to check if a whole number (1.0) is a double in TypeScript](https://www.reddit.com/r/typescript/comments/hj9qtk/how_to_check_if_a_whole_number_10_is_a_double_in/)
- url: https://www.reddit.com/r/typescript/comments/hj9qtk/how_to_check_if_a_whole_number_10_is_a_double_in/
---
I have a library for developers with a method like this:

        function add (value: number | string) {
          if (typeof value === 'number') {
            if (value % 1 == 0) {
              // post to database as an int
            } else {
              // post to database as a double
            }
          }
          ...
        }

The problem here arises when *value* is something like 2.0 or 10.0, this gets treated as an int. However, I need those values to be passed as doubles. How can I check for these 'whole numbers' (10.0 or 2.0) to be doubles?
## [9][Prints a dependency graph in dot format for your typescript project.](https://www.reddit.com/r/typescript/comments/hj7hct/prints_a_dependency_graph_in_dot_format_for_your/)
- url: https://github.com/PSeitz/ts-dependency-graph
---

## [10][Help digging a return type out of a big library (Google Cloud)](https://www.reddit.com/r/typescript/comments/hj6evi/help_digging_a_return_type_out_of_a_big_library/)
- url: https://www.reddit.com/r/typescript/comments/hj6evi/help_digging_a_return_type_out_of_a_big_library/
---
[https://googleapis.dev/nodejs/text-to-speech/latest/v1.TextToSpeechClient.html](https://googleapis.dev/nodejs/text-to-speech/latest/v1.TextToSpeechClient.html)

&gt;synthesizeSpeech(request, optionsopt) → {Promise}  
&gt;  
&gt;Synthesizes speech synchronously: receive results after all text input has been processed.

The library comment on this method:

    * @returns {Promise} - The promise which resolves to an array.
         *   The first element of the array is an object representing [SynthesizeSpeechResponse]{@link google.cloud.texttospeech.v1.SynthesizeSpeechResponse}.
         *   The promise has a method named "cancel" which cancels the ongoing API call.
         */

And on that page:

[https://cloud.google.com/text-to-speech/docs/reference/rpc/google.cloud.texttospeech.v1#google.cloud.texttospeech.v1.SynthesizeSpeechResponse](https://cloud.google.com/text-to-speech/docs/reference/rpc/google.cloud.texttospeech.v1#google.cloud.texttospeech.v1.SynthesizeSpeechResponse)

&gt;SynthesizeSpeechResponse  
&gt;  
&gt;The message returned to the client by the SynthesizeSpeech method.

&amp;#x200B;

My code so far:

    const TextToSpeechClient = require("@google-cloud/text-to-speech").TextToSpeechClient;
    
    // Cannot find namespace 'TextToSpeechClient'.ts(2503)
          : Promise&lt;TextToSpeechClient.SynthesizeSpeechResponse&gt; 

Full method below. I got close but I'm still rough with digging around an API, trying to extract these obscure types.

I want to learn how to do it though. Marking `any` is easy, specifying the correct type on a fetch especially is a lot more insightful and easy to backtrace later.

Thanks for any pointers on getting it right.

&amp;#x200B;

Full method (PS: I plan to split the fetch and save next. Maybe also implement a named `options` interface)

       protected async fetchAndWriteAudio(options: {
          input : { text : string }
          , voice : { languageCode : string, ssmlGender: voiceGender }
          , audioConfig : { audioEncoding : string }
       }, fileNameAndPath: string)
          : Promise&lt;TextToSpeechClient.SynthesizeSpeechResponse&gt; {
    
          const textToSpeech = new TextToSpeechClient();
          const writeFileAsync = util.promisify(writeFile);
    
          try {
             const [audioResponse] = await textToSpeech.synthesizeSpeech(options);
             await writeFileAsync(fileNameAndPath, audioResponse.audioContent);
          }
          catch(error) { console.log(error); }
       }
## [11][Compiler gives no warning with wrong extended class fed into a constructor](https://www.reddit.com/r/typescript/comments/hj67q0/compiler_gives_no_warning_with_wrong_extended/)
- url: https://www.reddit.com/r/typescript/comments/hj67q0/compiler_gives_no_warning_with_wrong_extended/
---
Is there a way for the compiler to give warning/error because I am giving Y class instead of Z class on P constructor?

    class X { }
    
    class Y extends X { } 
    class Z extends X { }
    
    const y = new Y()
    
    class P { 
        constructor(private z: Z) { }
    
        get get(): Z { 
            return this.z  
        } 
    }
    const p = new P(y) 
    console.log(p.get)

&amp;#x200B;
