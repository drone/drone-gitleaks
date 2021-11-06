// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

type match struct {
	Line            string `json:"line"`
	Linenumber      int    `json:"lineNumber"`
	Offender        string `json:"offender"`
	Offenderentropy int    `json:"offenderEntropy"`
	Commit          string `json:"commit"`
	Repo            string `json:"repo"`
	Repourl         string `json:"repoURL"`
	Leakurl         string `json:"leakURL"`
	Rule            string `json:"rule"`
	Commitmessage   string `json:"commitMessage"`
	Author          string `json:"author"`
	Email           string `json:"email"`
	File            string `json:"file"`
	Date            string `json:"date"`
	Tags            string `json:"tags"`
}
