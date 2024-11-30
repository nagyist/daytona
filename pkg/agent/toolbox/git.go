// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package toolbox

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type CloneRequest struct {
	URL      string `json:"url" binding:"required"`
	Path     string `json:"path" binding:"required"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CommitRequest struct {
	Path    string `json:"path" binding:"required"`
	Message string `json:"message" binding:"required"`
	Author  string `json:"author" binding:"required"`
	Email   string `json:"email" binding:"required"`
}

type PushRequest struct {
	Path     string `json:"path" binding:"required"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type BranchRequest struct {
	Path string `json:"path" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type CommitInfo struct {
	Hash      string    `json:"hash"`
	Author    string    `json:"author"`
	Email     string    `json:"email"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func cloneRepository(c *gin.Context) {
	var req CloneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	options := &git.CloneOptions{
		URL:      req.URL,
		Progress: nil,
	}

	if req.Username != "" && req.Password != "" {
		options.Auth = &http.BasicAuth{
			Username: req.Username,
			Password: req.Password,
		}
	}

	_, err := git.PlainClone(req.Path, false, options)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Repository cloned successfully"})
}

func getStatus(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		c.JSON(400, gin.H{"error": "path is required"})
		return
	}

	repo, err := git.PlainOpen(path)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	worktree, err := repo.Worktree()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	status, err := worktree.Status()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, status)
}

func commitChanges(c *gin.Context) {
	var req CommitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo, err := git.PlainOpen(req.Path)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	worktree, err := repo.Worktree()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_, err = worktree.Add(".")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	commit, err := worktree.Commit(req.Message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  req.Author,
			Email: req.Email,
			When:  time.Now(),
		},
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"hash": commit.String()})
}

func pushChanges(c *gin.Context) {
	var req PushRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo, err := git.PlainOpen(req.Path)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	options := &git.PushOptions{}
	if req.Username != "" && req.Password != "" {
		options.Auth = &http.BasicAuth{
			Username: req.Username,
			Password: req.Password,
		}
	}

	err = repo.Push(options)
	if err != nil && err != git.NoErrAlreadyUpToDate {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Changes pushed successfully"})
}

func listBranches(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		c.JSON(400, gin.H{"error": "path is required"})
		return
	}

	repo, err := git.PlainOpen(path)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	branches, err := repo.Branches()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var branchList []string
	branches.ForEach(func(ref *plumbing.Reference) error {
		branchList = append(branchList, ref.Name().Short())
		return nil
	})

	c.JSON(200, branchList)
}

func createBranch(c *gin.Context) {
	var req BranchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repo, err := git.PlainOpen(req.Path)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	worktree, err := repo.Worktree()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = worktree.Checkout(&git.CheckoutOptions{
		Create: true,
		Branch: plumbing.NewBranchReferenceName(req.Name),
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Branch created successfully"})
}

func getCommitHistory(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		c.JSON(400, gin.H{"error": "path is required"})
		return
	}

	repo, err := git.PlainOpen(path)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ref, err := repo.Head()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	commits, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var history []CommitInfo
	commits.ForEach(func(commit *object.Commit) error {
		history = append(history, CommitInfo{
			Hash:      commit.Hash.String(),
			Author:    commit.Author.Name,
			Email:     commit.Author.Email,
			Message:   commit.Message,
			Timestamp: commit.Author.When,
		})
		return nil
	})

	c.JSON(200, history)
}
