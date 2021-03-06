package main

import (
	. "github.com/OUCC/syaro/logger"
	"github.com/OUCC/syaro/setting"
	"github.com/OUCC/syaro/util"
	"github.com/OUCC/syaro/wikiio"

	"code.google.com/p/go.net/html"

	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	RE_BRACKET        = "^\\[[^\\]]+\\]$"
	RE_DOUBLE_BRACKET = "\\[\\[[^\\]]+\\]\\]"
)

var (
	reBracket       = regexp.MustCompile(RE_BRACKET)
	reDoubleBracket = regexp.MustCompile(RE_DOUBLE_BRACKET)
)

func processWikiLink(n *html.Node, currentDir string) {
	if n.Type != html.TextNode {
		return
	}

	p := n.Parent
	nx := n.NextSibling

	for {
		s := n.Data
		indices := reDoubleBracket.FindStringIndex(s)

		if len(indices) != 0 { // double bracket fount
			name := s[indices[0]+2 : indices[1]-2] // [[name]]
			Log.Debug("bracket tag found: [[%s]]", name)

			// text before <a> tag
			n.Data = s[:indices[0]]

			// <a> tag
			a := html.Node{
				Type: html.ElementNode,
				Data: "a",
			}

			if files := searchPage(name, currentDir); len(files) != 0 { // page found
				// TODO avoid ambiguous page
				Log.Debug("%d pages found", len(files))
				Log.Debug("select %s", files[0].WikiPath())
				a.Attr = []html.Attribute{html.Attribute{
					Key: "href",
					Val: string(files[0].URLPath()),
				},
				}
			} else { // page not found
				Log.Debug("no page found")
				a.Attr = []html.Attribute{
					html.Attribute{
						Key: "class",
						Val: "notfound",
					},
					html.Attribute{
						Key: "href",
						Val: setting.UrlPrefix + "/error/404?data=" + url.QueryEscape(name),
					},
				}
			}

			// text in <a> tag
			a.AppendChild(&html.Node{
				Type: html.TextNode,
				Data: name,
			})

			p.InsertBefore(&a, nx)

			// text after <a> tag
			p.InsertBefore(&html.Node{
				Type: html.TextNode,
				Data: s[indices[1]:],
			}, nx)

			if nx != nil {
				n = nx.PrevSibling
			}
		} else {
			break
		}
	}
}

func processWikiLink2(n *html.Node, currentDir string) {
	if n.Type != html.ElementNode || n.Data != "a" {
		return
	}

	c := n.FirstChild
	if c == nil || c.Type != html.TextNode || !reBracket.MatchString(c.Data) {
		return
	}

	name := c.Data[1 : len(c.Data)-1]
	link := ""
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link = attr.Val
		}
	}
	if link == "" {
		return
	}

	Log.Debug("bracket tag found: [[%s]](%s)", name, link)

	c.Data = name

	if files := searchPage(link, currentDir); len(files) != 0 { //page found
		// TODO avoid ambiguous page
		Log.Debug("%d pages found", len(files))
		Log.Debug("select %s", files[0].WikiPath())

		n.Attr = []html.Attribute{html.Attribute{
			Key: "href",
			Val: string(files[0].URLPath()),
		}}
	} else { // page not found
		Log.Debug("no page found")
		n.Attr = []html.Attribute{
			html.Attribute{
				Key: "href",
				Val: setting.UrlPrefix + "/error/404?data=" + url.QueryEscape(name),
			},
			html.Attribute{
				Key: "class",
				Val: "notfound",
			},
		}
	}
}

func searchPage(name string, currentDir string) []*wikiio.WikiFile {
	if name == "" {
		return nil
	}

	// unescape for searching
	name = html.UnescapeString(name)

	if filepath.IsAbs(name) {
		// search name as absolute path
		// example: /piyo /poyo/pyon.ext
		return searchPageByAbsPath(name)
	} else if strings.Contains(name, "/") || util.IsMarkdown(name) {
		// search name as relative path
		// example: ./hoge ../fuga.ext puyo.ext
		return searchPageByRelPath(name, currentDir)
	} else {
		// search name as base name
		// example: abc
		return searchPageByBaseName(name)
	}
}

func searchPageByAbsPath(abspath string) []*wikiio.WikiFile {
	Log.Debug("main.searchPageByAbsPath(%s)", abspath)
	file, _ := wikiio.Load(abspath)
	if file == nil {
		return nil
	}
	return []*wikiio.WikiFile{file}
}

func searchPageByRelPath(relpath, currentDir string) []*wikiio.WikiFile {
	Log.Debug("main.searchPageByRelPath(%s, %s)", relpath, currentDir)
	wpath := filepath.Join(currentDir, relpath)
	file, _ := wikiio.Load(wpath)
	if file == nil {
		return nil
	}
	return []*wikiio.WikiFile{file}
}

func searchPageByBaseName(baseName string) []*wikiio.WikiFile {
	Log.Debug("main.searchPageByBaseName(%s)", baseName)
	files, _ := wikiio.Search(baseName)
	return files
}
