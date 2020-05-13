# golang
## [1][What's coming in Go 1.15](https://www.reddit.com/r/golang/comments/ginp22/whats_coming_in_go_115/)
- url: https://lwn.net/SubscriberLink/820217/47ed80088c03b18d/
---

## [2][Too modern Go application? Building a serverless application with Google Cloud Run and Firebase](https://www.reddit.com/r/golang/comments/giwx9q/too_modern_go_application_building_a_serverless/)
- url: https://threedots.tech/post/serverless-cloud-run-firebase-modern-go-application/
---

## [3][🍄VPN supporting authentications such as Google OpenID Connect or AWS IAM ... ETC, over GRPC.](https://www.reddit.com/r/golang/comments/givybr/vpn_supporting_authentications_such_as_google/)
- url: https://github.com/gjbae1212/grpc-vpn
---

## [4][🔍 Announcing observe: Observe a website and get an e-mail if something changes.](https://www.reddit.com/r/golang/comments/givvf3/announcing_observe_observe_a_website_and_get_an/)
- url: https://github.com/dominikbraun/observe
---

## [5][Tutorial on building a desktop app with goey](https://www.reddit.com/r/golang/comments/giy708/tutorial_on_building_a_desktop_app_with_goey/)
- url: https://bitbucket.org/rj/goey/wiki/Tutorial-Hello
---

## [6][Ensmallening Go binaries by prohibiting comparisons](https://www.reddit.com/r/golang/comments/gifbju/ensmallening_go_binaries_by_prohibiting/)
- url: https://dave.cheney.net/2020/05/09/ensmallening-go-binaries-by-prohibiting-comparisons
---

## [7][🚀lazyhub - Terminal UI Client for GitHub using gocui.](https://www.reddit.com/r/golang/comments/gia5zz/lazyhub_terminal_ui_client_for_github_using_gocui/)
- url: https://www.reddit.com/r/golang/comments/gia5zz/lazyhub_terminal_ui_client_for_github_using_gocui/
---
&amp;#x200B;

https://i.redd.it/4gtiwcxqtby41.gif

[https://github.com/ryo-ma/lazyhub](https://github.com/ryo-ma/lazyhub)

* 🚀Check the trending repositories on GitHub today
* 🔍Search repositories
* 📘Read the README
* 📄Copy the clone command to clipboard
* 💻Open the repository page on your browser
## [8][PickleIt - A Gui based version control system for binary files](https://www.reddit.com/r/golang/comments/gii6wc/pickleit_a_gui_based_version_control_system_for/)
- url: https://www.reddit.com/r/golang/comments/gii6wc/pickleit_a_gui_based_version_control_system_for/
---
I once saw this meme, and wondered if I could build a solution for it. I settled on Go and [Qt](https://github.com/therecipe/qt) for the GUI and built pickleIt. Its early stages, and currently I have only compiled for OSX (as thats all I have). Please see the readme in the repository, and if you want to try it on OSX (or tell me that it doesn't work....) there is a download link near the top!  
To be clear, this was a side project, and at this stage, I'm wondering what people think.  
[Link to Code/Download Link here](https://gitlab.com/amlwwalker/pickleit)   
Feedback massively appreciated

&amp;#x200B;

[PickleIt Gui](https://preview.redd.it/2wt6mkjyvdy41.png?width=2624&amp;format=png&amp;auto=webp&amp;s=805225ee7dce35ca2297a06a4eacd65e167f2597)

[The meme that got me going](https://preview.redd.it/vrr10nz8vdy41.png?width=766&amp;format=png&amp;auto=webp&amp;s=b7be1496f7de84a39b8aaa50282352652bcc7ab3)
## [9][Is Go's concurrency model which uses channels and goroutines for writing concurrent code an example of function coloring?](https://www.reddit.com/r/golang/comments/gixpgu/is_gos_concurrency_model_which_uses_channels_and/)
- url: https://www.reddit.com/r/golang/comments/gixpgu/is_gos_concurrency_model_which_uses_channels_and/
---
I was watching this talk https://www.youtube.com/watch?v=zeLToGnjIUM which discusses various ways to write async code in different languages, the speaker discussed Go and said since you have to use channels to get return values from goroutines this is like function coloring http://journal.stuffwithstuff.com/2015/02/01/what-color-is-your-function/

now I am not exactly sure what function coloring is I am assuming that it is a syntax in which the langauge starts looking more like a DSL and becomes hard to reason about.

While programming in Go I have found that it is a joy to write concurrent code and its very easy to reason about as compared to any of the other async await syntax which doesn't really make sense (although the Zig async await syntax does look much cleaner and straightforward)

I wanted to the community's opinion on this whether goroutines-channels style concurrency can be considered an example of function colouring?
## [10][How should I test API wrapper?](https://www.reddit.com/r/golang/comments/gixc6c/how_should_i_test_api_wrapper/)
- url: https://www.reddit.com/r/golang/comments/gixc6c/how_should_i_test_api_wrapper/
---
Hi I'm currently working on API wrapper and decided that I'll learn how to test Golang code. I have such method in my struct:

    func (l *Librus) CreateSession() error {
    	postData := url.Values{}
    	postData.Set("username", l.Username)
    	postData.Set("password", l.Password)
    	postData.Set("librus_long_term_token", "1")
    	postData.Set("grant_type", "password")
    
    	// request
    	req, err := http.NewRequest("POST", host+"OAuth/Token", strings.NewReader(postData.Encode()))
    	// add headers
    	for _, h := range Headers {
    		req.Header.Set(h.Key, h.Value)
    	}
    
    	if err != nil {
    		return err
    	}
    
    	// response
    	res, err := l.Client.Do(req)
    	if err != nil {
    		return err
    	}
    	defer res.Body.Close()
    
    	// check response code
    	if res.StatusCode != http.StatusOK {
    		return fmt.Errorf("Error status code, wanted: %v, got: %v", http.StatusOK, res.StatusCode)
    	}
    
    	// decode json response
    	okResponse := new(OKResponse)
    	err = json.NewDecoder(res.Body).Decode(okResponse)
    	if err != nil {
    		return err
    	}
    
    	// change authorization header
    	Headers[0].Value = "Bearer " + okResponse.AccessToken
    
    	return nil
    }

I wrote mock for http.Client and it looks like this:

    type MockClient struct {
    	ExpectedRes *http.Response
    	Req         *http.Request
    	Err         error
    	DoFunc      func(req *http.Request) (*http.Response, error)
    }
    
    func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
    	m.Req = req
    	if m.DoFunc != nil {
    		return m.DoFunc(req)
    	}
    	return m.ExpectedRes, m.Err
    }

Test:

    func TestSuccessCreateSession(t *testing.T) {
    	client := &amp;mocks.MockClient{
    		DoFunc: func(req *http.Request) (*http.Response, error) {
    			return &amp;http.Response{
    				StatusCode: http.StatusOK,
    			}, nil
    		},
    	}
    
    	l := &amp;golibrus.Librus{
    		Client:   client,
    		Username: "test",
    		Password: "test",
    	}
    
    	err = l.CreateSession()
    	assert.Nil(t, err)
    }

&amp;#x200B;

I tried writing test for this 'CreateSession' method but it returns error on 'defer.res.Body.Close()', I think it causes it because 'Body' is empty. Do you guys have any advice? I read a lot of articles how to mock but I still don't understand it :(.
