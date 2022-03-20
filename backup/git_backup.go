package backup

import (
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func Db_backup() {

	gitRep, err := git.PlainOpen("")

	if err != nil {
		//todo
		return
	}

	tree, err := gitRep.Worktree()
	if err != nil {
		//todo
		return
	}

	_, err = tree.Add(".")
	if err != nil {
		//todo
		return
	}

	treeCommit, error := tree.Commit("The first commit ", &git.CommitOptions{
		All: true,
		Author: &object.Signature{
			Name:  "weixin",
			Email: "kingson4wu@gmail.com",
			When:  time.Now(),
		},
	})

	if error != nil {
		return
	}

	_, error = gitRep.CommitObject(treeCommit)
	if error != nil {
		return
	}

	error = gitRep.Push(&git.PushOptions{})

	if error != nil {
		return
	}

}

//https://pkg.go.dev/gopkg.in/src-d/go-git.v4#Open
