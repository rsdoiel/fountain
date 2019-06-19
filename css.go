// Package fountain is a Golang package supporting Fountain screenplay markup.
//
// css.go manages setting up CSS for either inline style elements or links.
//
package fountain

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func getCSS() string {
	var (
		CSS      string
		override bool
	)
	// 1. Find where we've put any custom CSS
	if _, err := os.Stat("fountain.css"); os.IsNotExist(err) == false {
		CSS = "fountain.css"
		override = true
	} else if _, err := os.Stat(path.Join("css", "fountain.css")); os.IsNotExist(err) == false {
		CSS = path.Join("css", "fountain.css")
		override = true
	}
	if override {
		src, err := ioutil.ReadFile(CSS)
		if err != nil {
			log.Printf("%s", err)
		}
		return fmt.Sprintf("%s", src)
	}
	// 2. Otherwise provide default
	return createElement("style", []string{}, fmt.Sprintf(`
/**
 * fountain.css - CSS for displaying foutain2html generated HTML.
 * It was inspired by scrippet.css found on the Fountain
 * website at https://fountain.io/_css/scrippets.css which is attributed
 * to John August, updated in 2012.
 *
 * 2019-06-18, RSD
 */

.fountain {
	margin: 0;
	padding: 0;
	display: block;
}

.fountain {
	max-width: 400px;
	background: #fffffc;
	color: #000000;
	padding: 5px 14px 15px 14px !important;
	clear: both;
	margin-bottom: 2.5em;
	margin-top: 2.5em;
	border-radius: 3px;
}

section.title-page, section.script {
	width: 36em;
	padding-left: 1em;
	padding-bottom: 2em;
	margin-bottom: 2em;
	border: 1px solid #d2d2d2;
}

.title-page,
.script {
	height: %s;
}

.title {
	position: relative;
	top: 12em;
	text-align: center;
	padding-left: 33%%;
	padding-right: 33%%;
	text-transform: uppercase;
	text-decoration: underline;
	margin-top: 1em;
	margin-bottom: 1em;
}

.author {
	position: relative;
	top: 13em;
	text-align: center;
	padding-left: 33%%;
	padding-right: 33%%;
	margin-top: 0em;
	margin-bottom: 0em;
}

.draft-date, .date {
	position: relative;
	top: 14em;
	text-align: center;
	padding-left: 33%%;
	padding-right: 33%%;
	margin-top: 0;
	margin-bottom: 6em;

}

.copyright {
	position:relative;
	top:35em;
	display: block;
	padding: 0;   	
	margin: 0;
	text-align: left;
	text-transform: none;
	text-decoration: none;
}

.contact {
	position:relative;
	top: 36em;
	display: block;
	padding: 0;
	margin: 0;
	text-align: left;
	text-transform: none;
	text-decoration: none;
}

.script {
	padding-top: 2em;
	padding-left: 0;
	padding-right: 0;
	padding-bottom: 2em;
}

.scene-heading,
.action,
.character,
.parenthetical,
.transition,
.dialogue  {
	font: 12px/14px Courier, "Courier New", monospace;
    text-align: left !important;
	letter-spacing: 0 !important;
	margin-top: 0px !important;
	margin-bottom: 0px !important;
	display: block;
}

.scene-heading,
.action,
.character {
	padding-top: 1.5ex !important;
	display: block;
}

.action {
	padding-right: 5%% !important;
	font-size: 12px !important;
	line-height: 14px !important;
}

.character {
	padding-left: 40%% !important;
}

.dialogue {
	padding-left: 20%% !important;
	padding-right: 20%% !important;
}

.parenthetical {
	display: block;
	padding-left: 32%% !important;
	padding-right: 30%% !important;
}

.dialogue + .parenthetical {
	padding-bottom: 0 !important;
}

.left-align {
	float: left;
	padding-left: 2em;
	text-align: left;
}

.centered {
	padding-left: 33%%;
	padding-right: 33%%;
	text-align: center;
}

.right-align {
	float: right;
	padding-right: 2em;
	text-align: right;
}

.empty {
	display: none;
	height: 0;
}

section.fountain {
	-webkit-box-shadow:
		1px 1px 5px rgba(0,0,0,.1),
		inset -2px -2px 2px white
		;
	-moz-box-shadow:
		1px 1px 5px rgba(0,0,0,.1),
		inset -2px -2px 2px white
		;
	box-shadow:
		1px 1px 5px rgba(0,0,0,.1),
		inset -2px -2px 2px white
		;
	border: none;
	background-image:
	-webkit-gradient(
		linear,
		100%% 100%%,
		50%% 0%%,
		from(#f9f9f9), to(white));
	background-image:
		-moz-linear-gradient(
		    left top,
		    rgb(255,255,255) 29%%,
		    rgb(247,247,247) 100%%
		);
	border: 1px solid #d2d2d2;
}
`, SectionHeight))
}

func getCSSLink() string {
	return fmt.Sprintf("<link rel=%q href=%q>\n", "stylesheet", CSS)
}
