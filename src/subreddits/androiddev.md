# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/eunf8c/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/eunf8c/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - January 27, 2020](https://www.reddit.com/r/androiddev/comments/eumgji/weekly_questions_thread_january_27_2020/)
- url: https://www.reddit.com/r/androiddev/comments/eumgji/weekly_questions_thread_january_27_2020/
---
This thread is for simple questions that don't warrant their own thread (although we suggest checking the sidebar, [the wiki](http://www.reddit.com/r/androiddev/wiki/), [our Discord](https://discord.gg/D2cNrqX), or [Stack Overflow](http://stackoverflow.com) before posting). Examples of questions:

* How do I pass data between my Activities?
* Does anyone have a link to the source for the AOSP messaging app?
* Is it possible to programmatically change the color of the status bar without targeting API 21?

**Important: Downvotes are strongly discouraged in this thread. Sorting by new is strongly encouraged.**

Large code snippets don't read well on reddit and take up a lot of space, so please don't paste them in your comments. Consider linking [Gists](https://gist.github.com) instead.

Have a question about the subreddit or otherwise for /r/androiddev mods? [We welcome your mod mail!](http://www.reddit.com/message/compose?to=%2Fr%2Fandroiddev)

Also, please don't link to Play Store pages or ask for feedback on this thread. Save those for the App Feedback threads we host on Saturdays.

Looking for all the Questions threads? Want an easy way to locate this week's thread? Click [this link](https://www.reddit.com/r/androiddev/search?q=title%3A%22questions+thread%22+author%3A%22AutoModerator%22&amp;restrict_sr=on&amp;sort=new&amp;t=all)!
## [3][Dynamic screens using server-driven UI in Android](https://www.reddit.com/r/androiddev/comments/eufs6x/dynamic_screens_using_serverdriven_ui_in_android/)
- url: https://proandroiddev.com/dynamic-screens-using-server-driven-ui-in-android-262f1e7875c1
---

## [4][Auto read OTP android with SMS User Consent API](https://www.reddit.com/r/androiddev/comments/eu9wy6/auto_read_otp_android_with_sms_user_consent_api/)
- url: https://www.reddit.com/r/androiddev/comments/eu9wy6/auto_read_otp_android_with_sms_user_consent_api/
---
I have created a sample app that verify OTP using SMS User Consent API. 

  

https://preview.redd.it/epnw0o85u5d41.png?width=770&amp;format=png&amp;auto=webp&amp;s=7353709e992d8ceb12be677c3981a6951c880896

[https://androidwave.com/auto-read-otp-android-user-consent-api/](https://androidwave.com/auto-read-otp-android-user-consent-api/)

\#android #androiddev #androidwave
## [5][How come some annoying recruiters bypass Do Not Disturb and even the zeroed media volume and still harass me with ringing calls? Is this a privilege they pay for?](https://www.reddit.com/r/androiddev/comments/eunskg/how_come_some_annoying_recruiters_bypass_do_not/)
- url: https://www.reddit.com/r/androiddev/comments/eunskg/how_come_some_annoying_recruiters_bypass_do_not/
---

## [6][Android WebView not working with CipherInputStream](https://www.reddit.com/r/androiddev/comments/eunsia/android_webview_not_working_with_cipherinputstream/)
- url: https://www.reddit.com/r/androiddev/comments/eunsia/android_webview_not_working_with_cipherinputstream/
---
I am building an Application which will be showing downloaded Encrypted web-pages (web pages contains might contain webm files also).

To achieve this goal I override **shouldInpterceptRequest** method which returns **WebResourceResponse**. The constructor of **WebResourceResponse** needs multiple argument one of which is **InputStream**. 

Every thing is working fine When I am using Decrypted files and getting **FileInputStream** and passing the same to **WebResourceResponse**

    private WebResourceResponse getFile(String filePath) {
        FileInputStream fileInputStream = null;
        String mimeType = "";
        String encoding = null;
        switch (filePath.substring(filePath.lastIndexOf('.'))) {
            case ".html" : mimeType =  "text/html";
                            encoding = "utf-8";
            break;

            case ".png" : mimeType = "image/png";
            break;

            case ".webm" : mimeType = "video/webm";
            break;

            case ".css" : mimeType = "text/css";
                encoding = "utf-8";
                break;

            case ".js" : mimeType = "text/javascript";
                encoding = "utf-8";
                break;
        }
       
           fileInputStream = new FileInputStream(new File(path, filePath));
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }
        return new WebResourceResponse(mimeType, encoding, fileInputStream);
    }

But When I am passing **CipherInputStream** (While using encrypted file) its not working. 

        private WebResourceResponse getFile(String filePath) {
        String mimeType = "";
        String encoding = null;
        switch (filePath.substring(filePath.lastIndexOf('.'))) {
            case ".html" : mimeType =  "text/html";
                            encoding = "utf-8";
            break;

            case ".png" : mimeType = "image/png";
            break;

            case ".webm" : mimeType = "video/webm";
            break;

            case ".css" : mimeType = "text/css";
                encoding = "utf-8";
                break;

            case ".js" : mimeType = "text/javascript";
                encoding = "utf-8";
                break;
        }
        TestFileInputStream TestFileInputStream = null;
        TestCipherInputStream testCipherInputStream = null;
        try {
            if(!filePath.contains("favicon.ico") ) {

                testFileInputStream = new TestFileInputStream(new File(path, filePath));
                testCipherInputStream = new TestCipherInputStream(testFileInputStream, cipher);
                fileInputStreamArrayList.add(testFileInputStream);
                cipherInputStreamArrayList.add(testCipherInputStream);
            }
            //fileInputStream = new FileInputStream(new File(path, filePath));
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }
        return new WebResourceResponse(mimeType, encoding, testCipherInputStream);
    }
// Custom Class so that CipherInputStream will not Closed 

    public class TestCipherInputStream extends CipherInputStream {
        private boolean canClose = false;
        public TestCipherInputStream(InputStream is, Cipher c) {
            super(is, c);
        }

        @Override
        public void close() throws IOException {
            if(canClose)
                super.close();
        }

        public void canCloseStream() {
            canClose = true;
        }
    }
// Custom Class so that fileInputStream will not Closed 

    public class TestFileInputStream extends FileInputStream {
        private boolean canClose = false;

        public TestFileInputStream(File file) throws FileNotFoundException {
            super(file);
        }


        @Override
        public void close() throws IOException {
            if(canClose)
                super.close();
        }

        public void canCloseStream() {
            canClose = true;
        }
    }

Its shows syntax error in js, html etc file. I thought it might be an error because of synchronisation (When using cipherInputStream that time before getting the whole js file WebView started calling other files) So I put some delay and its working. 

But then I faced **one more issue** (My web-Page have one scroll feature which when user scroll it fetch different quality of video) When user scroll **shouldInpterceptRequest** was called and I am passing the right video but its not displaying.

When I tried to debug that I saw that WebView not calling **available()** and **read()** method over that stream.


Thank You.
## [7][How-To In-App Subscriptions for an app?](https://www.reddit.com/r/androiddev/comments/eun8vf/howto_inapp_subscriptions_for_an_app/)
- url: https://www.reddit.com/r/androiddev/comments/eun8vf/howto_inapp_subscriptions_for_an_app/
---
Hi!

I have a free app with basic functionality, but I would like to add many more. In order to do so, I've planned to implement in-app purchase (subscription) so people wanting premium features and to support my development can do so as it'd take plenty time and research from my part.

But I don't know where to start, I've already looked StackOverflow and other sites. My questions:

* Is it a good idea to do so? (I'm approaching it right?)
* How to handle in-app subscriptions? (Firebase?)
* Should I make it so code (modules?) will be downloaded into user's device when the subscription is brought?
* App's functionality, for now, doesn't require a server but should I make an auth system (sign-in with Google account for example) to verify premium users?
## [8][Android Biometric Authentication with Face Auth](https://www.reddit.com/r/androiddev/comments/eumwx0/android_biometric_authentication_with_face_auth/)
- url: https://www.reddit.com/r/androiddev/comments/eumwx0/android_biometric_authentication_with_face_auth/
---
Hi all, 

I recently noticed some banking apps adding in their change logs that the app can now be unlocked using your face through the [biometric prompt API](https://developer.android.com/training/sign-in/biometric-auth). This is a feature I am interested in testing and getting working in some past apps I have worked on, however, I cannot seem to find a list of phones that are supported. I currently use a OnePlus 6T which does support system face unlock but doesn't use your face for the biometric prompt. I believe it's something to do with google deeming their technology not secure enough or whatever. 

Does anyone know of a phone or a list of phones that support this new API feature as I will be looking to purchase one to test. I imagine it would be the latest pixel phones but I don't want to buy one to find out it doesn't work on it!

Thanks in advance
## [9][How to make drawers in Jetpack Compose](https://www.reddit.com/r/androiddev/comments/eumpl8/how_to_make_drawers_in_jetpack_compose/)
- url: https://android.jlelse.eu/exploring-drawers-in-jetpack-compose-3131e6f1b07b
---

## [10][Call Mainactivity's method](https://www.reddit.com/r/androiddev/comments/eumncs/call_mainactivitys_method/)
- url: https://www.reddit.com/r/androiddev/comments/eumncs/call_mainactivitys_method/
---
Greetings,

I'm trying to call a method for my mainactivity, to no avail. On the root of my project, I have 2 files, Bth.cs and MainActivity.cs

For context, my knowledge in Xamarin is... Lackluster. I'm self taught, as the project was basically thrown on my lap. I learned C#, but Xamarin is brand new to me, I only got a few days of tutorials on my back. The project was created a few years ago and there is absolutely NO documentation and there are barely any comments (most of them don't help in any ways). Now onto the problem:

&amp;#x200B;

In Mainactivity, I have a timer that runs in the background to check for user activity. If nothing happen within 5 minutes, shut down the app (necessary, as we are looking to facture the time spent, so we need relatively accurate readings)

Bth is the code partaining to the barcode readers we use.

I've managed to pinpoint The part of the code that run whenever I scan a barcode. I would like to be able to call the Method in MainActivity that reset the variables used by the timer. But I cannot find a way to reach MainActivity. If that helps, know that Bth task that runs the connection to the barcode reader is an async method. Bth itself isn't an activity.

Here are the relevant signatures:

    //MainActivity, which is an activity
    public class MainActivity : global::Xamarin.Forms.Platform.Android.FormsAppCompatActivity{}
    
    //The method that resets the timer
    public void ResetTimer(){}
    
    //Bth class, IBth being an interface
    public class Bth : IBth
    
    //loop being the async task run all the time in the background.
    private async Task loop(string name, int sleepTime, bool readAsCharArray)

Thank you for your help.
## [11][How can I display a large online pdf file?](https://www.reddit.com/r/androiddev/comments/eum97x/how_can_i_display_a_large_online_pdf_file/)
- url: https://www.reddit.com/r/androiddev/comments/eum97x/how_can_i_display_a_large_online_pdf_file/
---
I am trying to save a lot of pdf files on a firebase database, and display them at will using webview. My problem is that some of them are large(100 pages or so), and when using google drive to display the file it says it is too large. Downloading the file is also no good since it could take a long time to download, and id rather it displays it page by page like google docs does. Any idea how I can do so?
## [12][Offline-first Android app with local database and online sync](https://www.reddit.com/r/androiddev/comments/eue3c3/offlinefirst_android_app_with_local_database_and/)
- url: https://www.reddit.com/r/androiddev/comments/eue3c3/offlinefirst_android_app_with_local_database_and/
---
I'm developing an android app that needs to function well offline, but I also want to be able to backup and sync the database with the cloud for paid users. Almost all data in the app will be user-created. If I didn't care about backup and sync, then Room database would be all I needed. But I'm not sure how to arrange backup and sync across devices for Room/SQLite.

I've been learning about Firestore as an alternative, but I have two concerns:

1. Reading and writing the same data that is already stored locally (cache/offline persistence turned on) causing unnecessarily high charges. The user is expected to access the app multiple times a day for very brief periods. I've read that after 30 minutes of inactivity, all the data will be re-fetched, which would be a LOT of read charges!
2. The speed of queries when offline. Apparently reading data from the cache only is slow.

Has anyone used Firestore for an offline-first app? Is there a way to use both Room for free users who only want data on their device and Firestore for paid users who want backup and sync? What would be an appropriate alternative? I'm pretty much a newbie at android dev. I have one app released, but it's all offline (Room database), so this is my first time looking at online solutions.

TIA.
