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
## [2][Am I being pedantic when it comes to helper methods in classes?](https://www.reddit.com/r/typescript/comments/hqar0k/am_i_being_pedantic_when_it_comes_to_helper/)
- url: https://www.reddit.com/r/typescript/comments/hqar0k/am_i_being_pedantic_when_it_comes_to_helper/
---
For the most part, I try to program in a functional style - simple, reusable functions, accessing state, where needed, via closure, etc.  And certainly \*anything\* you can write as a class in JS/TS can be written as functions.  But sometimes classes are a simpler, easier solution and I find that when you have complex state to manage, they're nice to use. (Basically, when I start thinking that something will require a monad, I think it might be okay to switch to classes. :) ) 

The question occurs when I have a class definition that Is more or less useless outside the class, but is purely functional - that is, no access to state, no mutations. 

Now, I would insist these functions be defined \*outside the class\* so that they can be independently unit tested.  (Though to keep concerns together, I'd put them in the same \*file\* where the class is the default export).  But is that being pedantic?    


In other words, I'm refactoring this: 

```typescript
export default class Match {
  private checkCondition = (condition: Condition, input: any): boolean =&gt; {
    if (typeof condition === "function") {
      return condition(input);
    }
    return condition === input;
  };
}
```

to this:
```typescript
export const checkCondition = (condition: Condition, input: any): boolean =&gt; {
  if (typeof condition === "function") {
    return condition(input);
  }
  return condition === input;
};

export default class Match {
  /* OPTIONAL!
  private checkCondition = checkCondition;
  */
}
```
## [3][Express.js course with TypeScript Lesson 2 — Apollo &amp; WebSockets](https://www.reddit.com/r/typescript/comments/hqcfqa/expressjs_course_with_typescript_lesson_2_apollo/)
- url: https://medium.com//express-js-course-with-typescript-lesson-2-apollo-websockets-7eb2063186bf?source=friends_link&amp;sk=123f4b00527f5080717f1cbfa6f34ac2
---

## [4][Access mapped type by index](https://www.reddit.com/r/typescript/comments/hqc4h2/access_mapped_type_by_index/)
- url: https://www.reddit.com/r/typescript/comments/hqc4h2/access_mapped_type_by_index/
---
I'm trying to wrap my head around the following behavior, say I have the following type definitions

    type RMap = {
        a: { foo: string }
        b: { foo: unknown }
    }
    
    type RType = keyof RMap
    
    type RState = {
        [k: string]: {
            s: RMap[RType]
        }
    }

Given these, what's the difference between f1 and f2 below? Why does the first one not compile?

    declare const map: RState
    
    function f1&lt;R extends RType&gt;(k: string, f: (r: RMap[R]) =&gt; void): void {
        f(map[k].s)
    }
    
    function f2(k: string, f: (r: RMap[RType]) =&gt; void): void {
        f(map[k].s)
    }
## [5][Having a hard time Typing a ref that is a function that gets assigned to ref.curent](https://www.reddit.com/r/typescript/comments/hq8tnc/having_a_hard_time_typing_a_ref_that_is_a/)
- url: https://www.reddit.com/r/typescript/comments/hq8tnc/having_a_hard_time_typing_a_ref_that_is_a/
---
Code:

    const memoizedHandler = useCallback((e) =&gt; handler(e), dependencies);useEventListener({ eventName: 'click', handler: memoizedHandler }); // &lt;-- error here
    ERROR:Error:(17, 42) TS2741: Property 'current' is missing in type '(e: any) =&gt; void' but required in type 'RefObject&lt;HTMLDivElement&gt;'.

I am assigning to "current", WHY do I need it on the function I am trying to assign to it?

which says it is originating from:

    useEffect(() =&gt; {
      savedHandler.current = handler;
    }, [handler]);
    
    useEffect(() =&gt; {
      const isSupported = element &amp;&amp; element.addEventListener;
      if (!isSupported) return;
    
      const eventListener = (event: React.ChangeEvent&lt;HTMLDivElement&gt;) =&gt; 
      savedHandler?.current?.(event);
    
    // error at "current" says TS2349: This expression is not callable.
    // Type 'RefObject&lt;HTMLDivElement&gt;' has no call signatures

getting so frustrated. lol
## [6][Code review request of one of my classes for a Node app](https://www.reddit.com/r/typescript/comments/hpy9qo/code_review_request_of_one_of_my_classes_for_a/)
- url: https://www.reddit.com/r/typescript/comments/hpy9qo/code_review_request_of_one_of_my_classes_for_a/
---
I do not claim to be an advanced TS programmer yet but I'd like an interviewer to at least think "this guy can contribute value fairly quickly".

The class below interacts with 2 Google APIs to translate text to a foreign language, and also build audio files from a text-to-speech library.

I don't think anyone wants to get passed an entire repo to review. but if you do want to see how the class gets utilized here it is: [https://github.com/SeanDez/foreignSentenceRepeater](https://github.com/SeanDez/foreignSentenceRepeater)

The class I would like feedback on (posted below): [https://github.com/SeanDez/foreignSentenceRepeater/blob/master/src/classes/sentenceBuilder/AudioMaker.ts](https://github.com/SeanDez/foreignSentenceRepeater/blob/master/src/classes/sentenceBuilder/AudioMaker.ts)

    import fs, { writeFile } from "fs";
    import path from "path";
    import util from "util";
    import readLine from "readline-sync";
    import { audioParentFolderPath, silenceFolderPath, oneSecondPause, twoSecondPause, threeSecondPause, fourSecondPause, fiveSecondPause } from "../../globals";
    const audioConcat = require("audioconcat");
    const tmp = require("tmp");
    
    import Sentence from "./Sentence";
    import Utilities from "../Utilities";
    import ConfigData from "../setupWizard/ConfigDataInterface";
    import ForeignPhraseDefinitionPair from "./ForeignPhraseDefinitionPairInterface";
    import {translationDirection, voiceGender, audioEncoding} from "../minorTypes";
    import {TextToSpeechClient} from "@google-cloud/text-to-speech";
    import WordFile, {contentTypes} from "./WordFile";
    
    
    const {TranslationServiceClient} = require('@google-cloud/translate');
    
    
    export default class AudioMaker {
       // --------------- Properties
       private configData: Readonly&lt;ConfigData&gt;;
       private sentence: Sentence;
    
       
       // --------------- Constructor
       constructor(configData: Readonly&lt;ConfigData&gt;, sentence: Sentence) {
          this.configData = configData;
          this.sentence = sentence;
       }
    
    
       // --------------- Public Methods
    
       /* 
          Silence files are kept in {rootDir}/silences. 
          
          Delays from 1 to 12 seconds are available in 1 
          second increments
    
          Note: The audio combiner adds 2 seconds
       */
      public calculatePauseDuration(wordCount: number)
      : number {
          // word 1 gets 2 seconds automatically from the audio combiner.
          // this means all values are low by 2 seconds
          let pauseDuration = wordCount - 1;
          if (pauseDuration &gt; 12) pauseDuration = 12;
    
          return pauseDuration;
       }
    
       public makeSentenceFolder(): string {
          const subfolderPath = path.join(audioParentFolderPath, this.sentence.folderName);
    
          fs.mkdirSync(subfolderPath);
    
          return subfolderPath;
       }
    
    
       /* 
          Makes an audio of the sentence, in the sentence's subfolder
       */
       public async makeSentenceTrack(prefix: string): Promise&lt;void&gt; {
    
          /**** Setup Audio Options ****/
          const sharedOptions = {
             voice : { ssmlGender : voiceGender.male }
             , audioConfig : {audioEncoding : "OGG_OPUS" as audioEncoding }
          }
    
          let foreignSentenceText;
          try {
             foreignSentenceText = await this.textTranslate(
                this.sentence.englishVersion
                , translationDirection.toForeign);   
          }
          catch(error) { 
             console.log(error); 
             throw new Error(error);
          }
    
          const foreignAudioRequest = {
             ...sharedOptions
             , voice : { 
                ...sharedOptions.voice
                , languageCode : this.configData.languageCode 
             }
             , input : {text : foreignSentenceText}
          }
    
    
          const englishText = this.sentence.englishVersion;
          const englishAudioRequest = {
             ...sharedOptions
             , voice : { 
                ...sharedOptions.voice
                , languageCode : "en" 
             }
             , input : {text : englishText}
          }
    
          /**** Create 1st sentence audio ****/
          const tempFolder = tmp.dirSync({unsafeCleanup: true});
          
          const foreignAudioName = `${prefix} - ${foreignSentenceText}.ogg`;
          const foreignAudioTempPath = path.join(`${tempFolder}`, foreignAudioName);
    
          const englishAudioName = `${prefix} - ${this.sentence.folderName}.ogg`;
          const englishAudioTempPath = path.join(tempFolder, englishAudioName);
          
          this.fetchAndWriteAudio(foreignAudioRequest, foreignAudioTempPath);
          this.fetchAndWriteAudio(englishAudioRequest, englishAudioTempPath);
    
    
          /**** Add Pause ****/
          // 2 for 1st word, plus 1 per word thereafter
          let mainPauseDuration: number = 2 + this.sentence.foreignWordCount;
          if (mainPauseDuration &gt; 12) { mainPauseDuration = 12 }
          const pauseFilePath = path.join(silenceFolderPath, `${mainPauseDuration}.ogg`);
          
    
          /**** Setup final audio file structure ****/
          const singlePassStructure = [englishAudioTempPath, pauseFilePath, foreignAudioTempPath, threeSecondPause];
    
          let endStructure = singlePassStructure;
          for (let i = 1; i &lt;= this.configData.numberOfRepeats - 1; i++) {
             const repeatStructure = [englishAudioTempPath, twoSecondPause, foreignAudioTempPath, threeSecondPause];
    
             endStructure.concat(repeatStructure);
          }
    
          /**** Save To Production Subfolder ****/
          const finalSaveFolderPath: string = path.join(audioParentFolderPath, this.sentence.folderName);
    
    
          this.combineAndSave(
             endStructure
             , finalSaveFolderPath
             , {prefix}
          )
       }
    
    
       public async makeWordAudioFile(): Promise&lt;void&gt; {
          // sets data to this.sentence.foreignPhraseDefinitionPairs
          await this.gatherAllForeignWordsAndDefinitionsFromUser();
    
          /**** Build word def audios to temp directory ****/
          const tempFolder: ReturnType&lt;typeof tmp.dirSync&gt; = tmp.dirSync({unsafeCleanup : true});
    
          this.buildWordDefinitionAudiosToTempFolder(tempFolder);
    
          const tempFileNames: Array&lt;string&gt; = fs.readdirSync(tempFolder);
          
          /**** Convert from filenames to WordFile objects  ****/
          const wordFileObjects: Array&lt;WordFile&gt; = tempFileNames.map(fileName =&gt; {
             return new WordFile(fileName, tempFolder);
          });
    
          /**** Setup order of files to be combined, including silences ****/
          const finalAudioOrder: string[] = this.setAudioOrderFromWordFileObjects(wordFileObjects);
    
    
          /**** Build Single Production File ****/
          const productionFileName = `2 - all words and definitions.ogg`;
          const finalSaveFolderPath: string = path.join(audioParentFolderPath, this.sentence.folderName);
          const fullSavePath = path.join(finalSaveFolderPath, productionFileName);
    
          this.combineAndSave(
             finalAudioOrder
             , fullSavePath
             , {name: productionFileName}
          )
       }
    
    
       /* 
          Copies file to same folder with a different filename
    
          @param copiedFileName. Should contain the full filename including extension
       */
       public duplicateTrack(
          prefixMatcher: string
          , copiedFileName: string): void {
    
          const targetAudioFolder = path.join(audioParentFolderPath, this.sentence.folderName);
    
          const audioFileNames: string[] = fs.readdirSync(targetAudioFolder);
          
          const regexMatcherFromBeginning: RegExp = new RegExp(`^${prefixMatcher}`);
          
          const targetFile: string = audioFileNames.filter(filename =&gt; {
             const matchFound: boolean = regexMatcherFromBeginning.test(filename);
             
             return matchFound;
          })[0];
    
          const sourceFileNameAndPath = path.join(targetAudioFolder, targetFile);
          const copiedFileNameAndPath = path.join(targetAudioFolder, copiedFileName);
    
          fs.copyFileSync(sourceFileNameAndPath, copiedFileNameAndPath);
       }
    
    
       // --------------- Internal Methods
    
       protected parseFileContents(filepath = path.join(__dirname, "../../../sentences.txt")): Array&lt;string&gt; {
          const sentenceCandidates = fs
             .readFileSync(filepath)
             .toString()
             .split("\n");
    
          return sentenceCandidates;
       }
    
    
       protected async textTranslate(
             wordPhraseOrSentence: string
             , direction: translationDirection
          ) : Promise&lt;string&gt; {
          const translationClient = new TranslationServiceClient();
    
          let sourceLanguage: string;
          let targetLanguage: string;
          if (direction === translationDirection.toEnglish) {
             sourceLanguage = this.configData.languageCode;
             targetLanguage = "en";
          } else {
             sourceLanguage = "en";
             targetLanguage = this.configData.languageCode;
          }
    
          const options = {
             parent: `projects/${this.configData.projectId}`
             , contents: [wordPhraseOrSentence]
             , mimeType: 'text/plain'
             , sourceLanguageCode: sourceLanguage
             , targetLanguageCode: targetLanguage
          }
    
          try {
             const [response] = await translationClient.translateText(options);
             const {translations}: { translations: string[] } = response;
             return translations[0];
          }
          catch (error) {
             console.error(error.details)
             throw new Error(error.details);
          }
    
       }
       
    
       /* 
          Asks for the foreign word
          
          Gets a definition, asks for confirmation or an adjustment.
    
          Returns an object with the foreign word and definition pair
          
          Returns false if user marks "done" commands
       */
       protected async getForeignWordAndDefinition()
          : Promise&lt;ForeignPhraseDefinitionPair | false&gt; {
          console.log("Please copy and paste the (next) foreign word in the sentence here. Or type -d or --done when all words in the sentence have been specified.")
          const userInput = readLine.question();
    
          const userHasExited = this.isDone(userInput);
          if (userHasExited) return false;
    
          const foreignWord = userInput;
    
          const googleOfferedDefinition: string = await this.textTranslate(
             foreignWord
             , translationDirection.toEnglish
          );
    
          console.log("Type your own contextual definition for this word now. It will be used during audio translation. Or, press ENTER without typing anything to accept the following default definition from Google Translate:")
          console.log(googleOfferedDefinition);
          const userDefinition: string = readLine.question();
    
          let acceptedDefinition = googleOfferedDefinition;
          if (userDefinition !== "") {
             acceptedDefinition = userDefinition;
          } 
    
          // shape the object and return it
          const foreignPhraseDefinitionPair: ForeignPhraseDefinitionPair = {
             foreignPhrase: foreignWord
             , englishDefinition: acceptedDefinition
          }
    
          return foreignPhraseDefinitionPair;
       }
    
    
       /* 
          used to exit a user input loop
       */
       private isDone(userInput: string): boolean {
          if (userInput === "-d" || userInput === "--done") {
             return true;
          }
    
          return false;
       }
    
    
       /* 
          send a text-to-speech request
          catcheds the audio stream. Saves to file
       */
       protected async fetchAndWriteAudio(
          request: Readonly&lt;{
             input : { text : string }
             , voice : { languageCode : string, ssmlGender: voiceGender }
             , audioConfig : { 
                audioEncoding : "AUDIO_ENCODING_UNSPECIFIED" | "LINEAR16" | "MP3" | "OGG_OPUS" }
          }&gt;
          , fileNameAndPath: string
       ) : Promise&lt;ReturnType&lt;typeof TextToSpeechClient.prototype.synthesizeSpeech&gt;&gt; {
    
          const textToSpeech = new TextToSpeechClient();
          const writeFileAsync = util.promisify(writeFile);
    
          try {
             const [audioResponse] = await textToSpeech.synthesizeSpeech(request);
             await writeFileAsync(fileNameAndPath, audioResponse.audioContent!);
          }
          catch(error) { console.log(error); }
       }
    
       
       protected setAudioOrderFromWordFileObjects(wordFiles: Array&lt;WordFile&gt;): Array&lt;string&gt; {
          const finalAudioStructure: string[] = [];
    
          wordFiles.forEach(wordFile =&gt; {
             const hasBeginningPause = wordFile.beforePausePadding &gt; 0;
             if (hasBeginningPause) {
                const beginningPauseFile = path.join(silenceFolderPath, `${wordFile.beforePausePadding}.ogg`);
                finalAudioStructure.push(beginningPauseFile);
             }
    
             finalAudioStructure.push(wordFile.fullFilePath);
    
             const hasEndingPause = wordFile.beforePausePadding &gt; 0;
             if (hasEndingPause) {
                const endingPauseFile = path.join(silenceFolderPath, `${wordFile.afterPausePadding}.ogg`);
                finalAudioStructure.push(endingPauseFile);
             }
    
          });
    
          return finalAudioStructure;
       }
    
    
       /* 
          save an array of audios to a production folder
    
          argument 3 takes a prefix, or full filename/extension
       */
       protected combineAndSave &lt;NamingOption extends {prefix: string, name: string}&gt; (
          audiosAndPauseFiles: Array&lt;string&gt;
          , savePath: string
          , fileNameOptions : 
             Pick&lt;NamingOption, "prefix"&gt; |
             Pick&lt;NamingOption, "name"&gt;
       ) : void {
          let finalFileSavePath: string;
          
          // used to save sentence files
          if ("prefix" in fileNameOptions) {
             finalFileSavePath = path.join(savePath, `${fileNameOptions.prefix} - ${this.sentence.folderName}.ogg`);
          } else {
             finalFileSavePath = path.join(savePath, name);
          }
    
          audioConcat(audiosAndPauseFiles)
             .concat(finalFileSavePath)
             .on("start", (command: any) =&gt; {
                console.log(`ffmpeg build process started on file at: ${finalFileSavePath}`);
             })
             .on("end", (output: any) =&gt; {
                console.log(`Sucessfully created file at: ${finalFileSavePath}`);
             })
             .on("error", (error: any, stdout: any, stderr: any) =&gt; {
                console.log('error', error);
                console.log('stdout', stdout);
                console.log('stderr', stderr);
             });
       }
    
     
       protected buildWordDefinitionAudiosToTempFolder(
          tempFolder: ReturnType&lt;typeof tmp.dirSync&gt;
       ) : void {
          let pairNumber: number = 1;
    
          this.sentence.foreignPhraseDefinitionPairs.forEach(wordDefinitionPair =&gt; {
             /**** Setup request object ****/
             const sharedRequestOptions = {
                voice : { ssmlGender : voiceGender.female }
                , audioConfig : { audioEncoding : "OGG_OPUS" as audioEncoding }
             }
    
             const foreignWordOptions = {
                ...sharedRequestOptions
                , voice : {
                   ...sharedRequestOptions.voice
                   , languageCode : this.configData.languageCode
                }
                , input : {text: wordDefinitionPair.foreignPhrase}
             }
    
    
             const englishDefinitionOptions = {
                ...sharedRequestOptions
                , voice : {
                   ...sharedRequestOptions.voice
                   , languageCode : "en"
                }
                , input : {text: wordDefinitionPair.englishDefinition}
             }
    
             /**** Setup file names &amp; paths ****/
    
             // foreign words are assigned 1 in the second position. english words are assigned 2
             const foreignWordFileName = `${pairNumber}1 - foreign word - ${wordDefinitionPair.englishDefinition}.ogg`;
             const foreignWordFullPath = path.join(tempFolder, foreignWordFileName);
    
             const englishDefinitionFileName = `${pairNumber}2 - definition - ${wordDefinitionPair.englishDefinition}.ogg`;
             const englishDefinitionFullPath = path.join(tempFolder, foreignWordFileName);
    
    
             /**** Translate and save ****/
             for (let i = 1; i &lt;= this.configData.numberOfRepeats; i++) {
                this.fetchAndWriteAudio(foreignWordOptions, foreignWordFullPath);
                this.fetchAndWriteAudio(englishDefinitionOptions, englishDefinitionFullPath);
       
                pairNumber += 1;   
             }
    
             pairNumber++;
          });
    
       }
    
    
       /* 
          Pushes all gathered data to: this.sentence.foreignPhraseDefinitionPairs
       */
       protected async gatherAllForeignWordsAndDefinitionsFromUser() : Promise&lt;void&gt; {
          enum sequentialAdjectives {
             first = "first"
             , next = "next"
          }
          let sequentialAdjective: sequentialAdjectives = sequentialAdjectives.first;
          let continueLooping: boolean = true;
    
          while (continueLooping) {
             console.log(`Please enter the ${sequentialAdjective} foreign language word and press ENTER.`)
             
             if (sequentialAdjective === sequentialAdjectives.next) {
                console.log(`Or type "--done" or "-d" (no quotes) to complete word definitions for this sentence.`);
             }
    
             const foreignWordUserInput: string = readLine.question();
    
             if (sequentialAdjective === sequentialAdjectives.first) {
                sequentialAdjective = sequentialAdjectives.next;
              }
    
             /**** Exit Check ****/
             const userExited: boolean = this.isDone(foreignWordUserInput);
    
             if (userExited) {
                continueLooping = false;
             }
             else {
                /**** Fetch Google Translation ****/
                const googlesuggestedTranslation: string = await this.textTranslate(foreignWordUserInput, translationDirection.toEnglish);
    
                /**** Ask user to accept, or override ****/
                console.log(`The translation returned by Google for ${foreignWordUserInput} is: ${googlesuggestedTranslation}`);
                console.log(`Press ENTER (without entering any text) to accept this translation. Or, type your own custom definition, and press ENTER.`);
    
                const definitonUserInput = readLine.question();
    
                let acceptedDefinition;
                if (definitonUserInput === "") {
                   acceptedDefinition = googlesuggestedTranslation;
                } else {
                   acceptedDefinition = definitonUserInput;
                }
    
                const wordDefinitionPair: ForeignPhraseDefinitionPair = {
                   foreignPhrase : foreignWordUserInput
                   , englishDefinition : acceptedDefinition
                };
    
                this.sentence.foreignPhraseDefinitionPairs.push(wordDefinitionPair);
             }
          }
       }
    
    }

Why protected methods? Makes it easy to extend and then test in jest as I (currently) start the Jest testing and debugging process. I'll turn them private before marking the repo beta. If you don't like this strategy I'm all ears as to how you would go from draft stage to testing to production, I know my process there is kinda fuzzy. Maybe i should have drafted and tested together.

Thanks for any suggestions; I will try my best to apply them to my next project.
## [7][Importing .ts files without extension](https://www.reddit.com/r/typescript/comments/hps207/importing_ts_files_without_extension/)
- url: https://www.reddit.com/r/typescript/comments/hps207/importing_ts_files_without_extension/
---
Hi, I recently migrated a vue project to typescript manually without cli, and I managed to set everything up except importing ts files. Somehow I need to specify the .ts on the import path.

Can someone please help?

&amp;#x200B;

Edit: I found the solution, the problem is I manually install typescript so there's no shims-vue.d.ts generated, just put it on the src of the your directory at it should work fine  


&amp;#x200B;
```
-project dir
|-src
 |-shims-vue.d.ts
|-tsconfig.json
```

I swear I hate setting up these kind of config
## [8][After learning Rust, I was surprised that the concept of "match" (also in Haskell) didn't make it to ES2020/TS4.0. So I created my own little class for it. Thoughts?](https://www.reddit.com/r/typescript/comments/hphpu8/after_learning_rust_i_was_surprised_that_the/)
- url: https://gist.github.com/brianboyko/1decb88a793b92c06ad7d13f73d88854
---

## [9][Q: Using yarn workspace but Typescript isn't aware of it](https://www.reddit.com/r/typescript/comments/hpqnrm/q_using_yarn_workspace_but_typescript_isnt_aware/)
- url: https://www.reddit.com/r/typescript/comments/hpqnrm/q_using_yarn_workspace_but_typescript_isnt_aware/
---
When I import something from an another workspace I get `Cannot find module @org/common/Movie or its corresponding type...` from the TS compiler/Intellisense. How can make TS aware of my workspace structure?

Edit: for others driving by, the solution was to add `references` to the importing module's tsconfig pointing to the imported (shared) module's folder `path`. and adding `"composite": "true"` to the tsconfig of the imported (shared) module.
## [10][Type '"X"' does not satisfy the constraint 'keyof NamingOption'.ts(2344)](https://www.reddit.com/r/typescript/comments/hpranq/type_x_does_not_satisfy_the_constraint_keyof/)
- url: https://www.reddit.com/r/typescript/comments/hpranq/type_x_does_not_satisfy_the_constraint_keyof/
---
       protected combineAndSave &lt;NamingOption&gt;(
          audiosAndPauseFiles: Array&lt;string&gt;
          , savePath: string
          , fileNameOptions : 
             Pick&lt;NamingOption, "prefix"&gt; |
             Pick&lt;NamingOption, "name"&gt;
       ) : void {
    
    /*
    Type '"prefix"' does not satisfy the constraint 'keyof NamingOption'.ts(2344)
    Type '"name"' does not satisfy the constraint 'keyof NamingOption'.ts(2344)
    */

I'm still rough with generics, why isn't `NamingOption` being treated as amorphous in shape above? I thought that's the point of a generic type, maybe my concept of them is wrong, or maybe it's a syntax error.

PS: The intent of the above is that argument 3 receive an object with a single key, for either `prefix` or `name` (not both). I know an object union could also do this but I want to learn the generic type way.
## [11][Object argument: Any way to enforce one of two keys must be present?](https://www.reddit.com/r/typescript/comments/hpg1mr/object_argument_any_way_to_enforce_one_of_two/)
- url: https://www.reddit.com/r/typescript/comments/hpg1mr/object_argument_any_way_to_enforce_one_of_two/
---
       protected combineAndSave(
          audiosAndPauseFiles: Array&lt;string&gt;
          , savePath: string
          , fileNameOptions : {
             prefix?: string
             , name?: string
          }
       ) : void {
          const {prefix, name} = fileNameOptions;
          let finalFileSavePath: string;
          
          // used to save sentence files
          if (prefix) {
             finalFileSavePath = `${savePath}/${prefix} - ${this.sentence.folderName}.ogg`;
          } else {
             finalFileSavePath = `${savePath}/${name}`;
          }

I think there's a hole above. Looks like someone may be able to pass an empty object, or even exclude the 3rd argument completely, which would hit the `else` clause with an `undefined` variable and cause a `TypeError` (if this is wrong please let me know).

Is there any way to constrain the `fileNameOptions` object so that one of the keys must be specified?

I tried this

          , fileNameOptions : Partial&lt;{
             prefix?: string
             , name?: string
          }&gt;

But on review of the docs this apparently does nothing the dual `?` don't already do.

I may leave it anyways for code clarity, if anyone has an opinion on that please share.
