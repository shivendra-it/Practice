package main

import fb "github.com/huandu/facebook"

func setupFBVersion() {
	fb.Version = "v2.6"

	// it's possible to specify version per session.
	session := &fb.Session{}
	session.Version = "V2.6"
}
