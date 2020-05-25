# golang
## [1][Difference between "type a string" and "type a = string"](https://www.reddit.com/r/golang/comments/gpxab2/difference_between_type_a_string_and_type_a_string/)
- url: https://www.reddit.com/r/golang/comments/gpxab2/difference_between_type_a_string_and_type_a_string/
---
I was asked this in an interview, with not much context around it. I am unable to find relevant documentation around this. Please redirect me to the correct resources.
## [2][Instrumentation in Go](https://www.reddit.com/r/golang/comments/gq6aiv/instrumentation_in_go/)
- url: https://gbws.io/articles/instrumentation-in-go/
---

## [3][What is a goroutine? And what is their size?](https://www.reddit.com/r/golang/comments/gqat1d/what_is_a_goroutine_and_what_is_their_size/)
- url: https://tpaschalis.github.io/goroutines-size/
---

## [4][Switch from C++ to Go](https://www.reddit.com/r/golang/comments/gq7rkv/switch_from_c_to_go/)
- url: https://www.reddit.com/r/golang/comments/gq7rkv/switch_from_c_to_go/
---
Hello, I am new to the go lang. I have just played around a little bit. Currently, we have a big desktop application that uses the OpenCV library for image processing. Now we need to create a new cleaner version of this program and start from a new base. Currently, we are using C++. 
I was thinking if it would be a good idea to replace C++ by Go. An important factor for my is if there is an equivalent library for OpenCV (or maybe can I use OpenCV?) and is there a good free/cheap GUI framework for it to create a desktop application.
## [5][What is the purpose of fs and gs registers in golang?](https://www.reddit.com/r/golang/comments/gq4pfh/what_is_the_purpose_of_fs_and_gs_registers_in/)
- url: https://www.reddit.com/r/golang/comments/gq4pfh/what_is_the_purpose_of_fs_and_gs_registers_in/
---
I was looking through the (intel) assembly of a compiled binary written in golang and I found multiple instructions using fs and gs segment registers with address range covering whole 4GB space (was looking on a 32bit system)

