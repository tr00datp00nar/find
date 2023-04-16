package web

import (
	"github.com/tr00datp00nar/fn"
)

func searchBrave(query string) {
	q := "https://search.brave.com/search?q=" + query
	fn.OpenURL(q)
}

func searchGithub(query string) {
	q := "https://github.com/search?q=" + query
	fn.OpenURL(q)
}

func searchWikipedia(query string) {
	q := "https://en.wikipedia.org/wiki/" + query
	fn.OpenURL(q)
}

func searchStackOverflow(query string) {
	q := "https://stackoverflow.com/search?q=" + query
	fn.OpenURL(q)
}

func searchYoutube(query string) {
	q := "https://www.youtube.com/results?search_query=" + query
	fn.OpenURL(q)
}

func searchGoogle(query string) {
	q := "https://www.google.com/search?q=" + query
	fn.OpenURL(q)
}

func searchAmazon(query string) {
	q := "https://www.amazon.com/s?k=" + query
	fn.OpenURL(q)
}

func searchWolfram(query string) {
	q := "https://www.wolframalpha.com/input?i=" + query
	fn.OpenURL(q)
}

func searchMyAnimeList(query string) {
	q := "https://myanimelist.net/search/all?q=" + query
	fn.OpenURL(q)
}
