package web

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"

	"github.com/tr00datp00nar/fn"
)

var Cmd = &Z.Cmd{
	Name:        `web`,
	Usage:       `[help]`,
	Version:     `v0.1.0`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_web),
	Description: help.D(_web),
	Commands: []*Z.Cmd{
		help.Cmd,
		amazonCmd,
		braveCmd,
		githubCmd,
		googleCmd,
		malCmd,
		stackOverflowCmd,
		wikipediaCmd,
		wolframCmd,
		youtubeCmd,
	},
}

var amazonCmd = &Z.Cmd{
	Name:        `amazon`,
	Usage:       `QUERY`,
	Version:     `v0.1.0`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_amazon),
	Description: help.D(_amazon),
	Commands: []*Z.Cmd{
		help.Cmd,
	},

	Call: func(_ *Z.Cmd, args ...string) error {
		a := Z.ArgsOrIn(args)
		query := fn.RawUrlEncode(a)
		searchAmazon(query)
		return nil
	},
}

var braveCmd = &Z.Cmd{
	Name:        `brave`,
	Usage:       `QUERY`,
	Version:     `v0.1.0`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_brave),
	Description: help.D(_brave),
	Commands: []*Z.Cmd{
		help.Cmd,
	},

	Call: func(_ *Z.Cmd, args ...string) error {
		a := Z.ArgsOrIn(args)
		query := fn.RawUrlEncode(a)
		searchBrave(query)
		return nil
	},
}

var githubCmd = &Z.Cmd{
	Name:        `github`,
	Aliases:     []string{`gh`},
	Usage:       `QUERY`,
	Version:     `v0.1.0`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_github),
	Description: help.D(_github),
	Commands: []*Z.Cmd{
		help.Cmd,
	},

	Call: func(_ *Z.Cmd, args ...string) error {
		a := Z.ArgsOrIn(args)
		query := fn.RawUrlEncode(a)
		searchGithub(query)
		return nil
	},
}

var googleCmd = &Z.Cmd{
	Name:        `google`,
	Usage:       `QUERY`,
	Version:     `v0.1.0`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_google),
	Description: help.D(_google),
	Commands: []*Z.Cmd{
		help.Cmd,
	},

	Call: func(_ *Z.Cmd, args ...string) error {
		a := Z.ArgsOrIn(args)
		query := fn.RawUrlEncode(a)
		searchGoogle(query)
		return nil
	},
}

var stackOverflowCmd = &Z.Cmd{
	Name:        `stackoverflow`,
	Usage:       `QUERY`,
	Version:     `v0.1.0`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_stackoverflow),
	Description: help.D(_stackoverflow),
	Commands: []*Z.Cmd{
		help.Cmd,
	},

	Call: func(_ *Z.Cmd, args ...string) error {
		a := Z.ArgsOrIn(args)
		query := fn.RawUrlEncode(a)
		searchStackOverflow(query)
		return nil
	},
}

var wikipediaCmd = &Z.Cmd{
	Name:        `wikipedia`,
	Usage:       `QUERY`,
	Aliases:     []string{`wiki`},
	Version:     `v0.1.0`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_wikipedia),
	Description: help.D(_wikipedia),
	Commands: []*Z.Cmd{
		help.Cmd,
	},

	Call: func(_ *Z.Cmd, args ...string) error {
		a := Z.ArgsOrIn(args)
		query := fn.RawUrlEncode(a)
		searchWikipedia(query)
		return nil
	},
}

var wolframCmd = &Z.Cmd{
	Name:        `wolfram`,
	Usage:       `QUERY`,
	Version:     `v0.1.0`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_wolfram),
	Description: help.D(_wolfram),
	Commands: []*Z.Cmd{
		help.Cmd,
	},

	Call: func(_ *Z.Cmd, args ...string) error {
		a := Z.ArgsOrIn(args)
		query := fn.RawUrlEncode(a)
		searchWolfram(query)
		return nil
	},
}

var youtubeCmd = &Z.Cmd{
	Name:        `youtube`,
	Usage:       `QUERY`,
	Aliases:     []string{`yt`},
	Version:     `v0.1.0`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_youtube),
	Description: help.D(_youtube),
	Commands: []*Z.Cmd{
		help.Cmd,
	},

	Call: func(_ *Z.Cmd, args ...string) error {
		a := Z.ArgsOrIn(args)
		query := fn.RawUrlEncode(a)
		searchYoutube(query)
		return nil
	},
}

var malCmd = &Z.Cmd{
	Name:      `mal`,
	Usage:     `[help]`,
	Version:   `v0.0.1`,
	Copyright: `Copyright Micah Nadler 2023`,
	License:   `Apache-2.0`,
	// Summary:     help.S(_mal),
	// Description:  help.D(_mal),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		a := Z.ArgsOrIn(args)
		query := fn.RawUrlEncode(a)
		searchMyAnimeList(query)
		return nil
	},
}