I searched around and I found [this stack overflow answer](https://stackoverflow.com/questions/10810203/what-is-the-fs-gs-register-intended-for#10810340) which says that their usage is OS specific. Some common usage I found is to access thread specific storage.

One usage of gs in golang that I could grasp (and not fully understand, please provide insights here) is to get more stack in runtime, by reading this snippet from [golang blog](https://golang.org/doc/asm):

    TEXT main.main(SB) /tmp/x.go
    0x10501c0	MOVQ GS:0x30, CX  # GS used here
    0x10501c9	CMPQ 0x10(CX), SP 
    0x10501cd	JBE 0x1050203 
    ...
    0x1050202	RET 
    0x1050203	CALL runtime.morestack_noctxt(SB)
    0x1050208	JMP main.main(SB)

I wanted to know other usages of gs and fs registers in golang (especially in go runtime). Can someone point me to resources which could include this information?
## [6][Is updating go version via homebrew safe?](https://www.reddit.com/r/golang/comments/gq9bgb/is_updating_go_version_via_homebrew_safe/)
- url: https://www.reddit.com/r/golang/comments/gq9bgb/is_updating_go_version_via_homebrew_safe/
---
I would like to update my go 1.11 to go 1.13 via homebrew in macOS. Did anyone encounter any problems with it? Or is the best way is to uninstall it and reinstall it via homebrew better?
## [7][Error checks extraction, good practice?](https://www.reddit.com/r/golang/comments/gq7nje/error_checks_extraction_good_practice/)
- url: https://www.reddit.com/r/golang/comments/gq7nje/error_checks_extraction_good_practice/
---
Hi there! I've been using Go for some time now, but I have a question of something I do, and I don't know if it actually improves the code readability or hurts it.

I have a project of making a "Cards against Humanity" online game, and I have this function:

    func putBlackCardInPlay(g *cah.GameState) error {
    	if g.BlackCardInPlay != nilBlackCard {
    		return errors.New("Tried to put a black card in play but there is already a black card in play")
    	}
    	if g.Phase == cah.Finished {
    		return errors.New("Tried to put a black card in play but the game has already finished")
    	}
    	if len(g.BlackDeck) == 0 {
    		return errorEmptyBlackDeck{}
    	}
    	g.BlackCardInPlay = g.BlackDeck[0]
    	g.BlackDeck = g.BlackDeck[1:]
    	g.Phase = cah.SinnersPlaying
    	g.CurrRound++
    	return nil
    }

I chose this function because it's kind of simple so it's a good example, but sometimes I have like 5 checks before I actually do something, so then I extract the checks to another function, and end up with something like this:

    func putBlackCardInPlay(g *cah.GameState) error {
    	if err := putBlackCardInPlayChecks(g); err != nil {
    		return err
    	}
    	g.BlackCardInPlay = g.BlackDeck[0]
    	g.BlackDeck = g.BlackDeck[1:]
    	g.Phase = cah.SinnersPlaying
    	g.CurrRound++
    	return nil
    }
    
    func putBlackCardInPlayChecks(g *cah.GameState) error {
    	if g.BlackCardInPlay != nilBlackCard {
    		return errors.New("Tried to put a black card in play but there is already a black card in play")
    	}
    	if g.Phase == cah.Finished {
    		return errors.New("Tried to put a black card in play but the game has already finished")
    	}
    	if len(g.BlackDeck) == 0 {
    		return errorEmptyBlackDeck{}
    	}
    	return nil
    }

I have not seeing something like this in other projects, so I wonder if it would be considered a bad practice.

Thank you in advance!
## [8][Code styles for tests in Go](https://www.reddit.com/r/golang/comments/gq70c1/code_styles_for_tests_in_go/)
- url: https://www.reddit.com/r/golang/comments/gq70c1/code_styles_for_tests_in_go/
---
Hi everyone. The most popular way of writing tests in Go is Table testing. Usually code styles for tests is not a big deal but some aspects of this approach are not very convenient if you use an IDE for development. I wrote an article about how I write tests in Go now which avoids some of the problems of table testing. Would be great if someone can share their opinion on this. Maybe I'm missing something [https://medium.com/@alexander\_yappo/code-styles-for-tests-in-go-b803b4455ffe](https://medium.com/@alexander_yappo/code-styles-for-tests-in-go-b803b4455ffe)
## [9][Why we need to convert a string to byte slice before saving to a file?](https://www.reddit.com/r/golang/comments/gpqysh/why_we_need_to_convert_a_string_to_byte_slice/)
- url: https://www.reddit.com/r/golang/comments/gpqysh/why_we_need_to_convert_a_string_to_byte_slice/
---
Example https://gobyexample.com/writing-files
## [10][How to use encrypted private key for TLS?](https://www.reddit.com/r/golang/comments/gq6pw7/how_to_use_encrypted_private_key_for_tls/)
- url: https://www.reddit.com/r/golang/comments/gq6pw7/how_to_use_encrypted_private_key_for_tls/
---
I generated a certificate and private key with password using OpenSSL. Now I want to use it in Go. I found this [old issue with some instructions](https://golang.org/issue/10181), but it didn't work for me. `x509.DecryptPEMBlock` returns "x509: no DEK-Info header in block". Here is an example key file, password is 1234.

    -----BEGIN ENCRYPTED PRIVATE KEY-----
    MIIBSzBOBgkqhkiG9w0BBQ0wQTApBgkqhkiG9w0BBQwwHAQIQAY0IsXMhucCAggA
    MAwGCCqGSIb3DQIJBQAwFAYIKoZIhvcNAwcECHi4EIJK+T6FBIH4Fatl16Lwznm/
    jIhKygStjhIlpww0A0aZDp/D0eEJpXzvPgRZWf2xhlf5gzTMblQ2XkNrbu/OWOOS
    f+qx//lh30WTFYOwu0ZWBuxGnjDQav2nc+GKRfzCWbTvgdj8EOKi3vgt8PkuBZWp
    IwX0GRrLLd19EmC/VpZ6zAoJIxeE2Oc76tBREJCs5T8o+4Y28rgo/mXbMJmxpdAK
    ncWa4y0f1IEcjdw2u3I8csvtwUIj6WjVLkrS1R3I0DS9jEbs0rZ9uORk5aFatzre
    ccfQA0JI0n15QPX8dGh/RnWmpzpGXMxShiwn434KGD/Fa0mZeQex26chknoV3YE=
    -----END ENCRYPTED PRIVATE KEY-----
